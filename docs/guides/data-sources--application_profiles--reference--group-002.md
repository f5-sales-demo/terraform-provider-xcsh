---
page_title: "xcsh_application_profiles reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles reference."
---

# xcsh_application_profiles reference

<a id="canonical-07e21d78c6f6fb3d4e3e8997b087d93696f5e930e574757c2467e7cb04d74083"></a>

## virtual_server.auto_last_hop.auto_last_hop_enable — virtual_server.auto_last_hop.auto_last_hop_enable / 624e9cdfbbb2 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.auto_last_hop](data-sources--application_profiles--reference--group-001.md#canonical-67b69d939156306d219f0b653719624f46811a1991a32753e6fdfd42e0c3de4f)
- virtual_server.auto_last_hop.auto_last_hop_enable

<a id="canonical-82206f3b2f1975ec40ed44e55f1960fe96510f1ec9d22b955464cbd70026bafb"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for auto last hop enable.

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

<a id="canonical-0317505de158ef812660157138c71e0d8f1803c5b0d168127b78aef8ae356764"></a>

## Direct properties — virtual_server.auto_last_hop.auto_last_hop_enable / 624e9cdfbbb2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5c793c84ba48fd4c6d94b58189662cd3c0141807970f71fddb5000695c62078b"></a>

## Next pages — virtual_server.auto_last_hop.auto_last_hop_enable / 624e9cdfbbb2 / 4

