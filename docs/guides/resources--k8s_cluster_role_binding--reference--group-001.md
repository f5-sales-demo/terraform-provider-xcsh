---
page_title: "xcsh_k8s_cluster_role_binding reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role_binding reference."
---

# xcsh_k8s_cluster_role_binding reference

<a id="canonical-daa8129e558b368a8f6eac480ce48c930cc68a1d199f4322ed9f9803d63de406"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7af192114202af3cfa20640b67192eb4d36134e76ef3333c891dff68f97c16d"></a>

## Property reference — Property reference / a12ccad41cef / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md#canonical-63630feae3ae0d980292a5b75763d0fbd203d3335f8367da4d44f60fef31c646)
- Property reference

<a id="canonical-461a83b4c921ca5fdff49821010fed990f35779244c91bd26914e41c5481f54e"></a>

## Direct properties — Property reference / a12ccad41cef / 3

<a id="canonical-ac9c96b94fe9450818dc04f2e3a33edc9b1a1393ec8feb4983e2f79d4f0b9f9a"></a>

<a id="canonical-8f5a3eb7a1ea06d89d52e90f09ad418258edc31a92849fbf598fc462b265d5e0"></a>

## annotations property — Property reference / a12ccad41cef / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

<a id="canonical-210c3c9afb1d40d4cb18dd05f8577dd5a3d8cc1ef618e41a05f0f52695c8a2e8"></a>

<a id="canonical-a3df6d5fd77a20f91a74cf7c0a16273c7f51b5de3501e08d3a275a36d0a6dac2"></a>

## description property — Property reference / a12ccad41cef / 5

Type: `"string"`. Optional.

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

<a id="canonical-ff136ee527641a81d5afe7bb286a604452c3fbd19f018d309c083777600d2b3f"></a>

<a id="canonical-41142f3b8674c23dfd869bccddf013ea23c14a7b3e22d3f7d496c692591c362d"></a>

## disable property — Property reference / a12ccad41cef / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

<a id="canonical-d94b5d0a48bc761cad66c38994f03b559f77811245e9eb9b8c309513a0724e83"></a>

<a id="canonical-21d85f3006e0f9cfd325bcdc442adfe05c7a6b6472f65f6674d30e4c6ccc4ab1"></a>

## id property — Property reference / a12ccad41cef / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [k8s_cluster_role](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-80bd75eb53e2bb7077e746f5841a3df5ff04c47505be8fa8039ec31069d4f519): complete subsection reference.

<a id="canonical-b9ec129384b633cf05e4afdb42c567a18de5a86edd0faf2c4b8ddb08c8f9507f"></a>

<a id="canonical-ff4b2f5be61c2905a749eabf2a5a48d6f73a303a87a73984ff1af66a896ada96"></a>

## labels property — Property reference / a12ccad41cef / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-d60e81cef8f74d97a2f868c51e7790614084884ea3eae0ef80fea27e0714410d"></a>

<a id="canonical-ae78ef96e0358d892fa5a4052025c97be59530145932f274969b73f3cfc5380b"></a>

## name property — Property reference / a12ccad41cef / 9

Type: `"string"`. Required.

Name of the K8S Cluster Role Binding. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
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

<a id="canonical-8eafb71c67ccabdacf52b4dd8fdc8b6a92de214d100c2a3a5a6a7a705c30281c"></a>

<a id="canonical-9877f9430833aa99284f2401fc7196640c49e4eed12bb9ca6299399f883c6f12"></a>

## namespace property — Property reference / a12ccad41cef / 10

Type: `"string"`. Required.

Namespace where the K8S Cluster Role Binding is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
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

- [subjects](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-e2b7b1f98eac911389cdf181f7ddd1e26e111909b2786a50379b42ad1cb1af7e): complete subsection reference.

- [timeouts](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-f21d4b794a2af044c3a3c43da791401b31b60855b2f3dfb029c08164a69e1bd0): complete subsection reference.

<a id="canonical-f853a6d76f2eda88f13fd44b2090f276c9f356759043f2892b0aeaf638f3d6c1"></a>

