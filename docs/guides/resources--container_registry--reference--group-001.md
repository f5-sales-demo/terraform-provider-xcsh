---
page_title: "xcsh_container_registry reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_container_registry reference."
---

# xcsh_container_registry reference

<a id="canonical-4e616e01aa1598e2fc141ba33f7c752da2e7a66e59191d4c102a8adb6557faa3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33cebbce06eb73c45abe480f368bea9365d30bf9927e7a2ce26285172a27c786"></a>

## Property reference — Property reference / 80f8915612d9 / 2

Breadcrumbs:

- [xcsh_container_registry](../resources/container_registry.md#canonical-d9354f0149bcec7d60dd61bf8c2faf0289b61a944b3928d453a60e87522be1fe)
- Property reference

<a id="canonical-b4a88c3466b2640d2f54c183defaea8d61fde7854e8af852c0233fddac7d9e33"></a>

## Direct properties — Property reference / 80f8915612d9 / 3

<a id="canonical-f016c5376d6ef438b1e8ee74722e060afcc723d254ca5eff7821b54b264a1cb5"></a>

<a id="canonical-3f00ddf76dafcfa8e308330fa885b2ae1bac856fb7eb8ead9a5d40e144398632"></a>

## annotations property — Property reference / 80f8915612d9 / 4

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

<a id="canonical-411fcaa8965bea4c5e5a62c79f6a6ed8f2b22068c73221e7ccb4879734f9ba9d"></a>

<a id="canonical-5198359f82606f393f2f7d15e25c42ab9b572f8023abc5529802ea4000918b7f"></a>

## description property — Property reference / 80f8915612d9 / 5

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

<a id="canonical-5974b5b28e6a5de81446b5199a46d06a13314b48153eadeba283b58f4eeb6b30"></a>

<a id="canonical-8862873a87f7e4572bb5da09b0fa24a8f17bc78103ccb38597b55d3e4566768e"></a>

## disable property — Property reference / 80f8915612d9 / 6

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

<a id="canonical-19d782bb93cc09f4f26358e11c31384d3d0233368ae65d3e79d5954db3380588"></a>

<a id="canonical-3525c201100f7d2eccbc5450b459d0bcef4491679a3f19e3b714e75e2c54403c"></a>

## email property — Property reference / 80f8915612d9 / 7

Type: `"string"`. Optional, Computed.

Email. Email used for the registry.

Upstream description:

Email used for the registry.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "email",
    "formatDescription": "RFC 5322 email address, max 254 characters",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
    "validation": {
      "rfc": "RFC 5322"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.email": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.email": "true"
  }
}
```

<a id="canonical-ed1c30462b02e31bad96c41888b303d5c0845694e358750720deef62dad5a9c4"></a>

<a id="canonical-ab6dff1be54e9fbc7bd144259a4d9cca2a0321d9e7a76be57c9a77c8465f889f"></a>

## id property — Property reference / 80f8915612d9 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-461d6ab8e716abdb6ce563a795942741599c3bb067d826008685c69518c4610a"></a>

<a id="canonical-cf4072dc981bf425365587cf51f4815c52f6640a191b22bc3103272b34388890"></a>

## labels property — Property reference / 80f8915612d9 / 9

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

<a id="canonical-24a29117be7cb3e97153bd66469e272ab37a343e14699e90eefd93ed4bd94f7d"></a>

<a id="canonical-8bbcfd8ec3298908fc59a4d4c2bed04befb1f337af41e431ab623ce1528ed13c"></a>

## name property — Property reference / 80f8915612d9 / 10

Type: `"string"`. Required.

Name of the Container Registry. Must be unique within the namespace.

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

<a id="canonical-d2f6fa801cf1c8cd0d5e50d2e1d661c0fa0f537cfa74b07b653fd4ad9b42fb0c"></a>

<a id="canonical-e2aa18206e0578e0216a02cb5e964936cfc00a21cc3d9d72d32ab5d07271a1c4"></a>

## namespace property — Property reference / 80f8915612d9 / 11

Type: `"string"`. Required.

Namespace where the Container Registry is created.

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

- [password](resources--container_registry--reference--group-001.md#canonical-bfc4842267fc635196a8710e5f991005dd73c95e97df089feebe338708b210cc): complete subsection reference.

<a id="canonical-dbaaade7a4a9672a296d66429c9d6dbb02da46a470941b12bd386052f6685016"></a>

<a id="canonical-60bd1df803ec94af0ba62a0bf340427165ba50b7d17c6fdbe73df9ccb46b7194"></a>

## registry property — Property reference / 80f8915612d9 / 12

Type: `"string"`. Required.

Fully qualified name of the registry login server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [timeouts](resources--container_registry--reference--group-001.md#canonical-9f566e8474c98272d7a45bdda4a03ebc7c8aaa0c8a664de216696b684c915244): complete subsection reference.

<a id="canonical-1f8dda57063c3823f80d64278423280045d4c374bfc6aab6fd55723dc45ddc9c"></a>

<a id="canonical-72218cbab434c5df25e2a1c847c92f604d64489c17bef9f6f5dd896405e399f6"></a>

## user_name property — Property reference / 80f8915612d9 / 13

Type: `"string"`. Required.

User Name. Username used to access the registry.

Upstream description:

Username used to access the registry.

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-898a7a0eedf6b06b235e1092ae48fc3f116ef5dfce486bfbf2ac491710e5a23b"></a>

## All schema paths — Property reference / 80f8915612d9 / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--container_registry--reference--group-001.md#canonical-f016c5376d6ef438b1e8ee74722e060afcc723d254ca5eff7821b54b264a1cb5) |
| `description` | [description](resources--container_registry--reference--group-001.md#canonical-411fcaa8965bea4c5e5a62c79f6a6ed8f2b22068c73221e7ccb4879734f9ba9d) |
| `disable` | [disable](resources--container_registry--reference--group-001.md#canonical-5974b5b28e6a5de81446b5199a46d06a13314b48153eadeba283b58f4eeb6b30) |
| `email` | [email](resources--container_registry--reference--group-001.md#canonical-19d782bb93cc09f4f26358e11c31384d3d0233368ae65d3e79d5954db3380588) |
| `id` | [id](resources--container_registry--reference--group-001.md#canonical-ed1c30462b02e31bad96c41888b303d5c0845694e358750720deef62dad5a9c4) |
| `labels` | [labels](resources--container_registry--reference--group-001.md#canonical-461d6ab8e716abdb6ce563a795942741599c3bb067d826008685c69518c4610a) |
| `name` | [name](resources--container_registry--reference--group-001.md#canonical-24a29117be7cb3e97153bd66469e272ab37a343e14699e90eefd93ed4bd94f7d) |
| `namespace` | [namespace](resources--container_registry--reference--group-001.md#canonical-d2f6fa801cf1c8cd0d5e50d2e1d661c0fa0f537cfa74b07b653fd4ad9b42fb0c) |
| `password` | [password](resources--container_registry--reference--group-001.md#canonical-6c277754435430acc53bcfa5c0d782c33e843e2e49b310cc10378e03c55d4eaa) |
| `password.blindfold_secret_info` | [password.blindfold_secret_info](resources--container_registry--reference--group-001.md#canonical-c3a4daf92dafaa9096aa8543d5f67edb5ec2fe221da52c7c28f89792d29888bf) |
| `password.blindfold_secret_info.decryption_provider` | [password.blindfold_secret_info.decryption_provider](resources--container_registry--reference--group-001.md#canonical-099d8ad9a054ff4a26a244293280b89fc643a13074727bf713e432fe963c329c) |
| `password.blindfold_secret_info.location` | [password.blindfold_secret_info.location](resources--container_registry--reference--group-001.md#canonical-e4930cb04e20c091c5db445ed38332dcb34d95f89b604ef262a41f092207b051) |
| `password.blindfold_secret_info.store_provider` | [password.blindfold_secret_info.store_provider](resources--container_registry--reference--group-001.md#canonical-b68187b7d94d172133dec3ac70470a2be205fa9426c977e0f1d7155270de4ffa) |
| `password.clear_secret_info` | [password.clear_secret_info](resources--container_registry--reference--group-001.md#canonical-abc503d13d548c99be9146355f4005615fa74624fdbe42daf766e470e57b128d) |
| `password.clear_secret_info.provider_ref` | [password.clear_secret_info.provider_ref](resources--container_registry--reference--group-001.md#canonical-3a45b45e7bc0818285affe13b08ab7f8825cf68786afaf0bbdc56dc1f6dc23b9) |
| `password.clear_secret_info.url` | [password.clear_secret_info.url](resources--container_registry--reference--group-001.md#canonical-98c7e0eb28b2eab40479f9cc9fca5730614d780ec01163dd65a7e7be05c3cea2) |
| `registry` | [registry](resources--container_registry--reference--group-001.md#canonical-dbaaade7a4a9672a296d66429c9d6dbb02da46a470941b12bd386052f6685016) |
| `timeouts` | [timeouts](resources--container_registry--reference--group-001.md#canonical-8470bafa6128d257703d671043520873009717cb7e24de1af5aff682f1b53921) |
| `timeouts.create` | [timeouts.create](resources--container_registry--reference--group-001.md#canonical-c0ac067b91a9532f20bbde178afc3ba1486a66e14421c823b567fc57e02649c0) |
| `timeouts.delete` | [timeouts.delete](resources--container_registry--reference--group-001.md#canonical-d63d78a8899d16b4e5a94cb632cae3a0286d723494884da92721f04528772881) |
| `timeouts.read` | [timeouts.read](resources--container_registry--reference--group-001.md#canonical-70a9a0f1500eeb83835d9b0ebe36bfd03ab8c3c0949d6edcb1d9b719ae56c4b4) |
| `timeouts.update` | [timeouts.update](resources--container_registry--reference--group-001.md#canonical-25e2772f549af7c1bb429ab87390812cfc2f123529327dff0a1a101dbf295f97) |
| `user_name` | [user_name](resources--container_registry--reference--group-001.md#canonical-1f8dda57063c3823f80d64278423280045d4c374bfc6aab6fd55723dc45ddc9c) |

<a id="canonical-b8030948bde17045cf548cf9090d3bbf351bca541a26d111178a542c29d56c0d"></a>

## Next pages — Property reference / 80f8915612d9 / 15

- [password](resources--container_registry--reference--group-001.md#canonical-bfc4842267fc635196a8710e5f991005dd73c95e97df089feebe338708b210cc)
- [timeouts](resources--container_registry--reference--group-001.md#canonical-9f566e8474c98272d7a45bdda4a03ebc7c8aaa0c8a664de216696b684c915244)
- [xcsh_container_registry](../resources/container_registry.md#canonical-d9354f0149bcec7d60dd61bf8c2faf0289b61a944b3928d453a60e87522be1fe)

<a id="canonical-bfc4842267fc635196a8710e5f991005dd73c95e97df089feebe338708b210cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8e74a58717cac22a1a0bbd3b080542833e980ca037057e0f2134235269dceee"></a>

## password — password / bcde27771627 / 2

Breadcrumbs:

- [xcsh_container_registry](../resources/container_registry.md#canonical-d9354f0149bcec7d60dd61bf8c2faf0289b61a944b3928d453a60e87522be1fe)
- [Property reference](resources--container_registry--reference--group-001.md#canonical-4e616e01aa1598e2fc141ba33f7c752da2e7a66e59191d4c102a8adb6557faa3)
- password

<a id="canonical-6c277754435430acc53bcfa5c0d782c33e843e2e49b310cc10378e03c55d4eaa"></a>

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

<a id="canonical-0b6a52ccd51d79be826ff53abf96f89ed3a8bf17f7fd1cb0e1c5c5e7329bf8c5"></a>

## Direct properties — password / bcde27771627 / 3

- [blindfold_secret_info](resources--container_registry--reference--group-001.md#canonical-91d6081e3a74d296bb98a06665d96748c9e5a38c74b80c858e8e46ef8d44d1a1): complete subsection reference.

- [clear_secret_info](resources--container_registry--reference--group-001.md#canonical-0b6d48d6eca065a0116260fd77e24eef9b7236026e3648d82acbe291d36d9ac3): complete subsection reference.

<a id="canonical-6f1bad33e55476063e0b31032a937770983ad877f4c96ff4932620a83e778cd0"></a>

## Next pages — password / bcde27771627 / 4

- [password.blindfold_secret_info](resources--container_registry--reference--group-001.md#canonical-91d6081e3a74d296bb98a06665d96748c9e5a38c74b80c858e8e46ef8d44d1a1)
- [password.clear_secret_info](resources--container_registry--reference--group-001.md#canonical-0b6d48d6eca065a0116260fd77e24eef9b7236026e3648d82acbe291d36d9ac3)
- [Property reference](resources--container_registry--reference--group-001.md#canonical-4e616e01aa1598e2fc141ba33f7c752da2e7a66e59191d4c102a8adb6557faa3)
- [xcsh_container_registry](../resources/container_registry.md#canonical-d9354f0149bcec7d60dd61bf8c2faf0289b61a944b3928d453a60e87522be1fe)

<a id="canonical-91d6081e3a74d296bb98a06665d96748c9e5a38c74b80c858e8e46ef8d44d1a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2205907bf3a4fe4b553e0f514705c3c9b836b8cd1feca4ae84a9960add1e985a"></a>

## password.blindfold_secret_info — password.blindfold_secret_info / eda78aedc06d / 2

Breadcrumbs:

- [xcsh_container_registry](../resources/container_registry.md#canonical-d9354f0149bcec7d60dd61bf8c2faf0289b61a944b3928d453a60e87522be1fe)
- [Property reference](resources--container_registry--reference--group-001.md#canonical-4e616e01aa1598e2fc141ba33f7c752da2e7a66e59191d4c102a8adb6557faa3)
- [password](resources--container_registry--reference--group-001.md#canonical-bfc4842267fc635196a8710e5f991005dd73c95e97df089feebe338708b210cc)
- password.blindfold_secret_info

<a id="canonical-c3a4daf92dafaa9096aa8543d5f67edb5ec2fe221da52c7c28f89792d29888bf"></a>

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

<a id="canonical-38cf6f6fb2d589af7ce857cd742ffed78573a31d0efb37a822a534b72472f5f5"></a>

## Direct properties — password.blindfold_secret_info / eda78aedc06d / 3

<a id="canonical-099d8ad9a054ff4a26a244293280b89fc643a13074727bf713e432fe963c329c"></a>

<a id="canonical-f11fa0151bfbba7769553109929d5030edc84d430952b7bd3e1342daeb1cef87"></a>

## decryption_provider property — password.blindfold_secret_info / eda78aedc06d / 4

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

<a id="canonical-e4930cb04e20c091c5db445ed38332dcb34d95f89b604ef262a41f092207b051"></a>

<a id="canonical-f54a43fdd03a7b6d005a60ff13a5e67cc4bb3e56955761ababb4bb07555231f6"></a>

## location property — password.blindfold_secret_info / eda78aedc06d / 5

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

<a id="canonical-b68187b7d94d172133dec3ac70470a2be205fa9426c977e0f1d7155270de4ffa"></a>

<a id="canonical-d017cc5c0a024510f08850961eec8031ce669b6da35f7f63b270bcfb802a635c"></a>

## store_provider property — password.blindfold_secret_info / eda78aedc06d / 6

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

<a id="canonical-3c0748041bd973cefe7b13a36641f6a6ae99c721f9b488fc9906f303f8b05a56"></a>

## Next pages — password.blindfold_secret_info / eda78aedc06d / 7

- [password](resources--container_registry--reference--group-001.md#canonical-bfc4842267fc635196a8710e5f991005dd73c95e97df089feebe338708b210cc)
- [xcsh_container_registry](../resources/container_registry.md#canonical-d9354f0149bcec7d60dd61bf8c2faf0289b61a944b3928d453a60e87522be1fe)

<a id="canonical-0b6d48d6eca065a0116260fd77e24eef9b7236026e3648d82acbe291d36d9ac3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4fb543c88e033756ea825abee292cac9597c6be4097d2c6d0dfb08f1b319c8a"></a>

## password.clear_secret_info — password.clear_secret_info / f15f3b20674e / 2

Breadcrumbs:

- [xcsh_container_registry](../resources/container_registry.md#canonical-d9354f0149bcec7d60dd61bf8c2faf0289b61a944b3928d453a60e87522be1fe)
- [Property reference](resources--container_registry--reference--group-001.md#canonical-4e616e01aa1598e2fc141ba33f7c752da2e7a66e59191d4c102a8adb6557faa3)
- [password](resources--container_registry--reference--group-001.md#canonical-bfc4842267fc635196a8710e5f991005dd73c95e97df089feebe338708b210cc)
- password.clear_secret_info

<a id="canonical-abc503d13d548c99be9146355f4005615fa74624fdbe42daf766e470e57b128d"></a>

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

<a id="canonical-8a84e48b5d664d20fa05fba007f64bd21d6892e38184f3ee19bb0bbb6e9dfcac"></a>

## Direct properties — password.clear_secret_info / f15f3b20674e / 3

<a id="canonical-3a45b45e7bc0818285affe13b08ab7f8825cf68786afaf0bbdc56dc1f6dc23b9"></a>

<a id="canonical-e382cdf2e5b658ace0ed11ef28a6c81ebf6b96f94212d4f1bdff01b112ac5cef"></a>

## provider_ref property — password.clear_secret_info / f15f3b20674e / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-98c7e0eb28b2eab40479f9cc9fca5730614d780ec01163dd65a7e7be05c3cea2"></a>

<a id="canonical-676d4a403741ff00d18f9fe221ec91114171998d12e1768e802fd74e4d5a7768"></a>

## url property — password.clear_secret_info / f15f3b20674e / 5

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

<a id="canonical-44c0c1db7a09d33e70ecba7a1d419ae54a6eb60d94e57c9d1d4e3292b4e6c558"></a>

## Next pages — password.clear_secret_info / f15f3b20674e / 6

- [password](resources--container_registry--reference--group-001.md#canonical-bfc4842267fc635196a8710e5f991005dd73c95e97df089feebe338708b210cc)
- [xcsh_container_registry](../resources/container_registry.md#canonical-d9354f0149bcec7d60dd61bf8c2faf0289b61a944b3928d453a60e87522be1fe)

<a id="canonical-9f566e8474c98272d7a45bdda4a03ebc7c8aaa0c8a664de216696b684c915244"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63bce1fd0f2ab6bb1c8cb7f8f99c0baa043222bd68f1d4d45db5b12a61fa0678"></a>

## timeouts — timeouts / 7157958c1dd5 / 2

Breadcrumbs:

- [xcsh_container_registry](../resources/container_registry.md#canonical-d9354f0149bcec7d60dd61bf8c2faf0289b61a944b3928d453a60e87522be1fe)
- [Property reference](resources--container_registry--reference--group-001.md#canonical-4e616e01aa1598e2fc141ba33f7c752da2e7a66e59191d4c102a8adb6557faa3)
- timeouts

<a id="canonical-8470bafa6128d257703d671043520873009717cb7e24de1af5aff682f1b53921"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-b14c681f0713266c6ae2aaef72296a0e10defa2e1e4d721891c791d7b9508417"></a>

## Direct properties — timeouts / 7157958c1dd5 / 3

<a id="canonical-c0ac067b91a9532f20bbde178afc3ba1486a66e14421c823b567fc57e02649c0"></a>

<a id="canonical-4295806a4d35029785d4971994b285e3d26600b62d9061be20bdf6d9275bf822"></a>

## create property — timeouts / 7157958c1dd5 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-d63d78a8899d16b4e5a94cb632cae3a0286d723494884da92721f04528772881"></a>

<a id="canonical-303c51044ba0d38d17069500ab28b39e518f47f163d0fd9feda97d0e288053c9"></a>

## delete property — timeouts / 7157958c1dd5 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-70a9a0f1500eeb83835d9b0ebe36bfd03ab8c3c0949d6edcb1d9b719ae56c4b4"></a>

<a id="canonical-39402fcde043892614ae063d99d6e8bb35425fe9b21b32ff6de04b1ef2e673d9"></a>

## read property — timeouts / 7157958c1dd5 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-25e2772f549af7c1bb429ab87390812cfc2f123529327dff0a1a101dbf295f97"></a>

<a id="canonical-6cac724549f813d67b7ffb72794d638a8fde1a627028f173f2b7d3c1add29c7c"></a>

## update property — timeouts / 7157958c1dd5 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-a052d2081525685ec67de06856706fb47ea42828714c2fb358312ee676b715e8"></a>

## Next pages — timeouts / 7157958c1dd5 / 8

- [Property reference](resources--container_registry--reference--group-001.md#canonical-4e616e01aa1598e2fc141ba33f7c752da2e7a66e59191d4c102a8adb6557faa3)
- [xcsh_container_registry](../resources/container_registry.md#canonical-d9354f0149bcec7d60dd61bf8c2faf0289b61a944b3928d453a60e87522be1fe)
