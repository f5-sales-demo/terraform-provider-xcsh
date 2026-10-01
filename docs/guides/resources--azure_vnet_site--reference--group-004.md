---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-b11dc4b392459d9565a8e12b0150163903d1a790d3307686ce82c2115b2bdd2a"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.connections / ac00a68133a2 / 6

- [ingress_egress_gw.hub.express_route_enabled.connections.metadata](resources--azure_vnet_site--reference--group-004.md#canonical-125d647b90ded50c3640c5d22d3a0958f5206e32e0d02c3c28e2da893c15a051)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription](resources--azure_vnet_site--reference--group-004.md#canonical-b1ea10a801b96c1ced7b48ce5703417915fa3b81171e92a1a1ec69b4f928f2cf)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-125d647b90ded50c3640c5d22d3a0958f5206e32e0d02c3c28e2da893c15a051"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7350f93111c12da4f93e68d2b992bef421a2daa010db28b278a32055b9a51f1"></a>

## ingress_egress_gw.hub.express_route_enabled.connections.metadata — ingress_egress_gw.hub.express_route_enabled.connections.metadata / c2efad32ad8d / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [ingress_egress_gw.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-003.md#canonical-1215fd19c723f8ab3343f9813287edacb86397c626e5a83dbe3bd419902326c5)
- ingress_egress_gw.hub.express_route_enabled.connections.metadata

<a id="canonical-c9171ea1450e33fc1ee43ad8c7dd20f31822866a8ddcd07c39f7efbb0ffe88d9"></a>

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

<a id="canonical-a1c7e60995f59b3d43a17e7eceda78384bc878883a15837d3f00ae2b171c17cd"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.connections.metadata / c2efad32ad8d / 3

<a id="canonical-af0cc1ef1d93a5820488f76fe5e3a3f9fced28b4fee0db7c60281d5f667a6979"></a>

<a id="canonical-bad10ada5960ebfa0367f59f79aeb5247af1881f670f366610e9c3962747d07b"></a>

## description_spec property — ingress_egress_gw.hub.express_route_enabled.connections.metadata / c2efad32ad8d / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-090aa739572c5a3284e613d26efc8213e8654409b1ca5229f47b4fa17a7765da"></a>

<a id="canonical-db0ccfecf68db812da1f7c2c28f7b74001a6d21ee36406ece55b21f12efcf88f"></a>

## name property — ingress_egress_gw.hub.express_route_enabled.connections.metadata / c2efad32ad8d / 5

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

<a id="canonical-85a77936d437a5450ff3a36d7266d11f1d687e08f1a3b9d7174af962b37cb115"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.connections.metadata / c2efad32ad8d / 6

- [ingress_egress_gw.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-003.md#canonical-1215fd19c723f8ab3343f9813287edacb86397c626e5a83dbe3bd419902326c5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-b1ea10a801b96c1ced7b48ce5703417915fa3b81171e92a1a1ec69b4f928f2cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2971493719aac6c9f3524f66aa6eaf9fae6ce50faa59ce39222dc9ef905fb270"></a>

## ingress_egress_gw.hub.express_route_enabled.connections.other_subscription — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription / eac84fda50bf / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [ingress_egress_gw.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-003.md#canonical-1215fd19c723f8ab3343f9813287edacb86397c626e5a83dbe3bd419902326c5)
- ingress_egress_gw.hub.express_route_enabled.connections.other_subscription

<a id="canonical-634f244176c99d957443bccc43ecd1764e6770c7135343124090a6702ee2c204"></a>

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

<a id="canonical-bc2f958a58fff2fd48abc18ada5f0aadeca429ecfa2610297f393f3951e87b4f"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription / eac84fda50bf / 3

- [authorized_key](resources--azure_vnet_site--reference--group-004.md#canonical-8a53064efecc9071401abaf6a5ae24b40019cefbd74de93aa457b14a93f2788b): complete subsection reference.

<a id="canonical-fce733552eb69ee68954ee285edf913c3f18701214108256859f0c9e20fea3e3"></a>

<a id="canonical-a3824454ab3768c9fc687bb57b2307d23a63be4e432e552f1d03d8c62d054ae2"></a>

## circuit_id property — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription / eac84fda50bf / 4

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

<a id="canonical-7b5f2005f11236a04ddbdc47e1d7228b1ebac050b919472866df3dbb976b7375"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription / eac84fda50bf / 5

- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](resources--azure_vnet_site--reference--group-004.md#canonical-8a53064efecc9071401abaf6a5ae24b40019cefbd74de93aa457b14a93f2788b)
- [ingress_egress_gw.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-003.md#canonical-1215fd19c723f8ab3343f9813287edacb86397c626e5a83dbe3bd419902326c5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-8a53064efecc9071401abaf6a5ae24b40019cefbd74de93aa457b14a93f2788b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e5674d8f617fef82e2a83508f9bca5054f831d0661168742f566bb3881a4749"></a>

## ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 2151597f060c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [ingress_egress_gw.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-003.md#canonical-1215fd19c723f8ab3343f9813287edacb86397c626e5a83dbe3bd419902326c5)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription](resources--azure_vnet_site--reference--group-004.md#canonical-b1ea10a801b96c1ced7b48ce5703417915fa3b81171e92a1a1ec69b4f928f2cf)
- ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key

<a id="canonical-d9cb3ff370bf754a01f270f16d5669444e78671100d05392a397dc901e32651d"></a>

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

<a id="canonical-c8d7a6ec69ce357954ff9c18e78945717b5f807687dfe5684cb5d796b7b45d62"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 2151597f060c / 3

- [blindfold_secret_info](resources--azure_vnet_site--reference--group-004.md#canonical-fa67f280f30b5251a8f056d0d9fe36c0595a6f8353ba315b26eebb67db120223): complete subsection reference.

- [clear_secret_info](resources--azure_vnet_site--reference--group-004.md#canonical-3365b4c578954205d7326b13612251f9a4a719ba65b3778c2867ac9844278b1d): complete subsection reference.

<a id="canonical-15f0eee0c87ea53a4822a765a263c379b1a9db870b80754b92307281820b27d5"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 2151597f060c / 4

- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info](resources--azure_vnet_site--reference--group-004.md#canonical-fa67f280f30b5251a8f056d0d9fe36c0595a6f8353ba315b26eebb67db120223)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info](resources--azure_vnet_site--reference--group-004.md#canonical-3365b4c578954205d7326b13612251f9a4a719ba65b3778c2867ac9844278b1d)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription](resources--azure_vnet_site--reference--group-004.md#canonical-b1ea10a801b96c1ced7b48ce5703417915fa3b81171e92a1a1ec69b4f928f2cf)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-fa67f280f30b5251a8f056d0d9fe36c0595a6f8353ba315b26eebb67db120223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b41ea2a08bd4ee3a0009d8067b2202a39f26fbeb67c6fcd83f54bc27f53dc51"></a>

## ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 2c89b3359fb3 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [ingress_egress_gw.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-003.md#canonical-1215fd19c723f8ab3343f9813287edacb86397c626e5a83dbe3bd419902326c5)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription](resources--azure_vnet_site--reference--group-004.md#canonical-b1ea10a801b96c1ced7b48ce5703417915fa3b81171e92a1a1ec69b4f928f2cf)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](resources--azure_vnet_site--reference--group-004.md#canonical-8a53064efecc9071401abaf6a5ae24b40019cefbd74de93aa457b14a93f2788b)
- ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info

<a id="canonical-63b84b99e733407e0ac445811b012934cfa6b28ddfc6dc97fc30c60345311a23"></a>

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

<a id="canonical-e15ee4cd2fcfd2b26c10f526a5fe1831b8d046b175d79338b0f9922a5a9cc850"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 2c89b3359fb3 / 3

<a id="canonical-ee5427a2860cf5e2e868259ecba29838a5286db7793959972143b2561ed88690"></a>

<a id="canonical-4a9b7eae0488fab5385c12b8b757e1434a19942818a9b18cc7d160b7d00be409"></a>

## decryption_provider property — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 2c89b3359fb3 / 4

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

<a id="canonical-8dda8c9bd5344b4a38c49a6ea12b26ef4710edbe5c3b751c7df7453625e3696b"></a>

<a id="canonical-bcbd6863966e3c700dfd4eeaa8bfb9d091b67ba5268808d6dc70540922fede65"></a>

## location property — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 2c89b3359fb3 / 5

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

<a id="canonical-b1cab9cd7d6c2e62ad271e30eb219d1722018e7499455614d3a235e4e70884ff"></a>

<a id="canonical-028a2e09b42b40ae330bde803345cf8397f4d610e710851c6223f30dc336c068"></a>

## store_provider property — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 2c89b3359fb3 / 6

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

<a id="canonical-03fae5d477ecd6ee55f1e2c03adb03fab69cdb0c6813056f5b8b443f0d04a41d"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 2c89b3359fb3 / 7

- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](resources--azure_vnet_site--reference--group-004.md#canonical-8a53064efecc9071401abaf6a5ae24b40019cefbd74de93aa457b14a93f2788b)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-3365b4c578954205d7326b13612251f9a4a719ba65b3778c2867ac9844278b1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dea20b5be7fcd6fe1838804e7c00ae3f553e4e44e39e2863ce6279572e12fdec"></a>

## ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 06085b72f528 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [ingress_egress_gw.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-003.md#canonical-1215fd19c723f8ab3343f9813287edacb86397c626e5a83dbe3bd419902326c5)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription](resources--azure_vnet_site--reference--group-004.md#canonical-b1ea10a801b96c1ced7b48ce5703417915fa3b81171e92a1a1ec69b4f928f2cf)
- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](resources--azure_vnet_site--reference--group-004.md#canonical-8a53064efecc9071401abaf6a5ae24b40019cefbd74de93aa457b14a93f2788b)
- ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info

<a id="canonical-588bcd6e699fae0814ef9dd240e919ab1445fecbe76c34154aa6d97b663c9aad"></a>

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

<a id="canonical-471e16b2fd784831c605d6c3bc20a6ff43c0e57475077ab4079878e322477a3d"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 06085b72f528 / 3

<a id="canonical-4c29decdd595958ca9ba85747e75ff58e973d7dc874f6b3fbbef46b5f322eb35"></a>

<a id="canonical-f4fcbe522ae96db08706e8b7bbf76dc7a5914470586cc4ec2113fdabf227bf7c"></a>

## provider_ref property — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 06085b72f528 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-6c16227be6f9f3a8c5bb53540b95de4da316619d0c9a45d5f79ab7e506d169f6"></a>

<a id="canonical-20bdb7ed987fa74d7202cb138e58bd1666cc7974ea136a25b7a39ffc45ff485d"></a>

## url property — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 06085b72f528 / 5

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

<a id="canonical-bafc67087b514a1a285e6d304fd79b38bf8a447070339ef131911f1a13e88522"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.autho / 06085b72f528 / 6

- [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](resources--azure_vnet_site--reference--group-004.md#canonical-8a53064efecc9071401abaf6a5ae24b40019cefbd74de93aa457b14a93f2788b)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-50ac0535c6086004d634d6025d4a637a3c5b58d76b0caf8617673766d16f2858"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8b50a387b3c9e0751b7f0f9932080263638704c809529126a8b6b1e18d44a72"></a>

## ingress_egress_gw.hub.express_route_enabled.do_not_advertise_to_route_server — ingress_egress_gw.hub.express_route_enabled.do_not_advertise_to_route_server / d251551f6bca / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- ingress_egress_gw.hub.express_route_enabled.do_not_advertise_to_route_server

<a id="canonical-34a3b6c90495637fc03a7290ff1f01cc82442fc621e177bcc196bc306bfeebf2"></a>

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

<a id="canonical-5401872946cbf4bbea467b931a9a54815616a5af9ed9faaaef4dada7d0dbb4c9"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.do_not_advertise_to_route_server / d251551f6bca / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-48875ccd57f32f5e89c85d64f442f19e4baa2caea35bb3fa9104b985a8641db4"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.do_not_advertise_to_route_server / d251551f6bca / 4

- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-59bc55040bfb2d3d1829e8ec5af9d8b18d29dabc8777c4c33f4f2a74c1b44bce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c887ecafd45d7031c4af84ee3d38de1dc9ba8811d6d8703e8378f3901d836cfb"></a>

## ingress_egress_gw.hub.express_route_enabled.gateway_subnet — ingress_egress_gw.hub.express_route_enabled.gateway_subnet / 5672ecf9d042 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- ingress_egress_gw.hub.express_route_enabled.gateway_subnet

<a id="canonical-2ae4cae7ae25ff0207045b25f35130cc2c890deea374b87eab0ce1354306e4a4"></a>

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

<a id="canonical-819f6c7b72f8fdf0bc14e457d339395e4841a563566be149f83190cfa3ea062d"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.gateway_subnet / 5672ecf9d042 / 3

- [auto](resources--azure_vnet_site--reference--group-004.md#canonical-ae8de25b9c6ce83f231df051ba7cd544ed0b8ff9d85eec515a17fdbbce67b0d6): complete subsection reference.

- [subnet](resources--azure_vnet_site--reference--group-004.md#canonical-df227b34931cc8f71e1874dfdd788d42ac5bcab9dae93e68cc028d4eae4eee6b): complete subsection reference.

- [subnet_param](resources--azure_vnet_site--reference--group-004.md#canonical-09a3f49c9797f9937f506bb84530ece059933cf0396f386f780a3dee831f11e0): complete subsection reference.

<a id="canonical-e6353037ebd3524bb558f4b2bb6ebec66481e002b93517100a3f52690fe7fe32"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.gateway_subnet / 5672ecf9d042 / 4

- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.auto](resources--azure_vnet_site--reference--group-004.md#canonical-ae8de25b9c6ce83f231df051ba7cd544ed0b8ff9d85eec515a17fdbbce67b0d6)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet](resources--azure_vnet_site--reference--group-004.md#canonical-df227b34931cc8f71e1874dfdd788d42ac5bcab9dae93e68cc028d4eae4eee6b)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param](resources--azure_vnet_site--reference--group-004.md#canonical-09a3f49c9797f9937f506bb84530ece059933cf0396f386f780a3dee831f11e0)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-ae8de25b9c6ce83f231df051ba7cd544ed0b8ff9d85eec515a17fdbbce67b0d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87483d383f6bdfbbd32d9610b6cef9b1579472b5bc26ae21e75d6e46a432c0d8"></a>

## ingress_egress_gw.hub.express_route_enabled.gateway_subnet.auto — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.auto / 72e53e9b312c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-004.md#canonical-59bc55040bfb2d3d1829e8ec5af9d8b18d29dabc8777c4c33f4f2a74c1b44bce)
- ingress_egress_gw.hub.express_route_enabled.gateway_subnet.auto

<a id="canonical-c5c864c601ceadfec06bbb0ac1a67071f98a7ea2668b92e5425a1a2a2048111c"></a>

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

<a id="canonical-9c12324a002587dba18f4bbdd042edc2bdb6ad57ddd6bec4ca80515694dcd108"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.auto / 72e53e9b312c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-12323097eba105ca5e273a468eb285f037367377527f3abbd775fac675dc2fb1"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.auto / 72e53e9b312c / 4

- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-004.md#canonical-59bc55040bfb2d3d1829e8ec5af9d8b18d29dabc8777c4c33f4f2a74c1b44bce)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-df227b34931cc8f71e1874dfdd788d42ac5bcab9dae93e68cc028d4eae4eee6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f9afc03770ab6a4487e59b190e39ed0e58140b840bf5ceb415a633dbe090887"></a>

## ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet / 4b1b37d9265f / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-004.md#canonical-59bc55040bfb2d3d1829e8ec5af9d8b18d29dabc8777c4c33f4f2a74c1b44bce)
- ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet

<a id="canonical-f4412328f9594c8907042fc0973987d8edd1c2598c50cd8c4a5ad202e35becbd"></a>

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

<a id="canonical-24808021f419e3802023d2697109fdc9d12ee019c22e0120375ecbc289012b3d"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet / 4b1b37d9265f / 3

<a id="canonical-fe96217a94459d3c074e0aa3662f3374897e4b25be5802cf6f2130b21084d0d5"></a>

<a id="canonical-05e48d33e755602adbc54d200ff30cb8fe350f98f338629d388b3fb023e1c867"></a>

## subnet_resource_grp property — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet / 4b1b37d9265f / 4

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

- [vnet_resource_group](resources--azure_vnet_site--reference--group-004.md#canonical-3a2b75f801f0b2a2110695781d4987675f37b72ff29f14c6ea4e1c5609fc0613): complete subsection reference.

<a id="canonical-66bc1a7691a755d8ebbe62473077a6d1e4ea1b1f8c0d64bd857750224f612a61"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet / 4b1b37d9265f / 5

- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group](resources--azure_vnet_site--reference--group-004.md#canonical-3a2b75f801f0b2a2110695781d4987675f37b72ff29f14c6ea4e1c5609fc0613)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-004.md#canonical-59bc55040bfb2d3d1829e8ec5af9d8b18d29dabc8777c4c33f4f2a74c1b44bce)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-3a2b75f801f0b2a2110695781d4987675f37b72ff29f14c6ea4e1c5609fc0613"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-670f8f20a608575cbaa7416fb1492a3da3b162008ef5c635d39d12a76e1b97d5"></a>

## ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_ / c08fd27fc1f3 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-004.md#canonical-59bc55040bfb2d3d1829e8ec5af9d8b18d29dabc8777c4c33f4f2a74c1b44bce)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet](resources--azure_vnet_site--reference--group-004.md#canonical-df227b34931cc8f71e1874dfdd788d42ac5bcab9dae93e68cc028d4eae4eee6b)
- ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group

<a id="canonical-3a0649a4ddaa562600493ce9793134c9467626ac62b7ecd5f773173517546f16"></a>

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

<a id="canonical-8be5ffe32a155e44d6dd0c657d9061b1f95078f9f09bebc6ad36ed50bed4a4ac"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_ / c08fd27fc1f3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cd942716457803a61d875010f48396e85c7bc6294733bd1930d1938169b9f6c7"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_ / c08fd27fc1f3 / 4

- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet](resources--azure_vnet_site--reference--group-004.md#canonical-df227b34931cc8f71e1874dfdd788d42ac5bcab9dae93e68cc028d4eae4eee6b)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-09a3f49c9797f9937f506bb84530ece059933cf0396f386f780a3dee831f11e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bea2c60d81ef7a6c4c8e95a4f928e372421fc80a27a0d666d42a394780785b2f"></a>

## ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param / 14f311d3fe36 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-004.md#canonical-59bc55040bfb2d3d1829e8ec5af9d8b18d29dabc8777c4c33f4f2a74c1b44bce)
- ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param

<a id="canonical-52640d8187c26a21410b187d13b6e26379d9ed4985d41e8c6b13601f94794d66"></a>

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

<a id="canonical-0ae685d2ac0903b2e57b7289b656923f09ce861224ce6b534f0090010244522b"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param / 14f311d3fe36 / 3

<a id="canonical-e8b4631593256b417b84b409c1c9307d45a98393e2878b2cc81fc72984ef3bbe"></a>

<a id="canonical-52e3555bc5a41d87c35a2562912378e4ea47b2dc85a7b68e845434ff5be9a44b"></a>

## ipv4 property — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param / 14f311d3fe36 / 4

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

<a id="canonical-247025d49a6725cd4711763c3c61c4edd742f65668f4b0d4e3c6591e17d9642f"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param / 14f311d3fe36 / 5

- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-004.md#canonical-59bc55040bfb2d3d1829e8ec5af9d8b18d29dabc8777c4c33f4f2a74c1b44bce)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-938e5cca7722126b40865d713ac78f3b0247806313b252296342fe51b97e734a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aece7083d2d7f0990bc98b085056445c618d601dfcdc53447ed050930e1bd3ab"></a>

## ingress_egress_gw.hub.express_route_enabled.route_server_subnet — ingress_egress_gw.hub.express_route_enabled.route_server_subnet / cca75788d529 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- ingress_egress_gw.hub.express_route_enabled.route_server_subnet

<a id="canonical-f253d4dd5bf8abc1e66e407de962c9f03be559fe42f09a416729fe93c54e0cdf"></a>

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

<a id="canonical-24d8e4f4053164acb17a292fb0ded97c5c322791910a60d7620a5215f54c634f"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.route_server_subnet / cca75788d529 / 3

- [auto](resources--azure_vnet_site--reference--group-004.md#canonical-9beda5ccf844fcd08ed1d0d8748d675238fc3998b6ed98c4ef48243d73bdaad7): complete subsection reference.

- [subnet](resources--azure_vnet_site--reference--group-004.md#canonical-7791cc7f2253896dcbfc791ade49082b73f2447cc1f603a89ba964c9f557f02d): complete subsection reference.

- [subnet_param](resources--azure_vnet_site--reference--group-004.md#canonical-4c9fed782ad332e14f0bf2993de752a93960a67f064b3655549683d8b8e4dc6d): complete subsection reference.

<a id="canonical-24daeb2f96d992afa26e328831b5e690407bb88de4668987350d91b2a7a80f9d"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.route_server_subnet / cca75788d529 / 4

- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto](resources--azure_vnet_site--reference--group-004.md#canonical-9beda5ccf844fcd08ed1d0d8748d675238fc3998b6ed98c4ef48243d73bdaad7)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet](resources--azure_vnet_site--reference--group-004.md#canonical-7791cc7f2253896dcbfc791ade49082b73f2447cc1f603a89ba964c9f557f02d)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param](resources--azure_vnet_site--reference--group-004.md#canonical-4c9fed782ad332e14f0bf2993de752a93960a67f064b3655549683d8b8e4dc6d)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-9beda5ccf844fcd08ed1d0d8748d675238fc3998b6ed98c4ef48243d73bdaad7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dee3f5db1ac0b736fb95c523124d71839b170087d04875bee6948c7a6aab032a"></a>

## ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto / ac0011c6d02f / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-004.md#canonical-938e5cca7722126b40865d713ac78f3b0247806313b252296342fe51b97e734a)
- ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto

<a id="canonical-7e409f066f8585549425229b440644c7796c85d6275879caee1b29155a19e788"></a>

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

<a id="canonical-114dc07496f753c015ef2339deeef05fbe5caf043a476d306bcffaef8d3d0a63"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto / ac0011c6d02f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9aba15cb23dbbc458ed57d0d9dd285c00d93bc2c75b5ab9a603eeb0b0837826c"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto / ac0011c6d02f / 4

- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-004.md#canonical-938e5cca7722126b40865d713ac78f3b0247806313b252296342fe51b97e734a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-7791cc7f2253896dcbfc791ade49082b73f2447cc1f603a89ba964c9f557f02d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eca0cc661faeb98dcb7df422038bc78234cd7cd8ef10df9dd9b81e0f3dfa70d3"></a>

## ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet / 0d8c67c776f0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-004.md#canonical-938e5cca7722126b40865d713ac78f3b0247806313b252296342fe51b97e734a)
- ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet

<a id="canonical-e625a1e4ae239a1951527c2f3f59a90f5d0f6c60378d9e5a181636ff78a3d732"></a>

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

<a id="canonical-35004446c2ab6755378444310341cf2f6008543b6d7850878ecaed4d86fd34b1"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet / 0d8c67c776f0 / 3

<a id="canonical-c19762181263f4024528dd6ce4fd88a742e114743a9e21be0da57bbd37cc69a1"></a>

<a id="canonical-a8a8052d692369459abbca921dbcb071ff6e6319e66069ebe18b80bf4d9d4a46"></a>

## subnet_resource_grp property — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet / 0d8c67c776f0 / 4

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

- [vnet_resource_group](resources--azure_vnet_site--reference--group-004.md#canonical-5e234bb352b059c6ad5e5cece2c4589038dfc2cf1a4a3df4ad40aab0061eadca): complete subsection reference.

<a id="canonical-c20fb7400130c4d7211d06cbb80d58b5a9f5ce5af9766ab1eda55bee69128e23"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet / 0d8c67c776f0 / 5

- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group](resources--azure_vnet_site--reference--group-004.md#canonical-5e234bb352b059c6ad5e5cece2c4589038dfc2cf1a4a3df4ad40aab0061eadca)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-004.md#canonical-938e5cca7722126b40865d713ac78f3b0247806313b252296342fe51b97e734a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-5e234bb352b059c6ad5e5cece2c4589038dfc2cf1a4a3df4ad40aab0061eadca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d65120d5438aaf7647105cff98e84ea1a834ac15b255eeeedc77f28dd9c31ef"></a>

## ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.vnet_reso / ba4f2b49daa2 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-004.md#canonical-938e5cca7722126b40865d713ac78f3b0247806313b252296342fe51b97e734a)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet](resources--azure_vnet_site--reference--group-004.md#canonical-7791cc7f2253896dcbfc791ade49082b73f2447cc1f603a89ba964c9f557f02d)
- ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group

<a id="canonical-9b61b7408a1f4a024fb54b2d996df31ef8ba5dc346b351ba1780fcd03d1c184b"></a>

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

<a id="canonical-d5cfd0a62ef7ee6416176ddbe062a249051317da7eddeef4fb7a1c9bf3d2e383"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.vnet_reso / ba4f2b49daa2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-52dd2b89acad0ed3345e30754addf64406268e8780d3636e6d827c9e17880e7f"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.vnet_reso / ba4f2b49daa2 / 4

- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet](resources--azure_vnet_site--reference--group-004.md#canonical-7791cc7f2253896dcbfc791ade49082b73f2447cc1f603a89ba964c9f557f02d)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-4c9fed782ad332e14f0bf2993de752a93960a67f064b3655549683d8b8e4dc6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e0c1257dc3c576c41c4bc0ce35075754787d084c80b0d22b77d527ad3f88863"></a>

## ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param / 216c33346735 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-004.md#canonical-938e5cca7722126b40865d713ac78f3b0247806313b252296342fe51b97e734a)
- ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param

<a id="canonical-ce710566f24266f6bf849444b9a4eaef0f168b3cd94ac25d9d452a2f8a81b162"></a>

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

<a id="canonical-5c7ad5f629555c0907963c6abe11aa94b008dac15f24f2b68fe141ee66cb71ec"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param / 216c33346735 / 3

<a id="canonical-bfdb93cdfdbe82921050d2ccc5a9fd03a0b130182e498581a2b45fe6d9b3ae6a"></a>

<a id="canonical-5793e4100a999c2c0df71d4c5908bf22d56e1542023bfc1610ac1c6b1b97af6c"></a>

## ipv4 property — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param / 216c33346735 / 4

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

<a id="canonical-84c4641281c28e5fa46f485bac90eb6197d19c01bfd47941eac1ae15e4fb08a3"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param / 216c33346735 / 5

- [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-004.md#canonical-938e5cca7722126b40865d713ac78f3b0247806313b252296342fe51b97e734a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-82f6c3f5833bfe2c29c206e718f3e0a4102616d5efdfd44a62bba1d0e004028d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b0e9c39e1ebb5c94b6a6a835b482e12e53daa5eac97590e097fb801bf0008e0"></a>

## ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route — ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route / ef42b520ef24 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route

<a id="canonical-bb3f79d938a2aadeaef29dcfffc9ea8032636033e1e1718a823739b06446e0fd"></a>

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

<a id="canonical-87ab5248f5398d3797e4cab27f06f6564b6af0ab38320543046855f520bb3349"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route / ef42b520ef24 / 3

<a id="canonical-58bcceee371600e3231e378fe3c785fa03d6c84cc03693fcad05bfc2ba54717b"></a>

<a id="canonical-8626badabb16cee0b6775d37342c4196cc9715eec6f6ebab08e30d312ede46a7"></a>

## cloudlink_network_name property — ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route / ef42b520ef24 / 4

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

<a id="canonical-cfc99571a8b8e4768a8e008b4f228a85d7baeba52cd3fbb893a7545a9f4ce3e6"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route / ef42b520ef24 / 5

- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-93906348aa0b9959bea6d19131eeafe5ebb2b8e93ee7b2a1ba4d512796a35ae9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce9e5de9c1eb9ae67c5ed019754f92a67c5cf9b169fc3827df4e590fe1d8d29d"></a>

## ingress_egress_gw.hub.express_route_enabled.site_registration_over_internet — ingress_egress_gw.hub.express_route_enabled.site_registration_over_internet / 150bfa84c01f / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- ingress_egress_gw.hub.express_route_enabled.site_registration_over_internet

<a id="canonical-9cd8121da88a2ed7e79f6be64a792e8b215d80c8e96ba5095da25f55be516bbb"></a>

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

<a id="canonical-4dc43c5881fd11cb877996b408f7e0625d34b70f26b6a6b5ced5489a236ed83c"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.site_registration_over_internet / 150bfa84c01f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bfcf282f9fe652986fd9dc3f5529041ba5a4d06ce2b23693b9ae6b28173ecf94"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.site_registration_over_internet / 150bfa84c01f / 4

- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-7dba6e5e8d841ac5e2904bdc02d0b664a4f31ff05ee52244f69c9f9fa30143c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-568cdf18c42402b21c566728d90b03fb8934b0b0eaed66a04d3a6827ae064375"></a>

## ingress_egress_gw.hub.express_route_enabled.sku_ergw1az — ingress_egress_gw.hub.express_route_enabled.sku_ergw1az / a9f1c894afb1 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- ingress_egress_gw.hub.express_route_enabled.sku_ergw1az

<a id="canonical-5428e5d4940771bc60edc1a08bfdcd701c2635ab9d926c182a484dbd0426ba8e"></a>

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

<a id="canonical-e766b88bc7439252bf4271251c7886ed252a6f0f4a7e597529bf5d66085648f1"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.sku_ergw1az / a9f1c894afb1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-db5f56859692cbb466a51f61b1e3c18331be754edbf5ad72c186ae353b70ea7f"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.sku_ergw1az / a9f1c894afb1 / 4

- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-f8989fe72a73418ae82d56df9e933a674db9904e771bbd2bcbe90b07c5beb641"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2c8e72502f3f48315f91baecf8e330146dfcbd5a5c4e426bbf359ea5cd42ffc"></a>

## ingress_egress_gw.hub.express_route_enabled.sku_ergw2az — ingress_egress_gw.hub.express_route_enabled.sku_ergw2az / e86bb4dd39a4 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- ingress_egress_gw.hub.express_route_enabled.sku_ergw2az

<a id="canonical-3c14eb8648bcfbce9a202c81e56399ad0fd4c7c78f842719a5a6c52643adf476"></a>

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

<a id="canonical-73795bba546516361fc6a1100c152fdcfe95a3ad1b9c961f4a93c406d59d168c"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.sku_ergw2az / e86bb4dd39a4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-03dc454e156eee52f8b96b884e470c887b27a1b8a26318c0ab5539960fc42486"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.sku_ergw2az / e86bb4dd39a4 / 4

- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-24ad12bb3fc6296b708ae203bec8e786a0e511d715b65915a84262fff129f3e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d07222e1d6102f5b072803d4c5428da9561b33ad846c25aa9963c374523fed3"></a>

## ingress_egress_gw.hub.express_route_enabled.sku_high_perf — ingress_egress_gw.hub.express_route_enabled.sku_high_perf / a7ffa010abc2 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- ingress_egress_gw.hub.express_route_enabled.sku_high_perf

<a id="canonical-dbd49232fbf23c44f9159fd49e77d68fcd0361f57aa449b64c153431fd87789d"></a>

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

<a id="canonical-fa424c6725fb196b47b117fdc55d6bf8841b343488a15d8af4c843c2c6630ec3"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.sku_high_perf / a7ffa010abc2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ed22cadbe8933fbda581f7f3b8d7fff817ef1aed8cd1d35003e92b06bccd370c"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.sku_high_perf / a7ffa010abc2 / 4

- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-0848949928c30f9f12b8526503c55d0e48c821d41b914c206d64155c68721076"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a5e16aa9d5eca05e5612e606296c6b23733545a11fb5ec938f6044b7fb8e853"></a>

## ingress_egress_gw.hub.express_route_enabled.sku_standard — ingress_egress_gw.hub.express_route_enabled.sku_standard / 03a3b3a4e152 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- ingress_egress_gw.hub.express_route_enabled.sku_standard

<a id="canonical-28116f7d89a3233f5860f9af0b175077849d0f79b31cf54debbda3946378094a"></a>

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

<a id="canonical-91604fc74bbd61a7c36ce8831c6ed8dca6658fdf4d362e0236bc715776330847"></a>

## Direct properties — ingress_egress_gw.hub.express_route_enabled.sku_standard / 03a3b3a4e152 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-af68d48e4f26488a24b32395142a16b92f3efc6112da7bd1302993538c84ee08"></a>

## Next pages — ingress_egress_gw.hub.express_route_enabled.sku_standard / 03a3b3a4e152 / 4

- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--reference--group-003.md#canonical-5c7a21a851b0ac0b3f473db0b6736f3c8223747b28317a1a79c294434e6ea911)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-f8cced872d5b22379d7c2e9152ef94b4d0746e01dff9b62189c752cd445803b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d7cb5f25eeca98f7b1e01859c70becadcf193df860fde89528aacc97190f924"></a>

## ingress_egress_gw.hub.spoke_vnets — ingress_egress_gw.hub.spoke_vnets / 2512077e99bf / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- ingress_egress_gw.hub.spoke_vnets

<a id="canonical-7b9199dba00663450e4d20a2b6a5fba72394ce3ce60b2ad5824b15ecdcdf56fc"></a>

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

<a id="canonical-484123d075e2b1b75f9ef845ff2c35370088dc86dc796abe8b03f5a73cf692a6"></a>

## Direct properties — ingress_egress_gw.hub.spoke_vnets / 2512077e99bf / 3

- [auto](resources--azure_vnet_site--reference--group-004.md#canonical-2f3fc7fbce39e8af9e0fab1e83c094466974316a14d9ee9b75a63bf0923fe682): complete subsection reference.

- [labels](resources--azure_vnet_site--reference--group-004.md#canonical-c2b61852e1fee481217ddbd475a4c80044916427a7b98591429c94830a3a16a7): complete subsection reference.

- [manual](resources--azure_vnet_site--reference--group-004.md#canonical-16acfff6f6fd3f683aca3de881f5e785c3339380317e6444089d27e747040fdc): complete subsection reference.

- [vnet](resources--azure_vnet_site--reference--group-004.md#canonical-b23785f031ff7eec3f52ede411ba9638edb96e868e038f64d9efdf3915a91bc8): complete subsection reference.

<a id="canonical-9f7a3efc3d088f86f398ce7148175e0bdb6a89d6e4ac770c22e8ff6f6be84cb9"></a>

## Next pages — ingress_egress_gw.hub.spoke_vnets / 2512077e99bf / 4

- [ingress_egress_gw.hub.spoke_vnets.auto](resources--azure_vnet_site--reference--group-004.md#canonical-2f3fc7fbce39e8af9e0fab1e83c094466974316a14d9ee9b75a63bf0923fe682)
- [ingress_egress_gw.hub.spoke_vnets.labels](resources--azure_vnet_site--reference--group-004.md#canonical-c2b61852e1fee481217ddbd475a4c80044916427a7b98591429c94830a3a16a7)
- [ingress_egress_gw.hub.spoke_vnets.manual](resources--azure_vnet_site--reference--group-004.md#canonical-16acfff6f6fd3f683aca3de881f5e785c3339380317e6444089d27e747040fdc)
- [ingress_egress_gw.hub.spoke_vnets.vnet](resources--azure_vnet_site--reference--group-004.md#canonical-b23785f031ff7eec3f52ede411ba9638edb96e868e038f64d9efdf3915a91bc8)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-2f3fc7fbce39e8af9e0fab1e83c094466974316a14d9ee9b75a63bf0923fe682"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3efb9a645a6470c796eb23c490581fd4dba015eea1c74754548bbb9e27caa15e"></a>

## ingress_egress_gw.hub.spoke_vnets.auto — ingress_egress_gw.hub.spoke_vnets.auto / 785ff5ed4055 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.spoke_vnets](resources--azure_vnet_site--reference--group-004.md#canonical-f8cced872d5b22379d7c2e9152ef94b4d0746e01dff9b62189c752cd445803b3)
- ingress_egress_gw.hub.spoke_vnets.auto

<a id="canonical-482507c3328651ad9379174453d33a60f372b47c03d63dc5d8540619b091a8d9"></a>

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

<a id="canonical-c06e773335a5f1d5d5807fb00052f8e3de5348ba718aa2741d731cc5009c65e0"></a>

## Direct properties — ingress_egress_gw.hub.spoke_vnets.auto / 785ff5ed4055 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c7ed354e91477d4dfd4a466b7b437b9882aa6ed28c1043bab3d55d00388dbeb1"></a>

## Next pages — ingress_egress_gw.hub.spoke_vnets.auto / 785ff5ed4055 / 4

- [ingress_egress_gw.hub.spoke_vnets](resources--azure_vnet_site--reference--group-004.md#canonical-f8cced872d5b22379d7c2e9152ef94b4d0746e01dff9b62189c752cd445803b3)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-c2b61852e1fee481217ddbd475a4c80044916427a7b98591429c94830a3a16a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2604c993b9b0b3790aaee3968f3edcda59b4d1e4bce95da754dab26aff260a0a"></a>

## ingress_egress_gw.hub.spoke_vnets.labels — ingress_egress_gw.hub.spoke_vnets.labels / 705da575b936 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.spoke_vnets](resources--azure_vnet_site--reference--group-004.md#canonical-f8cced872d5b22379d7c2e9152ef94b4d0746e01dff9b62189c752cd445803b3)
- ingress_egress_gw.hub.spoke_vnets.labels

<a id="canonical-efaf2a6f8ca3e6184b5469a8987483c3b94b39f7d5e0b8565376aa0252812593"></a>

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

<a id="canonical-2b9e9c1d901249854baa1cccdb513a37269087f6ac01ac74250de63dfacc8680"></a>

## Direct properties — ingress_egress_gw.hub.spoke_vnets.labels / 705da575b936 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e5a820050864ec5230f17b6ea9f2ad23cf8ab96ecc801131b04b97ffd9dc4ef4"></a>

## Next pages — ingress_egress_gw.hub.spoke_vnets.labels / 705da575b936 / 4

- [ingress_egress_gw.hub.spoke_vnets](resources--azure_vnet_site--reference--group-004.md#canonical-f8cced872d5b22379d7c2e9152ef94b4d0746e01dff9b62189c752cd445803b3)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-16acfff6f6fd3f683aca3de881f5e785c3339380317e6444089d27e747040fdc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9e5c8c54305976ff0fefa5200aeca71b6a397dbb8fa0e160e38698a8dc51ab8"></a>

## ingress_egress_gw.hub.spoke_vnets.manual — ingress_egress_gw.hub.spoke_vnets.manual / 0904f9b1d745 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.spoke_vnets](resources--azure_vnet_site--reference--group-004.md#canonical-f8cced872d5b22379d7c2e9152ef94b4d0746e01dff9b62189c752cd445803b3)
- ingress_egress_gw.hub.spoke_vnets.manual

<a id="canonical-629eead760a420e5991871fd7720721c52d068aaad0d80aad985eef5652fa1d0"></a>

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

<a id="canonical-2ffec5c540994f186fa9342a4f3154643d5b5ef3800256209f9f227b723ecb9f"></a>

## Direct properties — ingress_egress_gw.hub.spoke_vnets.manual / 0904f9b1d745 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-30e8754ae3e605170ca1cf7a6b49d8e9a11c3d6c6034dc0adc3765d74b328d5b"></a>

## Next pages — ingress_egress_gw.hub.spoke_vnets.manual / 0904f9b1d745 / 4

- [ingress_egress_gw.hub.spoke_vnets](resources--azure_vnet_site--reference--group-004.md#canonical-f8cced872d5b22379d7c2e9152ef94b4d0746e01dff9b62189c752cd445803b3)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-b23785f031ff7eec3f52ede411ba9638edb96e868e038f64d9efdf3915a91bc8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bdbe172c0c09bcaa5571b8fb1fc155c1b919ce84f572a6d81be1877b32cec547"></a>

## ingress_egress_gw.hub.spoke_vnets.vnet — ingress_egress_gw.hub.spoke_vnets.vnet / 501ff478c614 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.spoke_vnets](resources--azure_vnet_site--reference--group-004.md#canonical-f8cced872d5b22379d7c2e9152ef94b4d0746e01dff9b62189c752cd445803b3)
- ingress_egress_gw.hub.spoke_vnets.vnet

<a id="canonical-4974138c3eebb544dac6aedc7a7fc1688f30a65f995a1adb1bcd7d54867fa8b9"></a>

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

<a id="canonical-26d103a1418c5c13b199b9f6a24fb007da160bc1f960fede336ada857e1e5afa"></a>

## Direct properties — ingress_egress_gw.hub.spoke_vnets.vnet / 501ff478c614 / 3

- [f5_orchestrated_routing](resources--azure_vnet_site--reference--group-004.md#canonical-b418c9247bf0c451bb462d5f265e0a25e49057eb8e92ade201d74318f802367a): complete subsection reference.

- [manual_routing](resources--azure_vnet_site--reference--group-004.md#canonical-44896317b67c09fecf62aca2b22d6a4b6f0d10624e0bba207d3f53bd5d2f3265): complete subsection reference.

<a id="canonical-87bc026cf7783c2239cbc664287ce341391d167c720ab865cb382a7db0f79ccb"></a>

<a id="canonical-a364da3c9aca1d76b9e8960f8d06c31010ce74fa6bca6c1f1f1b8a67e9629aab"></a>

## resource_group property — ingress_egress_gw.hub.spoke_vnets.vnet / 501ff478c614 / 4

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

<a id="canonical-1862c0130c2284a88e25d7e75a403e1d3161ffe081b9ed759ef40014722ace6a"></a>

<a id="canonical-3c01f0647f201d18c69cdb558090857bfb3478e19fae09662a724eb0f538b443"></a>

## vnet_name property — ingress_egress_gw.hub.spoke_vnets.vnet / 501ff478c614 / 5

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

<a id="canonical-1d6d59ba491713fcf10966633b9c950efb4e8c1d987decb64988755e7175edfc"></a>

## Next pages — ingress_egress_gw.hub.spoke_vnets.vnet / 501ff478c614 / 6

- [ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing](resources--azure_vnet_site--reference--group-004.md#canonical-b418c9247bf0c451bb462d5f265e0a25e49057eb8e92ade201d74318f802367a)
- [ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing](resources--azure_vnet_site--reference--group-004.md#canonical-44896317b67c09fecf62aca2b22d6a4b6f0d10624e0bba207d3f53bd5d2f3265)
- [ingress_egress_gw.hub.spoke_vnets](resources--azure_vnet_site--reference--group-004.md#canonical-f8cced872d5b22379d7c2e9152ef94b4d0746e01dff9b62189c752cd445803b3)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-b418c9247bf0c451bb462d5f265e0a25e49057eb8e92ade201d74318f802367a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-baa1e5fdd98ef8cb2835604eae566d3545c81b6b6f5a13646049661bd4cff42d"></a>

## ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing — ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing / 13c70d2bd398 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.spoke_vnets](resources--azure_vnet_site--reference--group-004.md#canonical-f8cced872d5b22379d7c2e9152ef94b4d0746e01dff9b62189c752cd445803b3)
- [ingress_egress_gw.hub.spoke_vnets.vnet](resources--azure_vnet_site--reference--group-004.md#canonical-b23785f031ff7eec3f52ede411ba9638edb96e868e038f64d9efdf3915a91bc8)
- ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing

<a id="canonical-74b87cc1225de957ad238d21484901046595cbbe02f44a1223a0f170976e65c0"></a>

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

<a id="canonical-60b17508c8467b4c65cab2ab6f3c483401cb0734873ad404c11ae7e39bd877c5"></a>

## Direct properties — ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing / 13c70d2bd398 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7b31217b5f4ddaf1d504b9e8c2069ffcb833ea1aad5291cf1c39a7e275a61506"></a>

## Next pages — ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing / 13c70d2bd398 / 4

- [ingress_egress_gw.hub.spoke_vnets.vnet](resources--azure_vnet_site--reference--group-004.md#canonical-b23785f031ff7eec3f52ede411ba9638edb96e868e038f64d9efdf3915a91bc8)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-44896317b67c09fecf62aca2b22d6a4b6f0d10624e0bba207d3f53bd5d2f3265"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aec187ede0eb97f6548ab360c621d66f9cf5f29a0bdde96d1ec2c88704477203"></a>

## ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing — ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing / 9aca000d9833 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.hub](resources--azure_vnet_site--reference--group-003.md#canonical-f0badbe55875efad13f9e3ccfb09ac96b2bfc72749944bf90c98d0b6a003b17a)
- [ingress_egress_gw.hub.spoke_vnets](resources--azure_vnet_site--reference--group-004.md#canonical-f8cced872d5b22379d7c2e9152ef94b4d0746e01dff9b62189c752cd445803b3)
- [ingress_egress_gw.hub.spoke_vnets.vnet](resources--azure_vnet_site--reference--group-004.md#canonical-b23785f031ff7eec3f52ede411ba9638edb96e868e038f64d9efdf3915a91bc8)
- ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing

<a id="canonical-f5181a3a80061f8dd13a0ee797a88bf97132d1e5eaac83897f7e9112b2aeb23e"></a>

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

<a id="canonical-51dcb72b633f0565bc3f01f62b9254d9dd17a28d6944370b99909f44fdd342cd"></a>

## Direct properties — ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing / 9aca000d9833 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1930a768efa4b1179ab4cc950f214453005981acbc802ce243d4f303baa09703"></a>

## Next pages — ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing / 9aca000d9833 / 4

- [ingress_egress_gw.hub.spoke_vnets.vnet](resources--azure_vnet_site--reference--group-004.md#canonical-b23785f031ff7eec3f52ede411ba9638edb96e868e038f64d9efdf3915a91bc8)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-0e085f32699a6dcb11563177e00bba4ef69f0be1ce651fcd655a9aefe833c4fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5752bccf6ed08d852cc9260fe320c672c26f581ab258e5073e0ce760a5273fe4"></a>

## ingress_egress_gw.inside_static_routes — ingress_egress_gw.inside_static_routes / 6f59b6d23c8f / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- ingress_egress_gw.inside_static_routes

<a id="canonical-3578fe458a28fe07ed060fbf617a4b0e4586e8db37fab6ef377170772e11ee54"></a>

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

<a id="canonical-0b2d87947dbb968c37e1031408bbb2abf551c2c1322214141962755458f49023"></a>

## Direct properties — ingress_egress_gw.inside_static_routes / 6f59b6d23c8f / 3

- [static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-654a731bccf29cd913f33ebe55a6062d87a5765369d708ecf24938c5774e65d4): complete subsection reference.

<a id="canonical-d220a6df08114c51f29299165e4636beb591dfc3b3bd529d3cdaf004ef6bc214"></a>

## Next pages — ingress_egress_gw.inside_static_routes / 6f59b6d23c8f / 4

- [ingress_egress_gw.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-654a731bccf29cd913f33ebe55a6062d87a5765369d708ecf24938c5774e65d4)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-654a731bccf29cd913f33ebe55a6062d87a5765369d708ecf24938c5774e65d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0426c4abe5256e4117886e1b992afad2c3b613af7a7bb25c11658567c31bc3fd"></a>

## ingress_egress_gw.inside_static_routes.static_route_list — ingress_egress_gw.inside_static_routes.static_route_list / 67d97600231e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.inside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-0e085f32699a6dcb11563177e00bba4ef69f0be1ce651fcd655a9aefe833c4fc)
- ingress_egress_gw.inside_static_routes.static_route_list

<a id="canonical-a4ca28915bf8d1f6e44cc90341670f16c215792b35e5c068901feab1382b3233"></a>

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

<a id="canonical-94f0928dcae735eb80706771ade7d2a1016fda304c24c29f62e667c23eead6b5"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list / 67d97600231e / 3

- [custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-f4da2f181499b62c32e9fd67d8bd8bd0fedf7873a34e74dc3ab5a82d076740ea): complete subsection reference.

<a id="canonical-7561ff7586c0c24897268c4deb30695931efe82870e16026bce9eba3bf64c965"></a>

<a id="canonical-7d7fbac090c071f86fb620f327b00afbd7ca5a3cb886c5fa0ee2c1d75ef75d04"></a>

## simple_static_route property — ingress_egress_gw.inside_static_routes.static_route_list / 67d97600231e / 4

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

<a id="canonical-0974c1c23b691434b20d43cc8a37ffe289c4b08c1efe689550c8e63f87cbf3d3"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list / 67d97600231e / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-f4da2f181499b62c32e9fd67d8bd8bd0fedf7873a34e74dc3ab5a82d076740ea)
- [ingress_egress_gw.inside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-0e085f32699a6dcb11563177e00bba4ef69f0be1ce651fcd655a9aefe833c4fc)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-f4da2f181499b62c32e9fd67d8bd8bd0fedf7873a34e74dc3ab5a82d076740ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6edefa829ed553350f2c8db8e02a8df35ddd0e244353c878db2c1fc7d4620171"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / 9dd85cfb23b9 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.inside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-0e085f32699a6dcb11563177e00bba4ef69f0be1ce651fcd655a9aefe833c4fc)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-654a731bccf29cd913f33ebe55a6062d87a5765369d708ecf24938c5774e65d4)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route

<a id="canonical-3f10240db090226f3ca14470485d557ee31067fd80faf8acc14cf4b8747e1018"></a>

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

<a id="canonical-fbf2b8414ea6468b85424d4624dfbbad2cd9f8fb64644d0a6e681b66606f6bf2"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / 9dd85cfb23b9 / 3

<a id="canonical-788d04e3fb7dd9110f4310fa69e0280e5d22793e317657ab58e8c1c30b517e6b"></a>

<a id="canonical-5e3bb6d33039b371235e7e2f0162dff3c1fd1cd0016c0cb80b17361e813aeb3d"></a>

## attrs property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / 9dd85cfb23b9 / 4

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

- [labels](resources--azure_vnet_site--reference--group-004.md#canonical-2b589c103c46178e7552deff462def6fe9c032af947ce3635a00baa80def0d84): complete subsection reference.

- [nexthop](resources--azure_vnet_site--reference--group-004.md#canonical-b6c0c4dfd7d0fd51c2ada37d8045145a62914f02caa209e41737e6e6be74a87d): complete subsection reference.

- [subnets](resources--azure_vnet_site--reference--group-004.md#canonical-c5a728cc40e29d823e28c4e94a3a553179425ef5b922e23d029f02f9685c17b2): complete subsection reference.

<a id="canonical-5c504b8320944a3f45d3e73abf8ab51dab1a6bd65d87b997e9aa66acb2dfce0f"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / 9dd85cfb23b9 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](resources--azure_vnet_site--reference--group-004.md#canonical-2b589c103c46178e7552deff462def6fe9c032af947ce3635a00baa80def0d84)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-004.md#canonical-b6c0c4dfd7d0fd51c2ada37d8045145a62914f02caa209e41737e6e6be74a87d)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-004.md#canonical-c5a728cc40e29d823e28c4e94a3a553179425ef5b922e23d029f02f9685c17b2)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-654a731bccf29cd913f33ebe55a6062d87a5765369d708ecf24938c5774e65d4)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-2b589c103c46178e7552deff462def6fe9c032af947ce3635a00baa80def0d84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0524960e527cc0d16c147bb5c1a85956ed9e5357d885bdb0d9ac1ecc7d413ee"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.lab / 14bc8fb89830 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.inside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-0e085f32699a6dcb11563177e00bba4ef69f0be1ce651fcd655a9aefe833c4fc)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-654a731bccf29cd913f33ebe55a6062d87a5765369d708ecf24938c5774e65d4)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-f4da2f181499b62c32e9fd67d8bd8bd0fedf7873a34e74dc3ab5a82d076740ea)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-ac76a48a7f6ca222d92c037a132f9413f42649aede1fef9127d37e9086bcf4dd"></a>

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

<a id="canonical-cd8462178785ca4e8f08a2c0e54fe9caecb572bc5d6c47e13b1ed07d0bf3e9b5"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.lab / 14bc8fb89830 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fe8049dec9faff42ad52e052f963285e95c4111e9ff4a6e622f80f0f8cf49896"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.lab / 14bc8fb89830 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-f4da2f181499b62c32e9fd67d8bd8bd0fedf7873a34e74dc3ab5a82d076740ea)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-b6c0c4dfd7d0fd51c2ada37d8045145a62914f02caa209e41737e6e6be74a87d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48a93f608c6ebe45d91fd3dc6d6e43e7fd498f40f60da68c88e94a0432adc7ba"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 2de1205ae246 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.inside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-0e085f32699a6dcb11563177e00bba4ef69f0be1ce651fcd655a9aefe833c4fc)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-654a731bccf29cd913f33ebe55a6062d87a5765369d708ecf24938c5774e65d4)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-f4da2f181499b62c32e9fd67d8bd8bd0fedf7873a34e74dc3ab5a82d076740ea)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-9f92d25f082b995a176f2973017566c112a2b94a95100e890771099e8451df6c"></a>

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

<a id="canonical-01d64beb881eec86c129bdc5d024bdd00625f34a6ea9e82d99ea604d07c855fc"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 2de1205ae246 / 3

- [interface](resources--azure_vnet_site--reference--group-004.md#canonical-01f18a094c203557588814018490e66c261e991faa2f2437d5160329eb88a3dd): complete subsection reference.

- [nexthop_address](resources--azure_vnet_site--reference--group-004.md#canonical-5e0118763bcad783c319589c83e5a9939a6d4c93ddac43087b878f479ebd9eca): complete subsection reference.

<a id="canonical-fcf330abcc47aa94bd5dd20b4df307fecc64afa03466490ff0e80487e0fba42f"></a>

<a id="canonical-0209ffa63a34963624645615a2b0031d37bd7f5f1e08e5e9e323db5ee81c1f14"></a>

## type property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 2de1205ae246 / 4

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

<a id="canonical-606a915eeee7360bfdedb5b4abfe06e6df13c06c1844ae757c5972695759a27d"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 2de1205ae246 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--azure_vnet_site--reference--group-004.md#canonical-01f18a094c203557588814018490e66c261e991faa2f2437d5160329eb88a3dd)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-004.md#canonical-5e0118763bcad783c319589c83e5a9939a6d4c93ddac43087b878f479ebd9eca)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-f4da2f181499b62c32e9fd67d8bd8bd0fedf7873a34e74dc3ab5a82d076740ea)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-01f18a094c203557588814018490e66c261e991faa2f2437d5160329eb88a3dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ede0e168f9a99fca1a1f2b9e21b5f15d2ddbb7f8a10efe1c03ca7bfb2767adf8"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / d513c3e99908 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.inside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-0e085f32699a6dcb11563177e00bba4ef69f0be1ce651fcd655a9aefe833c4fc)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-654a731bccf29cd913f33ebe55a6062d87a5765369d708ecf24938c5774e65d4)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-f4da2f181499b62c32e9fd67d8bd8bd0fedf7873a34e74dc3ab5a82d076740ea)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-004.md#canonical-b6c0c4dfd7d0fd51c2ada37d8045145a62914f02caa209e41737e6e6be74a87d)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-73ac8cb4a6d258153ac840a57ec76fd35cbe9c9682c5fa8e1e217944ea7c706a"></a>

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

<a id="canonical-108d3cf3236c2360814c69a8c6940ad04153641873c331cb2829549433890e8f"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / d513c3e99908 / 3

<a id="canonical-72dad42b61c211b4b7fce92ecf6a98eb7f8a433c0fd30aa485b16b1630f98545"></a>

<a id="canonical-61330a95688de4e05f26a6057a27b184078f71e4453e6ad8576a8bc120c3d07f"></a>

## kind property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / d513c3e99908 / 4

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

<a id="canonical-983e4e8cc71de3382f848377afc301d1844b2ef2cf5264744e658084217267fe"></a>

<a id="canonical-0e10280597f18bd617980b293367d0f93e346fc783bd42d842428c26ea6bd2b5"></a>

## name property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / d513c3e99908 / 5

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

<a id="canonical-04d63fff64fba78d2d47b4668ec76409a397f9101950f53f7585e6aaa2780101"></a>

<a id="canonical-83b1058983c804ad744463298e699095e4f2914952a842c718177ed47ad1bec5"></a>

## namespace property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / d513c3e99908 / 6

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

<a id="canonical-b041364e612a3b97779a5db8cafe54353249752ad1c8e423253dd9f69d483e6b"></a>

<a id="canonical-114c662b31c2f099628234d0f951e09f5f2e9145203e4781c32dd2b1fd4b69e6"></a>

## tenant property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / d513c3e99908 / 7

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

<a id="canonical-e31341cbb775c7d32c1402fda448d133ab1dfe7c9f94b5ddd46b1b37c5ef4328"></a>

<a id="canonical-ac4e51d5a934c745cf26fedc15101fdda6b0dd4ffc80e80891aba1d79a024b61"></a>

## uid property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / d513c3e99908 / 8

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

<a id="canonical-6d6bd9206422fe488a80f8c2418012896573f96dc9aa39fb4425e125e6121020"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / d513c3e99908 / 9

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-004.md#canonical-b6c0c4dfd7d0fd51c2ada37d8045145a62914f02caa209e41737e6e6be74a87d)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-5e0118763bcad783c319589c83e5a9939a6d4c93ddac43087b878f479ebd9eca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e499dba0c49fa2f830191ae7d6f0d9160b8692879fb1c2301750a730b746dbe9"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 11f3318ac131 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.inside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-0e085f32699a6dcb11563177e00bba4ef69f0be1ce651fcd655a9aefe833c4fc)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-654a731bccf29cd913f33ebe55a6062d87a5765369d708ecf24938c5774e65d4)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-f4da2f181499b62c32e9fd67d8bd8bd0fedf7873a34e74dc3ab5a82d076740ea)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-004.md#canonical-b6c0c4dfd7d0fd51c2ada37d8045145a62914f02caa209e41737e6e6be74a87d)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-0aaa0cfe60e6668f2c4361ce25b8ef5c39c972d9fcfa285f635132a85dd14f2a"></a>

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

<a id="canonical-3f6a20b55329349f8a8dca095b152492090866c0a7b0d1d4d3635a1dbc96abfd"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 11f3318ac131 / 3

- [dual_stack](resources--azure_vnet_site--reference--group-004.md#canonical-85124f4339ddf056c8b6191b7a562313acc7f9ffa9c885ae43d73f16527803cc): complete subsection reference.

- [ipv4](resources--azure_vnet_site--reference--group-004.md#canonical-3328d33f7c9a285951852319e278cda223cb6c0e70ca8f57a512c32ae5edc207): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-004.md#canonical-32773985f4014b0285043e7d43645b65b5a488343e7523e3546bb85a81abda2a): complete subsection reference.

<a id="canonical-57b00ec9dd46cb17ccb913e37e1a9968a2fe29f6af05a03225a5470bcd34b8e5"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 11f3318ac131 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-004.md#canonical-85124f4339ddf056c8b6191b7a562313acc7f9ffa9c885ae43d73f16527803cc)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--azure_vnet_site--reference--group-004.md#canonical-3328d33f7c9a285951852319e278cda223cb6c0e70ca8f57a512c32ae5edc207)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--azure_vnet_site--reference--group-004.md#canonical-32773985f4014b0285043e7d43645b65b5a488343e7523e3546bb85a81abda2a)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-004.md#canonical-b6c0c4dfd7d0fd51c2ada37d8045145a62914f02caa209e41737e6e6be74a87d)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-85124f4339ddf056c8b6191b7a562313acc7f9ffa9c885ae43d73f16527803cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f27e5ec69c7e589cb77c4a657343476b147cb0bc896e653deaa0f25ace84dd33"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 70f5c254c75b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.inside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-0e085f32699a6dcb11563177e00bba4ef69f0be1ce651fcd655a9aefe833c4fc)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-654a731bccf29cd913f33ebe55a6062d87a5765369d708ecf24938c5774e65d4)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-f4da2f181499b62c32e9fd67d8bd8bd0fedf7873a34e74dc3ab5a82d076740ea)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-004.md#canonical-b6c0c4dfd7d0fd51c2ada37d8045145a62914f02caa209e41737e6e6be74a87d)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-004.md#canonical-5e0118763bcad783c319589c83e5a9939a6d4c93ddac43087b878f479ebd9eca)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-4f709d1ea3b33e544fa8f2b7277ba5263680739981d117ecc37ea5773d7474c2"></a>

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

<a id="canonical-c03a358b0fbb4ffd4ccbb9cb6abe970a06c591cf2062339739682eeb643fb1f3"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 70f5c254c75b / 3

- [ipv4](resources--azure_vnet_site--reference--group-004.md#canonical-cacb94fa1ea1d3285748b6b651caf7af2ded4ed9c31bbd9fb98c2107c0bc249d): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-004.md#canonical-5a85afa318356145dd154e63b6082d5a0e8d1bf55fd5f27dacd487e511e5dd6f): complete subsection reference.

<a id="canonical-9fc64de352f8cf57c2182edd7525f32af0fdccf6a2a83d3624c990918d01d912"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 70f5c254c75b / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--azure_vnet_site--reference--group-004.md#canonical-cacb94fa1ea1d3285748b6b651caf7af2ded4ed9c31bbd9fb98c2107c0bc249d)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--azure_vnet_site--reference--group-004.md#canonical-5a85afa318356145dd154e63b6082d5a0e8d1bf55fd5f27dacd487e511e5dd6f)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-004.md#canonical-5e0118763bcad783c319589c83e5a9939a6d4c93ddac43087b878f479ebd9eca)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-cacb94fa1ea1d3285748b6b651caf7af2ded4ed9c31bbd9fb98c2107c0bc249d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a06d4ceec768eda257808473fd0bd1132f745ed75f8e6e1dd4453fd47d282fd"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / ff5edd24a5d7 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.inside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-0e085f32699a6dcb11563177e00bba4ef69f0be1ce651fcd655a9aefe833c4fc)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-654a731bccf29cd913f33ebe55a6062d87a5765369d708ecf24938c5774e65d4)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-f4da2f181499b62c32e9fd67d8bd8bd0fedf7873a34e74dc3ab5a82d076740ea)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-004.md#canonical-b6c0c4dfd7d0fd51c2ada37d8045145a62914f02caa209e41737e6e6be74a87d)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-004.md#canonical-5e0118763bcad783c319589c83e5a9939a6d4c93ddac43087b878f479ebd9eca)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-004.md#canonical-85124f4339ddf056c8b6191b7a562313acc7f9ffa9c885ae43d73f16527803cc)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-7e4a87da78a213b37f1f4718ee016a6bbce74fb048c5abd94238ff77a10911d5"></a>

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

<a id="canonical-bb7ceecd3e58b9a8351bba1ece8af49945858e3413c617607082c6b62e6d20b2"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / ff5edd24a5d7 / 3

<a id="canonical-94b75dab51e9cca85a8af12627f5ca3d9163b4911094f6bf26c58aaf33d1a84a"></a>

<a id="canonical-ea3f62ddfd7433b64c477737dad97b67cac03b4eec01ee7584342dc01560dd61"></a>

## addr property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / ff5edd24a5d7 / 4

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

<a id="canonical-209873102f29ed892aeeff4d3ee80014b0351fa0fd252c609a72bdfb753b33e4"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / ff5edd24a5d7 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-004.md#canonical-85124f4339ddf056c8b6191b7a562313acc7f9ffa9c885ae43d73f16527803cc)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-5a85afa318356145dd154e63b6082d5a0e8d1bf55fd5f27dacd487e511e5dd6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e5db1ba3459ecf50758df0b77aac3cae950dff625de62a1c26241011dbb32ea2"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / d32f26cfb1b1 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.inside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-0e085f32699a6dcb11563177e00bba4ef69f0be1ce651fcd655a9aefe833c4fc)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-654a731bccf29cd913f33ebe55a6062d87a5765369d708ecf24938c5774e65d4)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-f4da2f181499b62c32e9fd67d8bd8bd0fedf7873a34e74dc3ab5a82d076740ea)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-004.md#canonical-b6c0c4dfd7d0fd51c2ada37d8045145a62914f02caa209e41737e6e6be74a87d)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-004.md#canonical-5e0118763bcad783c319589c83e5a9939a6d4c93ddac43087b878f479ebd9eca)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-004.md#canonical-85124f4339ddf056c8b6191b7a562313acc7f9ffa9c885ae43d73f16527803cc)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-0228a6fc406b8cbc4f310cf3743f4534a953bb327de5eff7d91ab851c352ca8e"></a>

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

<a id="canonical-961bd2c055697c633526b2171d6fd2a2ef3c021eefdda945541c4101303c7434"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / d32f26cfb1b1 / 3

<a id="canonical-e7e1b4ebdb98ad25e7834583c272540dac77e2ee0991a8a1119f5185ef75660e"></a>

<a id="canonical-75eede38b6f6bc4082ba47a192b172b6bb98a2667a4794e4950c06857215a6f9"></a>

## addr property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / d32f26cfb1b1 / 4

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

<a id="canonical-440c05b9dd45fc17b83c388417a38e08e4615a498bc7769cd161771c8987ec37"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / d32f26cfb1b1 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-004.md#canonical-85124f4339ddf056c8b6191b7a562313acc7f9ffa9c885ae43d73f16527803cc)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-3328d33f7c9a285951852319e278cda223cb6c0e70ca8f57a512c32ae5edc207"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-965fd7e7e13e0d65b4cd03fcbf441919a569ef249e3cbfa037980ae69adcfd2a"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 30f42ee1187b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.inside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-0e085f32699a6dcb11563177e00bba4ef69f0be1ce651fcd655a9aefe833c4fc)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-654a731bccf29cd913f33ebe55a6062d87a5765369d708ecf24938c5774e65d4)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-f4da2f181499b62c32e9fd67d8bd8bd0fedf7873a34e74dc3ab5a82d076740ea)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-004.md#canonical-b6c0c4dfd7d0fd51c2ada37d8045145a62914f02caa209e41737e6e6be74a87d)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-004.md#canonical-5e0118763bcad783c319589c83e5a9939a6d4c93ddac43087b878f479ebd9eca)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-04692ffe155367e57430b63ddf2d858a3dba44018b4ddfb2d954ed5f85926bcd"></a>

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

<a id="canonical-330a2ab1c1675b36dd4f7366b30be191fe54a27b523d576cb270a42be70dea18"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 30f42ee1187b / 3

<a id="canonical-c060a06be1134e670a45f43216f52cc042c5dc23b61b5c7a641d2438d55b544c"></a>

<a id="canonical-bdd0865f1f96a19316aaa5d96b451cddb9c40ecc8fcf2eb1330521be476484cc"></a>

## addr property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 30f42ee1187b / 4

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

<a id="canonical-5415c0c9ea5360fbdb1363ea58ef5acd5e2863a1f0b5f173ebe288705c82de77"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 30f42ee1187b / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-004.md#canonical-5e0118763bcad783c319589c83e5a9939a6d4c93ddac43087b878f479ebd9eca)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-32773985f4014b0285043e7d43645b65b5a488343e7523e3546bb85a81abda2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a29b71c0910f218a0e96fdfe33a2042a34ec12baeb4be810e11a74226758a44"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 7dad1e6ca25b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.inside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-0e085f32699a6dcb11563177e00bba4ef69f0be1ce651fcd655a9aefe833c4fc)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-654a731bccf29cd913f33ebe55a6062d87a5765369d708ecf24938c5774e65d4)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-f4da2f181499b62c32e9fd67d8bd8bd0fedf7873a34e74dc3ab5a82d076740ea)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-004.md#canonical-b6c0c4dfd7d0fd51c2ada37d8045145a62914f02caa209e41737e6e6be74a87d)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-004.md#canonical-5e0118763bcad783c319589c83e5a9939a6d4c93ddac43087b878f479ebd9eca)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-178db7da27f7dd3ac99484eba92887a7d6386e62a3896273916eb43ea4c75f5b"></a>

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

<a id="canonical-a665f1fa05c39d6593b9990740cf8b9026f355b1e4a9d413423f841786076b26"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 7dad1e6ca25b / 3

<a id="canonical-6e3e6ebf2efdb750b2713a8d943db42f62328f91326e9a07cc8babd381841669"></a>

<a id="canonical-c83b69f7ccc8a9281f6bf9159fb86a9c68f46ee33f0aa6264b692327f90367a7"></a>

## addr property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 7dad1e6ca25b / 4

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

<a id="canonical-992bb34d700a2718f99840058d7b95019a9e3c3982d8fcf1354a53062bad929d"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 7dad1e6ca25b / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-004.md#canonical-5e0118763bcad783c319589c83e5a9939a6d4c93ddac43087b878f479ebd9eca)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-c5a728cc40e29d823e28c4e94a3a553179425ef5b922e23d029f02f9685c17b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f34ab47371417a18253f3e3d9c7ccb1c213d89bbc85fa51c2c17096424f2fbdd"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 3370c8f183ab / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.inside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-0e085f32699a6dcb11563177e00bba4ef69f0be1ce651fcd655a9aefe833c4fc)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-654a731bccf29cd913f33ebe55a6062d87a5765369d708ecf24938c5774e65d4)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-f4da2f181499b62c32e9fd67d8bd8bd0fedf7873a34e74dc3ab5a82d076740ea)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-c19f1d3c0694839f839f93d66297184263dd842c545568dc91875e08ead8c82d"></a>

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

<a id="canonical-5447881f55679fcdeeb3dda88c1bc5a8a31291096434edd3a856a1b49dd81a61"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 3370c8f183ab / 3

- [ipv4](resources--azure_vnet_site--reference--group-004.md#canonical-29c2b5a660048f186c80aa41cfb5e9de71ceca621087a7e79fc1eab5706aa71e): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-004.md#canonical-3c8b93305c1bca4672d5b5fac979b1fc1471f2c0484999e803d7da422cf10508): complete subsection reference.

<a id="canonical-2b68644d26be58de80e980aec426ee5d72b02042e8963a3aeca7885fd532bb96"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 3370c8f183ab / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--azure_vnet_site--reference--group-004.md#canonical-29c2b5a660048f186c80aa41cfb5e9de71ceca621087a7e79fc1eab5706aa71e)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--azure_vnet_site--reference--group-004.md#canonical-3c8b93305c1bca4672d5b5fac979b1fc1471f2c0484999e803d7da422cf10508)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-f4da2f181499b62c32e9fd67d8bd8bd0fedf7873a34e74dc3ab5a82d076740ea)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-29c2b5a660048f186c80aa41cfb5e9de71ceca621087a7e79fc1eab5706aa71e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-82a6adf90047cc574ab61c7a1a1bca7903ccf2b084763bec4b27d205ea6170cd"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 75634473db58 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.inside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-0e085f32699a6dcb11563177e00bba4ef69f0be1ce651fcd655a9aefe833c4fc)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-654a731bccf29cd913f33ebe55a6062d87a5765369d708ecf24938c5774e65d4)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-f4da2f181499b62c32e9fd67d8bd8bd0fedf7873a34e74dc3ab5a82d076740ea)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-004.md#canonical-c5a728cc40e29d823e28c4e94a3a553179425ef5b922e23d029f02f9685c17b2)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-bef5961d586a9144e282c5cc26cb7b0bae96a49eb6d13d1d30e21e9275b2d17a"></a>

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

<a id="canonical-dd41c5a38c64a23b7c758ddc301faf70ae8f7b0340e220086432ad14df11ad4b"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 75634473db58 / 3

<a id="canonical-2b88f50e9fc39aff92334527aa8e6ed5d042237dee0d6523c5ec1e732727021a"></a>

<a id="canonical-f96d0c4fc26ce84504a5aa36bdcd100b596c0f0a6111345462d1281d58e3a443"></a>

## plen property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 75634473db58 / 4

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

<a id="canonical-8015ecd44c239654497e5320b9954f6805525f1427be9b868d19bed32e24e4cf"></a>

<a id="canonical-f922de2c376cd99cc4ba06329ce7728602e73ed9e4281261e2c7d44bd21561be"></a>

## prefix property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 75634473db58 / 5

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

<a id="canonical-bfadd3a221e706054f7b5497a234ae75cbafaa90d8560671349dd913eb64356d"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 75634473db58 / 6

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-004.md#canonical-c5a728cc40e29d823e28c4e94a3a553179425ef5b922e23d029f02f9685c17b2)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-3c8b93305c1bca4672d5b5fac979b1fc1471f2c0484999e803d7da422cf10508"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e66c35c29d9cd2d015f59c3f907e472b631166e5774be56b6b463fd9715ac48a"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 557f449d2dbf / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.inside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-0e085f32699a6dcb11563177e00bba4ef69f0be1ce651fcd655a9aefe833c4fc)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-654a731bccf29cd913f33ebe55a6062d87a5765369d708ecf24938c5774e65d4)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-f4da2f181499b62c32e9fd67d8bd8bd0fedf7873a34e74dc3ab5a82d076740ea)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-004.md#canonical-c5a728cc40e29d823e28c4e94a3a553179425ef5b922e23d029f02f9685c17b2)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-f46c623a86a2522a8529aa89d1641ce0b2545d86479e35a27628914dfc39229d"></a>

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

<a id="canonical-6373cd4fed682222a5e7ddb3a214aff06cabf973643207e692c574497f6292fe"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 557f449d2dbf / 3

<a id="canonical-d99c4272987924a6e5220cbbd4624a50ceb5b33207dda7eed0026cd407e3c534"></a>

<a id="canonical-3ebaccb64d5ff5f8de1dc73b736e618622f42f0aa03534d4ff65754730a07af4"></a>

## plen property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 557f449d2dbf / 4

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

<a id="canonical-a4968ae26176a291f0e91d904aeb05591c9ba7479ac7d33f8ef53d169bd27fff"></a>

<a id="canonical-74f99b3519c32ae04bea468d927442589e1c2eb510d5019f10701dbcb4a1c0f6"></a>

## prefix property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 557f449d2dbf / 5

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

<a id="canonical-10c8136eff95c5e27213544ca098bd8ddd3760a38a20ece28f9fd4db79032934"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 557f449d2dbf / 6

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-004.md#canonical-c5a728cc40e29d823e28c4e94a3a553179425ef5b922e23d029f02f9685c17b2)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-23723538e4cf5e88cd06b08a4af2bc7aeb0ec7df4e5f05dc9746ace9ce12d239"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6109a570ebc1be9dd4f56b5ddcdd274692225d27c5a84757cc7e8af92e1dd54e"></a>

## ingress_egress_gw.no_dc_cluster_group — ingress_egress_gw.no_dc_cluster_group / ea0e9a248f3e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- ingress_egress_gw.no_dc_cluster_group

<a id="canonical-dbca2ef6af46bb22d4ffa15a1b5cbab79fdf9fa4b8f002009373c4ff97049d5a"></a>

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

<a id="canonical-97d8283ad927a8436017243d292da4eb33edb5a4a343f5e983d6d69c437b64b9"></a>

## Direct properties — ingress_egress_gw.no_dc_cluster_group / ea0e9a248f3e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e74b11880a33467e349647da776a8eb136b42b5f49d5cf1dd2d29f686f150e53"></a>

## Next pages — ingress_egress_gw.no_dc_cluster_group / ea0e9a248f3e / 4

- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-ea31407c8dfc25365176830e3ee1842d6431014e0e5c8d8d31f65d93ee7c0bda"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0f4a5fd58d4a09087e04f947ccb298b13f4c4ec6a07606ec717bf57a4cd0f16"></a>

## ingress_egress_gw.no_forward_proxy — ingress_egress_gw.no_forward_proxy / 9b05d5d5340e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- ingress_egress_gw.no_forward_proxy

<a id="canonical-a03e990eb069651d98c438ff1c6c9546a1a393f3eb83c33e386f429439a00dca"></a>

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

<a id="canonical-16f241856785fec64b88084ab357b84b46366a3871b7c5591e8d5068ba1ab554"></a>

## Direct properties — ingress_egress_gw.no_forward_proxy / 9b05d5d5340e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4c47ccf6c3a83277dfb696ea78706fe1e75f4c37fca4baecec54653341782ce8"></a>

## Next pages — ingress_egress_gw.no_forward_proxy / 9b05d5d5340e / 4

- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-bf45ee4ee97e911082f35b39ca244cf2f494596a1f1ad2780c75f5da59783c86"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6fad02651476746def58fbaefac9f3ebf0a9e64b3ade239c6868b67d777db176"></a>

## ingress_egress_gw.no_global_network — ingress_egress_gw.no_global_network / 1bc7c1a66b53 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- ingress_egress_gw.no_global_network

<a id="canonical-30389884057d364e89a93f71c8410f53f6644e2fe992e0f546f9cebf1a3fa983"></a>

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

<a id="canonical-f18276e484c7d3d3abb7f3cf826d4faf52f251f0d67fb7a8d7f88c6035140bf4"></a>

## Direct properties — ingress_egress_gw.no_global_network / 1bc7c1a66b53 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b8fdc685b3701e8ad2440634e15b70c3259c8419ecd7b570935a4c92d040d586"></a>

## Next pages — ingress_egress_gw.no_global_network / 1bc7c1a66b53 / 4

- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-bbef5af7f643ab0709b5886f4b6c98b8e23ac6d26a2b0d7b1a1b85bbbee0917f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88d20419488e58e8f82d34930a3f07c0ef4bf6e75ab9ccd2cf00c78a687bbe58"></a>

## ingress_egress_gw.no_inside_static_routes — ingress_egress_gw.no_inside_static_routes / 63a9c2c91dcb / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- ingress_egress_gw.no_inside_static_routes

<a id="canonical-6a514e998e9c3e14098583c0c46d3b7a87ac5a0adb9982097fe49fa05b1d46fe"></a>

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

<a id="canonical-388f38a200edb8bc27a9b8f046ff44d141d823dc8edbf8001909f4e27ec95263"></a>

## Direct properties — ingress_egress_gw.no_inside_static_routes / 63a9c2c91dcb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e60669db5be3cce3dc68d487a81c34b8abb03ad98298dd08f6f40c1017efcb5d"></a>

## Next pages — ingress_egress_gw.no_inside_static_routes / 63a9c2c91dcb / 4

- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-e30ef0e5e0c317c1a910c4b6310bceb6aad459ba73713b0ade29ffd0abdb3324"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f81c1457b0e52876e7bf87287f4677f9c87b026aca05fd213c620154974745c"></a>

## ingress_egress_gw.no_network_policy — ingress_egress_gw.no_network_policy / 9504dd0ca3d0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- ingress_egress_gw.no_network_policy

<a id="canonical-ee3cfe4f149649604ed8ad96627c39e081a13125ae28dc38a5896001f6cf6d11"></a>

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

<a id="canonical-64e22f6711e6a782bfc0f92bbf0cee6acecf91b3199db1a56e27b9fc667c7118"></a>

## Direct properties — ingress_egress_gw.no_network_policy / 9504dd0ca3d0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-69c6ba801dea96b49ff2872deb841722cd8bf3173939b0008c859dfcc091b7ca"></a>

## Next pages — ingress_egress_gw.no_network_policy / 9504dd0ca3d0 / 4

- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-8133d5ff2c93a8ac6a894d8da09060f464fae672cc8d94b01d38221524e851ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52ccaa402df5dc7b0d3d739212946050120be18f05770606996ff7701810221c"></a>

## ingress_egress_gw.no_outside_static_routes — ingress_egress_gw.no_outside_static_routes / 370dfc07152b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- ingress_egress_gw.no_outside_static_routes

<a id="canonical-555d1cfad5a14bcbd95ac8164655e92dcb82ea7647e655d3ffa01d13199b9441"></a>

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

<a id="canonical-c5c1ea51ed8747c12388dbd18611cd8827f632b9990c3b1a5c2805cbf5fbd2b1"></a>

## Direct properties — ingress_egress_gw.no_outside_static_routes / 370dfc07152b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-900168f2f580265c40e36a37deeb686d9bdc985b96236cd1e76bc748d0c374e5"></a>

## Next pages — ingress_egress_gw.no_outside_static_routes / 370dfc07152b / 4

- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-9a1f7c492923bf7fc978a0a5249f52f5b70111d387ccdc15d3ffed775abe16d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5276277fa74c366022319ffcf71999a67c43d4604bca8b709ef205dc20e3117e"></a>

## ingress_egress_gw.not_hub — ingress_egress_gw.not_hub / ec09e2febe85 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- ingress_egress_gw.not_hub

<a id="canonical-37ab1a7f37cca8c91f1300fa2ef489a8102bab82711fd4e7ed062c1f314a4a92"></a>

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
not_hub = {}
```

<a id="canonical-02b7c99711bbb4d55233f2503d5ec1482e453ba83d612877e91b9d83b3a89ce6"></a>

## Direct properties — ingress_egress_gw.not_hub / ec09e2febe85 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3a3f10e00774dc4f7e0086e164030a41b17c90739b4a14df911d96f91cf26bd2"></a>

## Next pages — ingress_egress_gw.not_hub / ec09e2febe85 / 4

- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-b2cf69d861e0a57ea1c50f0929ee2b476c5e3e8282d0f7a92aa6d14e7b6faa6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2fc566f93563da8e8dcf42e552ddf0f794d8c49b661ce55ce4cdc04dea31d414"></a>

## ingress_egress_gw.outside_static_routes — ingress_egress_gw.outside_static_routes / 930b6333e696 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- ingress_egress_gw.outside_static_routes

<a id="canonical-e6a6a5cafd03fb40a69d8d4a0e75a87903453ea782d9cbe72ae27c89837ba9f0"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside static routes.

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
outside_static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-221949e8b6cc4edac0cc7114daa980409b405840061153a137b8c34c49ab6442"></a>

## Direct properties — ingress_egress_gw.outside_static_routes / 930b6333e696 / 3

- [static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-18c00d1739a185869280757aedf5629b9dbe3f6360e3ae5c1315ab3868f8e1bb): complete subsection reference.

<a id="canonical-798f0a48f9d485800efe1a8bd6ce99e73decfc33bb4c3efd6aef0423653c04db"></a>

## Next pages — ingress_egress_gw.outside_static_routes / 930b6333e696 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-18c00d1739a185869280757aedf5629b9dbe3f6360e3ae5c1315ab3868f8e1bb)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-18c00d1739a185869280757aedf5629b9dbe3f6360e3ae5c1315ab3868f8e1bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e59882cdd7347d05ac2191f8c068ea5cece137b2dd44272844de4faedcf1e464"></a>

## ingress_egress_gw.outside_static_routes.static_route_list — ingress_egress_gw.outside_static_routes.static_route_list / 7ab904515d53 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.outside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-b2cf69d861e0a57ea1c50f0929ee2b476c5e3e8282d0f7a92aa6d14e7b6faa6a)
- ingress_egress_gw.outside_static_routes.static_route_list

<a id="canonical-a8ec4fddd2f9d5b8ab0a977cd473099717bdb46df06421a3bd9d714c46205904"></a>

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

<a id="canonical-80a773c006996aca23282f45ce881650dc809dd02e0dc7644647ea76cb77d311"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list / 7ab904515d53 / 3

- [custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-1c64750fca39996634713907f2ecaa89ef1367766e0469bd3e57de47cc089f81): complete subsection reference.

<a id="canonical-e85166af9894fae7bd7ba498d61005700dd8a6aa9eac7bf8fd57595d3492b626"></a>

<a id="canonical-9ea788556fdeaa79085e6bf2aa2175d63447943723e6d3e4ab842723bc11659a"></a>

## simple_static_route property — ingress_egress_gw.outside_static_routes.static_route_list / 7ab904515d53 / 4

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

<a id="canonical-16768375465a981432de227da193cfbfbf5c508a7249d2201d6749af44ad00f6"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list / 7ab904515d53 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-1c64750fca39996634713907f2ecaa89ef1367766e0469bd3e57de47cc089f81)
- [ingress_egress_gw.outside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-b2cf69d861e0a57ea1c50f0929ee2b476c5e3e8282d0f7a92aa6d14e7b6faa6a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-1c64750fca39996634713907f2ecaa89ef1367766e0469bd3e57de47cc089f81"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
