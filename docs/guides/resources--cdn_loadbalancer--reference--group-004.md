---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-2b4ea4e07d7b42bd3197658dc1ef35a5052b4f4d2c43d4d3442d218dddb2f414"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules / fa701a3b97ad / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-004.md#canonical-61e158d8fead158a86e53a10a9afebcf843615f7b8b333f47d46a57361f0fee9): complete subsection reference.

- [any_url](resources--cdn_loadbalancer--reference--group-004.md#canonical-a18405c6b9a5d1185855e3c84f194358ad74a4353d0793e4ef8c0dd7dea889c6): complete subsection reference.

- [api_endpoint](resources--cdn_loadbalancer--reference--group-004.md#canonical-686ad16da0b696d1716f17ba3636118f4c751e25b900e4722167628a93f97b30): complete subsection reference.

- [api_groups](resources--cdn_loadbalancer--reference--group-004.md#canonical-9798890fc4a51c4c0346379195b2df5f82f17a1b376052afcdac48339dc1f40b): complete subsection reference.

<a id="canonical-a3b8fc7174ba7e0e756cc1b4174a584af2f88ddc1895824920775055b92326f0"></a>

<a id="canonical-61f9054a6442d1b6a0ac3f4a8464d912db02011c54e3c1e6da1d8afe8c23e8a6"></a>

## base_path property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules / fa701a3b97ad / 4

Type: `"string"`. Optional.

Exclusive with \[any\_url api\_endpoint api\_groups\] The base path which this validation applies
to.

Upstream description:

Exclusive with \[any\_url api\_endpoint api\_groups\] The base path which this validation applies
to.

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

- [client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6): complete subsection reference.

- [request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83): complete subsection reference.

<a id="canonical-d280a0e3ceff9bdea37d0c2cba11f20743bccf9f3ed37344097f6deb80c26a3e"></a>

<a id="canonical-841802e6d8a575fb126f126b11f7738114c7b5a168555b042b15e1067a971cdb"></a>

## specific_domain property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules / fa701a3b97ad / 5

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For example:
api.example.com.

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

<a id="canonical-1ed75f89df048c64b564b4eab5809eeea6f1c491eefd06f37edf28d3dadc13e5"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules / fa701a3b97ad / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain](resources--cdn_loadbalancer--reference--group-004.md#canonical-61e158d8fead158a86e53a10a9afebcf843615f7b8b333f47d46a57361f0fee9)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url](resources--cdn_loadbalancer--reference--group-004.md#canonical-a18405c6b9a5d1185855e3c84f194358ad74a4353d0793e4ef8c0dd7dea889c6)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint](resources--cdn_loadbalancer--reference--group-004.md#canonical-686ad16da0b696d1716f17ba3636118f4c751e25b900e4722167628a93f97b30)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups](resources--cdn_loadbalancer--reference--group-004.md#canonical-9798890fc4a51c4c0346379195b2df5f82f17a1b376052afcdac48339dc1f40b)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-61e158d8fead158a86e53a10a9afebcf843615f7b8b333f47d46a57361f0fee9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0ed91594f41e44bf38516805958413d6ff5623ff9b8aa73b5c3d654483955d4"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain / e29b885a0c01 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain

<a id="canonical-1b7cdb520c58b92b631f8f88ef3e84d291804349c2b8abc353730892ec334e84"></a>

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

<a id="canonical-dba451afcd582f10a5de6cadce35f3a8365a6820e4a12d2b33404d6c2db3c3c5"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain / e29b885a0c01 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e1181c0da2aa0cff9fe55f1cfe42cd396af8ae8b5e955c5db078e34584fd921a"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain / e29b885a0c01 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a18405c6b9a5d1185855e3c84f194358ad74a4353d0793e4ef8c0dd7dea889c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-953ee1b77e07fdd5ef120a91c2d21090a896e536b2098d4ef7dde017af02ba8e"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url / 9c24f64679aa / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url

<a id="canonical-b4bad424569a8212fc692e1d4967c8c4e071bceb5a0ccbb820cc1b51cb0cabd1"></a>

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
any_url = {}
```

<a id="canonical-1951f7739a858c7c00e8a50d6dc7e097cb731a99ff7aaa79d52b41136563e8ff"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url / 9c24f64679aa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f1efcefad0bc898333aa7d3ebd17a5fda45a824ebee854c939b4fa8a069b0d7d"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url / 9c24f64679aa / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-686ad16da0b696d1716f17ba3636118f4c751e25b900e4722167628a93f97b30"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc014a2613b882af8c8f807911eacd3cccb92ced578ee789cf9a312875bb2cb7"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoin / 5c677e1764c9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint

<a id="canonical-b49bd4a3d3f860a8479a83cd82ab84e2e7bc5264b7860d55e1616d96b933d822"></a>

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

<a id="canonical-af608ccbb4d3907d711630a380f1b2e35ea1659c11d7d42fa9a8e2de1243e58d"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoin / 5c677e1764c9 / 3

<a id="canonical-6b64d7ba8bac16557eebad702b394e4d59e766d617c8872470a6fd6ef8fee80b"></a>

<a id="canonical-9bd83633504d1f62b8e5971270c7aefc984878c14cf9d4cb43f7a0622b8e3d56"></a>

## methods property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoin / 5c677e1764c9 / 4

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

<a id="canonical-40c36542357a8a48ef2affede7f4c5e50d82ea012d42383bd00c64465c9b193d"></a>

<a id="canonical-5f5e9f481d6488f2c7b5471a104e829816deb4d6c267d9b1330573f68f455b39"></a>

## path property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoin / 5c677e1764c9 / 5

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

<a id="canonical-e9c2c09577443ab040ecf238ed823f3157c0b44f295295a8383eadae0635a458"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoin / 5c677e1764c9 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-9798890fc4a51c4c0346379195b2df5f82f17a1b376052afcdac48339dc1f40b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f543baf36c83899653a6e90219693f86b668f6654646eabb61aa4a00d245481"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups / 4fceba518eb7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups

<a id="canonical-d6ac24f3a1dfeae79f9bb2dd77ea19578422311e53c98501e983b2553cd69efa"></a>

Type: `"object"`. single nested block, Optional.

API Groups.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("api_groups")}
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
api_groups {
  # Configure direct properties listed below.
}
```

<a id="canonical-eeb40d76d8b98ddec68079e58a2e53a5bd58f4e666adbc8efbe745794395e04d"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups / 4fceba518eb7 / 3

<a id="canonical-0b5907896142135e626c03fb1cf5d5bdff9a77a34022db44b36dbf15c1c1b25b"></a>

<a id="canonical-de054d7a4d00d4b2bc20bb27d938090dee724315e38fd996967cde168cf2e9f3"></a>

## api_groups property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups / 4fceba518eb7 / 4

Type: `["list", "string"]`. Optional.

API Groups. Group or collection configuration

Upstream description:

Group or collection configuration

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-bff73694e23bf62777edc02350cdc43dc8228ef57f3ed014c3afb4bb20eadc62"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups / 4fceba518eb7 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-236b35bdc224ff1983fb879fbf098034b9e1fb57faad64b589536872835d0ff9"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 3638a17b7203 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher

<a id="canonical-0842e796df1bf3d628c5da45ec4fcc499535d0a942e4ca7e4889d77e7ba3c72b"></a>

Type: `"object"`. single nested block, Optional.

Client Matcher. Client conditions for matching a rule.

Upstream description:

Client conditions for matching a rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_client",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("client_selector",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
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
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-ip_asn_choice": "[\"any_ip\",\"asn_list\",\"asn_matcher\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

Terraform syntax:

```terraform
client_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-e6aae69ec7fa1ec11850d08e25efc5aa0ef8522a19ba45e0b7f8ec2c88ae3b65"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 3638a17b7203 / 3

- [any_client](resources--cdn_loadbalancer--reference--group-004.md#canonical-b2f60362640b0a449e9fad0c79a1c6de61222b93646210653808449f4d45f4df): complete subsection reference.

- [any_ip](resources--cdn_loadbalancer--reference--group-004.md#canonical-5ce0c40f65cde0182bb73c440f96f92ab576c1da1ddbd22cf7cc1bdc364c98bc): complete subsection reference.

- [asn_list](resources--cdn_loadbalancer--reference--group-004.md#canonical-94fed78ee5fb76728728a9c02680b6d113eeea08ea430c3eb16b3732731a365b): complete subsection reference.

- [asn_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-cec1d71d12f28cfe48e9e2a317ba37066aded4238e8e6e1078bac5ed78fec4c9): complete subsection reference.

- [client_selector](resources--cdn_loadbalancer--reference--group-004.md#canonical-98d37ce5e21534888452ccaf22e11c09cbd7614e69f736872027e464bcf018bc): complete subsection reference.

- [ip_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-77e8d7cd8d6c23924d55df93295322d7d22483df3fe12dae81613e1ca1abb1b8): complete subsection reference.

- [ip_prefix_list](resources--cdn_loadbalancer--reference--group-004.md#canonical-e124c5bd13d90d69bc81aea8e07e5eb4334d2cf305a110e05b6983dd8d5ef101): complete subsection reference.

- [ip_threat_category_list](resources--cdn_loadbalancer--reference--group-004.md#canonical-83594fd7ad25439a1b7854936b93d8749f72926e26d177707a10508861d137c5): complete subsection reference.

- [tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9676752f2f8543eb87055c3eed15bec800f2fd459df5485c1280b35c1fafe53d): complete subsection reference.

<a id="canonical-408fa7deaf44b4678a6f9c112bc9a5b28d28023047c3c4ecc857189b8613ee0f"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 3638a17b7203 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client](resources--cdn_loadbalancer--reference--group-004.md#canonical-b2f60362640b0a449e9fad0c79a1c6de61222b93646210653808449f4d45f4df)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip](resources--cdn_loadbalancer--reference--group-004.md#canonical-5ce0c40f65cde0182bb73c440f96f92ab576c1da1ddbd22cf7cc1bdc364c98bc)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list](resources--cdn_loadbalancer--reference--group-004.md#canonical-94fed78ee5fb76728728a9c02680b6d113eeea08ea430c3eb16b3732731a365b)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-cec1d71d12f28cfe48e9e2a317ba37066aded4238e8e6e1078bac5ed78fec4c9)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector](resources--cdn_loadbalancer--reference--group-004.md#canonical-98d37ce5e21534888452ccaf22e11c09cbd7614e69f736872027e464bcf018bc)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-77e8d7cd8d6c23924d55df93295322d7d22483df3fe12dae81613e1ca1abb1b8)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list](resources--cdn_loadbalancer--reference--group-004.md#canonical-e124c5bd13d90d69bc81aea8e07e5eb4334d2cf305a110e05b6983dd8d5ef101)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list](resources--cdn_loadbalancer--reference--group-004.md#canonical-83594fd7ad25439a1b7854936b93d8749f72926e26d177707a10508861d137c5)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9676752f2f8543eb87055c3eed15bec800f2fd459df5485c1280b35c1fafe53d)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-b2f60362640b0a449e9fad0c79a1c6de61222b93646210653808449f4d45f4df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5fbbd875bf44e941c65814ed3f0ee06a7c1fe958e6663d990acebfa767160a80"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 2aa5fb068283 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client

<a id="canonical-94ada38c59ce331241f680fc15f6e3cc0281422e4341243565cb6f8237189f10"></a>

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
any_client = {}
```

<a id="canonical-f89a7d1951d71a6126c603f8426bc8bd0280a1aecc320a542bcd90bc8c075a38"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 2aa5fb068283 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7ac4241e916b2afd75a39407cbca71a8de47a4c92eecf6676f55f0da82a68c69"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 2aa5fb068283 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-5ce0c40f65cde0182bb73c440f96f92ab576c1da1ddbd22cf7cc1bdc364c98bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d1b253dca7e99eb1a0a4286755a4d810c27157068e191e2ce8174ce35f3814f"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 65dcb1ee9923 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip

<a id="canonical-31cdf5abdfc26a2b7751009b60b28a1f6443fb95f2b2b124113d50c7db72b7f4"></a>

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
any_ip = {}
```

<a id="canonical-887c6eb5ea8ed0e2cf619784dd1c0727148d00f92971f571e0276531230c7fa7"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 65dcb1ee9923 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-358f3a1a3969877fb3755a990286d6517029c6340b8b832b8542909edf7f6dbe"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 65dcb1ee9923 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-94fed78ee5fb76728728a9c02680b6d113eeea08ea430c3eb16b3732731a365b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ecff5cbb9f424e6f067b47f26dbfa83a13cb1c970c83915804829b25659b862"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / ed1f5f585027 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list

<a id="canonical-5963fd46996b4e3a233debe54187729c01eb1f632806de407901751c8d5eae37"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-b62f8cdde9799c833c367fc86d2a4bfe90ab687f24ce61b8b3e282dd2476e3a2"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / ed1f5f585027 / 3

<a id="canonical-f3f0a8833988ae28d725f8f9a0fd14a3388809fc236e5b1b45224ab571efd7a7"></a>

<a id="canonical-1a9b26ebdd4da2a298960d66aa7e23b3dd7e1f45cf4fb53cbcbe8176b4a27c5a"></a>

## as_numbers property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / ed1f5f585027 / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-22f9ed2976c7b9057616d009a8c3783926d7b3835386a72d84b3dd202b3c5dca"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / ed1f5f585027 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-cec1d71d12f28cfe48e9e2a317ba37066aded4238e8e6e1078bac5ed78fec4c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0d36cbd32a41919230af82dae1f15bd235f98a789b27e0cba21a07bea1a22a5"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 3510c332df46 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher

<a id="canonical-10353b9b652167df4d4f5b63347c3227070f34432ec96b8321545a20b8b4d25d"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0a47e9a80672b8142c8ac9c47ad0a5281c66e3e5129c26029a07e55b50ee0f44"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 3510c332df46 / 3

- [asn_sets](resources--cdn_loadbalancer--reference--group-004.md#canonical-132b982194eb24a9f47e36065ae30bb56e777e04a4ec9b14e5710f91fc326ef8): complete subsection reference.

<a id="canonical-cb412b04bea7bd9945e1eb8985b0c639d5df5ad9a3380b9d030631d63a9e8871"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 3510c332df46 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets](resources--cdn_loadbalancer--reference--group-004.md#canonical-132b982194eb24a9f47e36065ae30bb56e777e04a4ec9b14e5710f91fc326ef8)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-132b982194eb24a9f47e36065ae30bb56e777e04a4ec9b14e5710f91fc326ef8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1971c2be3fc9dfe02da956dbda404c1e1e23332da00a0b3a9a47d6d5c45ee2e5"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / acd680accf5b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-cec1d71d12f28cfe48e9e2a317ba37066aded4238e8e6e1078bac5ed78fec4c9)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-b652dc410a0c453f13a80e2c3183cc434d8058a5573ac99be294b3eb060e5584"></a>

Type: `"object"`. list nested block, Optional.

List of references to bgp\_asn\_set objects.

Upstream description:

A list of references to bgp\_asn\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-2527202971e1511638721a755fe8b0cb4761780f0487c4d1f9195430ae5ecaf3"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / acd680accf5b / 3

<a id="canonical-b7537b51c02c37e64da6888bca49fc33345ea3eb72f509303c323a8a381b969e"></a>

<a id="canonical-5a4ce6f73acd9897d79249411c7961bfb8434a8a89b044afb159e35e02160d82"></a>

## kind property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / acd680accf5b / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-e06370f9f83b8ea48332e02f1112619147629bbc487079004678e73072b03ab7"></a>

<a id="canonical-194b2c7b139a7a9bf56615e4a9f1a76ac51aa7b1579d39498a0484a395c25c81"></a>

## name property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / acd680accf5b / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-543d131a0e385f48488187f9601a76107a4695c01f48b6770c5d0a1ed2e52613"></a>

<a id="canonical-b0da929c21cde4c0b6c3369cb3f40bc23b7068fc55b76990fe0fcff08dcc7291"></a>

## namespace property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / acd680accf5b / 6

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

<a id="canonical-028fccb8c265b3aa8f8db20c83472e084598f043baaa660a1839241165337fe7"></a>

<a id="canonical-f63bf41b8420475c8731e36f9c40950357c775dfa361e688a28ac47c1f619f29"></a>

## tenant property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / acd680accf5b / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-484ab2c34d5178758e934a4a47d15f5d20bb27266b784d7b192d0a91a0b3c183"></a>

<a id="canonical-f4d021018a318a7abf0595c26605bd6942dca6f8457870a2f1d5f75f21314d20"></a>

## uid property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / acd680accf5b / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-e256e7a653f70e4abad8ff12fb0f0aa4792ce516f6dd10e4a1dcc2e68519982e"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / acd680accf5b / 9

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-cec1d71d12f28cfe48e9e2a317ba37066aded4238e8e6e1078bac5ed78fec4c9)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-98d37ce5e21534888452ccaf22e11c09cbd7614e69f736872027e464bcf018bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a15228bb5419fae324021a0177d5cc0d5d887314cfe095b99ba7523a78042e3"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 26babc7e5a77 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector

<a id="canonical-79d6167ff31e002b44fae0a9860e7dfdff6622e8744207a25a186f7183fcede6"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-4d13443738aee8b248b3cfd91234cb010ae05a65fdf66b385ad2be72fe7b390a"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 26babc7e5a77 / 3

<a id="canonical-63e2c5240510fce4acf11d207e16d32a4a56251a3a3249087aa52c592692b094"></a>

<a id="canonical-e4796be113da5475afa724e0126ae772c9f5b64a34f78bef28f0160c808dd83f"></a>

## expressions property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 26babc7e5a77 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-dfc206aba127c9657ce0c02b641fb30ead6b838af04313a4407a18612b376e12"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 26babc7e5a77 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-77e8d7cd8d6c23924d55df93295322d7d22483df3fe12dae81613e1ca1abb1b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01d76f9e6e02585576bd4ac094092fb877a8147d19978e8c831de8ed7aa0fad9"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / fd65eb2e32fe / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher

<a id="canonical-c73d35866c2e2e65e1548412c9f11da083f5c9c90b4b943862ebf7287d99abb6"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-c463597b71c15455839468223166cd59cd756e87dcc6cbbb0af5fb88c63ac9c8"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / fd65eb2e32fe / 3

<a id="canonical-d5cc3c80e6e6e9406e6f31a103a99efaf639d80add71cab7695bab6327175d37"></a>

<a id="canonical-54274bb41f9fbae10e55966de3cc74cc38edb50b04df6c84b44a13d5bcc2e64c"></a>

## invert_matcher property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / fd65eb2e32fe / 4

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [prefix_sets](resources--cdn_loadbalancer--reference--group-004.md#canonical-6f31363ed08585eef02532187d19ddb52bbdc1a473d964ac48182ac252dd25f7): complete subsection reference.

<a id="canonical-a352781bf1488dd53bead52a010486649c963ffcacd54204d1612e36dc22c2a5"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / fd65eb2e32fe / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets](resources--cdn_loadbalancer--reference--group-004.md#canonical-6f31363ed08585eef02532187d19ddb52bbdc1a473d964ac48182ac252dd25f7)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-6f31363ed08585eef02532187d19ddb52bbdc1a473d964ac48182ac252dd25f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eed6cac2c854dad9c457a071ea1d85a664ce12eb00b9e37407d33b5dfc08ef83"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / fb36456ffb7b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-77e8d7cd8d6c23924d55df93295322d7d22483df3fe12dae81613e1ca1abb1b8)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-8748cec08816c9fa624999e87f15c197c83d9b781742c53cd4c79bb3cf06f3e5"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-926a6292d104d03f78b65bf80d46b7b95faf43353da66129a55f6212ad317190"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / fb36456ffb7b / 3

<a id="canonical-c8c1e5b91d0e4c786eadf720db387998abf7cfa42dfe8aad8b6dbdcbd0a33931"></a>

<a id="canonical-57cd6abc4fd206e0ffb4c7be1cddeb2b94555345c17268e61536fb6524690c28"></a>

## kind property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / fb36456ffb7b / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-a7c304c19d5ebfed984270e8775d74fcfb7fd9bf9a6b4827c81608810a1ea6ff"></a>

<a id="canonical-7ac4b77877b7005ae8d32cf233a27970d8b2d06a7e22a6552b1f4c280a17528b"></a>

## name property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / fb36456ffb7b / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-d90f3760f6a4fc42ef1898d01cf893be66512e1bc011428f7a2c9cfee71a0fe4"></a>

<a id="canonical-80333d9ced0997e30cde8274f02b93ea0af95efbbaed397d3c729b09227a8314"></a>

## namespace property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / fb36456ffb7b / 6

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

<a id="canonical-a45922789138542ab5fbc2c6a6b0043bf9f21904bb8e109a811157f7a68ba790"></a>

<a id="canonical-b1d9eb04f6111c849062171a2b9f1667779d733e3b74b0b34438357119b7a8db"></a>

## tenant property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / fb36456ffb7b / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3a602186446f95b2fb8378ba36abf645e55e2b1a7dd4b6d846317e472f4b2962"></a>

<a id="canonical-579c42c4296302a582536982cbf50f6ffbfdc38bdb6e48390abb581627b20e80"></a>

## uid property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / fb36456ffb7b / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-2c03f947a48586933173a05a645082839b0335b99f43e4ce942defeb1bdec370"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / fb36456ffb7b / 9

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-77e8d7cd8d6c23924d55df93295322d7d22483df3fe12dae81613e1ca1abb1b8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-e124c5bd13d90d69bc81aea8e07e5eb4334d2cf305a110e05b6983dd8d5ef101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-614f182df6be1327edbe02f39f5c06fd792dfb1eb28d9f441cda8af281f6d4d1"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / b6b794bc4fac / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list

<a id="canonical-e6e0f31f57d64a38809674052a7b79a8ed73f467d0df825bc0728f41f1f92852"></a>

Type: `"object"`. single nested block, Optional.

List of IP Prefix strings to match against.

Receipt-pinned upstream constraints:

```json
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
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-c31b1a3bfa656e9d2c806e06ae2695cbb997c2fdbc0590672cfac49fa285c80f"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / b6b794bc4fac / 3

<a id="canonical-15b4fd8a54b4c64b9380f43b3f3639851586864d76ac335440a87c46e5f6b974"></a>

<a id="canonical-d1e32b70011e867883f10ffb9c94a8cfc5f49c731d52cba9baaf42e18d02ac98"></a>

## invert_match property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / b6b794bc4fac / 4

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2606f49742bd1d3684704c791c85c41cf77938bd6c30ff96e22266b4a70c885b"></a>

<a id="canonical-0fcf031d624b90ab4832217eb07eb3913ceeb285bd33d8b4dde4c6adc4653f34"></a>

## ip_prefixes property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / b6b794bc4fac / 5

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-80549e3cd47ad8e806dd4ffcf80e161023827e4796b21bfb628b908e6c8d2e3a"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / b6b794bc4fac / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-83594fd7ad25439a1b7854936b93d8749f72926e26d177707a10508861d137c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2fdbb55d008dd05597435cd31d1ec098d99c949e16899d44c841f98bf11cf040"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / e54205dbc880 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list

<a id="canonical-56d3a4049b40302c58207b6bba81ebb14712e63259866840ddbf5e84ba206061"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List Type. List of IP threat categories.

Upstream description:

List of IP threat categories.

Provider validators and defaults (from schema source):

```go
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
ip_threat_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-196da33beb5c4889bf750fe9abecef4e90e9fd79792527d4d2c5d941584f67b0"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / e54205dbc880 / 3

<a id="canonical-b91e2d271bb631fbcbb38a0d48f07834d6d653cdf8cc445d8fc5d2679f197d51"></a>

<a id="canonical-e27d8483d6b265a9fe7aedcc8b91f7ab48b8bbf7dc24240a4b1e712591471d22"></a>

## ip_threat_categories property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / e54205dbc880 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

Upstream description:

The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-82784e283f106903fc384449a6308f7dc8beba6764b578a05ceaee9b50d929c7"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / e54205dbc880 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-9676752f2f8543eb87055c3eed15bec800f2fd459df5485c1280b35c1fafe53d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0dea9b38d06bd20a5972849bcd5aa40597915e675215a361fea4628f35efa3a"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 80ed284c1962 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-1f0bd90b95115ad0de9a424a2b3fe11ac3599189de13d14efa17af58075427b2"></a>

Type: `"object"`. single nested block, Optional.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

Receipt-pinned upstream constraints:

```json
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
tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-08cb7b7288ccc921e4a5599307a5240aad7a10f468c056c6be78cfeb20339e39"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 80ed284c1962 / 3

<a id="canonical-44e600e6c8b5e93ab6647ab45f24721198b35f50edbe383a43cd74f68a7293a4"></a>

<a id="canonical-19f549980f627979864005c4f24738027c2238a17a3335d639d1f8c9b9f7b7dd"></a>

## classes property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 80ed284c1962 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-94c011bdb39f617362b8d46dbfd85830843c28a0d32e6a0ca6aef89553a49bc6"></a>

<a id="canonical-eee02fa7e6390df264faceb76791ca6932ac65da522f5c3075b12085a8c7fe0b"></a>

## exact_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 80ed284c1962 / 5

Type: `["list", "string"]`. Optional.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5c9d0fa929155783af4223594aaaf433d70ddca368be3ace2c9fbc2e32d94b80"></a>

<a id="canonical-2c084e99fb2a5335c39ca06ea3db6dd8cf05e5e4859bc1b5385c072510e8046c"></a>

## excluded_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 80ed284c1962 / 6

Type: `["list", "string"]`. Optional.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-67e43553cb8ac6a6185f1c01096bd92fe9649173994a530c2be2fc1b6d401857"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matc / 80ed284c1962 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-9ada8ff9def97ac368ec33c908d92f986568fc1a577071c9ed0715c97ac6fad6)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a78b83c0767486ebef7477f65d348810ed18b5fb75412caf41db7a243ac9a999"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / b87adf3f0709 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher

<a id="canonical-a65ba1411011dfeaf7e980e1eba92f0bc1a60b09e4a72a0a35c349bfa513c012"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for request matcher.

Upstream description:

Request conditions for matching a rule.

Receipt-pinned upstream constraints:

```json
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
request_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-38f6947924c9058c870c10dc3d9af5dee9e67975ca3e4dbaa0e9bd09a1aab29e"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / b87adf3f0709 / 3

- [cookie_matchers](resources--cdn_loadbalancer--reference--group-004.md#canonical-ec9ba755287bd5659cc142113336bc46afb0a317fa4c177c20454e6af7760fbb): complete subsection reference.

- [headers](resources--cdn_loadbalancer--reference--group-004.md#canonical-d2695ed2ff15af01e3948089f5fde961715f347c0897b03d5e2328021e2ef633): complete subsection reference.

- [jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-6a34a0029ea207a973fe6ed5021fdc98ec7aa5a4895872cf62b6abb80bf609e5): complete subsection reference.

- [query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-7a2316de11baf30a1ba28b16646eee1007ed8c965c69abd3e391e55b4b0558bb): complete subsection reference.

<a id="canonical-7b8fb3c93e0aa4b700805e9fccb6f3773a6aaf4d71471aab3e766e73b08845d0"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / b87adf3f0709 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-004.md#canonical-ec9ba755287bd5659cc142113336bc46afb0a317fa4c177c20454e6af7760fbb)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-004.md#canonical-d2695ed2ff15af01e3948089f5fde961715f347c0897b03d5e2328021e2ef633)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-6a34a0029ea207a973fe6ed5021fdc98ec7aa5a4895872cf62b6abb80bf609e5)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-7a2316de11baf30a1ba28b16646eee1007ed8c965c69abd3e391e55b4b0558bb)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ec9ba755287bd5659cc142113336bc46afb0a317fa4c177c20454e6af7760fbb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5041c511a5432c6819c16a49aa0f25bf7d688caaa7b17ec93dbe8ad765ca3a1"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / a1bc3babd8d4 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers

<a id="canonical-91f4ad1c9cc45ad06f0f644e846669a28e7511d099af04ea67759cb57109b107"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-84da1d3361b7b8a484eb519754899e98cf501d98dcb6653482d5cdc813f31305"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / a1bc3babd8d4 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-f4794a625c69f61b5c5aefa228040375e3059d44f8e140217357d6dcecdf3e0b): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-ec936de1e72ca7a9d541b92d158d7ec5e447dba158233f6f7f29087eaa64f71b): complete subsection reference.

<a id="canonical-751885cd372c4044d8b0b50e2238a123b989f6238407eb43da005c93f0a607ad"></a>

<a id="canonical-e63249d829824ee947c0a4a22b7c0fcd6dc8e526d94f1d4c2d6b25efc5679de6"></a>

## invert_matcher property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / a1bc3babd8d4 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--cdn_loadbalancer--reference--group-004.md#canonical-603380cce8f42e6af9cf8985c87cddc37ed60e77d138851ea39d6039082f0047): complete subsection reference.

<a id="canonical-b031fdff2dedb3cebb8fa0d92600b25a89ccf6c81b24f5f0172103a61599dc83"></a>

<a id="canonical-f60d820a754c1981b24721a5b59c2cc738fc1031718ab0490aa28422147dfcce"></a>

## name property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / a1bc3babd8d4 / 5

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-91db300ac0e73c88bece4dac04c21e13a219d6863f0aaa7a346f4cc1bd0113a2"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / a1bc3babd8d4 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-f4794a625c69f61b5c5aefa228040375e3059d44f8e140217357d6dcecdf3e0b)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-ec936de1e72ca7a9d541b92d158d7ec5e447dba158233f6f7f29087eaa64f71b)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item](resources--cdn_loadbalancer--reference--group-004.md#canonical-603380cce8f42e6af9cf8985c87cddc37ed60e77d138851ea39d6039082f0047)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f4794a625c69f61b5c5aefa228040375e3059d44f8e140217357d6dcecdf3e0b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d5b055e693c06db5ae972c8100c3f51dca42268d1fd2d9997ad57d48d69328f"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 855660da0272 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-004.md#canonical-ec9ba755287bd5659cc142113336bc46afb0a317fa4c177c20454e6af7760fbb)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-af3e3b1541f2e80800805994781302eaaada0345c5b3f218be3cd4bb5dc80749"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-68568ec1fe150e2dd036985e647810f3953b8a68600553dc82ea2e146416454d"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 855660da0272 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c46dbfc91e5c990a4d47bfcabf7cc4104554882c4a4c2d61073abe05b59c588f"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 855660da0272 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-004.md#canonical-ec9ba755287bd5659cc142113336bc46afb0a317fa4c177c20454e6af7760fbb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ec936de1e72ca7a9d541b92d158d7ec5e447dba158233f6f7f29087eaa64f71b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f99e683b7bd50ee37cd9b515d574e12ef2c46f884ee8c819aaadfaa6c7a5522e"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 98aa2dcc3aaa / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-004.md#canonical-ec9ba755287bd5659cc142113336bc46afb0a317fa4c177c20454e6af7760fbb)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-df5fd5bf446f32fbb60331c1d4bc95f711fd2a39002fa66f8187e832fc555df5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-6d3c00110fbce316c2efb0ec3e17950a05ea1371ea2c3fbe4c518de13add4925"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 98aa2dcc3aaa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9ee53b6f5a54c0472775eb103ef3feaaf968e99b684943065ee6d15af0a91825"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 98aa2dcc3aaa / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-004.md#canonical-ec9ba755287bd5659cc142113336bc46afb0a317fa4c177c20454e6af7760fbb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-603380cce8f42e6af9cf8985c87cddc37ed60e77d138851ea39d6039082f0047"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd933bee777df11d3a3394684242d551e91926210d6de8b734d35dc25a843d26"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 44c0cd75e8fa / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-004.md#canonical-ec9ba755287bd5659cc142113336bc46afb0a317fa4c177c20454e6af7760fbb)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item

<a id="canonical-c1b99feda29db6c801d80876c30f6b16b55169e51a8c2f4076e8532e70f878eb"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-419f91c3e89e9a054032317b442df0be5647d6115867861206ba7e8a2763717d"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 44c0cd75e8fa / 3

<a id="canonical-1ba5ede5a604762d3453de4dca226ff74f01f85426bf972c8777f8e515b11858"></a>

<a id="canonical-2350b978a8101c2ccaf79de17735009ff47856d5985558331f1b19c103ab4686"></a>

## exact_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 44c0cd75e8fa / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
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

<a id="canonical-deaeff9237df06acd8de6d77918f9f1848713cdfab0f6fae77c21a771bb07ff2"></a>

<a id="canonical-31ce1a8452255d2476c76bc3f5ccad37dd4c42a753f47d172b9fb4c0c60bd7a9"></a>

## regex_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 44c0cd75e8fa / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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

<a id="canonical-b778f40e071d52ecfc1bfcba2e9c3b019e1049205d73c3c4f22b5d62fcd9d94e"></a>

<a id="canonical-edf68914d0bdcb02a6ba32d3dbb4ed3005052d2764e22420676aae8e8790148e"></a>

## transformers property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 44c0cd75e8fa / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-6ed80f3f6c6fa809ef7610115dd5139686b6dde3871c253b285b95d4f809b549"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 44c0cd75e8fa / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-004.md#canonical-ec9ba755287bd5659cc142113336bc46afb0a317fa4c177c20454e6af7760fbb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-d2695ed2ff15af01e3948089f5fde961715f347c0897b03d5e2328021e2ef633"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40492a238576eaea4bafaf6ff33c7f72698904bfba542295e7ba95f27c481eed"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / bcb0e952484e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers

<a id="canonical-8c947d432bb65182dca8fd845e319645f0fa434f5acfcc29999737ffafef34e6"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-6a3705910e8b5e93a5bb9b6583a571e726b3ad9eea217bc1602ff83bf4d39d0c"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / bcb0e952484e / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-c7f25f207e415df6ebc47f5ef0ed13ba135bfbc87d7991bf5a963e0a9fd4cffb): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-7e3e1a982176fd6737ffd69b9153b4461c4aa92f5f7ad4bc2d1e7125b4f2641a): complete subsection reference.

<a id="canonical-d89ad26980db6c9451cf7f8eb9fc9070e5f1a6ee2fc8d9597f6c0b36d437bf28"></a>

<a id="canonical-a8ae0dfd031e6e15d1e00a2bb4dcc0ab41c3946df6f738d3d37d1042bc95cc9b"></a>

## invert_matcher property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / bcb0e952484e / 4

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--cdn_loadbalancer--reference--group-004.md#canonical-47444ba010dafb1bebfa167d71a174dbd9cd28f4f0fd4ca06daef5db8cc6e23e): complete subsection reference.

<a id="canonical-4342356d9a884166b5577a85daca80bee49d95678fdd649a4e41b06deeb0439e"></a>

<a id="canonical-d1b761c7dfd9ebb1b31ca18570cf006553fbf8e1f73eb4ffc78465a19d4d490a"></a>

## name property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / bcb0e952484e / 5

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-4376110b8077ff5d9ba40ce313a43e31654d69b61aaed343b2d4ddd05a89d158"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / bcb0e952484e / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-c7f25f207e415df6ebc47f5ef0ed13ba135bfbc87d7991bf5a963e0a9fd4cffb)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-7e3e1a982176fd6737ffd69b9153b4461c4aa92f5f7ad4bc2d1e7125b4f2641a)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item](resources--cdn_loadbalancer--reference--group-004.md#canonical-47444ba010dafb1bebfa167d71a174dbd9cd28f4f0fd4ca06daef5db8cc6e23e)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-c7f25f207e415df6ebc47f5ef0ed13ba135bfbc87d7991bf5a963e0a9fd4cffb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a4f47a52991da4f85dbd2a7470a1506b1c2db27e08531786dcd5ac5e476a5c4"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / d388da9b9615 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-004.md#canonical-d2695ed2ff15af01e3948089f5fde961715f347c0897b03d5e2328021e2ef633)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present

<a id="canonical-5638be001c672c723bc297c405675339d85109404614daff4aca37a3abb7388d"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-55dc54c2768340a237b4ada3c38e617ddef73501f349f56ce7d7bc88cf8f55be"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / d388da9b9615 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2c83a830dbf217ff2bbab3a74756be2065018ffecb627ba48d0f509723a493ea"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / d388da9b9615 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-004.md#canonical-d2695ed2ff15af01e3948089f5fde961715f347c0897b03d5e2328021e2ef633)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-7e3e1a982176fd6737ffd69b9153b4461c4aa92f5f7ad4bc2d1e7125b4f2641a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f9f1b779d2dc836dd5207c461783f95d1c52e24739d49686c976d9ba374b951"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / ee9545dbcf10 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-004.md#canonical-d2695ed2ff15af01e3948089f5fde961715f347c0897b03d5e2328021e2ef633)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present

<a id="canonical-19f8f0b7b2a30f1e8a83d352ade8357cdd8df1904849e8ff033df078f1b04eb0"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-87b3fc35cae45f9659e51acc58e6decbb51a33283c36b48151d9c57460cc0061"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / ee9545dbcf10 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-84ce654608fa342452a0117799baa369c12376abaf147610fb9150cd4552b398"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / ee9545dbcf10 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-004.md#canonical-d2695ed2ff15af01e3948089f5fde961715f347c0897b03d5e2328021e2ef633)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-47444ba010dafb1bebfa167d71a174dbd9cd28f4f0fd4ca06daef5db8cc6e23e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4f35507839d153961ab63225d3739ca96c98ffa5788b8180d1c7a0a6599df30"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / a45900e97000 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-004.md#canonical-d2695ed2ff15af01e3948089f5fde961715f347c0897b03d5e2328021e2ef633)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item

<a id="canonical-e05c744c1ac111fb782903051e37064c0c7d50b140d65d4045d784a816691afa"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-3dbf28233de7cdccef2eb268d247e0e0a73e478ff7b20326b2c24d034cd567c7"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / a45900e97000 / 3

<a id="canonical-8aa7904eb494b3d69275cd1398cdc562d34a0680c7eea0258162ae8a41777d2d"></a>

<a id="canonical-7e71c32860f4d0e920f41e10babca14f04fd980d34b22e78d93f5fd42b5bbaed"></a>

## exact_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / a45900e97000 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
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

<a id="canonical-b9c4cc4da50ed40eac0169319478143405c947a4adc9023bc18777bc99f74d47"></a>

<a id="canonical-5af80afad79e2ab241ab58e099f7a84e45334a246ef83b9e17d34747e5b4306c"></a>

## regex_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / a45900e97000 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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

<a id="canonical-18da751aeebd8a5a8e7c240294ad092521878b4aa28d31e65e0960b655682ee3"></a>

<a id="canonical-ac791852d3d373cd6385903f9e1dd9dbdb1a6d73df86d169f9b21d2c1b0b4849"></a>

## transformers property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / a45900e97000 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-18bab8fc6aea0f52bea4c79b16a14dc07d1d91b91194b7cd2b421e1b63d22bd4"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / a45900e97000 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-004.md#canonical-d2695ed2ff15af01e3948089f5fde961715f347c0897b03d5e2328021e2ef633)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-6a34a0029ea207a973fe6ed5021fdc98ec7aa5a4895872cf62b6abb80bf609e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-262b17c66693dec3bd73c171b1d5d3cb4566d8bcded9999df6695e0d189ffb7e"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 1f03cfa6dea4 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims

<a id="canonical-f2dd1dfe3f21f266df1a14a95b297ae4637a4452ab4748580cab7c283a5d3f50"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
jwt_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-bc1932dd6a7261604cedf3a705994ed61e00b52d78091984fb0648dc4215cd88"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 1f03cfa6dea4 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-12f6d3c1300933348b84bac8ae223372ee907f99e3e550d0e3d945fd9f925254): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-0b24c8224675b284e18310d5f461fd763d45a402b10ce01f470fe58339e8d0c6): complete subsection reference.

<a id="canonical-883ec89faa8f108bf76d781c347c43bfa94a75a781e2a14a77afe0d7bd6a2a84"></a>

<a id="canonical-30f016ea3f9b59af2b5bd7f0351974293930a4a1a13d529c20e2efb8490eb485"></a>

## invert_matcher property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 1f03cfa6dea4 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--cdn_loadbalancer--reference--group-004.md#canonical-fc1f197791fb9f6b90f8f9e24c69923a2d35b1502cbba935f328cbbd6d9482b1): complete subsection reference.

<a id="canonical-144e20fc43d9844f10f8b81b7cd012b6f385cc1f9fe86355a00a59fd1552e541"></a>

<a id="canonical-2db64f3ec70cf4acd6f169ae7579c238aa8adf97f4e9db4760e6c229cc7ba4ba"></a>

## name property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 1f03cfa6dea4 / 5

Type: `"string"`. Optional.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-4b68070c011e811d57b569666e93ecc81d5dea1d59167234f9867397f1954054"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 1f03cfa6dea4 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-12f6d3c1300933348b84bac8ae223372ee907f99e3e550d0e3d945fd9f925254)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-0b24c8224675b284e18310d5f461fd763d45a402b10ce01f470fe58339e8d0c6)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item](resources--cdn_loadbalancer--reference--group-004.md#canonical-fc1f197791fb9f6b90f8f9e24c69923a2d35b1502cbba935f328cbbd6d9482b1)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-12f6d3c1300933348b84bac8ae223372ee907f99e3e550d0e3d945fd9f925254"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9da2f892340ddca425a7f271b80eac0c75ee500b4b63ed52d227bcf875aa7df8"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 496c41388857 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-6a34a0029ea207a973fe6ed5021fdc98ec7aa5a4895872cf62b6abb80bf609e5)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-4abb1575c74ab545418939aab55e0ad718285004cad5d32be411207f7f0344be"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-405de653d222535778e2f554c32101cad4916dc73cb52da97e610474273ee49f"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 496c41388857 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bdb9e59c17508402f1c2043a7061501aabea371ce46a1276788b45f1163baed0"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 496c41388857 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-6a34a0029ea207a973fe6ed5021fdc98ec7aa5a4895872cf62b6abb80bf609e5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-0b24c8224675b284e18310d5f461fd763d45a402b10ce01f470fe58339e8d0c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8634d00854bfea0ceeca4f4e2d40b04d3cc6467d1073c20950e70dfa2c5f7248"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / af053cea3cc5 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-6a34a0029ea207a973fe6ed5021fdc98ec7aa5a4895872cf62b6abb80bf609e5)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present

<a id="canonical-2deaeb503007af9103683bcd0d23fdaaa883658d1d112e58c0c7ef234259f0ca"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-3380d6327e18232c65bd26bc430050214cb659c91a50df224daa0162701ccd9f"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / af053cea3cc5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e3b23729aa2fdabeaaa339205f8b9fa5ca07bd6bb5f0a0dbed863bd4b2eada8b"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / af053cea3cc5 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-6a34a0029ea207a973fe6ed5021fdc98ec7aa5a4895872cf62b6abb80bf609e5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-fc1f197791fb9f6b90f8f9e24c69923a2d35b1502cbba935f328cbbd6d9482b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f4b80b1b7627f2815076a1ae0adc5d827cecd83f58a427c29af5756492b5547"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 21c2998465dc / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-6a34a0029ea207a973fe6ed5021fdc98ec7aa5a4895872cf62b6abb80bf609e5)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item

<a id="canonical-cac8dafd416893be211412b9d60b3072a6181e43453c686eff869044854e2e8e"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-b2d3ed1ca7626405312fedd6f612cfa6131fa2aba5cc7cd4de9352842303fe3e"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 21c2998465dc / 3

<a id="canonical-2bb52b4e737678c2f9d4400de0fcb9c0f4fdbb9fa68b88c88754c6f1f61c037e"></a>

<a id="canonical-6f559bb264b098b5561a20e18af2fada7b40b07451143488ea0dd60bf0b9e991"></a>

## exact_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 21c2998465dc / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
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

<a id="canonical-e9fad3c79e38d376b5ae870922cfab0969e937bb645f50152ab08d4ee65a1067"></a>

<a id="canonical-0028f3b245a043592dc40fbf202a4c32808a6c6fa3f9090e9b8b0e4ed75d7d45"></a>

## regex_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 21c2998465dc / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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

<a id="canonical-328e5777318fedb89d802c2ef5eafc7b1c1fe1c001f961f7d9dafd0988b489e3"></a>

<a id="canonical-59c092b9a564aa79db9a7d90e80b2c9d2a09b65b08c59dc95b1323171c2eb3d3"></a>

## transformers property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 21c2998465dc / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-62fb41ea4cc2f10a036a65f30e3eb2654c00242d12bbee720f17e1c925cecb56"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 21c2998465dc / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-6a34a0029ea207a973fe6ed5021fdc98ec7aa5a4895872cf62b6abb80bf609e5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-7a2316de11baf30a1ba28b16646eee1007ed8c965c69abd3e391e55b4b0558bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2563feb579c1ff7501ab1172e35c06134846d1cfeff35580e209afb8820ddf31"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 75854c38c642 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params

<a id="canonical-347737982b22686a1094ba04462acf5b7f6d926a15216386169c6f67f146edb4"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-cc0fdbdaae457a59fdc9ebd622f4f4f1fce02c604e5f9209221c6b67090d41cd"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 75854c38c642 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-3ce5d48c0f0876a32dc3b50830b60a53947a7d17f1090af8c01d9e4535a433a9): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-5f1a5a7470e8c92094e2d27235505e9f60c1e45b8d638347e6b8e4cb3f01144e): complete subsection reference.

<a id="canonical-491018119f2912ea6ca2d12f1f8dd7dfb1c7408892472b0ffad47f5a60cae9a9"></a>

<a id="canonical-a980c83f835d62e8344058bc075b49904f6185dc342f086241aa9edf7de10366"></a>

## invert_matcher property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 75854c38c642 / 4

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--cdn_loadbalancer--reference--group-004.md#canonical-408a619113fba5ed9de634f613954addcd984ebd6316b3d38411881e6e7ed054): complete subsection reference.

<a id="canonical-ac9132033d6ce992c91fbc9b97700cce454d3b64265484106e2a7d5e6a6c15d9"></a>

<a id="canonical-659945925f06ea6e83773417d9a433829a4b21fedb5c0720fcad5cab4c021518"></a>

## key property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 75854c38c642 / 5

Type: `"string"`. Optional.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-67f6dcef55c5124d91226e48ce66fe4b7e3bbb4a62edae3d9e3179a575e57e3f"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / 75854c38c642 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-3ce5d48c0f0876a32dc3b50830b60a53947a7d17f1090af8c01d9e4535a433a9)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-5f1a5a7470e8c92094e2d27235505e9f60c1e45b8d638347e6b8e4cb3f01144e)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item](resources--cdn_loadbalancer--reference--group-004.md#canonical-408a619113fba5ed9de634f613954addcd984ebd6316b3d38411881e6e7ed054)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-3ce5d48c0f0876a32dc3b50830b60a53947a7d17f1090af8c01d9e4535a433a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e05a064e13951f26ebbfcad74874e029b81270f98a9079c8b5a9eac062fc2b8d"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / f6dfaa669331 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-7a2316de11baf30a1ba28b16646eee1007ed8c965c69abd3e391e55b4b0558bb)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present

<a id="canonical-1af085c8387453ac238d2b118266f2dc2d38d1f1c00aa3fb44dd47c687b24575"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-279e3cf5dacd18226fda092ea82d990f7a024b50930c66f4f9b3275577b5bd4f"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / f6dfaa669331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-979aa40b92996414231075c6b9edaa7a872f8826bca567fd7326d7ca5354d617"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / f6dfaa669331 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-7a2316de11baf30a1ba28b16646eee1007ed8c965c69abd3e391e55b4b0558bb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-5f1a5a7470e8c92094e2d27235505e9f60c1e45b8d638347e6b8e4cb3f01144e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-562fcdd27c3ffaf71837a44cea78cd026da25e7274bca83bc58c793413d805aa"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / dd402d35cd02 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-7a2316de11baf30a1ba28b16646eee1007ed8c965c69abd3e391e55b4b0558bb)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present

<a id="canonical-a4412ac6063003deb209f3bec04b0775588afe83c97053c42763e4f96a7851bd"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-78ce02eaa405f8797a2319f699aa9ae560c18ccdb9005dd4e0d91c62ef0ae066"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / dd402d35cd02 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-61c342cd0d528284dd110c941ed49f40307e85940e1a4eb2f4ab81d1c32feb91"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / dd402d35cd02 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-7a2316de11baf30a1ba28b16646eee1007ed8c965c69abd3e391e55b4b0558bb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-408a619113fba5ed9de634f613954addcd984ebd6316b3d38411881e6e7ed054"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f606f3a55c8db22d65657bfd5a43488b7ffe110619f6743ff1958f7088541db5"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / c8219f24254f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-b4677332387d4db9a5c6c9909ba1618697db0a597fc4caaa2101969639c5b3d4)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-7d2d60a3fdec5e32961cbf7739c5aae0b075abe4b41506a44564b09d7920109c)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-f5545020a54885c00292c353bf52409e8536d461b06f1c183847ad2900048d83)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-7a2316de11baf30a1ba28b16646eee1007ed8c965c69abd3e391e55b4b0558bb)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item

<a id="canonical-a1b114d2df613c48ada8b81ecaeddd7d0402a771282fd66cef2327087b82230b"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-c676c89f5889622b0c9ef864b60149cb458d60fb93feecad9c75bbcc0e34aa9d"></a>

## Direct properties — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / c8219f24254f / 3

<a id="canonical-3e6413a75844df8cc1a656e1c8597d38a18391149254868464c3fe46e977edda"></a>

<a id="canonical-944b44a07be685d96968bf4514c1f53fa9b41d0f3f6ade208a655f5528435097"></a>

## exact_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / c8219f24254f / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
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

<a id="canonical-43f06192c36f792f09b770a73134bfd4ae659d519e582c3bf5b6371d0e2fb86a"></a>

<a id="canonical-2a02cfc7ad5bed178adcfc6f22afcb0af87315d88c83ffa5748930970c40246f"></a>

## regex_values property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / c8219f24254f / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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

<a id="canonical-15a3530ef436213fbdb24905aad0233c2f11a0db32e8f2ad67bc76090bdf4a04"></a>

<a id="canonical-58af214442e511e73e5ffd2e6224c9f2810de9a8684f2b75bfd28081045d0a47"></a>

## transformers property — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / c8219f24254f / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-15a828943045ece222577868ac0c22f2bca11b650228e9e536144d53c739abb8"></a>

## Next pages — api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_mat / c8219f24254f / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-7a2316de11baf30a1ba28b16646eee1007ed8c965c69abd3e391e55b4b0558bb)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-bf19a4c115b91ec1bf931c1af2143aeac22dace353069021748e22ac840dedc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9173169e90ddf594a2e71d137ab22bdfde894f9027e725144d5ac68c3c173b8e"></a>

## api_rate_limit.custom_ip_allowed_list — api_rate_limit.custom_ip_allowed_list / 7aa11faca80c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- api_rate_limit.custom_ip_allowed_list

<a id="canonical-091d39f02eb4e27e186917249c6036461fa985ba10900e7da2c0ba771e2a3e29"></a>

Type: `"object"`. single nested block, Optional.

IP Allowed list using existing ip\_prefix\_set objects.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rate_limiter_allowed_prefixes")}
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
custom_ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-cdd28edaeccbf5e533d50fea9f8ad4b087f82eac76f4b999a5811206c63b0d61"></a>

## Direct properties — api_rate_limit.custom_ip_allowed_list / 7aa11faca80c / 3

- [rate_limiter_allowed_prefixes](resources--cdn_loadbalancer--reference--group-004.md#canonical-dd8ea50887fb1523c35516425ff8baf12fb6fb252a46102fd80406abee5ae99b): complete subsection reference.

<a id="canonical-0be44f0804451199235a9ff1ab161294c5e44c1d3c359c737224b7fb31ad1754"></a>

## Next pages — api_rate_limit.custom_ip_allowed_list / 7aa11faca80c / 4

- [api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes](resources--cdn_loadbalancer--reference--group-004.md#canonical-dd8ea50887fb1523c35516425ff8baf12fb6fb252a46102fd80406abee5ae99b)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-dd8ea50887fb1523c35516425ff8baf12fb6fb252a46102fd80406abee5ae99b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f45a7bd8a85c81ccdb325cb4ece39cf9af3577c70cbdc5585a6202ef45e3915"></a>

## api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes — api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / cc22da1797ba / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.custom_ip_allowed_list](resources--cdn_loadbalancer--reference--group-004.md#canonical-bf19a4c115b91ec1bf931c1af2143aeac22dace353069021748e22ac840dedc7)
- api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes

<a id="canonical-f6765d00ba3ede1331872dc2f9707aaac18afa014df63a86323cc4aafd819500"></a>

Type: `"object"`. list nested block, Optional.

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Upstream description:

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
    "minItems": 1
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
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
rate_limiter_allowed_prefixes {
  # Configure direct properties listed below.
}
```

<a id="canonical-06e8fda00f6642ad6559a4227f99029955c91ebf31a23e4ed854424464409479"></a>

## Direct properties — api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / cc22da1797ba / 3

<a id="canonical-e27099d03e9d5d9cc78de92902eaadc3115255cdb21712aa61f2060db5767f03"></a>

<a id="canonical-4958b4cbb9411171b3c6ae2a5953a11319d55dd4bbec706067eecc812df2589f"></a>

## name property — api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / cc22da1797ba / 4

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

<a id="canonical-da4221e9d9ee00c230f97a9ca1d233441ef2d0d5b25c27352838e979ea1a5828"></a>

<a id="canonical-4d6a4ca246b8cd4520bb277deaf8204672e48b29c5b05cd331be8d19da46e5df"></a>

## namespace property — api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / cc22da1797ba / 5

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

<a id="canonical-e3a2e444aa0fe902ea2c017907596ebc2bfe4858c1b0f77e85a2b81cc6ce7480"></a>

<a id="canonical-0af29b353aa9792b394d5d39c6fa0f568ed7bcf62ff606d7385c719e68617c3e"></a>

## tenant property — api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / cc22da1797ba / 6

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

<a id="canonical-7cf187a6052b077d4069b32e37cba87557a82e366eaac1720df84cafb846549d"></a>

## Next pages — api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / cc22da1797ba / 7

- [api_rate_limit.custom_ip_allowed_list](resources--cdn_loadbalancer--reference--group-004.md#canonical-bf19a4c115b91ec1bf931c1af2143aeac22dace353069021748e22ac840dedc7)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-5d427794f4feeabdadd16d86b5fd2bdef15819e2c3ad286911e2ecf14cfa66e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f920c7add552a0a16aa3826aca61eb80c4150258d23815aa717b9f74880f4c0"></a>

## api_rate_limit.ip_allowed_list — api_rate_limit.ip_allowed_list / 6d970c3d4b3e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- api_rate_limit.ip_allowed_list

<a id="canonical-97c1a90292c002ff3d2ec9492bea2d46e19da13630c437b74c1ef8fd60672398"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
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
ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-935eeae2339d083faeb1d780ace8d2a5a2ccd2cd8b17241a2d7f5575d242b0b3"></a>

## Direct properties — api_rate_limit.ip_allowed_list / 6d970c3d4b3e / 3

<a id="canonical-60aedc7bbe87e1c83c43e187584a59031afde97c70de78432bb7ba5805563817"></a>

<a id="canonical-60c29fd1696e6440995c02ee1ce35840dae88647fd7ce0d0eddfa3924efb190b"></a>

## prefixes property — api_rate_limit.ip_allowed_list / 6d970c3d4b3e / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-474ee47ac5c58cccb6980105b80b1a700cb0687ea779eef89f97865f9de9a7a3"></a>

## Next pages — api_rate_limit.ip_allowed_list / 6d970c3d4b3e / 5

- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f7df1ce35a4800d64ab75ae9727d8f056dc71f777bf9282354dea1e70b750388"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7753c1efc2bcc7c72de2c202b629dab0243365e87ed8fe29c681614ba969e5cc"></a>

## api_rate_limit.no_ip_allowed_list — api_rate_limit.no_ip_allowed_list / f04bdbecf376 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- api_rate_limit.no_ip_allowed_list

<a id="canonical-a5210261b3e8f2f27861bd1ed0dc04cca6c1e3aa28ea657bb0b1d71b2e0ed427"></a>

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
no_ip_allowed_list = {}
```

<a id="canonical-870d0e44ece7e79b22626d25c9e049ff55a4233e2be9335ead9673fc4c7c9c30"></a>

## Direct properties — api_rate_limit.no_ip_allowed_list / f04bdbecf376 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e7acf32b646e83c1487735f6345219a5ddf1fabda734b554688d3a6f4c64a603"></a>

## Next pages — api_rate_limit.no_ip_allowed_list / f04bdbecf376 / 4

- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c5ebb776e398ba4d89d954c4f9d83837703acba82876235a32d67a46a9e5b6e"></a>

## api_rate_limit.server_url_rules — api_rate_limit.server_url_rules / 953c941b4a8a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- api_rate_limit.server_url_rules

<a id="canonical-c82b6c3b079ffe0f4fd4e1a0eec3217da489923b25023b6fb06a90c0c519cb87"></a>

Type: `"object"`. list nested block, Optional.

Ordered domain or base-path rules for path-scoped rate limiting. Each rule must choose exactly one
rate\_limiter\_choice: inline\_rate\_limiter or ref\_rate\_limiter.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("base_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain"),
  validators.ConflictingListObjectAttributes("inline_rate_limiter",
    "ref_rate_limiter")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

Terraform syntax:

```terraform
server_url_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-6dfa4a68e13a2278004fd8626c633506bcd72d551458fd35a892df54352ef59e"></a>

## Direct properties — api_rate_limit.server_url_rules / 953c941b4a8a / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-004.md#canonical-5b23f59a9c745ba228b148111c69e2a61287f0efe8ab3767405999b17b9240c6): complete subsection reference.

<a id="canonical-e4088c1f1f810c9d5b2cf58f9c5fffd1ab6f1396da5105c3eacc438f51b558c4"></a>

<a id="canonical-14d982a0f73f445a025c4ea1593f1f2cc26786e7e586878f3029c5437f46a181"></a>

## api_group property — api_rate_limit.server_url_rules / 953c941b4a8a / 4

Type: `"string"`. Optional.

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with 'x-F5 Distributed..

Upstream description:

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with "x-F5 Distributed
Cloud-API-group" extensions inside swaggers.

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

<a id="canonical-ded9ae1f42e55177130a8e8b4b9f50280cc506673768aa4c0f7baece2274982c"></a>

<a id="canonical-c9317416adfacab9c19281db1fc4b348669a18995211af82b4205bbf90ba9fcc"></a>

## base_path property — api_rate_limit.server_url_rules / 953c941b4a8a / 5

Type: `"string"`. Optional.

Base Path. Prefix of the request path.

Upstream description:

Prefix of the request path.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275): complete subsection reference.

- [inline_rate_limiter](resources--cdn_loadbalancer--reference--group-005.md#canonical-393655f0c8b97058a5d7e5b2d287e80ff14bd3369485e1e2c4a9165d13f6d549): complete subsection reference.

- [ref_rate_limiter](resources--cdn_loadbalancer--reference--group-005.md#canonical-6455b887a64665e97c723c6bbd85c3c03ecc56318d08f0dee33becd1a7bc652f): complete subsection reference.

- [request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422): complete subsection reference.

<a id="canonical-29967c1a55fef493c98d15de0708df6c87df032d20b57140722adec9c7cb5be5"></a>

<a id="canonical-3dad9b1aa24ac6eb54bb62660aeeb876c33449ee72c39c558cf6090014878d89"></a>

## specific_domain property — api_rate_limit.server_url_rules / 953c941b4a8a / 6

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

<a id="canonical-f78a05a014e946bee45429562ebd2fa8a20141c8465963a40195b0fe8e808b52"></a>

## Next pages — api_rate_limit.server_url_rules / 953c941b4a8a / 7

- [api_rate_limit.server_url_rules.any_domain](resources--cdn_loadbalancer--reference--group-004.md#canonical-5b23f59a9c745ba228b148111c69e2a61287f0efe8ab3767405999b17b9240c6)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--cdn_loadbalancer--reference--group-005.md#canonical-393655f0c8b97058a5d7e5b2d287e80ff14bd3369485e1e2c4a9165d13f6d549)
- [api_rate_limit.server_url_rules.ref_rate_limiter](resources--cdn_loadbalancer--reference--group-005.md#canonical-6455b887a64665e97c723c6bbd85c3c03ecc56318d08f0dee33becd1a7bc652f)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-f58e4a2bf6386ee36e45a3c28c6b34ee61b7613b7b852544427ffe9250b32422)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-5b23f59a9c745ba228b148111c69e2a61287f0efe8ab3767405999b17b9240c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ff81bf0a42319e137ab507ef64a567ac3d011f0d2a9a3c48dc651ac22442bb0"></a>

## api_rate_limit.server_url_rules.any_domain — api_rate_limit.server_url_rules.any_domain / cbdf299676de / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- api_rate_limit.server_url_rules.any_domain

<a id="canonical-fbb1bfd5a82f2dfe63e94d429bf2439f26c881b9791597776f46f18889906348"></a>

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

<a id="canonical-50732914fb5e201f0383504318399b5c21d253d2db82400021d1c105e1e5da64"></a>

## Direct properties — api_rate_limit.server_url_rules.any_domain / cbdf299676de / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ed03c82340bd2ee84c6e6094e00ab93fb2a84e56e1150a8fdd9ec8cec8950744"></a>

## Next pages — api_rate_limit.server_url_rules.any_domain / cbdf299676de / 4

- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55e237690e0e9743cc5b2303cacbbb9991c42cda4a4e4a709876bf60ef7d2140"></a>

## api_rate_limit.server_url_rules.client_matcher — api_rate_limit.server_url_rules.client_matcher / 350aa29481f1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- api_rate_limit.server_url_rules.client_matcher

<a id="canonical-ca6c6c059383f5dc12c5ffa465c4aa7435bec32db2e7cc89ea3c4755a0f08c22"></a>

Type: `"object"`. single nested block, Optional.

Client Matcher. Client conditions for matching a rule.

Upstream description:

Client conditions for matching a rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_client",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("client_selector",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
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
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-ip_asn_choice": "[\"any_ip\",\"asn_list\",\"asn_matcher\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

Terraform syntax:

```terraform
client_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-1fcbfa8a751ecdf54980c7364c86921237e3fd00cb16c71bb2b77f12e2c0fb2a"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher / 350aa29481f1 / 3

- [any_client](resources--cdn_loadbalancer--reference--group-004.md#canonical-41675dc238f207ec7becd73ba3e237ed353b6cf806dcb7247c885ee951f6b0b4): complete subsection reference.

- [any_ip](resources--cdn_loadbalancer--reference--group-004.md#canonical-e541e674816c15bf242450728cc5cedb6e92152f35ff0b9035b6087aa5e12215): complete subsection reference.

- [asn_list](resources--cdn_loadbalancer--reference--group-005.md#canonical-c99d6513f352a985f593897b8f86f593471932f73c0fe005d016c01e19ba8c2b): complete subsection reference.

- [asn_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-db9b2bfe02531f7d46d5a91231b1c2bde4624d19ffe377c35dd9d0a14b023cfe): complete subsection reference.

- [client_selector](resources--cdn_loadbalancer--reference--group-005.md#canonical-0535fca52232cdfb0faf772dc41efee44bad609f21ce23324b156358dff54a8b): complete subsection reference.

- [ip_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-9c55db87686e1d62df013b59ba5fff3b2253597dd72994c171adeea78bcf40ab): complete subsection reference.

- [ip_prefix_list](resources--cdn_loadbalancer--reference--group-005.md#canonical-59e6fca4a1bccd722f194595f12b3708fcb25fa15a1d709fd0257483726bebc3): complete subsection reference.

- [ip_threat_category_list](resources--cdn_loadbalancer--reference--group-005.md#canonical-c041a8794d03af4f38420db7346f31412d98ac4b8c7710912d024310840cced7): complete subsection reference.

- [tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-1c0e7b8fa2c29802bd4ae73f2433ad98e3dfc9ca347233d65bf563e916bce200): complete subsection reference.

<a id="canonical-580f24d73b7e2f43722cbb5ee5f39d08a7573592d96d38e05a4dd4fef714c450"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher / 350aa29481f1 / 4

- [api_rate_limit.server_url_rules.client_matcher.any_client](resources--cdn_loadbalancer--reference--group-004.md#canonical-41675dc238f207ec7becd73ba3e237ed353b6cf806dcb7247c885ee951f6b0b4)
- [api_rate_limit.server_url_rules.client_matcher.any_ip](resources--cdn_loadbalancer--reference--group-004.md#canonical-e541e674816c15bf242450728cc5cedb6e92152f35ff0b9035b6087aa5e12215)
- [api_rate_limit.server_url_rules.client_matcher.asn_list](resources--cdn_loadbalancer--reference--group-005.md#canonical-c99d6513f352a985f593897b8f86f593471932f73c0fe005d016c01e19ba8c2b)
- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-db9b2bfe02531f7d46d5a91231b1c2bde4624d19ffe377c35dd9d0a14b023cfe)
- [api_rate_limit.server_url_rules.client_matcher.client_selector](resources--cdn_loadbalancer--reference--group-005.md#canonical-0535fca52232cdfb0faf772dc41efee44bad609f21ce23324b156358dff54a8b)
- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-9c55db87686e1d62df013b59ba5fff3b2253597dd72994c171adeea78bcf40ab)
- [api_rate_limit.server_url_rules.client_matcher.ip_prefix_list](resources--cdn_loadbalancer--reference--group-005.md#canonical-59e6fca4a1bccd722f194595f12b3708fcb25fa15a1d709fd0257483726bebc3)
- [api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list](resources--cdn_loadbalancer--reference--group-005.md#canonical-c041a8794d03af4f38420db7346f31412d98ac4b8c7710912d024310840cced7)
- [api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-1c0e7b8fa2c29802bd4ae73f2433ad98e3dfc9ca347233d65bf563e916bce200)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-41675dc238f207ec7becd73ba3e237ed353b6cf806dcb7247c885ee951f6b0b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f40b88caa63fb0e12a634db39eeb1eb3c6a36551f2b2eaa5a0d84307a39c8b1"></a>

## api_rate_limit.server_url_rules.client_matcher.any_client — api_rate_limit.server_url_rules.client_matcher.any_client / deb40fd83efc / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-9f90de04758943e80fca37f031ad0bf08d6b7cceea0026c2b70d9a707c5bbe68)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1ab5d122dfe6c6ac11012a142134afe6e97361b4f9529e5009c34b4d39a7f7cf)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- api_rate_limit.server_url_rules.client_matcher.any_client

<a id="canonical-db51546202067f86f42e294edb7fe06a9aa2df1e55badaa557d591e76d1beed9"></a>

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
any_client = {}
```

<a id="canonical-a824a9b7299cec5c0bc3671da83598ea2586a498e4b2edd54025e8226b64d1db"></a>

## Direct properties — api_rate_limit.server_url_rules.client_matcher.any_client / deb40fd83efc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e97412b2daf80054abaf9bba1d34b54d26d8ae2e844baa3b729e3a1f8e831cd1"></a>

## Next pages — api_rate_limit.server_url_rules.client_matcher.any_client / deb40fd83efc / 4

- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-b57675087227f85e31d2f1e413d8dfa0895ab0596b75681660d5755b97a89275)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-e541e674816c15bf242450728cc5cedb6e92152f35ff0b9035b6087aa5e12215"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
