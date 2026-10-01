---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-9314b628bf1cd27af121acf6948332819cd353718fa578d6b977533105797a4a"></a>

## routes.simple_route.query_params.retain_all_params — routes.simple_route.query_params.retain_all_params / 909cd238fb01 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.query_params](resources--http_loadbalancer--reference--group-025.md#canonical-79028ea11feedc89453fd837580235418e7e524944afa89f9e7523949cc38b51)
- routes.simple_route.query_params.retain_all_params

<a id="canonical-532e0aa5b9cb80e311a1f654faa9df8bfed0a23d05a2f379d902d28adac86d9c"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for retain all params.

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
retain_all_params = {}
```

<a id="canonical-b328f7217800bb43d4ef849c39546c6b99f5e6e85db513352f7ffebda38115c2"></a>

## Direct properties — routes.simple_route.query_params.retain_all_params / 909cd238fb01 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3e7becb8be820578518db50280c9636baf8e79cc3bcbb881e2ee0b68564f6fa4"></a>

## Next pages — routes.simple_route.query_params.retain_all_params / 909cd238fb01 / 4

- [routes.simple_route.query_params](resources--http_loadbalancer--reference--group-025.md#canonical-79028ea11feedc89453fd837580235418e7e524944afa89f9e7523949cc38b51)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c47009539eef3bdcdb8d2b7bb2e345b80b42d65bebf56615f0a60ba75fe78afe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9cd8a87d4dd5d41d1ff92cc54cdf16e1b43faf7b89d4152056f77d42d14d233"></a>

## sensitive_data_disclosure_rules — sensitive_data_disclosure_rules / ade4471c4eff / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- sensitive_data_disclosure_rules

<a id="canonical-b1c7ac9794930cfb9e33ae6955946963aa0e462f674b852c03bbe22049971d5d"></a>

Type: `"object"`. single nested block, Optional.

Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API
responses.

Receipt-pinned upstream constraints:

```json
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
sensitive_data_disclosure_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-6c011bad1803e37f0f989da4b481ebf81cba003ca961400d2e5923e557398f9a"></a>

## Direct properties — sensitive_data_disclosure_rules / ade4471c4eff / 3

