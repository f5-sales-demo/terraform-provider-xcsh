---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-c3b9637c5bb16d175a6363078fc42f964c382f3db77d569a7f35bb71aa5fe184"></a>

## custom_storage_config.storage_class_list.storage_classes — custom_storage_config.storage_class_list.storage_classes / dac3a9d51630 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_class_list](resources--voltstack_site--reference--group-005.md#canonical-e0ae1a1e35a34c03930df0f57bfcbedbc5f0465ea38c6de68c853d4c27620ace)
- custom_storage_config.storage_class_list.storage_classes

<a id="canonical-d2600de719c8d2c1e920f7b51577accc4b9f5fe9cd95c511f996f60e8466b4ee"></a>

Type: `"object"`. list nested block, Optional.

List of Storage Classes. List of custom storage classes.

Upstream description:

List of custom storage classes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("storage_class_name",
    "storage_device"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "hpe_storage"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "netapp_trident"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "pure_service_orchestrator"),
  validators.ConflictingListObjectAttributes("hpe_storage",
    "netapp_trident"),
  validators.ConflictingListObjectAttributes("hpe_storage",
    "pure_service_orchestrator"),
  validators.ConflictingListObjectAttributes("netapp_trident",
    "pure_service_orchestrator")}
```

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage_classes {
  # Configure direct properties listed below.
}
```

<a id="canonical-bbfd6d5e0541d0804a260db33aa8ababd4eb9f1c780d605cb3988dcf33cf595d"></a>

## Direct properties — custom_storage_config.storage_class_list.storage_classes / dac3a9d51630 / 3

<a id="canonical-6c3ceeb87d0e7e20f2df4a2b849ca569baa3545675dda2401346654211af9436"></a>

<a id="canonical-c43785442f37684df7ad2b59eb0a68ab7abad8246084f42b923c5463798811b0"></a>

## advanced_storage_parameters property — custom_storage_config.storage_class_list.storage_classes / dac3a9d51630 / 4

Type: `["map", "string"]`. Optional.

Advanced Parameters. Map of parameter name and string value.

Upstream description:

Map of parameter name and string value.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-0949e29ca53ee697f51faf613a3a277a42d7a5b966f2dc1b10e34e1848b49805"></a>

<a id="canonical-68a6972ba929dfa7f156d198a552dc75eb77d800cf9c22b23879d0e526ad1765"></a>

## allow_volume_expansion property — custom_storage_config.storage_class_list.storage_classes / dac3a9d51630 / 5

Type: `"bool"`. Optional.

Allow Volume Expansion. Allow volume expansion.

Upstream description:

Allow volume expansion.

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

- [custom_storage](resources--voltstack_site--reference--group-006.md#canonical-49fa8d42518f741859865d73f44435467cbb3328735d21ceba41913be7b39d18): complete subsection reference.

<a id="canonical-7c59335f829c4a941d20cf1b7b6d0de322427dd0d0941b31897a86e2e5aa1b01"></a>

<a id="canonical-6eb8819d0de80bc67b9ca80ebb72ed46d58bb5f2c5c0b42234c6ae9041d7bbc8"></a>

## default_storage_class property — custom_storage_config.storage_class_list.storage_classes / dac3a9d51630 / 6

Type: `"bool"`. Optional.

Make this storage class default storage class for the K8s cluster.

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

<a id="canonical-f72a742a4f049f6d0903f1ac53918594d803308b25e707c61aac7fa3dbb53e4d"></a>

<a id="canonical-18d9dfb362f86ffeb4d6edbf14e193b2ee9fdb47eef9c029cf844749dd25a514"></a>

## description_spec property — custom_storage_config.storage_class_list.storage_classes / dac3a9d51630 / 7

Type: `"string"`. Optional.

Storage Class Description. Description for this storage class.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-8e7b3706ed3f7d6083aa64aa424094d0d99a926b2649db3de6f8dbcef8e4eed2): complete subsection reference.

- [netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-c49ea2bdb7e038eedcc98a321caa6cbf93bbbceba45adfd875aa4394fbda86b0): complete subsection reference.

- [pure_service_orchestrator](resources--voltstack_site--reference--group-006.md#canonical-873d174188f816d931e178ee74bce7b30d699654304322dae0c7a69fb8884ee0): complete subsection reference.

<a id="canonical-bc03465e5f599f53c88254ca23581c4e3b73174267ec4d66a5760c6e1aee0d26"></a>

<a id="canonical-983a2d1c14611c3bf4d80b8ac2ca8fb94ae03cfb73aef2879e4f17454b1fce40"></a>

## reclaim_policy property — custom_storage_config.storage_class_list.storage_classes / dac3a9d51630 / 8

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Reclaim Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 16,
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
    "ves.io.schema.rules.string.max_len": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16"
  }
}
```

<a id="canonical-804d8bb8c14035eeced712d489c3d951d7f6045d19a0e9224f2905358505a2ec"></a>

<a id="canonical-09e66bcbe040a272f899ef45b6c1251c23d35821bd7964dee1c0495c60d0bced"></a>

## storage_class_name property — custom_storage_config.storage_class_list.storage_classes / dac3a9d51630 / 9

Type: `"string"`. Optional.

Name of the storage class as it will appear in K8s.

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
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-24a3639824b1a016bba8766038c28b86ac626b74def927447068777a6f46db0a"></a>

<a id="canonical-886f588c9b9d96a8080ed50561a9627e83cd4c6b445bf6c0dd5cd956e58da8a2"></a>

## storage_device property — custom_storage_config.storage_class_list.storage_classes / dac3a9d51630 / 10

Type: `"string"`. Optional.

Storage device that this class will use. The Device name defined at previous step.

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
    "byteLength": {
      "max": 64,
      "min": 1
    },
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-f4808f0590a4501f10c9b5012036e498cbe200b48f6844e9c95ffa2f862ee587"></a>

## Next pages — custom_storage_config.storage_class_list.storage_classes / dac3a9d51630 / 11

- [custom_storage_config.storage_class_list.storage_classes.custom_storage](resources--voltstack_site--reference--group-006.md#canonical-49fa8d42518f741859865d73f44435467cbb3328735d21ceba41913be7b39d18)
- [custom_storage_config.storage_class_list.storage_classes.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-8e7b3706ed3f7d6083aa64aa424094d0d99a926b2649db3de6f8dbcef8e4eed2)
- [custom_storage_config.storage_class_list.storage_classes.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-c49ea2bdb7e038eedcc98a321caa6cbf93bbbceba45adfd875aa4394fbda86b0)
- [custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrator](resources--voltstack_site--reference--group-006.md#canonical-873d174188f816d931e178ee74bce7b30d699654304322dae0c7a69fb8884ee0)
- [custom_storage_config.storage_class_list](resources--voltstack_site--reference--group-005.md#canonical-e0ae1a1e35a34c03930df0f57bfcbedbc5f0465ea38c6de68c853d4c27620ace)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-49fa8d42518f741859865d73f44435467cbb3328735d21ceba41913be7b39d18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd1e11fdf4244cb7d112655cd9a8250e61ad2f16407ee3767e58aebd3c5757be"></a>

## custom_storage_config.storage_class_list.storage_classes.custom_storage — custom_storage_config.storage_class_list.storage_classes.custom_storage / de3beba5fb55 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_class_list](resources--voltstack_site--reference--group-005.md#canonical-e0ae1a1e35a34c03930df0f57bfcbedbc5f0465ea38c6de68c853d4c27620ace)
- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-005.md#canonical-f0dd8ae1ed52d721e52531f94ab23974df15c76f232962d465a16dd35fee2c6f)
- custom_storage_config.storage_class_list.storage_classes.custom_storage

<a id="canonical-af3a1b8be3992143e743b1857dcd95f49a0987fa2b48ece1850f648da673cb99"></a>

Type: `"object"`. single nested block, Optional.

Custom Storage Class allows to insert Kubernetes storageclass definition which will be applied into
given site.

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
custom_storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-a26ceb0cb6a7477769019ca8f0343f9a6acc2a9766e2ee1f20bb68d1f74bf71a"></a>

## Direct properties — custom_storage_config.storage_class_list.storage_classes.custom_storage / de3beba5fb55 / 3

<a id="canonical-9a936daf1779cde7a4938ef009de7fa25aa55b4eb2245b1df69a31e81a827376"></a>

<a id="canonical-5c1955c495f400d40db2fe464e6e8115261d33533c02ab2920859282c8aef93a"></a>

## yaml property — custom_storage_config.storage_class_list.storage_classes.custom_storage / de3beba5fb55 / 4

Type: `"string"`. Optional.

Storage Class YAML. K8s YAML for StorageClass.

Upstream description:

K8s YAML for StorageClass.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "Valid parseable YAML",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "validation": {
      "customRule": "Must be valid YAML"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-39e826898734a34b8e3a8b01b9cfc8441798513fb633f5d8e4eada2a03f428f1"></a>

## Next pages — custom_storage_config.storage_class_list.storage_classes.custom_storage / de3beba5fb55 / 5

- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-005.md#canonical-f0dd8ae1ed52d721e52531f94ab23974df15c76f232962d465a16dd35fee2c6f)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-8e7b3706ed3f7d6083aa64aa424094d0d99a926b2649db3de6f8dbcef8e4eed2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e5e8db11301991516025a55506d08ebfaeda2e6ed3af7e0e938b2f10f05a848b"></a>

## custom_storage_config.storage_class_list.storage_classes.hpe_storage — custom_storage_config.storage_class_list.storage_classes.hpe_storage / a36e6efdee2c / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_class_list](resources--voltstack_site--reference--group-005.md#canonical-e0ae1a1e35a34c03930df0f57bfcbedbc5f0465ea38c6de68c853d4c27620ace)
- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-005.md#canonical-f0dd8ae1ed52d721e52531f94ab23974df15c76f232962d465a16dd35fee2c6f)
- custom_storage_config.storage_class_list.storage_classes.hpe_storage

<a id="canonical-4b8b83bd99d59518378df94652cdea1bf209626704440ba47560cf978af4467c"></a>

Type: `"object"`. single nested block, Optional.

Storage class Device configuration for HPE Storage.

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
hpe_storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-c74cd17447636fc0e6f975a8db7da0183d53783cb82d1095cc33bf4d5ffaed22"></a>

## Direct properties — custom_storage_config.storage_class_list.storage_classes.hpe_storage / a36e6efdee2c / 3

<a id="canonical-67953b7895a473b1a776229c51a78d55b4fb5fc5a3fbf695bae37f832defb6c3"></a>

<a id="canonical-da6ba653c6a7296cdc81fde859b130ca58b335fbf8515c35cd829c30a836bb74"></a>

## allow_mutations property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / a36e6efdee2c / 4

Type: `"string"`. Optional.

Mutation can override specified parameters.

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

<a id="canonical-51f9098f070ea0d3aecbe45fc4acb53df6a7fca413aa36661fce4173ff8cdc7b"></a>

<a id="canonical-86e7d4ffb72be2acdeee8026d53995b7c3cb97f9fb578794262b960b86eab2d2"></a>

## allow_overrides property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / a36e6efdee2c / 5

Type: `"string"`. Optional.

AllowOverrides. PVC can override specified parameters.

Upstream description:

PVC can override specified parameters.

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

<a id="canonical-5a03f603aeb1437c924b54a4d1874243ac5c6c7211090c8c553fa58e773e674d"></a>

<a id="canonical-12b87c02494af399340c0a78c95c8170dc4b576d393945b73eaf9342011ea7d0"></a>

## dedupe_enabled property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / a36e6efdee2c / 6

Type: `"bool"`. Optional.

Indicates that the volume should enable deduplication.

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

<a id="canonical-48330b8b572bdecfafda570f7665955b14b6724bce1b198246924d28b2ddf3a7"></a>

<a id="canonical-195a4133e9658924f3a76a69211fe5c875cd11f90b9bfb01103c7be47d3a1175"></a>

## description_spec property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / a36e6efdee2c / 7

Type: `"string"`. Optional.

The SecretName parameter is used to identify name of secret to identify backend storage's auth
information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

<a id="canonical-d6382d0fcf139a32f1954343cd3b6b660d9322876972582ec801460905d1e155"></a>

<a id="canonical-4de9265fb3f4a833a0fe9b82bfda79167a763fad482526ec954628332360bf97"></a>

## destroy_on_delete property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / a36e6efdee2c / 8

Type: `"bool"`. Optional.

Indicates the backing Nimble volume (including snapshots) should be destroyed when the PVC is
deleted.

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

<a id="canonical-d74b90367a375fe305a7a28a2c776c5a1b07519645afa4fe3abb277bb3a6b707"></a>

<a id="canonical-723fa397ca884c3f18748631ace55cb3aa3b3672da4cdbe4c9cda0864ceea7ee"></a>

## encrypted property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / a36e6efdee2c / 9

Type: `"bool"`. Optional.

Indicates that the volume should be encrypted.

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

<a id="canonical-584ef68d3b10589ee62f25b47ae93ee23ef6950585909dbfb6051199f0eb052e"></a>

<a id="canonical-d414ef79f85edf8e5bb88c011e785dbf91480d25a0de5f07873cce22c4a01d9e"></a>

## folder property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / a36e6efdee2c / 10

Type: `"string"`. Optional.

The name of the folder in which to place the volume.

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

<a id="canonical-1f26fef7a41171e15ba5b7fe83c263108438cf9df1f6768c00f7ff6336e17bf7"></a>

<a id="canonical-e9da3c99c524fb3c8c5c4a652e9914fde9fd978129c806743bdcaf49c0c5df8d"></a>

## limit_iops property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / a36e6efdee2c / 11

Type: `"string"`. Optional.

LimitIops. The IOPS limit of the volume.

Upstream description:

The IOPS limit of the volume.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "int64",
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

<a id="canonical-6a4a5d8aec2b5f99fa9548a57b8afc6dd8d641d8c31e9085c754c5a255f9e713"></a>

<a id="canonical-cde23c8630b11e5383ebbcc771b1910975f74691cb9fb11db1567b2f631edb39"></a>

## limit_mbps property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / a36e6efdee2c / 12

Type: `"string"`. Optional.

LimitMbps. The IOPS limit of the volume.

Upstream description:

The IOPS limit of the volume.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "int64",
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

<a id="canonical-bf88128d26aa5add61b8ae42b5cce04939d72f60946bd8ff4a31d583d2f6417d"></a>

<a id="canonical-01a6851dcfc515f5c9e0142c304bb15ff10ae66a31731562c197f122e8905ad8"></a>

## performance_policy property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / a36e6efdee2c / 13

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

The name of the performance policy to assign to the volume.

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

<a id="canonical-33c051f13f1ea023672ccca0e57624248afdbf4ba35310c8e0e60f18b1e6e295"></a>

<a id="canonical-b8361692d1a1a40c75aef26f8a1257045722168b423d5f2378000a2a98b1ada9"></a>

## pool property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / a36e6efdee2c / 14

Type: `"string"`. Optional.

The name of the pool in which to place the volume.

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

<a id="canonical-b4677377fa874b18382272c9ed360ad24105c45a614f4f72fda11201b6071c4c"></a>

<a id="canonical-de957363345478e8c39bf6e0c0053f3fea8e2026c1fec88d1dc47a21d13e8c14"></a>

## protection_template property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / a36e6efdee2c / 15

Type: `"string"`. Optional.

The name of the performance policy to assign to the volume.

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

<a id="canonical-c51fe0173deef3eadf1cc22f63d24d99f40d98db93c5502a7904f6b583cbf310"></a>

<a id="canonical-0d90f9c7226c22e0f611e91c6dcead2fa590239041f1032d7c6198466da867d5"></a>

## secret_name property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / a36e6efdee2c / 16

Type: `"string"`. Optional.

The SecretName parameter is used to identify name of secret to identify backend storage's auth
information.

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

<a id="canonical-87339e550c377e7786d0f9b2504372af5a98eb7e9880ef8f78eefcf18666276a"></a>

<a id="canonical-9eac2b62dfa84d725bd1f8556a74b57c620591f1e10ae8ac3d517b5a46f9e434"></a>

## secret_namespace property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / a36e6efdee2c / 17

Type: `"string"`. Optional.

The SecretNamespace parameter is used to identify name of namespace where secret resides.

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

<a id="canonical-7b9338bfb634a712505599aec1f88fdd701ff1293239c16596e78cdfdbabfedd"></a>

<a id="canonical-de12c9eeacd642e61918fe2a8398a353f73cbc00a43ca155019be018b5d65942"></a>

## sync_on_detach property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / a36e6efdee2c / 18

Type: `"bool"`. Optional.

Indicates that a snapshot of the volume should be synced to the replication partner each time it is
detached from a node.

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

<a id="canonical-3c2151ab373684397c35ef066d2bd843dafd2dc45c2033531d52a3f8625a038a"></a>

<a id="canonical-d32524735801be328ad7ed59444d37c2af8fb81e9d18fced4e7eaf06e9a9628e"></a>

## thick property — custom_storage_config.storage_class_list.storage_classes.hpe_storage / a36e6efdee2c / 19

Type: `"bool"`. Optional.

Indicates that the volume should be thick provisioned.

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

<a id="canonical-3d25648f9a1d81ad1365de06510ad3b018e828bd5f0235d4ad989cb573aff666"></a>

## Next pages — custom_storage_config.storage_class_list.storage_classes.hpe_storage / a36e6efdee2c / 20

- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-005.md#canonical-f0dd8ae1ed52d721e52531f94ab23974df15c76f232962d465a16dd35fee2c6f)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-c49ea2bdb7e038eedcc98a321caa6cbf93bbbceba45adfd875aa4394fbda86b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04c3a01f293e920e233dcc6c39ae8213464219a6f375221e69337143a5665049"></a>

## custom_storage_config.storage_class_list.storage_classes.netapp_trident — custom_storage_config.storage_class_list.storage_classes.netapp_trident / bf1be5e899b7 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_class_list](resources--voltstack_site--reference--group-005.md#canonical-e0ae1a1e35a34c03930df0f57bfcbedbc5f0465ea38c6de68c853d4c27620ace)
- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-005.md#canonical-f0dd8ae1ed52d721e52531f94ab23974df15c76f232962d465a16dd35fee2c6f)
- custom_storage_config.storage_class_list.storage_classes.netapp_trident

<a id="canonical-fa7e163f3c73e69f2af6855ab0cfb897b1d6a8b4575124c95076e3084ee71842"></a>

Type: `"object"`. single nested block, Optional.

Storage class Device configuration for NetApp Trident.

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
netapp_trident {
  # Configure direct properties listed below.
}
```

