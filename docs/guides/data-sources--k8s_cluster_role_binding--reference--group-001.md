---
page_title: "xcsh_k8s_cluster_role_binding reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role_binding reference."
---

# xcsh_k8s_cluster_role_binding reference

<a id="canonical-bc8de5758cacd31e5791b1d4bd3baf93288975699ce14c1facc574a0223a8a97"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ff4495a000b4ab8452cd491c5e624d1f7edd049f56dee50bab731baef83aebf"></a>

## Property reference — Property reference / 0261c19dff7b / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-89aa41927c7201d228ed76b699b17dcb5b4c9bac74cfeefe90e715c6cd60d347)
- Property reference

<a id="canonical-8b9fda7fb16520ca91ff85e5af3a1f000abd5511f4ac74f2462ecc598f292641"></a>

## Direct properties — Property reference / 0261c19dff7b / 3

<a id="canonical-79dd889f70f5c5cd4c237775133bf7e89a0d97d10d3a483e402e3ce6b2b96106"></a>

<a id="canonical-bdecee0bf15eae7d0795ee959ada28de1f628712ce78fb9acb829cb6b200e1c7"></a>

## annotations property — Property reference / 0261c19dff7b / 4

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

<a id="canonical-379a7135e8fa459827fe90a2ed933fb46a783100c6889937df8df1b098f7d15b"></a>

<a id="canonical-096c6e941d6ee6638a286b3b713f7e03e8bc0956ea86c0847143228b285f00fc"></a>

## description property — Property reference / 0261c19dff7b / 5

Type: `"string"`. Computed.

Description of the K8SClusterRoleBinding.

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

<a id="canonical-277698258758f90f5664721531e0b5ababdfd2bba44f0f06dfef8cad4ac8f9f7"></a>

<a id="canonical-8d464501e6a1310e49fa6d68ba5c5cede1e21e456216011d4e094955bff30e75"></a>

## id property — Property reference / 0261c19dff7b / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [k8s_cluster_role](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-711439d734548af578ebdcbee7e77fdf843a0a81900197ff7a32fdcf11e9f201): complete subsection reference.

<a id="canonical-ca4a1199a6de7e1520fcc0ec5c82f21a4570cbe72efc612541c16376f6ab5327"></a>

<a id="canonical-cd9b28df6f6fcb48a6b8a22fc5f75944460d8f43c53025fcd75e6cdccacf0ec6"></a>

## labels property — Property reference / 0261c19dff7b / 7

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

<a id="canonical-785b6da14dbe2f114f5b9616cd9c7359bd35cc9283c1b7ca91db9059d641caaa"></a>

<a id="canonical-5b698836acb32be3524c2a4498197a5e1c93159c9521e2bcdfeb919bebb41584"></a>

## name property — Property reference / 0261c19dff7b / 8

Type: `"string"`. Required.

Name of the K8SClusterRoleBinding.

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

<a id="canonical-a979a226e03925e2301c566b8d89b66d46ee5f1018d9e25e2439942fea6fc834"></a>

<a id="canonical-f6b9215d0a8eafcd4d41c1341ab262c75fb570709ce09535dbfb983495f3abf5"></a>

## namespace property — Property reference / 0261c19dff7b / 9

Type: `"string"`. Required.

Namespace where the K8SClusterRoleBinding exists.

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

- [subjects](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-639f073759ca5af74a41d91ef26f58de007e87da6ce8765e3ad5965b602342f4): complete subsection reference.

<a id="canonical-a16c4b6be22018d7f967f1b67238d00a6e10029e4526b29915b08521f8f74ad1"></a>

