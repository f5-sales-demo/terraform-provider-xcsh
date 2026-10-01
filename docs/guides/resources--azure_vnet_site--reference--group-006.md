---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-2eb2da4db85c9049d69d5fcbd9891cab1e860f9ad3cba6d8d820bb87d8a14560"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections — ingress_egress_gw_ar.hub.express_route_enabled.connections / 8c399a1467c7 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- ingress_egress_gw_ar.hub.express_route_enabled.connections

<a id="canonical-0e279f1dd1960d507d2a4f74dcd082d61a9e48f3ccf3447932f78eb0ef8fa201"></a>

Type: `"object"`. list nested block, Optional.

Add the ExpressRoute Circuit Connections to this site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("circuit_id",
    "other_subscription")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
connections {
  # Configure direct properties listed below.
}
```

<a id="canonical-06c62bc667f8cbdc7d4fd89df12a5b58d3cfe8538cca621346d708b46e2e2b6d"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.connections / 8c399a1467c7 / 3

<a id="canonical-d1fc35ec806930f652a91a2238391e4c0be62c8b8e62b15fd2cb1d32bf08cfdd"></a>

<a id="canonical-53aca3b599ea04d99f0f48e40562361524c55ac9912faf3778d73226eb6fcfc9"></a>

## circuit_id property — ingress_egress_gw_ar.hub.express_route_enabled.connections / 8c399a1467c7 / 4

Type: `"string"`. Optional.

Exclusive with \[other\_subscription\] ExpressRoute Circuit is in same subscription as the site.

Upstream description:

Exclusive with \[other\_subscription\] ExpressRoute Circuit is in same subscription as the site.

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

- [metadata](resources--azure_vnet_site--reference--group-006.md#canonical-d7d10c1eade06334bb3dfc750cd8d3b73ab586f0b3ff33d8dc4632585f8cf2e3): complete subsection reference.

- [other_subscription](resources--azure_vnet_site--reference--group-006.md#canonical-47f7c0fac6968977cbad6cb62ec4a27737580f51e260982b2381f35a4e7415c7): complete subsection reference.

<a id="canonical-fe1d9551265d5b4468d2597fd6923c9231da03867a2c4d94574e0a7c8f077fd1"></a>

<a id="canonical-f5fe985367077151ec68fe723a9650515ed9c9fa153db5412628564055a26e6a"></a>

## weight property — ingress_egress_gw_ar.hub.express_route_enabled.connections / 8c399a1467c7 / 5

Type: `"number"`. Optional.

The weight (or priority) for the routes received from this connection. The. Defaults to \`10\`.

Upstream description:

The weight (or priority) for the routes received from this connection. The default value is 10.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-96ebf953b0fd95ef403ff5c533fbfd0ba1c6b27ae38a33276b11379c05f278cf"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.connections / 8c399a1467c7 / 6

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata](resources--azure_vnet_site--reference--group-006.md#canonical-d7d10c1eade06334bb3dfc750cd8d3b73ab586f0b3ff33d8dc4632585f8cf2e3)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](resources--azure_vnet_site--reference--group-006.md#canonical-47f7c0fac6968977cbad6cb62ec4a27737580f51e260982b2381f35a4e7415c7)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-d7d10c1eade06334bb3dfc750cd8d3b73ab586f0b3ff33d8dc4632585f8cf2e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce3e6a6ae6cb449086812c2c5527ee6710f8750435bfd76f2dad92743e93c061"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata — ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata / cbb2bc52cb89 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-005.md#canonical-fcdf88fba2fd9f53af8b5af69def1d59f4eadf9f27bc096dac177d6f5c05bf93)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata

<a id="canonical-2a6c5d38bd687a78bdc97e08d5a9c19c5d40ac350f12a645ad650d972d020946"></a>

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

<a id="canonical-e42df6a44479b513468a19928480a2c6aedb5b5791466cfaad620271a2a9f329"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata / cbb2bc52cb89 / 3

<a id="canonical-be96f1ed07bdeeeb9894bfbc66148db66cead0396a163f0903cbfcd463d36978"></a>

<a id="canonical-67f6bc51530d7d078573c0fc15e91734ca67a8405b671a8e1e40c296b7d5c166"></a>

## description_spec property — ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata / cbb2bc52cb89 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-7839e6ae83882d8c501ec6ea26eb7a85beb0dbe63fc2f93b7edf18fa32560486"></a>

<a id="canonical-6dfa68bb58d647ad6699cddacfb3c3ade0a08342f33475d8debc0c3a7c5efe45"></a>

## name property — ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata / cbb2bc52cb89 / 5

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

<a id="canonical-914e1ace56bc0487bbe0e71dfb66792160ac88a6539f0eb3fbdf209287278505"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata / cbb2bc52cb89 / 6

- [ingress_egress_gw_ar.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-005.md#canonical-fcdf88fba2fd9f53af8b5af69def1d59f4eadf9f27bc096dac177d6f5c05bf93)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-47f7c0fac6968977cbad6cb62ec4a27737580f51e260982b2381f35a4e7415c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab76bd2772dc36e53e5b6bfb18833d44afff048ac17f6c693ed900cea83d4355"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription / aa448231d14b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-005.md#canonical-fcdf88fba2fd9f53af8b5af69def1d59f4eadf9f27bc096dac177d6f5c05bf93)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription

<a id="canonical-2f22655c65a65a410761e4b0f63e001860c8f586e30a7174996e3bd24c25c057"></a>

Type: `"object"`. single nested block, Optional.

Express Route Circuit Config From Other Subscription.

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
other_subscription {
  # Configure direct properties listed below.
}
```

<a id="canonical-93340d934259283998c5d6ccc3ba8f231065d72589a6597002f012994353bc30"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription / aa448231d14b / 3