<a id="canonical-07de3fa92efeb46a2720ea7967cdf72c29f9c0798c57850019fbb13bcddb9e33"></a>

## Direct properties — custom_storage_config.storage_class_list.storage_classes.netapp_trident / bf1be5e899b7 / 3

- [selector](resources--voltstack_site--reference--group-006.md#canonical-5ba1e5b7879bad244270d3bf6af8f4fa0068cad5a17f297484e0edff43bced40): complete subsection reference.

<a id="canonical-96c726b403f79722a00f4d30f188a6dfdee5b7fe6401a08d7aab5151aff7a4fc"></a>

<a id="canonical-2e5140719deb0bfcb6d2a8f9c22e4bd63a8692e3477d19be9de6899440583e40"></a>

## storage_pools property — custom_storage_config.storage_class_list.storage_classes.netapp_trident / bf1be5e899b7 / 4

Type: `"string"`. Optional.

The storagePools parameter is used to further restrict the set of pools that match any specified
attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

<a id="canonical-1d99304ed2e963d90583574d251addfbd45c0d6b891133f3a86761f6ce0dbf81"></a>

## Next pages — custom_storage_config.storage_class_list.storage_classes.netapp_trident / bf1be5e899b7 / 5

- [custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector](resources--voltstack_site--reference--group-006.md#canonical-5ba1e5b7879bad244270d3bf6af8f4fa0068cad5a17f297484e0edff43bced40)
- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-005.md#canonical-f0dd8ae1ed52d721e52531f94ab23974df15c76f232962d465a16dd35fee2c6f)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-5ba1e5b7879bad244270d3bf6af8f4fa0068cad5a17f297484e0edff43bced40"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b401250283ba7ea2892dd6da8587ec326f2c15d70cb53c9d907ba22c3cdf99a"></a>

## custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector — custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector / 10c6eba339ee / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_class_list](resources--voltstack_site--reference--group-005.md#canonical-e0ae1a1e35a34c03930df0f57bfcbedbc5f0465ea38c6de68c853d4c27620ace)
- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-005.md#canonical-f0dd8ae1ed52d721e52531f94ab23974df15c76f232962d465a16dd35fee2c6f)
- [custom_storage_config.storage_class_list.storage_classes.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-c49ea2bdb7e038eedcc98a321caa6cbf93bbbceba45adfd875aa4394fbda86b0)
- custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector

<a id="canonical-171f73da868ee52b57b2974545caa21db1cb23481d5172d6a172a3f71d4108f4"></a>

Type: `"object"`. single nested block, Optional.

Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a
volume. The volume will have the aspects defined in the chosen virtual pool.

Upstream description:

Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a
volume. The volume will have the aspects defined in the chosen virtual pool.

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
selector {}
```

<a id="canonical-03af8e845901ca932e143160849cd2e6b30587fbd4b3c8e43bcea09f18337036"></a>

## Direct properties — custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector / 10c6eba339ee / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1152b4af328df0503ece21e4f253dbf28b6780ffd512568a9babd5082fb13390"></a>

## Next pages — custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector / 10c6eba339ee / 4

- [custom_storage_config.storage_class_list.storage_classes.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-c49ea2bdb7e038eedcc98a321caa6cbf93bbbceba45adfd875aa4394fbda86b0)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-873d174188f816d931e178ee74bce7b30d699654304322dae0c7a69fb8884ee0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1efcf1393223d2b52a5c0d41b7841cdc6c7cd111091d274d6433a104b06f595e"></a>

## custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrator — custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrat / ed88169188ec / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_class_list](resources--voltstack_site--reference--group-005.md#canonical-e0ae1a1e35a34c03930df0f57bfcbedbc5f0465ea38c6de68c853d4c27620ace)
- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-005.md#canonical-f0dd8ae1ed52d721e52531f94ab23974df15c76f232962d465a16dd35fee2c6f)
- custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrator

<a id="canonical-e3bece06a0c1b37db157b6f35a17bf98abd99321b486291819439e79d547c83a"></a>

Type: `"object"`. single nested block, Optional.

Storage class Device configuration for Pure Service Orchestrator.

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
pure_service_orchestrator {
  # Configure direct properties listed below.
}
```

<a id="canonical-1a00c3790a5d90862af36df63a7ac4fc9f210d540bb738e5f9ab8c69bbd3f5c2"></a>

## Direct properties — custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrat / ed88169188ec / 3

<a id="canonical-8059c078ea871f9237ca8b32dc4a8aaa2db043c4dd614c60411ce730f55d6d6d"></a>

<a id="canonical-ed75470dcef4f2817ad7fc25160013f4c01036dc1292d22ae28e110d17594d17"></a>

## backend property — custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrat / ed88169188ec / 4

Type: `"string"`. Optional.

\[Enum: block|file\] Defines type of Pure storage backend block or file. The volume will have the
aspects defined in the chosen virtual pool. Possible values are \`block\`, \`file\`.

Upstream description:

Defines type of Pure storage backend block or file. The volume will have the aspects defined in the
chosen virtual pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("block",
    "file"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "block",
    "file"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"block\\\",\\\"file\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"block\\\",\\\"file\\\"]"
  }
}
```

<a id="canonical-e6b4d58280c4541c87cc8d35064e3df8c9ebb067cc9e9f9d4de0046e86f35a9d"></a>

<a id="canonical-5f99b6686dd1c3f7d277e73e8079d7fcea2bb04bfac9b3e8ecc1b3ab199ee1d7"></a>

## bandwidth_limit property — custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrat / ed88169188ec / 5

Type: `"string"`. Optional.

It must be between 1 MB/s and 512 GB/s. Enter the size as a number (bytes must be multiple of 512)
or number with a single character unit symbol. Valid unit symbols are K, M, G, representing KiB,
MiB, and GiB.

Upstream description:

It must be between 1 MB/s and 512 GB/s. Enter the size as a number (bytes must be multiple of 512)
or number with a single character unit symbol. Valid unit symbols are K, M, G, representing KiB,
MiB, and GiB.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(12),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 12,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 12,
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
    "ves.io.schema.rules.string.max_len": "12"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "12"
  }
}
```

<a id="canonical-1f844f9bd85742edb9bda754a48dcf7fce5fc9dad81a329c1b790cb909f8b3ca"></a>

<a id="canonical-6c8b1d8827bab495ebb43872d19f0ff9bc44b2d3121b07decb2081c4fd081e49"></a>

## iops_limit property — custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrat / ed88169188ec / 6

Type: `"number"`. Optional.

Enable IOPS limitation. It must be between 100 and 100 million. If value is 0, IOPS limit is not
defined.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 100, Maximum: 100000000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100000000,
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
    "ves.io.schema.rules.uint32.ranges": "0,100-100000000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,100-100000000"
  }
}
```