- [sensitive_data_types_in_response](resources--http_loadbalancer--reference--group-026.md#canonical-8b63c4c54c0db4d252acc3fbde9f25bc9d8c3fec46490680e5e0702fac27e68c): complete subsection reference.

<a id="canonical-c64600b49695da7d3e0139617593e608c4e6817201b32804981d5f8b10ad35bc"></a>

## Next pages — sensitive_data_disclosure_rules / ade4471c4eff / 4

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--reference--group-026.md#canonical-8b63c4c54c0db4d252acc3fbde9f25bc9d8c3fec46490680e5e0702fac27e68c)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8b63c4c54c0db4d252acc3fbde9f25bc9d8c3fec46490680e5e0702fac27e68c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-385c5a079db4da2040d3e89ddd9d19ed8b086843a7c3a431d5babb34a4bfe89d"></a>

## sensitive_data_disclosure_rules.sensitive_data_types_in_response — sensitive_data_disclosure_rules.sensitive_data_types_in_response / 2793fdf5a275 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [sensitive_data_disclosure_rules](resources--http_loadbalancer--reference--group-026.md#canonical-c47009539eef3bdcdb8d2b7bb2e345b80b42d65bebf56615f0a60ba75fe78afe)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response

<a id="canonical-aaa095a33846f39d5cd89abcde27962b383759d2fc39975575c818e98d949f11"></a>

Type: `"object"`. list nested block, Optional.

Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API
responses.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("mask",
    "report")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
sensitive_data_types_in_response {
  # Configure direct properties listed below.
}
```

<a id="canonical-15b2595472cae09913a5c667b659627f1785fbcb958624f74410dab2c7ec8f34"></a>

## Direct properties — sensitive_data_disclosure_rules.sensitive_data_types_in_response / 2793fdf5a275 / 3

- [api_endpoint](resources--http_loadbalancer--reference--group-026.md#canonical-d397ec119500e0a29b8c5a952e31bd6e1105d20636e4deeb5840bfd2e166f684): complete subsection reference.

- [body](resources--http_loadbalancer--reference--group-026.md#canonical-bb646ae42e5c3fac294409a205913f815a74ebee615908ed32d345806dae2885): complete subsection reference.

- [mask](resources--http_loadbalancer--reference--group-026.md#canonical-7931f9884e4ea4ebc5af8d4c75552c077bb7079052ca4deaefa8b203948f1cf3): complete subsection reference.

- [report](resources--http_loadbalancer--reference--group-026.md#canonical-4de659478e3cb21bc734c7a598c9e31bc7b4ebe021e06e6db0f807e0ded43c35): complete subsection reference.

<a id="canonical-7aea36d1f19905bbdd0eb84ed292eb10532baca06530439202bcab280c1ee272"></a>

## Next pages — sensitive_data_disclosure_rules.sensitive_data_types_in_response / 2793fdf5a275 / 4

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint](resources--http_loadbalancer--reference--group-026.md#canonical-d397ec119500e0a29b8c5a952e31bd6e1105d20636e4deeb5840bfd2e166f684)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.body](resources--http_loadbalancer--reference--group-026.md#canonical-bb646ae42e5c3fac294409a205913f815a74ebee615908ed32d345806dae2885)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask](resources--http_loadbalancer--reference--group-026.md#canonical-7931f9884e4ea4ebc5af8d4c75552c077bb7079052ca4deaefa8b203948f1cf3)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.report](resources--http_loadbalancer--reference--group-026.md#canonical-4de659478e3cb21bc734c7a598c9e31bc7b4ebe021e06e6db0f807e0ded43c35)
- [sensitive_data_disclosure_rules](resources--http_loadbalancer--reference--group-026.md#canonical-c47009539eef3bdcdb8d2b7bb2e345b80b42d65bebf56615f0a60ba75fe78afe)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d397ec119500e0a29b8c5a952e31bd6e1105d20636e4deeb5840bfd2e166f684"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-910fe71ab96761e9c670962d420fe03e4f2f7be3647a7d0b334499fd7605bcda"></a>

## sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint — sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint / 4b5c066c397e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [sensitive_data_disclosure_rules](resources--http_loadbalancer--reference--group-026.md#canonical-c47009539eef3bdcdb8d2b7bb2e345b80b42d65bebf56615f0a60ba75fe78afe)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--reference--group-026.md#canonical-8b63c4c54c0db4d252acc3fbde9f25bc9d8c3fec46490680e5e0702fac27e68c)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint

<a id="canonical-eb6a9471b4e1e1984c074b8ca0ea3d80b73de22ab2d1bc7c876cf3de678e8fc6"></a>

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

<a id="canonical-927d2f033e21847b2b16d337ba96189e51f4d36aaee5652131beef405b03180a"></a>

## Direct properties — sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint / 4b5c066c397e / 3

<a id="canonical-b54a51a98695a2e01b6bf4df4539d6a7ca444d142edb63183658d4a93f04fe90"></a>

<a id="canonical-7dbba58b13ead36202827cad9c217c589ff5028f80fe3c8ea1e6c5c593b7e8ea"></a>

## methods property — sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint / 4b5c066c397e / 4

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

<a id="canonical-ae01d0538c39818c01f5f29e345be619905e416e178e1300886c30ce2a4327a4"></a>

<a id="canonical-c178e35a3ec24abc166a931227d740555e43d477c0e44111bc39b5f32d7455e4"></a>

## path property — sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint / 4b5c066c397e / 5

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

<a id="canonical-ad6931d152f2b699906aad1bad1a1f6f5dd7f7d3151ca79e03dd3d97ac5ad593"></a>

## Next pages — sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint / 4b5c066c397e / 6

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--reference--group-026.md#canonical-8b63c4c54c0db4d252acc3fbde9f25bc9d8c3fec46490680e5e0702fac27e68c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-bb646ae42e5c3fac294409a205913f815a74ebee615908ed32d345806dae2885"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4fccac06678dae1e0bacf483292d4c090921f01b87e388a7eaed0a21a0db7ba1"></a>

## sensitive_data_disclosure_rules.sensitive_data_types_in_response.body — sensitive_data_disclosure_rules.sensitive_data_types_in_response.body / ebd46ccf63d5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [sensitive_data_disclosure_rules](resources--http_loadbalancer--reference--group-026.md#canonical-c47009539eef3bdcdb8d2b7bb2e345b80b42d65bebf56615f0a60ba75fe78afe)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--reference--group-026.md#canonical-8b63c4c54c0db4d252acc3fbde9f25bc9d8c3fec46490680e5e0702fac27e68c)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.body

<a id="canonical-adcd613f8f00726f3318dc2c9540e327c0e28321144bdd1b073e1fb01e437f59"></a>

Type: `"object"`. single nested block, Optional.

Body Section Masking OPTIONS. OPTIONS for HTTP Body Masking.

Upstream description:

OPTIONS for HTTP Body Masking.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("fields")}
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
body {
  # Configure direct properties listed below.
}
```

<a id="canonical-17bcbe7d5f6a1384fe0adcadb93253a27bdc5851b79d1642113a78237e379e8f"></a>

## Direct properties — sensitive_data_disclosure_rules.sensitive_data_types_in_response.body / ebd46ccf63d5 / 3

<a id="canonical-01b302947eccafa3fe44bc13286e1174378f00caa294f11187e377f4e3e33069"></a>

<a id="canonical-4b26bda9aac71036b36e68ef7466ae9e32a62c318d5b6701a563df12cc36b5c1"></a>

## fields property — sensitive_data_disclosure_rules.sensitive_data_types_in_response.body / ebd46ccf63d5 / 4

Type: `["list", "string"]`. Optional.

List of JSON Path field values. Use square brackets with an underscore \[\_\] to indicate array
elements (e.g., person.emails\[\_\]). To reference JSON keys that contain spaces, enclose the entire
path in double quotes.

Upstream description:

List of JSON Path field values. Use square brackets with an underscore \[\_\] to indicate array
elements (e.g., person.emails\[\_\]). To reference JSON keys that contain spaces, enclose the entire
path in double quotes. For example: "person.first name".

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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
    "ves.io.schema.rules.repeated.items.string.json_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.json_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-073de911fc397a80e10b28e5b031826d06768b734b7a37a27dedbf763080e032"></a>

## Next pages — sensitive_data_disclosure_rules.sensitive_data_types_in_response.body / ebd46ccf63d5 / 5

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--reference--group-026.md#canonical-8b63c4c54c0db4d252acc3fbde9f25bc9d8c3fec46490680e5e0702fac27e68c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7931f9884e4ea4ebc5af8d4c75552c077bb7079052ca4deaefa8b203948f1cf3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-896c8e13378a7d69e167f6e941823deaac4ff4c8684ba9fc45907ed85e764392"></a>

## sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask — sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask / de005702cf7a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [sensitive_data_disclosure_rules](resources--http_loadbalancer--reference--group-026.md#canonical-c47009539eef3bdcdb8d2b7bb2e345b80b42d65bebf56615f0a60ba75fe78afe)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--reference--group-026.md#canonical-8b63c4c54c0db4d252acc3fbde9f25bc9d8c3fec46490680e5e0702fac27e68c)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask

<a id="canonical-f6779dfde9a1ad3550276794782bf947da1d344867be511965f350c3bdef3fa8"></a>

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
mask = {}
```

<a id="canonical-5f53ca51701dd4dee543ffbee5d48021e92829b0b78e2842c89bfc8e0aa720d9"></a>

## Direct properties — sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask / de005702cf7a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-62cf55d7f1f0930e51f71a86e053c86a022edfb3e29fe5a468c68d604fae7ab3"></a>

## Next pages — sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask / de005702cf7a / 4

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--reference--group-026.md#canonical-8b63c4c54c0db4d252acc3fbde9f25bc9d8c3fec46490680e5e0702fac27e68c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4de659478e3cb21bc734c7a598c9e31bc7b4ebe021e06e6db0f807e0ded43c35"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1b1b396eaa9afc50f0b159ce8e8494664bfb95fc85fbc03787cefb66146fc2f"></a>

## sensitive_data_disclosure_rules.sensitive_data_types_in_response.report — sensitive_data_disclosure_rules.sensitive_data_types_in_response.report / 9e08abe90847 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [sensitive_data_disclosure_rules](resources--http_loadbalancer--reference--group-026.md#canonical-c47009539eef3bdcdb8d2b7bb2e345b80b42d65bebf56615f0a60ba75fe78afe)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--reference--group-026.md#canonical-8b63c4c54c0db4d252acc3fbde9f25bc9d8c3fec46490680e5e0702fac27e68c)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.report

<a id="canonical-d8e16c80a08d51f307ed03a3c7270e48dd7b1b5ca489fa3701f1e55cc432fbec"></a>

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
report = {}
```

<a id="canonical-3b6e768a964d901449eea3bb19bfe6faf75ccf06aec802a04b15de6d0fd9d125"></a>

## Direct properties — sensitive_data_disclosure_rules.sensitive_data_types_in_response.report / 9e08abe90847 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-53fd0bd7b7064fda63a0506c0a8b2f3935f9c88569074784c199ebaab3bb4f6e"></a>

## Next pages — sensitive_data_disclosure_rules.sensitive_data_types_in_response.report / 9e08abe90847 / 4

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--reference--group-026.md#canonical-8b63c4c54c0db4d252acc3fbde9f25bc9d8c3fec46490680e5e0702fac27e68c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f0e4b12924ffd6a5474ec9ccc1e5f3bb10984f997a330e17cd29b09c63f46f08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a452ffc98df59195c9f23f8da388389e03b30d99e88cfb88c98eea05793ada34"></a>

## sensitive_data_policy — sensitive_data_policy / 422951ad8c49 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- sensitive_data_policy

<a id="canonical-bde87177d5899e4b570b68001e43fcceb5df0837c3f0b2e71fe8eeebe618be2a"></a>

Type: `"object"`. single nested block, Optional.

Policy configuration for this feature.

Upstream description:

Settings for data type policy.

Receipt-pinned upstream constraints:

```json
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
sensitive_data_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-cb56254eb6aefe7f9863b4e6a2aa6ea8b846961a2c5547b39758e61dd0d23c2a"></a>

## Direct properties — sensitive_data_policy / 422951ad8c49 / 3

- [sensitive_data_policy_ref](resources--http_loadbalancer--reference--group-026.md#canonical-6b182ed30465b0f806ec37845330919e7f43f2e01e2d9ea0b7dfeda144769217): complete subsection reference.

<a id="canonical-41b1b45cbff51d5953f1e0ff54c709d898a049efbe57ade082d763d94cfb458e"></a>

## Next pages — sensitive_data_policy / 422951ad8c49 / 4

- [sensitive_data_policy.sensitive_data_policy_ref](resources--http_loadbalancer--reference--group-026.md#canonical-6b182ed30465b0f806ec37845330919e7f43f2e01e2d9ea0b7dfeda144769217)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6b182ed30465b0f806ec37845330919e7f43f2e01e2d9ea0b7dfeda144769217"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15e4ba9b92d30fc29682330870c3b205938d4f54ed598111af1627be4c802123"></a>

## sensitive_data_policy.sensitive_data_policy_ref — sensitive_data_policy.sensitive_data_policy_ref / aad20c5aee92 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [sensitive_data_policy](resources--http_loadbalancer--reference--group-026.md#canonical-f0e4b12924ffd6a5474ec9ccc1e5f3bb10984f997a330e17cd29b09c63f46f08)
- sensitive_data_policy.sensitive_data_policy_ref

<a id="canonical-2d083475e3133731a6d7d2f55100443e5e53daae241987b6f36e48469ed3dbe1"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

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

Terraform syntax:

```terraform
sensitive_data_policy_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-d0e3cef6aa1a568304ec3d7b23abc578ad58e969e754eb548ad5c36e3706e873"></a>

## Direct properties — sensitive_data_policy.sensitive_data_policy_ref / aad20c5aee92 / 3

<a id="canonical-efe5465deae4bee7e419e8cf74952bb713f24cd2acac1d52cbcf4ef62ec9bd9c"></a>

<a id="canonical-a180c8ab12f5a147e3a23931d03d9f341522396429a3d5e4badbf9ecc179fe5b"></a>

## name property — sensitive_data_policy.sensitive_data_policy_ref / aad20c5aee92 / 4

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

<a id="canonical-b68e747729f8c42b6505c326941e67417eacc2c53e641cb32461372bc0bd1cb9"></a>

<a id="canonical-2f483a2c1adff44b52830af63ccc0b3c531208b6b5c9f839d5c883c997a7392f"></a>

## namespace property — sensitive_data_policy.sensitive_data_policy_ref / aad20c5aee92 / 5

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

<a id="canonical-00e868b64c93e618980ed4902259ccf9e5fa61a55b12f22627fd6ff295875380"></a>

<a id="canonical-f0e425677bf1e8660a5cf0a1eb76d51790aaa5c21aeb664d5d709de20c404e13"></a>

## tenant property — sensitive_data_policy.sensitive_data_policy_ref / aad20c5aee92 / 6

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

<a id="canonical-870aef72cf1480021e304f25af1b4913e06a532db90cfd25d07d73f66f652597"></a>

## Next pages — sensitive_data_policy.sensitive_data_policy_ref / aad20c5aee92 / 7

- [sensitive_data_policy](resources--http_loadbalancer--reference--group-026.md#canonical-f0e4b12924ffd6a5474ec9ccc1e5f3bb10984f997a330e17cd29b09c63f46f08)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f8e76b68a32f7686ffecdddc1f58ed7947137696e500423d54f0161e3b096a55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a58e0a01ada6ebe5168bad12adb7742708f9291d7020816134874b30c9a7536"></a>

## service_policies_from_namespace — service_policies_from_namespace / 7a79ff78590b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- service_policies_from_namespace

<a id="canonical-dbbad8a7d71d77af9c830a7e918ae84ee7210609e1b8c5474ab44e3750fc8d69"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
service_policies_from_namespace = {}
```

<a id="canonical-a9377a51316dfa9eb6ca003b3d630e342a97b8a08d676aae70c5425ae694e28d"></a>

## Direct properties — service_policies_from_namespace / 7a79ff78590b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a109caba380c96810b2e530277ac2703de5b54abd48638713af4131b7bdcf177"></a>

## Next pages — service_policies_from_namespace / 7a79ff78590b / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7962d3ee77b3c66dba01ad107dbf06631d66f9f118a25473f3c0938cf819bb2c"></a>

## single_lb_app — single_lb_app / 941a01434058 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- single_lb_app

<a id="canonical-61edcd8f7e9848ea5675c000f776f6de16862c029a213391c63a3a015bbcac1f"></a>

Type: `"object"`. single nested block, Optional.

Specific settings for Machine learning analysis on this HTTP LB, independently from other LBs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_discovery",
    "enable_discovery"),
  validators.ConflictingObjectAttributes("disable_malicious_user_detection",
    "enable_malicious_user_detection")}
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
  "x-ves-oneof-field-api_discovery_choice": "[\"disable_discovery\",\"enable_discovery\"]",
  "x-ves-oneof-field-malicious_user_detection_choice": "[\"disable_malicious_user_detection\",\"enable_malicious_user_detection\"]"
}
```

Terraform syntax:

```terraform
single_lb_app {
  # Configure direct properties listed below.
}
```

<a id="canonical-1ab3fdd6f38e89c48c29c87a46b1f9d6c2fe253e17121012fb4cc4b0f66cf0d5"></a>

## Direct properties — single_lb_app / 941a01434058 / 3

- [disable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-42f31549ecf506646bc1cec827cb063d16df290c5df52af5ebe177dfb6d1a033): complete subsection reference.

- [disable_malicious_user_detection](resources--http_loadbalancer--reference--group-026.md#canonical-aa4f9dd6d6b267fd45c9e4e7311ab66dae8f2d1e6384130fc2675fba6814de8f): complete subsection reference.

- [enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f): complete subsection reference.

- [enable_malicious_user_detection](resources--http_loadbalancer--reference--group-026.md#canonical-3b8de8aa2df2e33331cd69c20c0ae1bee2de9b546d49200646a53f2a03708272): complete subsection reference.

<a id="canonical-268773358def78f768012de25db2d52c3b6d443c712d0b3105e9ff0ce9faef2b"></a>

## Next pages — single_lb_app / 941a01434058 / 4

- [single_lb_app.disable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-42f31549ecf506646bc1cec827cb063d16df290c5df52af5ebe177dfb6d1a033)
- [single_lb_app.disable_malicious_user_detection](resources--http_loadbalancer--reference--group-026.md#canonical-aa4f9dd6d6b267fd45c9e4e7311ab66dae8f2d1e6384130fc2675fba6814de8f)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [single_lb_app.enable_malicious_user_detection](resources--http_loadbalancer--reference--group-026.md#canonical-3b8de8aa2df2e33331cd69c20c0ae1bee2de9b546d49200646a53f2a03708272)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-42f31549ecf506646bc1cec827cb063d16df290c5df52af5ebe177dfb6d1a033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e43d2d173589e26771cdc21446d7e805d78c2981b78ae3a346f8c815da73a2b0"></a>

## single_lb_app.disable_discovery — single_lb_app.disable_discovery / 50d65beafa3f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- single_lb_app.disable_discovery

<a id="canonical-061ff2e146b6674d9ccf80cb0b2ae7d2cb11a011ea4e9fb44edcf246deded9a5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable discovery.

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
disable_discovery = {}
```

<a id="canonical-974797c1a579f7803c7f5e919630f372b2956be8ddf93b7e20d3c4b23aa3af32"></a>

## Direct properties — single_lb_app.disable_discovery / 50d65beafa3f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f3ed179ab678ca1e2715de94a83edfb9419cdfebc06fe51a21c019c670dea908"></a>

## Next pages — single_lb_app.disable_discovery / 50d65beafa3f / 4

- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-aa4f9dd6d6b267fd45c9e4e7311ab66dae8f2d1e6384130fc2675fba6814de8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98d88c6287fc9cb9400fef99b905710f26061da812016061209fa5cd0adccd47"></a>

## single_lb_app.disable_malicious_user_detection — single_lb_app.disable_malicious_user_detection / 83a8ca7f1b37 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- single_lb_app.disable_malicious_user_detection

<a id="canonical-98eeef6b1bd5d811d1c5a584177920d5186efb48e8fb9cc8bf0917051404f5ad"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable malicious user detection.

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
disable_malicious_user_detection = {}
```

<a id="canonical-d0c5d65d915dc6880402a0a0ee553382b14e9371c073e92cb9b358c72870db11"></a>

## Direct properties — single_lb_app.disable_malicious_user_detection / 83a8ca7f1b37 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-600a67a4a8e49db199482c6ed128acafdf646786a5fdd08e5b651455b23bcc65"></a>

## Next pages — single_lb_app.disable_malicious_user_detection / 83a8ca7f1b37 / 4

- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c971ffdbabbc1ebb204770cddbbf69581c1dbc57664cf2b9b3f590e884374e8"></a>

## single_lb_app.enable_discovery — single_lb_app.enable_discovery / 6d0a38f9e11e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- single_lb_app.enable_discovery

<a id="canonical-b584a6ba8fec4c47961f69a8dd4fef42610f8ccbfaf380958fad7e5fa35c3c3f"></a>

Type: `"object"`. single nested block, Optional.

Specifies the settings used for API discovery.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_api_auth_discovery",
    "default_api_auth_discovery"),
  validators.ConflictingObjectAttributes("disable_learn_from_redirect_traffic",
    "enable_learn_from_redirect_traffic")}
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
  "x-ves-oneof-field-api_discovery_settings_choice": "[\"custom_api_auth_discovery\",\"default_api_auth_discovery\"]",
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable_learn_from_redirect_traffic\",\"enable_learn_from_redirect_traffic\"]"
}
```

Terraform syntax:

```terraform
enable_discovery {
  # Configure direct properties listed below.
}
```

<a id="canonical-5dbe0720ffadffd395857a9129f485951f26727ce841478e3036a6b8e0a2dfce"></a>

## Direct properties — single_lb_app.enable_discovery / 6d0a38f9e11e / 3

- [api_crawler](resources--http_loadbalancer--reference--group-026.md#canonical-bd895f7d5fc980e7c41cfa7e68ec079cf979376708b050350a2df4257735b3c6): complete subsection reference.

- [api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-026.md#canonical-290d1ecaf61b258bc0f26ba76908c342c20ea187fbdb217fff8d74f7654f3028): complete subsection reference.

- [custom_api_auth_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-0f230849969e44c2f2ac246346b86e14f6a0843e340c1789ac232fd51679ced1): complete subsection reference.

- [default_api_auth_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-76d497446e6c858ad5778d2211bb253f9e943c331cef340d027714372cccd19c): complete subsection reference.

- [disable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-026.md#canonical-5d6606df66372cf747c7cfbd5108a1aeb91ee8aab307276c5c15a5f9a1ae4232): complete subsection reference.

- [discovered_api_settings](resources--http_loadbalancer--reference--group-026.md#canonical-29185b0644f16e70f840996434781ad14fcb85c7631936b2402bc5c5c69e10b5): complete subsection reference.

- [enable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-026.md#canonical-5302cce5c74a0a0cee944f31fb856f01ef1a39fb827527c4f6add5b9a2dc4970): complete subsection reference.

<a id="canonical-1617e64d8283886e1a93d865956402ab9334e93782da1f0361903e4ac6476363"></a>

## Next pages — single_lb_app.enable_discovery / 6d0a38f9e11e / 4

- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-026.md#canonical-bd895f7d5fc980e7c41cfa7e68ec079cf979376708b050350a2df4257735b3c6)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-026.md#canonical-290d1ecaf61b258bc0f26ba76908c342c20ea187fbdb217fff8d74f7654f3028)
- [single_lb_app.enable_discovery.custom_api_auth_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-0f230849969e44c2f2ac246346b86e14f6a0843e340c1789ac232fd51679ced1)
- [single_lb_app.enable_discovery.default_api_auth_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-76d497446e6c858ad5778d2211bb253f9e943c331cef340d027714372cccd19c)
- [single_lb_app.enable_discovery.disable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-026.md#canonical-5d6606df66372cf747c7cfbd5108a1aeb91ee8aab307276c5c15a5f9a1ae4232)
- [single_lb_app.enable_discovery.discovered_api_settings](resources--http_loadbalancer--reference--group-026.md#canonical-29185b0644f16e70f840996434781ad14fcb85c7631936b2402bc5c5c69e10b5)
- [single_lb_app.enable_discovery.enable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-026.md#canonical-5302cce5c74a0a0cee944f31fb856f01ef1a39fb827527c4f6add5b9a2dc4970)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-bd895f7d5fc980e7c41cfa7e68ec079cf979376708b050350a2df4257735b3c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db6feec0a6c6717982ca1d68fb5d922c73fa60fff5616e63cd033f7debd2a21e"></a>

## single_lb_app.enable_discovery.api_crawler — single_lb_app.enable_discovery.api_crawler / 34efd6a65f6c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- single_lb_app.enable_discovery.api_crawler

<a id="canonical-5f5e9ca399db6271b535b4cf2fa125c0abce6640e8375d2e981ef2467d06ebd2"></a>

Type: `"object"`. single nested block, Optional.

API Crawling. API Crawler message.

Upstream description:

API Crawler message.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("api_crawler_config",
    "disable_api_crawler")}
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
  "x-ves-oneof-field-api_crawler": "[\"api_crawler_config\",\"disable_api_crawler\"]"
}
```

Terraform syntax:

```terraform
api_crawler {
  # Configure direct properties listed below.
}
```

<a id="canonical-939ca2ff1815e33010d57d5bedaafe95268f91f55a02636a7ed8b4f7398c8ce3"></a>

## Direct properties — single_lb_app.enable_discovery.api_crawler / 34efd6a65f6c / 3

- [api_crawler_config](resources--http_loadbalancer--reference--group-026.md#canonical-8f6d639c170c0bd13c34d8bb96ad2ba8822f63f5803db5617273886cf8048716): complete subsection reference.

- [disable_api_crawler](resources--http_loadbalancer--reference--group-026.md#canonical-224059c3dc911c405de20583131137ca79cface5ff51b8febd4c7a9dea8d7154): complete subsection reference.

<a id="canonical-edab5d41607eea3b25657c7a83b803fbd9c453b1f9771243bcadbaa9b95bfca6"></a>

## Next pages — single_lb_app.enable_discovery.api_crawler / 34efd6a65f6c / 4

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-026.md#canonical-8f6d639c170c0bd13c34d8bb96ad2ba8822f63f5803db5617273886cf8048716)
- [single_lb_app.enable_discovery.api_crawler.disable_api_crawler](resources--http_loadbalancer--reference--group-026.md#canonical-224059c3dc911c405de20583131137ca79cface5ff51b8febd4c7a9dea8d7154)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8f6d639c170c0bd13c34d8bb96ad2ba8822f63f5803db5617273886cf8048716"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7bf5071df25be365f3052859288d6b1304bc6c6c539fb62dc08151216a3f484"></a>

## single_lb_app.enable_discovery.api_crawler.api_crawler_config — single_lb_app.enable_discovery.api_crawler.api_crawler_config / 71c19d7a33f5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-026.md#canonical-bd895f7d5fc980e7c41cfa7e68ec079cf979376708b050350a2df4257735b3c6)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config

<a id="canonical-9eb68c842e403e989e3576f8961e00cd2de94893c4b415b73d33e69d4ba97625"></a>

Type: `"object"`. single nested block, Optional.

Crawler Configure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains")}
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
api_crawler_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-d74e30aac9a869d3fdfc3103ec9aa28866cd2469ef978b5085ee54057236b5a7"></a>

## Direct properties — single_lb_app.enable_discovery.api_crawler.api_crawler_config / 71c19d7a33f5 / 3

- [domains](resources--http_loadbalancer--reference--group-026.md#canonical-9ba0ce8049ded7ef0dac39c4265b646c82d6c6cfc6c807629b8d87d0ff430903): complete subsection reference.

<a id="canonical-cad96e9d4634c4cf153b4e67282bb99c10015c8ac42baf93aacc0c58ae45fe6a"></a>

## Next pages — single_lb_app.enable_discovery.api_crawler.api_crawler_config / 71c19d7a33f5 / 4

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-026.md#canonical-9ba0ce8049ded7ef0dac39c4265b646c82d6c6cfc6c807629b8d87d0ff430903)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-026.md#canonical-bd895f7d5fc980e7c41cfa7e68ec079cf979376708b050350a2df4257735b3c6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9ba0ce8049ded7ef0dac39c4265b646c82d6c6cfc6c807629b8d87d0ff430903"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd1e23e664e2037595130965b2bf31a1cafd04b6ef3c149a891a9743e3d6cd5d"></a>

## single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains / e75b0d1b90d2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-026.md#canonical-bd895f7d5fc980e7c41cfa7e68ec079cf979376708b050350a2df4257735b3c6)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-026.md#canonical-8f6d639c170c0bd13c34d8bb96ad2ba8822f63f5803db5617273886cf8048716)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains

<a id="canonical-241f5f27939d3117668109e9fbd46f6526503d79a5bd05bb391a5b98967855c9"></a>

Type: `"object"`. list nested block, Optional.

Enter domains and their credentials to allow authenticated API crawling. You can only include
domains you own that are associated with this Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("domain")}
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

Terraform syntax:

```terraform
domains {
  # Configure direct properties listed below.
}
```

<a id="canonical-fc011218e98e7dcde6adc3dabf792247068d0aa271e3c50a2a7c3bacf27524e9"></a>

## Direct properties — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains / e75b0d1b90d2 / 3

<a id="canonical-ce787f3b3bcc3d04e2cebb39b2fd3f48c7ed0409bc67e34961c717f141d11012"></a>

<a id="canonical-7224f87c10250675254e60fb4b990a794710e2b036f1caaf3e3af769b3124b02"></a>

## domain property — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains / e75b0d1b90d2 / 4

Type: `"string"`. Optional.

Select the domain to execute API Crawling with given credentials.

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

- [simple_login](resources--http_loadbalancer--reference--group-026.md#canonical-910c3d6d86abbb811dbbb5900d641e59b952dd2d52420ad931f9966235c6c4b4): complete subsection reference.

<a id="canonical-c13ba6e6a008d231ddce24ecabee8dddf7543c8b72dfb3f8fcd6c383ab2e46f9"></a>

## Next pages — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains / e75b0d1b90d2 / 5

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-026.md#canonical-910c3d6d86abbb811dbbb5900d641e59b952dd2d52420ad931f9966235c6c4b4)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-026.md#canonical-8f6d639c170c0bd13c34d8bb96ad2ba8822f63f5803db5617273886cf8048716)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-910c3d6d86abbb811dbbb5900d641e59b952dd2d52420ad931f9966235c6c4b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84de59338a23c868214c4b13aaa62bd46d11a7b85656b9f8332d16205439d607"></a>

## single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 307bca4da31c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-026.md#canonical-bd895f7d5fc980e7c41cfa7e68ec079cf979376708b050350a2df4257735b3c6)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-026.md#canonical-8f6d639c170c0bd13c34d8bb96ad2ba8822f63f5803db5617273886cf8048716)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-026.md#canonical-9ba0ce8049ded7ef0dac39c4265b646c82d6c6cfc6c807629b8d87d0ff430903)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login

<a id="canonical-afeb1aec802cd2d694790718a558c9992781f878363748551a23e843f851622d"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for simple login.

Receipt-pinned upstream constraints:

```json
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
simple_login {
  # Configure direct properties listed below.
}
```

<a id="canonical-f1a3666454ae729807b795708a1afef0b8a0aaee9b22ca6ed64f08e4f131c486"></a>

## Direct properties — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 307bca4da31c / 3

- [password](resources--http_loadbalancer--reference--group-026.md#canonical-e62c4d7c140787dcea36f6355e11f63e2bbe0cfd1dfa022614088b7da3801154): complete subsection reference.

<a id="canonical-404e692b65322967636f9af9a9d423aa555d23bdbdb673140256309cac5db86c"></a>

<a id="canonical-31a24f9f420f6705e33789ca102598b314839f99b75678c7139e4a1c3beb18c3"></a>

## user property — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 307bca4da31c / 4

Type: `"string"`. Optional.

Enter the username to assign credentials for the selected domain to crawl.

Provider validators and defaults (from schema source):

```go
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-5773cbae9989a34602e44f865eb8e94e563e0a245642f37b216bd2522ba55edb"></a>

## Next pages — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 307bca4da31c / 5

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-026.md#canonical-e62c4d7c140787dcea36f6355e11f63e2bbe0cfd1dfa022614088b7da3801154)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-026.md#canonical-9ba0ce8049ded7ef0dac39c4265b646c82d6c6cfc6c807629b8d87d0ff430903)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e62c4d7c140787dcea36f6355e11f63e2bbe0cfd1dfa022614088b7da3801154"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7550cb9f038b2519fe612eab89794a0c64b447f867c72a6ce6b1989e8d57d8c0"></a>

## single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 1565e93351b9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-026.md#canonical-bd895f7d5fc980e7c41cfa7e68ec079cf979376708b050350a2df4257735b3c6)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-026.md#canonical-8f6d639c170c0bd13c34d8bb96ad2ba8822f63f5803db5617273886cf8048716)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-026.md#canonical-9ba0ce8049ded7ef0dac39c4265b646c82d6c6cfc6c807629b8d87d0ff430903)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-026.md#canonical-910c3d6d86abbb811dbbb5900d641e59b952dd2d52420ad931f9966235c6c4b4)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password

<a id="canonical-e03f456a81b25869b76726feca7aacbe1cbd2957fd88c750420318150cb3913f"></a>

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

<a id="canonical-2d53eeb0c53da643e8a1531eae9fe581583e1cda04bde3214fe0744c96e16c9a"></a>

## Direct properties — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 1565e93351b9 / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-026.md#canonical-2ef0af8161116eb7a3b483085f9283838c5b0217d53c40f5a3d3a35f57947731): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-026.md#canonical-c9bc7f3f5df068b58b2b107ef939622f3cedc6e574691be254cd2b5a6170c2b4): complete subsection reference.

<a id="canonical-1ab12236b87702d7187730f2c94d083ee32d174c455b65bea8b34208f445ca61"></a>

## Next pages — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 1565e93351b9 / 4

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info](resources--http_loadbalancer--reference--group-026.md#canonical-2ef0af8161116eb7a3b483085f9283838c5b0217d53c40f5a3d3a35f57947731)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info](resources--http_loadbalancer--reference--group-026.md#canonical-c9bc7f3f5df068b58b2b107ef939622f3cedc6e574691be254cd2b5a6170c2b4)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-026.md#canonical-910c3d6d86abbb811dbbb5900d641e59b952dd2d52420ad931f9966235c6c4b4)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2ef0af8161116eb7a3b483085f9283838c5b0217d53c40f5a3d3a35f57947731"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10f7382a2b709ac54119725d93905db396de4b6305df4827f14b99f4b973944a"></a>

## single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / f84ea475ab27 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-026.md#canonical-bd895f7d5fc980e7c41cfa7e68ec079cf979376708b050350a2df4257735b3c6)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-026.md#canonical-8f6d639c170c0bd13c34d8bb96ad2ba8822f63f5803db5617273886cf8048716)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-026.md#canonical-9ba0ce8049ded7ef0dac39c4265b646c82d6c6cfc6c807629b8d87d0ff430903)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-026.md#canonical-910c3d6d86abbb811dbbb5900d641e59b952dd2d52420ad931f9966235c6c4b4)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-026.md#canonical-e62c4d7c140787dcea36f6355e11f63e2bbe0cfd1dfa022614088b7da3801154)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info

<a id="canonical-2b35571d06b9c8974cb5925799e61ee64dfdc28eb9e135542228f4c65def34b9"></a>

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

<a id="canonical-bf0b83da0dc1a5a92d0358d933a9d048d36e7c0f44b6001197ede0ab2b1e0456"></a>

## Direct properties — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / f84ea475ab27 / 3

<a id="canonical-11d359bd90f2fa12b6e6848dd539b349cd629f007c22a5025384bb80b2132fc2"></a>

<a id="canonical-1d091b11ca6f0e803ff593dc4cf20e7fa7bf141e4f6226a960269096698b5873"></a>

## decryption_provider property — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / f84ea475ab27 / 4

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

<a id="canonical-26ad3811650011aa2679c9cad3e275a664a99ae08af2670367afb545c862a19f"></a>

<a id="canonical-814b401ed669fe9d2efa12a8079e9989819ceebc1b464114ed773446e87e9be0"></a>

## location property — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / f84ea475ab27 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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

<a id="canonical-c19109def512fa5406858461288986554fdf4d229cb9f9e32a2e167e155f5951"></a>

<a id="canonical-ec5f30a14877542d143f9c6f65024fa868205a281ab9cc519c828e390560b2f1"></a>

## store_provider property — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / f84ea475ab27 / 6

Type: `"string"`. Optional.

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

<a id="canonical-6f6a408b04ed15be2883123b46f378ab024d9cbd8681bf7a61c52409bf2c6154"></a>

## Next pages — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / f84ea475ab27 / 7

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-026.md#canonical-e62c4d7c140787dcea36f6355e11f63e2bbe0cfd1dfa022614088b7da3801154)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c9bc7f3f5df068b58b2b107ef939622f3cedc6e574691be254cd2b5a6170c2b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-787b969fb06698cf895b1b0d0b5859a5f4b42e82875940fa99e0008318ee9f87"></a>

## single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 507d2c1b39f2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-026.md#canonical-bd895f7d5fc980e7c41cfa7e68ec079cf979376708b050350a2df4257735b3c6)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-026.md#canonical-8f6d639c170c0bd13c34d8bb96ad2ba8822f63f5803db5617273886cf8048716)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-026.md#canonical-9ba0ce8049ded7ef0dac39c4265b646c82d6c6cfc6c807629b8d87d0ff430903)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-026.md#canonical-910c3d6d86abbb811dbbb5900d641e59b952dd2d52420ad931f9966235c6c4b4)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-026.md#canonical-e62c4d7c140787dcea36f6355e11f63e2bbe0cfd1dfa022614088b7da3801154)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

<a id="canonical-ced1d3c257a9aa5c14dba69d8325c231be565ca1c7eef738ba9ef75cf054d041"></a>

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

<a id="canonical-cd49da135ccc2b9d5d8d025d598b59a155173d31c31e8451a827d8814d85bf2b"></a>

## Direct properties — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 507d2c1b39f2 / 3

<a id="canonical-f58c5a1c207254d83b6eff62b3504db94ec5a4e31bc813719835ca6b1993dcda"></a>

<a id="canonical-0338102fa9ae09559b0532f4fe1ecd30d4bef529f484777abe31cfc45b760c77"></a>

## provider_ref property — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 507d2c1b39f2 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-07ef5f57b7935ac844dbc5a772cf4c317d1c1c0e1d3a35d97604a8b0ad6cf3cf"></a>

<a id="canonical-a0d9c6321e7a984de8448194aa5e88e404dff28a59117fafbf3c7b500479f1e5"></a>

## url property — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 507d2c1b39f2 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-466715e97a4f4c974453c424bd1bdf5d5c6841a67d43e69744811222308c9e1e"></a>

## Next pages — single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_log / 507d2c1b39f2 / 6

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-026.md#canonical-e62c4d7c140787dcea36f6355e11f63e2bbe0cfd1dfa022614088b7da3801154)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-224059c3dc911c405de20583131137ca79cface5ff51b8febd4c7a9dea8d7154"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7352c86730d334f2d3b97b6f18bd66b50585adacbeea0b41bd9281c5e9a7d6d"></a>

## single_lb_app.enable_discovery.api_crawler.disable_api_crawler — single_lb_app.enable_discovery.api_crawler.disable_api_crawler / 4260f0f29cda / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-026.md#canonical-bd895f7d5fc980e7c41cfa7e68ec079cf979376708b050350a2df4257735b3c6)
- single_lb_app.enable_discovery.api_crawler.disable_api_crawler

<a id="canonical-96c0bc63f55b25a34d8d84e9db97ef440d2598b70ff6f9d9bfe0b060ff4368e3"></a>

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
disable_api_crawler = {}
```

<a id="canonical-801a3de9f7c6b4feda38f28bca656e89ab54404babff0089d5894949381009bc"></a>

## Direct properties — single_lb_app.enable_discovery.api_crawler.disable_api_crawler / 4260f0f29cda / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e825bb5c80676232e3be76b72aebb1b03d1d0cfe204e0e4efa293fb8d082fb69"></a>

## Next pages — single_lb_app.enable_discovery.api_crawler.disable_api_crawler / 4260f0f29cda / 4

- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-026.md#canonical-bd895f7d5fc980e7c41cfa7e68ec079cf979376708b050350a2df4257735b3c6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-290d1ecaf61b258bc0f26ba76908c342c20ea187fbdb217fff8d74f7654f3028"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c1e077532524220e8808b654f6d897b6cd12d184b72ea530fcdddc6a659cf50"></a>

## single_lb_app.enable_discovery.api_discovery_from_code_scan — single_lb_app.enable_discovery.api_discovery_from_code_scan / 898b03e2466f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- single_lb_app.enable_discovery.api_discovery_from_code_scan

<a id="canonical-05693827b30a2b312b245f195c419ff0050436736d1dd5ef6e79db87d2d174e5"></a>

Type: `"object"`. single nested block, Optional.

Select Code Base and Repositories.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-33c87d6cc5317e503556f1c5d80beeb5f3f03eb2ff98254c94986ddcaacf8c57"></a>

## Direct properties — single_lb_app.enable_discovery.api_discovery_from_code_scan / 898b03e2466f / 3

- [code_base_integrations](resources--http_loadbalancer--reference--group-026.md#canonical-3a3780bf6d7f9475a936015d9db310d7870bd19544fd66b7d045dd6ebe9837aa): complete subsection reference.

<a id="canonical-be4f5a7028111b5ea78d2f984f1aca93b3cb6f77eb0bd7311a1d0239b4224712"></a>

## Next pages — single_lb_app.enable_discovery.api_discovery_from_code_scan / 898b03e2466f / 4

- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-026.md#canonical-3a3780bf6d7f9475a936015d9db310d7870bd19544fd66b7d045dd6ebe9837aa)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3a3780bf6d7f9475a936015d9db310d7870bd19544fd66b7d045dd6ebe9837aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c2367de3cff48a40544513db4c532f2ea36631a46989455e37e5c605b1f0fd5"></a>

## single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 096eb0ee8902 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-026.md#canonical-290d1ecaf61b258bc0f26ba76908c342c20ea187fbdb217fff8d74f7654f3028)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations

<a id="canonical-e7fb108ec00e39100a86df0d17f61ef4826da478a8bc845ddb5df859f48d641c"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for code base integrations.

Upstream description:

Configuration parameter for code base integrations

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-01a1c76d8368210a1755290f596ab9913f87be3eafee25aac8e7f295ee880b65"></a>

## Direct properties — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 096eb0ee8902 / 3

- [all_repos](resources--http_loadbalancer--reference--group-026.md#canonical-ab2735f360be1faa721fd23ae901aea547c45c3cb30861f02da17e4c5765cbf7): complete subsection reference.

- [code_base_integration](resources--http_loadbalancer--reference--group-026.md#canonical-13f95a48a0746900a01de533b0637eeb9399cee77b0e4eaf5cae08d203a728ee): complete subsection reference.

- [selected_repos](resources--http_loadbalancer--reference--group-026.md#canonical-18335fb4cea5343ff063648078b58882a410e5e145513d7173c2bcf79e17182f): complete subsection reference.

<a id="canonical-a08382a6bbcc9280a3c3bd917377e5e6c7a203f376e1dddfd361f8d15c2fee12"></a>

## Next pages — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 096eb0ee8902 / 4

- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos](resources--http_loadbalancer--reference--group-026.md#canonical-ab2735f360be1faa721fd23ae901aea547c45c3cb30861f02da17e4c5765cbf7)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration](resources--http_loadbalancer--reference--group-026.md#canonical-13f95a48a0746900a01de533b0637eeb9399cee77b0e4eaf5cae08d203a728ee)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos](resources--http_loadbalancer--reference--group-026.md#canonical-18335fb4cea5343ff063648078b58882a410e5e145513d7173c2bcf79e17182f)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-026.md#canonical-290d1ecaf61b258bc0f26ba76908c342c20ea187fbdb217fff8d74f7654f3028)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ab2735f360be1faa721fd23ae901aea547c45c3cb30861f02da17e4c5765cbf7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d47a9b2b61c415f6482fcd93bd36cc0ffe36f3fbd000debb65f91f39f77ff63"></a>

## single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 92b3bb7a3e78 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-026.md#canonical-290d1ecaf61b258bc0f26ba76908c342c20ea187fbdb217fff8d74f7654f3028)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-026.md#canonical-3a3780bf6d7f9475a936015d9db310d7870bd19544fd66b7d045dd6ebe9837aa)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos

<a id="canonical-dfebca56fc3d8b1f0d46fafc2c8c9e4bdc92d25c127ce59b24ee3121f3f033c0"></a>

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
all_repos = {}
```

<a id="canonical-03355ceda65ed9ddef2c76396177cdc15e44d1f0b4ed8dc1d08091a4e7e8b555"></a>

## Direct properties — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 92b3bb7a3e78 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-317bcc5e52a2da62c1ad7c4b0660edbc7ed02b8924ff71c8f024b11fb14d5368"></a>

## Next pages — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 92b3bb7a3e78 / 4

- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-026.md#canonical-3a3780bf6d7f9475a936015d9db310d7870bd19544fd66b7d045dd6ebe9837aa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-13f95a48a0746900a01de533b0637eeb9399cee77b0e4eaf5cae08d203a728ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2e3620240874b08fd035d799495f675d06a21e6defa6339f887c13693f7266f"></a>

## single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 1daed3f9f377 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-026.md#canonical-290d1ecaf61b258bc0f26ba76908c342c20ea187fbdb217fff8d74f7654f3028)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-026.md#canonical-3a3780bf6d7f9475a936015d9db310d7870bd19544fd66b7d045dd6ebe9837aa)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration

<a id="canonical-9ef83bdb03db2bb13948161afa57dcb85752a22f8558f944f552935bd3cb771b"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

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

Terraform syntax:

```terraform
code_base_integration {
  # Configure direct properties listed below.
}
```

<a id="canonical-89794c303202e36189fa4ec7516f0781216205aec7c0bc680722ea34a41249a8"></a>

## Direct properties — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 1daed3f9f377 / 3

<a id="canonical-e76ec93bb53d670c52b2a00126dffd4d7aabaf7535af7325505ecb306633b32c"></a>

<a id="canonical-39df0eb4f215c39f285681cce6216efa389c2204287b93fd273e4114f1f519e7"></a>

## name property — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 1daed3f9f377 / 4

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

<a id="canonical-eec3fc979049bb05a5989cccd5f18ac55c1874876f5e55c487ab476703d0f8cd"></a>

<a id="canonical-5e5368b4f429bdffbdd6d06937025e5d85af0f37d955424d0e7d3c9d634517c3"></a>

## namespace property — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 1daed3f9f377 / 5

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

<a id="canonical-701eed47b2be2e22b0e57e516a10b7ca49fce70b51874976d1f8c9c6a2d8451d"></a>

<a id="canonical-e8e115d71ca77d12f17809fcbc1d997d5219cdb417f3510eb88bf9a665a7e9b4"></a>

## tenant property — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 1daed3f9f377 / 6

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

<a id="canonical-0252218f6e796d0aea0296d800cdc88e307557d0fbb35f8ee9243b6581aef1e2"></a>

## Next pages — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / 1daed3f9f377 / 7

- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-026.md#canonical-3a3780bf6d7f9475a936015d9db310d7870bd19544fd66b7d045dd6ebe9837aa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-18335fb4cea5343ff063648078b58882a410e5e145513d7173c2bcf79e17182f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92db8ad8ad75a6bc010cb2b2cf6ca10e117ac162acec77c5bab287e13b033d16"></a>

## single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / b5f9bde5ac11 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-026.md#canonical-290d1ecaf61b258bc0f26ba76908c342c20ea187fbdb217fff8d74f7654f3028)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-026.md#canonical-3a3780bf6d7f9475a936015d9db310d7870bd19544fd66b7d045dd6ebe9837aa)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="canonical-7d46dd6019957eb57e9ac4f78822c58a6109ac133cf0b6490979cde23f7fd297"></a>

Type: `"object"`. single nested block, Optional.

Select which API repositories represent the LB applications.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-49e20f938d72a721f03f4279309217689bb4f9872a7ae5d07f0c98a29a5b57ca"></a>

## Direct properties — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / b5f9bde5ac11 / 3

<a id="canonical-d2ffb01ce74e1c6d0d98bf4baee08e9bb585d6a3b7749f4eac233450547f1152"></a>

<a id="canonical-8801a28ff59ba5632152a75f837ecf93d1629ce6b58ab60efe8e3359299cace8"></a>

## api_code_repo property — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / b5f9bde5ac11 / 4

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

<a id="canonical-47d85c36689b61eb5d67c19d4ae499b505ed3093d4b2c4b5e51b82009c0f2a30"></a>

## Next pages — single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integratio / b5f9bde5ac11 / 5

- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-026.md#canonical-3a3780bf6d7f9475a936015d9db310d7870bd19544fd66b7d045dd6ebe9837aa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0f230849969e44c2f2ac246346b86e14f6a0843e340c1789ac232fd51679ced1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2e081a48c93ec321f92f0ef68a4f964fdc26e9ed6b14e281e2f1dd4b9a80e57"></a>

## single_lb_app.enable_discovery.custom_api_auth_discovery — single_lb_app.enable_discovery.custom_api_auth_discovery / f4464e2063f1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- single_lb_app.enable_discovery.custom_api_auth_discovery

<a id="canonical-3fe7893e07338e1721736d2814359c9dcebec9686c0a07f3510828ae357456e0"></a>

Type: `"object"`. single nested block, Optional.

API Discovery Advanced Settings. API Discovery Advanced settings.

Upstream description:

API Discovery Advanced settings.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-b3ad79df8dc738b6280b8989e4714f3540f731b1ac0227c1ca79e70c25ec7f03"></a>

## Direct properties — single_lb_app.enable_discovery.custom_api_auth_discovery / f4464e2063f1 / 3

- [api_discovery_ref](resources--http_loadbalancer--reference--group-026.md#canonical-740b5343ea8cc056bc8aebb2b6f1032684178f6381d040489f4ccff879f159fb): complete subsection reference.

<a id="canonical-85ec4f77a17e606fd0375963a3b51184bbe34a7c684db751cc045cbd8da4922c"></a>

## Next pages — single_lb_app.enable_discovery.custom_api_auth_discovery / f4464e2063f1 / 4

- [single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref](resources--http_loadbalancer--reference--group-026.md#canonical-740b5343ea8cc056bc8aebb2b6f1032684178f6381d040489f4ccff879f159fb)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-740b5343ea8cc056bc8aebb2b6f1032684178f6381d040489f4ccff879f159fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edcb5ec784b02155c4cd99d7838d199f405117e7b9ae3e44bd3f524e31ca6acd"></a>

## single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref — single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref / 75618d8efdd5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [single_lb_app.enable_discovery.custom_api_auth_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-0f230849969e44c2f2ac246346b86e14f6a0843e340c1789ac232fd51679ced1)
- single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref

<a id="canonical-3aaf1ed4a44252f91eb39c7680c77ebefd122a57f12cd47678bc25ec4470c584"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

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

Terraform syntax:

```terraform
api_discovery_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-43819bf3d247fb2c3c7e8c765395d87c5d1f664a739069c2f4e93628d1d20b25"></a>

## Direct properties — single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref / 75618d8efdd5 / 3

<a id="canonical-59066995c2c7a8ead76d217f356edd3e8c8d383a1a75c2c0ad5d9fdc7660428b"></a>

<a id="canonical-31e3641fe82a4912873c0009656d03940f7a14ddda1c38e76b52feaf55539ac2"></a>

## name property — single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref / 75618d8efdd5 / 4

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

<a id="canonical-ae2f354e66dc34cc97a247f2f1ee7021e7ffb98fa2e3917a09c59e441d4c5788"></a>

<a id="canonical-c89ceceb418111df691df149091affa33e9e36d37a0e1b64f05a4324af43b0a5"></a>

## namespace property — single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref / 75618d8efdd5 / 5

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

<a id="canonical-2650dea200166da1d4f3fe7b9ce3370c381ccbcda2228b275a1eb89bcdd09370"></a>

<a id="canonical-2d561a97921f51f73f70d36f0c91262c131cac2f9f465ec2b821e63700be814e"></a>

## tenant property — single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref / 75618d8efdd5 / 6

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

<a id="canonical-0ec28fc4e95751a8f1f518ff06a352a94d3f3d7b7c76bb8ee0fa2b98997bc05a"></a>

## Next pages — single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref / 75618d8efdd5 / 7

- [single_lb_app.enable_discovery.custom_api_auth_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-0f230849969e44c2f2ac246346b86e14f6a0843e340c1789ac232fd51679ced1)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-76d497446e6c858ad5778d2211bb253f9e943c331cef340d027714372cccd19c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6699860860bbb46ed5ec2b32b6e2fbf70ed708e0018cd6639d923e68fc92738f"></a>

## single_lb_app.enable_discovery.default_api_auth_discovery — single_lb_app.enable_discovery.default_api_auth_discovery / 41d5d68b6201 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- single_lb_app.enable_discovery.default_api_auth_discovery

<a id="canonical-fd3347ed6f4695a81584c48fa6637c837a29894c4513fb31907876b6dbb347f8"></a>

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
default_api_auth_discovery = {}
```

<a id="canonical-a122dff811e27b9436bbf076635fc70fa78ca22eaa8cb01636843f21a4119547"></a>

## Direct properties — single_lb_app.enable_discovery.default_api_auth_discovery / 41d5d68b6201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b3fa5dbf60817212fa730301515a7a718fa02e01c7d4cd546e676586aa06208c"></a>

## Next pages — single_lb_app.enable_discovery.default_api_auth_discovery / 41d5d68b6201 / 4

- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5d6606df66372cf747c7cfbd5108a1aeb91ee8aab307276c5c15a5f9a1ae4232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-220fa1f5d2c79194ea61880994f9e89068ba9aff41b8317ab19c92f0ab3e0945"></a>

## single_lb_app.enable_discovery.disable_learn_from_redirect_traffic — single_lb_app.enable_discovery.disable_learn_from_redirect_traffic / 30fccf463a2a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- single_lb_app.enable_discovery.disable_learn_from_redirect_traffic

<a id="canonical-c04d157d6931b973bfb2ba87e790a61ebd80472214d1a765c9fae50b714e6294"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable learn from redirect traffic.

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
disable_learn_from_redirect_traffic = {}
```

<a id="canonical-11ed0870a9e40c043611b9134b234d21b4fccfda6ae7b972a344cde83df895ac"></a>

## Direct properties — single_lb_app.enable_discovery.disable_learn_from_redirect_traffic / 30fccf463a2a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fcad4d232a929f66b0b3ee93321de7775be23abf71dd0c365f5ccdde10d6f92c"></a>

## Next pages — single_lb_app.enable_discovery.disable_learn_from_redirect_traffic / 30fccf463a2a / 4

- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-29185b0644f16e70f840996434781ad14fcb85c7631936b2402bc5c5c69e10b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8414cb09324c4e12955c96bd3e7654c038956a81b5f9c0d148286f913cefef5"></a>

## single_lb_app.enable_discovery.discovered_api_settings — single_lb_app.enable_discovery.discovered_api_settings / 335ddb657eb3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- single_lb_app.enable_discovery.discovered_api_settings

<a id="canonical-3dcd23a794fba646665527396c5b2ecce4f2740b3717bbb3575fb2cb3c547e88"></a>

Type: `"object"`. single nested block, Optional.

Discovered API Settings. Configure Discovered API Settings.

Upstream description:

Configure Discovered API Settings.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-df5f08c4feb969cc0606c56ae4ed88afe482269aa1ef2ba392b01fa8316c1fba"></a>

## Direct properties — single_lb_app.enable_discovery.discovered_api_settings / 335ddb657eb3 / 3

<a id="canonical-2cfd2e9d40a84ae5aa9777a10c546f3af6a0199f68345a23f29ab22336f9f3e8"></a>

<a id="canonical-1df9cd1a35e57ee0984afabed3be40939531da3df26f63ac08c5543884ae99d2"></a>

## purge_duration_for_inactive_discovered_apis property — single_lb_app.enable_discovery.discovered_api_settings / 335ddb657eb3 / 4

Type: `"number"`. Optional.

Inactive discovered API will be deleted after configured duration.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-ecdf730874743a2b6e1e1e7456d0c2da42e0c977af99c4a139142e6d35dbdf9c"></a>

## Next pages — single_lb_app.enable_discovery.discovered_api_settings / 335ddb657eb3 / 5

- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5302cce5c74a0a0cee944f31fb856f01ef1a39fb827527c4f6add5b9a2dc4970"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23b908b176e6ec55d7802b420a3686a0acd5e64798d98cc806375fe55f318a6f"></a>

## single_lb_app.enable_discovery.enable_learn_from_redirect_traffic — single_lb_app.enable_discovery.enable_learn_from_redirect_traffic / 86295a7c4089 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- single_lb_app.enable_discovery.enable_learn_from_redirect_traffic

<a id="canonical-583afbd84f700466eaca8af430b12db59fb7778cf040bb5a099183af59093f91"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable learn from redirect traffic.

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
enable_learn_from_redirect_traffic = {}
```

<a id="canonical-285157fc2d6e3a2aedd1bcb0c95d16de80415d220541c174e199fa62d00bde17"></a>

## Direct properties — single_lb_app.enable_discovery.enable_learn_from_redirect_traffic / 86295a7c4089 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3c4689e00404203796ecd28a6b0b9d729dcfa55b97a0069931d0f5ddd4cc2827"></a>

## Next pages — single_lb_app.enable_discovery.enable_learn_from_redirect_traffic / 86295a7c4089 / 4

- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-026.md#canonical-937eaeec6db761b137c4ac861eab94b6a2e9868c74160bbabddc406cf1297b6f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3b8de8aa2df2e33331cd69c20c0ae1bee2de9b546d49200646a53f2a03708272"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-098fb02ffe7bee8dd3f43bca0c6b0b6fb462cfc5e1acf1014f0b5422612b9c15"></a>

## single_lb_app.enable_malicious_user_detection — single_lb_app.enable_malicious_user_detection / 86e1c96a4ccc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- single_lb_app.enable_malicious_user_detection

<a id="canonical-7cd9959db0a71b8ade56c375447b468cd1f0acaa6e93bbb9a888b98a08570fe1"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable malicious user detection.

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
enable_malicious_user_detection = {}
```

<a id="canonical-a44cc44210b27792f90b629b7bdc483744f9592b80319981b0e72d239caea0bd"></a>

## Direct properties — single_lb_app.enable_malicious_user_detection / 86e1c96a4ccc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4fb2d8fe0039a91637dd5bd55775865fd0bd723f76b3432307435bc5668a865a"></a>

## Next pages — single_lb_app.enable_malicious_user_detection / 86e1c96a4ccc / 4

- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-740ee67d743653650e483c667a0cdd4213e3aec04c7b984a7c4b1bf22d3dae63)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ffd5840bac3442e8871f0e6732a39bf72a5cf7ca9d57813bb3e3f7dac2c7c03a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1e9a964f9694b7765a5ea2a8f0b617ed7249502d363f04293aaa07df70bb710"></a>

## slow_ddos_mitigation — slow_ddos_mitigation / cb484e651d56 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- slow_ddos_mitigation

<a id="canonical-f972431591e3c838eb7524b8950a9639c69478e684b5d1713618187deee9f691"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: slow\_ddos\_mitigation, system\_default\_timeouts; Default: system\_default\_timeouts\]
'Slow and low' attacks tie up server resources, leaving none available for servicing requests from
actual users.

Upstream description:

"Slow and low" attacks tie up server resources, leaving none available for servicing requests from
actual users.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("request_headers_timeout"),
  validators.ConflictingObjectAttributes("disable_request_timeout",
    "request_timeout")}
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
  "x-ves-oneof-field-request_timeout_choice": "[\"disable_request_timeout\",\"request_timeout\"]"
}
```

OneOf alternatives in this subsection:

- [slow_ddos_mitigation](resources--http_loadbalancer--reference--group-026.md#canonical-f972431591e3c838eb7524b8950a9639c69478e684b5d1713618187deee9f691)
- [system_default_timeouts](resources--http_loadbalancer--reference--group-026.md#canonical-3c9a4d02eb146132e7026d7455c5076f01b2448fe2d47b87055400ad3874d050)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
slow_ddos_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-38a46c883fb51187f8dd2f78838290f4b965081d88a0a6589e405bc2012de558"></a>

## Direct properties — slow_ddos_mitigation / cb484e651d56 / 3

- [disable_request_timeout](resources--http_loadbalancer--reference--group-026.md#canonical-7fac4d4d6b944623ad25e8e886da7a4241d7902b444bfd30ad9857ab1436c50b): complete subsection reference.

<a id="canonical-beedef77b117349dda76ead10b0ec1e61a87c8e9da17201b7f7d6b88eda7abd9"></a>

<a id="canonical-d9e1fc3d62d493549affcdc7c7bda39eb328e199b6027a1abed8b67773550cb7"></a>

## request_headers_timeout property — slow_ddos_mitigation / cb484e651d56 / 4

Type: `"number"`. Optional.

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The milliseconds. This setting provides protection against Slowloris attacks. Defaults
to \`10000\`.

Upstream description:

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The default value is 10000 milliseconds. This setting provides protection against
Slowloris attacks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2000, 30000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  }
}
```

<a id="canonical-9b792229e71af5eb7744d4cd9bff142b42668c87967827c16c9ce7da744efead"></a>

<a id="canonical-c27694d42087666f93ed3a51b1610762b8abc222ba2ad6b3a70b2dbaee15f9a1"></a>

## request_timeout property — slow_ddos_mitigation / cb484e651d56 / 5

Type: `"number"`. Optional.

Exclusive with \[disable\_request\_timeout\].

Upstream description:

Exclusive with \[disable\_request\_timeout\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2000, 300000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  }
}
```

<a id="canonical-bd4af921159b4e0ad76bb27cd455388e6c27cdb80c67de8868f76276074ebfdf"></a>

## Next pages — slow_ddos_mitigation / cb484e651d56 / 6

- [slow_ddos_mitigation.disable_request_timeout](resources--http_loadbalancer--reference--group-026.md#canonical-7fac4d4d6b944623ad25e8e886da7a4241d7902b444bfd30ad9857ab1436c50b)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7fac4d4d6b944623ad25e8e886da7a4241d7902b444bfd30ad9857ab1436c50b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4dae26db054b3dc23619d35656d3e8fbe1e5a6d4287ea0b7a42584d96886dc46"></a>

## slow_ddos_mitigation.disable_request_timeout — slow_ddos_mitigation.disable_request_timeout / 2ba0b6c249c7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [slow_ddos_mitigation](resources--http_loadbalancer--reference--group-026.md#canonical-ffd5840bac3442e8871f0e6732a39bf72a5cf7ca9d57813bb3e3f7dac2c7c03a)
- slow_ddos_mitigation.disable_request_timeout

<a id="canonical-5156da6358b03db0de0118679a000561aadad53ffb11667fe0a2493411633626"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable request timeout.

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
disable_request_timeout = {}
```

<a id="canonical-d12c71d49782d5041c1f1000ccb8d4b82b7e360acdac5ec618e168afc366e892"></a>

## Direct properties — slow_ddos_mitigation.disable_request_timeout / 2ba0b6c249c7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-23da8bdba1b4728be878147b2c10e9e4af99fe7af9e5d715ac5f2536dea6b7a4"></a>

## Next pages — slow_ddos_mitigation.disable_request_timeout / 2ba0b6c249c7 / 4

- [slow_ddos_mitigation](resources--http_loadbalancer--reference--group-026.md#canonical-ffd5840bac3442e8871f0e6732a39bf72a5cf7ca9d57813bb3e3f7dac2c7c03a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1855ab42bff48be1c839dbc02e86ba75469ba9ba572e6e12d5a69afa2ba6d8c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a0ce92d029d69d3e4c7027c272848f936c4409ca732dd26649f3922251c8cf3"></a>

## source_ip_stickiness — source_ip_stickiness / b02055c5f5eb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- source_ip_stickiness

<a id="canonical-22b74d22d28f94180a0f3f4f765f6c85bd492b35036db22fb390f2d3ba3cfd15"></a>

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
source_ip_stickiness = {}
```

<a id="canonical-9a316278568205da4d549e6990a68c9645fdaff93194284f4abe071f67411728"></a>

## Direct properties — source_ip_stickiness / b02055c5f5eb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-856fcfad4d7a6d36f8dd2fc701f2a8e978763f913c95788771fc44349b3e32fa"></a>

## Next pages — source_ip_stickiness / b02055c5f5eb / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d544b47e171fe74c0dd1c3e1bb390b65307469826b276b138f3f4d3d59ebc155"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d21483f2c8c1b16c4d1ada899d4960731e40ec1606bceb2dec5d8b38c7ec8785"></a>

## system_default_timeouts — system_default_timeouts / 7d7f015987ba / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- system_default_timeouts

<a id="canonical-3c9a4d02eb146132e7026d7455c5076f01b2448fe2d47b87055400ad3874d050"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for system default timeouts.

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
system_default_timeouts = {}
```

<a id="canonical-4c2a3ae7fe11d5c91f6cb10cb6a3ebaa35f3eff67d0deb79d131c20edd5dfcbc"></a>

## Direct properties — system_default_timeouts / 7d7f015987ba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1532ac6068be78bc5be2715136091efafbfc416f69ba10dfbc78f64862de4aab"></a>

## Next pages — system_default_timeouts / 7d7f015987ba / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0d031c5e02b9b01edc14906dd7d479341ac318e527a84ba537a24a689885ff3a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-071dbdb5d93468a5ac2cd9d79df47858a3d30c3e1d90b5bd4bfe5619a205d44e"></a>

## timeouts — timeouts / 5b98c717e122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- timeouts

<a id="canonical-e974332bdf248fc0c4b6ce921147246b23e89040837fba6777de499367f27ca4"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-7a301b4a2357474e73fa241631809673d408c3d58703d77b71eadba8499d92dc"></a>

## Direct properties — timeouts / 5b98c717e122 / 3

<a id="canonical-7e8f873e2d8ce58bf1639e074442c19be1d60ca87534519d477a0201fe3f2e3c"></a>

<a id="canonical-1e4fbba9fb44c0e1700f3af1cbbf408b5f26776c134e7044702f0190292fb315"></a>

## create property — timeouts / 5b98c717e122 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-59fee319cf3a3815d7b10090ac7a67d817d852b0101f355972fc79baae8cfd1c"></a>

<a id="canonical-90cab9116799711e5d167c4117bb7d4b568f2c5e94ddd27034359d0f0f7b881f"></a>

## delete property — timeouts / 5b98c717e122 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-b977365941fe76bd71db61e3d8680df86e49dabbd4cf449b174beeba2ae31b4a"></a>

<a id="canonical-86c21603c17bd2fa503bca38922edf8f5c40fd6d958e91a7b69037b43a80a47d"></a>

## read property — timeouts / 5b98c717e122 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-9b1e5cae20bbe5aefb9ef2c1a3980e7a8899024096613a98a0f5bf0090d18317"></a>

<a id="canonical-ab6a9048746c7a61e571c3658a1b399ade1f8749014a76ad469a840c137c1892"></a>

## update property — timeouts / 5b98c717e122 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-ad1caa181fcbcfc968dbcf5353542ecff7d269cbb0d759ae96f34ff0abe361cc"></a>

## Next pages — timeouts / 5b98c717e122 / 8

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e0af00365f8f0ebfccab32942f6b2726b0806a55ec8ea4b3b523dd7cebb7e4c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ca09dab6b839e84ee1b52790d3cefd1f93dc0f79b35be4275542ca29eb55edf"></a>

## trusted_clients — trusted_clients / e7abe41a4555 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- trusted_clients

<a id="canonical-1f70e59f5010c13593a8490385884f509ac2142dd626a2533bd8efbeab57d46f"></a>

Type: `"object"`. list nested block, Optional.

Define rules to skip processing of one or more features such as WAF, Bot Defense etc.

Upstream description:

Define rules to skip processing of one or more features such as WAF, Bot Defense etc. For clients.

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
trusted_clients {
  # Configure direct properties listed below.
}
```

<a id="canonical-540922f1537f3ba9e3fd73716ecd36b821ec34ac65d4981dcf387536d9e90acd"></a>

## Direct properties — trusted_clients / e7abe41a4555 / 3

<a id="canonical-08966f7e0199b18f526f92d946659a8a7fef34bbbab25b34df3bbb9ac6071262"></a>

<a id="canonical-08dd4b712ecfd45bd12a54e9e5bf3eedabe72a213d328d86f7f9a9745d39959d"></a>

## actions property — trusted_clients / e7abe41a4555 / 4

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

<a id="canonical-852fdd31ddaa5361a0d50be61530d225b0e02a217b285c1c7e2e67e6e40a9386"></a>

<a id="canonical-c3cea6c5ec3cde7195198e7248f4eb14dc1eda8ad595d979a08af5d05f3f77a2"></a>

## as_number property — trusted_clients / e7abe41a4555 / 5

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

- [bot_skip_processing](resources--http_loadbalancer--reference--group-026.md#canonical-42c1fdf304c324e3d06cb3fd20f514dd570628b022089e0d50d934db6e4f5bf1): complete subsection reference.

<a id="canonical-b38ee6001373c697525dd3fff781c516765daecade1e2ed62c71a007a398ce15"></a>

<a id="canonical-81e4de7d7437847e84baac7b3361bb91616cb34d0631ea5c3472751ccdd6cad7"></a>

## expiration_timestamp property — trusted_clients / e7abe41a4555 / 6

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

- [http_header](resources--http_loadbalancer--reference--group-026.md#canonical-1c9a0891a9591c76a9f2def08fa77ccc3e43373cc0e8f2faa5078576369fb896): complete subsection reference.

<a id="canonical-2cb50fa5d19a8e42fe7e08db955e0c447fad0bcea2582fbf7bfd9a03a7482ba4"></a>

<a id="canonical-f0a7200c6ce9aeb6f36a72c45fc7bf820a0ea6de524d1a5a7d795e439d6a4020"></a>

## ip_prefix property — trusted_clients / e7abe41a4555 / 7

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

<a id="canonical-b8f4c84f0634667fbc9afa0485cb25801ae31ca18ad5719a24e064e31af7d87f"></a>

<a id="canonical-5e60dc3389c2640c294cfa8e5ed38ad2b98ff3eb12a5545f4fc0e00a5ef2ed61"></a>

## ipv6_prefix property — trusted_clients / e7abe41a4555 / 8

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

- [metadata](resources--http_loadbalancer--reference--group-026.md#canonical-fb981eb0ef88f76f0b3bc64b447c0ee71c6503c255faf982a0c7272c493692e7): complete subsection reference.

- [skip_processing](resources--http_loadbalancer--reference--group-026.md#canonical-e7bcde99df08c728198a5c68efcb47e0c72bc1dd0f1ea6af31a6dc91ff263ec9): complete subsection reference.

<a id="canonical-8188e7a445d6544b8c740c999df7d5757a4566d8396c7ecb0047cbc3767aeda8"></a>

<a id="canonical-8e831d5ed084f82ae247612b355668039ad4be05a0c6f0f91becbcb0c025afbc"></a>

## user_identifier property — trusted_clients / e7abe41a4555 / 9

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

- [waf_skip_processing](resources--http_loadbalancer--reference--group-026.md#canonical-2c8105133c2f94ca96e84f382661a67bdc7ab96f0487eb20d28d0a306dcd5e42): complete subsection reference.

<a id="canonical-7e87f450a51435d59d4f9092d869f6c1db68ec18a546b6c9c530ab9e651d90d7"></a>

## Next pages — trusted_clients / e7abe41a4555 / 10

- [trusted_clients.bot_skip_processing](resources--http_loadbalancer--reference--group-026.md#canonical-42c1fdf304c324e3d06cb3fd20f514dd570628b022089e0d50d934db6e4f5bf1)
- [trusted_clients.http_header](resources--http_loadbalancer--reference--group-026.md#canonical-1c9a0891a9591c76a9f2def08fa77ccc3e43373cc0e8f2faa5078576369fb896)
- [trusted_clients.metadata](resources--http_loadbalancer--reference--group-026.md#canonical-fb981eb0ef88f76f0b3bc64b447c0ee71c6503c255faf982a0c7272c493692e7)
- [trusted_clients.skip_processing](resources--http_loadbalancer--reference--group-026.md#canonical-e7bcde99df08c728198a5c68efcb47e0c72bc1dd0f1ea6af31a6dc91ff263ec9)
- [trusted_clients.waf_skip_processing](resources--http_loadbalancer--reference--group-026.md#canonical-2c8105133c2f94ca96e84f382661a67bdc7ab96f0487eb20d28d0a306dcd5e42)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-42c1fdf304c324e3d06cb3fd20f514dd570628b022089e0d50d934db6e4f5bf1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d76e255a5dabe35452612687d70929945a1478ada75268663e4548fc51d002c4"></a>

## trusted_clients.bot_skip_processing — trusted_clients.bot_skip_processing / e66a6f8b56e4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [trusted_clients](resources--http_loadbalancer--reference--group-026.md#canonical-e0af00365f8f0ebfccab32942f6b2726b0806a55ec8ea4b3b523dd7cebb7e4c9)
- trusted_clients.bot_skip_processing

<a id="canonical-cd1e195a487cb7e2a93833e31bca4e397ead773985d2427ffef8b07640c399df"></a>

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
bot_skip_processing = {}
```

<a id="canonical-046d63c6845b10ee76c87ddccb4b5f52111f372d268039927e37009e0a1c31c2"></a>

## Direct properties — trusted_clients.bot_skip_processing / e66a6f8b56e4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-387d7f06737469e55ad057d2315c5ab7db66ae0428404f8f87145d1395c3af3c"></a>

## Next pages — trusted_clients.bot_skip_processing / e66a6f8b56e4 / 4

- [trusted_clients](resources--http_loadbalancer--reference--group-026.md#canonical-e0af00365f8f0ebfccab32942f6b2726b0806a55ec8ea4b3b523dd7cebb7e4c9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1c9a0891a9591c76a9f2def08fa77ccc3e43373cc0e8f2faa5078576369fb896"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b692f87d1fee6e01800f00a5ec145df77abcdbcfeb565a7213a9829592607468"></a>

## trusted_clients.http_header — trusted_clients.http_header / d58fcda685c3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [trusted_clients](resources--http_loadbalancer--reference--group-026.md#canonical-e0af00365f8f0ebfccab32942f6b2726b0806a55ec8ea4b3b523dd7cebb7e4c9)
- trusted_clients.http_header

<a id="canonical-6d517ccdcb5b0001e4c08fbc9c1ca6fc5b6b24518d60e6aa0942800858271600"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http header.

Upstream description:

Request header name and value pairs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("headers")}
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
http_header {
  # Configure direct properties listed below.
}
```

<a id="canonical-d7968d61ba28a81899980ba0ab3ab0fc11fc1d90a60fc7fcf9b6286a6591941c"></a>

## Direct properties — trusted_clients.http_header / d58fcda685c3 / 3

- [headers](resources--http_loadbalancer--reference--group-026.md#canonical-a7d27062b0ce66fac82d9ca2f774d8b6f90719d9a97d66987726f8bd5cae0b49): complete subsection reference.

<a id="canonical-5ebd9b23f400e185888c2f3ca6ae333d9b41528310fd1e7ccf5de2f24997d9a4"></a>

## Next pages — trusted_clients.http_header / d58fcda685c3 / 4

- [trusted_clients.http_header.headers](resources--http_loadbalancer--reference--group-026.md#canonical-a7d27062b0ce66fac82d9ca2f774d8b6f90719d9a97d66987726f8bd5cae0b49)
- [trusted_clients](resources--http_loadbalancer--reference--group-026.md#canonical-e0af00365f8f0ebfccab32942f6b2726b0806a55ec8ea4b3b523dd7cebb7e4c9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a7d27062b0ce66fac82d9ca2f774d8b6f90719d9a97d66987726f8bd5cae0b49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2c83fc69954d29a513ebd3c53efad32ab80da5ffcef234e140239a88b56cbdc"></a>

## trusted_clients.http_header.headers — trusted_clients.http_header.headers / a1b8995c6221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [trusted_clients](resources--http_loadbalancer--reference--group-026.md#canonical-e0af00365f8f0ebfccab32942f6b2726b0806a55ec8ea4b3b523dd7cebb7e4c9)
- [trusted_clients.http_header](resources--http_loadbalancer--reference--group-026.md#canonical-1c9a0891a9591c76a9f2def08fa77ccc3e43373cc0e8f2faa5078576369fb896)
- trusted_clients.http_header.headers

<a id="canonical-9a11daec30313f697f16991bfae9002bde834204ec32021b9bb3ac4f30d5bb26"></a>

Type: `"object"`. list nested block, Optional.

List of HTTP header name and value pairs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "presence"),
  validators.ConflictingListObjectAttributes("exact",
    "regex"),
  validators.ConflictingListObjectAttributes("presence",
    "regex")}
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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-88b731c5ab9c124958c6698e47f0f672a7f3dcb2f230a3e45efa534610ca509d"></a>

## Direct properties — trusted_clients.http_header.headers / a1b8995c6221 / 3

<a id="canonical-bc997e42a429ffdd517711b9174ffce11f1d47cfdf6749083256c346ba69444c"></a>

<a id="canonical-b5cef7c7b858cf5ab2a858848f2309d39d79268b6685c97f3883740b5fc27972"></a>

## exact property — trusted_clients.http_header.headers / a1b8995c6221 / 4

Type: `"string"`. Optional.

Exclusive with \[presence regex\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regex\] Header value to match exactly.

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

<a id="canonical-cad65a7043de51da75e243d5d456f1078e2762822dd78b8b4f92937d86662e81"></a>

<a id="canonical-777a74f068244381c94f4e796abed87f6c5a4f3a0aec0173b63e4e041b76b36e"></a>

## invert_match property — trusted_clients.http_header.headers / a1b8995c6221 / 5

Type: `"bool"`. Optional.

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

<a id="canonical-df6e52840046b1373dc3a08346751c221d153fc1ee5f8b730606a0359eaf143d"></a>

<a id="canonical-48ebc48fda42d2e71c7ee354436bbb7ba639523a912f7051d18100d4ab3838ee"></a>

## name property — trusted_clients.http_header.headers / a1b8995c6221 / 6

Type: `"string"`. Optional.

Name. Name of the header.

Upstream description:

Name of the header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
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

<a id="canonical-84a1925ba134442a55cf77aeb165db7134b1b13d025e6cbadf0f37f724c77364"></a>

<a id="canonical-08f3910f129cfb9eb90de767f665e5350d0fd0fb3c162df470df4b8a5412a1aa"></a>

## presence property — trusted_clients.http_header.headers / a1b8995c6221 / 7

Type: `"bool"`. Optional.

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

<a id="canonical-0366b6bf20b8cfe18636c8c62ae20efce742679deaf62bb96b941aed889955c3"></a>

<a id="canonical-56e81393fb37bb32e46080c0ac22a5610f55e6da7000f2020e0b3864f4ea5f5e"></a>

## regex property — trusted_clients.http_header.headers / a1b8995c6221 / 8

Type: `"string"`. Optional.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

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

<a id="canonical-f15c3580bcbc2f7241c418dec572afebcf16a94b4c6bd1ea4a2c166894dc61f8"></a>

## Next pages — trusted_clients.http_header.headers / a1b8995c6221 / 9

- [trusted_clients.http_header](resources--http_loadbalancer--reference--group-026.md#canonical-1c9a0891a9591c76a9f2def08fa77ccc3e43373cc0e8f2faa5078576369fb896)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-fb981eb0ef88f76f0b3bc64b447c0ee71c6503c255faf982a0c7272c493692e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8a75b9ae6b0a4124f65cd8dab35d665b394446140f8694823a5148a37d5da02"></a>

## trusted_clients.metadata — trusted_clients.metadata / 3a90e4ba2bbf / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [trusted_clients](resources--http_loadbalancer--reference--group-026.md#canonical-e0af00365f8f0ebfccab32942f6b2726b0806a55ec8ea4b3b523dd7cebb7e4c9)
- trusted_clients.metadata

<a id="canonical-a936946662668c3d8214b59dcf380eb083ee0cb41aa7d370caad8fd0cbc35fa5"></a>

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

<a id="canonical-74c1dbc1cf5188383b4d00a3dbfb5e2d380f3d67806a8759527d975329b914d3"></a>

## Direct properties — trusted_clients.metadata / 3a90e4ba2bbf / 3

<a id="canonical-ad74a090e70121a1d087182980a56bcd8837f7d325fff1fed92ecad29b9ef910"></a>

<a id="canonical-f52dd8fcc855eb69587924697c075af33034214d8fe37f0fcde4126f3aec1a93"></a>

## description_spec property — trusted_clients.metadata / 3a90e4ba2bbf / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-8551d622c5c5abeefb0a6be4325c33a8b2c756d0bb14ad67b23431cf0c86420f"></a>

<a id="canonical-8f397bf16240d4ae5ee3c1f3d1ac1fbe4caeba895ae528437094a89b86350820"></a>

## name property — trusted_clients.metadata / 3a90e4ba2bbf / 5

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

<a id="canonical-ddbfb46d948ef7c8615f133cafc44b862daa95bbb50a416d8e78449acde07e6d"></a>

## Next pages — trusted_clients.metadata / 3a90e4ba2bbf / 6

- [trusted_clients](resources--http_loadbalancer--reference--group-026.md#canonical-e0af00365f8f0ebfccab32942f6b2726b0806a55ec8ea4b3b523dd7cebb7e4c9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e7bcde99df08c728198a5c68efcb47e0c72bc1dd0f1ea6af31a6dc91ff263ec9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce9c1ce03007b5c83456c12ed55e94a57667fcf8281fd1729adf5f50d96e2276"></a>

## trusted_clients.skip_processing — trusted_clients.skip_processing / 80be359e609c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [trusted_clients](resources--http_loadbalancer--reference--group-026.md#canonical-e0af00365f8f0ebfccab32942f6b2726b0806a55ec8ea4b3b523dd7cebb7e4c9)
- trusted_clients.skip_processing

<a id="canonical-27e55ff091a11d4379eec613b27177495456f1882e9dba7bc942136bba4c001c"></a>

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
skip_processing = {}
```

<a id="canonical-99b01fa3b25f5a43fc8ac50575c26385153155a4e92b476977b46a70b213d070"></a>

## Direct properties — trusted_clients.skip_processing / 80be359e609c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3ca78ed1e43abcd558e3f32514d06ae48276e4275761dc5fbb3612b0c77309eb"></a>

## Next pages — trusted_clients.skip_processing / 80be359e609c / 4

- [trusted_clients](resources--http_loadbalancer--reference--group-026.md#canonical-e0af00365f8f0ebfccab32942f6b2726b0806a55ec8ea4b3b523dd7cebb7e4c9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2c8105133c2f94ca96e84f382661a67bdc7ab96f0487eb20d28d0a306dcd5e42"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c29707b0fcc8464477809671cc15063b5ab01f702655daf759c6a676f5a785a"></a>

## trusted_clients.waf_skip_processing — trusted_clients.waf_skip_processing / e6affd5db53f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [trusted_clients](resources--http_loadbalancer--reference--group-026.md#canonical-e0af00365f8f0ebfccab32942f6b2726b0806a55ec8ea4b3b523dd7cebb7e4c9)
- trusted_clients.waf_skip_processing

<a id="canonical-9d0e002f14aeefe4e904ffef1092afb425dc98aa2f58e0e3e700cab094367c08"></a>

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
waf_skip_processing = {}
```

<a id="canonical-3e72a4a06f5ea542758205ed4e4fd7aef3212a6ffc407f74c6614c4b98d19ffc"></a>

## Direct properties — trusted_clients.waf_skip_processing / e6affd5db53f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5e45e35efac2eb69b2716a000e8a05512bd6d17f319c5f3c540b551750b083c0"></a>

## Next pages — trusted_clients.waf_skip_processing / e6affd5db53f / 4

- [trusted_clients](resources--http_loadbalancer--reference--group-026.md#canonical-e0af00365f8f0ebfccab32942f6b2726b0806a55ec8ea4b3b523dd7cebb7e4c9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-fe1f51e76f8a4491a2c7ef8f9cc82d6a6df68859896b3b9775334e6f613cac00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2284f94bec446fa0ded14189c03a75496ef031c2fde4e28090624fdefe0a638f"></a>

## user_id_client_ip — user_id_client_ip / 9124bb6aac2c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- user_id_client_ip

<a id="canonical-0a3dc764cefa1a49ecbc2d5e1c00ffbefc32df510a0276b9311a54d7859dce04"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: user\_id\_client\_ip, user\_identification\] Enable this option. Defaults to \`map\[\]\`.
Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [user_id_client_ip](resources--http_loadbalancer--reference--group-026.md#canonical-0a3dc764cefa1a49ecbc2d5e1c00ffbefc32df510a0276b9311a54d7859dce04)
- [user_identification](resources--http_loadbalancer--reference--group-026.md#canonical-7f99149161da3f6337dfaacc6d4f9cf573ebdc17847ca329eb1e7a8dad79c88c)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
user_id_client_ip = {}
```

<a id="canonical-55cc9c21e979e536418a824cb91b65e6d46798f6133ddb5c38031d976b037b53"></a>

## Direct properties — user_id_client_ip / 9124bb6aac2c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-34543c513f09718f61477574b9087edff8c40a77dbf6b3770d5b7bbe7f4db1c3"></a>

## Next pages — user_id_client_ip / 9124bb6aac2c / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0e0670b13d60febf0c438e4f8a73b5794f67bb35ecad84e1001cef3948fb2a2b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2ee4cfb8cac35594774fa1abb7405ca8baa1e0dae81ad3650654a09d141f91b"></a>

## user_identification — user_identification / b43901f7fa8b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- user_identification

<a id="canonical-7f99149161da3f6337dfaacc6d4f9cf573ebdc17847ca329eb1e7a8dad79c88c"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

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

Terraform syntax:

```terraform
user_identification {
  # Configure direct properties listed below.
}
```

<a id="canonical-26bbb8df78fca654e42d4e95b5cf4a9738ec8f181becf0e4b078d63dfcf5c5e1"></a>

## Direct properties — user_identification / b43901f7fa8b / 3

<a id="canonical-58b72283e0f1df9a836481a2f11e7f1a4db777527620be350e30a1a7c8016e37"></a>

<a id="canonical-cb98e5ca5660e94f7bfa9e915a5e1661eb974e76334436e949c19dfc5cc100d2"></a>

## name property — user_identification / b43901f7fa8b / 4

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

<a id="canonical-c0f583134f30c6336c8edf447506f66a4502315170d5c35bf8c9c728fb576d13"></a>

<a id="canonical-7eb80f887ed73ae06bb0fde8b269fb26a4d5ecf90efb0c802a9a7722ef999500"></a>

## namespace property — user_identification / b43901f7fa8b / 5

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

<a id="canonical-361708f865fcbb9e06f8045f4a1be01060a6be79ed3abd75456928dccb1c17bc"></a>

<a id="canonical-da2fbd5e61f9f2903f344a4aac2e18aa87fbe8bf202dc6bb0753f220f65a153e"></a>

## tenant property — user_identification / b43901f7fa8b / 6

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

<a id="canonical-d792292b9c95e40634b86d34b7e023a9bb17d4072309b9a06b18f22ff6f554ca"></a>

## Next pages — user_identification / b43901f7fa8b / 7

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-86635926dff9a406ce1b9e922504e4198e3c10c37ace3a589dca358b0134d051"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78432d214f1a7d3f42c02138e181f0904cfd9707b8d63a8485e994fd1ee6ed19"></a>

## waf_exclusion — waf_exclusion / 6400ab810b26 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- waf_exclusion

<a id="canonical-cce3822eb9ca0a68ef819ca2b200bc4bacc2d0a4fd56bee7c1466bf6ac8428c4"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for waf exclusion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("waf_exclusion_inline_rules",
    "waf_exclusion_policy")}
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
  "x-ves-oneof-field-waf_exclusion_choice": "[\"waf_exclusion_inline_rules\",\"waf_exclusion_policy\"]"
}
```

Terraform syntax:

```terraform
waf_exclusion {
  # Configure direct properties listed below.
}
```

<a id="canonical-df1032167ab94343a3773d1ff341e2941b28e2685313ce4825ccbe81ecb9843a"></a>

## Direct properties — waf_exclusion / 6400ab810b26 / 3

- [waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-026.md#canonical-91119c921169fc0a7de77ce55f29cbbac40c31b4c5e7c2bd36e4da1bcdc10f64): complete subsection reference.

- [waf_exclusion_policy](resources--http_loadbalancer--reference--group-027.md#canonical-b0b6330d7e2f48f585888321c33d373bf16b58f34991df340655b6284c5c3dcf): complete subsection reference.

<a id="canonical-831a8919f4d5378a9724d9d9e0677862274ad7075d8988da906f04bdfe8c1941"></a>

## Next pages — waf_exclusion / 6400ab810b26 / 4

- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-026.md#canonical-91119c921169fc0a7de77ce55f29cbbac40c31b4c5e7c2bd36e4da1bcdc10f64)
- [waf_exclusion.waf_exclusion_policy](resources--http_loadbalancer--reference--group-027.md#canonical-b0b6330d7e2f48f585888321c33d373bf16b58f34991df340655b6284c5c3dcf)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-91119c921169fc0a7de77ce55f29cbbac40c31b4c5e7c2bd36e4da1bcdc10f64"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9bf89f12112eb8493fd5eb7077d0f12e18d2dcb6ac637d2c4617df7c8a72473"></a>

## waf_exclusion.waf_exclusion_inline_rules — waf_exclusion.waf_exclusion_inline_rules / 4b1f82c960f7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [waf_exclusion](resources--http_loadbalancer--reference--group-026.md#canonical-86635926dff9a406ce1b9e922504e4198e3c10c37ace3a589dca358b0134d051)
- waf_exclusion.waf_exclusion_inline_rules

<a id="canonical-051455b5f279fb2233a64f0edbff7a0102b91bfeefaef841698a1de7422c950d"></a>

Type: `"object"`. single nested block, Optional.

List of WAF exclusion rules that will be applied inline.

Upstream description:

A list of WAF exclusion rules that will be applied inline.

Receipt-pinned upstream constraints:

```json
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
waf_exclusion_inline_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-94876db75892b8598263f1649dabc5f091406bc253e83bff5049de359c8905a6"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules / 4b1f82c960f7 / 3

- [rules](resources--http_loadbalancer--reference--group-026.md#canonical-71a42163514283b6129451520df4d5f8fa82b34374c264bbfd256857d1bc5b52): complete subsection reference.

<a id="canonical-5639c80636553e5513a3cb9fc44364571eedca3041cd543a92bde7cc92cee0d5"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules / 4b1f82c960f7 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-026.md#canonical-71a42163514283b6129451520df4d5f8fa82b34374c264bbfd256857d1bc5b52)
- [waf_exclusion](resources--http_loadbalancer--reference--group-026.md#canonical-86635926dff9a406ce1b9e922504e4198e3c10c37ace3a589dca358b0134d051)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-71a42163514283b6129451520df4d5f8fa82b34374c264bbfd256857d1bc5b52"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8888b16d7f70eda582df8836a793d75cc944ad8eb030cc25a08f405e3bdcd30"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules — waf_exclusion.waf_exclusion_inline_rules.rules / 9941d8ad3240 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [waf_exclusion](resources--http_loadbalancer--reference--group-026.md#canonical-86635926dff9a406ce1b9e922504e4198e3c10c37ace3a589dca358b0134d051)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-026.md#canonical-91119c921169fc0a7de77ce55f29cbbac40c31b4c5e7c2bd36e4da1bcdc10f64)
- waf_exclusion.waf_exclusion_inline_rules.rules

<a id="canonical-525fd71d8bb6d05502aeaa905691d984da3662630e350a694529ac6fb2e196cd"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of WAF Exclusions specific to this Load Balancer.

Upstream description:

An ordered list of WAF Exclusions specific to this Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "exact_value"),
  validators.ConflictingListObjectAttributes("any_domain",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_prefix"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_regex"),
  validators.ConflictingListObjectAttributes("app_firewall_detection_control",
    "waf_skip_processing"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("path_prefix",
    "path_regex")}
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
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-6393adbb6944625654da1f1c66e7f7b23f6c600a500ad06a1497ce2c4ea74035"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules / 9941d8ad3240 / 3

- [any_domain](resources--http_loadbalancer--reference--group-026.md#canonical-6ff2be725c63ed34234306db9cf753141eb3992c9ea5303743895c738f2eb77a): complete subsection reference.

- [any_path](resources--http_loadbalancer--reference--group-026.md#canonical-3479607619954c9ef143cbc20e8a5cea8b344873273a7d09a4febdead42d156d): complete subsection reference.

- [app_firewall_detection_control](resources--http_loadbalancer--reference--group-026.md#canonical-230eedcf2bb3d3dde95c70bfe0cc41e6e23e53f8e449e5cdacbc7a4b32cb209a): complete subsection reference.

<a id="canonical-795474109a0f4183d012368d49e857c7bfe92f3d2725dd90166139a74ec7539c"></a>

<a id="canonical-0310102cb03832a64679f18fb3b48a71b18e4a658234a057573860e8989293be"></a>

## exact_value property — waf_exclusion.waf_exclusion_inline_rules.rules / 9941d8ad3240 / 4

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

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

<a id="canonical-d7273974bf2750c41f1c75de46905edc7c646d920ca236c10c4fa12853477e9b"></a>

<a id="canonical-ec53b95d3d01392c2a6391b1bbb5361ba390736b7c1e8ae6eb1620379177ad7f"></a>

## expiration_timestamp property — waf_exclusion.waf_exclusion_inline_rules.rules / 9941d8ad3240 / 5

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
  }
}
```

- [metadata](resources--http_loadbalancer--reference--group-027.md#canonical-d8c9d0257d9dd5cafba57c04b0cf33abe1b6b8d3439687628f56b2750e459b74): complete subsection reference.

<a id="canonical-0fa5f3e5544cdfbd6576be38a66d26a69ac3c3f99603f11afe09516faddee874"></a>

<a id="canonical-9b1cbfed62c43f6852b5086efac1038d49b5a0a59c4109d3c78f927a9e2991cc"></a>

## methods property — waf_exclusion.waf_exclusion_inline_rules.rules / 9941d8ad3240 / 6

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

<a id="canonical-6b10a91e91e771553ea843c5633b873d305852c68cf35217f7c51af8534ad90f"></a>

<a id="canonical-67df5f08b691620baf0454832d0c3b4e9e0c069fd26bc44b7763db34d71ac342"></a>

## path_prefix property — waf_exclusion.waf_exclusion_inline_rules.rules / 9941d8ad3240 / 7

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths).

Upstream description:

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-22b0f9bef2b9973fb2777c1c37bfa5a80e46fa51662134c7e3b36a0ae361708b"></a>

<a id="canonical-a302ac155ab8176581f3d4e262c29824af1cd054b6bcf3679c1f3a8ac25f36ef"></a>

## path_regex property — waf_exclusion.waf_exclusion_inline_rules.rules / 9941d8ad3240 / 8

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_prefix\] Define the regex for the path. For example, the regex
^/.\*$ will match on all paths.

Upstream description:

Exclusive with \[any\_path path\_prefix\] Define the regex for the path. For example, the regex
^/.\*$ will match on all paths.

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
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-38516394115347d7aa8ed297474a62026d67af299a3418c30cb6d4a0f651f420"></a>

<a id="canonical-5b1dbce3ceb8663a0b6e85c2be99008580bd7864b5c12003a433098dce073a1b"></a>

## suffix_value property — waf_exclusion.waf_exclusion_inline_rules.rules / 9941d8ad3240 / 9

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

- [waf_skip_processing](resources--http_loadbalancer--reference--group-027.md#canonical-6204469c6e5ae9521249e2fec82fce2dd8dacbf615b1a7ba22163ebcab5c0c8a): complete subsection reference.

<a id="canonical-3f197fef80a6c98a1c1c3b58972dceaeb9c7e8a4a44a6782754033c411de1f08"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules / 9941d8ad3240 / 10

- [waf_exclusion.waf_exclusion_inline_rules.rules.any_domain](resources--http_loadbalancer--reference--group-026.md#canonical-6ff2be725c63ed34234306db9cf753141eb3992c9ea5303743895c738f2eb77a)
- [waf_exclusion.waf_exclusion_inline_rules.rules.any_path](resources--http_loadbalancer--reference--group-026.md#canonical-3479607619954c9ef143cbc20e8a5cea8b344873273a7d09a4febdead42d156d)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--http_loadbalancer--reference--group-026.md#canonical-230eedcf2bb3d3dde95c70bfe0cc41e6e23e53f8e449e5cdacbc7a4b32cb209a)
- [waf_exclusion.waf_exclusion_inline_rules.rules.metadata](resources--http_loadbalancer--reference--group-027.md#canonical-d8c9d0257d9dd5cafba57c04b0cf33abe1b6b8d3439687628f56b2750e459b74)
- [waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing](resources--http_loadbalancer--reference--group-027.md#canonical-6204469c6e5ae9521249e2fec82fce2dd8dacbf615b1a7ba22163ebcab5c0c8a)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-026.md#canonical-91119c921169fc0a7de77ce55f29cbbac40c31b4c5e7c2bd36e4da1bcdc10f64)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6ff2be725c63ed34234306db9cf753141eb3992c9ea5303743895c738f2eb77a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f40fb834931a47cc46fafb1a7cb1d32d8d26bf668998cfb11cf331f9e618b2df"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.any_domain — waf_exclusion.waf_exclusion_inline_rules.rules.any_domain / a2d7e6c4dd68 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [waf_exclusion](resources--http_loadbalancer--reference--group-026.md#canonical-86635926dff9a406ce1b9e922504e4198e3c10c37ace3a589dca358b0134d051)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-026.md#canonical-91119c921169fc0a7de77ce55f29cbbac40c31b4c5e7c2bd36e4da1bcdc10f64)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-026.md#canonical-71a42163514283b6129451520df4d5f8fa82b34374c264bbfd256857d1bc5b52)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_domain

<a id="canonical-356b820f7b60d6e3f3d5b42c1f91bd1ecaf934640b439d21f0bda8508c9ef370"></a>

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

<a id="canonical-de6aad29488c0196b20465de7e60ce2dc5413c6539229dfac103169591121847"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.any_domain / a2d7e6c4dd68 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f75a8d9fe85a7b7c130e386446d9f971f0fa5d564c0acd2c1bc929c1882df2db"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.any_domain / a2d7e6c4dd68 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-026.md#canonical-71a42163514283b6129451520df4d5f8fa82b34374c264bbfd256857d1bc5b52)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3479607619954c9ef143cbc20e8a5cea8b344873273a7d09a4febdead42d156d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53e31abd79670b9961b7fdd926b71e8d6c759dea723643ec47bf398d18ce62bf"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.any_path — waf_exclusion.waf_exclusion_inline_rules.rules.any_path / 09bd5bd30487 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [waf_exclusion](resources--http_loadbalancer--reference--group-026.md#canonical-86635926dff9a406ce1b9e922504e4198e3c10c37ace3a589dca358b0134d051)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-026.md#canonical-91119c921169fc0a7de77ce55f29cbbac40c31b4c5e7c2bd36e4da1bcdc10f64)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-026.md#canonical-71a42163514283b6129451520df4d5f8fa82b34374c264bbfd256857d1bc5b52)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_path

<a id="canonical-75083a9abb9dcb3eb6878500b80d4c4ee55d8ad659a5a83de243a356b9a22b4e"></a>

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
any_path = {}
```

<a id="canonical-eef90dcc55c2b88ea6a06bd237baa8b6eba0cd39687b7de7d948864a989ba597"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.any_path / 09bd5bd30487 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6a2b07bab6481be9efad7c81b204146e36a38a9cbab441690dc22577b9a4bf4b"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.any_path / 09bd5bd30487 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-026.md#canonical-71a42163514283b6129451520df4d5f8fa82b34374c264bbfd256857d1bc5b52)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-230eedcf2bb3d3dde95c70bfe0cc41e6e23e53f8e449e5cdacbc7a4b32cb209a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db8944efb5007ea4fa9c261427de544f4b1fd86f70992b4d24a371ee39d28c43"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control / f591916091ef / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [waf_exclusion](resources--http_loadbalancer--reference--group-026.md#canonical-86635926dff9a406ce1b9e922504e4198e3c10c37ace3a589dca358b0134d051)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-026.md#canonical-91119c921169fc0a7de77ce55f29cbbac40c31b4c5e7c2bd36e4da1bcdc10f64)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-026.md#canonical-71a42163514283b6129451520df4d5f8fa82b34374c264bbfd256857d1bc5b52)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control

<a id="canonical-43447880ad8bec5af4d246b8386e59b4c940a99329b87d705bb2aa0109583043"></a>

Type: `"object"`. single nested block, Optional.

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded
from triggering on the defined match criteria.

Receipt-pinned upstream constraints:

```json
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
app_firewall_detection_control {
  # Configure direct properties listed below.
}
```

<a id="canonical-77bd6d880fc4bb725a07a5b0317bbfd83c587ddd9d9f1f8567eafc645021a5bf"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control / f591916091ef / 3

- [exclude_attack_type_contexts](resources--http_loadbalancer--reference--group-026.md#canonical-e1655d4c9b680d58683a1fe355060fe72d29710c9391f87536a9da07f6e3ec78): complete subsection reference.

- [exclude_bot_name_contexts](resources--http_loadbalancer--reference--group-027.md#canonical-070278d972be648bf5e6453cf32400a5ded9a9ba4eb81ae59b5a7c21c4c7d41f): complete subsection reference.

- [exclude_signature_contexts](resources--http_loadbalancer--reference--group-027.md#canonical-3fbc3d363c75d4d2003c8ade3ad7203b7c03e4c62b749a3680c527f068c57a0c): complete subsection reference.

- [exclude_violation_contexts](resources--http_loadbalancer--reference--group-027.md#canonical-abc0c9411e5290b7ef1f9af78613eb6bd87cb5cb85a606e6682b89155fddac8c): complete subsection reference.

<a id="canonical-bf9367b9b12d3993d77d53cf11060d149c954dac362d261669794743b88a2d71"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control / f591916091ef / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts](resources--http_loadbalancer--reference--group-026.md#canonical-e1655d4c9b680d58683a1fe355060fe72d29710c9391f87536a9da07f6e3ec78)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts](resources--http_loadbalancer--reference--group-027.md#canonical-070278d972be648bf5e6453cf32400a5ded9a9ba4eb81ae59b5a7c21c4c7d41f)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts](resources--http_loadbalancer--reference--group-027.md#canonical-3fbc3d363c75d4d2003c8ade3ad7203b7c03e4c62b749a3680c527f068c57a0c)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts](resources--http_loadbalancer--reference--group-027.md#canonical-abc0c9411e5290b7ef1f9af78613eb6bd87cb5cb85a606e6682b89155fddac8c)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-026.md#canonical-71a42163514283b6129451520df4d5f8fa82b34374c264bbfd256857d1bc5b52)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e1655d4c9b680d58683a1fe355060fe72d29710c9391f87536a9da07f6e3ec78"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4087fa5c7c8e8fe1835579d66d9b6020ab5841353935ab0c43176bb4ce1d39d"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 04e4e7f0d189 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [waf_exclusion](resources--http_loadbalancer--reference--group-026.md#canonical-86635926dff9a406ce1b9e922504e4198e3c10c37ace3a589dca358b0134d051)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-026.md#canonical-91119c921169fc0a7de77ce55f29cbbac40c31b4c5e7c2bd36e4da1bcdc10f64)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-026.md#canonical-71a42163514283b6129451520df4d5f8fa82b34374c264bbfd256857d1bc5b52)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--http_loadbalancer--reference--group-026.md#canonical-230eedcf2bb3d3dde95c70bfe0cc41e6e23e53f8e449e5cdacbc7a4b32cb209a)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-b8a53603654d2e173d02791512764e14c6019a73e3cd555b71da46b2ce317d59"></a>

Type: `"object"`. list nested block, Optional.

Exclude an entire attack type only in the named context. For migrated per-parameter exceptions,
prefer this over signature-ID exclusions because one payload can trigger several signatures;
unrelated parameters and attack types remain protected.

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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_attack_type_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-8690b6aa2760c308e7d0a5997d7b385f8fe5a1a6e7435fec235b6b3775533ec8"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 04e4e7f0d189 / 3

<a id="canonical-184df22bb47b61eeed8dca5e658c616adac8398c1205cab6ff3684b2a4762697"></a>