- [authorized_key](resources--azure_vnet_site--reference--group-006.md#canonical-ee0fb6b26c013c7bf15a417eb47673fff580da69e2353a60c7d23a2e6778eee3): complete subsection reference.

<a id="canonical-54b0e3b244632a187dd7371154ee541d0855751da43617b07f4bbc6bcde34c89"></a>

<a id="canonical-73e3458e9f72e5a10d640668ccf0401add045a8f26b5c35d618b9dfbfbf01854"></a>

## circuit_id property — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription / aa448231d14b / 4

Type: `"string"`. Optional.

Circuit ID. Circuit ID.

Upstream description:

Circuit ID.

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

<a id="canonical-af0ba73ddba90527c872db2ada76a353fdf0bd012abebc58493eafdac4dd91a0"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription / aa448231d14b / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](resources--azure_vnet_site--reference--group-006.md#canonical-ee0fb6b26c013c7bf15a417eb47673fff580da69e2353a60c7d23a2e6778eee3)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-005.md#canonical-fcdf88fba2fd9f53af8b5af69def1d59f4eadf9f27bc096dac177d6f5c05bf93)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-ee0fb6b26c013c7bf15a417eb47673fff580da69e2353a60c7d23a2e6778eee3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b39284286d0f993b47cc64523734ee941f0645a8989bb578a724391fa092ed92"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / cb3159b9c768 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-005.md#canonical-fcdf88fba2fd9f53af8b5af69def1d59f4eadf9f27bc096dac177d6f5c05bf93)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](resources--azure_vnet_site--reference--group-006.md#canonical-47f7c0fac6968977cbad6cb62ec4a27737580f51e260982b2381f35a4e7415c7)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key

<a id="canonical-f5e8d07e8b64b4e7158ac4be6784ebef88c8bb2899e71e3ede0ce77d434a5ea4"></a>

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
authorized_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-fc73fd81d8a6d6b6c6a9f651b9771a8135161082ad15583a217ee50ecb470d42"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / cb3159b9c768 / 3

- [blindfold_secret_info](resources--azure_vnet_site--reference--group-006.md#canonical-a5172b728c28986576fe6bc6da3f2b154fa45c6095135b9e7e0403d5cfd20c26): complete subsection reference.

- [clear_secret_info](resources--azure_vnet_site--reference--group-006.md#canonical-9dd7fc6c1c5a0fb4400ad751680c246e0ca6e022b0403c08c8a978e3d47261c4): complete subsection reference.

<a id="canonical-2bb76896e7db1e54a851dcca058f1c1d7c1ad27e60560f0783091305f4e1e1a2"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / cb3159b9c768 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info](resources--azure_vnet_site--reference--group-006.md#canonical-a5172b728c28986576fe6bc6da3f2b154fa45c6095135b9e7e0403d5cfd20c26)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info](resources--azure_vnet_site--reference--group-006.md#canonical-9dd7fc6c1c5a0fb4400ad751680c246e0ca6e022b0403c08c8a978e3d47261c4)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](resources--azure_vnet_site--reference--group-006.md#canonical-47f7c0fac6968977cbad6cb62ec4a27737580f51e260982b2381f35a4e7415c7)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-a5172b728c28986576fe6bc6da3f2b154fa45c6095135b9e7e0403d5cfd20c26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b318885c3a48a648975aa40338aabf4bb586f1e6eb6d1506b42ef4ceaf1a15b4"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / ba41efc842f4 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-005.md#canonical-fcdf88fba2fd9f53af8b5af69def1d59f4eadf9f27bc096dac177d6f5c05bf93)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](resources--azure_vnet_site--reference--group-006.md#canonical-47f7c0fac6968977cbad6cb62ec4a27737580f51e260982b2381f35a4e7415c7)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](resources--azure_vnet_site--reference--group-006.md#canonical-ee0fb6b26c013c7bf15a417eb47673fff580da69e2353a60c7d23a2e6778eee3)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info

<a id="canonical-28b4f0f09f705883113892d8b719ba9bad33b3222b1005b777a04c5fb17be9c1"></a>

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

<a id="canonical-49a109d2cf2fe5c9ff1a2b22e0ea20e730110d9c0f03729b34132a60158c85d2"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / ba41efc842f4 / 3

<a id="canonical-33d2ce7c8ac5eafca8486196f1a3170335a0739a3e28ed57c3e3ba786055412f"></a>

<a id="canonical-98900d38a7a2cad3854688b333d8ca814b047acf88e1cfb3395c5a1df7c9bf9f"></a>

## decryption_provider property — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / ba41efc842f4 / 4

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

<a id="canonical-144c7576aba03072369dd363d1b3805c5351d0ea2f2acea5d4111ce9d3e3bbed"></a>

<a id="canonical-cc73d58f699e23185861df15261cfa393853931928361ba85d075053abde84e5"></a>

## location property — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / ba41efc842f4 / 5

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

<a id="canonical-3bd896f02738c99576d67d84d8386171137f0713eb243eeb46991322d28f2e03"></a>

<a id="canonical-e66a4554e427f0e6c17859717779f3862aa9a3b8271b7778c89b4d8c6245ecf2"></a>

## store_provider property — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / ba41efc842f4 / 6

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

<a id="canonical-00f869b85ce279c19f25099970f2f5f19a45279d983fd6a16c43559f75edf9ad"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / ba41efc842f4 / 7

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](resources--azure_vnet_site--reference--group-006.md#canonical-ee0fb6b26c013c7bf15a417eb47673fff580da69e2353a60c7d23a2e6778eee3)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-9dd7fc6c1c5a0fb4400ad751680c246e0ca6e022b0403c08c8a978e3d47261c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a29b594efab0aaacdef5b29ee2be44f713259cec8d807391d584e2bb23c15b61"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / f2e44393bc16 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-005.md#canonical-fcdf88fba2fd9f53af8b5af69def1d59f4eadf9f27bc096dac177d6f5c05bf93)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](resources--azure_vnet_site--reference--group-006.md#canonical-47f7c0fac6968977cbad6cb62ec4a27737580f51e260982b2381f35a4e7415c7)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](resources--azure_vnet_site--reference--group-006.md#canonical-ee0fb6b26c013c7bf15a417eb47673fff580da69e2353a60c7d23a2e6778eee3)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info

<a id="canonical-ed7cadd2a86a593e50154e76251910cbb26d72c84f71b05374dab865d5d89504"></a>

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

<a id="canonical-5cad0d089e9269c3e3bd94a7cfca220073f32245e8564d314abda3e2d0e15602"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / f2e44393bc16 / 3

<a id="canonical-1622ec453292786fc0def9e7a106cf0b2ef3c2aa9c59752a74fc5b10655d7d7c"></a>

<a id="canonical-48c60e07fe754d1a1ffb056bf419a8b1e7d959700d3d20f437d2011cc9d10c9d"></a>

## provider_ref property — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / f2e44393bc16 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-f481137201d0c79879fe2c27915002e0e9d04f33e96f05afac159b159e895533"></a>

<a id="canonical-4331c1e2f9219d3dd345a2219586a93e3cfabde934e6343a5ffcad41a7339f27"></a>

## url property — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / f2e44393bc16 / 5

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

<a id="canonical-e769d1e2809e460ee173ad4615788bf088611e7264a0b202a8df1361a4cf77f7"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / f2e44393bc16 / 6

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](resources--azure_vnet_site--reference--group-006.md#canonical-ee0fb6b26c013c7bf15a417eb47673fff580da69e2353a60c7d23a2e6778eee3)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-08064fd4d5b38fbe45aeafeab03fb56140e7fc9ef75ecf45d4cbf3acc3c399b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-481c2f9428f570a336e9876960aaa0aab5fef29b94ed59212a1d76f616d59827"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server — ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server / 57c0dba27560 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server

<a id="canonical-70baf1e9179b182dc6b6dfbcfc63c07f02a78335eb3c74966717b9c08eeb33a3"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for do not advertise to route server.

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
do_not_advertise_to_route_server = {}
```

<a id="canonical-4f26ef3b45fbc639679a974664da0b04d87e79be8e3f635072eb866f393308c3"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server / 57c0dba27560 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3f77a07ac36acd6f667948fa3a276b70031835c1b25bc413bd5f19d1c51e76aa"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server / 57c0dba27560 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-e7406a16de739030a0cf19472d9037dcf845eaf6d76d69ef183a4db0e09a0219"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94ca6cc6929fcd7edd336701c002bdae5922c51bd326fb1bee3e7a46d90819db"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet / 25c52460aaac / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet

<a id="canonical-c279d8ea5f5df216e2205dd31329262247e2bd56965af8505b15c1662184926b"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for gateway subnet.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto",
    "subnet"),
  validators.ConflictingObjectAttributes("auto",
    "subnet_param"),
  validators.ConflictingObjectAttributes("subnet",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"auto\",\"subnet\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
gateway_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-2ee8ca20b8efb8f8d0c2c1e6d81648b2fcdb5c7a59067f70d3e0bdbe3b75b72f"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet / 25c52460aaac / 3

- [auto](resources--azure_vnet_site--reference--group-006.md#canonical-f54c4f88ca8abfbc79e6fbeacda52b4a943d1c5b0dccfb667c88f7633f1795d5): complete subsection reference.

- [subnet](resources--azure_vnet_site--reference--group-006.md#canonical-9ee08c49a7cbbf5840190d2f35dc35fd1fbd2c59920208fe06af666659cd932a): complete subsection reference.

- [subnet_param](resources--azure_vnet_site--reference--group-006.md#canonical-9ea4922848c860de0374374440287ff7a0309ea68ae3be7e75b87bf7c53f7349): complete subsection reference.

<a id="canonical-5ca9c1c629563320bcce354195714123d507c66a7b08eee4ad6ee09814fccc39"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet / 25c52460aaac / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto](resources--azure_vnet_site--reference--group-006.md#canonical-f54c4f88ca8abfbc79e6fbeacda52b4a943d1c5b0dccfb667c88f7633f1795d5)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet](resources--azure_vnet_site--reference--group-006.md#canonical-9ee08c49a7cbbf5840190d2f35dc35fd1fbd2c59920208fe06af666659cd932a)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param](resources--azure_vnet_site--reference--group-006.md#canonical-9ea4922848c860de0374374440287ff7a0309ea68ae3be7e75b87bf7c53f7349)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-f54c4f88ca8abfbc79e6fbeacda52b4a943d1c5b0dccfb667c88f7633f1795d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f6561aaa934d2591d2557cc3f5bb132e4b918b2f3e180427e678b4a8ff59f47"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto / 141b38876389 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-e7406a16de739030a0cf19472d9037dcf845eaf6d76d69ef183a4db0e09a0219)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto

<a id="canonical-a1276cf00eaf9cf71c520fae102c10b8e349ca9df5cb42e7f047384a411b6e8d"></a>

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
auto = {}
```

<a id="canonical-12f8508e0a84acfc953a949534bac1f14a9d65ea3e95290e56d5fbae9f22237c"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto / 141b38876389 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8f8e867ecb8ce3e8b99c4d655f9321da36b5aec0758637f15e7c59bc853514bf"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto / 141b38876389 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-e7406a16de739030a0cf19472d9037dcf845eaf6d76d69ef183a4db0e09a0219)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-9ee08c49a7cbbf5840190d2f35dc35fd1fbd2c59920208fe06af666659cd932a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9afa2070538ae7064597d42032657c62b00bf7a21e75b8a5a1d7379c5c6db44"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet / 1509c41c2382 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-e7406a16de739030a0cf19472d9037dcf845eaf6d76d69ef183a4db0e09a0219)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet

<a id="canonical-3df552f8d78c63042b2e32e9b381d291ccd525ff5f73828a1b1d636c548b3d1e"></a>

Type: `"object"`. single nested block, Optional.

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or
RouteServerSubnet).

Upstream description:

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or RouteServerSubnet)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("subnet_resource_grp",
    "vnet_resource_group")}
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
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

Terraform syntax:

```terraform
subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-09716258f76a8737fd63f853e9d4af3d196a9196a14248e62f07032910f5cf26"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet / 1509c41c2382 / 3

<a id="canonical-158ed7196be62522f1710a6dc50ed3b6ddb32b234ee202adf83b2c3bf47e8026"></a>

<a id="canonical-8ba9b229acbdf3db3fc3fe43210f467777250357fd23fdd5bc59fe8ea4040708"></a>

## subnet_resource_grp property — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet / 1509c41c2382 / 4

Type: `"string"`. Optional.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

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

- [vnet_resource_group](resources--azure_vnet_site--reference--group-006.md#canonical-b27f65da9376c40e8ae2115f6feb743375fb6400703283e3b50e2515940c10ff): complete subsection reference.

<a id="canonical-c792bdd5293e793f9f5f60cd8120ea928d0b978f9c111a7eda45884c7a30cbac"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet / 1509c41c2382 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group](resources--azure_vnet_site--reference--group-006.md#canonical-b27f65da9376c40e8ae2115f6feb743375fb6400703283e3b50e2515940c10ff)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-e7406a16de739030a0cf19472d9037dcf845eaf6d76d69ef183a4db0e09a0219)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-b27f65da9376c40e8ae2115f6feb743375fb6400703283e3b50e2515940c10ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd8d8f62fefdfe7070aa9b0947e11b8539d4187c78e55c4e24e9e01cfe8605c3"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resour / b1e69e0ff3ce / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-e7406a16de739030a0cf19472d9037dcf845eaf6d76d69ef183a4db0e09a0219)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet](resources--azure_vnet_site--reference--group-006.md#canonical-9ee08c49a7cbbf5840190d2f35dc35fd1fbd2c59920208fe06af666659cd932a)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group

<a id="canonical-de328d4a756315c7ff26067e1b292ea9c5548456837ace276e5eafdd28267036"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for vnet resource group.

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
vnet_resource_group = {}
```

<a id="canonical-e85d64235aec934b4d91a7680852fba5c2b533da0fe13cc9049b77d57ee4ebff"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resour / b1e69e0ff3ce / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b86ff8cc16ad579c46813e48c83e1d4fc5d70b2b94df87a7378a380bd7386f70"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resour / b1e69e0ff3ce / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet](resources--azure_vnet_site--reference--group-006.md#canonical-9ee08c49a7cbbf5840190d2f35dc35fd1fbd2c59920208fe06af666659cd932a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-9ea4922848c860de0374374440287ff7a0309ea68ae3be7e75b87bf7c53f7349"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d8e610120c12350c31f9892d37b71ed29e6a8189f263a259307f2ee554853ef"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param / d21293839a9e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-e7406a16de739030a0cf19472d9037dcf845eaf6d76d69ef183a4db0e09a0219)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param

<a id="canonical-96f3d0681476f7652114ed436cc80f66f6acbdd2f82f1ccfd58d5762c93e5eda"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-05b0fa44fcb8b1f9e58096f8bccf2b5bbafa9beda7d249697765a560e4098fb3"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param / d21293839a9e / 3

<a id="canonical-9ac0cadc810121beeea26a22240b679addb0757f4fcfd5e977802509a2a5c001"></a>

<a id="canonical-601f08db79fe1e920aa4fdf722c8fd386c09dd30fe8f5ef9f0fa3363f2a8cd45"></a>

