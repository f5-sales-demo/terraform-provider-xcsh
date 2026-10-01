---
page_title: "xcsh_crl reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_crl reference."
---

# xcsh_crl reference

<a id="canonical-c2c016728d6cc79099c35ae246c02f69a282f25fbac91acb6b380a38ba11c747"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27460c0d799999699feb4ff83847c3960371f19fe25af9d6bf5c363a5b35009a"></a>

## Property reference — Property reference / 8df667a7129e / 2

Breadcrumbs:

- [xcsh_crl](../data-sources/crl.md#canonical-73f9576382ba605e0c52a90b24f3c47e46f2fa2fe6a252eff356e5dcb2c0e943)
- Property reference

<a id="canonical-e6342bc563655e92e41e51942ea29f960e50dc60b145cf5e6e179acc51a1b1fb"></a>

## Direct properties — Property reference / 8df667a7129e / 3

<a id="canonical-0ffc9469f0096cdf6175dde58d2297c921794618c944278d2a70b7a05474fa7d"></a>

<a id="canonical-35ebb34da6cc517b5e69ded227dbc9e4cc0b5b2aa3cba51e57d1ff13acfbd141"></a>

## annotations property — Property reference / 8df667a7129e / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

<a id="canonical-d3d9113b90eb247fd938650c3d049a0845657ca5eba27f76accb428f1b670754"></a>

<a id="canonical-b94a08a20ad386fb36ba9817dffd7fc243847d83fe8b5f47c3250f352cec358b"></a>

## description property — Property reference / 8df667a7129e / 5

Type: `"string"`. Computed.

Description of the CRL.

Upstream description:

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

- [http_access](data-sources--crl--reference--group-001.md#canonical-4aa87e7995a0400aaedc2a39d9c859c6240bfba132883b6d4fe11d5ce565c5e5): complete subsection reference.

<a id="canonical-5ee27e2daad2ac096d4ffe9ab02b1db2da5c63f5b0745c2b77ada1c78c4914b8"></a>

<a id="canonical-ed00f66480781c4132a7770e69a9ed61a7a4204ad062ea98c0a97622e8cbc9f5"></a>

## id property — Property reference / 8df667a7129e / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-d76f5795a90e77a3cf8ba993c1b4b5cff597a1c17f26d461dc5e0327ac16f9ae"></a>

<a id="canonical-2a94af31ebcbff1e389455b7b66857443b488b0f255f8e7b323ba7b2dd761454"></a>

## labels property — Property reference / 8df667a7129e / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-2c31db0e68e9c5cbc54a279b1c9a49d8db755df855a601a8be6737a017f506cc"></a>

<a id="canonical-4d691b05c60dffca9a699ec5e8c48b3c310cac855fbb15467b1e78d3af50e34c"></a>

## name property — Property reference / 8df667a7129e / 8

Type: `"string"`. Required.

Name of the CRL.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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

<a id="canonical-9186bd6ee7fa44a49325386aab786f5af6ec0cc32906de165b67e45d54927b12"></a>

<a id="canonical-bee8f7a56d5412b1327457c3004a11eaf089d9779f51e91c6a4eac74e9c6cb4d"></a>

## namespace property — Property reference / 8df667a7129e / 9

Type: `"string"`. Required.

Namespace where the CRL exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

<a id="canonical-6cc71b8fe15e74e1bd9e5d1704796269abd9a11a50a64cccffbb598156163c0b"></a>

<a id="canonical-4e3cd88e24a352b10495a3ce1e4fdca032e9cfb03a2a0746f2809c59453ba2a6"></a>

## refresh_interval property — Property reference / 8df667a7129e / 10

Type: `"number"`. Computed.

CRL Refresh interval. CRL refresh interval, in hours.

Upstream description:

CRL refresh interval, in hours.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 168,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 6
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "6",
    "ves.io.schema.rules.uint32.lte": "168"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "6",
    "ves.io.schema.rules.uint32.lte": "168"
  }
}
```

<a id="canonical-ee35c62d2e599d2d2877a1a9bf4542e2399134b8846c38132f92023999c54124"></a>

<a id="canonical-d850daa4e9fc9289f224efb1f695f5ff0c0746471ff8bb4312564a1e8371e334"></a>

## server_address property — Property reference / 8df667a7129e / 11

Type: `"string"`. Computed.

CRL Server address. CRL server address or hostname.

Upstream description:

CRL server address or hostname.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
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
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-5db570a26747a474d535f359b1e3162f526c1032e6ec8be489dd4af6c90bfd8c"></a>

<a id="canonical-4a19c50e8d5d422725c09377ed011b130f625bdc1736a6350e812105d8677085"></a>

## server_port property — Property reference / 8df667a7129e / 12

Type: `"number"`. Computed.

CRL Server Port. Set CRL Server port number.

Upstream description:

Set CRL Server port number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-e1dd7cee661d1f244c5e393691a48b64dd81e4d150bbd69d80724a1b2ae8cfa0"></a>

<a id="canonical-ec5789f8fa6d747d376da1152412a08210c129820ec45493f29afc9284705d5a"></a>

## timeout property — Property reference / 8df667a7129e / 13

Type: `"number"`. Computed.

CRL download timeout. CRL download wait time, in seconds.

Upstream description:

CRL download wait time, in seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "180"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "180"
  }
}
```