<a id="canonical-b81395bfa0a1771b0c9ec9c644a3c65236590d910f56adc988248ee92c9102b0"></a>

## Next pages — custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrat / ed88169188ec / 7

- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--reference--group-005.md#canonical-f0dd8ae1ed52d721e52531f94ab23974df15c76f232962d465a16dd35fee2c6f)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-071b1e9cecd2e144eea010b705f1a3726acd7ec6dc044e04001b1620f4af4ec2"></a>

## custom_storage_config.storage_device_list — custom_storage_config.storage_device_list / e19c27119185 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- custom_storage_config.storage_device_list

<a id="canonical-1e90697a9aa62ca34b72288bd7c9c6d44da2c15f595affdec12997716fd08539"></a>

Type: `"object"`. single nested block, Optional.

Add additional custom storage classes in Kubernetes for this fleet.

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
storage_device_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-a723ac1440f267ee1bc711afbf1e5248c74cea66ff6f5b6e456d041dce7e8186"></a>

## Direct properties — custom_storage_config.storage_device_list / e19c27119185 / 3

- [storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb): complete subsection reference.

<a id="canonical-e000a062d2e9406e648c803ac615534b95bef4cd52655224e995754ff0d216c9"></a>

## Next pages — custom_storage_config.storage_device_list / e19c27119185 / 4

- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b1bc4675050817d84d75ca67667118d4a60011de41ea6ba107b2b68e14d36c3"></a>

## custom_storage_config.storage_device_list.storage_devices — custom_storage_config.storage_device_list.storage_devices / f16e1a7ae43d / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- custom_storage_config.storage_device_list.storage_devices

<a id="canonical-836b6c28dce376fb1d2da185f8db78caa1cbac6eef53fb1f1c04b1f4867f4c34"></a>

Type: `"object"`. list nested block, Optional.

List of Storage Devices. List of custom storage devices.

Upstream description:

List of custom storage devices.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("storage_device"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "hpe_storage"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "netapp_trident"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "pure_service_orchestrator"),
  validators.ConflictingListObjectAttributes("hpe_storage",
    "netapp_trident"),
  validators.ConflictingListObjectAttributes("hpe_storage",
    "pure_service_orchestrator"),
  validators.ConflictingListObjectAttributes("netapp_trident",
    "pure_service_orchestrator")}
```

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage_devices {
  # Configure direct properties listed below.
}
```

<a id="canonical-e90c9af756b50af9207940b68007ea3d0b68a3d29d0408c06a4f06a82f953265"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices / f16e1a7ae43d / 3

<a id="canonical-6cdcb41d19255378982ffa3e4ad0634c5f9db742211ece8c89785bf377742217"></a>

<a id="canonical-e115f123a3077105d574e267f3e4662e8933b61c3b5bd2d073121dd5d13c3f3d"></a>

## advanced_advanced_parameters property — custom_storage_config.storage_device_list.storage_devices / f16e1a7ae43d / 4

Type: `["map", "string"]`. Optional.

Advanced Parameters. Map of parameter name and string value.

Upstream description:

