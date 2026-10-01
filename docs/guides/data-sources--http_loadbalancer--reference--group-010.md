---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-11b87c9778aeb5f42c8df12861b8b09569b10a76ec16c656e1ca702ed88e6f3b"></a>

## api_specification.validation_disabled — api_specification.validation_disabled / 265d1366ee3c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- api_specification.validation_disabled

<a id="canonical-a7af21c7532a01897081e12a705cedce3dfec2a5d29240f8dff680bcd96921f6"></a>

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

<a id="canonical-9bf2afe6c595e6f51c93a38afbb708b852d2b15c2dd7779346975b81fc964968"></a>

## Direct properties — api_specification.validation_disabled / 265d1366ee3c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5b35bedff0ab92ae754df48ac129165744b3d54a3cd77266fc132651275804e5"></a>

## Next pages — api_specification.validation_disabled / 265d1366ee3c / 4

- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb1a6c085ab33dd3d974ef30264088cab646dc5353bd83c0a050edd9f063812e"></a>

## api_testing — api_testing / cd7691eea668 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- api_testing

<a id="canonical-349aeaac87a8c8a832bffeb65ed124a7ea5358bbc965253c7d08642b2877ab02"></a>

Type: `"single"`. Computed.

\[OneOf: api\_testing, disable\_api\_testing; Default: disable\_api\_testing\] API Testing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-frequency_choice": "[\"every_day\",\"every_month\",\"every_week\"]"
}
```

OneOf alternatives in this subsection:

- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-349aeaac87a8c8a832bffeb65ed124a7ea5358bbc965253c7d08642b2877ab02)
- [disable_api_testing](data-sources--http_loadbalancer--reference--group-017.md#canonical-bca8ea2132d6126063f053b84a89cc22d57c2e7fdb82651d031def1788f0961f)

Select alternatives according to the provider validators above.

<a id="canonical-0faa9e39917e65febb365ca0fdc7c66d008c38c5c6653670664e2df3c154aa2b"></a>

## Direct properties — api_testing / cd7691eea668 / 3

<a id="canonical-26eab44995b2e6622a3a7a6d1c24584a94e39da3e29a958dce63334e567336b6"></a>

<a id="canonical-bc8b7ca9ae0b09b0e91f005eac6dd7bfea0302cfc7500c776c73af0fa10a33e0"></a>

## custom_header_value property — api_testing / cd7691eea668 / 4

Type: `"string"`. Computed.

Add x-F5-API-testing-identifier header value to prevent security flags on API testing traffic.

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

- [domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb): complete subsection reference.

- [every_day](data-sources--http_loadbalancer--reference--group-010.md#canonical-d16625fa6562db275fe1d05b341d9879205d06f27dd3b9eb4d65e09fdb5acd00): complete subsection reference.

- [every_month](data-sources--http_loadbalancer--reference--group-010.md#canonical-0a02f2b4f94f3a8749ec70ad4653a1171e6c169a17f723204b4ece102e1b5924): complete subsection reference.

- [every_week](data-sources--http_loadbalancer--reference--group-010.md#canonical-04a403d827689939d030abd23a47f97de7b553937de3f4b12a68a4dfacb969ff): complete subsection reference.

<a id="canonical-cdba0b721be7c1b0f6abeee2c798f6bf1bb580a64c7d1e292d7ac26f845bec3d"></a>

## Next pages — api_testing / cd7691eea668 / 5

- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [api_testing.every_day](data-sources--http_loadbalancer--reference--group-010.md#canonical-d16625fa6562db275fe1d05b341d9879205d06f27dd3b9eb4d65e09fdb5acd00)
- [api_testing.every_month](data-sources--http_loadbalancer--reference--group-010.md#canonical-0a02f2b4f94f3a8749ec70ad4653a1171e6c169a17f723204b4ece102e1b5924)
- [api_testing.every_week](data-sources--http_loadbalancer--reference--group-010.md#canonical-04a403d827689939d030abd23a47f97de7b553937de3f4b12a68a4dfacb969ff)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04fc3292b7da6fd180a7e65fe47b0201e8d4e7a235d355d4440baf87d9d2cd78"></a>

## api_testing.domains — api_testing.domains / ec0ed7301074 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- api_testing.domains

<a id="canonical-be23fcf3ccb7d609dcecc4ec18de685936d916e1b1d9461ca339fa499bb617e3"></a>

Type: `"list"`. Computed.

Add and configure testing domains and credentials.

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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-aecba00091dc3377eae78cc31d85947a6774c68f0d6e3a5aa7368b67c98257e6"></a>

## Direct properties — api_testing.domains / ec0ed7301074 / 3

<a id="canonical-071d206cb9affe78521ea1e06a3ae043684e6bcf9945324fd97f42cb1df6050f"></a>

<a id="canonical-bca7eade9a0aac0fe5e475e5fd5df8c3dbb1d7adf9f6b54da8b114d13f798f06"></a>

## allow_destructive_methods property — api_testing.domains / ec0ed7301074 / 4

Type: `"bool"`. Computed.

Enable to allow API Testing to execute against destructive methods. Use with caution as these may
modify or DELETE data.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1): complete subsection reference.

<a id="canonical-53efd356f355e9d1d2c5f06d9e7d35e25072bf518949e16d702a289c53591011"></a>

<a id="canonical-e714eb6265df904f63502caf2c0c55c0870ea0ab41e0c17903d7e17d29b0fa34"></a>

## domain property — api_testing.domains / ec0ed7301074 / 5

Type: `"string"`. Computed.

Add your testing environment domain. Be aware that running tests on a production domain can impact
live applications, as API testing cannot distinguish between production and testing environments.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-2eec85698a0946ad05ff8c76bcfbf1539bbe9f8e6ed13a63ac7ce991d64d13c3"></a>

## Next pages — api_testing.domains / ec0ed7301074 / 6

- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8fdda99e9d8ce358ef8139044ac8213e69538c98beb03bcaba7ce9930e39ac00"></a>

## api_testing.domains.credentials — api_testing.domains.credentials / f7a3ef1c4f73 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- api_testing.domains.credentials

<a id="canonical-f2e4a000de939491116db6f305d1d98a847236baccc0554d7df0db6ab754df50"></a>

Type: `"list"`. Computed.

Add credentials for API testing to use in the selected environment.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-341da101e6d30cd3b7f5b90029d41a09ff8e849866cc9b5a78a78909d23a869f"></a>

## Direct properties — api_testing.domains.credentials / f7a3ef1c4f73 / 3

- [admin](data-sources--http_loadbalancer--reference--group-010.md#canonical-505012b3df1ac809619e0407e4d5d161a70be114fb33cc8341d94b030d460a87): complete subsection reference.

- [api_key](data-sources--http_loadbalancer--reference--group-010.md#canonical-ef7dfee2b7b52027e49a0e9d727e19417e976a40a116621ae995d9e04b36426f): complete subsection reference.

- [basic_auth](data-sources--http_loadbalancer--reference--group-010.md#canonical-b3c25a97b23dd1588c93610432003c1f863335c6b0b2abdaa9b96e9e1deaef99): complete subsection reference.

- [bearer_token](data-sources--http_loadbalancer--reference--group-010.md#canonical-d89982c7f459e67e9e9dd2c66d215f93d88d9ff3614d34e0cefb68f962525926): complete subsection reference.

<a id="canonical-97caf3079a8597344e1632f090a350461ddfb5229fa77bcc3d3bc9dcefa8f7b5"></a>

<a id="canonical-762e92445c7850538c343aa780e5345de6da38c300124a900943030570f97430"></a>

## credential_name property — api_testing.domains.credentials / f7a3ef1c4f73 / 4

Type: `"string"`. Computed.

Enter a unique name for the credentials used in API testing.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
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

- [login_endpoint](data-sources--http_loadbalancer--reference--group-010.md#canonical-c2e3cc2aa570b84c3978c9e016233fe86dd77038c9f1bf8dbf3675d2b928e760): complete subsection reference.

- [standard](data-sources--http_loadbalancer--reference--group-010.md#canonical-9be1f76ac7a312663fdb0f387237db7e2495c0358902458a124d7acc7d55866c): complete subsection reference.

<a id="canonical-deb4f35ae9f6b03d5da59e349db0ee3c7637d02fe00742f83412d9563401d57d"></a>

## Next pages — api_testing.domains.credentials / f7a3ef1c4f73 / 5

- [api_testing.domains.credentials.admin](data-sources--http_loadbalancer--reference--group-010.md#canonical-505012b3df1ac809619e0407e4d5d161a70be114fb33cc8341d94b030d460a87)
- [api_testing.domains.credentials.api_key](data-sources--http_loadbalancer--reference--group-010.md#canonical-ef7dfee2b7b52027e49a0e9d727e19417e976a40a116621ae995d9e04b36426f)
- [api_testing.domains.credentials.basic_auth](data-sources--http_loadbalancer--reference--group-010.md#canonical-b3c25a97b23dd1588c93610432003c1f863335c6b0b2abdaa9b96e9e1deaef99)
- [api_testing.domains.credentials.bearer_token](data-sources--http_loadbalancer--reference--group-010.md#canonical-d89982c7f459e67e9e9dd2c66d215f93d88d9ff3614d34e0cefb68f962525926)
- [api_testing.domains.credentials.login_endpoint](data-sources--http_loadbalancer--reference--group-010.md#canonical-c2e3cc2aa570b84c3978c9e016233fe86dd77038c9f1bf8dbf3675d2b928e760)
- [api_testing.domains.credentials.standard](data-sources--http_loadbalancer--reference--group-010.md#canonical-9be1f76ac7a312663fdb0f387237db7e2495c0358902458a124d7acc7d55866c)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-505012b3df1ac809619e0407e4d5d161a70be114fb33cc8341d94b030d460a87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c45942835b18c75d384fe366c2e7b90784b4f1c74b17f2f66aff37d511a8a41"></a>

## api_testing.domains.credentials.admin — api_testing.domains.credentials.admin / 8ad12c355d13 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- api_testing.domains.credentials.admin

<a id="canonical-20777871eca4dbccc853d415dd23f7d9c4b12a2ef577b2e58a3d292c499bcf59"></a>

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

<a id="canonical-359135bd4fd85c9f7d85cd789cd13a95b87f9a88dd36386bcf106f0906420ef8"></a>

## Direct properties — api_testing.domains.credentials.admin / 8ad12c355d13 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-34fc2da7029a4101fd3a32b8be8d07c7a27ca7ace1cd81e6fa42861583a13617"></a>

## Next pages — api_testing.domains.credentials.admin / 8ad12c355d13 / 4

- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ef7dfee2b7b52027e49a0e9d727e19417e976a40a116621ae995d9e04b36426f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3217608991e599deb5968101117602377838e661c1936a6db68cb7b2b7588895"></a>

## api_testing.domains.credentials.api_key — api_testing.domains.credentials.api_key / 447e9bf924d9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- api_testing.domains.credentials.api_key

<a id="canonical-f9e270ebf80f374c73688ca624a1b8e54e995d088d9426c5e64382bbe88c5425"></a>

Type: `"single"`. Computed.

API Key

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2993a15622b5352024ded2574c4a219d505a692ec10dc794ed25e69fa9fda357"></a>

## Direct properties — api_testing.domains.credentials.api_key / 447e9bf924d9 / 3

<a id="canonical-375f63d22a7b39c724d71237fc87446a4197c9a06106ef807f2feaff99615d7d"></a>

<a id="canonical-7cc0fda80e844d21925e55fd12a1ea62b07247be447e67a7b6fb076e4a1b2a27"></a>

## key property — api_testing.domains.credentials.api_key / 447e9bf924d9 / 4

Type: `"string"`. Computed.

Key. Cryptographic key material

Upstream description:

Cryptographic key material

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [value](data-sources--http_loadbalancer--reference--group-010.md#canonical-8d80656772336e38d982c1a585cb1737b293dd7718e5d6ea897ad302f8a1eac3): complete subsection reference.

<a id="canonical-f3083fec618e4f0415ff2175c33c1e7c763439ca7d86e9d8ac94652513677b68"></a>

## Next pages — api_testing.domains.credentials.api_key / 447e9bf924d9 / 5

- [api_testing.domains.credentials.api_key.value](data-sources--http_loadbalancer--reference--group-010.md#canonical-8d80656772336e38d982c1a585cb1737b293dd7718e5d6ea897ad302f8a1eac3)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8d80656772336e38d982c1a585cb1737b293dd7718e5d6ea897ad302f8a1eac3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-276f051c38b7af47e4aa2b4b9df3673801c96c6db283f2aecb13c5d1e4a12455"></a>

## api_testing.domains.credentials.api_key.value — api_testing.domains.credentials.api_key.value / 02a025b610e4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- [api_testing.domains.credentials.api_key](data-sources--http_loadbalancer--reference--group-010.md#canonical-ef7dfee2b7b52027e49a0e9d727e19417e976a40a116621ae995d9e04b36426f)
- api_testing.domains.credentials.api_key.value

<a id="canonical-af0ad25cb293a6e4ef56343b13e1d46497e69ce6907a7be86e350c85076eda70"></a>

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

<a id="canonical-6974041a789998b81458e25d7b9559e3cffc08a82f0c517c42d847d2f442c97f"></a>

## Direct properties — api_testing.domains.credentials.api_key.value / 02a025b610e4 / 3

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-010.md#canonical-60d92d3fcbf98eabb9afdefad8f07b89625852c26baacf0998d180f874ccdc50): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-010.md#canonical-ff3d5f0addfa7649245d126c821d52df6238a290010cd25bad6aeaf52eec9336): complete subsection reference.

<a id="canonical-0a2938304703b9dd58a144bd6d475ca85239607159f25df731b01a79b0e6cea7"></a>

## Next pages — api_testing.domains.credentials.api_key.value / 02a025b610e4 / 4

- [api_testing.domains.credentials.api_key.value.blindfold_secret_info](data-sources--http_loadbalancer--reference--group-010.md#canonical-60d92d3fcbf98eabb9afdefad8f07b89625852c26baacf0998d180f874ccdc50)
- [api_testing.domains.credentials.api_key.value.clear_secret_info](data-sources--http_loadbalancer--reference--group-010.md#canonical-ff3d5f0addfa7649245d126c821d52df6238a290010cd25bad6aeaf52eec9336)
- [api_testing.domains.credentials.api_key](data-sources--http_loadbalancer--reference--group-010.md#canonical-ef7dfee2b7b52027e49a0e9d727e19417e976a40a116621ae995d9e04b36426f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-60d92d3fcbf98eabb9afdefad8f07b89625852c26baacf0998d180f874ccdc50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14447e12922f3f92be8125159d1f41f4759d94513513496d051d5bea090faacd"></a>

## api_testing.domains.credentials.api_key.value.blindfold_secret_info — api_testing.domains.credentials.api_key.value.blindfold_secret_info / 1d23ab175862 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- [api_testing.domains.credentials.api_key](data-sources--http_loadbalancer--reference--group-010.md#canonical-ef7dfee2b7b52027e49a0e9d727e19417e976a40a116621ae995d9e04b36426f)
- [api_testing.domains.credentials.api_key.value](data-sources--http_loadbalancer--reference--group-010.md#canonical-8d80656772336e38d982c1a585cb1737b293dd7718e5d6ea897ad302f8a1eac3)
- api_testing.domains.credentials.api_key.value.blindfold_secret_info

<a id="canonical-64609ef4384a7959b07867f3adcaa9efae4383276fc4616b3c2174fd1886ff75"></a>

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

<a id="canonical-d4fb1aafbd630c232a8cdfb82a6c115dafed3ff1557a434ce02bf3a5f6da72de"></a>

## Direct properties — api_testing.domains.credentials.api_key.value.blindfold_secret_info / 1d23ab175862 / 3

<a id="canonical-b3d296528e2578b77461a82f747c1f79db55f859ae14b3a62edc9b1515560edd"></a>

<a id="canonical-f5d3b15c0f4239a3b04fb0df2519fb04439d55fc7a6e6de2d36af272a684add7"></a>

## decryption_provider property — api_testing.domains.credentials.api_key.value.blindfold_secret_info / 1d23ab175862 / 4

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

<a id="canonical-3325bcfbdd90d6d1890c5b51d73be50896ad73e896118883731f62e9d26f12a8"></a>

<a id="canonical-a5077ea9ae4da9506d6dcdddfb26ed69d62aa4b9dc8b68eb0da702def14886f1"></a>

## location property — api_testing.domains.credentials.api_key.value.blindfold_secret_info / 1d23ab175862 / 5

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

<a id="canonical-28d359f8b8bd9d0a95e121fabc97ddf4bb01d759aa7c9936539946ae79e5ee2d"></a>

<a id="canonical-29c09270e7352e1e2a11cccd0db5448dc10f53ebc4e7dc1c3402482986ec303e"></a>

## store_provider property — api_testing.domains.credentials.api_key.value.blindfold_secret_info / 1d23ab175862 / 6

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

<a id="canonical-f8168782e14092d70c55f3a60c87af668854f391e722b90b06b4f262034d62eb"></a>

## Next pages — api_testing.domains.credentials.api_key.value.blindfold_secret_info / 1d23ab175862 / 7

- [api_testing.domains.credentials.api_key.value](data-sources--http_loadbalancer--reference--group-010.md#canonical-8d80656772336e38d982c1a585cb1737b293dd7718e5d6ea897ad302f8a1eac3)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ff3d5f0addfa7649245d126c821d52df6238a290010cd25bad6aeaf52eec9336"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71cafc805899703a27f3dc859c9a6ff05f999c3cf631c4d696d7515626339d93"></a>

## api_testing.domains.credentials.api_key.value.clear_secret_info — api_testing.domains.credentials.api_key.value.clear_secret_info / 0f11e810aa88 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- [api_testing.domains.credentials.api_key](data-sources--http_loadbalancer--reference--group-010.md#canonical-ef7dfee2b7b52027e49a0e9d727e19417e976a40a116621ae995d9e04b36426f)
- [api_testing.domains.credentials.api_key.value](data-sources--http_loadbalancer--reference--group-010.md#canonical-8d80656772336e38d982c1a585cb1737b293dd7718e5d6ea897ad302f8a1eac3)
- api_testing.domains.credentials.api_key.value.clear_secret_info

<a id="canonical-19808bfb39de7916443a9dd25bf27fd59e5c329fbe6802f2fd844ede89d8feab"></a>

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

<a id="canonical-509aab1cc3c98f8b0b1862e990895310ba6ef1efbc9a52a7f5c127e4519611aa"></a>

## Direct properties — api_testing.domains.credentials.api_key.value.clear_secret_info / 0f11e810aa88 / 3

<a id="canonical-54691aa713b3e9e0f1679eceef09d2547442c8940bc0881b8802d1870bb683c7"></a>

<a id="canonical-bad047d71096a1b150e4b4dc8cd824d7e4c76593603747fd637f010113b4fdf9"></a>

## provider_ref property — api_testing.domains.credentials.api_key.value.clear_secret_info / 0f11e810aa88 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-beea330c41cde1c7edffc8c1f257dd2874fde61fe728bc9d57cc5b7c6c8d349c"></a>

<a id="canonical-3abfbaf1837abac6286b66c95f2275a23e6cdca5a3bccfab2fc791a7da933ea9"></a>

## url property — api_testing.domains.credentials.api_key.value.clear_secret_info / 0f11e810aa88 / 5

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

<a id="canonical-79191b3979715a623899b04594a7b3b9307ed0842354da3435a3230f7675c88e"></a>

## Next pages — api_testing.domains.credentials.api_key.value.clear_secret_info / 0f11e810aa88 / 6

- [api_testing.domains.credentials.api_key.value](data-sources--http_loadbalancer--reference--group-010.md#canonical-8d80656772336e38d982c1a585cb1737b293dd7718e5d6ea897ad302f8a1eac3)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b3c25a97b23dd1588c93610432003c1f863335c6b0b2abdaa9b96e9e1deaef99"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-291220c3c460621b71c3781117fabae8781be21ee0b3c8359f83dcf4afef2ce9"></a>

## api_testing.domains.credentials.basic_auth — api_testing.domains.credentials.basic_auth / 292fdfe76ae1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- api_testing.domains.credentials.basic_auth

<a id="canonical-9816705bed3968199b62d0342c47dffbc07c6e348313ba28560b44102e3d1a12"></a>

Type: `"single"`. Computed.

Basic Authentication.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-51d95ba05bc8e1a5f87287f58a635a06d5b4a86df4f4aa4f696053c20e02beab"></a>

## Direct properties — api_testing.domains.credentials.basic_auth / 292fdfe76ae1 / 3

- [password](data-sources--http_loadbalancer--reference--group-010.md#canonical-f7ea23906dcb05ca7571a9c196e2b1011dd78c7f04a4bcb46ccb91aff46da563): complete subsection reference.

<a id="canonical-20687170227506f1ffa5d0176bbc756e932cdff7914beb1c2e8937b5b59eae1e"></a>

<a id="canonical-0f1b7d03ec81c162291fe3868090db151c6cccde40caf3878623a1577ed4f782"></a>

## user property — api_testing.domains.credentials.basic_auth / 292fdfe76ae1 / 4

Type: `"string"`. Computed.

User. Configuration parameter for user

Upstream description:

Configuration parameter for user

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-ae0e7bcebc2e1d52ce095bf9c8233f9ab0ddb937b5e9ae619b3504dbdced04d5"></a>

## Next pages — api_testing.domains.credentials.basic_auth / 292fdfe76ae1 / 5

- [api_testing.domains.credentials.basic_auth.password](data-sources--http_loadbalancer--reference--group-010.md#canonical-f7ea23906dcb05ca7571a9c196e2b1011dd78c7f04a4bcb46ccb91aff46da563)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f7ea23906dcb05ca7571a9c196e2b1011dd78c7f04a4bcb46ccb91aff46da563"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d106add91257318cb7a615c8c53b730ebb227d1a25a5d3254e2727b925464600"></a>

## api_testing.domains.credentials.basic_auth.password — api_testing.domains.credentials.basic_auth.password / b72e994a8fae / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- [api_testing.domains.credentials.basic_auth](data-sources--http_loadbalancer--reference--group-010.md#canonical-b3c25a97b23dd1588c93610432003c1f863335c6b0b2abdaa9b96e9e1deaef99)
- api_testing.domains.credentials.basic_auth.password

<a id="canonical-4d642a78f5c5c295276ce1c6d92ba0378f5dc87a405245828aacdecc7303978a"></a>

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

<a id="canonical-3cf1134ecdd9239a2c69376f137c6d40d7da7244ab6bbede13e70fe369921fe6"></a>

## Direct properties — api_testing.domains.credentials.basic_auth.password / b72e994a8fae / 3

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-010.md#canonical-fb8710423f1a0013469320a93ecf190666a54383ce6fffd7f662ae6fbe62c305): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-010.md#canonical-be984292cf393e294b537707180904a2ec6d83396a9f31a48cea1c3584196eae): complete subsection reference.

<a id="canonical-2d51b000aa3464243b0fe7bd439f0894218c5c8d476c2b4d86cb185cf9430905"></a>

## Next pages — api_testing.domains.credentials.basic_auth.password / b72e994a8fae / 4

- [api_testing.domains.credentials.basic_auth.password.blindfold_secret_info](data-sources--http_loadbalancer--reference--group-010.md#canonical-fb8710423f1a0013469320a93ecf190666a54383ce6fffd7f662ae6fbe62c305)
- [api_testing.domains.credentials.basic_auth.password.clear_secret_info](data-sources--http_loadbalancer--reference--group-010.md#canonical-be984292cf393e294b537707180904a2ec6d83396a9f31a48cea1c3584196eae)
- [api_testing.domains.credentials.basic_auth](data-sources--http_loadbalancer--reference--group-010.md#canonical-b3c25a97b23dd1588c93610432003c1f863335c6b0b2abdaa9b96e9e1deaef99)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-fb8710423f1a0013469320a93ecf190666a54383ce6fffd7f662ae6fbe62c305"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f90f2246a0279b92f9c74f393972c03fb6137e3b8631a586848f09ea719864f"></a>

## api_testing.domains.credentials.basic_auth.password.blindfold_secret_info — api_testing.domains.credentials.basic_auth.password.blindfold_secret_info / f6835a1c3e55 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- [api_testing.domains.credentials.basic_auth](data-sources--http_loadbalancer--reference--group-010.md#canonical-b3c25a97b23dd1588c93610432003c1f863335c6b0b2abdaa9b96e9e1deaef99)
- [api_testing.domains.credentials.basic_auth.password](data-sources--http_loadbalancer--reference--group-010.md#canonical-f7ea23906dcb05ca7571a9c196e2b1011dd78c7f04a4bcb46ccb91aff46da563)
- api_testing.domains.credentials.basic_auth.password.blindfold_secret_info

<a id="canonical-085a78cb24e1373023bf12a622a45899ad4a3a7c999433f3f6552ca9b26cf241"></a>

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

<a id="canonical-650683942ef055fef88829496e142c010b8e1da6c647fac6c35e0fe51ca4e88c"></a>

## Direct properties — api_testing.domains.credentials.basic_auth.password.blindfold_secret_info / f6835a1c3e55 / 3

<a id="canonical-ed7b7ba4127d767941d51dfb660e04550dbcf527c474a10069e446afc45cf7c2"></a>

<a id="canonical-f5d7baa47c10c970df879512d5cc082a052e8f3b0b43d8b5d6c2f5cf416b4b49"></a>

## decryption_provider property — api_testing.domains.credentials.basic_auth.password.blindfold_secret_info / f6835a1c3e55 / 4

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

<a id="canonical-d268f59c1d22f141c6d0a0c670b30c371387be0ec985e0a6f40287a969268e99"></a>

<a id="canonical-d858fc0a49a567a12697af603c535937176cd0f2040e2ebf19153608af69be2d"></a>

## location property — api_testing.domains.credentials.basic_auth.password.blindfold_secret_info / f6835a1c3e55 / 5

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

<a id="canonical-acaf2259829d696c918273ac17800e91f2be779ff1cdda3451d3c503d01d7c68"></a>

<a id="canonical-29e10340b7e45acae740d8b400cf8d7f877999853a2782598a9ac9a75b014a04"></a>

## store_provider property — api_testing.domains.credentials.basic_auth.password.blindfold_secret_info / f6835a1c3e55 / 6

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

<a id="canonical-32e827e744bdabd8afaa6242d81efb3a2b4f077044d1a0307fae37113f3e6b41"></a>

## Next pages — api_testing.domains.credentials.basic_auth.password.blindfold_secret_info / f6835a1c3e55 / 7

- [api_testing.domains.credentials.basic_auth.password](data-sources--http_loadbalancer--reference--group-010.md#canonical-f7ea23906dcb05ca7571a9c196e2b1011dd78c7f04a4bcb46ccb91aff46da563)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-be984292cf393e294b537707180904a2ec6d83396a9f31a48cea1c3584196eae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b881ff979c3c4341d8904f9d0b85b5d2bab51862de617eb2101e39f512b736c4"></a>

## api_testing.domains.credentials.basic_auth.password.clear_secret_info — api_testing.domains.credentials.basic_auth.password.clear_secret_info / fb99cc218c28 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- [api_testing.domains.credentials.basic_auth](data-sources--http_loadbalancer--reference--group-010.md#canonical-b3c25a97b23dd1588c93610432003c1f863335c6b0b2abdaa9b96e9e1deaef99)
- [api_testing.domains.credentials.basic_auth.password](data-sources--http_loadbalancer--reference--group-010.md#canonical-f7ea23906dcb05ca7571a9c196e2b1011dd78c7f04a4bcb46ccb91aff46da563)
- api_testing.domains.credentials.basic_auth.password.clear_secret_info

<a id="canonical-c1812e66792c91a00881cae4100bc714f39897f07977cfd5f6991ff188264fc7"></a>

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

<a id="canonical-554fcce9f804988ddc0e96ca25de183a54bd5eb1e48d20550ec6413050eee842"></a>

## Direct properties — api_testing.domains.credentials.basic_auth.password.clear_secret_info / fb99cc218c28 / 3

<a id="canonical-44b718de5a3f2a395b35d1f999a95753bf2407e2a6495a76196161c69de4da16"></a>

<a id="canonical-a2700a91f878c1b66f52eb0b55cd1fe3f08e0f0d08ebaa3f72893735526bf4a0"></a>

## provider_ref property — api_testing.domains.credentials.basic_auth.password.clear_secret_info / fb99cc218c28 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-db870aa930f6d9ae4e2386e7cad0ebe3b9365156f8afecaad58d853383def8fe"></a>

<a id="canonical-15c2b07858a3166526b8af80e7cdeb5763177cb45c9891490188d2dbae1bbe73"></a>

## url property — api_testing.domains.credentials.basic_auth.password.clear_secret_info / fb99cc218c28 / 5

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

<a id="canonical-883606b0dd1bdae16534eb1cbac5aff1f6c241badc853c80080c92e48ddaaf73"></a>

## Next pages — api_testing.domains.credentials.basic_auth.password.clear_secret_info / fb99cc218c28 / 6

- [api_testing.domains.credentials.basic_auth.password](data-sources--http_loadbalancer--reference--group-010.md#canonical-f7ea23906dcb05ca7571a9c196e2b1011dd78c7f04a4bcb46ccb91aff46da563)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d89982c7f459e67e9e9dd2c66d215f93d88d9ff3614d34e0cefb68f962525926"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b54a80dfe02ee38a1db566a6a8157830f5a783dd6a6a7e930bb6e798f41d02f"></a>

## api_testing.domains.credentials.bearer_token — api_testing.domains.credentials.bearer_token / f8a3ed8c57dd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- api_testing.domains.credentials.bearer_token

<a id="canonical-0ea793d1b03545b07e41910a6c32a61e11dcf3d0aca6129e6b52967686b94b2b"></a>

Type: `"single"`. Computed.

Configuration parameter for bearer token.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-77468acab7969a564a0971f093094547c58a8bd32808baba6a94d604f4428760"></a>

## Direct properties — api_testing.domains.credentials.bearer_token / f8a3ed8c57dd / 3

- [token](data-sources--http_loadbalancer--reference--group-010.md#canonical-2a6345740fd504e387d960bb4ac3779717ac006b7bea160e08a253c873aed714): complete subsection reference.

<a id="canonical-d6c7bcd70c852d3034bbf17bc5649374a9d19929ab8ee89f904c23a207d354ff"></a>

## Next pages — api_testing.domains.credentials.bearer_token / f8a3ed8c57dd / 4

- [api_testing.domains.credentials.bearer_token.token](data-sources--http_loadbalancer--reference--group-010.md#canonical-2a6345740fd504e387d960bb4ac3779717ac006b7bea160e08a253c873aed714)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2a6345740fd504e387d960bb4ac3779717ac006b7bea160e08a253c873aed714"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e04746b0f160420f0e12db1ded7ad5a01fbe0287a4eaec291dd408e4d1aafd86"></a>

## api_testing.domains.credentials.bearer_token.token — api_testing.domains.credentials.bearer_token.token / ab9317aab119 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- [api_testing.domains.credentials.bearer_token](data-sources--http_loadbalancer--reference--group-010.md#canonical-d89982c7f459e67e9e9dd2c66d215f93d88d9ff3614d34e0cefb68f962525926)
- api_testing.domains.credentials.bearer_token.token

<a id="canonical-defa66e3db9f414bcd834fc03f6eb7bbca418dc3f6766c8236f8db2233e60b1e"></a>

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

<a id="canonical-659e3cba82c5a816fa43f069dffac3bdc012911398b1a7167ff7bde132d3c223"></a>

## Direct properties — api_testing.domains.credentials.bearer_token.token / ab9317aab119 / 3

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-010.md#canonical-3f1df1e12bae3e9446ed5eac1460239594d3e3e5bdb1889aad5951ac72b8c7e3): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-010.md#canonical-1ccaf03ecc4eaef50561cd7971c09beb8ae692ab4a90497bb5793d6c1e160247): complete subsection reference.

<a id="canonical-d05d88d6789020f121013a2284eb5fad61059206d69f7b2b148c43e3e4aac640"></a>

## Next pages — api_testing.domains.credentials.bearer_token.token / ab9317aab119 / 4

- [api_testing.domains.credentials.bearer_token.token.blindfold_secret_info](data-sources--http_loadbalancer--reference--group-010.md#canonical-3f1df1e12bae3e9446ed5eac1460239594d3e3e5bdb1889aad5951ac72b8c7e3)
- [api_testing.domains.credentials.bearer_token.token.clear_secret_info](data-sources--http_loadbalancer--reference--group-010.md#canonical-1ccaf03ecc4eaef50561cd7971c09beb8ae692ab4a90497bb5793d6c1e160247)
- [api_testing.domains.credentials.bearer_token](data-sources--http_loadbalancer--reference--group-010.md#canonical-d89982c7f459e67e9e9dd2c66d215f93d88d9ff3614d34e0cefb68f962525926)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3f1df1e12bae3e9446ed5eac1460239594d3e3e5bdb1889aad5951ac72b8c7e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-278cc902a1928966b318f72e842b360e667a12230a77a5a4a155a9008b78f358"></a>

## api_testing.domains.credentials.bearer_token.token.blindfold_secret_info — api_testing.domains.credentials.bearer_token.token.blindfold_secret_info / 1c5735becbf9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- [api_testing.domains.credentials.bearer_token](data-sources--http_loadbalancer--reference--group-010.md#canonical-d89982c7f459e67e9e9dd2c66d215f93d88d9ff3614d34e0cefb68f962525926)
- [api_testing.domains.credentials.bearer_token.token](data-sources--http_loadbalancer--reference--group-010.md#canonical-2a6345740fd504e387d960bb4ac3779717ac006b7bea160e08a253c873aed714)
- api_testing.domains.credentials.bearer_token.token.blindfold_secret_info

<a id="canonical-cd0c7e25aa0f1254d05106a37c4d566fc47aaacd105ebcb64ac966072cc3f013"></a>

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

<a id="canonical-24248014ef8e00e0d1245ff69863c995cb92d1029b2d1c9e82b8c8ed4e86ca09"></a>

## Direct properties — api_testing.domains.credentials.bearer_token.token.blindfold_secret_info / 1c5735becbf9 / 3

<a id="canonical-ff05e6eceb5031552c6f9ac526c890644811e99f581a5fa5c3b0e848680f3b86"></a>

<a id="canonical-99652b92cfef5688a23fa6bb5fe85d6657a1a9592cde05f86c5d6c3de04c8b83"></a>

## decryption_provider property — api_testing.domains.credentials.bearer_token.token.blindfold_secret_info / 1c5735becbf9 / 4

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

<a id="canonical-4c9fe18d9eaac9c3d60f1a56b910d6fba2a0b5227bd01b3f092649221e80b290"></a>

<a id="canonical-ee09f547f2efb0a0c7a6512a71f4ca5f1c616ea4da00de6ef3855c9d65ee488e"></a>

## location property — api_testing.domains.credentials.bearer_token.token.blindfold_secret_info / 1c5735becbf9 / 5

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

<a id="canonical-b38b635123f831d9a4feb5192d54c1f154e3308a26d126fb585f95fe82ff943d"></a>

<a id="canonical-32f053b074bfe0a1cfcbcf2b4e13623e470a4fdbdf82d17144e2b1644eeeabcb"></a>

## store_provider property — api_testing.domains.credentials.bearer_token.token.blindfold_secret_info / 1c5735becbf9 / 6

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

<a id="canonical-4d5fc20a9e18f8e44346de0b326bd6617f2eee2e80ec1e9eb2c3fb8d3314dfdf"></a>

## Next pages — api_testing.domains.credentials.bearer_token.token.blindfold_secret_info / 1c5735becbf9 / 7

- [api_testing.domains.credentials.bearer_token.token](data-sources--http_loadbalancer--reference--group-010.md#canonical-2a6345740fd504e387d960bb4ac3779717ac006b7bea160e08a253c873aed714)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1ccaf03ecc4eaef50561cd7971c09beb8ae692ab4a90497bb5793d6c1e160247"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3cddeb15294de74229e2864018e898c241ec71586a2b47c51ebc7ca35cee8df"></a>

## api_testing.domains.credentials.bearer_token.token.clear_secret_info — api_testing.domains.credentials.bearer_token.token.clear_secret_info / 45fd5d4dbf79 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- [api_testing.domains.credentials.bearer_token](data-sources--http_loadbalancer--reference--group-010.md#canonical-d89982c7f459e67e9e9dd2c66d215f93d88d9ff3614d34e0cefb68f962525926)
- [api_testing.domains.credentials.bearer_token.token](data-sources--http_loadbalancer--reference--group-010.md#canonical-2a6345740fd504e387d960bb4ac3779717ac006b7bea160e08a253c873aed714)
- api_testing.domains.credentials.bearer_token.token.clear_secret_info

<a id="canonical-e568ac1dbdae1f8e2d9997a0301ee4913515f291e5c09fcece1e1437be595c04"></a>

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

<a id="canonical-c3230695ba8bdb58ddc45d017afbb7f04c95b65359472e703a3872da5dc1eadd"></a>

## Direct properties — api_testing.domains.credentials.bearer_token.token.clear_secret_info / 45fd5d4dbf79 / 3

<a id="canonical-c1de3504e1d18508569b6e1fb5c9a6c7483cefec62a57614f9c2f324b4b69842"></a>

<a id="canonical-812ab2ec39ca961f2339497a5160fc406dc66817e02fd81aab16c9226a95de0e"></a>

## provider_ref property — api_testing.domains.credentials.bearer_token.token.clear_secret_info / 45fd5d4dbf79 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-99c8936d385256a571fbe80b7348fa28390e486e6f10917d8b7ebd075e1470ce"></a>

<a id="canonical-f85d8772a7a27d333086e69bebe0ef6f7b489886149e5d82a503ec9356c29ccb"></a>

## url property — api_testing.domains.credentials.bearer_token.token.clear_secret_info / 45fd5d4dbf79 / 5

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

<a id="canonical-a4d1258484454a542ac7ab56e8c36d89afb8818514a8dede167ec84e33a5126e"></a>

## Next pages — api_testing.domains.credentials.bearer_token.token.clear_secret_info / 45fd5d4dbf79 / 6

- [api_testing.domains.credentials.bearer_token.token](data-sources--http_loadbalancer--reference--group-010.md#canonical-2a6345740fd504e387d960bb4ac3779717ac006b7bea160e08a253c873aed714)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c2e3cc2aa570b84c3978c9e016233fe86dd77038c9f1bf8dbf3675d2b928e760"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f094edf53a5dab3f14ad3477576e86be583d9f6c81de8b628486035543b39db"></a>

## api_testing.domains.credentials.login_endpoint — api_testing.domains.credentials.login_endpoint / 3af5224d8f11 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- api_testing.domains.credentials.login_endpoint

<a id="canonical-c38a01d4b7a0377c4045d74743a0d7c804c794568862ef37152b39c95fd5265e"></a>

Type: `"single"`. Computed.

Login Endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9b53d0ffa9d0f3b304d47982e2ae94fdb02ecb62a349643c598b6d1d25bbdb94"></a>

## Direct properties — api_testing.domains.credentials.login_endpoint / 3af5224d8f11 / 3

- [json_payload](data-sources--http_loadbalancer--reference--group-010.md#canonical-eb72622ae6ac75a5c4ad46a9bbc70b59d527a541ce7fc0400e61bce5920da0f1): complete subsection reference.

<a id="canonical-2f5f211c1d98e9aeadff01e6c1e1bb6eb3fef991797d66a9df2511e549e4041d"></a>

<a id="canonical-2d226d277d56ccb18723a09c83c80e4aab1918412186bf4fa7fd0d6353389581"></a>

## method property — api_testing.domains.credentials.login_endpoint / 3af5224d8f11 / 4

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ae12050e98bacf80c8671883e3c1dac6a5e140b89b7656e2245d6c6b3869340a"></a>

<a id="canonical-e37cace2e8b72b8addec05fdf6f47cc235d333d25446c41f429878eb6d5c1da2"></a>

## path property — api_testing.domains.credentials.login_endpoint / 3af5224d8f11 / 5

Type: `"string"`. Computed.

Path. URL path for the endpoint

Upstream description:

URL path for the endpoint

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "1024"
  }
}
```