## All schema paths — Property reference / a12ccad41cef / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-ac9c96b94fe9450818dc04f2e3a33edc9b1a1393ec8feb4983e2f79d4f0b9f9a) |
| `description` | [description](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-210c3c9afb1d40d4cb18dd05f8577dd5a3d8cc1ef618e41a05f0f52695c8a2e8) |
| `disable` | [disable](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-ff136ee527641a81d5afe7bb286a604452c3fbd19f018d309c083777600d2b3f) |
| `id` | [id](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-d94b5d0a48bc761cad66c38994f03b559f77811245e9eb9b8c309513a0724e83) |
| `k8s_cluster_role` | [k8s_cluster_role](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-fa5f9f9589850c5e7128e92c3a9aae15be4596b39382c1482341935254bf762e) |
| `k8s_cluster_role.name` | [k8s_cluster_role.name](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-62b7b78ed93cf44e82c13f431e23a467fc386902d5d57014e8a2b0877c3640c4) |
| `k8s_cluster_role.namespace` | [k8s_cluster_role.namespace](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-7f6ede185d068782edbc136638495dd586790b984c4db8aabb2f1977374ba060) |
| `k8s_cluster_role.tenant` | [k8s_cluster_role.tenant](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-9814d6ee7f042145db7ec1e6d3d47eafb9430264fc57c7d543829b3de24504f4) |
| `labels` | [labels](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-b9ec129384b633cf05e4afdb42c567a18de5a86edd0faf2c4b8ddb08c8f9507f) |
| `name` | [name](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-d60e81cef8f74d97a2f868c51e7790614084884ea3eae0ef80fea27e0714410d) |
| `namespace` | [namespace](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-8eafb71c67ccabdacf52b4dd8fdc8b6a92de214d100c2a3a5a6a7a705c30281c) |
| `subjects` | [subjects](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-5bfab2fdc18904f1d49cdc778fc172dd5618c3acf45eb41286282aefeebe7596) |
| `subjects.group` | [subjects.group](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-c1e1fd998ec380e92fe82110784cca4a8b57852db4e323d226c308b875f2bfdb) |
| `subjects.service_account` | [subjects.service_account](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-ea144b3e8e17e4b36aaa8a616e87df47962cbf98236edd363720e9508bf57796) |
| `subjects.service_account.name` | [subjects.service_account.name](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-3e4f5d586df08102c97f808c98ee02f691aaa83e1ff7a2a84214ac31c3d855ba) |
| `subjects.service_account.namespace` | [subjects.service_account.namespace](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-d522fb09bb97d92d249b45e5f76d14829c32f6ac0a81f9de24d55b33d06934e7) |
| `subjects.user` | [subjects.user](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-00f8ce40ccf6930382cf3aafee6411695d9707e619c8bc92eee056558b805e3a) |
| `timeouts` | [timeouts](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-dfd08a730acb4d947b9861a6aa5d2e1bad74dd8afb3cb00905de7ad5842c8ccd) |
| `timeouts.create` | [timeouts.create](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-b423b028bd7087d0de64559f13609fee8df564bff28a601e939c0385503b8642) |
| `timeouts.delete` | [timeouts.delete](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-9c0e83030f538e2a89a398d91554a5569925db5e959e43bd4077bcf5e25d152d) |
| `timeouts.read` | [timeouts.read](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-5584b56172565de18e2383b92846942b705619cb21162038fa612ba47bdb6d3b) |
| `timeouts.update` | [timeouts.update](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-ce212f568df9ffc1043f8c8cff2788ca487993d771c5658dfe05646c19b2faf4) |

<a id="canonical-9187bcc5ecbd35d882381b9509ae7fb55a9fc8c81a6cfffca6b4599a652f736b"></a>

## Next pages — Property reference / a12ccad41cef / 12

- [k8s_cluster_role](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-80bd75eb53e2bb7077e746f5841a3df5ff04c47505be8fa8039ec31069d4f519)
- [subjects](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-e2b7b1f98eac911389cdf181f7ddd1e26e111909b2786a50379b42ad1cb1af7e)
- [timeouts](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-f21d4b794a2af044c3a3c43da791401b31b60855b2f3dfb029c08164a69e1bd0)
- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md#canonical-63630feae3ae0d980292a5b75763d0fbd203d3335f8367da4d44f60fef31c646)

<a id="canonical-80bd75eb53e2bb7077e746f5841a3df5ff04c47505be8fa8039ec31069d4f519"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92d667c07ca37562190ef2df04e4103db8d0272dc1db09807a6f08dce6511a0c"></a>

## k8s_cluster_role — k8s_cluster_role / bfbe4e3087fa / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md#canonical-63630feae3ae0d980292a5b75763d0fbd203d3335f8367da4d44f60fef31c646)
- [Property reference](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-daa8129e558b368a8f6eac480ce48c930cc68a1d199f4322ed9f9803d63de406)
- k8s_cluster_role

