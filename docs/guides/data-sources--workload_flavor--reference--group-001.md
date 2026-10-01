---
page_title: "xcsh_workload_flavor reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_workload_flavor reference."
---

# xcsh_workload_flavor reference

<a id="canonical-9193b6067e4e8e493c8d2e08f2541bfa4d5b090c7074a32eae1dc6e93462dee9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d40a01888106cf793d8c5bbf00a03b0594bd0276b870dfa82f3342c28ad9669"></a>

## Property reference — Property reference / cde2debddf59 / 2

Breadcrumbs:

- [xcsh_workload_flavor](../data-sources/workload_flavor.md#canonical-8c8c0ecc2e6ee25fc807e107477b515e49f3e31343c02d3b3fb96b576dfd8acc)
- Property reference

<a id="canonical-2af589003cbd186f557c0c08f7881e199ff268a1cac0becf7f6a19825204948e"></a>

## Direct properties — Property reference / cde2debddf59 / 3

<a id="canonical-8124638f794b435ca9e047d5a9ca035fee7c97d26a4aeac8a5ba5b6767b112e9"></a>

<a id="canonical-2ec67854f55265e917666521d018e386689e03c38b243b8661bd0edcc173e261"></a>

## annotations property — Property reference / cde2debddf59 / 4

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

<a id="canonical-1449561cc82925dc45a9927d047aa5cd4912a967ae9177a663dcf6ffbf36d229"></a>

<a id="canonical-ce9a5b4a2acde15dc38505d524912d71a89f98399221e77efed797960c5f73b0"></a>

## description property — Property reference / cde2debddf59 / 5

Type: `"string"`. Computed.

Description of the WorkloadFlavor.

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

<a id="canonical-c524eb7aecc0cbb7fc3aa6eff6be9ae6f2470fd2fdb8f371e31ad565489ea1fb"></a>

<a id="canonical-3c4b8f0cbd4a738158338c793d06145a76c8921337e290fb5811d62b651e0374"></a>

## ephemeral_storage property — Property reference / cde2debddf59 / 6

Type: `"string"`. Computed.

Ephemeral storage in MiB (mebibyte) allocated for the workload\_flavor.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
    "ves.io.schema.rules.uint32.ranges": "1-6000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "1-6000"
  }
}
```

<a id="canonical-6e964ffd646fbfb59be9a98584f0b99b3b3823cb498eec78e957fd92766ceb71"></a>

<a id="canonical-2b84e66eaf7e93d49f911a733bc7db8de7aaf7403197af38b02a64eefad24303"></a>

## id property — Property reference / cde2debddf59 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-d81e9d4ba69c92e779662cdd8c5e57c2a94cf4b764b7c31f9c242fcd980f7458"></a>

<a id="canonical-9fd08b734ee49fd70b32c52b7ad6bc516d5024d97f90b9d1e642bf76ecb651be"></a>

## labels property — Property reference / cde2debddf59 / 8

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

<a id="canonical-04c0280da3dbbbc6220c202a62c7d79c9c78f781aeaf1ef3adf58c4d1df922f5"></a>

<a id="canonical-aea65405c926218d26e8461ac785d9815133415f93b399daa081155dae549665"></a>

## memory property — Property reference / cde2debddf59 / 9

Type: `"string"`. Computed.

Memory in MiB (mebibyte) allocated for the workload\_flavor.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
    "ves.io.schema.rules.uint32.ranges": "1-32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "1-32768"
  }
}
```

<a id="canonical-e9e29da6fb721d928c28ad23781e78854841fd5391ef2eebf01f57e1de2f00e2"></a>

<a id="canonical-6533fcb75ed55e524339a0a68f2efe019c880235f114dba6b97b49d7526af0c7"></a>

## name property — Property reference / cde2debddf59 / 10

Type: `"string"`. Required.

Name of the WorkloadFlavor.

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

<a id="canonical-fd946b67c23978db991b2dc0938cada3ef1dfe4f4edfb873324ef307c88f6222"></a>

<a id="canonical-1a61727f3c35f4964d7707ac6d5966ece11ef452fd4f7b697bca5e30520f7b28"></a>

## namespace property — Property reference / cde2debddf59 / 11

Type: `"string"`. Optional, Computed.

Namespace where the WorkloadFlavor exists.

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

<a id="canonical-0156958071eba80c5632b7a479d2fe14b2d4d49de50f3bab3a607123bacd123f"></a>

<a id="canonical-2bda2056f773402b4d9a08519fbc30f1fdfcd78828d1e150f5b1b1b9ef400892"></a>

## vcpus property — Property reference / cde2debddf59 / 12

Type: `"number"`. Computed.

Number of vCPUs allocated for the workload\_flavor. Each vCPU is a thread on a CPU core.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "8.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "8.0"
  }
}
```

<a id="canonical-0c785e30a6f9b9cbc3a9496c6bcb34d93575b3fa4cb13c30914ab4bff43a8dad"></a>

## All schema paths — Property reference / cde2debddf59 / 13

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--workload_flavor--reference--group-001.md#canonical-8124638f794b435ca9e047d5a9ca035fee7c97d26a4aeac8a5ba5b6767b112e9) |
| `description` | [description](data-sources--workload_flavor--reference--group-001.md#canonical-1449561cc82925dc45a9927d047aa5cd4912a967ae9177a663dcf6ffbf36d229) |
| `ephemeral_storage` | [ephemeral_storage](data-sources--workload_flavor--reference--group-001.md#canonical-c524eb7aecc0cbb7fc3aa6eff6be9ae6f2470fd2fdb8f371e31ad565489ea1fb) |
| `id` | [id](data-sources--workload_flavor--reference--group-001.md#canonical-6e964ffd646fbfb59be9a98584f0b99b3b3823cb498eec78e957fd92766ceb71) |
| `labels` | [labels](data-sources--workload_flavor--reference--group-001.md#canonical-d81e9d4ba69c92e779662cdd8c5e57c2a94cf4b764b7c31f9c242fcd980f7458) |
| `memory` | [memory](data-sources--workload_flavor--reference--group-001.md#canonical-04c0280da3dbbbc6220c202a62c7d79c9c78f781aeaf1ef3adf58c4d1df922f5) |
| `name` | [name](data-sources--workload_flavor--reference--group-001.md#canonical-e9e29da6fb721d928c28ad23781e78854841fd5391ef2eebf01f57e1de2f00e2) |
| `namespace` | [namespace](data-sources--workload_flavor--reference--group-001.md#canonical-fd946b67c23978db991b2dc0938cada3ef1dfe4f4edfb873324ef307c88f6222) |
| `vcpus` | [vcpus](data-sources--workload_flavor--reference--group-001.md#canonical-0156958071eba80c5632b7a479d2fe14b2d4d49de50f3bab3a607123bacd123f) |

<a id="canonical-d1fba6ebfe89c69e60578b94e8b6776fd2464bec2e87a5b5b01c475b05c111d5"></a>

## Next pages — Property reference / cde2debddf59 / 14

- [xcsh_workload_flavor](../data-sources/workload_flavor.md#canonical-8c8c0ecc2e6ee25fc807e107477b515e49f3e31343c02d3b3fb96b576dfd8acc)