<a id="canonical-d09422db9097f8969deae027445d6c5f8ec8e59da6af2315786d6c1211bb8f07"></a>

<a id="canonical-be35558627cf8dd15f4cbae39271a60e6567678f3bd74457ae636850da3839ed"></a>

## token_response_key property — api_testing.domains.credentials.login_endpoint / 3af5224d8f11 / 6

Type: `"string"`. Computed.

Specifies the key name used to extract the authentication token from the login response, such as
token or access\_token.

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

<a id="canonical-f4534112e3c7ab625cff09a972bb6ca43ad75a161ee75f14947535ffe1bf535f"></a>

## Next pages — api_testing.domains.credentials.login_endpoint / 3af5224d8f11 / 7

- [api_testing.domains.credentials.login_endpoint.json_payload](data-sources--http_loadbalancer--reference--group-010.md#canonical-eb72622ae6ac75a5c4ad46a9bbc70b59d527a541ce7fc0400e61bce5920da0f1)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-eb72622ae6ac75a5c4ad46a9bbc70b59d527a541ce7fc0400e61bce5920da0f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f40f81c7d215071c3826c8e84c159616827582487d0cda41973633810977634"></a>

## api_testing.domains.credentials.login_endpoint.json_payload — api_testing.domains.credentials.login_endpoint.json_payload / 2d268d642f50 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- [api_testing.domains.credentials.login_endpoint](data-sources--http_loadbalancer--reference--group-010.md#canonical-c2e3cc2aa570b84c3978c9e016233fe86dd77038c9f1bf8dbf3675d2b928e760)
- api_testing.domains.credentials.login_endpoint.json_payload

<a id="canonical-a94e429815c1f6d95ae65fdf660c681996735437899da2729409752a553fa0a7"></a>

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

<a id="canonical-3e7be4d6f315b06341b27cdb2220d7b138ca3b0115147d3d8de11b0e17b746ca"></a>

## Direct properties — api_testing.domains.credentials.login_endpoint.json_payload / 2d268d642f50 / 3

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-010.md#canonical-d38426d7ac17ed40a411040240863df7224ae9e7181e6d319ce2facce721bd50): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-010.md#canonical-f2f0ae424be07732fef5b6172ee5ac5e60f3e27c99d65d6887550f45ee957885): complete subsection reference.

<a id="canonical-2ceb4f9fa4995a4712b7da2bb8973db946bf2fdfb654a5b87b5020c5c87f7016"></a>

## Next pages — api_testing.domains.credentials.login_endpoint.json_payload / 2d268d642f50 / 4

- [api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_info](data-sources--http_loadbalancer--reference--group-010.md#canonical-d38426d7ac17ed40a411040240863df7224ae9e7181e6d319ce2facce721bd50)
- [api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info](data-sources--http_loadbalancer--reference--group-010.md#canonical-f2f0ae424be07732fef5b6172ee5ac5e60f3e27c99d65d6887550f45ee957885)
- [api_testing.domains.credentials.login_endpoint](data-sources--http_loadbalancer--reference--group-010.md#canonical-c2e3cc2aa570b84c3978c9e016233fe86dd77038c9f1bf8dbf3675d2b928e760)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d38426d7ac17ed40a411040240863df7224ae9e7181e6d319ce2facce721bd50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-456a1a33f909b6c14f75bc53cff3738681219f0abce9e16a571cf5fa61d1b4d0"></a>

## api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_info — api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_inf / 940bc86eb7d4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- [api_testing.domains.credentials.login_endpoint](data-sources--http_loadbalancer--reference--group-010.md#canonical-c2e3cc2aa570b84c3978c9e016233fe86dd77038c9f1bf8dbf3675d2b928e760)
- [api_testing.domains.credentials.login_endpoint.json_payload](data-sources--http_loadbalancer--reference--group-010.md#canonical-eb72622ae6ac75a5c4ad46a9bbc70b59d527a541ce7fc0400e61bce5920da0f1)
- api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_info

<a id="canonical-b0cc8553efb9463f359ff5e94f05da1480a83e5b8b30e4a0533f4cea01d1693f"></a>

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

<a id="canonical-5ea3eeae171184d2755741974353c244c1d60dc46aa51335d8bba0f84539ed08"></a>

## Direct properties — api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_inf / 940bc86eb7d4 / 3

<a id="canonical-7111f1a2b7a057f01024cff6a69124191a4fe531e45fe1d637df5d0537894829"></a>

<a id="canonical-99f161b489b23dc77e3230bce86335e23198542aeabf42ea1878cec6d4a47e3a"></a>

## decryption_provider property — api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_inf / 940bc86eb7d4 / 4

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

<a id="canonical-f7fbdcdd79aad540d9c2042f9b4a9453c08be14f7cf0ff37d77e71c25358265e"></a>

<a id="canonical-21d153cd0d620b692f4f0cf1ff59fce3ecc3260596d55907b703e58816513f3e"></a>

## location property — api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_inf / 940bc86eb7d4 / 5

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

<a id="canonical-914fd4a6c25c9427fb0ee67718a7f423bea84209d486024531d774c78e755271"></a>

<a id="canonical-b7897b7c2fad9b36cee5835c970b6f5ec07835a08be2d85e409e769f2b6df427"></a>

## store_provider property — api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_inf / 940bc86eb7d4 / 6

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

<a id="canonical-9bf81281e4178dffb61153563a75a886a5edb011d856c1962690f7c6bd33e140"></a>

## Next pages — api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_inf / 940bc86eb7d4 / 7

- [api_testing.domains.credentials.login_endpoint.json_payload](data-sources--http_loadbalancer--reference--group-010.md#canonical-eb72622ae6ac75a5c4ad46a9bbc70b59d527a541ce7fc0400e61bce5920da0f1)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f2f0ae424be07732fef5b6172ee5ac5e60f3e27c99d65d6887550f45ee957885"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf4088ccae8c0b885085797a4a2297441bee4e9ecf31e3df3e4bc9d315083921"></a>

## api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info — api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info / 8897ef8b349c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- [api_testing.domains.credentials.login_endpoint](data-sources--http_loadbalancer--reference--group-010.md#canonical-c2e3cc2aa570b84c3978c9e016233fe86dd77038c9f1bf8dbf3675d2b928e760)
- [api_testing.domains.credentials.login_endpoint.json_payload](data-sources--http_loadbalancer--reference--group-010.md#canonical-eb72622ae6ac75a5c4ad46a9bbc70b59d527a541ce7fc0400e61bce5920da0f1)
- api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info

<a id="canonical-6d8754ef93292d78ded705b869a0804009b31a8394947a698d635d048d692cad"></a>

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

<a id="canonical-eced485d9d1b83bc65fb6f348d78ce71818981b435dba701f8107189a439ab34"></a>

## Direct properties — api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info / 8897ef8b349c / 3

<a id="canonical-04bdb5fc29cd698f3e12bb53cc5edf4de03e205da449f2a32618fae6df7f7c64"></a>

<a id="canonical-c9de45d005e4a0df016bcc8c2b4e812c00dfb5e5b835b558ca933ea3817f799e"></a>

## provider_ref property — api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info / 8897ef8b349c / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-f8c189a4348866180520cd8d653806b899746cb91ebd7dcce7c7e42ba2300cdc"></a>

<a id="canonical-59cd7a56a4db18fef51ef18419fb89a0df088afa34ccbf5e7283a71bdbf14034"></a>

## url property — api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info / 8897ef8b349c / 5

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

<a id="canonical-f4318d76b8255b454f92fe046a0f43f13bc1ebb4361e91056686d2c8eaf4a48d"></a>

## Next pages — api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info / 8897ef8b349c / 6

- [api_testing.domains.credentials.login_endpoint.json_payload](data-sources--http_loadbalancer--reference--group-010.md#canonical-eb72622ae6ac75a5c4ad46a9bbc70b59d527a541ce7fc0400e61bce5920da0f1)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9be1f76ac7a312663fdb0f387237db7e2495c0358902458a124d7acc7d55866c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11aad2b9339060c464c8ffb76e8319fd44d57221906e4dc5f5c64985b7485b54"></a>

## api_testing.domains.credentials.standard — api_testing.domains.credentials.standard / 685ba3d5ddaf / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [api_testing.domains](data-sources--http_loadbalancer--reference--group-010.md#canonical-06aa0176b5ec49aa7c2ca1db640cd4fd89a65e8ce2b2df6f438f71d512cdcebb)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- api_testing.domains.credentials.standard

<a id="canonical-0fb6d6acc2e0bdc915b29d783988f462d1eeb71906b25be3e7062079cfb9db3c"></a>

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

<a id="canonical-6ca89654bc9ef65cca2e061b0eb375663a0b729d5977703854f3a827f417ac78"></a>

## Direct properties — api_testing.domains.credentials.standard / 685ba3d5ddaf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ff0b278542bbc76d07fbf46badd149ddd3d4e4c0985561410b40aa87fd96c734"></a>

## Next pages — api_testing.domains.credentials.standard / 685ba3d5ddaf / 4

- [api_testing.domains.credentials](data-sources--http_loadbalancer--reference--group-010.md#canonical-6425c428cd27ff6636951697d76d890607e19b63781e38131421f58e335c8fa1)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d16625fa6562db275fe1d05b341d9879205d06f27dd3b9eb4d65e09fdb5acd00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-376307a7f1278ffcee1c9758cf3ef7645f094d64ad4b054b4fc4862dd29500f1"></a>

## api_testing.every_day — api_testing.every_day / 1b6441bc7d0e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- api_testing.every_day

<a id="canonical-184c4a32a90aabfda8e86da8497e1d898a14173d37842fb1863f7a5c3a7cdfb3"></a>

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

<a id="canonical-d17f5d008aac3f20f40657be3163c970eb0f4fddca32e5027bf9f7e1e9ed2591"></a>

## Direct properties — api_testing.every_day / 1b6441bc7d0e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eb89ff576086a017c4c09fcd7ebf2474d49f9c9f86bb3d63f2e905e3b8ecf55b"></a>

## Next pages — api_testing.every_day / 1b6441bc7d0e / 4

- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0a02f2b4f94f3a8749ec70ad4653a1171e6c169a17f723204b4ece102e1b5924"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c308da8c72c735777ec9f72cfe89b074fdd54a23e40714a54c0da2af31423d4"></a>

## api_testing.every_month — api_testing.every_month / e649e6a8eabc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- api_testing.every_month

<a id="canonical-50fb4587eefd2153665ef359b72b2688c1c41e682aa9ab25a924e904ef24439f"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for every month.

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

<a id="canonical-d1aa99d9d2105c0b532d9e65f8f740b3efb59f5197d706f31f4fec840fa046d2"></a>

## Direct properties — api_testing.every_month / e649e6a8eabc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0acd3318f9d34f5aac619631834f5e2e20e778caac88dd0a4e21363392ac3968"></a>

## Next pages — api_testing.every_month / e649e6a8eabc / 4

- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-04a403d827689939d030abd23a47f97de7b553937de3f4b12a68a4dfacb969ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0fdf422d85b52256e4d41ccd4a8c6c625a290e64c36f5885d6d9223bc5fa7dcb"></a>

## api_testing.every_week — api_testing.every_week / d173e32fe859 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- api_testing.every_week

<a id="canonical-4fa43f49a1ec4ce844603dae55e3b9012b191620a7fbb37965f894bba6d00278"></a>

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

<a id="canonical-84aa30dde7f2dbf6a67e54c174e097bee7b10f69b727d72426617dcc6eda9c74"></a>

## Direct properties — api_testing.every_week / d173e32fe859 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-996ab6ed1ee94a7b87c1209d3ee8c94c1c6cbaaa81389bb714d5d95e76db77b2"></a>

## Next pages — api_testing.every_week / d173e32fe859 / 4

- [api_testing](data-sources--http_loadbalancer--reference--group-010.md#canonical-da1f350fc92754b7fb2659d2909df96d85c47c622ffaceb53e4510546bcdcb76)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a2c8a93200e97df3ae4c03b87e5920ac3ccf3a34ff325cacbecfe6c9a5e575f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ce4caa5ba639d550ff869ba427898c3478cc170db31743d4551d88d64df633d"></a>

## app_firewall — app_firewall / 318fa046add1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- app_firewall

<a id="canonical-7e56ab0b22a0a3225bec9d48086921b72a1c07264bf0e108da7d838e1313bc87"></a>

Type: `"single"`. Computed.

\[OneOf: app\_firewall, disable\_waf; Default: disable\_waf\] Type establishes a direct reference
from one object(the referrer) to another(the referred). Such a reference is in form of
tenant/namespace/name.

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

OneOf alternatives in this subsection:

- [app_firewall](data-sources--http_loadbalancer--reference--group-010.md#canonical-7e56ab0b22a0a3225bec9d48086921b72a1c07264bf0e108da7d838e1313bc87)
- [disable_waf](data-sources--http_loadbalancer--reference--group-017.md#canonical-b3ddcdbb722e309fe97016afb736bebfe8bf3480d4ad438b6a705f36bf04cfd7)

Select alternatives according to the provider validators above.

<a id="canonical-577f58249e22e436a2419eff63e509da8f8e72c1570b88d5c5d273ec245bbe4f"></a>

## Direct properties — app_firewall / 318fa046add1 / 3

<a id="canonical-a309dbac1a0e5a138f2e6154d5049cacb1c162e871529241827e3381510a3571"></a>

<a id="canonical-99441a6671abf3fee21725bfed3d26234e318903a048026b90b39a8f3a0e9c90"></a>

## name property — app_firewall / 318fa046add1 / 4

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

<a id="canonical-7b40a48c29e8f2305fd10b40ce938f6f2396702c5b3868b81842c50fa88eed5f"></a>

<a id="canonical-b5c4d0b7d50fa9728e113c74178d1f5b11f17a1747a7bad41ab5af3150cd4808"></a>

## namespace property — app_firewall / 318fa046add1 / 5

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

<a id="canonical-bccf61dd96aae6c830cb171d2da7fbf14c5b97efef705f75e3c2a79b6f773d87"></a>

<a id="canonical-01c6df4558a74ea9213438ed6af3b92fb308561b2aa294b0726d865e5976964e"></a>

## tenant property — app_firewall / 318fa046add1 / 6

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

<a id="canonical-159d3af2922b2b6a1f9dffdad5877f222f859eb67f9a2e5f296f6680fddc5a71"></a>

## Next pages — app_firewall / 318fa046add1 / 7

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e46710ba001bf8ddabcfccfbb5094215ce20da3a35b0ea48470159cc0e080303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72bae1a075da5332c8844eacfa913bd47a0ad13af065ba9747da334a9c429d73"></a>

## blocked_clients — blocked_clients / 59e3ba0b390d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- blocked_clients

<a id="canonical-82cb733e6e6e5e4fe2dc1938f7fc4121de18b65028b4548ca8e9adbf59eea2f4"></a>

Type: `"list"`. Computed.

Define rules to block IP Prefixes or AS numbers.

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

<a id="canonical-073b897f7ba68a41c4302e8645f59bd0b7a2fe6611b79e2c1a9a52d3035480d4"></a>

## Direct properties — blocked_clients / 59e3ba0b390d / 3

<a id="canonical-9ba1b3938e03dd3c8bb948e750dde4df8d24ba165974dd95c53272a7b4f831e5"></a>

<a id="canonical-28860b5d5f877d6f2554ffddcfefec6947d6501ef403840805549ae4c6cf5c36"></a>

## actions property — blocked_clients / 59e3ba0b390d / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-99e576492d94442d243a3f6c18aa640990bc07cb84ddd6dd9f01b7a6e87c044e"></a>

<a id="canonical-aa8c00a33dea0ca8bde3cac7f9b3b587717ceb3470e1f56795343b0814ce1bca"></a>

## as_number property — blocked_clients / 59e3ba0b390d / 5

Type: `"number"`. Computed.

Exclusive with \[http\_header ip\_prefix ipv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Upstream description:

Exclusive with \[http\_header ip\_prefix ipv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

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

- [bot_skip_processing](data-sources--http_loadbalancer--reference--group-010.md#canonical-89b1f160e9c5b67b565c80cebec0b7491d42fbe72567f05d3801d45968d067df): complete subsection reference.

<a id="canonical-2f88ac922e35944f4f33cbd846230feabfcf6fb2192d61e10935e15f1aef5442"></a>

<a id="canonical-240e019bc09cefc12a590c94bd040c1bede25e81d5f496fc021f107527f0f358"></a>

## expiration_timestamp property — blocked_clients / 59e3ba0b390d / 6

Type: `"string"`. Computed.

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

- [http_header](data-sources--http_loadbalancer--reference--group-010.md#canonical-5c1b98f99626888192e3ef6181bb0fc0f9936e83cdbf18b745d797f56ddcd20a): complete subsection reference.

<a id="canonical-3aac32835d716429741a1cea1fd037e67d58508c54d116604deb5931c4babe43"></a>

<a id="canonical-d94bb7e32944ce2e400a99e4aafb7d6639756b50b0a180db3663c900df3f1a13"></a>

## ip_prefix property — blocked_clients / 59e3ba0b390d / 7

Type: `"string"`. Computed.

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

<a id="canonical-8dd110b1d9d28525ab8064e0dfea8be0a439342cb03b16259e482131f1c71ab0"></a>

<a id="canonical-0e7d974b8d9474259c541363f5193bf4dd757952d16e9243a84dce8512684b84"></a>

## ipv6_prefix property — blocked_clients / 59e3ba0b390d / 8

Type: `"string"`. Computed.

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

- [metadata](data-sources--http_loadbalancer--reference--group-010.md#canonical-7a2529e60d610606409d0803a357a1d16a60f79b878f30a3e5a8079b092eb0c2): complete subsection reference.

- [skip_processing](data-sources--http_loadbalancer--reference--group-010.md#canonical-9965e647eba8dd09adeb3bc9b042c06d265c31e8b4a62451ae8bcf397aa29d5d): complete subsection reference.

<a id="canonical-e67429182e5ec4c35fc0423560d5db144928240fe129535056684471d659ec58"></a>

<a id="canonical-e5a05f491602eb541d6d33dc66a60327531ffd6d4ff6443ecd839c7c42ccdee1"></a>

## user_identifier property — blocked_clients / 59e3ba0b390d / 9

Type: `"string"`. Computed.

Exclusive with \[as\_number http\_header ip\_prefix ipv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

Upstream description:

Exclusive with \[as\_number http\_header ip\_prefix ipv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

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

- [waf_skip_processing](data-sources--http_loadbalancer--reference--group-010.md#canonical-5bf6cb533cc29130ca90c16ba7d9a325a1ce3f74038707bc0c16f987dff3b744): complete subsection reference.

<a id="canonical-1ce4b4267541e3b2bc235bba583e1e808d828b182a2610dbe76bb3d13826ce19"></a>

## Next pages — blocked_clients / 59e3ba0b390d / 10

- [blocked_clients.bot_skip_processing](data-sources--http_loadbalancer--reference--group-010.md#canonical-89b1f160e9c5b67b565c80cebec0b7491d42fbe72567f05d3801d45968d067df)
- [blocked_clients.http_header](data-sources--http_loadbalancer--reference--group-010.md#canonical-5c1b98f99626888192e3ef6181bb0fc0f9936e83cdbf18b745d797f56ddcd20a)
- [blocked_clients.metadata](data-sources--http_loadbalancer--reference--group-010.md#canonical-7a2529e60d610606409d0803a357a1d16a60f79b878f30a3e5a8079b092eb0c2)
- [blocked_clients.skip_processing](data-sources--http_loadbalancer--reference--group-010.md#canonical-9965e647eba8dd09adeb3bc9b042c06d265c31e8b4a62451ae8bcf397aa29d5d)
- [blocked_clients.waf_skip_processing](data-sources--http_loadbalancer--reference--group-010.md#canonical-5bf6cb533cc29130ca90c16ba7d9a325a1ce3f74038707bc0c16f987dff3b744)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-89b1f160e9c5b67b565c80cebec0b7491d42fbe72567f05d3801d45968d067df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd83a9865d8d0e06d44a1f6eac6aff50ee8715bd34e4aa2d400d287c96bf1a6d"></a>

## blocked_clients.bot_skip_processing — blocked_clients.bot_skip_processing / 1298d8525329 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [blocked_clients](data-sources--http_loadbalancer--reference--group-010.md#canonical-e46710ba001bf8ddabcfccfbb5094215ce20da3a35b0ea48470159cc0e080303)
- blocked_clients.bot_skip_processing

<a id="canonical-f24b9615c45145f3a612a5087b68e341b91fd1f4249758e90d294752f6e11949"></a>

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

<a id="canonical-bbe9d2912972a6ef5f6892e65fbb729a601ae4efc5de34fd3c9dba3021e00f01"></a>

## Direct properties — blocked_clients.bot_skip_processing / 1298d8525329 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fc87496476b8543a2c99c1cffe2a253ed4d8f10d2fedd09545d6d5b76a9c6137"></a>

## Next pages — blocked_clients.bot_skip_processing / 1298d8525329 / 4

- [blocked_clients](data-sources--http_loadbalancer--reference--group-010.md#canonical-e46710ba001bf8ddabcfccfbb5094215ce20da3a35b0ea48470159cc0e080303)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5c1b98f99626888192e3ef6181bb0fc0f9936e83cdbf18b745d797f56ddcd20a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b267282d66f85947e079bfa358f079edb59d0b7c4994266cb0ba9afd91295315"></a>

## blocked_clients.http_header — blocked_clients.http_header / d53f34333c81 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [blocked_clients](data-sources--http_loadbalancer--reference--group-010.md#canonical-e46710ba001bf8ddabcfccfbb5094215ce20da3a35b0ea48470159cc0e080303)
- blocked_clients.http_header

<a id="canonical-e65587fd13474e87298295e37be57a4f52d246c9ea0b36c4a2b1b8e71d68f08c"></a>

Type: `"single"`. Computed.

Configuration parameter for http header.

Upstream description:

Request header name and value pairs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ac5c8e43ca0febc36c8ebccf088128f2905ac0bc7a8420c5e58aeeab7b81e0ab"></a>

## Direct properties — blocked_clients.http_header / d53f34333c81 / 3

- [headers](data-sources--http_loadbalancer--reference--group-010.md#canonical-23e46ab75968bf9cc615a560f4bebd1caafcd6f28f44aa482712640a4e2595e1): complete subsection reference.

<a id="canonical-5b17800c5781fdfbdbd5641c57a97bb944b11946780061a000110f110707b396"></a>

## Next pages — blocked_clients.http_header / d53f34333c81 / 4

- [blocked_clients.http_header.headers](data-sources--http_loadbalancer--reference--group-010.md#canonical-23e46ab75968bf9cc615a560f4bebd1caafcd6f28f44aa482712640a4e2595e1)
- [blocked_clients](data-sources--http_loadbalancer--reference--group-010.md#canonical-e46710ba001bf8ddabcfccfbb5094215ce20da3a35b0ea48470159cc0e080303)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-23e46ab75968bf9cc615a560f4bebd1caafcd6f28f44aa482712640a4e2595e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bfb2923472e09c5ebef51dd48a5eff447a48e3a04e47452a4c85f34bf79f337d"></a>

## blocked_clients.http_header.headers — blocked_clients.http_header.headers / 995d4a439aea / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [blocked_clients](data-sources--http_loadbalancer--reference--group-010.md#canonical-e46710ba001bf8ddabcfccfbb5094215ce20da3a35b0ea48470159cc0e080303)
- [blocked_clients.http_header](data-sources--http_loadbalancer--reference--group-010.md#canonical-5c1b98f99626888192e3ef6181bb0fc0f9936e83cdbf18b745d797f56ddcd20a)
- blocked_clients.http_header.headers

<a id="canonical-e50293440dfa189f680ef3b1c864c90450ae5aa73afb0d456c5616232f0b30bf"></a>

Type: `"list"`. Computed.

List of HTTP header name and value pairs.

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
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-99814d0c9a9c87baa94e3e85aa96a0da0b3270648363748232941056ce2e1bb1"></a>

## Direct properties — blocked_clients.http_header.headers / 995d4a439aea / 3

<a id="canonical-1d36a927b4e16887cb977245420c8d527b88ba38b320a815455c8525199c452b"></a>

<a id="canonical-9384e3ef7fa0a981c482ff1cb01632bf53015b27a93b47ed197ffeeac1489b22"></a>

## exact property — blocked_clients.http_header.headers / 995d4a439aea / 4

Type: `"string"`. Computed.

Exclusive with \[presence regex\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regex\] Header value to match exactly.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-ba610d2b01145e732e1076f143364cbd95d34a5c35f331ee8992625fe64dfa16"></a>

<a id="canonical-b63d95769319caf8de34bcb1c402cdc389087389fdb9e8ccf1d3d4c93af57ed4"></a>

## invert_match property — blocked_clients.http_header.headers / 995d4a439aea / 5

Type: `"bool"`. Computed.

Invert the result of the match to detect missing header or non-matching value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-42af5ddaed2f92d8dd16eeb7484de88d92297231a77001f89c73c1ee4c8a8c63"></a>

<a id="canonical-47b754ebdfaf4fe9aef69a4f84996d6e6f73226112f6caa241ffc26b747ca9ad"></a>

## name property — blocked_clients.http_header.headers / 995d4a439aea / 6

Type: `"string"`. Computed.

Name. Name of the header.

Upstream description:

Name of the header.

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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-57a13547b34a563d4a3a3bd6d6213ab60074254edd3b0bdf7de979da9a786492"></a>

<a id="canonical-11bd0905aedd9bfd1817f4c296569dc79dc11ed9b3811e5055edaa25e8a88523"></a>

## presence property — blocked_clients.http_header.headers / 995d4a439aea / 7

Type: `"bool"`. Computed.

Exclusive with \[exact regex\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regex\] If true, check for presence of header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-6c81c357d9d2bc52e4cb5db117affa647ba97d2a6ee207d51a0ba1c4904b3283"></a>

<a id="canonical-489c385086d35431ba88b03a046c1d38c117fa1356b7d50cfc879e81409f5da0"></a>

## regex property — blocked_clients.http_header.headers / 995d4a439aea / 8

Type: `"string"`. Computed.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-d3191151e1a38a94662c8933b1a02806acf499f94d4f30a160494db187e4676f"></a>

## Next pages — blocked_clients.http_header.headers / 995d4a439aea / 9

- [blocked_clients.http_header](data-sources--http_loadbalancer--reference--group-010.md#canonical-5c1b98f99626888192e3ef6181bb0fc0f9936e83cdbf18b745d797f56ddcd20a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7a2529e60d610606409d0803a357a1d16a60f79b878f30a3e5a8079b092eb0c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-650c79dce14b474918e3790e5b4bce0861ab8fb000cd5df37613eabb254761a7"></a>

## blocked_clients.metadata — blocked_clients.metadata / d93a80c04469 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [blocked_clients](data-sources--http_loadbalancer--reference--group-010.md#canonical-e46710ba001bf8ddabcfccfbb5094215ce20da3a35b0ea48470159cc0e080303)
- blocked_clients.metadata

<a id="canonical-a65cec684306ce35744c393838d2e6a66c652c2f001cf04e2489e3a79f3ada81"></a>

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

<a id="canonical-6f07d935d219495e4a889b1aeda6f50996c1f100de5150cc788fa228278916ba"></a>

## Direct properties — blocked_clients.metadata / d93a80c04469 / 3

<a id="canonical-ad41b26afd878b26f0a24559b691e7cd9d9f4c1a06c3260ac6cc266dae51f1e3"></a>

<a id="canonical-f3a34edb12fc677da29dfa52003d967742bcdc1a5758318531d5319581436f6c"></a>

## description_spec property — blocked_clients.metadata / d93a80c04469 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-e0d1df2745e8830220fb978aa545a70b9ff195738492ee19b0b10645dfb966e0"></a>

<a id="canonical-e803cab6c205b23e4fe25caa2fc85b3b54f23b629f4e18fd8dbf2a5408cd11ed"></a>

## name property — blocked_clients.metadata / d93a80c04469 / 5

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

<a id="canonical-1ca6127b618eb491766cc68eba6d0ae0a42918533359fca5a98d490d4d8be460"></a>

## Next pages — blocked_clients.metadata / d93a80c04469 / 6

- [blocked_clients](data-sources--http_loadbalancer--reference--group-010.md#canonical-e46710ba001bf8ddabcfccfbb5094215ce20da3a35b0ea48470159cc0e080303)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9965e647eba8dd09adeb3bc9b042c06d265c31e8b4a62451ae8bcf397aa29d5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee9f516b132a5f405f04645f111e9b1e55558ec1d65833e724a7f7d51d976e37"></a>

## blocked_clients.skip_processing — blocked_clients.skip_processing / 60c4d183cf50 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [blocked_clients](data-sources--http_loadbalancer--reference--group-010.md#canonical-e46710ba001bf8ddabcfccfbb5094215ce20da3a35b0ea48470159cc0e080303)
- blocked_clients.skip_processing

<a id="canonical-59483d241a59ca7956944337a5715e12c28d9538fe6003248a6d2a3931373748"></a>

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

<a id="canonical-daf87d5ffb2ee68a5a1a1a54ad3bc641a392b968a8723a7b79a3a35b95df570a"></a>

## Direct properties — blocked_clients.skip_processing / 60c4d183cf50 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d23f22181d5ac9c05f979bc1598f979cc0f4a43b3a850f3b0ccc351d191fb8ea"></a>

## Next pages — blocked_clients.skip_processing / 60c4d183cf50 / 4

- [blocked_clients](data-sources--http_loadbalancer--reference--group-010.md#canonical-e46710ba001bf8ddabcfccfbb5094215ce20da3a35b0ea48470159cc0e080303)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5bf6cb533cc29130ca90c16ba7d9a325a1ce3f74038707bc0c16f987dff3b744"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cfabf8e82a5f37767b92139e098640bb4ec127070714cdaa3f8186d03b723532"></a>

## blocked_clients.waf_skip_processing — blocked_clients.waf_skip_processing / efc7e149a8df / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [blocked_clients](data-sources--http_loadbalancer--reference--group-010.md#canonical-e46710ba001bf8ddabcfccfbb5094215ce20da3a35b0ea48470159cc0e080303)
- blocked_clients.waf_skip_processing

<a id="canonical-bebcbbba2aa51aab7362360b07d160a2e376a85f0403011678dc8155b6580766"></a>

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

<a id="canonical-292d755f01b2882d7692256e864670282c2f349529f67c07aecb78fb24301f7e"></a>

## Direct properties — blocked_clients.waf_skip_processing / efc7e149a8df / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1aa4ba2d09b036d615041883351222e9564772f434d3b772bf4abe2ce1f68068"></a>

## Next pages — blocked_clients.waf_skip_processing / efc7e149a8df / 4

- [blocked_clients](data-sources--http_loadbalancer--reference--group-010.md#canonical-e46710ba001bf8ddabcfccfbb5094215ce20da3a35b0ea48470159cc0e080303)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c36f14b4c2b71759556cadaec862ae97a029b5a0eddaa0a3eef23edca51bc4e"></a>

## bot_defense — bot_defense / d71c753e7013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- bot_defense

<a id="canonical-5f2d1d977ad1ab886713468788b4d5e385b74d3258afde91f776605843253eeb"></a>

Type: `"single"`. Computed.

\[OneOf: bot\_defense, bot\_defense\_advanced\_protection, disable\_bot\_defense; Default:
disable\_bot\_defense\] Defines various configuration OPTIONS for Bot Defense Policy.

Upstream description:

This defines various configuration OPTIONS for Bot Defense Policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cors_support_choice": "[\"disable_cors_support\",\"enable_cors_support\"]"
}
```

OneOf alternatives in this subsection:

- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-5f2d1d977ad1ab886713468788b4d5e385b74d3258afde91f776605843253eeb)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-012.md#canonical-bebd7c928bc283be61d98e79034209eca0b599d7c29f0a2441423024c3414c13)
- [disable_bot_defense](data-sources--http_loadbalancer--reference--group-017.md#canonical-5c061bffc3b08eafa7c1378220ec071484387b114a62bbbf0083b591f4310159)

Select alternatives according to the provider validators above.

<a id="canonical-4be63a7193dfb5096a6978f2007e585e00679264bc2c1081e931e305514b5437"></a>

## Direct properties — bot_defense / d71c753e7013 / 3

- [disable_cors_support](data-sources--http_loadbalancer--reference--group-010.md#canonical-5e7887a7afd1ea761b5020a7930e1d2266a8c0fb10ceec324217e98c50e61a96): complete subsection reference.

- [enable_cors_support](data-sources--http_loadbalancer--reference--group-010.md#canonical-1913dffe822a76a34da5aa9366fa1e537b778df43ad314d20d94a6359c4eaf7d): complete subsection reference.

- [policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a): complete subsection reference.

<a id="canonical-31e0037366b76c851f87634f0dfebec95f131de620597e4aebc345d75427aa9f"></a>

<a id="canonical-6e59b3214ee9ddb02c9d556ee9ce188e3e4315371742421ddf7ff3b74b2bfac2"></a>

## regional_endpoint property — bot_defense / d71c753e7013 / 4

Type: `"string"`. Computed.

\[Enum: AUTO|US|EU|ASIA\] Defines a selection for Bot Defense region - AUTO: AUTO Automatic
selection based on client IP address - US: US US region - EU: EU European Union region - ASIA: ASIA
Asia region. Possible values are \`AUTO\`, \`US\`, \`EU\`, \`ASIA\`. Defaults to \`AUTO\`.

Upstream description:

Defines a selection for Bot Defense region

&#8203;- AUTO: AUTO

Automatic selection based on client IP address &#8203;- US: US

US region &#8203;- EU: EU

European Union region &#8203;- ASIA: ASIA

Asia region.

Receipt-pinned upstream constraints:

```json
{
  "default": "AUTO",
  "enum": [
    "AUTO",
    "US",
    "EU",
    "ASIA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5cf168fa3f81f64afd998aac5e0331e462253a692fa25a3e1c345f2a6de3580e"></a>

<a id="canonical-504d8b3b1830cb3807d1fc679984caf1e78f8a88441713be280fff1dac633c2b"></a>

## timeout property — bot_defense / d71c753e7013 / 5

Type: `"number"`. Computed.

The timeout for the inference check, in milliseconds.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-06a983828df3e13087d04550bea76fe01b07543cc6ddc988358d02390ba52edf"></a>

## Next pages — bot_defense / d71c753e7013 / 6

- [bot_defense.disable_cors_support](data-sources--http_loadbalancer--reference--group-010.md#canonical-5e7887a7afd1ea761b5020a7930e1d2266a8c0fb10ceec324217e98c50e61a96)
- [bot_defense.enable_cors_support](data-sources--http_loadbalancer--reference--group-010.md#canonical-1913dffe822a76a34da5aa9366fa1e537b778df43ad314d20d94a6359c4eaf7d)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5e7887a7afd1ea761b5020a7930e1d2266a8c0fb10ceec324217e98c50e61a96"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c999fe10edf932f9f1e7ed893ffff0773d0893771ed16776c74114fa47306dcf"></a>

## bot_defense.disable_cors_support — bot_defense.disable_cors_support / ba573c7de1e6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- bot_defense.disable_cors_support

<a id="canonical-d9bb1a866411165b44093091f4689d41b93e2e62fb809fa47ce570d735444d06"></a>

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

<a id="canonical-2dbfb245bc678c4be16efafebbb1800d2c0d49ed2ef6cb710ffb5b9b06d5eae3"></a>

## Direct properties — bot_defense.disable_cors_support / ba573c7de1e6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-34e55b3faf26f869100d040939dbdcce63e3a8028bcc6a5669980d23018cd7a2"></a>

## Next pages — bot_defense.disable_cors_support / ba573c7de1e6 / 4

- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1913dffe822a76a34da5aa9366fa1e537b778df43ad314d20d94a6359c4eaf7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09e05a737829d1b006b8ba38a83e0125dc8a875a938ac8716cea2cda6b4dcfef"></a>

## bot_defense.enable_cors_support — bot_defense.enable_cors_support / 046025b21cce / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- bot_defense.enable_cors_support

<a id="canonical-9eddb8e655900c5ae5f6b5f45627997ccdcfc72b091f7fab80ef3614dbf361ae"></a>

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

<a id="canonical-fd88a6c273657b2f30654f857a1ce3d17ae4b3b3e66bde64c5d476949bcefa1b"></a>

## Direct properties — bot_defense.enable_cors_support / 046025b21cce / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-884e59a86ff8aaa89b74e3f0b83866a56b5c2dfaf2ddf9de92faf0f94d0ae6c1"></a>

## Next pages — bot_defense.enable_cors_support / 046025b21cce / 4

- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0668dd4a68a891e41e9a76bd8fbdb5d6a5f8be1f21d8d9dde36b7cd1c8585e15"></a>

## bot_defense.policy — bot_defense.policy / c885166610b5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- bot_defense.policy

<a id="canonical-4d21e489fdab799856327e3797423174848288892509b8ed956de305f7072af0"></a>

Type: `"single"`. Computed.

Defines various configuration OPTIONS for Bot Defense policy.

Upstream description:

This defines various configuration OPTIONS for Bot Defense policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

<a id="canonical-532c627258f7821ebdd5e6ce1ed22a60c8d74252cf620b87a34d19620d394350"></a>

## Direct properties — bot_defense.policy / c885166610b5 / 3

- [disable_js_insert](data-sources--http_loadbalancer--reference--group-010.md#canonical-1569b0a569c2047dcc09508bc254387841afedb0ecdaee3049d24982915105c5): complete subsection reference.

- [disable_mobile_sdk](data-sources--http_loadbalancer--reference--group-010.md#canonical-1a3ef7ba25509045acef40126a6ad7e234f13d478ea22f114acca9660130201a): complete subsection reference.

<a id="canonical-51eb3c56789fae7646506047218deeb6afa20de31b42b805ca0ca5542549ec41"></a>

<a id="canonical-5a217e0daf2a5863fedb86b964d0604f2ba2d8bf69b72a3e0fadece9e4a2ee1b"></a>

## javascript_mode property — bot_defense.policy / c885166610b5 / 4

Type: `"string"`. Computed.

\[Enum: ASYNC\_JS\_NO\_CACHING|ASYNC\_JS\_CACHING|SYNC\_JS\_NO\_CACHING|SYNC\_JS\_CACHING\] Web
Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is cacheable Bot Defense JavaScript for telemetry collection is requested.. Possible values
are \`ASYNC\_JS\_NO\_CACHING\`, \`ASYNC\_JS\_CACHING\`, \`SYNC\_JS\_NO\_CACHING\`,
\`SYNC\_JS\_CACHING\`. Defaults to \`ASYNC\_JS\_NO\_CACHING\`.

Upstream description:

Web Client JavaScript Mode.

Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is non-cacheable
Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is cacheable Bot
Defense JavaScript for telemetry collection is requested synchronously, and it is non-cacheable Bot
Defense JavaScript for telemetry collection is requested synchronously, and it is cacheable.

Receipt-pinned upstream constraints:

```json
{
  "default": "ASYNC_JS_NO_CACHING",
  "enum": [
    "ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d3cacdfcb488cbdd8aeec506883cfe201331187732eeabd20ba1e164c54f21e7"></a>

<a id="canonical-e3b47ef785f8418067e48ce7119f55ca6b700866d0e1f33a8852a3fa6b3b8c65"></a>

## js_download_path property — bot_defense.policy / c885166610b5 / 5

Type: `"string"`. Computed.

Customize Bot Defense Client JavaScript path. If not specified, default

Upstream description:

Customize Bot Defense Client JavaScript path. If not specified, default \`/common.js\`

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
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

- [js_insert_all_pages](data-sources--http_loadbalancer--reference--group-010.md#canonical-43ca6eb1b26917902545de01bb8b1c982bcba863e1175cce12fdfd73a398d05a): complete subsection reference.

- [js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-010.md#canonical-83900217efc7117c4944aca72cdc833ac183120713a05e2c1689bdbf4f0c2189): complete subsection reference.

- [js_insertion_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-d6917238385568f63c3ce0b1391a716919e47c80879e8015a1a3c9796f033376): complete subsection reference.

- [mobile_sdk_config](data-sources--http_loadbalancer--reference--group-011.md#canonical-4811616a633e7ef2d289eaa66fe11c07419691ff5b67ab20f15316ccba1daa6b): complete subsection reference.

- [protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e): complete subsection reference.

<a id="canonical-9a3e4e0972cbc8147997259ba99a5403aae1ad98f4a74b3f6046bdbc72a28663"></a>

## Next pages — bot_defense.policy / c885166610b5 / 6

- [bot_defense.policy.disable_js_insert](data-sources--http_loadbalancer--reference--group-010.md#canonical-1569b0a569c2047dcc09508bc254387841afedb0ecdaee3049d24982915105c5)
- [bot_defense.policy.disable_mobile_sdk](data-sources--http_loadbalancer--reference--group-010.md#canonical-1a3ef7ba25509045acef40126a6ad7e234f13d478ea22f114acca9660130201a)
- [bot_defense.policy.js_insert_all_pages](data-sources--http_loadbalancer--reference--group-010.md#canonical-43ca6eb1b26917902545de01bb8b1c982bcba863e1175cce12fdfd73a398d05a)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-010.md#canonical-83900217efc7117c4944aca72cdc833ac183120713a05e2c1689bdbf4f0c2189)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-d6917238385568f63c3ce0b1391a716919e47c80879e8015a1a3c9796f033376)
- [bot_defense.policy.mobile_sdk_config](data-sources--http_loadbalancer--reference--group-011.md#canonical-4811616a633e7ef2d289eaa66fe11c07419691ff5b67ab20f15316ccba1daa6b)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-011.md#canonical-2794a05289a537590b99a108f366aac84c4e0adfa5932eb3d357c45ad5031b0e)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1569b0a569c2047dcc09508bc254387841afedb0ecdaee3049d24982915105c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b8214c469395d7434d9e4a248cab4d2bb1eb7f1d6c7b0d447192a576352d761"></a>

## bot_defense.policy.disable_js_insert — bot_defense.policy.disable_js_insert / b8a2001a6e89 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- bot_defense.policy.disable_js_insert

<a id="canonical-98fcbce419cdf8831887cbc969af5743dc21686cd1b8696d71e1d62e0d0fc6a0"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable js insert.

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

<a id="canonical-b6272e65194639b556d5886da8abcc888a52cd02d5cc431728c29a6d6a95d07d"></a>

## Direct properties — bot_defense.policy.disable_js_insert / b8a2001a6e89 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e8782ab10437c1a52eb0dfb2083fd873a332e8756d649ffd1222c40713b57079"></a>

## Next pages — bot_defense.policy.disable_js_insert / b8a2001a6e89 / 4

- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1a3ef7ba25509045acef40126a6ad7e234f13d478ea22f114acca9660130201a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a23028803015d6cf21d3a643567487193acd9a7647b0f3d2943a03835e8eadb3"></a>

## bot_defense.policy.disable_mobile_sdk — bot_defense.policy.disable_mobile_sdk / 6e8f4475ed45 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- bot_defense.policy.disable_mobile_sdk

<a id="canonical-89028d73ae4ca225278cf0afe776da72024352fcf2fa17797b4a478bd9d910fd"></a>

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

<a id="canonical-b6a3b0052b0086a2a97ee8c98f807782c5a605a24ce630502171bf026e9c2620"></a>

## Direct properties — bot_defense.policy.disable_mobile_sdk / 6e8f4475ed45 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ce6ce48e4a9a0fcb09d142bfd717fe138d8cbc22b83c9dc15eb1a6409133281f"></a>

## Next pages — bot_defense.policy.disable_mobile_sdk / 6e8f4475ed45 / 4

- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-43ca6eb1b26917902545de01bb8b1c982bcba863e1175cce12fdfd73a398d05a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7285717de0761d9bfabfa56a9a0deaa1b0a4479e9ee721f24b22e8d212e7012"></a>

## bot_defense.policy.js_insert_all_pages — bot_defense.policy.js_insert_all_pages / d7d953a5c222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- bot_defense.policy.js_insert_all_pages

<a id="canonical-24124fd85ac2fe9e7a2b3b79dadba3ae7ce4b428a7b62b24315188259db16d02"></a>

Type: `"single"`. Computed.

Insert Bot Defense JavaScript in all pages.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-cc67599bc562cbbcdf529ec3c008b15cc7eb07f280a6d194615cdf75b996c943"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages / d7d953a5c222 / 3

<a id="canonical-59b039d5e5ad5881854298a1a9c268e96c20d49522ce5ab6bf7ecdc1089547c5"></a>

<a id="canonical-7a0f8a789afda60db2ec3d8310d7cda5285616201649f574b4fc4874e5d54095"></a>

## javascript_location property — bot_defense.policy.js_insert_all_pages / d7d953a5c222 / 4

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-a9d6ff1dbf7cbb22993e965ba2cb19fb70e99b57468a268960da55194245aa83"></a>

## Next pages — bot_defense.policy.js_insert_all_pages / d7d953a5c222 / 5

- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-83900217efc7117c4944aca72cdc833ac183120713a05e2c1689bdbf4f0c2189"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-043a27636eff4837bef4c7a0724d7e19bc8636974883bdc012f8565689d293b8"></a>

## bot_defense.policy.js_insert_all_pages_except — bot_defense.policy.js_insert_all_pages_except / c731b35da937 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- bot_defense.policy.js_insert_all_pages_except

<a id="canonical-b196486b8f556eb2b7e8d98223bf9b60dfed9a04bd3130c016ba3a7e84ccaace"></a>

Type: `"single"`. Computed.

Insert Bot Defense JavaScript in all pages with the exceptions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-33350bd4be4b5f8efd0467ca5821240ac877e0bfcda821d1e4a22480eeb41453"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except / c731b35da937 / 3

- [exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-29e5fea8c7e803746f4a082260ca065911b776b40d881bc1b2dbba5dd8397279): complete subsection reference.

<a id="canonical-a5d46215fff27a47c88e8622fff39d5c0af447dcc97e503ee624c24ba63be278"></a>

<a id="canonical-688c2c0a1ebef6f5bd4257c42786891dd37131a8b5a67037f0dc48c084811f47"></a>

## javascript_location property — bot_defense.policy.js_insert_all_pages_except / c731b35da937 / 4

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-21c9e94bc73874723164080658b22b9e70ef20ac130bb9d957cbee753c63e2bc"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except / c731b35da937 / 5

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-29e5fea8c7e803746f4a082260ca065911b776b40d881bc1b2dbba5dd8397279)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-29e5fea8c7e803746f4a082260ca065911b776b40d881bc1b2dbba5dd8397279"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77af92216b1fe582c2d2dfb2f326f5e4fd5b65e9adc5f410c56eeca164af2b3c"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list — bot_defense.policy.js_insert_all_pages_except.exclude_list / 2067e221ee23 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-010.md#canonical-83900217efc7117c4944aca72cdc833ac183120713a05e2c1689bdbf4f0c2189)
- bot_defense.policy.js_insert_all_pages_except.exclude_list

<a id="canonical-cd8acb98f566bccadb9edf6bff2d6280b5c713653221f52c712bfa1494fcd2ff"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-59b9a7a095588ab94db1902929841828cfce90665b4a8fa029bacbc4f36f28c3"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list / 2067e221ee23 / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-010.md#canonical-0c0288ec4a7052e8950b43467a343ccbe7b6efb44a28d9f5b3dee225e61ed908): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-010.md#canonical-354831a9ce4bf0776f5034488a168a1d4deb101eb445a3ee9eb67705f6467394): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-010.md#canonical-1c29b935a14725bfeb4164ae81bb923493d32b79667d721be49556f0877a54c5): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-010.md#canonical-07c7efebc4ace78c8d8d452b3f0b97f0ed60ca966dba72bc23dfc0963aedabf9): complete subsection reference.

<a id="canonical-d85500d1240140efcb687c4206c87d183e9c21ca25fd8bd33dc2df632e2d1a54"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list / 2067e221ee23 / 4

- [bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain](data-sources--http_loadbalancer--reference--group-010.md#canonical-0c0288ec4a7052e8950b43467a343ccbe7b6efb44a28d9f5b3dee225e61ed908)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.domain](data-sources--http_loadbalancer--reference--group-010.md#canonical-354831a9ce4bf0776f5034488a168a1d4deb101eb445a3ee9eb67705f6467394)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata](data-sources--http_loadbalancer--reference--group-010.md#canonical-1c29b935a14725bfeb4164ae81bb923493d32b79667d721be49556f0877a54c5)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.path](data-sources--http_loadbalancer--reference--group-010.md#canonical-07c7efebc4ace78c8d8d452b3f0b97f0ed60ca966dba72bc23dfc0963aedabf9)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-010.md#canonical-83900217efc7117c4944aca72cdc833ac183120713a05e2c1689bdbf4f0c2189)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0c0288ec4a7052e8950b43467a343ccbe7b6efb44a28d9f5b3dee225e61ed908"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69471a2515e47c0a8869fb8eec84268a767ee54fda9d7c1702b539ae165afa66"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain — bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / cc0286bbb039 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-010.md#canonical-83900217efc7117c4944aca72cdc833ac183120713a05e2c1689bdbf4f0c2189)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-29e5fea8c7e803746f4a082260ca065911b776b40d881bc1b2dbba5dd8397279)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-6c32699585944996e15528d569bee945ef1a0aeec4478a7f261e2626535ac2d8"></a>

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

<a id="canonical-23d4d2e70239d349e0cee7eb2633836e0736af15e3671ac4d9cd7758d8239264"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / cc0286bbb039 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dbb20b97bf65321dca56d28208b695010869538aa728781b56b67a906fb8cd77"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / cc0286bbb039 / 4

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-29e5fea8c7e803746f4a082260ca065911b776b40d881bc1b2dbba5dd8397279)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-354831a9ce4bf0776f5034488a168a1d4deb101eb445a3ee9eb67705f6467394"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-697ecda538168cc862fa1f9c5eda32a74bb8e2db7fd2f13d4e80bcdbe5a42c4f"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.domain — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / f538d6587a7d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-010.md#canonical-83900217efc7117c4944aca72cdc833ac183120713a05e2c1689bdbf4f0c2189)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-29e5fea8c7e803746f4a082260ca065911b776b40d881bc1b2dbba5dd8397279)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-fde595aaa0d19c164707adbe5a5c548859c4fb89bd8e0c6ab68ce0b3aac586eb"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-4bf821fcbab6fd7e928175d5f14ea17fa8b637f97e98298705c2f2653c40de13"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / f538d6587a7d / 3

<a id="canonical-d5eab904d4ef7f6c328614818813984410617c722353a352821c7bfb9c2c1387"></a>

<a id="canonical-fc44bd3b19ede2d554fbae2fdd56825e4bfcc81bbb4fb464f815fd47f4defd19"></a>

## exact_value property — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / f538d6587a7d / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-4fc3070c8b6a13ed0add6b0e3e0b708cf71cf2109448de356e0b39a2fd0b9358"></a>

<a id="canonical-dcdc1dbb1e2390de06902ac9c722362a4a2be4ae55cb8b205a2f2a4519f019f8"></a>

## regex_value property — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / f538d6587a7d / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2837364eff00282137df07b45b47bc7080bae570fc17baa2eac29e55382bee3a"></a>

<a id="canonical-374a43d9bb25fe85834a7a32221e887cc0433c8991104369f84e4274963a2c4c"></a>

## suffix_value property — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / f538d6587a7d / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-455225720ad08401186b989593e1ec2a5f616fe87a84d266f5f30ff78b849bf1"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / f538d6587a7d / 7

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-29e5fea8c7e803746f4a082260ca065911b776b40d881bc1b2dbba5dd8397279)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1c29b935a14725bfeb4164ae81bb923493d32b79667d721be49556f0877a54c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-156381b49cd04db0cba357a0a31fe9688f6e51d71e2566f285cc5dd6376c94cb"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 45b8388a745a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-010.md#canonical-83900217efc7117c4944aca72cdc833ac183120713a05e2c1689bdbf4f0c2189)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-29e5fea8c7e803746f4a082260ca065911b776b40d881bc1b2dbba5dd8397279)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-5aa3e97e75993e0275cee19d9ab1f573f460075d3623e82fe5da888ed625dac9"></a>

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

<a id="canonical-e4971df2e28eb3f6b9c391b4a51b85b959531270675f322f68b540b23fe19097"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 45b8388a745a / 3

<a id="canonical-b3b2afcc960e3a20c20bdfca3c0d4c74370539fdd8093b6dd7fb5487fb270fd4"></a>

<a id="canonical-ddcffa45f33fc9eccbb3fe49670b2affda4bbbc6767e25877fdf962d580da5d5"></a>

## description_spec property — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 45b8388a745a / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0de1bd2ca6055f1e472ec1d7878eaddde3a94956d2a71de3ef8203d6f8b4191f"></a>

<a id="canonical-5185b414517b918350273ee6c979a89a3300c1f956cd4cedba95ce753f06c66a"></a>

## name property — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 45b8388a745a / 5

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

<a id="canonical-b00e7ae60ef215aeb45e5daaaa0f9238afcfd4209d3776c1b5ec378aec9b2d32"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 45b8388a745a / 6

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-29e5fea8c7e803746f4a082260ca065911b776b40d881bc1b2dbba5dd8397279)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-07c7efebc4ace78c8d8d452b3f0b97f0ed60ca966dba72bc23dfc0963aedabf9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34dfcfbc0a6c8cd76aa93dcab696654f51e2067296b1703c503acb507da28ac2"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.path — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 3ca38daff60c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-010.md#canonical-83900217efc7117c4944aca72cdc833ac183120713a05e2c1689bdbf4f0c2189)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-29e5fea8c7e803746f4a082260ca065911b776b40d881bc1b2dbba5dd8397279)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-f89d787814caef70d8e1ed63676f6f35a1d6778ccf9bfbbd3d03aa5908998d14"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-7fd6baeadf8b0069ad2aa579c7d7cc6e95229bb6cd91d903d68db60be18ad664"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 3ca38daff60c / 3

<a id="canonical-aec56b02e64b67955c49ccda54341bf877f33478fbcf325e86ff57d7f8432d19"></a>

<a id="canonical-1c3df1840a7b14bb3c62ad39c608aad17718a95c178a2f1a03cb07de88dcd189"></a>

## path property — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 3ca38daff60c / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
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

<a id="canonical-135b9ee74d10f8c52ef650ec597b5194e3a27997acf15ccca1219ca68ffa63c7"></a>

<a id="canonical-ab7b9a772d688894faa190779a922c7185303253330c3b9302a1c5dd5b7fb69f"></a>

## prefix property — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 3ca38daff60c / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-4fd8f73fec1c70577d365b25371aebe04902275a6b0fcfb3192d5464a42e7e6a"></a>

<a id="canonical-564ca0a753f7fbc1ee4a1726adde28521e48356c2c242a183bd35d69cd929a61"></a>

## regex property — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 3ca38daff60c / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-7faf3de98f8a892b26273a65bb60fe72c6e4b445ac396a751a3ff6e832d04f50"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 3ca38daff60c / 7

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-29e5fea8c7e803746f4a082260ca065911b776b40d881bc1b2dbba5dd8397279)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d6917238385568f63c3ce0b1391a716919e47c80879e8015a1a3c9796f033376"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e14ea14f98b2d9976539f2207d84039a9a426f5e2b9b9ab4a8da6e9253873387"></a>

## bot_defense.policy.js_insertion_rules — bot_defense.policy.js_insertion_rules / 4dd24c4b2918 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- bot_defense.policy.js_insertion_rules

<a id="canonical-f38fe2b8d73f124b9201d1228c542a7df03f12823537e03977a9d5ab02778531"></a>

Type: `"single"`. Computed.

Defines custom JavaScript insertion rules for Bot Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Bot Defense Policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c0b0b39a3fc4c8b91cb03732723a454a14a75e0c2b59a19232371546cc6ffa9c"></a>

## Direct properties — bot_defense.policy.js_insertion_rules / 4dd24c4b2918 / 3

- [exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-6b93e395e7051a1e9c1b02d58354fb5deb5c1d6242a97feebb6bc91dd631484c): complete subsection reference.

- [rules](data-sources--http_loadbalancer--reference--group-011.md#canonical-3ff4e8f7cf7a4c310f885f64a358f571d87010f0ef1dd74d583490db0c2381d2): complete subsection reference.

<a id="canonical-8c4009b89c4008f2a465c67d3534d865da93bcd711f1bfbaafa705266b6a01ff"></a>

## Next pages — bot_defense.policy.js_insertion_rules / 4dd24c4b2918 / 4

- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-6b93e395e7051a1e9c1b02d58354fb5deb5c1d6242a97feebb6bc91dd631484c)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-011.md#canonical-3ff4e8f7cf7a4c310f885f64a358f571d87010f0ef1dd74d583490db0c2381d2)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6b93e395e7051a1e9c1b02d58354fb5deb5c1d6242a97feebb6bc91dd631484c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9097c5326a0ea51361429534dbb36ce587f703a2f193f6d3a4652bd9a334dee"></a>

## bot_defense.policy.js_insertion_rules.exclude_list — bot_defense.policy.js_insertion_rules.exclude_list / c9ef7c1930bf / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-d6917238385568f63c3ce0b1391a716919e47c80879e8015a1a3c9796f033376)
- bot_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-4b9d247e95e24e50d19d1657b09d53630ebe4e374b4e3320a09ebf0c2d05d5bd"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f1772b9e66d9edbff76cc3fc36dac1fe409c7207c378b0648807edadfc4ae1a6"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list / c9ef7c1930bf / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-010.md#canonical-31f2beefa506c5a753d7dbf9c61150be5bfe50878ddc2a740004dbcd426fe8f7): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-010.md#canonical-7bd7f446a00d9705f0020282d580ab74da62c4ec56afa4597024f88b8f0d2f65): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-010.md#canonical-5546adc417549d96bbc6d6a02af932e28a00e713b6fbe9d7ab975346c5029478): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-010.md#canonical-83178fff32f090b704e36c3f78666d6338d509843a41e2e818ef491581e6b124): complete subsection reference.

<a id="canonical-bef60671e9a08c0ed9e0ffeb84eafe8d254e0d55f9734532df48aa6f0283d3b7"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list / c9ef7c1930bf / 4

- [bot_defense.policy.js_insertion_rules.exclude_list.any_domain](data-sources--http_loadbalancer--reference--group-010.md#canonical-31f2beefa506c5a753d7dbf9c61150be5bfe50878ddc2a740004dbcd426fe8f7)
- [bot_defense.policy.js_insertion_rules.exclude_list.domain](data-sources--http_loadbalancer--reference--group-010.md#canonical-7bd7f446a00d9705f0020282d580ab74da62c4ec56afa4597024f88b8f0d2f65)
- [bot_defense.policy.js_insertion_rules.exclude_list.metadata](data-sources--http_loadbalancer--reference--group-010.md#canonical-5546adc417549d96bbc6d6a02af932e28a00e713b6fbe9d7ab975346c5029478)
- [bot_defense.policy.js_insertion_rules.exclude_list.path](data-sources--http_loadbalancer--reference--group-010.md#canonical-83178fff32f090b704e36c3f78666d6338d509843a41e2e818ef491581e6b124)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-d6917238385568f63c3ce0b1391a716919e47c80879e8015a1a3c9796f033376)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-31f2beefa506c5a753d7dbf9c61150be5bfe50878ddc2a740004dbcd426fe8f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c87cfc8fc52fc927b2fe5a7ea78b9c6bb6442658f1a37bff93a54e596d82d9f1"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.any_domain — bot_defense.policy.js_insertion_rules.exclude_list.any_domain / 1b5ca748dc9c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-d6917238385568f63c3ce0b1391a716919e47c80879e8015a1a3c9796f033376)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-6b93e395e7051a1e9c1b02d58354fb5deb5c1d6242a97feebb6bc91dd631484c)
- bot_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-9ca5403ddd1a541ab9ed1c3e6db541a92aa1b4aa690f7ec92040889c9b924759"></a>

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

<a id="canonical-9627058cf26450a84a03db1e1ff01caebcaf8ba87fb4da3c33893be7c5ffa3a6"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list.any_domain / 1b5ca748dc9c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e04486ce75a5e386eb4bc3e0fa8cb6d61f72cf74cb0a9a0eb32f84e95cdf3359"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list.any_domain / 1b5ca748dc9c / 4

- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-6b93e395e7051a1e9c1b02d58354fb5deb5c1d6242a97feebb6bc91dd631484c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7bd7f446a00d9705f0020282d580ab74da62c4ec56afa4597024f88b8f0d2f65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2ee22ea27a9fe964ed58c4d04482da904c530fd5741ec33061287f0352097bb"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.domain — bot_defense.policy.js_insertion_rules.exclude_list.domain / 44dc0612cff7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-d6917238385568f63c3ce0b1391a716919e47c80879e8015a1a3c9796f033376)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-6b93e395e7051a1e9c1b02d58354fb5deb5c1d6242a97feebb6bc91dd631484c)
- bot_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-93971ee9a7522f0664e4130709657807767cc0a102c6df5dfb86c4d1023abaaf"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-89439670b6d55070b2407062e77eb47f17330a5e0fa029f73625d70648076d81"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list.domain / 44dc0612cff7 / 3

<a id="canonical-05409c15df43535476745dc69ddc3406cbd689e6d464c99fbe165a8675ce99d0"></a>

<a id="canonical-63d401c1d39f64edf426d182e0d611a6b3589a13ee56f385fe94b90c7386e1c6"></a>

## exact_value property — bot_defense.policy.js_insertion_rules.exclude_list.domain / 44dc0612cff7 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-c956b80012ce27b6c05a82a12265a3bed742890c01208a7a5a769961333f8b56"></a>

<a id="canonical-43e7d15096213ed4b9782d1fe47820b33e2c939df856aa909d132a52283f9884"></a>

## regex_value property — bot_defense.policy.js_insertion_rules.exclude_list.domain / 44dc0612cff7 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-35968a94ccfd90070d33dda17914f4f58398e77b9e1440e2bad94b0f7cb4df3e"></a>

<a id="canonical-fb42420d562d0360ecfdf1f65fe672df9d83eca5ec8e3e23afe81a23dcb01a2a"></a>

## suffix_value property — bot_defense.policy.js_insertion_rules.exclude_list.domain / 44dc0612cff7 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-9abcf3dc134902035e8121254130cdf4f9d7aa3ac0fff343db3ffa6589ec8e6e"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list.domain / 44dc0612cff7 / 7

- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-6b93e395e7051a1e9c1b02d58354fb5deb5c1d6242a97feebb6bc91dd631484c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5546adc417549d96bbc6d6a02af932e28a00e713b6fbe9d7ab975346c5029478"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eeff49a3bd48c1883833b4afeb07c12964a76bbdee9592dc339a6bb03bd736b1"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.metadata — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 790819e7e9fb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [bot_defense](data-sources--http_loadbalancer--reference--group-010.md#canonical-99d3520c3439bc5ffc063e6729ac65ba2f27abf1fc5ec5792c7fbbf01c1cbc90)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-010.md#canonical-4245edabcce8eb95c680417cae7324c32f35ad206f693b6686e51f0c66a7f87a)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-d6917238385568f63c3ce0b1391a716919e47c80879e8015a1a3c9796f033376)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-6b93e395e7051a1e9c1b02d58354fb5deb5c1d6242a97feebb6bc91dd631484c)
- bot_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-a6bb6e2972101cb3c52b10f982d1e8af9507839d6e8753117e21be475f508920"></a>

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

<a id="canonical-f576d97e52fdd5a36e607f4eb93d5e234cb880d4a7209dbf4155e9e9d2a7ac0c"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 790819e7e9fb / 3

<a id="canonical-54bfba4c5c89f691518e5d717538111a605e7b4669ad8df25854238e351b35a9"></a>

<a id="canonical-f27db4eb9d899420a91b6ed9eb0faf2dab953cce5de71b52fda8d991d8d10f1f"></a>

## description_spec property — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 790819e7e9fb / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-7c5ec1fd3ef9179e79bd77178c4e0a1853a4b7b896f45a44e2f0688b47897c2a"></a>

<a id="canonical-d0cd8cdcd91abd8f53a079ad7f0cd7bf2b2d7c2502a3e3213852bd860e8b9018"></a>

## name property — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 790819e7e9fb / 5

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

<a id="canonical-ea7be76c20f39e651737778165ba5b6e6abba94daf70141e7b0292d71d0f1f7c"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 790819e7e9fb / 6

- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-6b93e395e7051a1e9c1b02d58354fb5deb5c1d6242a97feebb6bc91dd631484c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-83178fff32f090b704e36c3f78666d6338d509843a41e2e818ef491581e6b124"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