## ipv4 property — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param / d21293839a9e / 4

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-60a05d611f5033aead2eff19196b80bdb7e7689f5d7b04baa11be772709bfacb"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param / d21293839a9e / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-e7406a16de739030a0cf19472d9037dcf845eaf6d76d69ef183a4db0e09a0219)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-c55b5ed2085f9e5b41dcc160c63e37f695add2cb74be66c808f977126f7ea3cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5ef2599f55ef81950ca98336be3e5466b4afb611dfb860ddd13a3746c0c6386"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet / 8a3dd212bb2a / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet

<a id="canonical-c6ba4a2d788140b30b9008c1a834fc65d9ea7a48cdbd5448fe2be80c3cce5612"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for route server subnet.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto",
    "subnet"),
  validators.ConflictingObjectAttributes("auto",
    "subnet_param"),
  validators.ConflictingObjectAttributes("subnet",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"auto\",\"subnet\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
route_server_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-a0498562169b318c8407c87e9ca6759aef9b70e03a82b3957349e5bf5736c05d"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet / 8a3dd212bb2a / 3

- [auto](resources--azure_vnet_site--reference--group-006.md#canonical-e1fb55ac2220ab58ff7e94d9f3a4520c1988cc029f5144e0fa3edf15eefe39a0): complete subsection reference.

- [subnet](resources--azure_vnet_site--reference--group-006.md#canonical-8aa6e4e62c93925ef136e57bba08921ea38fa5152bff67446d2e83bf28372732): complete subsection reference.

- [subnet_param](resources--azure_vnet_site--reference--group-006.md#canonical-06cbe0ba5efd7a9547b6210f1a1188b2533377efa13bb4d246b231c93633bd6e): complete subsection reference.

<a id="canonical-2e1a83c2134f18eeca7cb77a1f7a84dec477b9422a150a4483ac3882113e7aab"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet / 8a3dd212bb2a / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto](resources--azure_vnet_site--reference--group-006.md#canonical-e1fb55ac2220ab58ff7e94d9f3a4520c1988cc029f5144e0fa3edf15eefe39a0)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet](resources--azure_vnet_site--reference--group-006.md#canonical-8aa6e4e62c93925ef136e57bba08921ea38fa5152bff67446d2e83bf28372732)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param](resources--azure_vnet_site--reference--group-006.md#canonical-06cbe0ba5efd7a9547b6210f1a1188b2533377efa13bb4d246b231c93633bd6e)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-e1fb55ac2220ab58ff7e94d9f3a4520c1988cc029f5144e0fa3edf15eefe39a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d6323bc2f6760c00a7d8b28612042d71e309b15f5e8a7bb605b350e458f3669"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto / 209347507854 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-c55b5ed2085f9e5b41dcc160c63e37f695add2cb74be66c808f977126f7ea3cd)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto

<a id="canonical-ed2ead2abf8299313f4791772c68595ac2f03b022cc0c4ddb245cf155a633979"></a>

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
auto = {}
```

<a id="canonical-fe66a6528f687fbe6c0a279fb76d9857b6707cd667ed09da2f1caa9cf753967e"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto / 209347507854 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3bda144dfa1318af6a00b68a84e2a9af5d0c5fa756c388c9e0e4de4f64fa9eb2"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto / 209347507854 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-c55b5ed2085f9e5b41dcc160c63e37f695add2cb74be66c808f977126f7ea3cd)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-8aa6e4e62c93925ef136e57bba08921ea38fa5152bff67446d2e83bf28372732"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac2521df9bdae1b042c49db71e8374c55dc7d8fa7a3f02822e8f69f27dfacae8"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet / 8e76f43f31cd / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-c55b5ed2085f9e5b41dcc160c63e37f695add2cb74be66c808f977126f7ea3cd)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet

<a id="canonical-8ce2a0a74091abf7f98f1e7b2975aa63ad8183dd550884e2d78029c5140d9c59"></a>

Type: `"object"`. single nested block, Optional.

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or
RouteServerSubnet).

Upstream description:

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or RouteServerSubnet)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("subnet_resource_grp",
    "vnet_resource_group")}
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
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

Terraform syntax:

```terraform
subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-79a35a8a98be616ae0155a9acb04304d0438e226b673b931130f4e6330fa247a"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet / 8e76f43f31cd / 3

<a id="canonical-c4d48b3a9ea630b39e76f0c29a099b88f235848630276604def67b4759d3e326"></a>

<a id="canonical-048a942f8cf947cd2651a6a0a366be08306c8b6471987e7048aa02dc3b21df9c"></a>

## subnet_resource_grp property — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet / 8e76f43f31cd / 4

Type: `"string"`. Optional.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

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

- [vnet_resource_group](resources--azure_vnet_site--reference--group-006.md#canonical-f3a850c15e41bcec60abb11bed58d0a55c778c581de8e268c0cb6b9ac060edfa): complete subsection reference.

<a id="canonical-0bbe7ed25032611103aacbaa254b89ff6a2a38af78633332604db16a84157ee2"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet / 8e76f43f31cd / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group](resources--azure_vnet_site--reference--group-006.md#canonical-f3a850c15e41bcec60abb11bed58d0a55c778c581de8e268c0cb6b9ac060edfa)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-c55b5ed2085f9e5b41dcc160c63e37f695add2cb74be66c808f977126f7ea3cd)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-f3a850c15e41bcec60abb11bed58d0a55c778c581de8e268c0cb6b9ac060edfa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d20577e3caf76a0ce7a26c986bef198d46284b9367326c2019bf132a44c7005"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_r / f85e05728eba / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-c55b5ed2085f9e5b41dcc160c63e37f695add2cb74be66c808f977126f7ea3cd)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet](resources--azure_vnet_site--reference--group-006.md#canonical-8aa6e4e62c93925ef136e57bba08921ea38fa5152bff67446d2e83bf28372732)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group

<a id="canonical-f70b0e5f7c68f2c1c7ed0c1e09fa612098471609f9e83a7e9f8d95cf09d6594d"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for vnet resource group.

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
vnet_resource_group = {}
```

<a id="canonical-8e13e29ec577bc178343cd981e98899cadf78bf2d3d10c1fa92d763b89f77172"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_r / f85e05728eba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-16fd92434a6ba1ea27a6847028c44bf84e94de294f0db80e3de8d25b043f39e3"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_r / f85e05728eba / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet](resources--azure_vnet_site--reference--group-006.md#canonical-8aa6e4e62c93925ef136e57bba08921ea38fa5152bff67446d2e83bf28372732)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-06cbe0ba5efd7a9547b6210f1a1188b2533377efa13bb4d246b231c93633bd6e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a3db8013c3da2b7064329b41a0c9cb5c8c71cc98cbc043bc40b1b679c9d376c"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param / a75439cd5721 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-c55b5ed2085f9e5b41dcc160c63e37f695add2cb74be66c808f977126f7ea3cd)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param

<a id="canonical-09ac06d622fff4cf053ca77e49c2994c42415a0b406bf8bdcb88998eff997a7b"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-d44dcdbbd468995059efac0dbf8960f48e95d08616d67b161609c40280aca9cd"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param / a75439cd5721 / 3

<a id="canonical-75e3f174e4e532b189028dd4f0f3695652090a6ae194ce2a28843df1d5f55edb"></a>

<a id="canonical-fa53b038553563e921fd086a8c9b40e20e6005219a1c14a3e222ac6a5064e37e"></a>

## ipv4 property — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param / a75439cd5721 / 4

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-372674da4e8768cdfda48f8228107aa64b1eca946e0a686c13ce35c49fca5ee6"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param / a75439cd5721 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-c55b5ed2085f9e5b41dcc160c63e37f695add2cb74be66c808f977126f7ea3cd)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-1f85c3c59439548e46dd657bce79e1340e9dd6d7db8ac7940d56f21a43a06f94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6b5ffffc6b273baed087ffb555ee8c7d571c5df1a24f21f9bc6fd13acc8001a"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route — ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_ro / cfa829a33070 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route

<a id="canonical-836a9902de69730afe0ec595854238366e1d947bb6a6e776aeeb66c0428794f3"></a>

Type: `"object"`. single nested block, Optional.

CloudLink ADN Network Config.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cloudlink_network_name")}
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
site_registration_over_express_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-b17c9488e208ba2ab91b5f8a525bcff1d7644d25b95fb8a6040fec86eb304968"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_ro / cfa829a33070 / 3

<a id="canonical-e11ef88e661eba7323d41d2da20abf0d3f60ffaacde10b9e78c62efa4b702a2e"></a>

<a id="canonical-b91f398e93433b22729cc1021134060770e4ea0170c1e8fcc8734eac0588ef2e"></a>

## cloudlink_network_name property — ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_ro / cfa829a33070 / 4

Type: `"string"`. Optional.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-aee804f50a9e2d9f111929904ede8fb890008a27a51ef6951721997fbb762041"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_ro / cfa829a33070 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-3f42cb467712abca4dfe5ff8a68d1334df921aa3a01af66ab5a00e69fa78a31f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e5f240a1dfa657ca0ae1aa35d17712d360deb917b046cfa4b2936ec684c952e"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet — ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet / 2888674557f8 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet

<a id="canonical-9a9d46728a60aa671833bbd8cf8b745b0dbb10ecad3a0ecc19c51a73b854faf9"></a>

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
site_registration_over_internet = {}
```

<a id="canonical-ca1c3a31fd8a681ceee2d5c67d9d0e437266561019bec48e8fadaa182c9fe216"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet / 2888674557f8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5db63a1f7de4bc7af86fde4e9d45b695114546ded29f067d3c916b6022fae995"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet / 2888674557f8 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-701512fd97c100e810c9dedbcc76e5fa8078c4853732272e77075918cb409a39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-155a60a696f903212c7534bc20c875cd8c21495367b4a55c0738814a3fe12144"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az — ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az / 5d929159bcb3 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az

<a id="canonical-212bc207aee89b5e2bdeb6b2348f3868bf1d8a625235ae36f8748074a119caf6"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for sku ergw1az.

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
sku_ergw1az = {}
```

<a id="canonical-41d9bd5dfe6e64e87483cd18396dedc75c1ba0d45a49d2c76ceca4bb437656dd"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az / 5d929159bcb3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-44f7bb5ea73a6831b91c7bcfa36a37a59c5754acc88c7e03673bbcbf205464f7"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az / 5d929159bcb3 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-803e25e21c265627f6e20f4b3acbe8803ebacdc0945a9ae747f5c051c3502582"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d0631aca8dc7dc378c430a6c9b97a7de4c1cc08d7456c7fb75a676190f99978"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az — ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az / 3bf65e0c87db / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az

<a id="canonical-65b38e0b976e51956c4afa96e7a8dff0ed68324e906b49276e3d028797c6f5bb"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for sku ergw2az.

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
sku_ergw2az = {}
```

<a id="canonical-65ee500fb0174351a5f6abbd5110e73d3823cfadfaccde2fe57b2b9dbdf21839"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az / 3bf65e0c87db / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1b9d6a120ad1d3694ea756420c5d3fa6f9e4f22e93006b20ad1e7ec5684de88f"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az / 3bf65e0c87db / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-450890e587b8316cd9f2649d6999aa0268b3ca8eb0ea5b1345f34de01d31edc8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-500c8e45cfd1a5abe4856c1db3d6cc6c8b3c66a8f2c551f8fac104d16c35bd38"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf — ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf / 486c9df15204 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf

<a id="canonical-5d15fa1806919be83db0ffebda5ffdb6eceb1b127c185c0719e212b5d958498c"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for sku high perf.

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
sku_high_perf = {}
```

<a id="canonical-ca0083f2135577a5de1d0975642dbb30326bf8108f36a8b6b08311ea1ceb2370"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf / 486c9df15204 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-74dae980ddf711f0b830fa80447bc6517e1cfe220e12090d07f350ee6fef6617"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf / 486c9df15204 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-374dea40b689927c70cc506099a6fa9a917e62876e218114c4b97b326a01f29a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8b7014d89c847ad4695ddc7c4fb29dd89541d5a6f9e27af13c7c55eeb2dd6f6"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.sku_standard — ingress_egress_gw_ar.hub.express_route_enabled.sku_standard / 0e7065d792f3 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- ingress_egress_gw_ar.hub.express_route_enabled.sku_standard

<a id="canonical-1a60094f76b34d9015233d465a27e64616954a1b76b05ae874c6aff98b44bbea"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for sku standard.

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
sku_standard = {}
```

<a id="canonical-284d82f46777aa9455a9557db42755cf1cc83ce61f896084f45318afe434cdd0"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.sku_standard / 0e7065d792f3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c679ffd90d5c41413ef0046d51af95bff393780509cf701a14ce2e97f645dc50"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.sku_standard / 0e7065d792f3 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-43a5d191b8eee55ba9b749e4293819d3e12e6b2dd09125fa56cbd5421215571b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-485cf052be52865d958e45938cd9560a9427813b5c1a017c979b337c3f54fded"></a>

## ingress_egress_gw_ar.hub.spoke_vnets — ingress_egress_gw_ar.hub.spoke_vnets / 9b3b0a5d71cd / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- ingress_egress_gw_ar.hub.spoke_vnets

<a id="canonical-e2e1ce5c87d1100ac5d2bcfc264b0807375cc75e67064e985e15b7fee06f7b76"></a>

Type: `"object"`. list nested block, Optional.

Spoke VNet Peering (Legacy). Spoke VNet Peering.

Upstream description:

Spoke VNet Peering.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("auto",
    "manual")}
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
spoke_vnets {
  # Configure direct properties listed below.
}
```

<a id="canonical-260778387fde3e433daf45d8df0b0255d9e184b67a2b22cbc604c6f1a9d25f7f"></a>

## Direct properties — ingress_egress_gw_ar.hub.spoke_vnets / 9b3b0a5d71cd / 3

- [auto](resources--azure_vnet_site--reference--group-006.md#canonical-6242b31e3914b3a330a4922e160d3707eda6fc6de728358638524bfd58d6e296): complete subsection reference.

- [labels](resources--azure_vnet_site--reference--group-006.md#canonical-90ab6d95fcf06ce32677e636986b179e14e9d66b16e0db0df0c3d24242718b5b): complete subsection reference.

- [manual](resources--azure_vnet_site--reference--group-006.md#canonical-e9bec934f67c8cd7ac03a95299cbdb0f86114f54aaf2e724e9b3f6ca20ab9b74): complete subsection reference.

- [vnet](resources--azure_vnet_site--reference--group-006.md#canonical-399719f6db0270ee460767d05740c73e5462c805cf2d2cfc3ff43e1691829c89): complete subsection reference.

<a id="canonical-de2c63a2d7192151462210e86d742560371bf58f55a6b174d0a3698a91f5d4cd"></a>

## Next pages — ingress_egress_gw_ar.hub.spoke_vnets / 9b3b0a5d71cd / 4

- [ingress_egress_gw_ar.hub.spoke_vnets.auto](resources--azure_vnet_site--reference--group-006.md#canonical-6242b31e3914b3a330a4922e160d3707eda6fc6de728358638524bfd58d6e296)
- [ingress_egress_gw_ar.hub.spoke_vnets.labels](resources--azure_vnet_site--reference--group-006.md#canonical-90ab6d95fcf06ce32677e636986b179e14e9d66b16e0db0df0c3d24242718b5b)
- [ingress_egress_gw_ar.hub.spoke_vnets.manual](resources--azure_vnet_site--reference--group-006.md#canonical-e9bec934f67c8cd7ac03a95299cbdb0f86114f54aaf2e724e9b3f6ca20ab9b74)
- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](resources--azure_vnet_site--reference--group-006.md#canonical-399719f6db0270ee460767d05740c73e5462c805cf2d2cfc3ff43e1691829c89)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-6242b31e3914b3a330a4922e160d3707eda6fc6de728358638524bfd58d6e296"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d398f4ca6e0b4aaceb852a732b7e93d4fb608331b275e919b9b4dbe4d5e5fff0"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.auto — ingress_egress_gw_ar.hub.spoke_vnets.auto / 3cc93a7939c6 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-43a5d191b8eee55ba9b749e4293819d3e12e6b2dd09125fa56cbd5421215571b)
- ingress_egress_gw_ar.hub.spoke_vnets.auto

<a id="canonical-0d2b408ae77a973dd02c7b2a5fa37477ffbea5eec830842e2a61ade50d738622"></a>

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
auto = {}
```

<a id="canonical-50bc8d2424c7ece944777ded6ad76c14644bfa9813ba07dbababf3814ac637e2"></a>

## Direct properties — ingress_egress_gw_ar.hub.spoke_vnets.auto / 3cc93a7939c6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-90e7b9e102bce7c3b2b862191e238b8fcdde8ad3af8bc1beb304b2a0ac75f93b"></a>

## Next pages — ingress_egress_gw_ar.hub.spoke_vnets.auto / 3cc93a7939c6 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-43a5d191b8eee55ba9b749e4293819d3e12e6b2dd09125fa56cbd5421215571b)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-90ab6d95fcf06ce32677e636986b179e14e9d66b16e0db0df0c3d24242718b5b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7cf952faef11ebdccebdb0c83afb8d65adfe4a266d810a02c6c1805f571efd17"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.labels — ingress_egress_gw_ar.hub.spoke_vnets.labels / 08db9c8a951b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-43a5d191b8eee55ba9b749e4293819d3e12e6b2dd09125fa56cbd5421215571b)
- ingress_egress_gw_ar.hub.spoke_vnets.labels

<a id="canonical-a31d99df1de376bfe4f9507ae80aeccd75a53892a84fd920394acd08caad0e72"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall
policy These labels used must be from known key and label defined in shared namespace.

Upstream description:

Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall
policy These labels used must be from known key and label defined in shared namespace.

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
labels {}
```

<a id="canonical-876f5034b43cd7bf15819bf28d67ef99640fac1cc084936c6f3fef077693431b"></a>

## Direct properties — ingress_egress_gw_ar.hub.spoke_vnets.labels / 08db9c8a951b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cd949018bfec86a59c73fa8c6895e7468d256e8398c4e94c6e894bb27e28b5ea"></a>

## Next pages — ingress_egress_gw_ar.hub.spoke_vnets.labels / 08db9c8a951b / 4

- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-43a5d191b8eee55ba9b749e4293819d3e12e6b2dd09125fa56cbd5421215571b)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-e9bec934f67c8cd7ac03a95299cbdb0f86114f54aaf2e724e9b3f6ca20ab9b74"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b3c79621b7106119191e6bd510889412a4e53b5159a10f3370579e02e765341"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.manual — ingress_egress_gw_ar.hub.spoke_vnets.manual / 673f976d93fb / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-43a5d191b8eee55ba9b749e4293819d3e12e6b2dd09125fa56cbd5421215571b)
- ingress_egress_gw_ar.hub.spoke_vnets.manual

<a id="canonical-9bf698d6d8297df60b910909dc70c5b207de6ded91d700759742b3383ea0726f"></a>

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
manual = {}
```

<a id="canonical-a1878872485f7b2a002b1f8184400f0f1b2a3acf9ff37cfbe1dbec5f2193dfee"></a>

## Direct properties — ingress_egress_gw_ar.hub.spoke_vnets.manual / 673f976d93fb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0a923667132d4d874e3455c542f6257bfcf0e86e1383721b1f92d5b8ad586bd7"></a>

## Next pages — ingress_egress_gw_ar.hub.spoke_vnets.manual / 673f976d93fb / 4

- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-43a5d191b8eee55ba9b749e4293819d3e12e6b2dd09125fa56cbd5421215571b)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-399719f6db0270ee460767d05740c73e5462c805cf2d2cfc3ff43e1691829c89"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1cc8362921a3b6a7ef55a7552a2c05543d44708b7b5a64dc1132f760da9c21e6"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.vnet — ingress_egress_gw_ar.hub.spoke_vnets.vnet / faa3c528225e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-43a5d191b8eee55ba9b749e4293819d3e12e6b2dd09125fa56cbd5421215571b)
- ingress_egress_gw_ar.hub.spoke_vnets.vnet

<a id="canonical-b1387411b5a9cffdd1b2418b84ca33d71a7eb62b0496b78198cb008770b53a7c"></a>

Type: `"object"`. single nested block, Optional.

Resource group and name of existing Azure VNet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("resource_group",
    "vnet_name"),
  validators.ConflictingObjectAttributes("f5_orchestrated_routing",
    "manual_routing")}
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
  "x-ves-oneof-field-routing_type": "[\"f5_orchestrated_routing\",\"manual_routing\"]"
}
```

Terraform syntax:

```terraform
vnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-6920fa1f4cd44aa446fcb9cad12da67e0fa46c2c716ad8bb9ccd02a801defb5c"></a>