<a id="canonical-fa5f9f9589850c5e7128e92c3a9aae15be4596b39382c1482341935254bf762e"></a>

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
k8s_cluster_role {
  # Configure direct properties listed below.
}
```

<a id="canonical-e3b8b21ec383d5694ad6a321dbe6496a5dd2d448060d7ddaf91212b4f25cf985"></a>

## Direct properties — k8s_cluster_role / bfbe4e3087fa / 3

<a id="canonical-62b7b78ed93cf44e82c13f431e23a467fc386902d5d57014e8a2b0877c3640c4"></a>

<a id="canonical-bef9624d361f7cc36137c64820c746d2e5472c65eb7e216a768be284d591ab0c"></a>

## name property — k8s_cluster_role / bfbe4e3087fa / 4

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

<a id="canonical-7f6ede185d068782edbc136638495dd586790b984c4db8aabb2f1977374ba060"></a>

<a id="canonical-11e1275fd49a184c5229160c5b60728743f68bb11dacb011a1061c1d5c8de631"></a>

## namespace property — k8s_cluster_role / bfbe4e3087fa / 5

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

<a id="canonical-9814d6ee7f042145db7ec1e6d3d47eafb9430264fc57c7d543829b3de24504f4"></a>

<a id="canonical-b5548ab31cb75744a75293e47bb8104b205b49b9c5a36426931d759c9e2066b7"></a>

## tenant property — k8s_cluster_role / bfbe4e3087fa / 6

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

<a id="canonical-cfbe10773a6f563e7bb79b3ede6e8430cf77397fa35c14d182379887bdbf80b8"></a>

## Next pages — k8s_cluster_role / bfbe4e3087fa / 7

- [Property reference](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-daa8129e558b368a8f6eac480ce48c930cc68a1d199f4322ed9f9803d63de406)
- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md#canonical-63630feae3ae0d980292a5b75763d0fbd203d3335f8367da4d44f60fef31c646)

<a id="canonical-e2b7b1f98eac911389cdf181f7ddd1e26e111909b2786a50379b42ad1cb1af7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab503b9de2b75f6f3e32215780ed8723feba38e3c7f1e3b12ae3ada31f2cebb1"></a>

## subjects — subjects / 6d7d4b81cc93 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md#canonical-63630feae3ae0d980292a5b75763d0fbd203d3335f8367da4d44f60fef31c646)
- [Property reference](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-daa8129e558b368a8f6eac480ce48c930cc68a1d199f4322ed9f9803d63de406)
- subjects

<a id="canonical-5bfab2fdc18904f1d49cdc778fc172dd5618c3acf45eb41286282aefeebe7596"></a>

Type: `"object"`. list nested block, Optional.

List of subjects (user, group or service account) to which this role is bound.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("group",
    "service_account"),
  validators.ConflictingListObjectAttributes("group",
    "user"),
  validators.ConflictingListObjectAttributes("service_account",
    "user")}
```

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

Terraform syntax:

```terraform
subjects {
  # Configure direct properties listed below.
}
```

<a id="canonical-8e2530f033389da3a3f942842850933e537f4017403b3718d079360cff81fac4"></a>

## Direct properties — subjects / 6d7d4b81cc93 / 3

<a id="canonical-c1e1fd998ec380e92fe82110784cca4a8b57852db4e323d226c308b875f2bfdb"></a>

<a id="canonical-3eb4f3aa825f46f18963a46c45107d3aa383f88ba59de9059320b772cef3435f"></a>

## group property — subjects / 6d7d4b81cc93 / 4

Type: `"string"`. Optional.

Exclusive with \[service\_account user\] Group ID of the user group.

Upstream description:

Exclusive with \[service\_account user\] Group ID of the user group.

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

- [service_account](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-5f2f87a962b556b6474b54dc9db475de41dcd75ce1b566c9427d302b4062c064): complete subsection reference.

<a id="canonical-00f8ce40ccf6930382cf3aafee6411695d9707e619c8bc92eee056558b805e3a"></a>

<a id="canonical-be5f3e67c42e1db1eedf1db3838b4cb200ecc79d3ccda982db465a0cbec475dd"></a>

## user property — subjects / 6d7d4b81cc93 / 5

Type: `"string"`. Optional.

Exclusive with \[group service\_account\] User ID of the user.

Upstream description:

Exclusive with \[group service\_account\] User ID of the user.

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

<a id="canonical-aaea4aabd4a6448f99534f70efabc7b0d4b88b446896d5d2eba5d3212926d003"></a>

## Next pages — subjects / 6d7d4b81cc93 / 6

- [subjects.service_account](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-5f2f87a962b556b6474b54dc9db475de41dcd75ce1b566c9427d302b4062c064)
- [Property reference](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-daa8129e558b368a8f6eac480ce48c930cc68a1d199f4322ed9f9803d63de406)
- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md#canonical-63630feae3ae0d980292a5b75763d0fbd203d3335f8367da4d44f60fef31c646)

<a id="canonical-5f2f87a962b556b6474b54dc9db475de41dcd75ce1b566c9427d302b4062c064"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2158d88edc8a4b07bffd2da18114c6d65b9445798182e63244a4c98fe158f6f"></a>

