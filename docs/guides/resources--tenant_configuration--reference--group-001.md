---
page_title: "xcsh_tenant_configuration reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tenant_configuration reference."
---

# xcsh_tenant_configuration reference

<a id="canonical-16136964bc77142e101e918b10e24c07bf5c7f99c8c5df90b4efd69072c8b3a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44e48ca50bafac76e56ab678ae0b8ec08eda5bd510bd4f5d11240ae95035a8be"></a>

## Property reference — Property reference / e3377aee4542 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)
- Property reference

<a id="canonical-25b288279442fabd2f5b7abe4ebad23bf0debdfa6318194cd7a67e91abb207c4"></a>

## Direct properties — Property reference / e3377aee4542 / 3

<a id="canonical-c94fa6f9b9d3553bbec8a0ab33fea4b594bf4361f3f281d8362458de6fef16a3"></a>

<a id="canonical-c60d36224e5a01f1116d30a2d58a7c5c3deeca7ff03ad7be39a5b216a99b3661"></a>

## annotations property — Property reference / e3377aee4542 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

- [brute_force_detection](resources--tenant_configuration--reference--group-001.md#canonical-821619c8d84bb2cefd992a812706b55e3c8c013ff89fcb52507934a4e8398408): complete subsection reference.

<a id="canonical-55bc511b2a2900f74ce2445d0861ee98b02e6f903a3811b4177b35aa6d81ed23"></a>

<a id="canonical-c01406e8240ff8090f9633c64f8b970b44c40adbe549054070dde7da50d4aa33"></a>

## description property — Property reference / e3377aee4542 / 5

Type: `"string"`. Optional.

Human readable description for the object.

<a id="canonical-890c9e81138b9b06c1dba91fe251caa0b56e90eadcd2db41ad111dddcd3bf62a"></a>

<a id="canonical-20b1e1fe0c28e56b89d0706b0cc230c054a7e86cf87380b93dbf6171afd3edad"></a>

## disable property — Property reference / e3377aee4542 / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

<a id="canonical-d962e3877806f18153aefdcd5e4995f643a08406e1916c62dc2c261f8ede82f1"></a>

<a id="canonical-fd541bd1f3b357dcabaf081a449a1c1d64317b163a43a5a35a81d3f2b69e980e"></a>

## id property — Property reference / e3377aee4542 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-85cad13c7f0b6edf84289d0562023989da419bf26fa1f279dd535afb5651e0d0"></a>

<a id="canonical-8b79f08329b7e82824e360d876cb41d611dd7ae818f96147329fda7cc0a51907"></a>

## labels property — Property reference / e3377aee4542 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

<a id="canonical-4d0fb22d12eafd58f1530eb700b03f1481b8a3a9691abc90e91ec2eb2255038f"></a>

<a id="canonical-ac122ee6e911429d8e3a696a59ad3a881fa0cd1600f4718edd3269889ae3b7f1"></a>

## name property — Property reference / e3377aee4542 / 9

Type: `"string"`. Required.

Name of the Tenant Configuration. Must be unique within the namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

<a id="canonical-afaf16d4901cd36f08ef44ae840872b25693aca2c2350b3703368f77fcae5b71"></a>

<a id="canonical-fdcfaf34d04e54416ecf9412519bec10078e155b59a024c8d0bdabc9e7c6e370"></a>

## namespace property — Property reference / e3377aee4542 / 10

Type: `"string"`. Required.

Namespace where the Tenant Configuration is created.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

- [password_policy](resources--tenant_configuration--reference--group-001.md#canonical-d22ac2c8491ce686c86e96f5f7e8666dcd3f0370f6f460088d0e6594a3c2a10a): complete subsection reference.

- [tenant_details](resources--tenant_configuration--reference--group-001.md#canonical-452a289512a12b40aecd668a1edafd36a78f77e364e271f9ea07d757fd73073f): complete subsection reference.

- [timeouts](resources--tenant_configuration--reference--group-001.md#canonical-2560c76f29b89582c64faf83495deba1b668a3386d8290a61ae8072cb0666184): complete subsection reference.

- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-c2388a7d0b4ae235a90c04846dd8c1fc6d3d1e30b66d38bc7a5af4139f3518ba): complete subsection reference.

<a id="canonical-a17e27657947f5a5e892bee50b3d73dba14c33868de863d934a566deda178840"></a>

## All schema paths — Property reference / e3377aee4542 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--tenant_configuration--reference--group-001.md#canonical-c94fa6f9b9d3553bbec8a0ab33fea4b594bf4361f3f281d8362458de6fef16a3) |
| `brute_force_detection` | [brute_force_detection](resources--tenant_configuration--reference--group-001.md#canonical-48fa94e6c975a8d147c4e7e08ed5b3643a3b30d88f2acfca9d2203c0f6b2a95c) |
| `brute_force_detection.max_login_failures` | [brute_force_detection.max_login_failures](resources--tenant_configuration--reference--group-001.md#canonical-189d398634de6d9d31c1db42800896299c218d60294eb7af149c15ad147ac020) |
| `description` | [description](resources--tenant_configuration--reference--group-001.md#canonical-55bc511b2a2900f74ce2445d0861ee98b02e6f903a3811b4177b35aa6d81ed23) |
| `disable` | [disable](resources--tenant_configuration--reference--group-001.md#canonical-890c9e81138b9b06c1dba91fe251caa0b56e90eadcd2db41ad111dddcd3bf62a) |
| `id` | [id](resources--tenant_configuration--reference--group-001.md#canonical-d962e3877806f18153aefdcd5e4995f643a08406e1916c62dc2c261f8ede82f1) |
| `labels` | [labels](resources--tenant_configuration--reference--group-001.md#canonical-85cad13c7f0b6edf84289d0562023989da419bf26fa1f279dd535afb5651e0d0) |
| `name` | [name](resources--tenant_configuration--reference--group-001.md#canonical-4d0fb22d12eafd58f1530eb700b03f1481b8a3a9691abc90e91ec2eb2255038f) |
| `namespace` | [namespace](resources--tenant_configuration--reference--group-001.md#canonical-afaf16d4901cd36f08ef44ae840872b25693aca2c2350b3703368f77fcae5b71) |
| `password_policy` | [password_policy](resources--tenant_configuration--reference--group-001.md#canonical-ca5900014c605370373120907771c6218532d4cb21e99d13d6ea35da0e60588f) |
| `password_policy.digits` | [password_policy.digits](resources--tenant_configuration--reference--group-001.md#canonical-7cba66b03d1fb1e5d6448ca8a3f8ec5d794c36f903ce1f4a253c562dcbc73068) |
| `password_policy.expire_password` | [password_policy.expire_password](resources--tenant_configuration--reference--group-001.md#canonical-b5b9acfb68fab7b9b791b59eb5be60af58f82176d4977a0b85d7c9456b05844c) |
| `password_policy.lowercase_characters` | [password_policy.lowercase_characters](resources--tenant_configuration--reference--group-001.md#canonical-fde746eb1abecdc040330a485d532bc0b14649301d5ae95f711cf17b06983bf8) |
| `password_policy.minimum_length` | [password_policy.minimum_length](resources--tenant_configuration--reference--group-001.md#canonical-6f09724d5a49f4f4cee40e6da8ec4efb4833ca37d400a2e1e89694c2c8504a5e) |
| `password_policy.not_recently_used` | [password_policy.not_recently_used](resources--tenant_configuration--reference--group-001.md#canonical-bc02e4de8199cba4df58d6de68aea9a063a7d84c82e4d82a3a9e85165a728044) |
| `password_policy.not_username` | [password_policy.not_username](resources--tenant_configuration--reference--group-001.md#canonical-694547d573ccf852309153f0a61f9b792dd637c513ec794a9ab013a7388c371a) |
| `password_policy.special_characters` | [password_policy.special_characters](resources--tenant_configuration--reference--group-001.md#canonical-35f3c0239e61050a34ae911828cd5e86dcbe624f976da146378e08812c591c23) |
| `password_policy.uppercase_characters` | [password_policy.uppercase_characters](resources--tenant_configuration--reference--group-001.md#canonical-f50dd904cf9b8fe8b362474d50496f5c703ebe313390007710cba5d146fe1a5c) |
| `tenant_details` | [tenant_details](resources--tenant_configuration--reference--group-001.md#canonical-ab19e16cb9b735c8a149c0869696a7ed8eeebe25c1784f220e5a12a51eb2f76d) |
| `tenant_details.display_name` | [tenant_details.display_name](resources--tenant_configuration--reference--group-001.md#canonical-8665e3f55149777bac2ad93219c4cd4bc425d44edd4ad35bde447cae9129129d) |
| `timeouts` | [timeouts](resources--tenant_configuration--reference--group-001.md#canonical-70b2fa7667f36ffb7b4b5b17d674654eb37e39146edff7e73430f16608e085b7) |
| `timeouts.create` | [timeouts.create](resources--tenant_configuration--reference--group-001.md#canonical-5f68907166c6207ee3fdcb2a4df71b43ec645d53b103172a0cc0c947964537b5) |
| `timeouts.delete` | [timeouts.delete](resources--tenant_configuration--reference--group-001.md#canonical-576465c0618990ea0946721e8de7306b3710ffef5e6d65dbbf72f8e899401f49) |
| `timeouts.read` | [timeouts.read](resources--tenant_configuration--reference--group-001.md#canonical-0c5735b6b4ab35b800a29d5fc09d5f6ae02ce5e7371cf7a914b2acc1ddf01e74) |
| `timeouts.update` | [timeouts.update](resources--tenant_configuration--reference--group-001.md#canonical-47930ed94ac1d3b1a8ab54f11a44032c6eb900bfceab575eef3336e2b773059d) |
| `user_session_expiration` | [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-91ac0514301375b0a02750bd5915a0f3db4dfc75de82dfe993541045323eb472) |
| `user_session_expiration.absolute_timeout` | [user_session_expiration.absolute_timeout](resources--tenant_configuration--reference--group-001.md#canonical-d10df19bb39dbbe0395e55c1cd57b1008f975f3b22b8a6115cc638b79da1ea7b) |
| `user_session_expiration.absolute_timeout.hours` | [user_session_expiration.absolute_timeout.hours](resources--tenant_configuration--reference--group-001.md#canonical-bf079e049dc4612363887688f8fe41c8019603d2a00f9b6812bbb2db0557de15) |
| `user_session_expiration.absolute_timeout.hours.duration` | [user_session_expiration.absolute_timeout.hours.duration](resources--tenant_configuration--reference--group-001.md#canonical-d6332dcbcfe284239d6d706fd04783adc542f5c53c3aeec261fae9d76ad69697) |
| `user_session_expiration.absolute_timeout.minutes` | [user_session_expiration.absolute_timeout.minutes](resources--tenant_configuration--reference--group-001.md#canonical-ff9011ad00654d15cba61cbe42f2ef89022699922ca3c75560c110156f3fafc3) |
| `user_session_expiration.absolute_timeout.minutes.duration` | [user_session_expiration.absolute_timeout.minutes.duration](resources--tenant_configuration--reference--group-001.md#canonical-f76dfe12793470231a9e16928511f900f5ce3e66313eb5d8c09f6fbc4f8e05f0) |
| `user_session_expiration.idle_timeout` | [user_session_expiration.idle_timeout](resources--tenant_configuration--reference--group-001.md#canonical-5c12ce410905f00420ddfd71d16183354e85f4e0c327a80c9d6033c81e8abfa6) |
| `user_session_expiration.idle_timeout.hours` | [user_session_expiration.idle_timeout.hours](resources--tenant_configuration--reference--group-001.md#canonical-693340cd5a297b2e7411a9285bd90b1f5831fba903a1e47f1a77c5fa83ac1f3a) |
| `user_session_expiration.idle_timeout.hours.duration` | [user_session_expiration.idle_timeout.hours.duration](resources--tenant_configuration--reference--group-001.md#canonical-cb6af5cc68fdaf381c8b51e7304c7a7f27763911cc43de5d0d9e1e3f21d93fea) |
| `user_session_expiration.idle_timeout.minutes` | [user_session_expiration.idle_timeout.minutes](resources--tenant_configuration--reference--group-001.md#canonical-473a4a67168c1711cf9882d7893c5050bc050eac27a6839d4a6fead1422e896a) |
| `user_session_expiration.idle_timeout.minutes.duration` | [user_session_expiration.idle_timeout.minutes.duration](resources--tenant_configuration--reference--group-001.md#canonical-892ec7d143dad9e846826752fe9064335a4f4ee6a11974d151c1d04475bd76ee) |

<a id="canonical-09972b6f649220b8166408895577bf251dae4d2f276f26c4a5b9f147bfd67233"></a>

## Next pages — Property reference / e3377aee4542 / 12

- [brute_force_detection](resources--tenant_configuration--reference--group-001.md#canonical-821619c8d84bb2cefd992a812706b55e3c8c013ff89fcb52507934a4e8398408)
- [password_policy](resources--tenant_configuration--reference--group-001.md#canonical-d22ac2c8491ce686c86e96f5f7e8666dcd3f0370f6f460088d0e6594a3c2a10a)
- [tenant_details](resources--tenant_configuration--reference--group-001.md#canonical-452a289512a12b40aecd668a1edafd36a78f77e364e271f9ea07d757fd73073f)
- [timeouts](resources--tenant_configuration--reference--group-001.md#canonical-2560c76f29b89582c64faf83495deba1b668a3386d8290a61ae8072cb0666184)
- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-c2388a7d0b4ae235a90c04846dd8c1fc6d3d1e30b66d38bc7a5af4139f3518ba)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)

<a id="canonical-821619c8d84bb2cefd992a812706b55e3c8c013ff89fcb52507934a4e8398408"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf76a4dbe76e65f982e072b1136206755cf1a1ffdb234b2f74db565f748140e4"></a>

## brute_force_detection — brute_force_detection / 2de426369f9d / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-16136964bc77142e101e918b10e24c07bf5c7f99c8c5df90b4efd69072c8b3a2)
- brute_force_detection

<a id="canonical-48fa94e6c975a8d147c4e7e08ed5b3643a3b30d88f2acfca9d2203c0f6b2a95c"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for brute force detection.

Receipt-pinned upstream constraints:

```json
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
brute_force_detection {
  # Configure direct properties listed below.
}
```

<a id="canonical-8691a3f03a626d22cb876701f44fe1bb0e94a451f397114009c8c8890af4f991"></a>

## Direct properties — brute_force_detection / 2de426369f9d / 3

<a id="canonical-189d398634de6d9d31c1db42800896299c218d60294eb7af149c15ad147ac020"></a>

<a id="canonical-b0ebd45163b927d9a05f979fcab5323380feef9c1d0efafebefe968c77491964"></a>

## max_login_failures property — brute_force_detection / 2de426369f9d / 4

Type: `"number"`. Optional.

How many failures before wait is triggered. When login failure count is hit, user will be
temporarily locked for a max duration of 15 minutes.

Upstream description:

How many failures before wait is triggered. When login failure count is hit, user will be
temporarily locked for a max duration of 15 minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
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
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-85c7435b89ecd4506d89787fcc2c6f48c0e861942fb331cbb876a177d618c40f"></a>

## Next pages — brute_force_detection / 2de426369f9d / 5

- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-16136964bc77142e101e918b10e24c07bf5c7f99c8c5df90b4efd69072c8b3a2)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)

<a id="canonical-d22ac2c8491ce686c86e96f5f7e8666dcd3f0370f6f460088d0e6594a3c2a10a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7474aede3fa2f1c9b86c08a7fda5daf674c51c1809205919ba3e47962dfb64d"></a>

## password_policy — password_policy / d9dc91ac7c61 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-16136964bc77142e101e918b10e24c07bf5c7f99c8c5df90b4efd69072c8b3a2)
- password_policy

<a id="canonical-ca5900014c605370373120907771c6218532d4cb21e99d13d6ea35da0e60588f"></a>

Type: `"object"`. single nested block, Optional.

Policy configuration for this feature.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("minimum_length")}
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
password_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-90727f18e594d3586f634e90d16f0f8633454df02a862aadffc85af9ed9d0e19"></a>

## Direct properties — password_policy / d9dc91ac7c61 / 3

<a id="canonical-7cba66b03d1fb1e5d6448ca8a3f8ec5d794c36f903ce1f4a253c562dcbc73068"></a>

<a id="canonical-4086aebe70481210c49f2200d2218446237af6c068b42564acdeafe6069ea940"></a>

## digits property — password_policy / d9dc91ac7c61 / 4

Type: `"number"`. Optional.

The number of digits required to be in the password string.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-b5b9acfb68fab7b9b791b59eb5be60af58f82176d4977a0b85d7c9456b05844c"></a>

<a id="canonical-1d3d3b488981edcf60abee52ca758fb0c1af03841f60fb1a70605a1def2b6ee2"></a>

## expire_password property — password_policy / d9dc91ac7c61 / 5

Type: `"number"`. Optional.

The number of days for which the password is valid. After the number of days has expired, the user
is required to change their password.

Upstream description:

The number of days for which the password is valid. After the number of days has expired, the user
is required to change their password.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(1080),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1080,
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
    "ves.io.schema.rules.uint32.lte": "1080"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1080"
  }
}
```

<a id="canonical-fde746eb1abecdc040330a485d532bc0b14649301d5ae95f711cf17b06983bf8"></a>

<a id="canonical-d149094c3d9fc173c921ba7681db3966a506ad5e9f24156f1c1b4d7a977f1515"></a>

## lowercase_characters property — password_policy / d9dc91ac7c61 / 6

Type: `"number"`. Optional.

The number of lower case letters required to be in the password string.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-6f09724d5a49f4f4cee40e6da8ec4efb4833ca37d400a2e1e89694c2c8504a5e"></a>

<a id="canonical-4fc93460e339c5c88a61039faecc335bcd6c2405a930afc81d0321a092361c25"></a>

## minimum_length property — password_policy / d9dc91ac7c61 / 7

Type: `"number"`. Optional.

Minimum Length. Minimum length of password.

Upstream description:

Minimum length of password.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(7),
}
```

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
    "minimum": 7
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "7"
  }
}
```

<a id="canonical-bc02e4de8199cba4df58d6de68aea9a063a7d84c82e4d82a3a9e85165a728044"></a>

<a id="canonical-6ea072eb350e26bf4aae2a11a22bb99cc273db6d26dcbf277aa8fb7911f30add"></a>

## not_recently_used property — password_policy / d9dc91ac7c61 / 8

Type: `"number"`. Optional.

Policy is used to restrict user from using previously used passwords. Number that's set determines
number of last passwords which user cannot use as new password.

Upstream description:

This policy is used to restrict user from using previously used passwords. Number that's set
determines number of last passwords which user cannot use as new password.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-694547d573ccf852309153f0a61f9b792dd637c513ec794a9ab013a7388c371a"></a>

<a id="canonical-bd04bbbb0e06e820707ab9d083872f243e7589873d207bf935a7711280ab9f25"></a>

## not_username property — password_policy / d9dc91ac7c61 / 9

Type: `"bool"`. Optional.

When set, the password is not allowed to be the same as the username.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-35f3c0239e61050a34ae911828cd5e86dcbe624f976da146378e08812c591c23"></a>

<a id="canonical-0699e9a163fbff48c135bf1cfa037800b448d6296123185e0fea31aaa2b9b497"></a>

## special_characters property — password_policy / d9dc91ac7c61 / 10

Type: `"number"`. Optional.

The number of special characters like '?!\#%$' required to be in the password string.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-f50dd904cf9b8fe8b362474d50496f5c703ebe313390007710cba5d146fe1a5c"></a>

<a id="canonical-96c38123692651b058165846c68455efa7ecf983056e88db87276d5feb12c156"></a>

## uppercase_characters property — password_policy / d9dc91ac7c61 / 11

Type: `"number"`. Optional.

The number of upper case letters required to be in the password string.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-4ca9016ac5afc66c665f4b21361fbb4c81bd7315402966a29ce4461429b6c4ae"></a>

## Next pages — password_policy / d9dc91ac7c61 / 12

- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-16136964bc77142e101e918b10e24c07bf5c7f99c8c5df90b4efd69072c8b3a2)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)

<a id="canonical-452a289512a12b40aecd668a1edafd36a78f77e364e271f9ea07d757fd73073f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65dc12b57ec4de1bbfc75c7e12bd5be176a828d6aa548fb5763dd0089a346ad1"></a>

## tenant_details — tenant_details / 8a7110248df3 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-16136964bc77142e101e918b10e24c07bf5c7f99c8c5df90b4efd69072c8b3a2)
- tenant_details

<a id="canonical-ab19e16cb9b735c8a149c0869696a7ed8eeebe25c1784f220e5a12a51eb2f76d"></a>

Type: `"object"`. single nested block, Optional.

BasicConfiguration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("display_name")}
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
tenant_details {
  # Configure direct properties listed below.
}
```

<a id="canonical-f99ad47b2432dd2d808d11f1117e08fa0041f9c9605a86a28afdfb3cdf40e2a4"></a>

## Direct properties — tenant_details / 8a7110248df3 / 3

<a id="canonical-8665e3f55149777bac2ad93219c4cd4bc425d44edd4ad35bde447cae9129129d"></a>

<a id="canonical-c9919a06587154f2c552f3620813cc6c22114016932d5cd42813f5313785fe56"></a>

## display_name property — tenant_details / 8a7110248df3 / 4

Type: `"string"`. Optional.

Changes the tenant name displayed during login without affecting your company’s domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[^<>&\\\"]+$"
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
    "ves.io.schema.rules.string.pattern": "^[^<>&\\\"]+$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^[^<>&\\\"]+$"
  }
}
```

<a id="canonical-2b059bda4c96812706581f933f87968f2b7bef8d66d43ea0592a3dd72d370638"></a>

## Next pages — tenant_details / 8a7110248df3 / 5

- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-16136964bc77142e101e918b10e24c07bf5c7f99c8c5df90b4efd69072c8b3a2)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)

<a id="canonical-2560c76f29b89582c64faf83495deba1b668a3386d8290a61ae8072cb0666184"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9217cbe62f91c65dcf5a914912fb7403d83fc9e617a73d3354093e3272ad192f"></a>

## timeouts — timeouts / 828798e48ac6 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-16136964bc77142e101e918b10e24c07bf5c7f99c8c5df90b4efd69072c8b3a2)
- timeouts

<a id="canonical-70b2fa7667f36ffb7b4b5b17d674654eb37e39146edff7e73430f16608e085b7"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3d39856ba527fffc71d7ab45309da8579089682a68d1db1484c0e42682492136"></a>

## Direct properties — timeouts / 828798e48ac6 / 3

<a id="canonical-5f68907166c6207ee3fdcb2a4df71b43ec645d53b103172a0cc0c947964537b5"></a>

<a id="canonical-6a66d0ecedbccfc6cdfd4b9153a4f110f88ce456e34990e6f8751b2896dc6333"></a>

## create property — timeouts / 828798e48ac6 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-576465c0618990ea0946721e8de7306b3710ffef5e6d65dbbf72f8e899401f49"></a>

<a id="canonical-cbe3526502f97f9df232861eeadfda9e86c7967a6d42ac73cc0fe56269d80f86"></a>

## delete property — timeouts / 828798e48ac6 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0c5735b6b4ab35b800a29d5fc09d5f6ae02ce5e7371cf7a914b2acc1ddf01e74"></a>

<a id="canonical-94feb0d738ac9706909e4f0935493adf5bfbf35f147244e620a9716c8c3e9e9a"></a>

## read property — timeouts / 828798e48ac6 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-47930ed94ac1d3b1a8ab54f11a44032c6eb900bfceab575eef3336e2b773059d"></a>

<a id="canonical-5c7e76d015c2513da08f230ea2a460e7d4af8cf0e61af255c3084cdd7d6fb83e"></a>

## update property — timeouts / 828798e48ac6 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-99510d73a528d8af0091555c56f37f375a494afd088b7f7057f38719c3693826"></a>

## Next pages — timeouts / 828798e48ac6 / 8

- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-16136964bc77142e101e918b10e24c07bf5c7f99c8c5df90b4efd69072c8b3a2)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)

<a id="canonical-c2388a7d0b4ae235a90c04846dd8c1fc6d3d1e30b66d38bc7a5af4139f3518ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75bc2ba2d69dcb82cc67d30517d4397abb97e2ddcdd1a39a7c7a6c5982530b94"></a>

## user_session_expiration — user_session_expiration / 9f165bd0dc6b / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-16136964bc77142e101e918b10e24c07bf5c7f99c8c5df90b4efd69072c8b3a2)
- user_session_expiration

<a id="canonical-91ac0514301375b0a02750bd5915a0f3db4dfc75de82dfe993541045323eb472"></a>

Type: `"object"`. single nested block, Optional.

Defines all session-related expiration for user sessions within a tenant's environment. Relationship
between session\_expiry and cookie\_expiry: - session\_expiry defines the 'absolute maximum
duration' of a session and enforces RE-authentication after this time. - cookie\_expiry defines
the..

Upstream description:

Defines all session-related expiration for user sessions within a tenant's environment. Relationship
between session\_expiry and cookie\_expiry: &#8203;- session\_expiry defines the 'absolute maximum
duration' of a session and enforces RE-authentication after this time. &#8203;- cookie\_expiry
defines the 'inactivity timeout', which resets on user activity and only logs out users after idle
periods. Together, these ensure the user is logged out when either the session reaches its maximum
age or the user is inactive for too long.

Receipt-pinned upstream constraints:

```json
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
user_session_expiration {
  # Configure direct properties listed below.
}
```

<a id="canonical-2c2e6ff37fcd177e515227db10505b7a7fe742365fd56670ecc2eaad971c7ac1"></a>

## Direct properties — user_session_expiration / 9f165bd0dc6b / 3

- [absolute_timeout](resources--tenant_configuration--reference--group-001.md#canonical-1d5b2abc3e80c6e61035f8fb3ff373d5a5fdb93b2d6ef40ae236508ab45dad25): complete subsection reference.

- [idle_timeout](resources--tenant_configuration--reference--group-001.md#canonical-8794f1a173eba1a10c3602dcf1711b64171fd5d034b9cb286e00d6ba7189c6b6): complete subsection reference.

<a id="canonical-e98778698b9939f28acaa15a2982a05cb0b0b5c926d687373668f2be943ecdd1"></a>

## Next pages — user_session_expiration / 9f165bd0dc6b / 4

- [user_session_expiration.absolute_timeout](resources--tenant_configuration--reference--group-001.md#canonical-1d5b2abc3e80c6e61035f8fb3ff373d5a5fdb93b2d6ef40ae236508ab45dad25)
- [user_session_expiration.idle_timeout](resources--tenant_configuration--reference--group-001.md#canonical-8794f1a173eba1a10c3602dcf1711b64171fd5d034b9cb286e00d6ba7189c6b6)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-16136964bc77142e101e918b10e24c07bf5c7f99c8c5df90b4efd69072c8b3a2)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)

<a id="canonical-1d5b2abc3e80c6e61035f8fb3ff373d5a5fdb93b2d6ef40ae236508ab45dad25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36269dd723c789a53140ac9ede8b1b6a58834888fa48d5027e5d414445416500"></a>

## user_session_expiration.absolute_timeout — user_session_expiration.absolute_timeout / a9f8f8a8bb02 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-16136964bc77142e101e918b10e24c07bf5c7f99c8c5df90b4efd69072c8b3a2)
- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-c2388a7d0b4ae235a90c04846dd8c1fc6d3d1e30b66d38bc7a5af4139f3518ba)
- user_session_expiration.absolute_timeout

<a id="canonical-d10df19bb39dbbe0395e55c1cd57b1008f975f3b22b8a6115cc638b79da1ea7b"></a>

Type: `"object"`. single nested block, Optional.

Represents the session expiration duration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("hours",
    "minutes")}
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
  "x-ves-oneof-field-unit_of_time": "[\"hours\",\"minutes\"]"
}
```

Terraform syntax:

```terraform
absolute_timeout {
  # Configure direct properties listed below.
}
```

<a id="canonical-0bb741273f0e60f9a5c2cc2dde4607cfb9d9735abb7712694f1fd7bf6cd505d5"></a>

## Direct properties — user_session_expiration.absolute_timeout / a9f8f8a8bb02 / 3

- [hours](resources--tenant_configuration--reference--group-001.md#canonical-3cc7883db7179c63f768244172ae8ed8a72157ad7634f380bf2c0ffa494a10c3): complete subsection reference.

- [minutes](resources--tenant_configuration--reference--group-001.md#canonical-995ab8627c0cb630bbc25af1c18f50464380804b22df48d6387a15e80697d64f): complete subsection reference.

<a id="canonical-6c42050416215d33fb262aa98edfd72137412fad652cd4fe7db8581142d375e2"></a>

## Next pages — user_session_expiration.absolute_timeout / a9f8f8a8bb02 / 4

- [user_session_expiration.absolute_timeout.hours](resources--tenant_configuration--reference--group-001.md#canonical-3cc7883db7179c63f768244172ae8ed8a72157ad7634f380bf2c0ffa494a10c3)
- [user_session_expiration.absolute_timeout.minutes](resources--tenant_configuration--reference--group-001.md#canonical-995ab8627c0cb630bbc25af1c18f50464380804b22df48d6387a15e80697d64f)
- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-c2388a7d0b4ae235a90c04846dd8c1fc6d3d1e30b66d38bc7a5af4139f3518ba)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)

<a id="canonical-3cc7883db7179c63f768244172ae8ed8a72157ad7634f380bf2c0ffa494a10c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-104180d3c308cb2d6e1ecf5ee86c7e3bede0cdaf44b84840f8a42e4a17174066"></a>

## user_session_expiration.absolute_timeout.hours — user_session_expiration.absolute_timeout.hours / 1cc6d5d4ac8c / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-16136964bc77142e101e918b10e24c07bf5c7f99c8c5df90b4efd69072c8b3a2)
- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-c2388a7d0b4ae235a90c04846dd8c1fc6d3d1e30b66d38bc7a5af4139f3518ba)
- [user_session_expiration.absolute_timeout](resources--tenant_configuration--reference--group-001.md#canonical-1d5b2abc3e80c6e61035f8fb3ff373d5a5fdb93b2d6ef40ae236508ab45dad25)
- user_session_expiration.absolute_timeout.hours

<a id="canonical-bf079e049dc4612363887688f8fe41c8019603d2a00f9b6812bbb2db0557de15"></a>

Type: `"object"`. single nested block, Optional.

Represents the session duration in hours.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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
hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-cdd0f621cbe7abd3860f05ee8c652c3e9bdb334f7eb8722030b9f4bee5a2accc"></a>

## Direct properties — user_session_expiration.absolute_timeout.hours / 1cc6d5d4ac8c / 3

<a id="canonical-d6332dcbcfe284239d6d706fd04783adc542f5c53c3aeec261fae9d76ad69697"></a>

<a id="canonical-366c3145a8097f98ac9c2c94396aa2be663a7d27d80386663bc22b2463662fc6"></a>

## duration property — user_session_expiration.absolute_timeout.hours / 1cc6d5d4ac8c / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 720),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 720,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "720"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "720"
  }
}
```

<a id="canonical-2afd125070b63aa0b72402c0f6e010ca442a50ca9ea4be2897e5e11cea951322"></a>

## Next pages — user_session_expiration.absolute_timeout.hours / 1cc6d5d4ac8c / 5

- [user_session_expiration.absolute_timeout](resources--tenant_configuration--reference--group-001.md#canonical-1d5b2abc3e80c6e61035f8fb3ff373d5a5fdb93b2d6ef40ae236508ab45dad25)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)

<a id="canonical-995ab8627c0cb630bbc25af1c18f50464380804b22df48d6387a15e80697d64f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-282b454b3008e04f48efb64ac97a65c1eef85ceda31746aa64dfd8d81e0c3a45"></a>

## user_session_expiration.absolute_timeout.minutes — user_session_expiration.absolute_timeout.minutes / 760b711693a3 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-16136964bc77142e101e918b10e24c07bf5c7f99c8c5df90b4efd69072c8b3a2)
- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-c2388a7d0b4ae235a90c04846dd8c1fc6d3d1e30b66d38bc7a5af4139f3518ba)
- [user_session_expiration.absolute_timeout](resources--tenant_configuration--reference--group-001.md#canonical-1d5b2abc3e80c6e61035f8fb3ff373d5a5fdb93b2d6ef40ae236508ab45dad25)
- user_session_expiration.absolute_timeout.minutes

<a id="canonical-ff9011ad00654d15cba61cbe42f2ef89022699922ca3c75560c110156f3fafc3"></a>

Type: `"object"`. single nested block, Optional.

Represents the session duration in minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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
minutes {
  # Configure direct properties listed below.
}
```

<a id="canonical-054795f58397283900020b7dbef05373edb87c4e2c7dc996be92a9af1add81c0"></a>

## Direct properties — user_session_expiration.absolute_timeout.minutes / 760b711693a3 / 3

<a id="canonical-f76dfe12793470231a9e16928511f900f5ce3e66313eb5d8c09f6fbc4f8e05f0"></a>

<a id="canonical-e419c845923a1e39935a0c3d2336be4959ebc5355d97457a465117f15cc7fb1d"></a>

## duration property — user_session_expiration.absolute_timeout.minutes / 760b711693a3 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(5, 43200),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 43200,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 5
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "5",
    "ves.io.schema.rules.uint32.lte": "43200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "5",
    "ves.io.schema.rules.uint32.lte": "43200"
  }
}
```

<a id="canonical-403fa6f18731955cd7bf138cb46987a40c4fc2f179bb8f30e05594fbc943031a"></a>

## Next pages — user_session_expiration.absolute_timeout.minutes / 760b711693a3 / 5

- [user_session_expiration.absolute_timeout](resources--tenant_configuration--reference--group-001.md#canonical-1d5b2abc3e80c6e61035f8fb3ff373d5a5fdb93b2d6ef40ae236508ab45dad25)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)

<a id="canonical-8794f1a173eba1a10c3602dcf1711b64171fd5d034b9cb286e00d6ba7189c6b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98879631fc7b0e8b29a20f199addcb190d47301e38cb34ed6b360399126403dd"></a>

## user_session_expiration.idle_timeout — user_session_expiration.idle_timeout / 7c390341fc6e / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-16136964bc77142e101e918b10e24c07bf5c7f99c8c5df90b4efd69072c8b3a2)
- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-c2388a7d0b4ae235a90c04846dd8c1fc6d3d1e30b66d38bc7a5af4139f3518ba)
- user_session_expiration.idle_timeout

<a id="canonical-5c12ce410905f00420ddfd71d16183354e85f4e0c327a80c9d6033c81e8abfa6"></a>

Type: `"object"`. single nested block, Optional.

Represents the cookie expiration duration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("hours",
    "minutes")}
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
  "x-ves-oneof-field-unit_of_time": "[\"hours\",\"minutes\"]"
}
```

Terraform syntax:

```terraform
idle_timeout {
  # Configure direct properties listed below.
}
```

<a id="canonical-a7a16b42fa9c051ddb8b315909508709fa1047a03c330b6e117346cc1258a4b6"></a>

## Direct properties — user_session_expiration.idle_timeout / 7c390341fc6e / 3

- [hours](resources--tenant_configuration--reference--group-001.md#canonical-06083d89bba8fbe6c374b2086817f06ed814175ddcf5013ad4488a1d2c109844): complete subsection reference.

- [minutes](resources--tenant_configuration--reference--group-001.md#canonical-c8e0a55bdb8dbe8e00eef2be0a0431926723c184c5a2e1f003882040dc316641): complete subsection reference.

<a id="canonical-d9904361e16a213cede78eb877f49da3b3ad1474f4710a88429051d0fc1acde1"></a>

## Next pages — user_session_expiration.idle_timeout / 7c390341fc6e / 4

- [user_session_expiration.idle_timeout.hours](resources--tenant_configuration--reference--group-001.md#canonical-06083d89bba8fbe6c374b2086817f06ed814175ddcf5013ad4488a1d2c109844)
- [user_session_expiration.idle_timeout.minutes](resources--tenant_configuration--reference--group-001.md#canonical-c8e0a55bdb8dbe8e00eef2be0a0431926723c184c5a2e1f003882040dc316641)
- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-c2388a7d0b4ae235a90c04846dd8c1fc6d3d1e30b66d38bc7a5af4139f3518ba)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)

<a id="canonical-06083d89bba8fbe6c374b2086817f06ed814175ddcf5013ad4488a1d2c109844"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d21e190ea58b743675730c8c9aa6cdabb615793dd33735bfdce04e09372ded1"></a>

## user_session_expiration.idle_timeout.hours — user_session_expiration.idle_timeout.hours / afe7002489fb / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-16136964bc77142e101e918b10e24c07bf5c7f99c8c5df90b4efd69072c8b3a2)
- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-c2388a7d0b4ae235a90c04846dd8c1fc6d3d1e30b66d38bc7a5af4139f3518ba)
- [user_session_expiration.idle_timeout](resources--tenant_configuration--reference--group-001.md#canonical-8794f1a173eba1a10c3602dcf1711b64171fd5d034b9cb286e00d6ba7189c6b6)
- user_session_expiration.idle_timeout.hours

<a id="canonical-693340cd5a297b2e7411a9285bd90b1f5831fba903a1e47f1a77c5fa83ac1f3a"></a>

Type: `"object"`. single nested block, Optional.

Represents the cookie duration in hours.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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
hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-b138172e7d22e859abb83b86a4237559eae79a6c7670c81160b730b682d6c109"></a>

## Direct properties — user_session_expiration.idle_timeout.hours / afe7002489fb / 3

<a id="canonical-cb6af5cc68fdaf381c8b51e7304c7a7f27763911cc43de5d0d9e1e3f21d93fea"></a>

<a id="canonical-fb3565ee56e415ff1ffa356b6b92c9de66e2b294ed925b1d82fac758076e6b43"></a>

## duration property — user_session_expiration.idle_timeout.hours / afe7002489fb / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 720),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 720,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "720"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "720"
  }
}
```

<a id="canonical-23b2b159232f5e89d64368fb8569534ccf325622ba4d63151e6b3a2999053c30"></a>

## Next pages — user_session_expiration.idle_timeout.hours / afe7002489fb / 5

- [user_session_expiration.idle_timeout](resources--tenant_configuration--reference--group-001.md#canonical-8794f1a173eba1a10c3602dcf1711b64171fd5d034b9cb286e00d6ba7189c6b6)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)

<a id="canonical-c8e0a55bdb8dbe8e00eef2be0a0431926723c184c5a2e1f003882040dc316641"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-390e14bedb7b9634be461d74c0e923bec3fd7739060674c9ad9492396fe10a55"></a>

## user_session_expiration.idle_timeout.minutes — user_session_expiration.idle_timeout.minutes / 14181b3ee5bc / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-16136964bc77142e101e918b10e24c07bf5c7f99c8c5df90b4efd69072c8b3a2)
- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-c2388a7d0b4ae235a90c04846dd8c1fc6d3d1e30b66d38bc7a5af4139f3518ba)
- [user_session_expiration.idle_timeout](resources--tenant_configuration--reference--group-001.md#canonical-8794f1a173eba1a10c3602dcf1711b64171fd5d034b9cb286e00d6ba7189c6b6)
- user_session_expiration.idle_timeout.minutes

<a id="canonical-473a4a67168c1711cf9882d7893c5050bc050eac27a6839d4a6fead1422e896a"></a>

Type: `"object"`. single nested block, Optional.

Represents the cookie duration in minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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
minutes {
  # Configure direct properties listed below.
}
```

<a id="canonical-4e78911affa5088d446233f02dab064776876e9d89f16d0261af6d554cb4bf77"></a>

## Direct properties — user_session_expiration.idle_timeout.minutes / 14181b3ee5bc / 3

<a id="canonical-892ec7d143dad9e846826752fe9064335a4f4ee6a11974d151c1d04475bd76ee"></a>

<a id="canonical-365ee33204f02304e9cb3ef6342a16d7cecfea134f0bad8416c333dce5dda94b"></a>

## duration property — user_session_expiration.idle_timeout.minutes / 14181b3ee5bc / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(5, 43200),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 43200,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 5
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "5",
    "ves.io.schema.rules.uint32.lte": "43200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "5",
    "ves.io.schema.rules.uint32.lte": "43200"
  }
}
```

<a id="canonical-c023dd4d6c95fbc08e8510d62826648dc129d1d75b689187f797075246a1e47f"></a>

## Next pages — user_session_expiration.idle_timeout.minutes / 14181b3ee5bc / 5

- [user_session_expiration.idle_timeout](resources--tenant_configuration--reference--group-001.md#canonical-8794f1a173eba1a10c3602dcf1711b64171fd5d034b9cb286e00d6ba7189c6b6)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)