## Direct properties — ingress_egress_gw_ar.hub.spoke_vnets.vnet / faa3c528225e / 3

- [f5_orchestrated_routing](resources--azure_vnet_site--reference--group-006.md#canonical-6f2651a24266e1b2ab65c48373e0ce3caea15d8b75c9674c7ff8ba82435011af): complete subsection reference.

- [manual_routing](resources--azure_vnet_site--reference--group-006.md#canonical-baa4b442daa2dc4eada12424487fd9b6624fdc44e677e6daf595493c380b27bd): complete subsection reference.

<a id="canonical-3868aceea6ac85614e2126c907f1d756d5517f13a10ba961c1ac48074afe58c3"></a>

<a id="canonical-2a6bb999965f1d017fd09dca37248f48cd2a596d9150abfe8634aedfb79277b6"></a>

## resource_group property — ingress_egress_gw_ar.hub.spoke_vnets.vnet / faa3c528225e / 4

Type: `"string"`. Optional.

Existing VNet Resource Group. Resource group of existing VNet.

Upstream description:

Resource group of existing VNet.

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

<a id="canonical-02262e7108994ee71ce1db4f2f174bbfdf3c0c9eb997f2456c35366b3c5391d8"></a>

<a id="canonical-666c470956750b43ddb0846369354c53be223b8833f7f6bccc06ed56becd76ea"></a>

## vnet_name property — ingress_egress_gw_ar.hub.spoke_vnets.vnet / faa3c528225e / 5

Type: `"string"`. Optional.

Existing VNet Name. Name of existing VNet.

Upstream description:

Name of existing VNet.

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

<a id="canonical-8e4e779c042f58b92994d1b68955c49914b44c7d993d041bec7c42faeba536e0"></a>

## Next pages — ingress_egress_gw_ar.hub.spoke_vnets.vnet / faa3c528225e / 6

- [ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing](resources--azure_vnet_site--reference--group-006.md#canonical-6f2651a24266e1b2ab65c48373e0ce3caea15d8b75c9674c7ff8ba82435011af)
- [ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing](resources--azure_vnet_site--reference--group-006.md#canonical-baa4b442daa2dc4eada12424487fd9b6624fdc44e677e6daf595493c380b27bd)
- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-43a5d191b8eee55ba9b749e4293819d3e12e6b2dd09125fa56cbd5421215571b)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-6f2651a24266e1b2ab65c48373e0ce3caea15d8b75c9674c7ff8ba82435011af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf414d44593dee3236cb432f35ac2ad1fae6df7473507da5bf2e639f3e8a2546"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing — ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing / a13bbbfc4c67 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-43a5d191b8eee55ba9b749e4293819d3e12e6b2dd09125fa56cbd5421215571b)
- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](resources--azure_vnet_site--reference--group-006.md#canonical-399719f6db0270ee460767d05740c73e5462c805cf2d2cfc3ff43e1691829c89)
- ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing

<a id="canonical-5b058314b6d64f8f2a9e671f40417402bab761158ec1f7a3dc1f425243e2f60b"></a>

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
f5_orchestrated_routing = {}
```

<a id="canonical-4659ae2b4934c5c39b3196f73230603278f416b1f0b2b695cf76116a0a63f8de"></a>

## Direct properties — ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing / a13bbbfc4c67 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-21ff18fe0c71f4add5c347e8566fbbf7d52c5bc55f9ee53a271b15049652ecf4"></a>

## Next pages — ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing / a13bbbfc4c67 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](resources--azure_vnet_site--reference--group-006.md#canonical-399719f6db0270ee460767d05740c73e5462c805cf2d2cfc3ff43e1691829c89)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-baa4b442daa2dc4eada12424487fd9b6624fdc44e677e6daf595493c380b27bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e273967dadba13a8a1be47391b774608f03ead78b4ecd995f484376006761e39"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing — ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing / 3a136ae83614 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-43a5d191b8eee55ba9b749e4293819d3e12e6b2dd09125fa56cbd5421215571b)
- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](resources--azure_vnet_site--reference--group-006.md#canonical-399719f6db0270ee460767d05740c73e5462c805cf2d2cfc3ff43e1691829c89)
- ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing

<a id="canonical-e0d2a4e4878be7263effbe40f52eec34316515a59520429e0f236bac28a2a3f9"></a>

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
manual_routing = {}
```

<a id="canonical-0747ced0a181ea0626fbd0a29dee3e4006459a33cb7b924b30a50d78e4d6729c"></a>

## Direct properties — ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing / 3a136ae83614 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-056ea3c2a7a3d03ab310e7b51bf5aa6360f704380b3a97e075fedbba7412886d"></a>

## Next pages — ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing / 3a136ae83614 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](resources--azure_vnet_site--reference--group-006.md#canonical-399719f6db0270ee460767d05740c73e5462c805cf2d2cfc3ff43e1691829c89)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-83daff971568976e1322aa005d7edac6fc071a1d6ad075cc47c167833c9f00a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9c175ed397c931fd5f139d380a891e16e69425d6919af6fc642df9910d74eb1"></a>

## ingress_egress_gw_ar.inside_static_routes — ingress_egress_gw_ar.inside_static_routes / de9153f67948 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.inside_static_routes

<a id="canonical-143a0d877e25568595b959d18db6422ef98a067f6deff39b5a1ce0a497f8827f"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inside static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_route_list")}
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
inside_static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-f9dc38080501f91ce77bb7bb1ec274ad8ee1bf4a33d15c47379e5ec860fbb15c"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes / de9153f67948 / 3

- [static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-bbb6ef38e4c455285c41df2c971ca0b35b311ff97db1241dc291d629c4575344): complete subsection reference.

<a id="canonical-0ba9e79444bfc1a08fa4ab8b4576d691e6c35d4f1e9bb559b56f31d86b0b8066"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes / de9153f67948 / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-bbb6ef38e4c455285c41df2c971ca0b35b311ff97db1241dc291d629c4575344)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-bbb6ef38e4c455285c41df2c971ca0b35b311ff97db1241dc291d629c4575344"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b932089b52761664c10e9eb3070ef60dc4ddc71fb491e8a6cef325633c118618"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list — ingress_egress_gw_ar.inside_static_routes.static_route_list / 56c344150ca6 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-83daff971568976e1322aa005d7edac6fc071a1d6ad075cc47c167833c9f00a2)
- ingress_egress_gw_ar.inside_static_routes.static_route_list

<a id="canonical-e80e5ec2b1ff586511b4defeb0d6694415fb36c34f2afb5651b1282e4ed18e06"></a>

Type: `"object"`. list nested block, Optional.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_static_route",
    "simple_static_route")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