Map of parameter name and string value.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [custom_storage](resources--voltstack_site--reference--group-006.md#canonical-fa072cab043524dcffe82eeec4adac9ee9c78e534c9252f64d349543cd19b3b1): complete subsection reference.

- [hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-408b49d4b3038d0233a1a15d15c44f8a25a6212c7840568f58f066e6961127e1): complete subsection reference.

- [netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b): complete subsection reference.

- [pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-de9ccfeb4756e83ff6cc51ab35fb939749876ea24008d3dcd033ebad9d591290): complete subsection reference.

<a id="canonical-ebecbddbecf7a04a053bef62f70f15d922656975857a98a17473cdd557bde46a"></a>

<a id="canonical-7b21f70d1f1c8aeb95ba715fcf9f1001c316c7af968216d5da71e05fd193c338"></a>

## storage_device property — custom_storage_config.storage_device_list.storage_devices / f16e1a7ae43d / 5

Type: `"string"`. Optional.

Storage Device. Storage device and device unit.

Upstream description:

Storage device and device unit.

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

<a id="canonical-958aae2d97494c445483f10932a53ba368ed513517a07f411901626071376256"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices / f16e1a7ae43d / 6

- [custom_storage_config.storage_device_list.storage_devices.custom_storage](resources--voltstack_site--reference--group-006.md#canonical-fa072cab043524dcffe82eeec4adac9ee9c78e534c9252f64d349543cd19b3b1)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-408b49d4b3038d0233a1a15d15c44f8a25a6212c7840568f58f066e6961127e1)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-de9ccfeb4756e83ff6cc51ab35fb939749876ea24008d3dcd033ebad9d591290)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-fa072cab043524dcffe82eeec4adac9ee9c78e534c9252f64d349543cd19b3b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0876d77c85568cbebf4f25a1ec67dde6422cddae78a254b99fa9ee8703c7c0d"></a>

## custom_storage_config.storage_device_list.storage_devices.custom_storage — custom_storage_config.storage_device_list.storage_devices.custom_storage / ec49427c2fe7 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- custom_storage_config.storage_device_list.storage_devices.custom_storage

<a id="canonical-70786fb27fe2727c850dcfcf3a9c6fe78954be19c535fe592f6d3bb864426005"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for custom storage.

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
custom_storage = {}
```

<a id="canonical-38389f0cdf133cfcd7526928796ed2089aaaa168d05a80c63ef6b4f65a0e346f"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.custom_storage / ec49427c2fe7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-25c4ca447b887092540c118ca48e2b6b70227e9cfc1dafe0f9c6221873ed8598"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.custom_storage / ec49427c2fe7 / 4

- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-408b49d4b3038d0233a1a15d15c44f8a25a6212c7840568f58f066e6961127e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09ff93c4a8c3d4b77e86353785baf6c84401118deea9e15b7afbf71d17e8adaf"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage — custom_storage_config.storage_device_list.storage_devices.hpe_storage / e44834352a2f / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage

<a id="canonical-3130564e48126ca5f5e495815938c51c1902bf46fd7a1d757852c5b1021351ef"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for hpe storage.

Upstream description:

Device configuration for HPE Storage.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("api_server_port",
    "username")}
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
hpe_storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-7f8b99548828fa478d8464f8f0cf321aff4085e528fea69b71b8dae050f10717"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.hpe_storage / e44834352a2f / 3

<a id="canonical-ebf9f3d4b66fbbcbcfa3127ca1bb42e71c35eb87d3e28c2d80ef47bb520474cb"></a>

<a id="canonical-c686b8b4f58da945797ac8a6ce22efa9199fc21438d8ca3552bc773ac1ea1330"></a>

## api_server_port property — custom_storage_config.storage_device_list.storage_devices.hpe_storage / e44834352a2f / 4

Type: `"number"`. Optional.

Storage server Port. Enter Storage Server Port.

Upstream description:

Enter Storage Server Port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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

- [iscsi_chap_password](resources--voltstack_site--reference--group-006.md#canonical-c22da9c161eede6f7df4ed9716cc207425a7cd9b4f3ecfc0fc1c42d4b8a37e28): complete subsection reference.

<a id="canonical-a749b8dc58912f88ee335298cf7e878c50683148cee989448d4222ad50c6fea7"></a>

<a id="canonical-7632b086278ef986ddac0574b404e46e0505a7fb7c6cd4a58aebae8288330c42"></a>

## iscsi_chap_user property — custom_storage_config.storage_device_list.storage_devices.hpe_storage / e44834352a2f / 5

Type: `"string"`. Optional.

Chap Username to connect to the HPE storage.

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

- [password](resources--voltstack_site--reference--group-006.md#canonical-0e8993c197e6fef63651797b3663aa6f5783a2b4ddad77fa078a75d1888ec801): complete subsection reference.

<a id="canonical-305c4c3ca73ee1f43286e64f3ece03ad7d3ae62eb7be20e939679e071fd462ca"></a>

<a id="canonical-d9bcb109bd9bdde47228649a296b0f8a99a7319f5724e545c0493330106774b8"></a>

## storage_server_ip_address property — custom_storage_config.storage_device_list.storage_devices.hpe_storage / e44834352a2f / 6

Type: `"string"`. Optional.

Storage Server IP address. Enter storage server IP address.

Upstream description:

Enter storage server IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-ab44fee02467cf824f07ad5924f0e0f5f8bcc2cf611efc7935d5f7758fb5ce22"></a>

<a id="canonical-b34f6e6e0e7e7ed35861d93cc0ca1eabb9882896341f7c6b23cdf6e4b75ed3d9"></a>

## storage_server_name property — custom_storage_config.storage_device_list.storage_devices.hpe_storage / e44834352a2f / 7

Type: `"string"`. Optional.

Storage Server Name. Enter storage server Name.

Upstream description:

Enter storage server Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-4416c54de5be47d15a4a70e69e6c1b94e04b569521e1325057a6d9c59090516f"></a>

<a id="canonical-aea36ffeb566c7bd9d5ce975681d16dd91f600185aa0624e0f40fd92207bcd1f"></a>

## username property — custom_storage_config.storage_device_list.storage_devices.hpe_storage / e44834352a2f / 8

Type: `"string"`. Optional.

Username to connect to the HPE storage management IP.

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
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-31cab38fefb8849d194136fb09ab6849550de8c21bd96faef3fa0697f9dc2de8"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.hpe_storage / e44834352a2f / 9

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--voltstack_site--reference--group-006.md#canonical-c22da9c161eede6f7df4ed9716cc207425a7cd9b4f3ecfc0fc1c42d4b8a37e28)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](resources--voltstack_site--reference--group-006.md#canonical-0e8993c197e6fef63651797b3663aa6f5783a2b4ddad77fa078a75d1888ec801)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-c22da9c161eede6f7df4ed9716cc207425a7cd9b4f3ecfc0fc1c42d4b8a37e28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a5a200d6e253b75a04244f12584ae2f082af715e8e5c02d2c52b619d7a27e53"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / 3c3caf9be65a / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-408b49d4b3038d0233a1a15d15c44f8a25a6212c7840568f58f066e6961127e1)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password

<a id="canonical-40f7ccc789b28fd115ae32997ec94488268761de54e166b45e34544e7d49b708"></a>

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
iscsi_chap_password {
  # Configure direct properties listed below.
}
```

<a id="canonical-dc689dd8dfb2eba8878be2746116270056418e5339565991f5cc22490e0427e7"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / 3c3caf9be65a / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-006.md#canonical-87049b5c6686c5626de7a3059da0fe3f9e5118c9105cecf7acbdce910a91b981): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-006.md#canonical-3f534f2db89141b4c0d44c7550338f2a954144da47af836e634254d6cf8e568a): complete subsection reference.

<a id="canonical-a3985b9bc80cf7cd2e3a26d45f9386169a340df3e69ca950cda3ca355072e387"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / 3c3caf9be65a / 4

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info](resources--voltstack_site--reference--group-006.md#canonical-87049b5c6686c5626de7a3059da0fe3f9e5118c9105cecf7acbdce910a91b981)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info](resources--voltstack_site--reference--group-006.md#canonical-3f534f2db89141b4c0d44c7550338f2a954144da47af836e634254d6cf8e568a)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-408b49d4b3038d0233a1a15d15c44f8a25a6212c7840568f58f066e6961127e1)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-87049b5c6686c5626de7a3059da0fe3f9e5118c9105cecf7acbdce910a91b981"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5aa8d270fa1cb64b917fedd443e1f20bd8052a2860f6b4c720d6c077ff53e549"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / a0bad2070f21 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-408b49d4b3038d0233a1a15d15c44f8a25a6212c7840568f58f066e6961127e1)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--voltstack_site--reference--group-006.md#canonical-c22da9c161eede6f7df4ed9716cc207425a7cd9b4f3ecfc0fc1c42d4b8a37e28)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info

<a id="canonical-c65ea7374d125be7bfd2c6b8c80a1c74532faa73a2a3712c9e174318a49967e1"></a>

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

<a id="canonical-a14a6a71de76a21f7bb0e25ebfeb6233fdaf3451449e70e81e0ae279cb20b079"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / a0bad2070f21 / 3

<a id="canonical-e75f00e57001af6f14f474a5843c9d90099820b22a089ed1585475cd509b9363"></a>

<a id="canonical-b988a53f85404cffdc914f3447c9f98348c44881b5ee2319a6d3e8ede3a990b9"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / a0bad2070f21 / 4

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

<a id="canonical-5de8e5b741f0c0f27df300e2e6b8c0a6f5c3a6a8e302928aefc1b0ed9a676049"></a>

<a id="canonical-76e84d71fbf292c25133fa7d31fd8fe755e12768bedc43ff6f1b703298cfb63e"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / a0bad2070f21 / 5

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

<a id="canonical-a5899ffd625665aad63ef19551b5d3f2a8d2c4f3b2622318d4da9921f97ef900"></a>

<a id="canonical-4c8e4223c260942b30aaece5badf91f33639fb3a0a374f8ac6c233085be24c97"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / a0bad2070f21 / 6

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

<a id="canonical-90fb1b54c481884b54fc4265112c34c9be773190698d5806d10c1166e057d468"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / a0bad2070f21 / 7

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--voltstack_site--reference--group-006.md#canonical-c22da9c161eede6f7df4ed9716cc207425a7cd9b4f3ecfc0fc1c42d4b8a37e28)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-3f534f2db89141b4c0d44c7550338f2a954144da47af836e634254d6cf8e568a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8438d702be9474a9365dd485802153676392604da5dd18f407f8c9b8d9417a8"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / 2accc2cf0f96 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-408b49d4b3038d0233a1a15d15c44f8a25a6212c7840568f58f066e6961127e1)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--voltstack_site--reference--group-006.md#canonical-c22da9c161eede6f7df4ed9716cc207425a7cd9b4f3ecfc0fc1c42d4b8a37e28)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info

<a id="canonical-e29e989500b2dff98af84f3a5a64862515507b00a3f988c48f00b574823e5f6b"></a>

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

<a id="canonical-081bcbe724e98a22f39be7e37325e658b8efe06d92b225229a0ea1cc6679355e"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / 2accc2cf0f96 / 3

<a id="canonical-c46b8c08a5a3b917e36eeedb5281c690824dcf7da6970afe03dea29368f2b187"></a>

<a id="canonical-da4a73fd4f9d10ab1fac0e1a013f53b925772b5ce425a4a93ee4dfee5457c268"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / 2accc2cf0f96 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-7b64b00457e4c5b382a6ab0c716a832f73a0223cd609e6dafd73f37463c02831"></a>

<a id="canonical-e45018523e5b1c995eefafd19e043fa7999f01265f463b41bf5f6304ad33707c"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / 2accc2cf0f96 / 5

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

<a id="canonical-28930cb22839ed31b3a0c5a4bee1e4c202c6d21b51db72781078134b10ef6822"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap / 2accc2cf0f96 / 6

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](resources--voltstack_site--reference--group-006.md#canonical-c22da9c161eede6f7df4ed9716cc207425a7cd9b4f3ecfc0fc1c42d4b8a37e28)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-0e8993c197e6fef63651797b3663aa6f5783a2b4ddad77fa078a75d1888ec801"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95a28dbb81eb39c7b303909dc83f8fbad3ca0c77a4656279c6956a34262b31e5"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.password — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password / ba96e9f81766 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-408b49d4b3038d0233a1a15d15c44f8a25a6212c7840568f58f066e6961127e1)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.password

<a id="canonical-65efb30752a33d9b5dd42b6f0d8b49a948a274eb075220dbd997d99bd540ba03"></a>

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

<a id="canonical-affed565a420775ae0d483c34f76b0a96975f9f77bb2cb200c576028c887b83a"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password / ba96e9f81766 / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-006.md#canonical-e4308967205be3a352a114e693b41f20cfe0f08e326bf1fca82a7726b228a5dc): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-006.md#canonical-312730f0207f41804051d362d25a2503c5d2fbb0a95125d4a4235e711a9d1d6f): complete subsection reference.

<a id="canonical-59aa37c2d17b7a5611e3afb60763568608ac08becd5a6f88aec35d5e35ef2ed0"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password / ba96e9f81766 / 4

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info](resources--voltstack_site--reference--group-006.md#canonical-e4308967205be3a352a114e693b41f20cfe0f08e326bf1fca82a7726b228a5dc)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.clear_secret_info](resources--voltstack_site--reference--group-006.md#canonical-312730f0207f41804051d362d25a2503c5d2fbb0a95125d4a4235e711a9d1d6f)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-408b49d4b3038d0233a1a15d15c44f8a25a6212c7840568f58f066e6961127e1)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-e4308967205be3a352a114e693b41f20cfe0f08e326bf1fca82a7726b228a5dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16080a8ffff87cc86a8187edbac08427d03dc77f94905e49106fabce5f86c7ca"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.b / 38087516db13 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-408b49d4b3038d0233a1a15d15c44f8a25a6212c7840568f58f066e6961127e1)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](resources--voltstack_site--reference--group-006.md#canonical-0e8993c197e6fef63651797b3663aa6f5783a2b4ddad77fa078a75d1888ec801)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info

<a id="canonical-53b70506f5ccb4dca5c248871ee232b7b4878024940973ef83a5fe9b99562019"></a>

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

<a id="canonical-5eed046ffb56ffa0fab3be143cca570e7897300324cd4053917ee5ab57aaef9c"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.b / 38087516db13 / 3

<a id="canonical-ba8a81756125cb9c323724961a7608a8d7dfe377ab644e164abc8fa8abaeed1e"></a>

<a id="canonical-699788ac9e88aa55e818ab09b7f8bd5911aae0263c7e658cfff55e37fa2dac42"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.b / 38087516db13 / 4

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

<a id="canonical-c5c880d9ac51b619ccb37025fa51ccd4337794c1cf3caa19b8e949e1b8440cff"></a>

<a id="canonical-4e85af291923e6578167dbf620f498091fd5f7b08e811026d5b8d373e297f437"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.b / 38087516db13 / 5

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

<a id="canonical-a26532e4436ee68351cb3614afc7fb022b845ecb8c6e3164992b01c0c9199d6f"></a>

<a id="canonical-4564361f0d8806044994312a23fa9c2423f15a84682626eec24a4b0375c1ff22"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.b / 38087516db13 / 6

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

<a id="canonical-b8fc28a59e473ffcd50ed024e99da1ae5202014958741bfa963da3610a2ef587"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.b / 38087516db13 / 7

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](resources--voltstack_site--reference--group-006.md#canonical-0e8993c197e6fef63651797b3663aa6f5783a2b4ddad77fa078a75d1888ec801)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-312730f0207f41804051d362d25a2503c5d2fbb0a95125d4a4235e711a9d1d6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bcb5eab1d2eb7932d531e32f357b38174c40e6cf70456d77865d728db6ee231c"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.c / 8ccc98617439 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](resources--voltstack_site--reference--group-006.md#canonical-408b49d4b3038d0233a1a15d15c44f8a25a6212c7840568f58f066e6961127e1)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](resources--voltstack_site--reference--group-006.md#canonical-0e8993c197e6fef63651797b3663aa6f5783a2b4ddad77fa078a75d1888ec801)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.clear_secret_info

<a id="canonical-d120e2b9f3a30a396d21644d11ae95bf149d9f53024217ac42c527d53440ebed"></a>

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

<a id="canonical-307b05c5fc3c8e9f32b47af44cecfb7296e12328e6cffbd15593af8270be6f0e"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.c / 8ccc98617439 / 3

<a id="canonical-7145d5987523d642e7659b7a4b77b1f88e82e7be9692b03370ee377c3ea9c86b"></a>

<a id="canonical-4b685b66fc7fe0f6aab7faa7f2a87f07405171a1d40209352f21bee169719aee"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.c / 8ccc98617439 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-b693d76ed7fa874f150b39e146e3e67cbbe10c5966cd3fb8dc730a1d83663f22"></a>

<a id="canonical-73dad99be3ae91e1f981b20d0ee12a2502b8cf58aa76169e9bf6799655027c04"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.c / 8ccc98617439 / 5

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

<a id="canonical-0719b51f2752b3fd4e15a4448fd42d1a40e4a787e8d276ccb7ebca49ae2412f4"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.c / 8ccc98617439 / 6

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](resources--voltstack_site--reference--group-006.md#canonical-0e8993c197e6fef63651797b3663aa6f5783a2b4ddad77fa078a75d1888ec801)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4763d4447533fa068967521d90258c9b0726e3f6860e4f561f2fe73b36ed3c9a"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident — custom_storage_config.storage_device_list.storage_devices.netapp_trident / a3d8581e6b49 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident

<a id="canonical-2162c85f22a2a9ae52e2eae667f909cf605d9df5fc88ea05011ce77e3ceb9cbf"></a>

Type: `"object"`. single nested block, Optional.

Device configuration for NetApp Trident Storage.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("netapp_backend_ontap_nas",
    "netapp_backend_ontap_san")}
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
  "x-ves-oneof-field-backend_choice": "[\"netapp_backend_ontap_nas\",\"netapp_backend_ontap_san\"]"
}
```

Terraform syntax:

```terraform
netapp_trident {
  # Configure direct properties listed below.
}
```

<a id="canonical-b538971624f6cae426eb9f2a068cf876638dc50baef67eea57232bd2c853e1b7"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident / a3d8581e6b49 / 3

- [netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67): complete subsection reference.

- [netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04): complete subsection reference.

<a id="canonical-573c33c5679b03024f663493e4faab4e5e910f1f10dab1814a6cd3ebed36e8d0"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident / a3d8581e6b49 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-caffafbfd84623f5aea173b436dec70b793b2ed42e3d390b76b0a74b15c6f93a"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas

<a id="canonical-c1ae0239ce995c0194c3b997764ef9b0bb7ef9ed3e433aa82f5b7de06b39eb06"></a>

Type: `"object"`. single nested block, Optional.

Configuration of storage backend for NetApp ONTAP NAS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_driver_name",
    "username"),
  validators.ConflictingObjectAttributes("data_lif_dns_name",
    "data_lif_ip"),
  validators.ConflictingObjectAttributes("management_lif_dns_name",
    "management_lif_ip")}
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
  "x-ves-oneof-field-data_lif": "[\"data_lif_dns_name\",\"data_lif_ip\"]",
  "x-ves-oneof-field-management_lif": "[\"management_lif_dns_name\",\"management_lif_ip\"]"
}
```

Terraform syntax:

```terraform
netapp_backend_ontap_nas {
  # Configure direct properties listed below.
}
```

<a id="canonical-93eeb43c261390a7197863828115a9c6025978ef8a3371a92c4fbd0c705dfb51"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 3

- [auto_export_cidrs](resources--voltstack_site--reference--group-006.md#canonical-0d097e10f5546006b98f29f8bd06540cbdc56273348a7f9cc6d44ea57ea6f4b1): complete subsection reference.

<a id="canonical-520b43fc24e015f6621ac9514f8302d166be9d73d2107fa8ded75f02a4be8658"></a>

<a id="canonical-7cb28d3772cfef3e106a360b0633e447c044707f643206887100a5059c8d344a"></a>

## auto_export_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 4

Type: `"bool"`. Optional.

Policy configuration for this feature.

Upstream description:

Enable automatic export policy creation and updating.

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

<a id="canonical-d5467be23db895547b6a79f1154103991f316319a9cb9da7177a32574eaccde1"></a>

<a id="canonical-929e02bae8fd88dd04601b477dc3f31d76de5fc60644431bcd755cf782aac9c6"></a>

## backend_name property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 5

Type: `"string"`. Optional.

Configuration of Backend Name. Driver is name + '\_' + dataLIF.

Upstream description:

Configuration of Backend Name. Driver is name + "\_" + dataLIF.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 50),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 50,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 50,
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
    "ves.io.schema.rules.string.max_len": "50",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "50",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-4e5175bda20a7500598a059e44b343a189154ec54e0e2a9bd5b900ade2f1f10f"></a>