## All schema paths — Property reference / 0261c19dff7b / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-79dd889f70f5c5cd4c237775133bf7e89a0d97d10d3a483e402e3ce6b2b96106) |
| `description` | [description](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-379a7135e8fa459827fe90a2ed933fb46a783100c6889937df8df1b098f7d15b) |
| `id` | [id](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-277698258758f90f5664721531e0b5ababdfd2bba44f0f06dfef8cad4ac8f9f7) |
| `k8s_cluster_role` | [k8s_cluster_role](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-f64dbda3e78796ce26bd0c7ae4fef9aa55c8fd2e2f49b6fa6b4fb00dc922c7ba) |
| `k8s_cluster_role.name` | [k8s_cluster_role.name](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-bf07f41778506775a517e74cc43cb7496589a64af9be5f40012697eb133f02d2) |
| `k8s_cluster_role.namespace` | [k8s_cluster_role.namespace](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-1650e5166acc16f57d6b460f92a7a6191654741255665244b6130245aee7f2d8) |
| `k8s_cluster_role.tenant` | [k8s_cluster_role.tenant](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-bb9705b258696dea670194a6c1884a616d74da20063a00e672c1b2357b1eeecc) |
| `labels` | [labels](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-ca4a1199a6de7e1520fcc0ec5c82f21a4570cbe72efc612541c16376f6ab5327) |
| `name` | [name](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-785b6da14dbe2f114f5b9616cd9c7359bd35cc9283c1b7ca91db9059d641caaa) |
| `namespace` | [namespace](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-a979a226e03925e2301c566b8d89b66d46ee5f1018d9e25e2439942fea6fc834) |
| `subjects` | [subjects](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-b49544a6c5441592f5a95fe5583ca7e8dbffb9c374c0fb2f2d9f69fa38fd5285) |
| `subjects.group` | [subjects.group](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-0204f48025e3e0c99345093bdb736332886437a729fe29d8edbebf3bc18933a4) |
| `subjects.service_account` | [subjects.service_account](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-4835fd8c01abc62489c78f8f65fe36911ec39814e2ad844ee8ef78ab22f50c4d) |
| `subjects.service_account.name` | [subjects.service_account.name](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-80a7d6ce6bb51bb1365601f6489c62f4ec193007fa70c4d86bcd233435d39830) |
| `subjects.service_account.namespace` | [subjects.service_account.namespace](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-d7145a10f4acad35ef4da6f9442008fd4a1eebd000ece80169103b3fd2e2949e) |
| `subjects.user` | [subjects.user](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-a679272f2a2ac79744633e1d989eea847d413b6a59121dd2c8fa37682d45e25d) |

<a id="canonical-6b29e75fa1a242ba43f7c1d8fbad8123a2211a2e75d8aed81baa4b695a71b998"></a>

## Next pages — Property reference / 0261c19dff7b / 11

- [k8s_cluster_role](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-711439d734548af578ebdcbee7e77fdf843a0a81900197ff7a32fdcf11e9f201)
- [subjects](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-639f073759ca5af74a41d91ef26f58de007e87da6ce8765e3ad5965b602342f4)
- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-89aa41927c7201d228ed76b699b17dcb5b4c9bac74cfeefe90e715c6cd60d347)

<a id="canonical-711439d734548af578ebdcbee7e77fdf843a0a81900197ff7a32fdcf11e9f201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65b82aea633ba2b1a6eaa851f233d01828ee7184707ede784c420ad73cfe2c2a"></a>

## k8s_cluster_role — k8s_cluster_role / 9dc4bac09616 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-89aa41927c7201d228ed76b699b17dcb5b4c9bac74cfeefe90e715c6cd60d347)
- [Property reference](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-bc8de5758cacd31e5791b1d4bd3baf93288975699ce14c1facc574a0223a8a97)
- k8s_cluster_role

<a id="canonical-f64dbda3e78796ce26bd0c7ae4fef9aa55c8fd2e2f49b6fa6b4fb00dc922c7ba"></a>

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

<a id="canonical-e8fa16d64b530adbc2e754a6bcf6b05723030e3f2e197e19ce86272c12876aa1"></a>

## Direct properties — k8s_cluster_role / 9dc4bac09616 / 3

<a id="canonical-bf07f41778506775a517e74cc43cb7496589a64af9be5f40012697eb133f02d2"></a>

<a id="canonical-b0f8d0a05ef0e89691f0070408bb6495afc53a092d58ad53c4d0e3acdeaa62f4"></a>

## name property — k8s_cluster_role / 9dc4bac09616 / 4

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

<a id="canonical-1650e5166acc16f57d6b460f92a7a6191654741255665244b6130245aee7f2d8"></a>

<a id="canonical-997b37d1c65301eb51cff91a11ea2c25523c9b2239bb54de8fc0f4d476d14402"></a>

## namespace property — k8s_cluster_role / 9dc4bac09616 / 5

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

<a id="canonical-bb9705b258696dea670194a6c1884a616d74da20063a00e672c1b2357b1eeecc"></a>

<a id="canonical-3e0fc2b70457b9de31512f028355e581ca0da03f44318e2f6e653efd90cb4d42"></a>

## tenant property — k8s_cluster_role / 9dc4bac09616 / 6

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

<a id="canonical-c50aeecdfa60d618612936738a1d841a926264424488170cb76b803e9d61b2ae"></a>

## Next pages — k8s_cluster_role / 9dc4bac09616 / 7

- [Property reference](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-bc8de5758cacd31e5791b1d4bd3baf93288975699ce14c1facc574a0223a8a97)
- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-89aa41927c7201d228ed76b699b17dcb5b4c9bac74cfeefe90e715c6cd60d347)

<a id="canonical-639f073759ca5af74a41d91ef26f58de007e87da6ce8765e3ad5965b602342f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9f92c2d579a7440ed25be026ec98a6c066961e2f5b3beb5b6f5243ff048b096"></a>