static_route_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-af8e6bc7e40d7d5b83b104c863fd1ca1f859492e7cba3ebc02047074fb101f99"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list / 56c344150ca6 / 3

- [custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-aa069d564ca1214a646e9caeb52bc3759274defc0ccb1df199f3669a9f6edbca): complete subsection reference.

<a id="canonical-a13ad8a34a9836c397b4717881d1eca11327c30e1cae79ebef940ac2ae1f7ffc"></a>

<a id="canonical-37d75149bb5c60ce131418c0075263ee01bb53793be8673b9e12973b3f09a176"></a>

## simple_static_route property — ingress_egress_gw_ar.inside_static_routes.static_route_list / 56c344150ca6 / 4

Type: `"string"`. Optional.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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

<a id="canonical-28035a9563e78dcb30ee772c6edd4d6deb57dd4fe41a41ab3fd86f4118c78986"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list / 56c344150ca6 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-aa069d564ca1214a646e9caeb52bc3759274defc0ccb1df199f3669a9f6edbca)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-83daff971568976e1322aa005d7edac6fc071a1d6ad075cc47c167833c9f00a2)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-aa069d564ca1214a646e9caeb52bc3759274defc0ccb1df199f3669a9f6edbca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77a73e1946d5538b830fb4fcafdbe9b481d9fcc65839bda46dac0f235761a33f"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route / 98d5327d286f / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-83daff971568976e1322aa005d7edac6fc071a1d6ad075cc47c167833c9f00a2)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-bbb6ef38e4c455285c41df2c971ca0b35b311ff97db1241dc291d629c4575344)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route

<a id="canonical-4a42bc32ee66a15a4ac3938c673ba58b338d3b99f414d1dcb9e93916647646c5"></a>

Type: `"object"`. single nested block, Optional.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnets")}
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
custom_static_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-7fca98a8cf8504d98fb4adb1529c009535137535c146e4e373ec0363deffb951"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route / 98d5327d286f / 3

<a id="canonical-00958e9a24df8addfba29037fcd503abab4f0c31e1dae31484ab4871ef572eee"></a>

<a id="canonical-b2aa1c4e06b06bb19fac43aafcd5cc92cf39a79444a984651fd0dee24c8f23e8"></a>

## attrs property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route / 98d5327d286f / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](resources--azure_vnet_site--reference--group-006.md#canonical-2e34d1413da767972317b6f73bb4e3a66fc6fbaa161f2ebb9308d2ca2a8fa940): complete subsection reference.

- [nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-326775f6fbdcb204cec51a62e462b14eb5292ccafaf0b26347aa9c9e940f8647): complete subsection reference.