<a id="canonical-08c4e2fcd1a4979187da09240a385a60f2d8ee4ed816ffc02da89fc2fb491e2a"></a>

## client_certificate property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 6

Type: `"string"`. Optional.

Please Enter Base64-encoded value of client certificate. Used for certificate-based auth.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [client_private_key](resources--voltstack_site--reference--group-006.md#canonical-fbd2945ef5184c3b0591084ff7b262f7fd0535025cbf646177f2139ec85a451c): complete subsection reference.

<a id="canonical-11366d8734a036da00dfeaf92cb4d63ed82bb479ed0d770dc348beb0b29517bf"></a>

<a id="canonical-8acea0bdd0757cc5d67bd25cd66a95d446baf225d5579c186b2431c43f347fa3"></a>

## data_lif_dns_name property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 7

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-eaa3aa0c214fac86d9b81c8547ecd555bf6d6bd1572cfab71abf874163465c7b"></a>

<a id="canonical-d35df7c0866e53a154d93e0c9c26955f6a867b71bda7b8cd22d02c02c633f698"></a>

## data_lif_ip property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 8

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Upstream description:

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-ec8418ca08f81facaafde78629c6342684df545e4da4b6ed43f7d998adbbff27"></a>

<a id="canonical-442d21806e435f8ef7e5d63debb05db90a585605aa5fec75eca8a5318d50245c"></a>

## labels property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 9

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-dee0c1257e3cdfeca816b49da30094b7d24233b319fd4c0aef60b8b2495fc5a0"></a>

<a id="canonical-80736563af7b134e0d9d191981f2c992bed0c2f09aab972bcd208c6bff19730e"></a>

## limit_aggregate_usage property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 10

Type: `"string"`. Optional.

Fail provisioning if usage is above this percentage. Not enforced by default.

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

<a id="canonical-3b4eeed31e695904ab7ad3a9c6a2f464bafda6afd88faf70051ee1c2a537bff0"></a>

<a id="canonical-0b5f45050c4ebdba1e15741676086a8290dfddb793abe452cd9c7574c5a5c61e"></a>

## limit_volume_size property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 11

Type: `"string"`. Optional.

Fail provisioning if requested volume size is above this value. Not enforced by default.

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

<a id="canonical-10d5f27201e598326391f36f3b7e45ff9bb5be12bf3b906f2e57adffa9e8054f"></a>

<a id="canonical-399c70a3d11c7fe06d264b4efcfb6fe94a25680813a38bc7ff8d28ba93f40beb"></a>

## management_lif_dns_name property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 12

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-5bf500e0bd81406ab8efcd1effaca23e63fccaaed00976de1c5135f378692b75"></a>

<a id="canonical-bbc5e9ca448945188083ff120a83b43e6052de6834adf794fed2b3153e37a257"></a>

## management_lif_ip property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 13

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Upstream description:

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-aef94edbfc829f62480e44684248844937ae8981beb1f7cbcf22b72b8607d5c4"></a>

<a id="canonical-6bfc171549b49422477875c31e206eb4c029bc9f458bd0acc9f78a98ed174ae2"></a>

## nfs_mount_options property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 14

Type: `"string"`. Optional.

Comma-separated list of NFS mount OPTIONS. Not enforced by default.

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

- [password](resources--voltstack_site--reference--group-006.md#canonical-17f0345598e7d6599e1e89bf075231de9358c9dacefdaa6a5dfa306bc4a7dcef): complete subsection reference.

<a id="canonical-4802be454c87f1a176afddc7c8fa0138fd207d7448cf8f22fc743ea8efc09d0e"></a>

<a id="canonical-8a8728d8d55510878b926abc6a19bc4b742f8411f1d9de030a1ecf14b9d53178"></a>

## region property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 15

Type: `"string"`. Optional.

Backend Region. Virtual Pool Region.

Upstream description:

Virtual Pool Region.

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

- [storage](resources--voltstack_site--reference--group-006.md#canonical-03982cb82cb2a7c06c78f2a4731d9c226936f28d9d1366820e4a391c5b60ca7a): complete subsection reference.

<a id="canonical-b5ed0de4917dd53b5e2cd8d005a31a31968c75f23de192f1366a21dfcea6da43"></a>

<a id="canonical-1abb5abc884cb52ace01a36ab37b70e9279b7239a03392f8210c681eebbc2855"></a>

## storage_driver_name property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 16

Type: `"string"`. Optional.

\[Enum: ontap-nas|ontap-nas-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-nas\`, \`ontap-nas-economy\`, \`ontap-nas-flexgroup\`.

Upstream description:

Configuration of Backend Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ontap-nas",
    "ontap-nas-economy",
    "ontap-nas-flexgroup"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ontap-nas",
    "ontap-nas-economy",
    "ontap-nas-flexgroup"
  ],
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-nas\\\",\\\"ontap-nas-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-nas\\\",\\\"ontap-nas-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  }
}
```

<a id="canonical-7ed1f385cdbb9fa9df5073c0e3d498c9312f7d3791b23e832a021a504921374f"></a>

<a id="canonical-4fa56995f32dccde90b6c27eca49ae67fea560eebf1054e1fe7dbf4066c5e28a"></a>

## storage_prefix property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 17

Type: `"string"`. Optional.

Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.

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

<a id="canonical-fa3d4c6cc013c44c13fddcc405141633e45b8d137b3a8d85a877f6c1a6be89cb"></a>

<a id="canonical-f65c072469f9bb865654be848a34ffcbed4d107b75f32d9a2948c611099957ba"></a>

## svm property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 18

Type: `"string"`. Optional.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-28ab0ffb0938ca8c0a9186252462a583fe2040b261e0e39ad16533b7da54a233"></a>

<a id="canonical-14d67dec64b3cc14685957f1562330ce449c61a8e6ac470077e7cf640119a5d8"></a>

## trusted_ca_certificate property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 19

Type: `"string"`. Optional.

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Upstream description:

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth..

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

<a id="canonical-01c87db1fdc72fad9e5979630ba227645a0b57e345fcb55696dced35fd02cfc0"></a>

<a id="canonical-b0e6e8a5aa8c03494f04945205649c0d1dd1ec1e3f7a66c5de5f960e64227b72"></a>

## username property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 20

Type: `"string"`. Optional.

Username. Username to connect to the cluster/SVM.

Upstream description:

Username to connect to the cluster/SVM.

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
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [volume_defaults](resources--voltstack_site--reference--group-006.md#canonical-9791b3bca89df31e32a868e53f74e9d82e641e4659846347cf587a04be86a84e): complete subsection reference.

<a id="canonical-1f0f69c1464d26a55d6a9ae34b0fd68fdaf000a7dada46fc456f2b84013c165b"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 8cac76976f39 / 21

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs](resources--voltstack_site--reference--group-006.md#canonical-0d097e10f5546006b98f29f8bd06540cbdc56273348a7f9cc6d44ea57ea6f4b1)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--voltstack_site--reference--group-006.md#canonical-fbd2945ef5184c3b0591084ff7b262f7fd0535025cbf646177f2139ec85a451c)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--voltstack_site--reference--group-006.md#canonical-17f0345598e7d6599e1e89bf075231de9358c9dacefdaa6a5dfa306bc4a7dcef)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--voltstack_site--reference--group-006.md#canonical-03982cb82cb2a7c06c78f2a4731d9c226936f28d9d1366820e4a391c5b60ca7a)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](resources--voltstack_site--reference--group-006.md#canonical-9791b3bca89df31e32a868e53f74e9d82e641e4659846347cf587a04be86a84e)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-0d097e10f5546006b98f29f8bd06540cbdc56273348a7f9cc6d44ea57ea6f4b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b7e0adeffd4531588ca1c1ea1a813361860495b2eb6a38ddc9f47e64ce9dfbc"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / b0c3505a9b7a / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs

<a id="canonical-4fca9dbfdc232a74e1942eaa1cfa7ce1baad5d0c43e6dfbeddedcd231016dab4"></a>

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
auto_export_cidrs {
  # Configure direct properties listed below.
}
```

<a id="canonical-3277b67e473389abc019fc5b5a7abc74c9dcbb33007a98006b63b31be7f27cf5"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / b0c3505a9b7a / 3

<a id="canonical-0dbfd9e6f683720cab8208a6fa201233c6102f465bf0199ca292e0e460b8b52a"></a>

<a id="canonical-3a0a72454fa2ef5c93f8f15a88dfbcc249cc9a1cdcb0ff50d2818c85ee6ae7ce"></a>

## prefixes property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / b0c3505a9b7a / 4

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

<a id="canonical-accd5f113c9d7c8fd9aafd3522773e11ea748607149bd8285af818b9b74ff18e"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / b0c3505a9b7a / 5

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-fbd2945ef5184c3b0591084ff7b262f7fd0535025cbf646177f2139ec85a451c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b29aca90ee5e57d9351d8181177d52fc779405a97dde257251c523fd6912d96"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ca873a014c41 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key

<a id="canonical-98b675812864de5fda7182a3fb66041a64ed935ca4a0106c04f1ea0c8724877c"></a>

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
client_private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-990accbfaf30da0530bc86da33626b4d4123a9a0ff1bb28da29e0159a792a3b7"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ca873a014c41 / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-006.md#canonical-02643eab84a47096cc9cda5b7875fa3adde67eb02da22d54d89123fbbe5d0886): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-006.md#canonical-473e0f2885e4c4b5f3ca380ee3ce143da99bd9e9c534485ba3a3a2df613e9c80): complete subsection reference.

<a id="canonical-9cf0c85bdfcf5d0eea65f8ec81a2f3502c3c40837e55fec31e2fa99d93eb68c6"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / ca873a014c41 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info](resources--voltstack_site--reference--group-006.md#canonical-02643eab84a47096cc9cda5b7875fa3adde67eb02da22d54d89123fbbe5d0886)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info](resources--voltstack_site--reference--group-006.md#canonical-473e0f2885e4c4b5f3ca380ee3ce143da99bd9e9c534485ba3a3a2df613e9c80)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-02643eab84a47096cc9cda5b7875fa3adde67eb02da22d54d89123fbbe5d0886"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-91ff71f76cf0c2fbc2ec13ba3108203e8e3d4f1d7acb4f3d348399d93466da41"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0c94d838ef16 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--voltstack_site--reference--group-006.md#canonical-fbd2945ef5184c3b0591084ff7b262f7fd0535025cbf646177f2139ec85a451c)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info

<a id="canonical-38603fdf8366c6ed61c2b69d64dcbb6f169c251011a45dd96f58434f3a19e4e8"></a>

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

<a id="canonical-5219b493326dba0f23883072d74ccd7a4795e84d8ff817c0d592c7341c98d082"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0c94d838ef16 / 3

<a id="canonical-fc6cf90b9e0834eabf50fedde5924d99cb1151aeeb8c024b2332e3467815f724"></a>

<a id="canonical-739ffede4836c0a195e9639b35d7f5e712b49fee912f9934bb3f66a98af4a0ff"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0c94d838ef16 / 4

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

<a id="canonical-0f016f1ef79c4f92a9802f5d860d881bba51eb6a66f16df5b3fec7e6c826da62"></a>

<a id="canonical-92bd831fc1c130b115ebe1c8a16ba76b2ad4602b83b5f72a44448cccb5138662"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0c94d838ef16 / 5

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

<a id="canonical-6c7c224079c8f5785d88c4e0e50d229de0df6af885f9006a883438ce971244e6"></a>

<a id="canonical-c82ccdf34f94fd86a94d7caf6dc7a97cb906bce9efdc0444c31b0bb107439c09"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0c94d838ef16 / 6

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

<a id="canonical-7eb16e72005f9534b47ff4b4dfd0c3ce39d54243068c27730507a9095d570163"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0c94d838ef16 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--voltstack_site--reference--group-006.md#canonical-fbd2945ef5184c3b0591084ff7b262f7fd0535025cbf646177f2139ec85a451c)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-473e0f2885e4c4b5f3ca380ee3ce143da99bd9e9c534485ba3a3a2df613e9c80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fb32db6155068b0a00237feabfa213fbcb031012317f57323e3b5d937be8dee"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2eb5f10d8c63 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--voltstack_site--reference--group-006.md#canonical-fbd2945ef5184c3b0591084ff7b262f7fd0535025cbf646177f2139ec85a451c)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info

<a id="canonical-ccdcca5e671c1331b734676ea56fb638d65b56c3917b6e0c6249ccb9560a4c18"></a>

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

<a id="canonical-a066893aa241a4a1d9a6ffbd0fc55bdbdc91b6a62701965d1a47d1bff71ad46c"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2eb5f10d8c63 / 3

<a id="canonical-9bc11a206f6479f63c861c10b6ccf939522d3e19f06e0e9329135f56a86474f4"></a>

<a id="canonical-c1e4f26ea5b64628cc9e30291f2e11b81fa0c107f51ce9d16c74d1a31aa501be"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2eb5f10d8c63 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0c9674a89fa40af15aefb0e082ee43db75ad628ab2ee69a7c99b4c7b6f7254c0"></a>

<a id="canonical-cfd2e9bdfb13c08f8f5a4c705eb4bbeb5b01ce9762fc1606ffaae5a455cfe045"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2eb5f10d8c63 / 5

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

<a id="canonical-89e4a7ea663069a44095666bfd7426f2195b0ccccc7304f0c83a59b14b602a61"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2eb5f10d8c63 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](resources--voltstack_site--reference--group-006.md#canonical-fbd2945ef5184c3b0591084ff7b262f7fd0535025cbf646177f2139ec85a451c)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-17f0345598e7d6599e1e89bf075231de9358c9dacefdaa6a5dfa306bc4a7dcef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54aa345e5b4acdca57e3cd63bca19a5c094b641a6ecb7b895cb4266ba151b56a"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 94f94bb634c4 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password

<a id="canonical-9a77aeea69ec4f64f8c9192ad07456b1d6edaabd29662bccea6f66a58a1b8241"></a>

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

<a id="canonical-13555929392c92318bf392f8894f9f170b3b71393c72a5681d3c657385fc9a76"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 94f94bb634c4 / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-006.md#canonical-d3bbb1c03f4b3c1c42f4a6deacccb5c0dc0836755c57c356da30e3dae5985801): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-006.md#canonical-67b45e274eb298edc095bdf8541d14a13cca6030c2a1431ba392064cc9ca6e9b): complete subsection reference.

<a id="canonical-62cc397a1c8666d6fc09cdd48afa811cfd990e8ecb8414ab3659d74b2e1c4bc4"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 94f94bb634c4 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info](resources--voltstack_site--reference--group-006.md#canonical-d3bbb1c03f4b3c1c42f4a6deacccb5c0dc0836755c57c356da30e3dae5985801)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info](resources--voltstack_site--reference--group-006.md#canonical-67b45e274eb298edc095bdf8541d14a13cca6030c2a1431ba392064cc9ca6e9b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-d3bbb1c03f4b3c1c42f4a6deacccb5c0dc0836755c57c356da30e3dae5985801"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b25462c5e44d2e0a3d43c39147d43dd4c0bd3b3a5afa9bcb558640e7b74dcb2c"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 40bd20d11480 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--voltstack_site--reference--group-006.md#canonical-17f0345598e7d6599e1e89bf075231de9358c9dacefdaa6a5dfa306bc4a7dcef)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info

<a id="canonical-2a16d2d1602f4e0cc2d2a74d304b7de8aeb29f404eb55452bfc01a56f896f5e1"></a>

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

<a id="canonical-06744d0e40775e10196cba92e0e9a5a9b7a6be8e9308ae23ca084bc93738ead0"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 40bd20d11480 / 3

<a id="canonical-ef38e6801c8752313254f6c9d91cce0e9f79f659829b41cd92d25f370d58b4f4"></a>

<a id="canonical-4966037f3fe2f900d35957cc865bda69fa9542d5209cd90b8019ef5199fc3c37"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 40bd20d11480 / 4

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

<a id="canonical-fee62d8bf7b477b3f7e3626ec8ac3801f7c2f81c73c9a693677651cc1946e8e2"></a>

<a id="canonical-f096bedd07b4dcd7e3b40c4e8577d6ee198e0b0ccb5842428e4b823890614cd0"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 40bd20d11480 / 5

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

<a id="canonical-589dbe17e0acf84770f530aeeabb86591653fb2953f110b12f94538790173554"></a>

<a id="canonical-06d6640206c908f7a1731d9247cf95a80ec8bb613573d5d70ac115464f6f81b4"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 40bd20d11480 / 6

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

<a id="canonical-19a68c6d4ed43c0fdffd22d268ba1979243f5657049e6a1b772cfec0d1111d9d"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 40bd20d11480 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--voltstack_site--reference--group-006.md#canonical-17f0345598e7d6599e1e89bf075231de9358c9dacefdaa6a5dfa306bc4a7dcef)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-67b45e274eb298edc095bdf8541d14a13cca6030c2a1431ba392064cc9ca6e9b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f616b073aa6c63c04ced2d7246f4844a9be43c2d3fec7e777aa5ddd7d4ae1328"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5077d13cd855 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--voltstack_site--reference--group-006.md#canonical-17f0345598e7d6599e1e89bf075231de9358c9dacefdaa6a5dfa306bc4a7dcef)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info

<a id="canonical-2584c1733bb723dc6fd37849911e7d15c98de6630f631f014caed6cf51828bc8"></a>

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

<a id="canonical-af59c195ec08e3ee2ce399e6900e5d3b79c3349a8e1b485d5146ec72a02d646d"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5077d13cd855 / 3

<a id="canonical-1a235ada72e876fc7890e4a2a4dddd25456ce8f874c7f09f4a3a3f2bb0b317f1"></a>

<a id="canonical-a620665e8ffd51d443929578a9f229a604c95ea8db624152c9ee859f01dcd03a"></a>

## provider_ref property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5077d13cd855 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-d5601e7d3d1429211cfef10baedb938ffecb2d92a087aed1cd0af6ee2e1d87d8"></a>

<a id="canonical-1d27ce1fee91a84d470bfd89dc19c431607bf027fb847f58da3b61b16915d206"></a>

## url property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5077d13cd855 / 5

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

<a id="canonical-599aceead85362852f5baef23154e6256a32da8c15cdb8fe10ac1abd028dc086"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5077d13cd855 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--voltstack_site--reference--group-006.md#canonical-17f0345598e7d6599e1e89bf075231de9358c9dacefdaa6a5dfa306bc4a7dcef)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-03982cb82cb2a7c06c78f2a4731d9c226936f28d9d1366820e4a391c5b60ca7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95ee590c678f54576e0ef4f30d27fda09eaca5e8c7562698f33f1edc2ab72ac5"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d1e90e129da1 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage

<a id="canonical-2eb2c64734635228eed9588a3035509d2268f9505647883421bf0cea9684c18d"></a>

Type: `"object"`. list nested block, Optional.

List of Virtual Storage Pool definitions which are referred back by Storage Class label match
selection.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-5dc382f78187b44c2d0ac79a14fa32c5eb495d7be1bb561613d954946fa0e337"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d1e90e129da1 / 3

<a id="canonical-c79cdf5a5b31d03576b49b4f1974b323715260ca6640407e67b5d728aab76d7b"></a>

<a id="canonical-61c0f2f941a9910fd59fd8a5ae6fb04b62c81c88f0a81a594abf155101200145"></a>

## labels property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d1e90e129da1 / 4

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match
selection.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [volume_defaults](resources--voltstack_site--reference--group-006.md#canonical-98559e1f0ae78dfd9e8b3e29d1f344260ccb0a7f24ee38fed72879414cde9d75): complete subsection reference.

<a id="canonical-3fc03558331b262eb6bf9edfd229398acdc52f5756ac06998a05001fb09bf6bc"></a>

<a id="canonical-2923a6582de76c6a6251673b13f968d7d95255a7d341f8dc1aecef4f87bc313f"></a>

## zone property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d1e90e129da1 / 5

Type: `"string"`. Optional.

Virtual Pool Zone. Virtual Storage Pool zone definition.

Upstream description:

Virtual Storage Pool zone definition.

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

<a id="canonical-3ff4b6f3683bdefcaf668f59c3b3607d2a299c4d58b5a4c4713c8f59476c6f4d"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d1e90e129da1 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](resources--voltstack_site--reference--group-006.md#canonical-98559e1f0ae78dfd9e8b3e29d1f344260ccb0a7f24ee38fed72879414cde9d75)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-98559e1f0ae78dfd9e8b3e29d1f344260ccb0a7f24ee38fed72879414cde9d75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca2fe31e189234eda88bffb6754f2020e914599136533b8d6684092425629745"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2a90054c6965 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--voltstack_site--reference--group-006.md#canonical-03982cb82cb2a7c06c78f2a4731d9c226936f28d9d1366820e4a391c5b60ca7a)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults

<a id="canonical-65bb9704876188e723cf5eee757558f84600d37ac06eebf2dcc74c55a45860f5"></a>

Type: `"object"`. single nested block, Optional.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "no_qos"),
  validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "qos_policy"),
  validators.ConflictingObjectAttributes("no_qos",
    "qos_policy")}
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
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

Terraform syntax:

```terraform
volume_defaults {
  # Configure direct properties listed below.
}
```

<a id="canonical-2b8f11607e3938cdb8487ff3ef7907c3a01fc9248c5401ddf255edd185e5a2ea"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2a90054c6965 / 3

<a id="canonical-30503bc706fbf9d9303d19ff8ecddd4985f0c27f44689b1791fb3c80440e0a02"></a>

<a id="canonical-cdd8304607625251644ce12a784db8d1c28ae5fe9e605940f3a40f96d47b697d"></a>

## adaptive_qos_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2a90054c6965 / 4

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-c3789e5d260d92bb6319f5bb35e6083c79893cfa4191eca3b114ace8740f79a6"></a>

<a id="canonical-da63d909518d94cacd495a781bebb9f0c40e93c313a79366b004ba11fb177129"></a>

## encryption property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2a90054c6965 / 5

Type: `"bool"`. Optional.

Enable Encryption. Enable NetApp volume encryption.

Upstream description:

Enable NetApp volume encryption.

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

<a id="canonical-94e937c05aba18e6a507c29a79a25034057dc5d8d9802934c7e7fe5c4ee7a4d3"></a>

<a id="canonical-414079b4b4b580320eebbf5070e3bf1ff4e9ba131a1fef04aad8afedd669f7b5"></a>

## export_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2a90054c6965 / 6

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Export policy to use.

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

- [no_qos](resources--voltstack_site--reference--group-006.md#canonical-c5fc55b6c6a9335468701a39c36e5086fc943039206bb03195eaae290a3df8c1): complete subsection reference.

<a id="canonical-da97fb98b23fa3b4cebae21f642b8539cd2f7be8da69b13382f87f3fb77b8667"></a>

<a id="canonical-161690824a757c5fa2ec8edbc3c94fbdd356f161c349c1938eba227d6713c90a"></a>

## qos_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2a90054c6965 / 7

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2a27785d8258a31c8cb5d91f3c6d7adc28fa0e37c225db5151e9fe813b8cdab0"></a>

<a id="canonical-dd6e0ee7832e6ca055549350da8dcddbb70bd41e7caa948a81305bf8528b303b"></a>

## security_style property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2a90054c6965 / 8

Type: `"string"`. Optional.

Security Style. Security style for new volumes.

Upstream description:

Security style for new volumes.

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

<a id="canonical-942bf9fe3c7825afc0f1c1567c87506bbf8ddb7ceba54976808bed7107b5d8fb"></a>

<a id="canonical-e663b36104ca7f33ecbb36ad7f04767f108e5db0a10864b9412e79ffae48a3b0"></a>

## snapshot_dir property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2a90054c6965 / 9

Type: `"bool"`. Optional.

Access to Snapshot Directory. Access to the .snapshot directory.

Upstream description:

Access to the .snapshot directory.

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

<a id="canonical-bdc884ca28a87076893ff739185719d4ef350f70ee9d2ea95078f786443b276e"></a>

<a id="canonical-25a92728c2a8790d4375fa90652e75e422d390a22e02d5fdbd6c2f57e3f5a8e0"></a>

## snapshot_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2a90054c6965 / 10

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Snapshot policy to use.

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

<a id="canonical-c7de3f59ed7886683f8e40374425fa90c896babb9a7ae0eec3f01b555498e47d"></a>

<a id="canonical-9a922c2847691c014550a7a85a4a5f16fc46d987e9e20de7b12af860ee5e6f41"></a>

## snapshot_reserve property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2a90054c6965 / 11

Type: `"string"`. Optional.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Upstream description:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

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

<a id="canonical-fce054b74574b3633434cbaaf0cdf5416a3ce15728e8ced48645d718c39b055b"></a>

<a id="canonical-7f40e63ab85e7fe2acfd81c49e7d21843319f20294ce5362324e7cad8b1daf90"></a>

## space_reserve property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2a90054c6965 / 12

Type: `"string"`. Optional.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("none",
    "thick"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-6100487b1d670d7881eb62d2b930bebc619376dee4d356fbaa956eb1658ea4d7"></a>

<a id="canonical-5bcd75605114f416bbd157fdf7f38dd9a5df3d86f19219b6b7284f16b3172574"></a>

## split_on_clone property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2a90054c6965 / 13

Type: `"bool"`. Optional.

Split a clone from its parent upon creation.

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

<a id="canonical-9a0b97e9ed2c6165d2b652471ef6c5e4ddf369848418bad05bdc08e7c9236320"></a>

<a id="canonical-1958f34b8ddcf767e3ca4a5202111d15c246fb4d4cd4f5463b04c6cd1553b0d1"></a>

## tiering_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2a90054c6965 / 14

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Tiering policy to use. "none" is default.

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

<a id="canonical-4c8c1019c1b5ebd47acf65c6b5ae49fc83c06c4b229cb871e89b827b0067221c"></a>

<a id="canonical-a8dc5618796557c18778bbde35b138d7fea4b05582de95bc802a6fa799ab5410"></a>

## unix_permissions property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2a90054c6965 / 15

Type: `"number"`. Optional.

Unix permission mode for new volumes. All allowed 777.

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

<a id="canonical-12c5545bd5287a8bcecfca102bf271ef84274ab3a8b2519a3fb2aed7901beeae"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 2a90054c6965 / 16

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos](resources--voltstack_site--reference--group-006.md#canonical-c5fc55b6c6a9335468701a39c36e5086fc943039206bb03195eaae290a3df8c1)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--voltstack_site--reference--group-006.md#canonical-03982cb82cb2a7c06c78f2a4731d9c226936f28d9d1366820e4a391c5b60ca7a)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-c5fc55b6c6a9335468701a39c36e5086fc943039206bb03195eaae290a3df8c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9efa65a5dab440f5711d3003e5c67def6fe07ca3783205736f1b0701af879d43"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 6b4e9955be11 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--voltstack_site--reference--group-006.md#canonical-03982cb82cb2a7c06c78f2a4731d9c226936f28d9d1366820e4a391c5b60ca7a)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](resources--voltstack_site--reference--group-006.md#canonical-98559e1f0ae78dfd9e8b3e29d1f344260ccb0a7f24ee38fed72879414cde9d75)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos

<a id="canonical-7c66cae977ca9f3bac063b17c75762c925dd1ef09ccc5d38500d13d1cda99ca1"></a>

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
no_qos = {}
```

<a id="canonical-449297bb6b305655e4d19be153d5fd754c761236eaa3650ce49eac54f68899a8"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 6b4e9955be11 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-52191d48df31896f3e0aac440c5bdc85f579c3931d524c3617ecd81f3de0294c"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 6b4e9955be11 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](resources--voltstack_site--reference--group-006.md#canonical-98559e1f0ae78dfd9e8b3e29d1f344260ccb0a7f24ee38fed72879414cde9d75)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-9791b3bca89df31e32a868e53f74e9d82e641e4659846347cf587a04be86a84e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f97152d3bcacf0f6e67315761c3469d8778ed2ccf6c33de4e936595a65971ab"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d5db81f41e99 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults

<a id="canonical-304e714b65d046bf3b4a2df5fa92fd3a6be71247e1267f2bc5a748fcdb43d204"></a>

Type: `"object"`. single nested block, Optional.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "no_qos"),
  validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "qos_policy"),
  validators.ConflictingObjectAttributes("no_qos",
    "qos_policy")}
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
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

Terraform syntax:

```terraform
volume_defaults {
  # Configure direct properties listed below.
}
```

<a id="canonical-bc73d185b0e6687df2677d8feb1e6977ca7276f406dd1f567ce27107017f854f"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d5db81f41e99 / 3

<a id="canonical-b77181b6ac684ce7bbf818c2ab399c8cfe7ec98a8fca9faa4f42a1a6e97b149b"></a>

<a id="canonical-deb237f3c9ccdcf8a911f380f741ac93157f4199599c44da11a0cb56fdb7f511"></a>

## adaptive_qos_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d5db81f41e99 / 4

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-fdb7ecd1743dfbc854838de68c939268c7c957b987a6258f02caaa73814b2683"></a>

<a id="canonical-1befb8725095bb17bac3751dc503150df91f64abb19c90e5442c98fbbe065c87"></a>

## encryption property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d5db81f41e99 / 5

Type: `"bool"`. Optional.

Enable Encryption. Enable NetApp volume encryption.

Upstream description:

Enable NetApp volume encryption.

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

<a id="canonical-11a7ced86ca98478438902a3c1d29272a7fdcdf4f15e7d5e29b511daae1d5726"></a>

<a id="canonical-89fb2f906252a3ab2a14924ec6ad405f8ec182d6780b751fdc4eb690af1fb10e"></a>

## export_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d5db81f41e99 / 6

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Export policy to use.

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

- [no_qos](resources--voltstack_site--reference--group-006.md#canonical-a2722ffafa558173d6f59bdce212c813e60fafd34bc13d8b67737938f9551e16): complete subsection reference.

<a id="canonical-5f1e37a1a415fe1c5d23e875d193d546101536af0e7f45d1a526499281a6c146"></a>

<a id="canonical-7ed7d4de690d2f9abc8da594346da8ba49d2c77b6996d2af657620c59e7d7990"></a>

## qos_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d5db81f41e99 / 7

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-456f5701321e3d78981f77f10f886fccfac6cd02bffa264715ab634d3fd06e9b"></a>

<a id="canonical-3941d4045639fd4cae37651ae97e9727df8b6c834b5fc61c18d0a8e287b5f90d"></a>

## security_style property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d5db81f41e99 / 8

Type: `"string"`. Optional.

Security Style. Security style for new volumes.

Upstream description:

Security style for new volumes.

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

<a id="canonical-31431df62a18c7b897caa1c9964541f031e7ebe240fc5f372f7679d169538542"></a>

<a id="canonical-3d7f482f8ca3669d68ee55061546dd2371616b309a37817902f8caa00d228266"></a>

## snapshot_dir property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d5db81f41e99 / 9

Type: `"bool"`. Optional.

Access to Snapshot Directory. Access to the .snapshot directory.

Upstream description:

Access to the .snapshot directory.

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

<a id="canonical-7eaa95e7365ed9180d5c214d2346400ecc18fa47b9d499c3e0d1a617e07f7361"></a>

<a id="canonical-4be32d4e8279c87e64734e2f6b94d73999ab28a1957906e400e48741953322b2"></a>

## snapshot_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d5db81f41e99 / 10

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Snapshot policy to use.

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

<a id="canonical-c8ff3507057a4f94b58c623a5163c0afa035d2dd326c45c64564fcc4f057be75"></a>

<a id="canonical-5dbe2cb38f0751b12c36d73cd6a4ed1399d9eab958fd3d1bcd4afe5c525b1838"></a>

## snapshot_reserve property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d5db81f41e99 / 11

Type: `"string"`. Optional.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Upstream description:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

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

<a id="canonical-8a8fcf5ba32891de0afe0025fc7dcd62539b19268a00d254a0830d7bef1f425d"></a>

<a id="canonical-0189dadaab2c11781529153fe546a1702833383d0d08cbfc813c6d069b30902e"></a>

## space_reserve property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d5db81f41e99 / 12

Type: `"string"`. Optional.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("none",
    "thick"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-37d6a24758c4e8effa2096ed03baafa71fe333dfa2d59d1362208805e79d0c64"></a>

<a id="canonical-c8e178718a1b219a28acbe60361d8cd7a00f37ce882d820bf4d77fba5908ef34"></a>

## split_on_clone property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d5db81f41e99 / 13

Type: `"bool"`. Optional.

Split a clone from its parent upon creation.

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

<a id="canonical-8da2a90f3e376de77b7ab34ebfd7fda5ee7362deddec52e26259bfc94b8789eb"></a>

<a id="canonical-70632421b1e0e7e5ae9240ce5d9231fc3446bbf1fe8107d72f534a4be8f5a931"></a>

## tiering_policy property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d5db81f41e99 / 14

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Tiering policy to use. "none" is default.

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

<a id="canonical-6d75abbf5c4464216e34e3ee51623d7d1fa1b135ace3e30347615bf1d4c8842b"></a>

<a id="canonical-561f959f49e92d402a4a2592317aa635fe786633a615b54718fe8f00b33e0df4"></a>

## unix_permissions property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d5db81f41e99 / 15

Type: `"number"`. Optional.

Unix permission mode for new volumes. All allowed 777.

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

<a id="canonical-dbc610048313412d0c8e6544452f06dfa0f9e05b7c12f4e57a98d9408c48eaa9"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / d5db81f41e99 / 16

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos](resources--voltstack_site--reference--group-006.md#canonical-a2722ffafa558173d6f59bdce212c813e60fafd34bc13d8b67737938f9551e16)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-a2722ffafa558173d6f59bdce212c813e60fafd34bc13d8b67737938f9551e16"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbede4ce6ea380e1c2a081cc6b763bcd7b254abb23a530570a7b029566e94a5f"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / fc59ed08ea72 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--reference--group-006.md#canonical-cfa91cfc456ead2f7ca143d9fd5eaed1e1527e68587fbc2f5a69c5951de24a67)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](resources--voltstack_site--reference--group-006.md#canonical-9791b3bca89df31e32a868e53f74e9d82e641e4659846347cf587a04be86a84e)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos

<a id="canonical-0f6baed04bc82a28199d2aab8ae9ebfbc69e7bea540a5eccdc9744cbf3163041"></a>

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
no_qos = {}
```

<a id="canonical-6ec869c5efc895ead5f5ed2b4148a5a195549310ab7292f0be9a9c7ec6716991"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / fc59ed08ea72 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-837c9cd21035a1dfd026bc87fe03edb9dcb606334d5e65d0cc7d7b6bacf17b5c"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / fc59ed08ea72 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](resources--voltstack_site--reference--group-006.md#canonical-9791b3bca89df31e32a868e53f74e9d82e641e4659846347cf587a04be86a84e)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7828dd1f58bd8b1f9b4ba9147728132785af72f5c4239881f04b84d59ad5809e"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5a70dea8e316 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san

<a id="canonical-ea749f320c1683e912e933048e6bd07e5275c69e160731dc34c3487da8d2784c"></a>

Type: `"object"`. single nested block, Optional.

Configuration of storage backend for NetApp ONTAP SAN.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_driver_name",
    "username"),
  validators.ConflictingObjectAttributes("data_lif_dns_name",
    "data_lif_ip"),
  validators.ConflictingObjectAttributes("management_lif_dns_name",
    "management_lif_ip"),
  validators.ConflictingObjectAttributes("no_chap",
    "use_chap")}
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
  "x-ves-oneof-field-chap_choice": "[\"no_chap\",\"use_chap\"]",
  "x-ves-oneof-field-data_lif": "[\"data_lif_dns_name\",\"data_lif_ip\"]",
  "x-ves-oneof-field-management_lif": "[\"management_lif_dns_name\",\"management_lif_ip\"]"
}
```

Terraform syntax:

```terraform
netapp_backend_ontap_san {
  # Configure direct properties listed below.
}
```

<a id="canonical-9aea91e2ccc6c7f1b899e387b31b4d4c9ac860e9cbdbb51ac4bc2513bb77116d"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5a70dea8e316 / 3

<a id="canonical-46cf2faebaf940e2dd2377660533dc9706fefd1d49a998e682194f180c53eaf2"></a>

<a id="canonical-43105073e1be0cba0aa74d3429f38315f9d0e113493dc26ab25dfa964ad236e4"></a>

## client_certificate property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5a70dea8e316 / 4

Type: `"string"`. Optional.

Please Enter Base64-encoded value of client certificate. Used for certificate-based auth.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [client_private_key](resources--voltstack_site--reference--group-006.md#canonical-bc80dd9d42d6825d0ec6d1e650739edaa480568a7018f14a7813e8c45d0a36fa): complete subsection reference.

<a id="canonical-d055c5353bf152268c383ea67ad3a1b7e6595be918da42b0f7a6f68253bed89e"></a>

<a id="canonical-6f0525492e1bd5f56380d536b14831df324abbaa39132d24f44000ce00067f11"></a>

## data_lif_dns_name property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5a70dea8e316 / 5

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-6cef8d66f9019028e244cb13cf6741a5884ec30f50aa939946b39ae0d37886c0"></a>

<a id="canonical-f83033a919c48e1eef4be32b6d8639d2776bfd7e43d41773695d4430d6302d12"></a>

## data_lif_ip property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5a70dea8e316 / 6

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Upstream description:

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-d7514b217f761fb234022c2529bff6a241416b9c150164ca8e094c2b8ce773a3"></a>

<a id="canonical-3b1a54244b28858fb81c2e4950e63f075d79778ebc38d3303d8bef9106002db5"></a>

## igroup_name property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5a70dea8e316 / 7

Type: `"string"`. Optional.

Name of the igroup for SAN volumes to use.

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-7c290fd3d1fa0f421a8df9eee6f874bd85602e8f77ee1ed8854844531046975c"></a>

<a id="canonical-cb44c5fe739f60e1c8681ba99be6d929675412ea3e80844246125ba2bab3cc4b"></a>

## labels property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5a70dea8e316 / 8

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-e71f7493b88745eae28f926d482c7c21d6617f059f8ed0d37ad5a9f1e9abc6a2"></a>

<a id="canonical-c573d94d8f688b063bbdd9c8e2ec07d7bcbdcf215599af490724d49c8ce00fb6"></a>

## limit_aggregate_usage property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5a70dea8e316 / 9

Type: `"number"`. Optional.

Fail provisioning if usage is above this percentage. Not enforced by default.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-c4de16b0e2b1dbdea6f65b1ea4f99e778e8bfe1cefb1b38a2ff94cb3f0567380"></a>

<a id="canonical-49151786d9d6be2eb276743aabf1de6bdd62827b354a59407117efde08929119"></a>

## limit_volume_size property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5a70dea8e316 / 10

Type: `"number"`. Optional.

Fail provisioning if requested volume size in GBi is above this value. Not enforced by default.

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

<a id="canonical-98c50a398b153cdf4112efcc20b8a0ed8e20caf390f9b8ad8ea350cf2d795551"></a>

<a id="canonical-b24a389540c04f65d10d5871f80142d8dfec62eb90457a20b08d922930cea2a2"></a>

## management_lif_dns_name property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5a70dea8e316 / 11

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-4a136fecf254927b924c6d75c349eb59e37535c714a5c26eebe2d78c20e23d85"></a>

<a id="canonical-421d4bb792b35712c658b143b415161e117749190e0106cf15b2a7957e0e2d51"></a>

## management_lif_ip property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5a70dea8e316 / 12

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Upstream description:

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

- [no_chap](resources--voltstack_site--reference--group-007.md#canonical-ddc300ee675f6654917a2597aaa201b4f5df02043b47315e45e3081cd991d4c8): complete subsection reference.

- [password](resources--voltstack_site--reference--group-007.md#canonical-9edae627e94b200f98ab5d8f9076889619e850d6fcf6336daa0d1f7b9dccaf6d): complete subsection reference.

<a id="canonical-7eb8a47a02972719d16056010b4cd6ebfe98d6903eb61900b173cece6bb8b93b"></a>

<a id="canonical-f1f0c800821164186f305565bbae8100410d76382be8b58d951c4ed864fbb4a2"></a>

## region property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5a70dea8e316 / 13

Type: `"string"`. Optional.

Backend Region. Virtual Pool Region.

Upstream description:

Virtual Pool Region.

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

- [storage](resources--voltstack_site--reference--group-007.md#canonical-8987b76c1b3b57fede2c84f453c17907fe933fe1ebcd9a0bea1ee6534c5da46a): complete subsection reference.

<a id="canonical-00c5b9f017e22151ebd953fd3f781a8d62452b293d025ff5b33ea6bb4c308b4e"></a>

<a id="canonical-3844f6e609ccbb7f9a3a7c3ec8d32d74f7bb6954414bf46f94ad1d6e849291da"></a>

## storage_driver_name property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5a70dea8e316 / 14

Type: `"string"`. Optional.

\[Enum: ontap-san|ontap-san-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-san\`, \`ontap-san-economy\`, \`ontap-nas-flexgroup\`.

Upstream description:

Configuration of Backend Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ontap-san",
    "ontap-san-economy",
    "ontap-nas-flexgroup"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ontap-san",
    "ontap-san-economy",
    "ontap-nas-flexgroup"
  ],
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-san\\\",\\\"ontap-san-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-san\\\",\\\"ontap-san-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  }
}
```

<a id="canonical-3f710119154a4150546fbb95feccd24b7c916f7e87f9cdb0b8435312006835c2"></a>

<a id="canonical-1b2a879acb31d5161bfa6c2d87399ed4490d40fde5b3f9c4a2f45eecebd33f1e"></a>

## storage_prefix property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5a70dea8e316 / 15

Type: `"string"`. Optional.

Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 80),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 80,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 80,
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
    "ves.io.schema.rules.string.max_len": "80",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "80",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-32a8026505d0c3c2e49af59a2f96869b72629779fbd135a4c6d7695fc0a78920"></a>

<a id="canonical-affd25a76577eff3044bf3d49beb6ecf82b6f2e8394772c945f56ee15aeb33ef"></a>

## svm property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5a70dea8e316 / 16

Type: `"string"`. Optional.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-7c10fac35bc77a9401b0cb348acc8cd8ed5e6c066f9ecdf7aa7f8855686b2254"></a>

<a id="canonical-733aa554576ff1415648d3549caa953beb7e3aeb23330cfa3a038d8df8cf2a7f"></a>

## trusted_ca_certificate property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5a70dea8e316 / 17

Type: `"string"`. Optional.

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Upstream description:

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth..

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [use_chap](resources--voltstack_site--reference--group-007.md#canonical-72930290386fbf5a47dda8bd8be3f15dfd77c8d7af59e1bb7cac4cdc40c893e9): complete subsection reference.

<a id="canonical-3919fe9f225f330805d6bf40912b0decfe319d4abd668a33a1d120566c4d236d"></a>

<a id="canonical-165a05310a0f16fd08f4a47dc9e43f0ccb93ac259dbe63cf2077685e5347b603"></a>

## username property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5a70dea8e316 / 18

Type: `"string"`. Optional.

Username. Username to connect to the cluster/SVM.

Upstream description:

Username to connect to the cluster/SVM.

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
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [volume_defaults](resources--voltstack_site--reference--group-007.md#canonical-571b06ed2c72749baa5e3ab02d7db10f83d9d762b9cc97ccd21fa4fc922d4e6a): complete subsection reference.

<a id="canonical-c6790215d895d8fa25ee129d6ebdd88dfc523647934c86e4dfe8314ec00535f7"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 5a70dea8e316 / 19

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--voltstack_site--reference--group-006.md#canonical-bc80dd9d42d6825d0ec6d1e650739edaa480568a7018f14a7813e8c45d0a36fa)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap](resources--voltstack_site--reference--group-007.md#canonical-ddc300ee675f6654917a2597aaa201b4f5df02043b47315e45e3081cd991d4c8)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--voltstack_site--reference--group-007.md#canonical-9edae627e94b200f98ab5d8f9076889619e850d6fcf6336daa0d1f7b9dccaf6d)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](resources--voltstack_site--reference--group-007.md#canonical-8987b76c1b3b57fede2c84f453c17907fe933fe1ebcd9a0bea1ee6534c5da46a)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--reference--group-007.md#canonical-72930290386fbf5a47dda8bd8be3f15dfd77c8d7af59e1bb7cac4cdc40c893e9)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](resources--voltstack_site--reference--group-007.md#canonical-571b06ed2c72749baa5e3ab02d7db10f83d9d762b9cc97ccd21fa4fc922d4e6a)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-bc80dd9d42d6825d0ec6d1e650739edaa480568a7018f14a7813e8c45d0a36fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0150c6a6abda3f9a745299a3414a8c216d49ab644b25447cd813302bf28ea3a4"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0bffb31d552a / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key

<a id="canonical-420d653e52a52c83408e939630a465e04458e0a47db4a89a6ac90fb443b6e0c3"></a>

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
client_private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-0dadc72dc32263221f962827b107ee7677e4cee1136732029f355dbefc55e021"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0bffb31d552a / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-006.md#canonical-c39ae4961d00c777174a4aeae7fbf39eb86a03ff91935ddaa549f0ebb3c41c45): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-006.md#canonical-6e3329904c994c85f1f209820d12c38fd5b2d986b376c0e8f8b334f9808294d2): complete subsection reference.

<a id="canonical-5b817201a4f7e06910039bae74aca5f5d08e911ecae8914c9e701e7efb48510e"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 0bffb31d552a / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info](resources--voltstack_site--reference--group-006.md#canonical-c39ae4961d00c777174a4aeae7fbf39eb86a03ff91935ddaa549f0ebb3c41c45)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info](resources--voltstack_site--reference--group-006.md#canonical-6e3329904c994c85f1f209820d12c38fd5b2d986b376c0e8f8b334f9808294d2)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-c39ae4961d00c777174a4aeae7fbf39eb86a03ff91935ddaa549f0ebb3c41c45"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f31c040c5d8a8585b279c6a0869255963d9682083f77d68ea5c9f79bc277c30e"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 7f19be01308d / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-5c3fbbbbb6e0e498ac091fef1a6bfafc186f2c323883fb2cf7e510f312242af3)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-c060815b4ab3fde701039359636a3b27bc82a9a3ff810c8c5508e88a5d7e0679)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-5d4123d6819a2cd6d689bfb3087d2c8283333211bfb0c4c92c098a21358777bb)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-b84c8fd249287c751da748ad56a43c4930e50e9f2150d940fb34e1375474c21b)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-631abb4ba0a2d2f35af76b0131e8a361ca50b9eef89361ed9dee41d900b74a04)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--voltstack_site--reference--group-006.md#canonical-bc80dd9d42d6825d0ec6d1e650739edaa480568a7018f14a7813e8c45d0a36fa)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info

<a id="canonical-0df9eea2c099a7c19fc92d5edc887022a3fbe7469832c7ae0eed8db8baafbd5a"></a>

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

<a id="canonical-624da5d2f079befc4eb4d64ef2ec25c5cf8519c1a795f5a8627bd847675514f8"></a>

## Direct properties — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 7f19be01308d / 3

<a id="canonical-9ea08441f176743cb43b8e46585f72a13192be79c691115ae57da57bcbee95e1"></a>

<a id="canonical-dda6f1f8ea3d10ca76da8b6de7016b3a7812077d3e84e87e191a323d04656a96"></a>

## decryption_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 7f19be01308d / 4

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

<a id="canonical-303854303129bbfe51e03dd0b0ad51de93f62bc7eadae4c6b8c0ea31d621c5da"></a>

<a id="canonical-8d9b8c241461cd1c97f5298743528b5551c24e757e3e6a8b4cad8d7b480f7551"></a>

## location property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 7f19be01308d / 5

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

<a id="canonical-1f8cc5c32b52f5d3c0ed4d417a34365ca3553385840e17d5b3f15453008e1957"></a>

<a id="canonical-7ea8ce3a9babc8d1b00abeaeaf022c73994d5d4b63bda6c87dbb170a2058a172"></a>

## store_provider property — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 7f19be01308d / 6

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

<a id="canonical-178ba382cea52e66fb9595b2068df8b74fe1c515cec4c8c70630566b37cd071b"></a>

## Next pages — custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_ / 7f19be01308d / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--voltstack_site--reference--group-006.md#canonical-bc80dd9d42d6825d0ec6d1e650739edaa480568a7018f14a7813e8c45d0a36fa)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-6e3329904c994c85f1f209820d12c38fd5b2d986b376c0e8f8b334f9808294d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