## subjects.service_account — subjects.service_account / 684ce254e725 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md#canonical-63630feae3ae0d980292a5b75763d0fbd203d3335f8367da4d44f60fef31c646)
- [Property reference](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-daa8129e558b368a8f6eac480ce48c930cc68a1d199f4322ed9f9803d63de406)
- [subjects](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-e2b7b1f98eac911389cdf181f7ddd1e26e111909b2786a50379b42ad1cb1af7e)
- subjects.service_account

<a id="canonical-ea144b3e8e17e4b36aaa8a616e87df47962cbf98236edd363720e9508bf57796"></a>

Type: `"object"`. single nested block, Optional.

ServiceAccountType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name",
    "namespace")}
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
service_account {
  # Configure direct properties listed below.
}
```

<a id="canonical-e7cfbcdcaec05863ec2a76396125081cba620f994ba2a3c741458babc3115bd8"></a>

## Direct properties — subjects.service_account / 684ce254e725 / 3

<a id="canonical-3e4f5d586df08102c97f808c98ee02f691aaa83e1ff7a2a84214ac31c3d855ba"></a>

<a id="canonical-4a544da0f676eb5ab6a2ccda884ace3acabd3f4118b5e659ab1ffb206fac7daf"></a>

## name property — subjects.service_account / 684ce254e725 / 4

Type: `"string"`. Optional.

Name. Name of the service account.

Upstream description:

Name of the service account.

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

<a id="canonical-d522fb09bb97d92d249b45e5f76d14829c32f6ac0a81f9de24d55b33d06934e7"></a>

<a id="canonical-3dca47889c225c0559899a9bbf83c8306296fce2d242878efc7ec5b7a82af021"></a>

## namespace property — subjects.service_account / 684ce254e725 / 5

Type: `"string"`. Optional, Computed.

Namespace. Namespace of the service account.

Upstream description:

Namespace of the service account.

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

<a id="canonical-741550e5d43c790a90a471e8caa26f83415d25620cb2afae523b9cf2cca400ce"></a>

## Next pages — subjects.service_account / 684ce254e725 / 6

- [subjects](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-e2b7b1f98eac911389cdf181f7ddd1e26e111909b2786a50379b42ad1cb1af7e)
- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md#canonical-63630feae3ae0d980292a5b75763d0fbd203d3335f8367da4d44f60fef31c646)

<a id="canonical-f21d4b794a2af044c3a3c43da791401b31b60855b2f3dfb029c08164a69e1bd0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff07cd2347e9d12941d7c60554fe8e3ecadc7f616071b0511952b109c7c0a1b6"></a>

## timeouts — timeouts / df06927c6340 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md#canonical-63630feae3ae0d980292a5b75763d0fbd203d3335f8367da4d44f60fef31c646)
- [Property reference](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-daa8129e558b368a8f6eac480ce48c930cc68a1d199f4322ed9f9803d63de406)
- timeouts

<a id="canonical-dfd08a730acb4d947b9861a6aa5d2e1bad74dd8afb3cb00905de7ad5842c8ccd"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-7645913f74f2659ae2242f1fa2748a4ef87e9996d66990648c3625eb5aa407f3"></a>

## Direct properties — timeouts / df06927c6340 / 3

<a id="canonical-b423b028bd7087d0de64559f13609fee8df564bff28a601e939c0385503b8642"></a>

<a id="canonical-0ced3bc0a74e5d1f5dcb95a3852836f4a40f49cb3edb7509d9111735040f0c37"></a>

## create property — timeouts / df06927c6340 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-9c0e83030f538e2a89a398d91554a5569925db5e959e43bd4077bcf5e25d152d"></a>

<a id="canonical-2e1db2f8a0123835cc83acf331f3c27d8f53b9e3185b7fae03b36c30702f7965"></a>

## delete property — timeouts / df06927c6340 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-5584b56172565de18e2383b92846942b705619cb21162038fa612ba47bdb6d3b"></a>

<a id="canonical-f0febe42b1c80108001b73166a351cd29114e16ed981d26c214aeb641a5a5acc"></a>

## read property — timeouts / df06927c6340 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-ce212f568df9ffc1043f8c8cff2788ca487993d771c5658dfe05646c19b2faf4"></a>

<a id="canonical-dadec23ff53d2d597bc853d29e55f291e46e098003e46c99ae7f26abed1a0a10"></a>

## update property — timeouts / df06927c6340 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-c2dc889a938c4b0f77162a56551149f73ef8b5368c67f09c23c0d5689fbf7eb0"></a>

## Next pages — timeouts / df06927c6340 / 8

- [Property reference](resources--k8s_cluster_role_binding--reference--group-001.md#canonical-daa8129e558b368a8f6eac480ce48c930cc68a1d199f4322ed9f9803d63de406)
- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md#canonical-63630feae3ae0d980292a5b75763d0fbd203d3335f8367da4d44f60fef31c646)