- [subnets](resources--azure_vnet_site--reference--group-006.md#canonical-7ddfbedfae52a377f36b2608995c0fea63c209d586d2a6160920fece2b734328): complete subsection reference.

<a id="canonical-8d673e82680952006be20e04a88d8de763e93a85c0c8a7c2779ba08cfa7b2d24"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route / 98d5327d286f / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.labels](resources--azure_vnet_site--reference--group-006.md#canonical-2e34d1413da767972317b6f73bb4e3a66fc6fbaa161f2ebb9308d2ca2a8fa940)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-326775f6fbdcb204cec51a62e462b14eb5292ccafaf0b26347aa9c9e940f8647)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-006.md#canonical-7ddfbedfae52a377f36b2608995c0fea63c209d586d2a6160920fece2b734328)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-bbb6ef38e4c455285c41df2c971ca0b35b311ff97db1241dc291d629c4575344)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-2e34d1413da767972317b6f73bb4e3a66fc6fbaa161f2ebb9308d2ca2a8fa940"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4cd60d5567754dabb2944d0a3592572f58d0a393434384384161c8a1e03f736e"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.labels — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / fee59092ebb0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-83daff971568976e1322aa005d7edac6fc071a1d6ad075cc47c167833c9f00a2)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-bbb6ef38e4c455285c41df2c971ca0b35b311ff97db1241dc291d629c4575344)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-aa069d564ca1214a646e9caeb52bc3759274defc0ccb1df199f3669a9f6edbca)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-38bc519d90287699eee81621366610c9071b0bdaa140d124488d15a6d85157ac"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this Static Route, these labels can be used in network policy.

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
labels {}
```

<a id="canonical-286de25f9aa12622601ac2ff62e59bd91e3e3280f4e34ba77fe646edff5409ad"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / fee59092ebb0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1e804e69dee97781f44b534b8a8ddb79063de37ebc84a99a849aaca20d8fbd7b"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / fee59092ebb0 / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-aa069d564ca1214a646e9caeb52bc3759274defc0ccb1df199f3669a9f6edbca)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-326775f6fbdcb204cec51a62e462b14eb5292ccafaf0b26347aa9c9e940f8647"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-647f292b9d8a56489f5cf484cd159fa372c95b43594d973c96f12dcada6435cf"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 21d0094ac867 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-83daff971568976e1322aa005d7edac6fc071a1d6ad075cc47c167833c9f00a2)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-bbb6ef38e4c455285c41df2c971ca0b35b311ff97db1241dc291d629c4575344)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-aa069d564ca1214a646e9caeb52bc3759274defc0ccb1df199f3669a9f6edbca)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-0ff30f65c767c21222af054ad434d727636e5af65f57f9660fff1af9d1871f82"></a>

Type: `"object"`. single nested block, Optional.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

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
nexthop {
  # Configure direct properties listed below.
}
```

<a id="canonical-e2b85836082561ca13543ff5823177fdfcd83c2be1f65079c9a117ea22ff8950"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 21d0094ac867 / 3

- [interface](resources--azure_vnet_site--reference--group-006.md#canonical-513c9ada984043238ab174fe3a5312b12ffcdab4d3c84708627f0426d88a86a9): complete subsection reference.

- [nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-23919c29de5c25d44a2689f7f2a6f7eef8e87bc516c2b156bae4f4e2c3775cd6): complete subsection reference.

<a id="canonical-414bf08df17fccb8888d28c102c5a9c7619e277f25fc6eaf4caa0e4c6a463673"></a>

<a id="canonical-d804ae36b897f6016bfd61753f09dbca1e169ac7e75a5bb4d8b49d354accfb18"></a>

## type property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 21d0094ac867 / 4

Type: `"string"`. Optional.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-058f79f47b317035e3182eebfcb050a1a115f651093c8b3421aa14956c5d3459"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 21d0094ac867 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--azure_vnet_site--reference--group-006.md#canonical-513c9ada984043238ab174fe3a5312b12ffcdab4d3c84708627f0426d88a86a9)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-23919c29de5c25d44a2689f7f2a6f7eef8e87bc516c2b156bae4f4e2c3775cd6)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-aa069d564ca1214a646e9caeb52bc3759274defc0ccb1df199f3669a9f6edbca)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-513c9ada984043238ab174fe3a5312b12ffcdab4d3c84708627f0426d88a86a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f433b88816a778c5759948c300a09a55aef39e5bcb7ef63143970e8f3e356d0e"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / b24747430723 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-83daff971568976e1322aa005d7edac6fc071a1d6ad075cc47c167833c9f00a2)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-bbb6ef38e4c455285c41df2c971ca0b35b311ff97db1241dc291d629c4575344)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-aa069d564ca1214a646e9caeb52bc3759274defc0ccb1df199f3669a9f6edbca)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-326775f6fbdcb204cec51a62e462b14eb5292ccafaf0b26347aa9c9e940f8647)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-d744f4f7c0bd609f2558ab53bedaa295b3271dbe42bb3a46afba9ae57591a0f9"></a>

Type: `"object"`. list nested block, Optional.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

Nexthop is network interface when type is "Network-Interface"

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-345f19372d43ef30006ae47e1f20efbf040133c41a3f6c10810f63b6711c287a"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / b24747430723 / 3

<a id="canonical-eb9ffead5f3f22012e49f8759a058bbc176b7651f1114cfdd41f3c85069fe3d1"></a>

<a id="canonical-94e00470f764d8fb4a9475b57921b40eeb33e275bc3f35d205b6d3a6ba3aec8f"></a>

## kind property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / b24747430723 / 4

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

<a id="canonical-71fc3eb46d02bcf63174dd249c66ea83df7605540958eb9f29a54cf4966fb32f"></a>

<a id="canonical-bbfac24fada15a4b58d07b697bab3586b7e3c9ac267825cfc90e34a87e1e4ff1"></a>

## name property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / b24747430723 / 5

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

<a id="canonical-d68a5c55f20efb4967cae95ad336a344b6e8983730fe36b9c31e5baeb58d5fd9"></a>

<a id="canonical-1247163ec784cf3ded58b7a639541678ceb8e586f9c108ae097ea7edef89c617"></a>

## namespace property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / b24747430723 / 6

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

<a id="canonical-c3313d18635c5d31783525ed4703ab92a0ebe070fddeebfa9bba4a3283ad90c9"></a>

<a id="canonical-d4efa4e8f779bef18f4e29b43aafa76a076a45c4c93ff8b1185d375897f281dd"></a>

## tenant property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / b24747430723 / 7

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

<a id="canonical-89e517aaf86d791b67af138d9bdd057096e1689dfff64f7041a1203f4bb4396e"></a>

<a id="canonical-6212e343a2aa150912e21e0025dcecda8abeeab07d5e1cf886b70eac86b9239a"></a>

## uid property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / b24747430723 / 8

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

<a id="canonical-c902c05d674d8a55fa4ce7e30a9d0a4442146968b5f324904d8e028b02d4852e"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / b24747430723 / 9

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-326775f6fbdcb204cec51a62e462b14eb5292ccafaf0b26347aa9c9e940f8647)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-23919c29de5c25d44a2689f7f2a6f7eef8e87bc516c2b156bae4f4e2c3775cd6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e50ef00e49510674318c463888faf3b86d0b28ef396f07a499533fbc2ed9a3b3"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 74476f013fd1 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-83daff971568976e1322aa005d7edac6fc071a1d6ad075cc47c167833c9f00a2)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-bbb6ef38e4c455285c41df2c971ca0b35b311ff97db1241dc291d629c4575344)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-aa069d564ca1214a646e9caeb52bc3759274defc0ccb1df199f3669a9f6edbca)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-326775f6fbdcb204cec51a62e462b14eb5292ccafaf0b26347aa9c9e940f8647)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-8bcc0916151761b70c8a3614b55e0c3051cdd0e0814c963586fbe248045bb18e"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
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
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

Terraform syntax:

```terraform
nexthop_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-5978edefe8f169f0ba57e0868b1010dde883188744f9307d97cd3859ebb29ea4"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 74476f013fd1 / 3

- [dual_stack](resources--azure_vnet_site--reference--group-006.md#canonical-ec3be13a315bc80417469087bac0a173b5347d2cf9f7c1db35fc647a807f89d2): complete subsection reference.

- [ipv4](resources--azure_vnet_site--reference--group-006.md#canonical-9265f1b6bb39111747e5ae9a3f66ab5cef586e0faaeadeaed7ca48e32237f317): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-006.md#canonical-85805219a3c724008ee77ce6f71bd897c7fbb73317c62d4d910a19ceda9f4619): complete subsection reference.

<a id="canonical-b7705d4f84026d93915f20cf4ae4d41687a78e7a21f715a89b1e390cc887fad0"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 74476f013fd1 / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-006.md#canonical-ec3be13a315bc80417469087bac0a173b5347d2cf9f7c1db35fc647a807f89d2)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--azure_vnet_site--reference--group-006.md#canonical-9265f1b6bb39111747e5ae9a3f66ab5cef586e0faaeadeaed7ca48e32237f317)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--azure_vnet_site--reference--group-006.md#canonical-85805219a3c724008ee77ce6f71bd897c7fbb73317c62d4d910a19ceda9f4619)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-326775f6fbdcb204cec51a62e462b14eb5292ccafaf0b26347aa9c9e940f8647)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-ec3be13a315bc80417469087bac0a173b5347d2cf9f7c1db35fc647a807f89d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4eed445a3a6627cbe85072c3d238eefaa64f9aac84faa79f981cdb144b825440"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / f76c4a4d80ed / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-83daff971568976e1322aa005d7edac6fc071a1d6ad075cc47c167833c9f00a2)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-bbb6ef38e4c455285c41df2c971ca0b35b311ff97db1241dc291d629c4575344)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-aa069d564ca1214a646e9caeb52bc3759274defc0ccb1df199f3669a9f6edbca)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-326775f6fbdcb204cec51a62e462b14eb5292ccafaf0b26347aa9c9e940f8647)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-23919c29de5c25d44a2689f7f2a6f7eef8e87bc516c2b156bae4f4e2c3775cd6)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-573ae38a864c5409c4c41b9671e52b2cec1f26f31f899c808fc7895e75c68dd8"></a>

Type: `"object"`. single nested block, Optional.

DualStackAddressType represents both IPv4 and IPv6 together.

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
dual_stack {
  # Configure direct properties listed below.
}
```

<a id="canonical-5534d1eca47bb0bb04b77efd110c14e7e55011498aa5bd3245c8a880d9b8d9d4"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / f76c4a4d80ed / 3