## subjects — subjects / ebbb2f90b570 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-89aa41927c7201d228ed76b699b17dcb5b4c9bac74cfeefe90e715c6cd60d347)
- [Property reference](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-bc8de5758cacd31e5791b1d4bd3baf93288975699ce14c1facc574a0223a8a97)
- subjects

<a id="canonical-b49544a6c5441592f5a95fe5583ca7e8dbffb9c374c0fb2f2d9f69fa38fd5285"></a>

Type: `"list"`. Computed.

List of subjects (user, group or service account) to which this role is bound.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1231b54c2738939732b16054b3ac79d9b395215a3c101a438758b2a96bd2dbc4"></a>

## Direct properties — subjects / ebbb2f90b570 / 3

<a id="canonical-0204f48025e3e0c99345093bdb736332886437a729fe29d8edbebf3bc18933a4"></a>

<a id="canonical-9c7d63d4786d7fabba6e5100d775d68431345f1cb2202948a621235da7b10540"></a>

## group property — subjects / ebbb2f90b570 / 4

Type: `"string"`. Computed.

Exclusive with \[service\_account user\] Group ID of the user group.

Upstream description:

Exclusive with \[service\_account user\] Group ID of the user group.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [service_account](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-32e2769d9e9272aeb7b8a0a4a60c4f4c3e182d8992ab7a45e68a4af831cfed7e): complete subsection reference.

<a id="canonical-a679272f2a2ac79744633e1d989eea847d413b6a59121dd2c8fa37682d45e25d"></a>

<a id="canonical-1d28d0a70e26c73cb2e57e0b4ce935ca43f136ff6c1e453705559bac4b455aeb"></a>

## user property — subjects / ebbb2f90b570 / 5

Type: `"string"`. Computed.

Exclusive with \[group service\_account\] User ID of the user.

Upstream description:

Exclusive with \[group service\_account\] User ID of the user.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-42c3ab1b526b3195581edf9a513e2a698073b86e47ef5e6d36768c63c66e2d7c"></a>

## Next pages — subjects / ebbb2f90b570 / 6

- [subjects.service_account](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-32e2769d9e9272aeb7b8a0a4a60c4f4c3e182d8992ab7a45e68a4af831cfed7e)
- [Property reference](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-bc8de5758cacd31e5791b1d4bd3baf93288975699ce14c1facc574a0223a8a97)
- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-89aa41927c7201d228ed76b699b17dcb5b4c9bac74cfeefe90e715c6cd60d347)

<a id="canonical-32e2769d9e9272aeb7b8a0a4a60c4f4c3e182d8992ab7a45e68a4af831cfed7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c332883b271f209db818c3e9866c4d0ea192da393f11c78efa56a5c914081a4d"></a>

## subjects.service_account — subjects.service_account / 2c645b59e94f / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-89aa41927c7201d228ed76b699b17dcb5b4c9bac74cfeefe90e715c6cd60d347)
- [Property reference](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-bc8de5758cacd31e5791b1d4bd3baf93288975699ce14c1facc574a0223a8a97)
- [subjects](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-639f073759ca5af74a41d91ef26f58de007e87da6ce8765e3ad5965b602342f4)
- subjects.service_account

<a id="canonical-4835fd8c01abc62489c78f8f65fe36911ec39814e2ad844ee8ef78ab22f50c4d"></a>

Type: `"single"`. Computed.

ServiceAccountType.

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

<a id="canonical-a4c98be37ce6fbbeca49af9d2d83030f37f9b6d6d6f810e3e3be56fa53c3c392"></a>

## Direct properties — subjects.service_account / 2c645b59e94f / 3

<a id="canonical-80a7d6ce6bb51bb1365601f6489c62f4ec193007fa70c4d86bcd233435d39830"></a>

<a id="canonical-b9a6262c311b7178956669af820c042e0da7972741fefa0ed85a944f7174360f"></a>

## name property — subjects.service_account / 2c645b59e94f / 4

Type: `"string"`. Computed.

Name. Name of the service account.

Upstream description:

Name of the service account.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-d7145a10f4acad35ef4da6f9442008fd4a1eebd000ece80169103b3fd2e2949e"></a>

<a id="canonical-ddad522589990513088c06daeb3d404d32a3d10f85496cdb73e56b93e9e9a724"></a>

## namespace property — subjects.service_account / 2c645b59e94f / 5

Type: `"string"`. Computed.

Namespace. Namespace of the service account.

Upstream description:

Namespace of the service account.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-aa8b66a2a98461d8e4298bf9bab4b19950b924d769b3d50a4a85170d3378bd29"></a>

## Next pages — subjects.service_account / 2c645b59e94f / 6

- [subjects](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-639f073759ca5af74a41d91ef26f58de007e87da6ce8765e3ad5965b602342f4)
- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-89aa41927c7201d228ed76b699b17dcb5b4c9bac74cfeefe90e715c6cd60d347)