<a id="canonical-820ef0ba69b84ff04313a3666fa0d0e10bc7c1335ad72a27fbad5b03aad4bf7b"></a>

## All schema paths — Property reference / 8df667a7129e / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--crl--reference--group-001.md#canonical-0ffc9469f0096cdf6175dde58d2297c921794618c944278d2a70b7a05474fa7d) |
| `description` | [description](data-sources--crl--reference--group-001.md#canonical-d3d9113b90eb247fd938650c3d049a0845657ca5eba27f76accb428f1b670754) |
| `http_access` | [http_access](data-sources--crl--reference--group-001.md#canonical-6bdb04cb446dda525e53a2c133f529e6d5948caf736ed5ecbad7d6336d28a2d8) |
| `http_access.path` | [http_access.path](data-sources--crl--reference--group-001.md#canonical-4161da7bf81001220e9c5319a5c145ecf733735b2ae8f98b63ac0f1ebcc62f6b) |
| `id` | [id](data-sources--crl--reference--group-001.md#canonical-5ee27e2daad2ac096d4ffe9ab02b1db2da5c63f5b0745c2b77ada1c78c4914b8) |
| `labels` | [labels](data-sources--crl--reference--group-001.md#canonical-d76f5795a90e77a3cf8ba993c1b4b5cff597a1c17f26d461dc5e0327ac16f9ae) |
| `name` | [name](data-sources--crl--reference--group-001.md#canonical-2c31db0e68e9c5cbc54a279b1c9a49d8db755df855a601a8be6737a017f506cc) |
| `namespace` | [namespace](data-sources--crl--reference--group-001.md#canonical-9186bd6ee7fa44a49325386aab786f5af6ec0cc32906de165b67e45d54927b12) |
| `refresh_interval` | [refresh_interval](data-sources--crl--reference--group-001.md#canonical-6cc71b8fe15e74e1bd9e5d1704796269abd9a11a50a64cccffbb598156163c0b) |
| `server_address` | [server_address](data-sources--crl--reference--group-001.md#canonical-ee35c62d2e599d2d2877a1a9bf4542e2399134b8846c38132f92023999c54124) |
| `server_port` | [server_port](data-sources--crl--reference--group-001.md#canonical-5db570a26747a474d535f359b1e3162f526c1032e6ec8be489dd4af6c90bfd8c) |
| `timeout` | [timeout](data-sources--crl--reference--group-001.md#canonical-e1dd7cee661d1f244c5e393691a48b64dd81e4d150bbd69d80724a1b2ae8cfa0) |

<a id="canonical-962c04e814bd51e63c0f08a3f9d6b9e3acfa6d2b6d3c3c3d470e48a64abb8540"></a>

## Next pages — Property reference / 8df667a7129e / 15

- [http_access](data-sources--crl--reference--group-001.md#canonical-4aa87e7995a0400aaedc2a39d9c859c6240bfba132883b6d4fe11d5ce565c5e5)
- [xcsh_crl](../data-sources/crl.md#canonical-73f9576382ba605e0c52a90b24f3c47e46f2fa2fe6a252eff356e5dcb2c0e943)

<a id="canonical-4aa87e7995a0400aaedc2a39d9c859c6240bfba132883b6d4fe11d5ce565c5e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2624fe6a037386bc4b7a48473b0922a3bf54361c0b6b996f522a9e11a41cd664"></a>

## http_access — http_access / e30cd70f443e / 2

Breadcrumbs:

- [xcsh_crl](../data-sources/crl.md#canonical-73f9576382ba605e0c52a90b24f3c47e46f2fa2fe6a252eff356e5dcb2c0e943)
- [Property reference](data-sources--crl--reference--group-001.md#canonical-c2c016728d6cc79099c35ae246c02f69a282f25fbac91acb6b380a38ba11c747)
- http_access

<a id="canonical-6bdb04cb446dda525e53a2c133f529e6d5948caf736ed5ecbad7d6336d28a2d8"></a>

Type: `"single"`. Computed.

Configuration parameter for http access.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f633ffb0c3366f58d37b44017c4728f5ce427fb8b2083a84575e908c181977a0"></a>

## Direct properties — http_access / e30cd70f443e / 3

<a id="canonical-4161da7bf81001220e9c5319a5c145ecf733735b2ae8f98b63ac0f1ebcc62f6b"></a>

<a id="canonical-62b52f27c7edc628d2155225d2992f0cb6677f219ebc2aef8cc73fc66453fc65"></a>

## path property — http_access / e30cd70f443e / 4

Type: `"string"`. Computed.

CRL File path. CRL file location.

Upstream description:

CRL file location.

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

<a id="canonical-9ed29575cfc52a632b986db457bdbdaba846574ab15bd5b1d4cc166eed7a189a"></a>

## Next pages — http_access / e30cd70f443e / 5

- [Property reference](data-sources--crl--reference--group-001.md#canonical-c2c016728d6cc79099c35ae246c02f69a282f25fbac91acb6b380a38ba11c747)
- [xcsh_crl](../data-sources/crl.md#canonical-73f9576382ba605e0c52a90b24f3c47e46f2fa2fe6a252eff356e5dcb2c0e943)