- [ipv4](resources--azure_vnet_site--reference--group-006.md#canonical-eaca118ae13dd05ac23dcaf311f24c6c7a2d5c22f3ed89241265f8c4a4c7ca35): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-006.md#canonical-ce88093308bf8cde7b4f51266a196ae85d5f3b1ec7d79af071a0498f23091e25): complete subsection reference.

<a id="canonical-e7a4a48c6a3fdd34119473467e2750cc1a4b1f073b01059e6c033580dfffbbc0"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / f76c4a4d80ed / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--azure_vnet_site--reference--group-006.md#canonical-eaca118ae13dd05ac23dcaf311f24c6c7a2d5c22f3ed89241265f8c4a4c7ca35)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--azure_vnet_site--reference--group-006.md#canonical-ce88093308bf8cde7b4f51266a196ae85d5f3b1ec7d79af071a0498f23091e25)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-23919c29de5c25d44a2689f7f2a6f7eef8e87bc516c2b156bae4f4e2c3775cd6)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-eaca118ae13dd05ac23dcaf311f24c6c7a2d5c22f3ed89241265f8c4a4c7ca35"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ad87557fa6166743f6b7c1ca5b9a8ee7bd9060dd93aab004e48c9b70789ff0a"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / db3e725f7912 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-83daff971568976e1322aa005d7edac6fc071a1d6ad075cc47c167833c9f00a2)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-bbb6ef38e4c455285c41df2c971ca0b35b311ff97db1241dc291d629c4575344)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-aa069d564ca1214a646e9caeb52bc3759274defc0ccb1df199f3669a9f6edbca)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-326775f6fbdcb204cec51a62e462b14eb5292ccafaf0b26347aa9c9e940f8647)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-23919c29de5c25d44a2689f7f2a6f7eef8e87bc516c2b156bae4f4e2c3775cd6)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-006.md#canonical-ec3be13a315bc80417469087bac0a173b5347d2cf9f7c1db35fc647a807f89d2)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-23b0b3949a7b0dadea829665acc37a080f2d966b8b5bb1835d37e5b4e4a3fd0c"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-f9133aaf75ed63c22adfa86b452a345008b10c656389606a5acab95030130910"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / db3e725f7912 / 3

<a id="canonical-6592545a0619fce03ee99c68ef2b05cbb861fbed77298766f8ec369e196858b9"></a>

<a id="canonical-1d8a116fcd1c8280c2f33f4d9c8d330f20f61339e2d14f6f4357a1f7eac91c07"></a>

## addr property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / db3e725f7912 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-d6053308ff8c828fe1842e904906eb5cb46ff13dfcc360ea9d3cc69de6e38578"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / db3e725f7912 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-006.md#canonical-ec3be13a315bc80417469087bac0a173b5347d2cf9f7c1db35fc647a807f89d2)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-ce88093308bf8cde7b4f51266a196ae85d5f3b1ec7d79af071a0498f23091e25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8d140bf0a42f1503fa7c63091af3be514584ef6d73ec68c9e8236e3f2fdd18f"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 1546835f29f6 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-83daff971568976e1322aa005d7edac6fc071a1d6ad075cc47c167833c9f00a2)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-bbb6ef38e4c455285c41df2c971ca0b35b311ff97db1241dc291d629c4575344)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-aa069d564ca1214a646e9caeb52bc3759274defc0ccb1df199f3669a9f6edbca)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-326775f6fbdcb204cec51a62e462b14eb5292ccafaf0b26347aa9c9e940f8647)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-23919c29de5c25d44a2689f7f2a6f7eef8e87bc516c2b156bae4f4e2c3775cd6)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-006.md#canonical-ec3be13a315bc80417469087bac0a173b5347d2cf9f7c1db35fc647a807f89d2)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-485a145529574b1299dbd08c7e1ef309963da60c5cf8d6d73d019ffbeef42945"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-dd1c3c1e0e14fd27e3de48f4338aa9d53e103926f137fb3507b7615bd319bd0f"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 1546835f29f6 / 3

<a id="canonical-aae48057b0a2464bb6f36ab738daf4e5cfd15225739e41e6c89c9b37da222987"></a>

<a id="canonical-b6e7a3215d6932d39efbcb802f02bd6749ee803924c595311e87674148b2cbed"></a>

## addr property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 1546835f29f6 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-bdef39cff6b525b5f47c8b4c557aa8e64c42e14c31ae9e97b77c6897536a23a5"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 1546835f29f6 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-006.md#canonical-ec3be13a315bc80417469087bac0a173b5347d2cf9f7c1db35fc647a807f89d2)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-9265f1b6bb39111747e5ae9a3f66ab5cef586e0faaeadeaed7ca48e32237f317"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33354adcb43cf93061422698381d1483bd75b10c5d0dfd36bac7bdbd24023579"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 61e53998e073 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-83daff971568976e1322aa005d7edac6fc071a1d6ad075cc47c167833c9f00a2)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-bbb6ef38e4c455285c41df2c971ca0b35b311ff97db1241dc291d629c4575344)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-aa069d564ca1214a646e9caeb52bc3759274defc0ccb1df199f3669a9f6edbca)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-326775f6fbdcb204cec51a62e462b14eb5292ccafaf0b26347aa9c9e940f8647)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-23919c29de5c25d44a2689f7f2a6f7eef8e87bc516c2b156bae4f4e2c3775cd6)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-2e5826c5be5f3997e0c8f1bb06d40956bb5ac67b1a1aa4dba04324d8cf7bb1fb"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-9d3c6845a40569f6235bd06e3f54f5533c5f278e4e4774646ea74c06e7b835a7"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 61e53998e073 / 3

<a id="canonical-270f2ff207bf2a3e3ce5e5e5a677d8c01f9e50e2406b95c39f9029f2859433fd"></a>

<a id="canonical-5ce06bf05faebc8f2340d3f279be66a8ded5c84f2503f7a57e67ebec00c80142"></a>

## addr property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 61e53998e073 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-5e9070e92e737b090fbb3b3714c11436b624f67f05a839e98e984071b955e855"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 61e53998e073 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-23919c29de5c25d44a2689f7f2a6f7eef8e87bc516c2b156bae4f4e2c3775cd6)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-85805219a3c724008ee77ce6f71bd897c7fbb73317c62d4d910a19ceda9f4619"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-914651cfa55d6b6b587376d6889248574baef1a07ed875bd9dd12ad4514d059f"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 816d66074f53 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-83daff971568976e1322aa005d7edac6fc071a1d6ad075cc47c167833c9f00a2)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-bbb6ef38e4c455285c41df2c971ca0b35b311ff97db1241dc291d629c4575344)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-aa069d564ca1214a646e9caeb52bc3759274defc0ccb1df199f3669a9f6edbca)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-326775f6fbdcb204cec51a62e462b14eb5292ccafaf0b26347aa9c9e940f8647)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-23919c29de5c25d44a2689f7f2a6f7eef8e87bc516c2b156bae4f4e2c3775cd6)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-02e7a4100c0215ea8c0dadf9389af5983249978bd274f1c465b52b15ba22f893"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-7b363e007ffb5dacbdc2c841b770f2436944671d81f739cb77b29ede3906eae2"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 816d66074f53 / 3

<a id="canonical-f39229f16636f1a43c6ac5143e173b0d24b32c5a19989f91675fb093b5a6d608"></a>

<a id="canonical-c8a9e661bcd5cc3e70f94c24ebdff949f5440caaf5d7a75c5e28b317bd67e228"></a>

## addr property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 816d66074f53 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-aba504933fd70183308067ad49b86b55a5725949baff1a39cc5fdc4a91df4a36"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 816d66074f53 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-23919c29de5c25d44a2689f7f2a6f7eef8e87bc516c2b156bae4f4e2c3775cd6)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-7ddfbedfae52a377f36b2608995c0fea63c209d586d2a6160920fece2b734328"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0a468f217830920a302f8e07bebe1fd94b913675fc3e605dab5a0e0a84f3dfa"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / e3286bbe0904 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-83daff971568976e1322aa005d7edac6fc071a1d6ad075cc47c167833c9f00a2)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-bbb6ef38e4c455285c41df2c971ca0b35b311ff97db1241dc291d629c4575344)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-aa069d564ca1214a646e9caeb52bc3759274defc0ccb1df199f3669a9f6edbca)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-d10dcd9d4871f3ef70f66023088ed68c73074ed1d8132da317b8a92ed477af19"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ipv4",
    "ipv6")}
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

Terraform syntax:

```terraform
subnets {
  # Configure direct properties listed below.
}
```

<a id="canonical-153d7e41071090f8cc88cbd24e2b626857a5cfd80cc64fefaff6438f4c9252ee"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / e3286bbe0904 / 3