- [virtual_server.auto_last_hop](data-sources--application_profiles--reference--group-001.md#canonical-67b69d939156306d219f0b653719624f46811a1991a32753e6fdfd42e0c3de4f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-5d73316e5d127760389f4de33f8f2c93af8832f8e4b48154b679834d4f5408a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e35f9b9ff2cbd7b05e2569a3f2fada20eb98ed1f44e5f111449b5dbee7b848b"></a>

## virtual_server.clone_pool_client — virtual_server.clone_pool_client / a4c5c1d2dc54 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.clone_pool_client

<a id="canonical-5f21ebe5b4511a050617e0b9479dedcf91e28d53d51ba4590cb00fdeee099a77"></a>

Type: `"list"`. Computed.

Replicates client-side traffic (that is, prior to address translation) to a member of the specified
pool.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-66c9ca350a5d91488bb259ac811cc8c182422c878dfd1b9ec2e8a3b12d3f60e0"></a>

## Direct properties — virtual_server.clone_pool_client / a4c5c1d2dc54 / 3

<a id="canonical-e73990620c630b475a7a61bd2756a3bc2d46a5a009c23dfd8944a3a0c3bd5fc3"></a>

<a id="canonical-a56c6687153f87148e6438a5d01cc5588089f3647424765dcaf7aff13ef94607"></a>

## kind property — virtual_server.clone_pool_client / a4c5c1d2dc54 / 4

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

<a id="canonical-79cf1ed6eef4c2cbcaff54c4e7ac7dfa31d1669f27278fa779163591c81ec621"></a>

<a id="canonical-d52022a02efb83ce4ca24efd39898a1dd1c12842ac0f346549f6097c26828a67"></a>

## name property — virtual_server.clone_pool_client / a4c5c1d2dc54 / 5

Type: `"string"`. Computed.

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

<a id="canonical-a5b6aad587df6b3fac47bb6af430e35a75fe798bc1a35251a10c4f2cfd2f7e6e"></a>

<a id="canonical-b235b6add44b9be26ab614d315a445a83c945e0ce23e585076659f07d3592500"></a>

## namespace property — virtual_server.clone_pool_client / a4c5c1d2dc54 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-84fa3e14e52dbe6ba8b066f9e832492a582421fbc7c017076a2a7a89521f3e7e"></a>

<a id="canonical-7d1b3d0e8400c92b365dd04cf382998196760512f8acde34861f9870a5777e36"></a>

## tenant property — virtual_server.clone_pool_client / a4c5c1d2dc54 / 7

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

<a id="canonical-6f564e025e781779fd26e0b7e730307710beb774113d27f266dd6c04f9eb5d90"></a>

<a id="canonical-a1653a0cc28c93bc086ae6922ceb7ad6b425175f738d8b7686d35e58a1bb524b"></a>

## uid property — virtual_server.clone_pool_client / a4c5c1d2dc54 / 8

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

<a id="canonical-fc1f92cb7177997834b7382b295d6acbedea939e2e7732e8ef11ad30fba4a8ae"></a>

## Next pages — virtual_server.clone_pool_client / a4c5c1d2dc54 / 9

- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-8ec873d6c734de5225aca2d4908de716882031bc8c0863c1dd26084af348358f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ce72f75eee0d223ff0fef8701742dd5b19a5481ffb987711bcd28040a31f38b"></a>

## virtual_server.clone_pool_server — virtual_server.clone_pool_server / 59b54b377f5c / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.clone_pool_server

<a id="canonical-29dd50e32aa5ed95979b7dd9e3ac8a4618b69ed68e6d42aabb77c57ed6acb868"></a>

Type: `"list"`. Computed.

Replicates server-side traffic (that is, prior to address translation) to a member of the specified
pool.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-a871cfd9033a3e466f804e4338101dac83401e51cecdd980baa34ec83055af20"></a>

## Direct properties — virtual_server.clone_pool_server / 59b54b377f5c / 3

<a id="canonical-6f59548b35299463534c643bc214049db390c75f2e5c76542479fea6ea0e05bd"></a>

<a id="canonical-20983aba2d2097375db7082ec3a54b2b97009bfaa3fb450520efde0c334b3dc3"></a>

## kind property — virtual_server.clone_pool_server / 59b54b377f5c / 4

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

<a id="canonical-b335996841b27c0c6c0fed4e8843eb67820ad20564c8a098dd7783410d063a2d"></a>

<a id="canonical-ba5b01c435fd29c5d201702f4f9408ff3ec50a5072329503303feab6c15f4208"></a>

## name property — virtual_server.clone_pool_server / 59b54b377f5c / 5

Type: `"string"`. Computed.

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

<a id="canonical-91e361ed2bdcf710b383c7b64e145e31e1fed4713369bd7904216b3ef83d7311"></a>

<a id="canonical-f10c7dab1ca23e0edb304643983c6adefbffe3957f44a44b8951615fa1c02cb6"></a>

## namespace property — virtual_server.clone_pool_server / 59b54b377f5c / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-0efca091478e84c6812815566b8b779a2cb3055caab9d4dcc2d068671713a711"></a>

<a id="canonical-08ce46391711092370cb8b9aa8ba3e919351ee53fde9bcc5af56608a328fb48d"></a>

## tenant property — virtual_server.clone_pool_server / 59b54b377f5c / 7

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

<a id="canonical-67be68133578d45274404e49ac1e87ac629dfdc3ea234e36ffb60878e38b9ce6"></a>

<a id="canonical-6872b70d5044d3af33185e1e13fedcc0ecd39bfdc16917daa2809fa4682da891"></a>

## uid property — virtual_server.clone_pool_server / 59b54b377f5c / 8

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

<a id="canonical-cae614867bba81a4644b54d3c7a8dc2e2d9789d0ebf3594c95bda155fea65251"></a>

## Next pages — virtual_server.clone_pool_server / 59b54b377f5c / 9

- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-0ff1a1f4c36341bce502bae92b575dac0a88378f391590e464696488a7ce3c5f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-559949e042fd5d72ea76859eb7d3ade2eb4f178bcde934e357e2e6ed01410c9e"></a>

## virtual_server.connection_rate_limit_mode — virtual_server.connection_rate_limit_mode / 187a0eded33f / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.connection_rate_limit_mode

<a id="canonical-ddc4ed26919c16086ceb2bcaefafcb34d867c9cc514e0c9716fc6cec179e319c"></a>

Type: `"single"`. Computed.

Configuration parameter for connection rate limit mode.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-connection_rate_limit_mode_choice": "[\"per_destination_address\",\"per_source_address\",\"per_source_destination_address\",\"per_virtual_server\",\"per_virtual_server_destination_address\",\"per_virtual_server_source_address\",\"per_virtual_server_source_destination_address\"]"
}
```

<a id="canonical-b4a271e1c8d891687564f6723a706454341a5fa077919c1e619cd7e219439c5f"></a>

## Direct properties — virtual_server.connection_rate_limit_mode / 187a0eded33f / 3

- [per_destination_address](data-sources--application_profiles--reference--group-002.md#canonical-421edf2b24d45a8820654fb5fa14a06f4754301c0ee01196a19e6d34d1abd2ae): complete subsection reference.

- [per_source_address](data-sources--application_profiles--reference--group-002.md#canonical-57ac8cd2c77c00a8abf4140d094857b8d9e77c0d2a06ccf83b263405ea0a06eb): complete subsection reference.

- [per_source_destination_address](data-sources--application_profiles--reference--group-002.md#canonical-fc3bafea12c7190d7394caa345545a61217510973a69873dfaca00404d887f82): complete subsection reference.

- [per_virtual_server](data-sources--application_profiles--reference--group-002.md#canonical-65c3820ebea2fdee7bd2e56d56c751210dddc6f4b753df82d0bfa0359803bb49): complete subsection reference.

- [per_virtual_server_destination_address](data-sources--application_profiles--reference--group-002.md#canonical-0a2e3aabafd37947ed516d4832a013879b23f25284fd0689f078a12a81912122): complete subsection reference.

- [per_virtual_server_source_address](data-sources--application_profiles--reference--group-002.md#canonical-6302eab90ac373106f695ae44706ac32cc66d3a645ae576c3a97593da6aeefec): complete subsection reference.

- [per_virtual_server_source_destination_address](data-sources--application_profiles--reference--group-002.md#canonical-4a64939cf3af5839cb5a2004b6986bd2775a5c4ad0378e4554a709ac0afc3a1d): complete subsection reference.

<a id="canonical-11848f706e606b47dd7a3aadbf56ce05a033caf42aea2a14e9edcbebde902500"></a>

## Next pages — virtual_server.connection_rate_limit_mode / 187a0eded33f / 4

- [virtual_server.connection_rate_limit_mode.per_destination_address](data-sources--application_profiles--reference--group-002.md#canonical-421edf2b24d45a8820654fb5fa14a06f4754301c0ee01196a19e6d34d1abd2ae)
- [virtual_server.connection_rate_limit_mode.per_source_address](data-sources--application_profiles--reference--group-002.md#canonical-57ac8cd2c77c00a8abf4140d094857b8d9e77c0d2a06ccf83b263405ea0a06eb)
- [virtual_server.connection_rate_limit_mode.per_source_destination_address](data-sources--application_profiles--reference--group-002.md#canonical-fc3bafea12c7190d7394caa345545a61217510973a69873dfaca00404d887f82)
- [virtual_server.connection_rate_limit_mode.per_virtual_server](data-sources--application_profiles--reference--group-002.md#canonical-65c3820ebea2fdee7bd2e56d56c751210dddc6f4b753df82d0bfa0359803bb49)
- [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address](data-sources--application_profiles--reference--group-002.md#canonical-0a2e3aabafd37947ed516d4832a013879b23f25284fd0689f078a12a81912122)
- [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address](data-sources--application_profiles--reference--group-002.md#canonical-6302eab90ac373106f695ae44706ac32cc66d3a645ae576c3a97593da6aeefec)
- [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address](data-sources--application_profiles--reference--group-002.md#canonical-4a64939cf3af5839cb5a2004b6986bd2775a5c4ad0378e4554a709ac0afc3a1d)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-421edf2b24d45a8820654fb5fa14a06f4754301c0ee01196a19e6d34d1abd2ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f964a65bc3c5259268de9b9502f25bc514890eb4e38cb857e763e2c102ca227"></a>

## virtual_server.connection_rate_limit_mode.per_destination_address — virtual_server.connection_rate_limit_mode.per_destination_address / b9aef08c21c2 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-002.md#canonical-0ff1a1f4c36341bce502bae92b575dac0a88378f391590e464696488a7ce3c5f)
- virtual_server.connection_rate_limit_mode.per_destination_address

<a id="canonical-ee9b3d2df982dc70d72c52f5f7503e8fcfee153548801069800663586acbc4c2"></a>

Type: `"single"`. Computed.

Destination Address Mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-e8b95d352ab9a570ad2629a55369e792903e8033e25f8a0fbb87c6f9300c898a"></a>

## Direct properties — virtual_server.connection_rate_limit_mode.per_destination_address / b9aef08c21c2 / 3

<a id="canonical-4119fea0293eb65fbd643d95bd1167fea7b5b3fcb3ebfc2ae29c5c9bbee723eb"></a>

<a id="canonical-18e62efd9df958b91b508a9e44d9e46c6e0dff5e77eb573227df4904036c0e97"></a>

## destination_mask property — virtual_server.connection_rate_limit_mode.per_destination_address / b9aef08c21c2 / 4

Type: `"number"`. Computed.

Configuration parameter for destination mask.

Upstream description:

Configuration parameter for destination mask

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
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-ee7c2aef0ac01778639b6e465ba65f6c5a5528ebc2b9153fa46be9d5f47394c5"></a>

## Next pages — virtual_server.connection_rate_limit_mode.per_destination_address / b9aef08c21c2 / 5

- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-002.md#canonical-0ff1a1f4c36341bce502bae92b575dac0a88378f391590e464696488a7ce3c5f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-57ac8cd2c77c00a8abf4140d094857b8d9e77c0d2a06ccf83b263405ea0a06eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16bd1b00d6edb54bf9bc5f4c78ab7bdc70540f93ba90b1a7200f81651d8581c0"></a>

## virtual_server.connection_rate_limit_mode.per_source_address — virtual_server.connection_rate_limit_mode.per_source_address / db7ac619cee2 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-002.md#canonical-0ff1a1f4c36341bce502bae92b575dac0a88378f391590e464696488a7ce3c5f)
- virtual_server.connection_rate_limit_mode.per_source_address

<a id="canonical-2aa3ed29cd305a7b4514b960e06e7ad9ee11027971ccaf319f4a3e5747c81b58"></a>

Type: `"single"`. Computed.

Source Address Mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-40b4dc952e85a11ca62e1067f807b6ab260414e94466efdd0b6433e6992029ce"></a>

## Direct properties — virtual_server.connection_rate_limit_mode.per_source_address / db7ac619cee2 / 3

<a id="canonical-edd5dd522d8a5793b3182557bc2acc450c04dc2d74a33901cac8299cd9d3c2e0"></a>

<a id="canonical-cfec3ce72e1365651d974d5fc54ed360b949a2da56e1e5c73ac6cbdf997af854"></a>

## source_mask property — virtual_server.connection_rate_limit_mode.per_source_address / db7ac619cee2 / 4

Type: `"number"`. Computed.

Configuration parameter for source mask.

Upstream description:

Configuration parameter for source mask

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
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-dbddecf1ea8ac9e383154cd0b478ff91f33c703d6b7fb80835215025363b5e22"></a>

## Next pages — virtual_server.connection_rate_limit_mode.per_source_address / db7ac619cee2 / 5

- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-002.md#canonical-0ff1a1f4c36341bce502bae92b575dac0a88378f391590e464696488a7ce3c5f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-fc3bafea12c7190d7394caa345545a61217510973a69873dfaca00404d887f82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-022552c71585d3163afe49b2324594ad7d0b3b420007381a63ec981d7fbec3c2"></a>

## virtual_server.connection_rate_limit_mode.per_source_destination_address — virtual_server.connection_rate_limit_mode.per_source_destination_address / 00007417d8c4 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-002.md#canonical-0ff1a1f4c36341bce502bae92b575dac0a88378f391590e464696488a7ce3c5f)
- virtual_server.connection_rate_limit_mode.per_source_destination_address

<a id="canonical-4d87d6d980d30242f54c620326ff6a882c728222afef67d62f9ab0ef52513eeb"></a>

Type: `"single"`. Computed.

Destination and Source Address Mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-63d247b972912ce23fdbb072ff96b64db7c70de0906991ddf231efa3cd34c3aa"></a>

## Direct properties — virtual_server.connection_rate_limit_mode.per_source_destination_address / 00007417d8c4 / 3

<a id="canonical-86d1bb5653dd131f715aa0998ff0b1544af48f06ed07075fce7fc58e6608696c"></a>

<a id="canonical-5564a3a158b3b24f4dadc63bedb20aa847765a36c08245987bd2476597bb4661"></a>

## destination_mask property — virtual_server.connection_rate_limit_mode.per_source_destination_address / 00007417d8c4 / 4

Type: `"number"`. Computed.

Configuration parameter for destination mask.

Upstream description:

Configuration parameter for destination mask

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
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-8e105434fcd467f29a3406f839675b58d960151b134be74c580dde4ef0a10ac8"></a>

<a id="canonical-cec751336f197f57a6223ab65a70ca7d1b2e86329d1603265e4e6c861e0aec9b"></a>

## source_mask property — virtual_server.connection_rate_limit_mode.per_source_destination_address / 00007417d8c4 / 5

Type: `"number"`. Computed.

Configuration parameter for source mask.

Upstream description:

Configuration parameter for source mask

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
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-9ac4a063d9bcef6f2f88041ee37a30c9aaa9b8665a13b4b71c94cea20445bb1f"></a>

## Next pages — virtual_server.connection_rate_limit_mode.per_source_destination_address / 00007417d8c4 / 6

- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-002.md#canonical-0ff1a1f4c36341bce502bae92b575dac0a88378f391590e464696488a7ce3c5f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-65c3820ebea2fdee7bd2e56d56c751210dddc6f4b753df82d0bfa0359803bb49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2e58c35a7af664aaa65c3ad1b6e7447eeee4812009b2a08dcaeeba794a01397"></a>

## virtual_server.connection_rate_limit_mode.per_virtual_server — virtual_server.connection_rate_limit_mode.per_virtual_server / 86d8821c706f / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-002.md#canonical-0ff1a1f4c36341bce502bae92b575dac0a88378f391590e464696488a7ce3c5f)
- virtual_server.connection_rate_limit_mode.per_virtual_server

<a id="canonical-7df149253f9ad02c14e4bf0e12b3efb6e1bb4a42a4ba591896be99d318b25272"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for per virtual server.

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

<a id="canonical-102c425e1a807796e522f12e3b518cba4140832ebe98a9d3fb0deae1d5696167"></a>

## Direct properties — virtual_server.connection_rate_limit_mode.per_virtual_server / 86d8821c706f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-82de78aaf92f743f0536b8227e1282bf53b0fe8737eef612ffb1da8075f52244"></a>

## Next pages — virtual_server.connection_rate_limit_mode.per_virtual_server / 86d8821c706f / 4

- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-002.md#canonical-0ff1a1f4c36341bce502bae92b575dac0a88378f391590e464696488a7ce3c5f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-0a2e3aabafd37947ed516d4832a013879b23f25284fd0689f078a12a81912122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eba4efa050de75897a94169b3e4fd9172d8ae8b48bbcdedc1715b27a6089d769"></a>

## virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address — virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address / fffbc0deffa5 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-002.md#canonical-0ff1a1f4c36341bce502bae92b575dac0a88378f391590e464696488a7ce3c5f)
- virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address

<a id="canonical-e83333c5c6c718611130819d4816b640e5fb4ee141c0e12828bd92ed6a744cbe"></a>

Type: `"single"`. Computed.

Destination Address Mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-bfe120c264c8789daec70adfe9f626af2e4eecb77d7fedafe900ae641bf22f6e"></a>

## Direct properties — virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address / fffbc0deffa5 / 3

<a id="canonical-84072962501041f56f4ef0f34f616b99537a30f6f6449164f1c6dfb798f6756e"></a>

<a id="canonical-c36f84bcce95f8de976b5a22974b2bde63857707dff052f684909127592d466d"></a>

## destination_mask property — virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address / fffbc0deffa5 / 4

Type: `"number"`. Computed.

Configuration parameter for destination mask.

Upstream description:

Configuration parameter for destination mask

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
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-a9a65d2ecd9e1c779f31a5e7df7229fa3350b48411b64aebd93a4093b28186a4"></a>

## Next pages — virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address / fffbc0deffa5 / 5

- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-002.md#canonical-0ff1a1f4c36341bce502bae92b575dac0a88378f391590e464696488a7ce3c5f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-6302eab90ac373106f695ae44706ac32cc66d3a645ae576c3a97593da6aeefec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6bc9f8a17f298cd5b03b6195afcb1b1041570da95c952eed214740822dccb3b1"></a>

## virtual_server.connection_rate_limit_mode.per_virtual_server_source_address — virtual_server.connection_rate_limit_mode.per_virtual_server_source_address / 5b2921de7342 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-002.md#canonical-0ff1a1f4c36341bce502bae92b575dac0a88378f391590e464696488a7ce3c5f)
- virtual_server.connection_rate_limit_mode.per_virtual_server_source_address

<a id="canonical-a0d0a7b40ae0fce1696739b70b3b56491e9d72f723e8965b3465b0eea355cc3b"></a>

Type: `"single"`. Computed.

Source Address Mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-04897351931d243ce57c5c80dae707c3233176bad5d642bb1125792a2f8663fd"></a>

## Direct properties — virtual_server.connection_rate_limit_mode.per_virtual_server_source_address / 5b2921de7342 / 3

<a id="canonical-24c66a1beda73999202cc1ed8465635f646b848e29dae19f8e403291fa8fea0a"></a>

<a id="canonical-a30b962806d668b0471a34f8545364f5f075455f97ac6364ffce12f75cb43aa5"></a>

## source_mask property — virtual_server.connection_rate_limit_mode.per_virtual_server_source_address / 5b2921de7342 / 4

Type: `"number"`. Computed.

Configuration parameter for source mask.

Upstream description:

Configuration parameter for source mask

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
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-f624022e111f2481b3f5cbb65c9d20ad9eabb538884b7780fc8864a9e299124d"></a>

## Next pages — virtual_server.connection_rate_limit_mode.per_virtual_server_source_address / 5b2921de7342 / 5

- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-002.md#canonical-0ff1a1f4c36341bce502bae92b575dac0a88378f391590e464696488a7ce3c5f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-4a64939cf3af5839cb5a2004b6986bd2775a5c4ad0378e4554a709ac0afc3a1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af4ca00369c1b3417fbb4ea091fc258764288820d48aea31e2743c73bb5217f0"></a>

## virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address — virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_ / b20c4a18b816 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-002.md#canonical-0ff1a1f4c36341bce502bae92b575dac0a88378f391590e464696488a7ce3c5f)
- virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address

<a id="canonical-91572c8df33fe9baa93d7d3ba6f0098a0db67d64471b83069b17772c992a7790"></a>

Type: `"single"`. Computed.

Destination and Source Address Mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-62d0814c3146ab595816040b6800cf0d992b5ddadfb6fc57997bc10d73fabcc6"></a>

## Direct properties — virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_ / b20c4a18b816 / 3

<a id="canonical-fd7594b3979db6b09048ca2410b77c323fdd2ae7e3e339fe6bdcaddd0f7dac3b"></a>

<a id="canonical-9c6ad0390911cd558cda4de7d0cb1bd5619985b1e12f9ea68d88f9ba7d220d6b"></a>

## destination_mask property — virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_ / b20c4a18b816 / 4

Type: `"number"`. Computed.

Configuration parameter for destination mask.

Upstream description:

Configuration parameter for destination mask

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
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-a1c6aa2f3b90685145c49e3d62d0118ee1cebf08d6ee0bc0003496f1e3b278d9"></a>

<a id="canonical-0a84a54799ddab369bfcce9fd22a7af40ef087a3f4246290e64d91ddd0c45b8e"></a>

## source_mask property — virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_ / b20c4a18b816 / 5

Type: `"number"`. Computed.

Configuration parameter for source mask.

Upstream description:

Configuration parameter for source mask

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
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-e139ac307398108320736c271606ee0007965f1fd60758e69b90b5c8108b43bf"></a>

## Next pages — virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_ / b20c4a18b816 / 6

- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-002.md#canonical-0ff1a1f4c36341bce502bae92b575dac0a88378f391590e464696488a7ce3c5f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-0366d1dd265cc60d0524a3f453cf707e64e9634ca637fbca12d78eb6b8463161"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71f77c2d29f32156c62b16978bc3c76487e77638bd88a433f07eddeb52e6c948"></a>

## virtual_server.default_persistence_profile — virtual_server.default_persistence_profile / f33f80fc3366 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.default_persistence_profile

<a id="canonical-1e7938358f045952100fcca22fb3c8c5af8d26119330c252db38be6e835badc5"></a>

Type: `"list"`. Computed.

Configuration parameter for default persistence profile.

Upstream description:

Configuration parameter for default persistence profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f1881e2b4a8a487dd56e6c971cb1cd25a4b8907166be76c9e97c49ee23f04288"></a>

## Direct properties — virtual_server.default_persistence_profile / f33f80fc3366 / 3

<a id="canonical-433ea78566aeca4d9a219d8be0a14e3ba23d8390a981c6585570e1f37be5b94a"></a>

<a id="canonical-43d3d47a6f40b3aeb3bc1e20d9d7585ea4ac9fe6876bbb000ee1c27bdc1e06f7"></a>

## kind property — virtual_server.default_persistence_profile / f33f80fc3366 / 4

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

<a id="canonical-3b2bf8ce26449710b2874c472ad621e91d86858c80b7a448d6928396045141cf"></a>

<a id="canonical-38decdebf5e7a67e0e852c350be0e9de835deb3e342267ed4d3d43580e4c4a94"></a>

## name property — virtual_server.default_persistence_profile / f33f80fc3366 / 5

Type: `"string"`. Computed.

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

<a id="canonical-99388833946b4fc32c505f79c91de91f980d24a5d56112b315e919905b86061e"></a>

<a id="canonical-e421f2225fe08c4686af7ac825a5c0a7275124b15c603cd225aca1b9b1926e2a"></a>

## namespace property — virtual_server.default_persistence_profile / f33f80fc3366 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-fc7d6097b501e3cf9d8a8a32f8f7d94b42115e38bb62ab38aa8f3825d9bdf0aa"></a>

<a id="canonical-eefb510b13b754d14d530780ef70b047b565046e460952c2745f1d27d3121d4e"></a>

## tenant property — virtual_server.default_persistence_profile / f33f80fc3366 / 7

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

<a id="canonical-8158350217910ac826077e3c21b7e4dfda3f5414c6d22e670c03d6e0e3560dd3"></a>

<a id="canonical-85d3cfacce66237930d1095b24e3181b9ce1d124c53f9e80c54a50dcbb95f634"></a>

## uid property — virtual_server.default_persistence_profile / f33f80fc3366 / 8

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

<a id="canonical-fa180dbbf81aa54e465534e568e62921f9d51799b67e3fe8cd91eb979c8261dd"></a>

## Next pages — virtual_server.default_persistence_profile / f33f80fc3366 / 9

- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-3ff91b0f7877a94f4912abc43e110504ba594829ab9540f690ae63e862c83ced"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-766b05c221281e7be381adb6df44e9c2688b69fc59f2329e256685c3c27db1ec"></a>

## virtual_server.default_pool — virtual_server.default_pool / e90d412d8083 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.default_pool

<a id="canonical-0e0dc9da9e9097c5e4ed16e17dab2b180e879646d4c59aec036a91444e298f11"></a>

Type: `"list"`. Computed.

Specifies the pool name that you want the virtual server to use as the default pool. A load
balancing virtual server sends traffic to this pool automatically, unless an iRule directs the
server to send the traffic to another pool instead.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-964bac7724d7abe355dd28448ff76c9fd665f17d4bdd11e2478ab461e27b7161"></a>

## Direct properties — virtual_server.default_pool / e90d412d8083 / 3

<a id="canonical-81ea534fc69e3c233e263ab48a72f3372532fdd1b5786d1d4df426b9a5420c08"></a>

<a id="canonical-bf2a31e0ac171bb8059fc4f5d049483adefe81b8fd0b97e0063f8c4d94bcec34"></a>

## kind property — virtual_server.default_pool / e90d412d8083 / 4

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

<a id="canonical-ce9f8dbb6af102ef239352868dc3273997050e554df2aac90a958b40c4d65c5c"></a>

<a id="canonical-bcb3ad41d6392cd8dda00469ed28822d4c9c55c7dc051539e56ee5df20028a9c"></a>

## name property — virtual_server.default_pool / e90d412d8083 / 5

Type: `"string"`. Computed.

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

<a id="canonical-a75c3cedc697cf7abab9925b31c70dcfa71680e2a26c2dcb54a6beb8a15f195a"></a>

<a id="canonical-360a6e1b1645ad6873a537d33161bca7e1d5e875b855359fa8c95383f18c9b96"></a>

## namespace property — virtual_server.default_pool / e90d412d8083 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-91b6f0c64af0da8fffb870d3c081db8ae5677377e48a02494b5647247c1a297c"></a>

<a id="canonical-c16836ef4656cf2a72d4796d948b56b1aa4fd6259ec2cde58fbff93d87553470"></a>

## tenant property — virtual_server.default_pool / e90d412d8083 / 7

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

<a id="canonical-5254c7e98f59b24713b906c8912f3c18b37f03895e16fc0258230db02332afdf"></a>

<a id="canonical-1e84bbdfdf012832804716a25894fb7a7834acfc1596241c9c582372be8955a2"></a>

## uid property — virtual_server.default_pool / e90d412d8083 / 8

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

<a id="canonical-2a329cf783204aeaf37e94201c4fcb4a78d2246eb25221095c27b402f13eb036"></a>

## Next pages — virtual_server.default_pool / e90d412d8083 / 9

- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-f7bc990e9c3e8f027c8654d878046f88f3442028146200f0ba190a9e1bf9222c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a5a55e00dc80ecf371b86d481b471bdd2986b957511ecc97ddaba1ba874a3c5"></a>

## virtual_server.fallback_persistence_profile — virtual_server.fallback_persistence_profile / 24ba4025fabf / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.fallback_persistence_profile

<a id="canonical-cda90e7c0c5a0f593ed8e46c75cb5c5efdb09d15e5b64f090866401902b39b0d"></a>

Type: `"list"`. Computed.

Configuration parameter for fallback persistence profile.

Upstream description:

Configuration parameter for fallback persistence profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-7267a3dab925c816d66015462bdf438413be85383aaea6430e937fc6a4969bc7"></a>

## Direct properties — virtual_server.fallback_persistence_profile / 24ba4025fabf / 3

<a id="canonical-0cc90ce266fecdd5334b561ee9ca37ae03b208c3d7e98bd434abef684736fef3"></a>

<a id="canonical-574f092d5d9874d092acf362ca2763b2b7aace477392cdfa5d2d559d521c9a5a"></a>

## kind property — virtual_server.fallback_persistence_profile / 24ba4025fabf / 4

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

<a id="canonical-8557cc993fa464c9b0a66b2a42d60a9b98ca1e4a0b5e1d2eea43618c5838e349"></a>

<a id="canonical-06b1a7726aae07619b2307df21ae1a66d1f6e6524ac541922d57974ee7647fb0"></a>

## name property — virtual_server.fallback_persistence_profile / 24ba4025fabf / 5

Type: `"string"`. Computed.

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

<a id="canonical-33cb47cece85343d84201118fd3393b306aec103f92dad1a438a07bdc171f954"></a>

<a id="canonical-6ed75b4d40955cd17e2f21884d95b81276e81cd8d41ccd180e64c9772ecf082c"></a>

## namespace property — virtual_server.fallback_persistence_profile / 24ba4025fabf / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-e13c1afcd9264424301c7b8aad7b0b2b7623b2ab7c6e472338a6cb2f8a8cdcd4"></a>

<a id="canonical-182361ace6150f1743361e201ffe89e45e089da2e1b60b49675895db9ed93f90"></a>

## tenant property — virtual_server.fallback_persistence_profile / 24ba4025fabf / 7

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

<a id="canonical-6ff9953699b348d69aa3a02c03109a3f17bea350534e9e9b0948231c2c762856"></a>

<a id="canonical-b1b1aacdd37b340c645e40aab5db71664769128ae4e67ca16452c3da731eebf7"></a>

## uid property — virtual_server.fallback_persistence_profile / 24ba4025fabf / 8

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

<a id="canonical-215494e08cc6d12c16e63ef94307dcb207086b5375a69daa4e4171735eee264a"></a>

## Next pages — virtual_server.fallback_persistence_profile / 24ba4025fabf / 9

- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-fe1e80eb31e642e61346da9d7773a14ffdb260dee3a620e9e85a918fd920f453"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4c7c11d6b19ad29c58314d1f98f8fde913dfb6d5d7bac31a626a96fd7a3a3b8"></a>

## virtual_server.fix_profile — virtual_server.fix_profile / 1fe617187e37 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.fix_profile

<a id="canonical-bfc05c64e7b23d9747294460ce2e0e992fc21840ad29934e049e9f688a7d631c"></a>

Type: `"list"`. Computed.

Configuration parameter for fix profile.

Upstream description:

Configuration parameter for fix profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-853822edb3647d745b329ac5e403eb43d92b27f413d8adedd97d9f44c2d00396"></a>

## Direct properties — virtual_server.fix_profile / 1fe617187e37 / 3

<a id="canonical-5aeff7824d7cb8b2081f461cb2b3eb09d7d3423147fb0377392bbd0d67e55ee8"></a>

<a id="canonical-f42088854c16105557559c30bff1402f5aa3353453dd415ecae8e31f51aba678"></a>

## kind property — virtual_server.fix_profile / 1fe617187e37 / 4

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

<a id="canonical-629d214d39851e606c254872a7485648b1609df8fb74fb7bfaed1ea5425fceeb"></a>

<a id="canonical-3accded2b005de5dcd1110b1ae935b86770f39b52e6f0ca0b1070a49ccc5f468"></a>

## name property — virtual_server.fix_profile / 1fe617187e37 / 5

Type: `"string"`. Computed.

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

<a id="canonical-4a9de7edee40ef15dad7c2e132c685fcd6856047c370094dd49da576ee8a3c21"></a>

<a id="canonical-7f7000bfe6e21d902eeb46f76b72a42d08d7019d31172c5a2055b410995cb0ec"></a>

## namespace property — virtual_server.fix_profile / 1fe617187e37 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-a485e75275af0ae1b06f0fca25fe574be8c9611b2d4f8b1f7ac7eb6c7e3b8b1d"></a>

<a id="canonical-4e85fe9d3987035379d7adb0da8590bd9ae2717f22aac0064b606165e8174d2a"></a>

## tenant property — virtual_server.fix_profile / 1fe617187e37 / 7

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

<a id="canonical-31452fb8d1876d15f68b4c9c4c9d5b82062803be61f4c1f197d85eeeda5ff6bd"></a>

<a id="canonical-c7e1b4fa63c445727bdc5c88eebe3e3a77bd90e67c4701b017a876e24078dd81"></a>

## uid property — virtual_server.fix_profile / 1fe617187e37 / 8

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

<a id="canonical-5a0e43905fa8e172f7bc748e62681161f704407905bd27f09cb48588d8686811"></a>

## Next pages — virtual_server.fix_profile / 1fe617187e37 / 9

- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f72d236dedf1e64e6963372eae9801743f9f51c980f244121ba43551dda5d0f"></a>

## virtual_server.http — virtual_server.http / 6278190634af / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.http

<a id="canonical-3e91f022ac2cc4b9431b39f5f986bd64a35359eb0043e2e380f955e3249e0712"></a>

Type: `"single"`. Computed.

HTTP profiles.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-12fbed29f4a4eb0e7f6ad1f72245072becf818ec61b4a146b1974242a1db9e3f"></a>

## Direct properties — virtual_server.http / 6278190634af / 3

- [client_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-01ebbcdf7f7c95b242b010c451b9fe50b9d732caa4083553c0c9797d9a6dc50b): complete subsection reference.

- [http2_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-a98584269a16df3481a9b4d1824fb4335dcf635cbe8646bfc209488276b045cf): complete subsection reference.

- [http2_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-65c61057ec3c45557904a7bce01f3d117b79f6bfb1352cf6e16a3878aca7a369): complete subsection reference.

- [http_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-028dcb61add03158ce4ca0f60adae3c5a6131b3b907984e5f82b5c3e0b112e4e): complete subsection reference.

- [http_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-9752b0a7a6c6a4aa0fc53feb8aed948334f2202f9dd5f9596b7a7bf6bb222ec4): complete subsection reference.

- [ocsp_profile](data-sources--application_profiles--reference--group-002.md#canonical-961a7596d157c6770105e18495ca336bf74c94e105d7c04398618ed3166e68a3): complete subsection reference.

- [server_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-458304cc1446f2d95fc7869afa9db53e8258c1f3012a4b0cb3e3c8fd0a486743): complete subsection reference.

- [stream_profile](data-sources--application_profiles--reference--group-002.md#canonical-1bb202e19e008d1055dbad6f8baa6e957fb69f30653c2453cdd58597b7392712): complete subsection reference.

- [tcp_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-aabf4e8cdcc5464df575b51291e8753c47fd7e645aadda772fbae9dbed91a772): complete subsection reference.

- [tcp_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-deae319c44458b60a8195c738bc85666a1ece52ca21f8b3700532265ba45af07): complete subsection reference.

- [websocket_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-a39a8f21a531c381b5984d77711593e653ecfec0ad9441ca3459fded11f1be4a): complete subsection reference.

- [websocket_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-001df5346292404bb45ad860e9932722005765a6f0ec139fa2c476efc7409272): complete subsection reference.

<a id="canonical-213a445b45256663516616ed46fcf8e041addee6326878d6a503114dc8b8c490"></a>

## Next pages — virtual_server.http / 6278190634af / 4

- [virtual_server.http.client_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-01ebbcdf7f7c95b242b010c451b9fe50b9d732caa4083553c0c9797d9a6dc50b)
- [virtual_server.http.http2_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-a98584269a16df3481a9b4d1824fb4335dcf635cbe8646bfc209488276b045cf)
- [virtual_server.http.http2_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-65c61057ec3c45557904a7bce01f3d117b79f6bfb1352cf6e16a3878aca7a369)
- [virtual_server.http.http_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-028dcb61add03158ce4ca0f60adae3c5a6131b3b907984e5f82b5c3e0b112e4e)
- [virtual_server.http.http_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-9752b0a7a6c6a4aa0fc53feb8aed948334f2202f9dd5f9596b7a7bf6bb222ec4)
- [virtual_server.http.ocsp_profile](data-sources--application_profiles--reference--group-002.md#canonical-961a7596d157c6770105e18495ca336bf74c94e105d7c04398618ed3166e68a3)
- [virtual_server.http.server_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-458304cc1446f2d95fc7869afa9db53e8258c1f3012a4b0cb3e3c8fd0a486743)
- [virtual_server.http.stream_profile](data-sources--application_profiles--reference--group-002.md#canonical-1bb202e19e008d1055dbad6f8baa6e957fb69f30653c2453cdd58597b7392712)
- [virtual_server.http.tcp_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-aabf4e8cdcc5464df575b51291e8753c47fd7e645aadda772fbae9dbed91a772)
- [virtual_server.http.tcp_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-deae319c44458b60a8195c738bc85666a1ece52ca21f8b3700532265ba45af07)
- [virtual_server.http.websocket_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-a39a8f21a531c381b5984d77711593e653ecfec0ad9441ca3459fded11f1be4a)
- [virtual_server.http.websocket_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-001df5346292404bb45ad860e9932722005765a6f0ec139fa2c476efc7409272)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-01ebbcdf7f7c95b242b010c451b9fe50b9d732caa4083553c0c9797d9a6dc50b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8cdc0ab154b392ec22b762d2f2aad00dce95ad887afb4f5d97a2d08e5c6d84ec"></a>

## virtual_server.http.client_ssl_profile — virtual_server.http.client_ssl_profile / 767b330ecf2f / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- virtual_server.http.client_ssl_profile

<a id="canonical-1b99ccfab4dcf31c2422084b09e6dc049efaac7c519638faf63e13a45c25d15e"></a>

Type: `"list"`. Computed.

Client SSL Profile. Client-side configuration

Upstream description:

Client-side configuration

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

<a id="canonical-66a8fb0bc20cfb4b4251a0621ecb563640b0670ceb539c408812e0cdfbc49769"></a>

## Direct properties — virtual_server.http.client_ssl_profile / 767b330ecf2f / 3

<a id="canonical-3c9b156b7de194c526a45972f3f772c860f15f7312cd83c1e7d18dbdfee4fbee"></a>

<a id="canonical-30d6de327e4a7c2aa6ec7ec09c8116329f87178f74ddd674725aed4e73b75242"></a>

## kind property — virtual_server.http.client_ssl_profile / 767b330ecf2f / 4

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

<a id="canonical-711c38fbd9fd7b68dadeab50257e8e709c2f2c04cf99d18451f7c0c3f88deda5"></a>

<a id="canonical-477b1155c238044fd483addcde6dbf1f6d75a50550bf03db06425f114980cee8"></a>

## name property — virtual_server.http.client_ssl_profile / 767b330ecf2f / 5

Type: `"string"`. Computed.

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

<a id="canonical-e2415dcd4f39a391695b994be809702d4938fe4e544a698dc4f68fcc2a77ff3a"></a>

<a id="canonical-862e66b646acdd2f8744057aa0899e02a11236d3050a3af682dabd25852c2466"></a>

## namespace property — virtual_server.http.client_ssl_profile / 767b330ecf2f / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-056b00b50a5c63c710215266e3dada4b7d9394ee3ea2b30c279032d9160b7a6a"></a>

<a id="canonical-9d2a842eae7d9383326f902bce8578ce698a427dcf41dd491fd7334c7b16a8a1"></a>

## tenant property — virtual_server.http.client_ssl_profile / 767b330ecf2f / 7

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

<a id="canonical-c364ae7f9dd008ff803a51590348c2c3656d0559676701bf2c7b1fa0473f1f41"></a>

<a id="canonical-b977ec6302e8f8a65bbd972e6ddb10d3b89ba035eaf9060de08143dfd53e5012"></a>

## uid property — virtual_server.http.client_ssl_profile / 767b330ecf2f / 8

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

<a id="canonical-96fff529c33ccb0b55d5790231a9a2481aadf88ac44b65251d4938bbe2a1ef6e"></a>

## Next pages — virtual_server.http.client_ssl_profile / 767b330ecf2f / 9

- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-a98584269a16df3481a9b4d1824fb4335dcf635cbe8646bfc209488276b045cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c4342171c47d63618e8a1d625ae50a47bbfef62a6fecea5185cd2cdae6e4837"></a>

## virtual_server.http.http2_client_profile — virtual_server.http.http2_client_profile / 3dd2dec7c43c / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- virtual_server.http.http2_client_profile

<a id="canonical-06c6df1e540c4dd546442e888f4da2c896616ef789c8e966db6ef71e526fa616"></a>

Type: `"list"`. Computed.

HTTP/2 Profile Client. Client-side configuration

Upstream description:

Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-b736fa1b5a51fe9dee07a2c5abfaa5e8eae23b876a3ed8d603b7e388082bb7f2"></a>

## Direct properties — virtual_server.http.http2_client_profile / 3dd2dec7c43c / 3

<a id="canonical-59e165efb900b4edee109bff124bde6e4da17e3620011aa3b609514e7dfe716a"></a>

<a id="canonical-dc1c6e8d2fa5a6458384d6675049d987206a5091c94f30ef035564469b369c1e"></a>

## kind property — virtual_server.http.http2_client_profile / 3dd2dec7c43c / 4

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

<a id="canonical-3f5e855d20df012a4f858e4c3f12a1055f7a2f9c6b1f019578bc5b3e7ed52c52"></a>

<a id="canonical-f26603d0e1dfa3114c1bdb029881d173ec37ff585bb6cb54da323a4faa3abfe9"></a>

## name property — virtual_server.http.http2_client_profile / 3dd2dec7c43c / 5

Type: `"string"`. Computed.

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

<a id="canonical-c56bb928b0b39278ccd126936a34f226c0531323dbd6225056228cb71ccd10dc"></a>

<a id="canonical-036475d76944c4471724958fdfff7bfebab63e1be447dc1a036189445a500b85"></a>

## namespace property — virtual_server.http.http2_client_profile / 3dd2dec7c43c / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-0b7a3703d03d128cafa33019ff4989cc3e244d3d38c7ba5a4b6c2b7f52c629d2"></a>

<a id="canonical-4a1616e074b56948e19c9c6c2a8906d42a1bababbe045ae072dd000da4f1e1cc"></a>

## tenant property — virtual_server.http.http2_client_profile / 3dd2dec7c43c / 7

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

<a id="canonical-f489a3f43f4cece41581224b05b40a327a72fa743dc76252da452f893b6c1b89"></a>

<a id="canonical-2fb85506ddf03247a9e75e605131f83271d5178126fdb1a4fff63a6e4d7e01ac"></a>

## uid property — virtual_server.http.http2_client_profile / 3dd2dec7c43c / 8

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

<a id="canonical-f71ef67c3f0a57bf36d5fcd303e33d76bdbdec48dfc91a7aa09f1bb2d9d9216a"></a>

## Next pages — virtual_server.http.http2_client_profile / 3dd2dec7c43c / 9

- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-65c61057ec3c45557904a7bce01f3d117b79f6bfb1352cf6e16a3878aca7a369"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7304256af7373f7d40ebf791a8ea483f30b31ad2013c749ab964221d10a15ecc"></a>

## virtual_server.http.http2_server_profile — virtual_server.http.http2_server_profile / 57f011e171f9 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- virtual_server.http.http2_server_profile

<a id="canonical-d0a400e20566b3e910996d9cbe1f9e9d857be08e131bf3bd0fcfd634badcf0a5"></a>

Type: `"list"`. Computed.

Configuration parameter for http2 server profile.

Upstream description:

Configuration parameter for http2 server profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-26626f780f7fd5a6e2ba5c561ede29920b7ccdd82578a0d4f12747ee12c0b1e9"></a>

## Direct properties — virtual_server.http.http2_server_profile / 57f011e171f9 / 3

<a id="canonical-c38726d21def8650362c2f3527826c47e0ece7f33877dd0bc5b997dc3fdab994"></a>

<a id="canonical-609dda77b4c29e357b5a21ea2d5f1b7beb717a2474a87fc40c42f582c835bda8"></a>

## kind property — virtual_server.http.http2_server_profile / 57f011e171f9 / 4

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

<a id="canonical-3b03ff90266efa29472fe8bc7279e2bdf47cc7388353bcecd506743f6b80891f"></a>

<a id="canonical-8c4a263b6743a32ddc0969751b6e0461fbf34f772feccb32b2da008f02025ae0"></a>

## name property — virtual_server.http.http2_server_profile / 57f011e171f9 / 5

Type: `"string"`. Computed.

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

<a id="canonical-8b7c4b7eba5c6a33eb4efb02a8d9f429e454606743d635caee9e74b55668cb6e"></a>

<a id="canonical-cfd6bb86f17d74b9d973d09cf872601bebb63a9828af8fc92ac5a1feaddaa39d"></a>

## namespace property — virtual_server.http.http2_server_profile / 57f011e171f9 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-ac8455c8336f13ff3fb9a1c97679c4257fea278bec09a0f23dc89134ff31b1b6"></a>

<a id="canonical-24429d078fa82d7e6554b8053d66359a74d91549e911432f0b050d0f56b5edc8"></a>

## tenant property — virtual_server.http.http2_server_profile / 57f011e171f9 / 7

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

<a id="canonical-29141dd2a1dfe3b767d738f453a9523e5d0226c49d180f898aa32a4394aef805"></a>

<a id="canonical-70095bc1ed297fad42cc5f33d00e0bd579ce0c06e20235f80ad611132cc17c9d"></a>

## uid property — virtual_server.http.http2_server_profile / 57f011e171f9 / 8

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

<a id="canonical-6ca6cc5c33af8ad4e9a1b03a0851094b8889851362365c5fba8e3eb6fc0da2b9"></a>

## Next pages — virtual_server.http.http2_server_profile / 57f011e171f9 / 9

- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-028dcb61add03158ce4ca0f60adae3c5a6131b3b907984e5f82b5c3e0b112e4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca937f483b7165ecb26929b080ac746f5d08583300dad63626ffbac2a8bea532"></a>

## virtual_server.http.http_client_profile — virtual_server.http.http_client_profile / aa369c4cf6d3 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- virtual_server.http.http_client_profile

<a id="canonical-1a9c78f5056622fe56eb7018efe01d68405639a8d3479713e9551510945d5449"></a>

Type: `"list"`. Computed.

HTTP Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-b4032caf7ea7700beb897e058387224c4786bc187634b6255e045aa6f3381062"></a>

## Direct properties — virtual_server.http.http_client_profile / aa369c4cf6d3 / 3

<a id="canonical-dcac4f14d5d7401edc7d714e03d40e34a11b49fd078f7b37e8103e91928b6cd8"></a>

<a id="canonical-b81c6ef25913d160389174e5a9ebaca6bb8095c7e6a48ae39da0c7f08c633a34"></a>

## kind property — virtual_server.http.http_client_profile / aa369c4cf6d3 / 4

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

<a id="canonical-8c10e3b32b4c6cd179b0ea7e3dfa98a81cf7ba6d2e9b5cfd898fb6e311ef8f07"></a>

<a id="canonical-2b21fb72a44dae9c5c2fe6cb725c7c63822706107d09cfc597406eaa6239935a"></a>

## name property — virtual_server.http.http_client_profile / aa369c4cf6d3 / 5

Type: `"string"`. Computed.

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

<a id="canonical-f2fc114ebd2b26c042bb336937faec1a1d9126e951c5db8766061928ebce2272"></a>

<a id="canonical-bfe1608123793cab67fc21583a63837bf05ab299838e2b1255e311bc996d86a8"></a>

## namespace property — virtual_server.http.http_client_profile / aa369c4cf6d3 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-9374dac7e768ff91a3e7807769171c05f16c522a624d638166be91b38b892193"></a>

<a id="canonical-8a52403cfd1d06b7969c8ceb68d674ef5b98fac06fa11da0896dc205709e01e3"></a>

## tenant property — virtual_server.http.http_client_profile / aa369c4cf6d3 / 7

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

<a id="canonical-0a32a30e69f7dc3a77363d45ab11ba22fc0be5bbcd730ebd3e5f162119f0c0aa"></a>

<a id="canonical-4f258ac21bb38a3bc1cf3ee30648e9177123e38958f04222cce7e3680e31718b"></a>

## uid property — virtual_server.http.http_client_profile / aa369c4cf6d3 / 8

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

<a id="canonical-9e09de2351c0cfb433be76ccfda0ac30b89a0164e697f7906393f649193a9235"></a>

## Next pages — virtual_server.http.http_client_profile / aa369c4cf6d3 / 9

- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-9752b0a7a6c6a4aa0fc53feb8aed948334f2202f9dd5f9596b7a7bf6bb222ec4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58a966f5a915e242bf20b4667bd8ab68df22d5a4aec174b40d7c438d47e3c5f1"></a>

## virtual_server.http.http_server_profile — virtual_server.http.http_server_profile / d5be7036dff6 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- virtual_server.http.http_server_profile

<a id="canonical-aa02b94dab5d90b52a777578cff8c404b6dbb4c4b331f51eb01d2d7cb6d078c0"></a>

Type: `"list"`. Computed.

Configuration parameter for http server profile.

Upstream description:

Configuration parameter for http server profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5c5030976f1aa756154562f0dce00cddb23489684a7d89dde90ffa4560f27a63"></a>

## Direct properties — virtual_server.http.http_server_profile / d5be7036dff6 / 3

<a id="canonical-28218f8978ae4c8a86401c2624d3a86ddd6b3e65cdb206bd37b0bebf8385591f"></a>

<a id="canonical-3a1a68a00a9916b5ec8385d574c5cbca5d8baa02b7b427627a8f2371c9243c9a"></a>

## kind property — virtual_server.http.http_server_profile / d5be7036dff6 / 4

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

<a id="canonical-3c2a5fd59f8e1ca48b7d3b930358be4f621061d918c342b30a461f9cb507c99b"></a>

<a id="canonical-721db043c233661617f036d829fb1dd74059b01b2aa913a8d4e0b9c5a904036f"></a>

## name property — virtual_server.http.http_server_profile / d5be7036dff6 / 5

Type: `"string"`. Computed.

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

<a id="canonical-295f4d50b395e1549e10dfeee6698fbcdef1a36ca7bd10a10943087eb03a428a"></a>

<a id="canonical-4d62d2a0de6bc43227625d5f51e99071cfd465af2fd0d5242479510341a94528"></a>

## namespace property — virtual_server.http.http_server_profile / d5be7036dff6 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-5bb3dea00e808d20f0d57e25fff056710e7706cbc655d1094dcc318fc2377230"></a>

<a id="canonical-391c81a26f49b34a40dea844c91e5c96a518ad6f5c459d510532faf4e9832ecb"></a>

## tenant property — virtual_server.http.http_server_profile / d5be7036dff6 / 7

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

<a id="canonical-95c824cf1cbf084fb538662682faffdae12fd79a6dc853c809bb2c4d7b5c8b7d"></a>

<a id="canonical-e74a8d344f174ea879476bc20b267dfbd0e63c1504b0d03557264a4fdf4ce4ce"></a>

## uid property — virtual_server.http.http_server_profile / d5be7036dff6 / 8

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

<a id="canonical-be075896e0f4b2c38803338d7d66cb775c2f8321642ce705820945f0d73bb727"></a>

## Next pages — virtual_server.http.http_server_profile / d5be7036dff6 / 9

- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-961a7596d157c6770105e18495ca336bf74c94e105d7c04398618ed3166e68a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46fca89dffb633d795c9570b31749d9905521a158e76872b93b928a0aa9f6f04"></a>

## virtual_server.http.ocsp_profile — virtual_server.http.ocsp_profile / b3f8b3964d83 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- virtual_server.http.ocsp_profile

<a id="canonical-3685b833adb363b1381de77d94856efc3319543397d6ead25e84e9b52b5c2d9a"></a>

Type: `"list"`. Computed.

Configuration parameter for ocsp profile.

Upstream description:

Configuration parameter for ocsp profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ce2f9b443acb697c6f9401f40b769fa0b44a00434b61935363948f53ef66282b"></a>

## Direct properties — virtual_server.http.ocsp_profile / b3f8b3964d83 / 3

<a id="canonical-6ce506af7c30ef5245e5ed626c87f056c87490c2b739245a8f215c6ff4a81fc2"></a>

<a id="canonical-e1ef42a64f8e558b9bdaf1531893081b353157e033f7c81388f9c53e842a0877"></a>

## kind property — virtual_server.http.ocsp_profile / b3f8b3964d83 / 4

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

<a id="canonical-b060b151431754ea49f5c55314dbfbd9a452ffacfb8eb267c5eec6fe57643a17"></a>

<a id="canonical-4d51ba14f0222d1546b59564323dfadfa8a4b4d00a2f7210acde7e46bdfeddce"></a>

## name property — virtual_server.http.ocsp_profile / b3f8b3964d83 / 5

Type: `"string"`. Computed.

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

<a id="canonical-47b08256932f3ee177d5e49867ca4ef96d1e2d7ca9140e7a2f39191695416afe"></a>

<a id="canonical-c07cd968cff96c37135119df5b42c8920a351c9c07cac6dc857ca8a84332543b"></a>

## namespace property — virtual_server.http.ocsp_profile / b3f8b3964d83 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-1028207a49e7e5abf5a7d10ac8df3dc80211692249018ac3e699e989b3d0f044"></a>

<a id="canonical-a1d08332cb7b3d3f1af02e5e22b29236554ca02a1cb79ce4924c929216ad4b6d"></a>

## tenant property — virtual_server.http.ocsp_profile / b3f8b3964d83 / 7

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

<a id="canonical-064e8993dfe95c8fd5511e4f9fa7c1820921a3ab953eece18f2b046dfaa2953a"></a>

<a id="canonical-6898d46c4630ba17698cfcca666daf32b0c3a45447871b938ced3a0c0cc2fcb0"></a>

## uid property — virtual_server.http.ocsp_profile / b3f8b3964d83 / 8

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

<a id="canonical-022825eb0dda63496f469a1589f1d541a58e7ca2509486e0503a25da6f572c7e"></a>

## Next pages — virtual_server.http.ocsp_profile / b3f8b3964d83 / 9

- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-458304cc1446f2d95fc7869afa9db53e8258c1f3012a4b0cb3e3c8fd0a486743"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33fb8b3789605a336734e23552578ec7ff5403ebd0915ea861d7803cb0ad5d57"></a>

## virtual_server.http.server_ssl_profile — virtual_server.http.server_ssl_profile / 047a2a9e34d0 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- virtual_server.http.server_ssl_profile

<a id="canonical-4224d0dbf16bde64a9f8a59817118481976c1baa8e977ff924313965b55cb1d2"></a>

Type: `"list"`. Computed.

Configuration parameter for server ssl profile.

Upstream description:

Configuration parameter for server ssl profile

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

<a id="canonical-8fb5294b74e855bd92f9053189369b203858c1832a925bee79041cc39c104833"></a>

## Direct properties — virtual_server.http.server_ssl_profile / 047a2a9e34d0 / 3

<a id="canonical-f519bad014b7b52c25d60a78e94dffc1263bbbc49c24fd8acc308671c27993a4"></a>

<a id="canonical-b210fec1a8a73c5a20d08912e59b0354f7844598e8e2678c356788ca7144f072"></a>

## kind property — virtual_server.http.server_ssl_profile / 047a2a9e34d0 / 4

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

<a id="canonical-4c27d911a00ba625bb1f48986327bd1c64f50bbcd834f783bd29c0e21a3051e3"></a>

<a id="canonical-f242966445e35f8ca53820c2f9ec8ee207a3c10a23ae4a7994e1e04fc03206dc"></a>

## name property — virtual_server.http.server_ssl_profile / 047a2a9e34d0 / 5

Type: `"string"`. Computed.

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

<a id="canonical-d76262d1b0fbe848118c8019a1d2b5ad0186a232703c29250ae79ae0da4d8b40"></a>

<a id="canonical-e81bb24f3779bd3fdb7213fffd1b51f4ef52979fed9ea113366288f8da03b8cf"></a>

## namespace property — virtual_server.http.server_ssl_profile / 047a2a9e34d0 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-cdb715882dfb96d730a73fe0711ae6220913e1e1d90c0c78d971ddda21e08ceb"></a>

<a id="canonical-1b5247039b44f90948507f7fab0a1dc9481ea10985c4c23037853a343dc1f176"></a>

## tenant property — virtual_server.http.server_ssl_profile / 047a2a9e34d0 / 7

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

<a id="canonical-5a196d42a153999a79504d8b24f348d948b6755d0acd5c1dad6db3e01c80a0d3"></a>

<a id="canonical-6eefcf85e62d627d4c5b574d0e6222f2cf30cd262ed41af44692c548103419ac"></a>

## uid property — virtual_server.http.server_ssl_profile / 047a2a9e34d0 / 8

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

<a id="canonical-7ee160634d0778e3e0d07916dea39d82ee5d1184fb9c60dbaa64c80c446f22e6"></a>

## Next pages — virtual_server.http.server_ssl_profile / 047a2a9e34d0 / 9

- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-1bb202e19e008d1055dbad6f8baa6e957fb69f30653c2453cdd58597b7392712"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7cd5c36c62a95cbac1022b614de640a970fc0c00a952de0646cbc82caa4cff2"></a>

## virtual_server.http.stream_profile — virtual_server.http.stream_profile / 016d3039416b / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- virtual_server.http.stream_profile

<a id="canonical-de9e1be889e52150a81b6d319541f21a17a83cb924b0094b266f36463cce7948"></a>

Type: `"list"`. Computed.

Configuration parameter for stream profile.

Upstream description:

Configuration parameter for stream profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-bb0be3c8a49195f6e380b9929e17d5e2a02897d17319d1ca019a474749d8264b"></a>

## Direct properties — virtual_server.http.stream_profile / 016d3039416b / 3

<a id="canonical-2c024fb5eaf1996d3740b872d9effb256c95aad4045b2a36b738aab2d7536885"></a>

<a id="canonical-e6539a296303e4f7799bd214eeeef6efce2f556c6ec95ed1653421886fae972f"></a>

## kind property — virtual_server.http.stream_profile / 016d3039416b / 4

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

<a id="canonical-b7e93084dd7b97c9ba6a8c19bfc6cdf48fa5e2794a08c124a5a0cb348d422367"></a>

<a id="canonical-4e63c8cc17eae76f679d7fecc371f30a0e34f8bfed2620709049d700026fdf16"></a>

## name property — virtual_server.http.stream_profile / 016d3039416b / 5

Type: `"string"`. Computed.

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

<a id="canonical-167aac626058f3d2b9d4c84e58ebb55b005384154d48e001e99d605dc07e5e56"></a>

<a id="canonical-6d6caa970d2437b840a2f15a1dd2567a3f753f495d3feac3f495f163f92ff9d1"></a>

## namespace property — virtual_server.http.stream_profile / 016d3039416b / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-25cc2c67df7661e38433d08e796c5c61840c7c1f63d669607ae8773975fc227f"></a>

<a id="canonical-02a32e66ae197610da4558dbd66309e01d2b11332265478a793beafee2269225"></a>

## tenant property — virtual_server.http.stream_profile / 016d3039416b / 7

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

<a id="canonical-17b7483175589887d8b915dbc7f9a87a1b770030a2dcc83bd63ada74a40b493d"></a>

<a id="canonical-6b18c0cc97b0b57849c78566386ef0499f3199b7f0b7b05c078daf7505d64115"></a>

## uid property — virtual_server.http.stream_profile / 016d3039416b / 8

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

<a id="canonical-dcb80cedf33386882f64e334e61dfcc26ce1c232fb2bf8d1700fcefd7265c40a"></a>

## Next pages — virtual_server.http.stream_profile / 016d3039416b / 9

- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-aabf4e8cdcc5464df575b51291e8753c47fd7e645aadda772fbae9dbed91a772"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d85a7fddbdfb0e39aa4ed1259302ebd11683afecfec4367bec9b5041c721c68d"></a>

## virtual_server.http.tcp_client_profile — virtual_server.http.tcp_client_profile / 1c00ba4f80e0 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- virtual_server.http.tcp_client_profile

<a id="canonical-c51b4e45bfd112c205832f158012f507f8410becb123841f1e3c181b17435e41"></a>

Type: `"list"`. Computed.

Protocol Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-532c894ce516c763c40ea754f9c82c86c43b3fb05d2a5497c45ceda2f67547e8"></a>

## Direct properties — virtual_server.http.tcp_client_profile / 1c00ba4f80e0 / 3

<a id="canonical-66d48c18ebd73937c21b69b14acc7dd7d0268116ea7326f48463a7b03b98be06"></a>

<a id="canonical-cae6929f4ee4e4ce755ef86ea5c663a138259fa7eddd9e785cd10812867dd9a5"></a>

## kind property — virtual_server.http.tcp_client_profile / 1c00ba4f80e0 / 4

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

<a id="canonical-fe7aa3cfb10df9f967570bf566b82b64941685c63b0f8026d5c289170a078ddc"></a>

<a id="canonical-afe3dce745419acaf47142aa503b1719a0ad597ec85f36c42c8e2d0a4afa4cad"></a>

## name property — virtual_server.http.tcp_client_profile / 1c00ba4f80e0 / 5

Type: `"string"`. Computed.

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

<a id="canonical-e8f5eebfecbe1ba7f3ea15b0299f23bf92256098b1f3c44f9dd468c36360c89e"></a>

<a id="canonical-4c55216766580d903eba2afbac635c75b3fab0484cc94639aa20df9aabf7b4f7"></a>

## namespace property — virtual_server.http.tcp_client_profile / 1c00ba4f80e0 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-5fdfd73ebad5b844f8665936c8d9c947bb187a0dced74a87cae9deb1b483f772"></a>

<a id="canonical-44fac8ed9befcfdf857660ca23f6a7b4765f5c6403cc536ee3ab7deaa6272b6f"></a>

## tenant property — virtual_server.http.tcp_client_profile / 1c00ba4f80e0 / 7

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

<a id="canonical-4f920a894ad3f619bcb1a3f38d09edac59280074a7ee06db65e0b831d5149f2f"></a>

<a id="canonical-4c0a8bb3d8af990cb39d5c79c775bc44f2faf049b98ae4817696ca3c695cbadf"></a>

## uid property — virtual_server.http.tcp_client_profile / 1c00ba4f80e0 / 8

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

<a id="canonical-2a7462073d0535750db12f355c536ece96adec739539cbbe4699ec0759a3900d"></a>

## Next pages — virtual_server.http.tcp_client_profile / 1c00ba4f80e0 / 9

- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-deae319c44458b60a8195c738bc85666a1ece52ca21f8b3700532265ba45af07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5005fc6f5455446898ecc2db41846f0a2ae05a4aa97aa310fbe41a5ae1aec008"></a>

## virtual_server.http.tcp_server_profile — virtual_server.http.tcp_server_profile / 31c44aa6f101 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- virtual_server.http.tcp_server_profile

<a id="canonical-4df66e759912115ff2c98d3352738a3ac31842b64b38ea894abe97cabc3ff691"></a>

Type: `"list"`. Computed.

Configuration parameter for tcp server profile.

Upstream description:

Configuration parameter for tcp server profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0180facef6c0f01069b4e8a6b84eb8b3ab72054cc96eff212733cd098b93ef7a"></a>

## Direct properties — virtual_server.http.tcp_server_profile / 31c44aa6f101 / 3

<a id="canonical-7bcfaf6b78e08a7cc9f2e1eabf0c6bb9ad017568837b2ce5ae9c58e054265ef1"></a>

<a id="canonical-a7d8701a287eda1ce46bd2b716ca2f5e4eedcf25ea1542664cb0e7221d1d6a00"></a>

## kind property — virtual_server.http.tcp_server_profile / 31c44aa6f101 / 4

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

<a id="canonical-f0f224ea108def99ee8ac8ee9d0586ffcae3a965af818d6ac842877e591e7612"></a>

<a id="canonical-cf497a5230abb5c0310886e2426b40761d6b8183e83d3b4b05e059e7d84fdfb5"></a>

## name property — virtual_server.http.tcp_server_profile / 31c44aa6f101 / 5

Type: `"string"`. Computed.

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

<a id="canonical-38d3a103eb7ddcc05f995d7aa09f13f0f66fda7cfa3da772c61ed97f13686976"></a>

<a id="canonical-33341b76a8fd00967594df60edff818906d9bdcc8fef0121fc3f9685c7c13d32"></a>

## namespace property — virtual_server.http.tcp_server_profile / 31c44aa6f101 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-dbca43842ede3bd74794df567f8bc3485eca859d80fb324708d234886ebc5797"></a>

<a id="canonical-f0db69686a50db2d8b6432d0b8e77f4668f3dbab0c5bd7477d7d52ed43ee2bbe"></a>

## tenant property — virtual_server.http.tcp_server_profile / 31c44aa6f101 / 7

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

<a id="canonical-14afc38dc1eecda70db78953b834f0a448c2c5792b93c25735fbbdfb898b59c6"></a>

<a id="canonical-e7d56f1ad1bfc9ad5601271ffcd5ff21426b57647866c548791e1c39c8263eb7"></a>

## uid property — virtual_server.http.tcp_server_profile / 31c44aa6f101 / 8

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

<a id="canonical-4cc52935ff9a98cac15c2e7bd14d7a286878e52d66cf7326aaad0b71f175cd4d"></a>

## Next pages — virtual_server.http.tcp_server_profile / 31c44aa6f101 / 9

- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-a39a8f21a531c381b5984d77711593e653ecfec0ad9441ca3459fded11f1be4a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0fa401d515c61be4b241f4f39e47210624dc6e8d21711653dc1e3eba7de56bc5"></a>

## virtual_server.http.websocket_client_profile — virtual_server.http.websocket_client_profile / c308c06284b1 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- virtual_server.http.websocket_client_profile

<a id="canonical-99a75be6986d1779943a76b5db2f7896887f33654e5bd138dd25fa01d7bce3d5"></a>

Type: `"list"`. Computed.

WebSocket Profile Client. Web-related configuration

Upstream description:

Web-related configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-9fe982e157e2919a0eef8f6b6a21a4e63a55b0e3ab86e8d7f33c11d47fa44eef"></a>

## Direct properties — virtual_server.http.websocket_client_profile / c308c06284b1 / 3

<a id="canonical-2aa93501420a94f1e282f2d374bb2ad7cbe15a76a659579d6bfa0d3f57934992"></a>

<a id="canonical-226bc04782d5e945d3671207f0bffbba7413bf931a210e3d2bcea2ba93193e4d"></a>

## kind property — virtual_server.http.websocket_client_profile / c308c06284b1 / 4

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

<a id="canonical-751a1dbae2b1543e08985ea76f49a1edf3ce431d052f078e551f2aaa65e87bee"></a>

<a id="canonical-1675e871905ed030a4003b8929e36c91c8b6a8b2e246aa3dadc5f625752dbfcd"></a>

## name property — virtual_server.http.websocket_client_profile / c308c06284b1 / 5

Type: `"string"`. Computed.

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

<a id="canonical-e24476187f8804bfb7ec1ce1699e13bff0a0fb9e3596f8570d93294b9f4302c3"></a>

<a id="canonical-d4167285f9e9d19ed62c8e1c1a3822637c934061bf788dc6ce18eaacb4520376"></a>

## namespace property — virtual_server.http.websocket_client_profile / c308c06284b1 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-b7c0714faf9100daae542ca41330b27e96951e29f0bfc16313e0b722046a1d38"></a>

<a id="canonical-edf10d4153e762a6aec0af54206c822f6fe975d4912b727f9731b083711d08c1"></a>

## tenant property — virtual_server.http.websocket_client_profile / c308c06284b1 / 7

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

<a id="canonical-41ab680b7ff1cb85c8c1feb78350640a67d0c7a23bcf1e44eb2d32a7b2610b35"></a>

<a id="canonical-6ec73eedbd8136a9b8d22b0ce7b00c7ad0b422614ff6cdd2ca4f42ea88e6aec3"></a>

## uid property — virtual_server.http.websocket_client_profile / c308c06284b1 / 8

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

<a id="canonical-8540aa84d2c1b089f79788506ec94bb77a744b010e3ab6deba7dabd66db79aac"></a>

## Next pages — virtual_server.http.websocket_client_profile / c308c06284b1 / 9

- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-001df5346292404bb45ad860e9932722005765a6f0ec139fa2c476efc7409272"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a8cf0813e8dd2bd732b479ea4879d6b75929053e6b0a81cd6d90e8125bca5da"></a>

## virtual_server.http.websocket_server_profile — virtual_server.http.websocket_server_profile / d6c5a1f9be8d / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- virtual_server.http.websocket_server_profile

<a id="canonical-a1eb4f060c8c0a690908021d81dc0f9a664894d562ec56a50a8e663a8970b9f6"></a>

Type: `"list"`. Computed.

WebSocket Profile Server. Web-related configuration

Upstream description:

Web-related configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1e2429efe7fd66576878b0014029ac8eebb0c71b51478a7e90361dc553521fa3"></a>

## Direct properties — virtual_server.http.websocket_server_profile / d6c5a1f9be8d / 3

<a id="canonical-97cf291a7d43a9b8c2dd6f91153df5220779bc46a5b3582509b89ec8c720b2bb"></a>

<a id="canonical-ad28f22d327b34c10fbf4bd69f4287e3483332dba82ea85045837e12e88bae93"></a>

## kind property — virtual_server.http.websocket_server_profile / d6c5a1f9be8d / 4

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

<a id="canonical-68f2a25d912a4387ae5ba30943dc5ee579feb41e2053f72bb1f075e0cb7eadce"></a>

<a id="canonical-8345333e8552594348d761366f239b58e8e3207574547774f68d73960214d791"></a>

## name property — virtual_server.http.websocket_server_profile / d6c5a1f9be8d / 5

Type: `"string"`. Computed.

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

<a id="canonical-4717e1c31e04546992e0c711608598b7e26928abbb5d94b09920cb8a3de99874"></a>

<a id="canonical-1e32f295f7459afd393258238e8a996fc0cae07119bef06628568d22c1ce8a6e"></a>

## namespace property — virtual_server.http.websocket_server_profile / d6c5a1f9be8d / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-5a068417c7abf27dde51081c3f247e510be672746a7775996998f3b24509fc25"></a>

<a id="canonical-5827f88809da7d8bb6d2357213ce8e71a49281777a129f4a5aedd8ee4f00e0e2"></a>

## tenant property — virtual_server.http.websocket_server_profile / d6c5a1f9be8d / 7

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

<a id="canonical-c7f6f90c55f70d2a9bb1200a7daee4fc6b5e7e8a4946a14a801864989e9f7297"></a>

<a id="canonical-c6e2cb9511208e13e67127a2b9f377471459e8acd07128b00b5ce006a4e14323"></a>

## uid property — virtual_server.http.websocket_server_profile / d6c5a1f9be8d / 8

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

<a id="canonical-3b076fec8423752d069bc22dbfd74a2459a420282ace06709740937c2464965b"></a>

## Next pages — virtual_server.http.websocket_server_profile / d6c5a1f9be8d / 9

- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-128cbeb7bf9769a705b460cdf807fb5fb3658c967f490e3b7a9ff3f2090407cd)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-615bb385fac6ad6f77d0823453d6a15946aee0742f22f3647be7043317352d9a"></a>

## virtual_server.http3 — virtual_server.http3 / 8c55a1d04fde / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.http3

<a id="canonical-59c74f35851384576f40767c66a9253e2fd71cb1d116e74ea1357b5355bd03d5"></a>

Type: `"single"`. Computed.

HTTP/3 profiles.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-e11487a9ebaf9a83f191753b19939e1652a1f766997d965d1512a82bb47218bd"></a>

## Direct properties — virtual_server.http3 / 8c55a1d04fde / 3

- [client_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-eb70ca317fb2d06c7fc5b34b53cc7fcddecd42f1f916b0998150f1417f6c645d): complete subsection reference.

- [http3_profile](data-sources--application_profiles--reference--group-002.md#canonical-f9694d905f5ed59e46f2049dc25346398c794f2324791fccd4f51ff3df9e4183): complete subsection reference.

- [http_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-b66e3a646fcbf7b72caf8fe05f1af181b92b54e01269c81f1a5ec6daeca5090a): complete subsection reference.

- [http_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-9e926ffa120880d737800577d5f7f1f8bc50b5b5c10eb79ea2f3bf55e305ea62): complete subsection reference.

- [quic_profile](data-sources--application_profiles--reference--group-002.md#canonical-101ee2b0155e045a0d0ec8ea5dd79cb2b9036b8fa2d33b50393bc846a5e4d855): complete subsection reference.

- [server_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-0f83f095fe81404d2299e99e87f995eb0d78e0b4db14a2399ad53966dc83bc9a): complete subsection reference.

- [tcp_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-27fb2e871e23647c6a7458543e07b049a96cd9235240733cd75e515a5805f0e4): complete subsection reference.

- [udp_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-616d98759d583ff5c5f6e7211368eb3587315a8c749d5fc0f8694299d92b9b84): complete subsection reference.

- [udp_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-e5bb4853b99a44516a584def2cccbfd2307e2a58f5c8a692404ea1c68d922f20): complete subsection reference.

<a id="canonical-194547bb3c5aad31de973b615fc7d3443b8d66120f20d11b78206a45a58d1b94"></a>

## Next pages — virtual_server.http3 / 8c55a1d04fde / 4

- [virtual_server.http3.client_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-eb70ca317fb2d06c7fc5b34b53cc7fcddecd42f1f916b0998150f1417f6c645d)
- [virtual_server.http3.http3_profile](data-sources--application_profiles--reference--group-002.md#canonical-f9694d905f5ed59e46f2049dc25346398c794f2324791fccd4f51ff3df9e4183)
- [virtual_server.http3.http_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-b66e3a646fcbf7b72caf8fe05f1af181b92b54e01269c81f1a5ec6daeca5090a)
- [virtual_server.http3.http_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-9e926ffa120880d737800577d5f7f1f8bc50b5b5c10eb79ea2f3bf55e305ea62)
- [virtual_server.http3.quic_profile](data-sources--application_profiles--reference--group-002.md#canonical-101ee2b0155e045a0d0ec8ea5dd79cb2b9036b8fa2d33b50393bc846a5e4d855)
- [virtual_server.http3.server_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-0f83f095fe81404d2299e99e87f995eb0d78e0b4db14a2399ad53966dc83bc9a)
- [virtual_server.http3.tcp_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-27fb2e871e23647c6a7458543e07b049a96cd9235240733cd75e515a5805f0e4)
- [virtual_server.http3.udp_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-616d98759d583ff5c5f6e7211368eb3587315a8c749d5fc0f8694299d92b9b84)
- [virtual_server.http3.udp_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-e5bb4853b99a44516a584def2cccbfd2307e2a58f5c8a692404ea1c68d922f20)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-eb70ca317fb2d06c7fc5b34b53cc7fcddecd42f1f916b0998150f1417f6c645d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9ab81a58b99d93150166078d6d9b555b3fb2fb678977a003f70013fb51c3cf8"></a>

## virtual_server.http3.client_ssl_profile — virtual_server.http3.client_ssl_profile / e3cebbe6c267 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367)
- virtual_server.http3.client_ssl_profile

<a id="canonical-2da862f0c0a27f3961b846ec9ac0c9b0fbb93d1c88aa2d038e607b54538e496b"></a>

Type: `"list"`. Computed.

Client SSL Profile. Client-side configuration

Upstream description:

Client-side configuration

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

<a id="canonical-64a7b4f890560495c6a722947c18b3c00a19c98317147878bfa7be31612ee8f2"></a>

## Direct properties — virtual_server.http3.client_ssl_profile / e3cebbe6c267 / 3

<a id="canonical-cb93c4023178aa786455638fbb10943269568649871b417a1e9bf321cbf74090"></a>

<a id="canonical-5f433f1bd965eec01b3e184c9bb263db72a10cdfb8f2f5a93497bc2c868cf779"></a>

## kind property — virtual_server.http3.client_ssl_profile / e3cebbe6c267 / 4

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

<a id="canonical-d215931e8a5ebd7489d85a5b7948e1ee017e66d2ed6051656301f14ae37b1e95"></a>

<a id="canonical-aeaee9395d0612d8344e70c9a6830cf74630b3727d147ed176f9a5298cdee917"></a>

## name property — virtual_server.http3.client_ssl_profile / e3cebbe6c267 / 5

Type: `"string"`. Computed.

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

<a id="canonical-a04f0ce00e8c93d3682dd00c1771fb94ac3649d4ea11d22484a4dd07d4887f66"></a>

<a id="canonical-f38ad33010c859dfebbdc4779fe3b3f3861288399d655f41aa3b2200faa66168"></a>

## namespace property — virtual_server.http3.client_ssl_profile / e3cebbe6c267 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-c792d1e98e463d8b88faf7fe4dc73dad20af8882e640d42228beeb23527acbbf"></a>

<a id="canonical-bf5bfa4f1e8e71ff510ecc3afe0f6c0cfcb2862a5f5846501c549f203a79c598"></a>

## tenant property — virtual_server.http3.client_ssl_profile / e3cebbe6c267 / 7

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

<a id="canonical-2504c99b5fff5fd0e127140bc408f47955d69197007bbd3b6b99fbd7f7877d11"></a>

<a id="canonical-2db3f64ccec5eb0d0b19ad18276dbf9251d5da78037e194c2fe3708638c1815f"></a>

## uid property — virtual_server.http3.client_ssl_profile / e3cebbe6c267 / 8

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

<a id="canonical-d878172b768744241d3d9714e3437ed1b811106b57d7283b6eddb5a862e387df"></a>

## Next pages — virtual_server.http3.client_ssl_profile / e3cebbe6c267 / 9

- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-f9694d905f5ed59e46f2049dc25346398c794f2324791fccd4f51ff3df9e4183"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-301896b519dd5e6d840c3ed8b455c072772e2cb7b32e5208873d54c967d1b77c"></a>

## virtual_server.http3.http3_profile — virtual_server.http3.http3_profile / 57de229ee70e / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367)
- virtual_server.http3.http3_profile

<a id="canonical-e85b342deb4612b3bfb19a47d01b564d9ba0f221961ff9f5a8f6a55e6d23a1be"></a>

Type: `"list"`. Computed.

Configuration parameter for http3 profile.

Upstream description:

Configuration parameter for http3 profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-546bd2d3c39731519c87a8184b7f4284459d244095c3fbc2d12d3f753417c383"></a>

## Direct properties — virtual_server.http3.http3_profile / 57de229ee70e / 3

<a id="canonical-a697d4db6a32b7ade68b1ef43f8833817c56cef1150f7787ba78f8236ce43693"></a>

<a id="canonical-3b71ab0e1a0ff4562a04b9bffc3d084b1dfccf6c0401900b0d7da9d59b7049a4"></a>

## kind property — virtual_server.http3.http3_profile / 57de229ee70e / 4

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

<a id="canonical-78f4c4e2ea0cec354bafe02699029e318b9f190dd6cbf233ea7092d70375a18a"></a>

<a id="canonical-618162f3933ba7faec461c25e1d05659356c9aa56c5933c413ea6fc596cade5b"></a>

## name property — virtual_server.http3.http3_profile / 57de229ee70e / 5

Type: `"string"`. Computed.

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

<a id="canonical-68374b757b28c671cd02be73291e051da9b891a7adc52b474329adb432800d87"></a>

<a id="canonical-e108ac6c5cf6b792861e4bd19cd91ed18490812cd4aed3b5d4f085f00bba5d49"></a>

## namespace property — virtual_server.http3.http3_profile / 57de229ee70e / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-00d6ddfdc23b96a84db1a7fd6f561749cf5a08bc8ccee037f93d599dfae00421"></a>

<a id="canonical-d01da8b7b5400734af444fa17fdba9db09982bc8ca9367d52fe504b13d9842e8"></a>

## tenant property — virtual_server.http3.http3_profile / 57de229ee70e / 7

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

<a id="canonical-48123a4b3a4706efb3a7f1aeb8bc539dd169e6ccdb967e97157245ce940b5679"></a>

<a id="canonical-edd26fd97cb026fdd4d296e4de0a68c5ebfd1f049bb044d7a01542360a7ec2a7"></a>

## uid property — virtual_server.http3.http3_profile / 57de229ee70e / 8

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

<a id="canonical-596c08473aea3b19e3056f0622cb64767c2dc5ccf66b29eaf5c03f5ea92c521f"></a>

## Next pages — virtual_server.http3.http3_profile / 57de229ee70e / 9

- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-b66e3a646fcbf7b72caf8fe05f1af181b92b54e01269c81f1a5ec6daeca5090a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a549014571afa93fcdc0d000f676908da4952ff428cc0de551357cd67046c10"></a>

## virtual_server.http3.http_client_profile — virtual_server.http3.http_client_profile / 423f5fec1733 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367)
- virtual_server.http3.http_client_profile

<a id="canonical-15ccf63ee6aac4c25cabb0155afb24aa61ccdf950196d80611c37ed485e026b5"></a>

Type: `"list"`. Computed.

HTTP Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ef3d86524779e42b9c594e33d95a651766af4c925c1ba363660b6245bc7f0273"></a>

## Direct properties — virtual_server.http3.http_client_profile / 423f5fec1733 / 3

<a id="canonical-a08a549efcb756630d46caf0562c8ae3d8ba79e0caf443dc1de73355a4cd98c0"></a>

<a id="canonical-ae48f83385be6d73e448450dd31a411e8deb35739c28ac790397735df6e164b7"></a>

## kind property — virtual_server.http3.http_client_profile / 423f5fec1733 / 4

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

<a id="canonical-69c49bdd9ba32b00437f5fe62c1ed0328e28b9e5b28597932e593ccc60c72a33"></a>

<a id="canonical-34fa2f24cd175c37d3577c43f4519132532e34e5ea0b9e32da8ac4acf608f630"></a>

## name property — virtual_server.http3.http_client_profile / 423f5fec1733 / 5

Type: `"string"`. Computed.

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

<a id="canonical-789e73b40dacf8a18cd57afd26b01f1fcf0fea0403c01c980b24a037ef689236"></a>

<a id="canonical-72d6f2278c1910f84a8d738fafb292f0b7d829c146167e5af711dd6b12602a0d"></a>

## namespace property — virtual_server.http3.http_client_profile / 423f5fec1733 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-a65df17694572abe989e93ca8d8e503b4b5c333d3ac39f990902e5fec9b8ab42"></a>

<a id="canonical-13b6df0e9268d9f3a23db7fdb53b8db244208eb1d821d824fb522b2c9f7f6b6c"></a>

## tenant property — virtual_server.http3.http_client_profile / 423f5fec1733 / 7

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

<a id="canonical-6ecf1e806d3064c15d60818b0c417637accd448088a2dc3213c1fc9c656a8034"></a>

<a id="canonical-4178d43b2c32e595e96e791da367753a5967772314f9c9414e2d0df9d7ffeeb3"></a>

## uid property — virtual_server.http3.http_client_profile / 423f5fec1733 / 8

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

<a id="canonical-f1d3ce1a71f2d892c8b3e0218c917dd9551db0e750ddf3ebec262f62707d42fd"></a>

## Next pages — virtual_server.http3.http_client_profile / 423f5fec1733 / 9

- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-9e926ffa120880d737800577d5f7f1f8bc50b5b5c10eb79ea2f3bf55e305ea62"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7dd8584fbe936fdeb01399f93c5e60db84cbb4a94973836b77546a6219c258e"></a>

## virtual_server.http3.http_server_profile — virtual_server.http3.http_server_profile / 30fd09a7ca40 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367)
- virtual_server.http3.http_server_profile

<a id="canonical-e9da22cf52ac6df0e313c16c640a2e85b733bb71aa5bcae661ce5d65db9d7424"></a>

Type: `"list"`. Computed.

Configuration parameter for http server profile.

Upstream description:

Configuration parameter for http server profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-77e4fe6f1c4621e3458306d08eea27842309f6db68cbec4677ebacf6fbf3a45f"></a>

## Direct properties — virtual_server.http3.http_server_profile / 30fd09a7ca40 / 3

<a id="canonical-111d9a5e6284823ea2f30caff78e087a42d3c32a46d7a6609eb9f1574fc1bf7e"></a>

<a id="canonical-93e0fa09761fd4e0f835f7a22a79b89d17ecffa689d3e20e4ffc1dae096378ac"></a>

## kind property — virtual_server.http3.http_server_profile / 30fd09a7ca40 / 4

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

<a id="canonical-07df6918dcdff9ff7f9c956fe7ed3662714af25bf37d8d3f8edad2c87b575e37"></a>

<a id="canonical-9bf592b9a25a914f6da7aef90ffc4d60edb820f7644ba881916152209b2e0e36"></a>

## name property — virtual_server.http3.http_server_profile / 30fd09a7ca40 / 5

Type: `"string"`. Computed.

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

<a id="canonical-a7ad5bdb5bce221058651701445caf32b44bf4519bac3ddec6355e07e8c128c1"></a>

<a id="canonical-83275418ce6e6cf4398e473b28a34933c35775c908daa83d1882cb2ea2fc78db"></a>

## namespace property — virtual_server.http3.http_server_profile / 30fd09a7ca40 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-8c943699acbd8b65173840a9c950db83106493f52c64898ade6d8e2039ab79da"></a>

<a id="canonical-80abc561b02a8545d8ed5a1f46ed9ebcec0d74d53b142e23dd458fff514f0e51"></a>

## tenant property — virtual_server.http3.http_server_profile / 30fd09a7ca40 / 7

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

<a id="canonical-5f885262e3bf5b7029b87b60187a1b4ab2bc337300bf03b3a9302fe711e8efe3"></a>

<a id="canonical-74b07a6b747a986a79c7f8ac094954de88619db14ab26cff1861244b290dc6d8"></a>

## uid property — virtual_server.http3.http_server_profile / 30fd09a7ca40 / 8

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

<a id="canonical-3e166e1fdd2dd82efc6bf736d2bd351ec6984df55f6ec471600da79fe2eb0275"></a>

## Next pages — virtual_server.http3.http_server_profile / 30fd09a7ca40 / 9

- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-101ee2b0155e045a0d0ec8ea5dd79cb2b9036b8fa2d33b50393bc846a5e4d855"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ff99912b73e376cb5e9b3c13edc5b0902a64840f6c510de63702a30f5950a2d"></a>

## virtual_server.http3.quic_profile — virtual_server.http3.quic_profile / 00de46deef65 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367)
- virtual_server.http3.quic_profile

<a id="canonical-b50e8a2eec57ed79fa1950be4865f9706f3a7d22f5f10988d017df3388e27137"></a>

Type: `"list"`. Computed.

Configuration parameter for quic profile.

Upstream description:

Configuration parameter for quic profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2016c672e56b92cdcfdc85414eb9b837327366da6bb7feb9bdcf4564cd1d3cb1"></a>

## Direct properties — virtual_server.http3.quic_profile / 00de46deef65 / 3

<a id="canonical-63f41adb59f2bdd84f6923bdfc9f41c9563bd19d5531f0833c9d4804ed5b80d5"></a>

<a id="canonical-ebf39254d356838e1726dd1b00e447c9fe4891007baa2372b200e9020daf2f54"></a>

## kind property — virtual_server.http3.quic_profile / 00de46deef65 / 4

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

<a id="canonical-15262949d436631a03dc62127a28342377c1d92b53f92355c06deb0b4d4189b9"></a>

<a id="canonical-293c8f10ae093bdd408036ebbae0be50d79fefc784a809e15ac13add945d20ee"></a>

## name property — virtual_server.http3.quic_profile / 00de46deef65 / 5

Type: `"string"`. Computed.

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

<a id="canonical-b857560b8b16d3bc873f12c788047b8f9dc94f76e8d9c1cd1c1520a863a82495"></a>

<a id="canonical-d87910f99424db6b9a15155779dc7a30a26186efc0cd416c60ae59de233a8b69"></a>

## namespace property — virtual_server.http3.quic_profile / 00de46deef65 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-9e9fb74b79a10d0cf7ec602952ba9f8f6e16bed43454c39019acb75fe4057f41"></a>

<a id="canonical-0daa15549012d76b817a7a34605d2ebd132fd0069c047dcc405acd5ac6bcb3a1"></a>

## tenant property — virtual_server.http3.quic_profile / 00de46deef65 / 7

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

<a id="canonical-8b828f3973545bea6c0fbce252fc3f8f8a9a7f7cb27f83a8be824734836f0ae4"></a>

<a id="canonical-9c5265e1d3b6809a1d3f28692e5c9f907be387753c5e5793e79c48dce1fca53e"></a>

## uid property — virtual_server.http3.quic_profile / 00de46deef65 / 8

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

<a id="canonical-d14fa649216fac293875ecd28d0c86c1bfdb5e4ba9e3fd7a4dce7306714c752d"></a>

## Next pages — virtual_server.http3.quic_profile / 00de46deef65 / 9

- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-0f83f095fe81404d2299e99e87f995eb0d78e0b4db14a2399ad53966dc83bc9a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-213c63ad9c03c5fb3fb54b191f0a6570ea0083b6f46613c33e7e9eae94585a30"></a>

## virtual_server.http3.server_ssl_profile — virtual_server.http3.server_ssl_profile / f0d2ecad9290 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367)
- virtual_server.http3.server_ssl_profile

<a id="canonical-a2e7936f9d7cb4e74f14f2a394bec22d0386db6d5be7d6e0ce7c59a5cff08a43"></a>

Type: `"list"`. Computed.

Configuration parameter for server ssl profile.

Upstream description:

Configuration parameter for server ssl profile

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

<a id="canonical-082a5d2e238f108053ee020cd932dd835b29df6a59a0669700e14c56c970eb45"></a>

## Direct properties — virtual_server.http3.server_ssl_profile / f0d2ecad9290 / 3

<a id="canonical-06cb7d1d5383b35318ee40487179b9f4a4f3d4605c3166198dbfb0b55917ed76"></a>

<a id="canonical-7596da7e2bc8008e429bdd46be48aa9b378c81a1532815fe28ab3290d508e955"></a>

## kind property — virtual_server.http3.server_ssl_profile / f0d2ecad9290 / 4

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

<a id="canonical-4e39c3ce0109b220b8f3be8591545016263585f55974626e228b7171a2c3b1d1"></a>

<a id="canonical-571453b675ae6fb3734e1a0c5e2fcb0287afba9f04cdd1beafac5b57059399f5"></a>

## name property — virtual_server.http3.server_ssl_profile / f0d2ecad9290 / 5

Type: `"string"`. Computed.

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

<a id="canonical-a8dc133b17618d8636ce04534f0405a04ad801ff2e469ffcd477fd8a0e981410"></a>

<a id="canonical-f10674c98591ebe2b6ce79f3a0bdd11f918cd1308dbc66d55e25287cf5951bab"></a>

## namespace property — virtual_server.http3.server_ssl_profile / f0d2ecad9290 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-19fd4072c5abb1c0b86db8daaeb388446f8019f3a4fa1dcf6f84176ccdf935cf"></a>

<a id="canonical-73d26f08639cecad1e740d5967bb0e7c0355fb71938394dd22cb1afc7d3c81e0"></a>

## tenant property — virtual_server.http3.server_ssl_profile / f0d2ecad9290 / 7

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

<a id="canonical-f99acca0cf79c1762b86dfc8a536d862a6884db0d93a45c62e4828877122e302"></a>

<a id="canonical-441d8f3254b0944fee9ff1c1a18db8110e9b2af6ff18a7883b3b8e544088389e"></a>

## uid property — virtual_server.http3.server_ssl_profile / f0d2ecad9290 / 8

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

<a id="canonical-313a45d2921860d9fc5dea3120e5311474a66c4c8482f7e2eb7ac72fc152ebbe"></a>

## Next pages — virtual_server.http3.server_ssl_profile / f0d2ecad9290 / 9

- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-27fb2e871e23647c6a7458543e07b049a96cd9235240733cd75e515a5805f0e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b20c24c6f85e6fff92315e9732b9e3e131074eed0e2f9cc854e5769ba1893f79"></a>

## virtual_server.http3.tcp_server_profile — virtual_server.http3.tcp_server_profile / 32e3403a172b / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367)
- virtual_server.http3.tcp_server_profile

<a id="canonical-959687a2571b8aa04770a89bcf8a57170286fecdf391e1e4b98ad4642f55a86d"></a>

Type: `"list"`. Computed.

Configuration parameter for tcp server profile.

Upstream description:

Configuration parameter for tcp server profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-11159148155f00c2b2bc5cdfd715d6059b93c9b38bbe19151a741f5ddf60f834"></a>

## Direct properties — virtual_server.http3.tcp_server_profile / 32e3403a172b / 3

<a id="canonical-bf7a34d8f639b21838edd1e5bb266a3ae8ae2bdf6485c2a7c4a3462fdedfb8c1"></a>

<a id="canonical-1dc9ca0f315839c3b326e9c7cadc4c77689af9f386830af5e82b20b8763037c9"></a>

## kind property — virtual_server.http3.tcp_server_profile / 32e3403a172b / 4

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

<a id="canonical-b9a75ec583315b8eef735c7224696b3fdfe2762e8c7b2e0d6b19829255c1515b"></a>

<a id="canonical-702917a418e6d291b3e8617c329bc410cf78cee834fb1138366592dc7cbc7447"></a>

## name property — virtual_server.http3.tcp_server_profile / 32e3403a172b / 5

Type: `"string"`. Computed.

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

<a id="canonical-965b89684d5da78027a927448455e8c737d0bffe4f15f0acc3a380f432da3aaa"></a>

<a id="canonical-ad459e8003b53e36a7ab17e4812d331aca63b6c4f9a5c9951e1e9c5a3a10bded"></a>

## namespace property — virtual_server.http3.tcp_server_profile / 32e3403a172b / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-b9a6a5b9e027e315db57fb96330cda7242fa2eeff9baef6cf5a59d99b1b7000b"></a>

<a id="canonical-11779d1ceea156ada6a8cbc08af035614ad327a7269d87371280d61f7ab80817"></a>

## tenant property — virtual_server.http3.tcp_server_profile / 32e3403a172b / 7

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

<a id="canonical-fd14f0436cf609ac4a2d163b734822591deff2fe37ac90b11bce2e681315563a"></a>

<a id="canonical-3cde9309b46bbb28ea5550f93135ad8446f6a3eeff344e4e01f7ba02621634d1"></a>

## uid property — virtual_server.http3.tcp_server_profile / 32e3403a172b / 8

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

<a id="canonical-958dee0af31360f5195667da75770441d022b2a9811d67ecf84c72eaede72207"></a>

## Next pages — virtual_server.http3.tcp_server_profile / 32e3403a172b / 9

- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-616d98759d583ff5c5f6e7211368eb3587315a8c749d5fc0f8694299d92b9b84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c76d59e8fb842d8f16d48f89f0c1ebeed340841dac3990e5864f695bbb964cc"></a>

## virtual_server.http3.udp_client_profile — virtual_server.http3.udp_client_profile / 75a5526819c1 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367)
- virtual_server.http3.udp_client_profile

<a id="canonical-d455d2e53cb96825d6cd79d9195cf5edd883d522a5a74ae99a2cf7fbd8635269"></a>

Type: `"list"`. Computed.

Protocol Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-6cf3ad1e75fc56b964d8a2f01e2b570d25f47502d88ad19931d48bd580060ec2"></a>

## Direct properties — virtual_server.http3.udp_client_profile / 75a5526819c1 / 3

<a id="canonical-a381ad507f416241c6412706080ecc5b2ed2cd13fc012e191ab9eeb4c3672709"></a>

<a id="canonical-e499ecf7a94ca67964b5f993674e7eece9966356a63b82a340d0b45afc96bbf6"></a>

## kind property — virtual_server.http3.udp_client_profile / 75a5526819c1 / 4

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

<a id="canonical-970aeb4e827cc02d63bf1975ff1f36d488566c0cf3b2b146b75061f9a15d42e0"></a>

<a id="canonical-9b3bb9e4b454f46d1579a58eb6bb3cf8cf131be47f5274a2ed6483f6c47e1f4c"></a>

## name property — virtual_server.http3.udp_client_profile / 75a5526819c1 / 5

Type: `"string"`. Computed.

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

<a id="canonical-a34f28c9b3fafef8b345bbd7c0312203c041075eca1187d6c03d817eed9427b1"></a>

<a id="canonical-0fd47396127b198c06066a13db8dbe6a4668524e2efa42c1ed7dd541a5175d2f"></a>

## namespace property — virtual_server.http3.udp_client_profile / 75a5526819c1 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-20c4e627d01cb49e1e9884f72ae8fff26f154bd64cd4342954b10a141d5262ed"></a>

<a id="canonical-4fe7b33f450748d70837d075f773f983242cf63fecc36b08bc0581d9837c87ef"></a>

## tenant property — virtual_server.http3.udp_client_profile / 75a5526819c1 / 7

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

<a id="canonical-df276a148c4bed5388121a127fc955c15fd1d25798c122f4801500b76586eb8d"></a>

<a id="canonical-4e837a5da8c56f1b9758468b5931ba22d0db7051f71d2ff584ff7c07a42c23bd"></a>

## uid property — virtual_server.http3.udp_client_profile / 75a5526819c1 / 8

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

<a id="canonical-7857f2608b663a0f3eb9ef3223fd1bd5d444f3ca3c9adf31d68f1cc2eb16aac6"></a>

## Next pages — virtual_server.http3.udp_client_profile / 75a5526819c1 / 9

- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-e5bb4853b99a44516a584def2cccbfd2307e2a58f5c8a692404ea1c68d922f20"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2cc75ac5dc7c1656e24a5c30b04a8f689fd0318b732aa6897251765548b5593d"></a>

## virtual_server.http3.udp_server_profile — virtual_server.http3.udp_server_profile / c93704e96ada / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367)
- virtual_server.http3.udp_server_profile

<a id="canonical-03e78a67a865975e00df35c6fc6ecc36f4d7901daae8cb33f5ae7f41a6de3463"></a>

Type: `"list"`. Computed.

Configuration parameter for udp server profile.

Upstream description:

Configuration parameter for udp server profile

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-14d5ca57133783c01a14a69795ff0cdf43fa7dae4c6627a5723f9dcb40780bda"></a>

## Direct properties — virtual_server.http3.udp_server_profile / c93704e96ada / 3

<a id="canonical-1e2bdafdc7f35c1e1f0b980a7b5f10930f597ffdb01291676a8a69d4385b1123"></a>

<a id="canonical-372011e61a803ecd9807f6cf55f78db53f12a3a8b99a7a730785530e9c2393e0"></a>

## kind property — virtual_server.http3.udp_server_profile / c93704e96ada / 4

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

<a id="canonical-99ad14f06ef68e0b7498c1f579016b8fe0232f645e30b9b76b3f8bf0a90d22e9"></a>

<a id="canonical-50cfb5a783061a5e6318d0b2ef4ea4bc6dd02b1e530fdd3beb5d97753442caa3"></a>

## name property — virtual_server.http3.udp_server_profile / c93704e96ada / 5

Type: `"string"`. Computed.

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

<a id="canonical-d71558dd0013a6608cc2386a6916ea10d565cd841aa9e65c671f459f090a52a2"></a>

<a id="canonical-9eb5804e1ddafc3b67bcfdcd52219524adb0f45a3f12853588620de22983e609"></a>

## namespace property — virtual_server.http3.udp_server_profile / c93704e96ada / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-978514a2159618d6be71d6d98e55fd16288786fc07a2c7418cd80abf2453bf1d"></a>

<a id="canonical-f510a49ecf446a12bdce544f8682a951f35f3435b5eec3b2a48b9ce2bffbbed9"></a>

## tenant property — virtual_server.http3.udp_server_profile / c93704e96ada / 7

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

<a id="canonical-6da45e49e79b962745dbefffe9421df37824154fe20f8190d6d328d36d2b4b7b"></a>

<a id="canonical-7682e39a523c4bd7244bb1d7a14673f7cdec8a963066d35b9d037003fc7357a3"></a>

## uid property — virtual_server.http3.udp_server_profile / c93704e96ada / 8

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

<a id="canonical-839964bc6086055efbefd955348baeef4254ba555cf219ae7bb5b389475475d5"></a>

## Next pages — virtual_server.http3.udp_server_profile / c93704e96ada / 9

- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-af9f2c4d452502e80b0646ae3e551a0d366427163c82c18e370f5d573a180367)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edbaa97b487a3bfc63dd891076ed50c2490fa112472fc3e5019450f5e7615358"></a>

## virtual_server.https — virtual_server.https / c14a30283bb5 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- virtual_server.https

<a id="canonical-d4590d5492772762dcd475e0f257fdb82dd3a054627a0738dc4f0a13978965db"></a>

Type: `"single"`. Computed.

HTTP profiles.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-af1179c0d8b86b116e999e161d180dc03a77792623a0f92cc3eb132b6abf4ee6"></a>

## Direct properties — virtual_server.https / c14a30283bb5 / 3

- [client_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-c10f2594f33b6eab9f734d6d1dec627e245a1fced6273176409e375f2e1e24a3): complete subsection reference.

- [http2_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-2e971dbdcdeaad04a7e67e910ccc1b35bf7237dc2dfd64442ca19d9b30922476): complete subsection reference.

- [http2_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-1f4d56a60c7aa9907aab9fa958c49f1db3289aa085bc89f77d289cc78512d70c): complete subsection reference.

- [http_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-4f2295adb60b5954627ad011694f3a310228b53f7068dabfe59974cabeba4f40): complete subsection reference.

- [http_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-5ab92633e91ab2b2e5e181271c742ebb366df4afb6db1748884f84225fa1f911): complete subsection reference.

- [ocsp_profile](data-sources--application_profiles--reference--group-003.md#canonical-fa2f89657d2a7c991f8eaf934c50ff599ed05d06e1368f775301b07573ad1824): complete subsection reference.

- [server_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-6e7660e2ef830fb2c83dd4360748cdcbef6b9b4a757cc62fd62105258fdb1cd7): complete subsection reference.

- [stream_profile](data-sources--application_profiles--reference--group-003.md#canonical-240a2b23fed325227dd634c21c987962bb8f31defa13b73f33971e6832a5aa13): complete subsection reference.

- [tcp_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-a2de95bc4627e2378ebccf89c853098db318d0565882d5ccc198ffefed9bbaf2): complete subsection reference.

- [tcp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-bb93372f0965c7716127d587edd71646800d126364cc343c2c10f4c0a4476d63): complete subsection reference.

- [websocket_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-93eedc7c0076712eee02e8427f787d05503b61bd07efcd3c772c90a582b01c04): complete subsection reference.

- [websocket_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-c6f803d16e3092136fe7867446111a7a15a6010c339dfa4ef2439b4ccc4779bd): complete subsection reference.

<a id="canonical-eb8e97291607c8a95ffe1fd3a9d5de2126790cd96618cf424a24c1696a10712f"></a>

## Next pages — virtual_server.https / c14a30283bb5 / 4

- [virtual_server.https.client_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-c10f2594f33b6eab9f734d6d1dec627e245a1fced6273176409e375f2e1e24a3)
- [virtual_server.https.http2_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-2e971dbdcdeaad04a7e67e910ccc1b35bf7237dc2dfd64442ca19d9b30922476)
- [virtual_server.https.http2_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-1f4d56a60c7aa9907aab9fa958c49f1db3289aa085bc89f77d289cc78512d70c)
- [virtual_server.https.http_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-4f2295adb60b5954627ad011694f3a310228b53f7068dabfe59974cabeba4f40)
- [virtual_server.https.http_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-5ab92633e91ab2b2e5e181271c742ebb366df4afb6db1748884f84225fa1f911)
- [virtual_server.https.ocsp_profile](data-sources--application_profiles--reference--group-003.md#canonical-fa2f89657d2a7c991f8eaf934c50ff599ed05d06e1368f775301b07573ad1824)
- [virtual_server.https.server_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-6e7660e2ef830fb2c83dd4360748cdcbef6b9b4a757cc62fd62105258fdb1cd7)
- [virtual_server.https.stream_profile](data-sources--application_profiles--reference--group-003.md#canonical-240a2b23fed325227dd634c21c987962bb8f31defa13b73f33971e6832a5aa13)
- [virtual_server.https.tcp_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-a2de95bc4627e2378ebccf89c853098db318d0565882d5ccc198ffefed9bbaf2)
- [virtual_server.https.tcp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-bb93372f0965c7716127d587edd71646800d126364cc343c2c10f4c0a4476d63)
- [virtual_server.https.websocket_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-93eedc7c0076712eee02e8427f787d05503b61bd07efcd3c772c90a582b01c04)
- [virtual_server.https.websocket_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-c6f803d16e3092136fe7867446111a7a15a6010c339dfa4ef2439b4ccc4779bd)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-c10f2594f33b6eab9f734d6d1dec627e245a1fced6273176409e375f2e1e24a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af10b546b02c40b2e4b1db7434bb82848d4e52ed1f9da077803fe1d2ee60b3ae"></a>

## virtual_server.https.client_ssl_profile — virtual_server.https.client_ssl_profile / 0ccd39f468a3 / 2

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-d0885eb035e95f1b1c36adc903324929bfd8a8166bc6eb75d22521793fb630ed)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-8519cbc749aafaa144cf506541d90c30aaa671f48a5df22ebcbf657678299cdc)
- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- virtual_server.https.client_ssl_profile

<a id="canonical-5ee982bf36322fe9f914ead4f15768695c61a7aa42ac737432d767aeac0da4b7"></a>

Type: `"list"`. Computed.

Client SSL Profile. Client-side configuration

Upstream description:

Client-side configuration

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

<a id="canonical-63ba2ce4d2738a47b4472a6ce67cd6066c25304542d3159f6fa0bf5300bbab1a"></a>

## Direct properties — virtual_server.https.client_ssl_profile / 0ccd39f468a3 / 3

<a id="canonical-5467672c814dd823424b77ddf32537a01c3ad96b70a16cb2742b59a9439f5060"></a>

<a id="canonical-b3859df3c83903acee9a0cf3bda61570c83731095881e83a18ccd5f3db8655b0"></a>

## kind property — virtual_server.https.client_ssl_profile / 0ccd39f468a3 / 4

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

<a id="canonical-e73f7703e9c1231810de98bcb5fb85445290849ec06696388d26e138e4a95942"></a>

<a id="canonical-a7e28db3e418e14663b3d761817516ab6fa91a3331ac95b2b1b989448a8b8cf3"></a>

## name property — virtual_server.https.client_ssl_profile / 0ccd39f468a3 / 5

Type: `"string"`. Computed.

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

<a id="canonical-8fd42f31ad1273c70c2963dafe51b4a9176233c77a0cfb883514b7aed323c460"></a>

<a id="canonical-a7c252d594ab613380b0fbac82c3b95e14239731dd69453a406e2eafc6495706"></a>

## namespace property — virtual_server.https.client_ssl_profile / 0ccd39f468a3 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-16544674d2ac2b82631e500d2cd66b44c9f4456e05f8c72e109514337dbe1104"></a>

<a id="canonical-f1d93c4a5106b2da28a8bd044cfb5b2076020146f4238670d948b44462e08cab"></a>

## tenant property — virtual_server.https.client_ssl_profile / 0ccd39f468a3 / 7

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

<a id="canonical-d24e2c915b892193f109ff629814de0b7d403ceffa4a51c1ef0a63adc82182aa"></a>

<a id="canonical-a5a62d4a93c180580ca5ba660caf9fc7301350b3e14f313a75376ed785b1ea0b"></a>

## uid property — virtual_server.https.client_ssl_profile / 0ccd39f468a3 / 8

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

<a id="canonical-63a568c0a4fe6ac24b9e66c467f0297b7c538c92ff73666133587054df5893ff"></a>

## Next pages — virtual_server.https.client_ssl_profile / 0ccd39f468a3 / 9

- [virtual_server.https](data-sources--application_profiles--reference--group-002.md#canonical-3913d79055f5178e921761a3717a87f1a62d9a7075f7880dbebea5e6152a523f)
- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-759630bb83acabf06c265dca13c07d88ae392193439f6c718a93ef52b0ecd37c)

<a id="canonical-2e971dbdcdeaad04a7e67e910ccc1b35bf7237dc2dfd64442ca19d9b30922476"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