- [ipv4](resources--azure_vnet_site--reference--group-006.md#canonical-a713763901a26328485d63d4630f513aff53f79ad812359a1e21cb70cd9d8572): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-006.md#canonical-8974ad675b4547805a9111e977ff5ee041e833408003f31d8be435d668325904): complete subsection reference.

<a id="canonical-735e7a81e9e48494ff66a41c7e51fcfc79367f9b432e120ce67c3b43bc29a192"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / e3286bbe0904 / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--azure_vnet_site--reference--group-006.md#canonical-a713763901a26328485d63d4630f513aff53f79ad812359a1e21cb70cd9d8572)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--azure_vnet_site--reference--group-006.md#canonical-8974ad675b4547805a9111e977ff5ee041e833408003f31d8be435d668325904)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-aa069d564ca1214a646e9caeb52bc3759274defc0ccb1df199f3669a9f6edbca)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-a713763901a26328485d63d4630f513aff53f79ad812359a1e21cb70cd9d8572"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27b54101bd453b925528e8593532ac62ee93e918391aa3923dd54c85b102c4fd"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / a33afb55d082 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-83daff971568976e1322aa005d7edac6fc071a1d6ad075cc47c167833c9f00a2)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-bbb6ef38e4c455285c41df2c971ca0b35b311ff97db1241dc291d629c4575344)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-aa069d564ca1214a646e9caeb52bc3759274defc0ccb1df199f3669a9f6edbca)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-006.md#canonical-7ddfbedfae52a377f36b2608995c0fea63c209d586d2a6160920fece2b734328)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-5df7840c6b953216baca61cd206b22d67c4e89ec6238d7e861ffb440c9722211"></a>

Type: `"object"`. single nested block, Optional.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-dafcfb4fd657c83006b3cf0af43e05e2b0a81c22ad5fd796c7e25c516fc489a2"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / a33afb55d082 / 3

<a id="canonical-243f44043332a147490704fac676a5fd639ab48e784ef47b73fb6cd1a4ecc81c"></a>

<a id="canonical-72d06c400cfa323330873b3733427b3118046f9621b5d6cd558184a7fa180030"></a>

## plen property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / a33afb55d082 / 4

Type: `"number"`. Optional.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-5a57cd6070c4f0d3d06597fbba864081ad4beb0107ebdfa1bb751eb5fe9f6aab"></a>

<a id="canonical-47afb16b0f7562b8fb75ecb9ef13ea2e76ae6516afb05eb6e3ad7e5920886814"></a>

## prefix property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / a33afb55d082 / 5

Type: `"string"`. Optional.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

<a id="canonical-7448daecc64709404cf68ab2a1571041ecbac1341e030aa646272a18eee092be"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / a33afb55d082 / 6

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-006.md#canonical-7ddfbedfae52a377f36b2608995c0fea63c209d586d2a6160920fece2b734328)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-8974ad675b4547805a9111e977ff5ee041e833408003f31d8be435d668325904"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f90e5af7d7631d22cb2e7d7d6cbc51395d08649520ecc83f399b6deefa9d365"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 9bb45fa25a2c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-83daff971568976e1322aa005d7edac6fc071a1d6ad075cc47c167833c9f00a2)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-bbb6ef38e4c455285c41df2c971ca0b35b311ff97db1241dc291d629c4575344)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-aa069d564ca1214a646e9caeb52bc3759274defc0ccb1df199f3669a9f6edbca)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-006.md#canonical-7ddfbedfae52a377f36b2608995c0fea63c209d586d2a6160920fece2b734328)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-b1a0c52479581c383f834f847199ee8fe8f25ef24889ab4f1c2eb6453d5b8147"></a>

Type: `"object"`. single nested block, Optional.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-1bff9aca87e46f1a0cd1d5407eca74746b8441f428feb6d4b261c317fb2dc80f"></a>

## Direct properties — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 9bb45fa25a2c / 3

<a id="canonical-9c6f288504d41942204038ce4f35f632296a9d1bb8db6c73437b392535742ab2"></a>

<a id="canonical-621e0725d9738d6872814c11fce2d60dc446b676af7d3d1e6c7837f8b02f5f73"></a>

## plen property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 9bb45fa25a2c / 4

Type: `"number"`. Optional.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-a67d511f37cff01b6dbb42743e48f09f4024a0c763b12e0b5c564e01555fe3ad"></a>

<a id="canonical-efeb0624b0b39bcbf4ab6bb5d98db31b21be472526630e7ec4fb95ff27b80afe"></a>

## prefix property — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 9bb45fa25a2c / 5

Type: `"string"`. Optional.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-71a75cce82773358e5a1ae7db8c08ce2dd64e166470691adc4a29ce25db919a8"></a>

## Next pages — ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route. / 9bb45fa25a2c / 6

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-006.md#canonical-7ddfbedfae52a377f36b2608995c0fea63c209d586d2a6160920fece2b734328)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-818b2f34104f14d32057cca2e534182d9ac2d33a53df4aa9b717a719ada0eb26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee1b22a53df0ed3d70402fcdaa1ae48304b5c0623eb5aca519dac47320bbc97c"></a>

## ingress_egress_gw_ar.no_dc_cluster_group — ingress_egress_gw_ar.no_dc_cluster_group / bdcd7a5b3bec / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.no_dc_cluster_group

<a id="canonical-38a6a9ef78974f0bc692e37ef0fb5483c9069ad16a548aa364e64a184363f1f7"></a>

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
no_dc_cluster_group = {}
```

<a id="canonical-af2a58ee8dc426704431ab92893f03a3bd81930c59ac40f03deddde9575177b8"></a>

## Direct properties — ingress_egress_gw_ar.no_dc_cluster_group / bdcd7a5b3bec / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6db7f9345ad7adf5c83b4353d3660db46628ab48d704c11f6b646512b6420ed5"></a>

## Next pages — ingress_egress_gw_ar.no_dc_cluster_group / bdcd7a5b3bec / 4

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-9bb715e41a041df39664d0319b1bee1c7e20c1a98800d02496ff67fc172ca6e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0dbe6a53dbf90d11e1e09fb44de5ec2069f6a6baed2893266b9ef9ea20dd7684"></a>

## ingress_egress_gw_ar.no_forward_proxy — ingress_egress_gw_ar.no_forward_proxy / 5e08618f439b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.no_forward_proxy

<a id="canonical-b11ff4e72fc686f869fec3242a1335d7adf23ace51e37a92bb4d96e041d1892f"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no forward proxy.

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
no_forward_proxy = {}
```

<a id="canonical-372532402d06356b5d7d78f66e7606e60403d8bc213a26fff3bb03a5b351ba4e"></a>

## Direct properties — ingress_egress_gw_ar.no_forward_proxy / 5e08618f439b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cb4020bce536dba1933f29c1519ab60b979e5f4e08f9cd07426ebe3bef49edbd"></a>

## Next pages — ingress_egress_gw_ar.no_forward_proxy / 5e08618f439b / 4

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-f51b02e48409cf9538906b61dbe13c3a9350e7fd69c4b25e12302657fab47215"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad645dde2b68c1a5fb6e86d71f2007e5f16fe591da1d756ee9907258c334f11f"></a>

## ingress_egress_gw_ar.no_global_network — ingress_egress_gw_ar.no_global_network / 8b44ffdef4bc / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.no_global_network

<a id="canonical-0fbd3f123f5126ae42a694992130a4f523c036307c344ee0fac692ff0698d4d4"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no global network.

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
no_global_network = {}
```

<a id="canonical-a7f3311de0eb16c8179af0e0a3c6ea8905a010080fed3735f0554df80efaa151"></a>

## Direct properties — ingress_egress_gw_ar.no_global_network / 8b44ffdef4bc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8be3e0600a36d20cdb0723cf45564744d98a913af5a508098f5fac53a80f7929"></a>

## Next pages — ingress_egress_gw_ar.no_global_network / 8b44ffdef4bc / 4

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-02d31b1e638725cd89df46178569cc3cb222f88646f90faf10df5185e21ecd9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad61c2d9f6b563e01670186f81b1960ff9e4695918e218a4e3661cca34a6bb8c"></a>

## ingress_egress_gw_ar.no_inside_static_routes — ingress_egress_gw_ar.no_inside_static_routes / 7a65a6d4d507 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.no_inside_static_routes

<a id="canonical-4934e698b4f145c220f10e58fa7f4cae119e7021d4f4ab5c49c462450cf4484c"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no inside static routes.

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
no_inside_static_routes = {}
```

<a id="canonical-38412127930a88838182cf65865d6aaaf9f44d3201dc8460a9efed637846484d"></a>

## Direct properties — ingress_egress_gw_ar.no_inside_static_routes / 7a65a6d4d507 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a527eee5ddf7ae863c7b2e9fe2a6c93fbe7c615a0151f5e5e20fb35aae2eb34c"></a>

## Next pages — ingress_egress_gw_ar.no_inside_static_routes / 7a65a6d4d507 / 4

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-86a905690a308b5667688711ac6d7232a89f386b2cb89604be9a4678fdd5b4d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0dd218318f0880d71fa85f30ed43ac65a03cdbba6504d4f725c8ee6b5d10cdc1"></a>

## ingress_egress_gw_ar.no_network_policy — ingress_egress_gw_ar.no_network_policy / 8d04417e06ae / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.no_network_policy

<a id="canonical-bd4ff337d7c41bdf3304de75822ede40655f5dc689342d918397c3ad7587584e"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
no_network_policy = {}
```

<a id="canonical-10a43aa66a2f984c2192749ca26ad26fe846529a41d26277aeabd0fdef9743eb"></a>

## Direct properties — ingress_egress_gw_ar.no_network_policy / 8d04417e06ae / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7fc248adff3e14772bee88e5018c23872dcc80fb50030c4aebe6831e85dd189b"></a>

## Next pages — ingress_egress_gw_ar.no_network_policy / 8d04417e06ae / 4

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-edbfe8c94180d225a7f8eb569f8efc122ae23ded1317587e2dde7c7063d4c534"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba37ebf6d498b5873c01de74da4b5a2b6a6565bfda4906a9828747d4f052b393"></a>

## ingress_egress_gw_ar.no_outside_static_routes — ingress_egress_gw_ar.no_outside_static_routes / af4f2d1012b0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.no_outside_static_routes

<a id="canonical-cfa97ff7af1f7dab249bde7498f93bd923198a632e3343dfe601061170a8b30f"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no outside static routes.

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
no_outside_static_routes = {}
```

<a id="canonical-63896aba2e493db86df980b06ea7005d87b8c6b7fc285955b46dd91dbe6adbf5"></a>

## Direct properties — ingress_egress_gw_ar.no_outside_static_routes / af4f2d1012b0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bb2d2ed1b4d1b8b7488fcc5cdd47c83efdb150ac9b37d249f65908a057bf2953"></a>

## Next pages — ingress_egress_gw_ar.no_outside_static_routes / af4f2d1012b0 / 4

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-3477bd025743eda61d04d12322ef9c35e50e230c97b2d323226575b1598833fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7dce432c02910b933ea0ac5dff3256c3fc06f458c125481476305b06f9b7cb80"></a>

## ingress_egress_gw_ar.node — ingress_egress_gw_ar.node / b2bff600988a / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.node

<a id="canonical-2ed4436c6031f452cc4244efbd86bcfcec39e3d92819fef79e71b7edaaf9ffd5"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating two interface Node in one AZ.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("fault_domain",
    "node_number",
    "update_domain")}
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
node {
  # Configure direct properties listed below.
}
```

<a id="canonical-9c56d0bb433992e56472c7ed7fdaf870921ba0529a5529f6412ef0ba7a76600a"></a>

## Direct properties — ingress_egress_gw_ar.node / b2bff600988a / 3

<a id="canonical-1b27c74ccf22f9b8d47e5e554773a4fa2eb2d7931fcac675d3fed7c06dea000a"></a>

<a id="canonical-69f7102e2a2139d9942d3faab5a71047dd886977a24f52df5d77a9efef9f79fa"></a>

## fault_domain property — ingress_egress_gw_ar.node / b2bff600988a / 4

Type: `"number"`. Optional.

Namuber of fault domains to be used while creating the availability set.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 3),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3,
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
    "ves.io.schema.rules.uint32.lte": "3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "3"
  }
}
```

- [inside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-db3c1f7f142e679d6543aac66b1c1bd3346b98047803d3c64fe13a0c566d9a59): complete subsection reference.

<a id="canonical-c83674e7dfd7859a685eabac720390e32a4cffc9784a6b39e7e25ee117dd0788"></a>

<a id="canonical-040a3352a4a90ac4957c2d9f3bd921af8a2f63d39ae887033969973b97b8671e"></a>

## node_number property — ingress_egress_gw_ar.node / b2bff600988a / 5

Type: `"number"`. Optional.

Number of main nodes to create, either 1 or 3.

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
    "ves.io.schema.rules.uint32.in": "[1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.in": "[1,3]"
  }
}
```

- [outside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-8f1e154b91ff6e6b188def3b996e66ea86a6f4263a2f7eb6ae15101440c8e74e): complete subsection reference.

<a id="canonical-4adcb77d8ab0b4366b5041da299853c79d6131cb066ba9050afa89a418ccc415"></a>

<a id="canonical-8ad9df2a1539a41fd0f21039dcec2af869c2d91dcf95fe0841133b5b8082c144"></a>

## update_domain property — ingress_egress_gw_ar.node / b2bff600988a / 6

Type: `"number"`. Optional.

Namuber of update domains to be used while creating the availability set.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
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
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```
