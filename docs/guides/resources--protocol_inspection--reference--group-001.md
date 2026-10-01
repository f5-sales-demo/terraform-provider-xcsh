---
page_title: "xcsh_protocol_inspection reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_inspection reference."
---

# xcsh_protocol_inspection reference

<a id="canonical-e4364e9cdf919f8070e0c9ff68ce35d1ccf1e8b653ca9312621219bdcca91de8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d76174516b46527fe07dec0c13f4a9bf8b1ff3f35860d5c4c02330a4ef40172a"></a>

## Property reference — Property reference / 98dcc9bbb9bd / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)
- Property reference

<a id="canonical-a8b55ffd6a9b23ceb9ba0015ca24fe2c78cdf939076507a596d5e29603f8a3e2"></a>

## Direct properties — Property reference / 98dcc9bbb9bd / 3

<a id="canonical-dbe152e2908953ef95d20aeae60261881dfa67b889cd47514a684249972a4f1d"></a>

<a id="canonical-d6c82f8837b40a5555d45b1baf5a9942bb0080cb91cbbcc751bb57d3a8fdc517"></a>

## action property — Property reference / 98dcc9bbb9bd / 4

Type: `"string"`. Optional, Computed.

\[Enum: ALLOW|DENY|DROP\] Action after inspection - ALLOW: Allow Allow traffic - DENY: Deny Throw
RST error for TCP and ICMP error for UDP - DROP: DROP Silently drop traffic. Possible values are
\`ALLOW\`, \`DENY\`, \`DROP\`. Defaults to \`ALLOW\`. Server applies default when omitted.

Upstream description:

Action after inspection

&#8203;- ALLOW: Allow

Allow traffic &#8203;- DENY: Deny

Throw RST error for TCP and ICMP error for UDP &#8203;- DROP: DROP

Silently drop traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ALLOW",
    "DENY",
    "DROP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ALLOW",
  "enum": [
    "ALLOW",
    "DENY",
    "DROP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-bf8c9aeaab4739ed146a0065f148aa576ba7a07354ff7da2a0d7cb129557e646"></a>

<a id="canonical-3df8089f5afdbff2ee22a0c5fec9bc332ecb67258d07dafd800e284fd3f19413"></a>

## annotations property — Property reference / 98dcc9bbb9bd / 5

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

<a id="canonical-a40a837afcc4a241e8cc520736347a56b00bcf23ad5fdc11988c2b8cbc9811f4"></a>

<a id="canonical-bd5a30025d4dbc419023e478f799e8c4c80fa625df9b0c7dffb5958583883822"></a>

## description property — Property reference / 98dcc9bbb9bd / 6

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

<a id="canonical-e31e53bb492473cafc90da2afb1294b8741076eff32ef8a169f07c64f33cc03a"></a>

<a id="canonical-8e78d33af202797868aca0e7569ae9dba88d30f3c3c5ac57839b6c2802e4099c"></a>

## disable property — Property reference / 98dcc9bbb9bd / 7

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

- [enable_disable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-cecd863a89400095b4a3e12d6f4540655ced3c84a79fd7d64de648f4427fd7e2): complete subsection reference.

- [enable_disable_signatures](resources--protocol_inspection--reference--group-001.md#canonical-53898f41c9aebb3ff5557e3017bfcec174a3dd1cd5b1cce9cf672bf2286adaaf): complete subsection reference.

<a id="canonical-1af65ffd43fa1e3096decafd8dc25920e9e7fcc38ca16d8587f307b580eccd2e"></a>

<a id="canonical-8d1b503a26ceeb438be7cbbf8f01c27091fe55a1a10be8243091d64b23adf16f"></a>

## id property — Property reference / 98dcc9bbb9bd / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-ae12424e644941ea859a4ec2e94069c8d262b8066118f03def7c0e5e85ce272a"></a>

<a id="canonical-34e4e91d8d9f28cf68717d178b73a5d7e0f97b784b46a54ff00927b0cd238957"></a>

## labels property — Property reference / 98dcc9bbb9bd / 9

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

<a id="canonical-a8af682c465e4dd66fe89873e69e663a74ee1f915971fd5b1628721950c9c0ff"></a>

<a id="canonical-88d48faabb100f5dbfa47a8061e63846bc48bac5d1e452821a7e0432dedc4fed"></a>

## name property — Property reference / 98dcc9bbb9bd / 10

Type: `"string"`. Required.

Name of the Protocol Inspection. Must be unique within the namespace.

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

<a id="canonical-1b9b5d0ab652fe8494f7c02a310a9c76cc8c29ba5be39786c0e1af04083d019b"></a>

<a id="canonical-fce89a67e25f1dde35307585837d1300aca2608300a4685863f2ba2765ed02f2"></a>

## namespace property — Property reference / 98dcc9bbb9bd / 11

Type: `"string"`. Required.

Namespace where the Protocol Inspection is created.

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

- [timeouts](resources--protocol_inspection--reference--group-001.md#canonical-a4036f65653afbcd837ce39f90c89b339377eac12c4dfa014da8f20f2ed1ca16): complete subsection reference.

<a id="canonical-382ab818df8e4f6e258a823aa7ac4d39e620e599b2e65bf3977380edd5d4d5ed"></a>

## All schema paths — Property reference / 98dcc9bbb9bd / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](resources--protocol_inspection--reference--group-001.md#canonical-dbe152e2908953ef95d20aeae60261881dfa67b889cd47514a684249972a4f1d) |
| `annotations` | [annotations](resources--protocol_inspection--reference--group-001.md#canonical-bf8c9aeaab4739ed146a0065f148aa576ba7a07354ff7da2a0d7cb129557e646) |
| `description` | [description](resources--protocol_inspection--reference--group-001.md#canonical-a40a837afcc4a241e8cc520736347a56b00bcf23ad5fdc11988c2b8cbc9811f4) |
| `disable` | [disable](resources--protocol_inspection--reference--group-001.md#canonical-e31e53bb492473cafc90da2afb1294b8741076eff32ef8a169f07c64f33cc03a) |
| `enable_disable_compliance_checks` | [enable_disable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-93013b51bf40dc0053d824d8125439ad5449f42316411a142f561deb7fdd68ac) |
| `enable_disable_compliance_checks.disable_compliance_checks` | [enable_disable_compliance_checks.disable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-d410a52c7eef327e4a0707bbf01427cf47e3f4cb6b0c5b6edf0d595d33ffdb7e) |
| `enable_disable_compliance_checks.enable_compliance_checks` | [enable_disable_compliance_checks.enable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-c6640b4295fa59a0364cd5a4ab66cfcc8a73fdfe4bf95446c9a0908df5917364) |
| `enable_disable_compliance_checks.enable_compliance_checks.name` | [enable_disable_compliance_checks.enable_compliance_checks.name](resources--protocol_inspection--reference--group-001.md#canonical-ff2a2001192e1af459c5b3f5f8d0374cf2c37f57e2ae688450b3f2932469c764) |
| `enable_disable_compliance_checks.enable_compliance_checks.namespace` | [enable_disable_compliance_checks.enable_compliance_checks.namespace](resources--protocol_inspection--reference--group-001.md#canonical-9ed77c8008ba965afe1519356c8dcb3ec4caefb65f1a0a0e36a2038667d2f383) |
| `enable_disable_compliance_checks.enable_compliance_checks.tenant` | [enable_disable_compliance_checks.enable_compliance_checks.tenant](resources--protocol_inspection--reference--group-001.md#canonical-6b7593bececfff279fd75d5a1a51e0971f7b42c372f9ff2b2742f8fc874c541b) |
| `enable_disable_signatures` | [enable_disable_signatures](resources--protocol_inspection--reference--group-001.md#canonical-86b1c634d267c0b9b3c331c326c971fffd848e752007deeb7c90bf0a74853e6c) |
| `enable_disable_signatures.disable_signature` | [enable_disable_signatures.disable_signature](resources--protocol_inspection--reference--group-001.md#canonical-537214204eb1600708d8ee182e10f8c5ea74ffeb7d7396a06a63d7fd186e1b45) |
| `enable_disable_signatures.enable_signature` | [enable_disable_signatures.enable_signature](resources--protocol_inspection--reference--group-001.md#canonical-1ad2ef3a28fdceec8c53d6c127ca357254431d16b6275d041e44866eac81101b) |
| `id` | [id](resources--protocol_inspection--reference--group-001.md#canonical-1af65ffd43fa1e3096decafd8dc25920e9e7fcc38ca16d8587f307b580eccd2e) |
| `labels` | [labels](resources--protocol_inspection--reference--group-001.md#canonical-ae12424e644941ea859a4ec2e94069c8d262b8066118f03def7c0e5e85ce272a) |
| `name` | [name](resources--protocol_inspection--reference--group-001.md#canonical-a8af682c465e4dd66fe89873e69e663a74ee1f915971fd5b1628721950c9c0ff) |
| `namespace` | [namespace](resources--protocol_inspection--reference--group-001.md#canonical-1b9b5d0ab652fe8494f7c02a310a9c76cc8c29ba5be39786c0e1af04083d019b) |
| `timeouts` | [timeouts](resources--protocol_inspection--reference--group-001.md#canonical-8797497dd00379dea37b0259508a9cff83e53f2473e5d54dba04752746ec0504) |
| `timeouts.create` | [timeouts.create](resources--protocol_inspection--reference--group-001.md#canonical-eada380d86ec84007863be4303783bc7874505923d5217efc80c2faadf9f3a8a) |
| `timeouts.delete` | [timeouts.delete](resources--protocol_inspection--reference--group-001.md#canonical-35321ec37f346bca37df006bb05cbf926f9000d661ae991f2e8e56baa765ebf2) |
| `timeouts.read` | [timeouts.read](resources--protocol_inspection--reference--group-001.md#canonical-4dddcbebbc17c0377fac813a419b5391083d48a5daac233124ffbf8268ce7a9a) |
| `timeouts.update` | [timeouts.update](resources--protocol_inspection--reference--group-001.md#canonical-5b57f0cf6a4b80d4b639ae897cdd2503e66989d57d45b67fca9323847286384f) |

<a id="canonical-5a714a0ddef75d07632909850eadf08b2d838ef7dc8eca4fc82ca1c9614db61b"></a>

## Next pages — Property reference / 98dcc9bbb9bd / 13

- [enable_disable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-cecd863a89400095b4a3e12d6f4540655ced3c84a79fd7d64de648f4427fd7e2)
- [enable_disable_signatures](resources--protocol_inspection--reference--group-001.md#canonical-53898f41c9aebb3ff5557e3017bfcec174a3dd1cd5b1cce9cf672bf2286adaaf)
- [timeouts](resources--protocol_inspection--reference--group-001.md#canonical-a4036f65653afbcd837ce39f90c89b339377eac12c4dfa014da8f20f2ed1ca16)
- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)

<a id="canonical-cecd863a89400095b4a3e12d6f4540655ced3c84a79fd7d64de648f4427fd7e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-464cce81568edad5d73aa544d7273a3ebdd57e25cafdf4c4a719ea6035b63f0b"></a>

## enable_disable_compliance_checks — enable_disable_compliance_checks / f01a0c69ca43 / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)
- [Property reference](resources--protocol_inspection--reference--group-001.md#canonical-e4364e9cdf919f8070e0c9ff68ce35d1ccf1e8b653ca9312621219bdcca91de8)
- enable_disable_compliance_checks

<a id="canonical-93013b51bf40dc0053d824d8125439ad5449f42316411a142f561deb7fdd68ac"></a>

Type: `"object"`. single nested block, Optional.

Enable Disable Compliance Checks Choice.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_compliance_checks",
    "enable_compliance_checks")}
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
  "x-ves-oneof-field-compliance_check_choice": "[\"disable_compliance_checks\",\"enable_compliance_checks\"]"
}
```

Terraform syntax:

```terraform
enable_disable_compliance_checks {
  # Configure direct properties listed below.
}
```

<a id="canonical-d22c3b5c15a795f1ff829183190b917cf613abe35712ace84198f1afec059c36"></a>

## Direct properties — enable_disable_compliance_checks / f01a0c69ca43 / 3

- [disable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-829e4ab7718cef2e5c09121d4da28822250572639eff8b8d708dd140584f92c0): complete subsection reference.

- [enable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-451b401d276dcd7f8ab6984e5ff27eed33a272e1b8ef25334088e87182533b75): complete subsection reference.

<a id="canonical-06ff0060a1404ee7c9d3bb876a4a622d208007c1e74a39dcd02da825afc1ae7d"></a>

## Next pages — enable_disable_compliance_checks / f01a0c69ca43 / 4

- [enable_disable_compliance_checks.disable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-829e4ab7718cef2e5c09121d4da28822250572639eff8b8d708dd140584f92c0)
- [enable_disable_compliance_checks.enable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-451b401d276dcd7f8ab6984e5ff27eed33a272e1b8ef25334088e87182533b75)
- [Property reference](resources--protocol_inspection--reference--group-001.md#canonical-e4364e9cdf919f8070e0c9ff68ce35d1ccf1e8b653ca9312621219bdcca91de8)
- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)

<a id="canonical-829e4ab7718cef2e5c09121d4da28822250572639eff8b8d708dd140584f92c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f5d74fbaa914dc117afe13626059ce8a02d7e31f7765af433b4cec79a005780"></a>

## enable_disable_compliance_checks.disable_compliance_checks — enable_disable_compliance_checks.disable_compliance_checks / ed4373217f92 / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)
- [Property reference](resources--protocol_inspection--reference--group-001.md#canonical-e4364e9cdf919f8070e0c9ff68ce35d1ccf1e8b653ca9312621219bdcca91de8)
- [enable_disable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-cecd863a89400095b4a3e12d6f4540655ced3c84a79fd7d64de648f4427fd7e2)
- enable_disable_compliance_checks.disable_compliance_checks

<a id="canonical-d410a52c7eef327e4a0707bbf01427cf47e3f4cb6b0c5b6edf0d595d33ffdb7e"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable compliance checks.

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
disable_compliance_checks = {}
```

<a id="canonical-6c7a2f3c53bf78f5075be0075352f515754916b7e163a44b0a7666ffeffdcb1d"></a>

## Direct properties — enable_disable_compliance_checks.disable_compliance_checks / ed4373217f92 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ca0924a0941107ee7009a43b400bb7f9980fc592f56b6006e67c0dac4d25da77"></a>

## Next pages — enable_disable_compliance_checks.disable_compliance_checks / ed4373217f92 / 4

- [enable_disable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-cecd863a89400095b4a3e12d6f4540655ced3c84a79fd7d64de648f4427fd7e2)
- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)

<a id="canonical-451b401d276dcd7f8ab6984e5ff27eed33a272e1b8ef25334088e87182533b75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-129cfe7c1afdc11e9613bf4856e3ea96ad9ef06cb081d3e350ca9bdf05ee3a1f"></a>

## enable_disable_compliance_checks.enable_compliance_checks — enable_disable_compliance_checks.enable_compliance_checks / d3dd34759313 / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)
- [Property reference](resources--protocol_inspection--reference--group-001.md#canonical-e4364e9cdf919f8070e0c9ff68ce35d1ccf1e8b653ca9312621219bdcca91de8)
- [enable_disable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-cecd863a89400095b4a3e12d6f4540655ced3c84a79fd7d64de648f4427fd7e2)
- enable_disable_compliance_checks.enable_compliance_checks

<a id="canonical-c6640b4295fa59a0364cd5a4ab66cfcc8a73fdfe4bf95446c9a0908df5917364"></a>

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
enable_compliance_checks {
  # Configure direct properties listed below.
}
```

<a id="canonical-d5b9f25f881c047141afdcd5762f84c90e52ae890baf15a59cc0cdd0543b8e2b"></a>

## Direct properties — enable_disable_compliance_checks.enable_compliance_checks / d3dd34759313 / 3

<a id="canonical-ff2a2001192e1af459c5b3f5f8d0374cf2c37f57e2ae688450b3f2932469c764"></a>

<a id="canonical-45b690402634934680b2b3cef953e1151ed382e5b2d01d266dd9337c5f5b8024"></a>

## name property — enable_disable_compliance_checks.enable_compliance_checks / d3dd34759313 / 4

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

<a id="canonical-9ed77c8008ba965afe1519356c8dcb3ec4caefb65f1a0a0e36a2038667d2f383"></a>

<a id="canonical-34c663949b4fa023927db07729c06bfa5119d6ee3f3c61dca83237ae23450a6e"></a>

## namespace property — enable_disable_compliance_checks.enable_compliance_checks / d3dd34759313 / 5

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

<a id="canonical-6b7593bececfff279fd75d5a1a51e0971f7b42c372f9ff2b2742f8fc874c541b"></a>

<a id="canonical-58095db11b192a11eea3439798f689905904b8087a1711bb435b3c5049db3bca"></a>

## tenant property — enable_disable_compliance_checks.enable_compliance_checks / d3dd34759313 / 6

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

<a id="canonical-627f7cb186c84a86492b95a002eeaa7231bbed7d871bb5f5f41a5d36bc041917"></a>

## Next pages — enable_disable_compliance_checks.enable_compliance_checks / d3dd34759313 / 7

- [enable_disable_compliance_checks](resources--protocol_inspection--reference--group-001.md#canonical-cecd863a89400095b4a3e12d6f4540655ced3c84a79fd7d64de648f4427fd7e2)
- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)

<a id="canonical-53898f41c9aebb3ff5557e3017bfcec174a3dd1cd5b1cce9cf672bf2286adaaf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fec2933a3c2cfbf4b68edae58774bc16a8b1eb5ef4a5b120f07d6efaf9f7414"></a>

## enable_disable_signatures — enable_disable_signatures / 4a882cc53bad / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)
- [Property reference](resources--protocol_inspection--reference--group-001.md#canonical-e4364e9cdf919f8070e0c9ff68ce35d1ccf1e8b653ca9312621219bdcca91de8)
- enable_disable_signatures

<a id="canonical-86b1c634d267c0b9b3c331c326c971fffd848e752007deeb7c90bf0a74853e6c"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for enable disable signatures.

Upstream description:

Enable Disable Signature Choice.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_signature",
    "enable_signature")}
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
  "x-ves-oneof-field-signature_choice": "[\"disable_signature\",\"enable_signature\"]"
}
```

Terraform syntax:

```terraform
enable_disable_signatures {
  # Configure direct properties listed below.
}
```

<a id="canonical-1ff21d3e89a0f039480cff5770d713b4b977a0a9662221497fcc81a83c041ea3"></a>

## Direct properties — enable_disable_signatures / 4a882cc53bad / 3

- [disable_signature](resources--protocol_inspection--reference--group-001.md#canonical-fc7e4a6ceabe60f2465919b2a3bd4a0825980059c20f33975a1266d717c69be6): complete subsection reference.

- [enable_signature](resources--protocol_inspection--reference--group-001.md#canonical-697711c68815f5cbe4fd83ade83e1cc3e0ef4b937e5d52a1740f422ee8c065f3): complete subsection reference.

<a id="canonical-a642bb06b2f5e9fe9f636c15d5eec389885d41781d55e3a3c53c94ae65aca52a"></a>

## Next pages — enable_disable_signatures / 4a882cc53bad / 4

- [enable_disable_signatures.disable_signature](resources--protocol_inspection--reference--group-001.md#canonical-fc7e4a6ceabe60f2465919b2a3bd4a0825980059c20f33975a1266d717c69be6)
- [enable_disable_signatures.enable_signature](resources--protocol_inspection--reference--group-001.md#canonical-697711c68815f5cbe4fd83ade83e1cc3e0ef4b937e5d52a1740f422ee8c065f3)
- [Property reference](resources--protocol_inspection--reference--group-001.md#canonical-e4364e9cdf919f8070e0c9ff68ce35d1ccf1e8b653ca9312621219bdcca91de8)
- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)

<a id="canonical-fc7e4a6ceabe60f2465919b2a3bd4a0825980059c20f33975a1266d717c69be6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90d6fe327cac49806f11e2e40094ec9ee1c1a0841974d864bd05ed8934ae3376"></a>

## enable_disable_signatures.disable_signature — enable_disable_signatures.disable_signature / bae5dcb64abb / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)
- [Property reference](resources--protocol_inspection--reference--group-001.md#canonical-e4364e9cdf919f8070e0c9ff68ce35d1ccf1e8b653ca9312621219bdcca91de8)
- [enable_disable_signatures](resources--protocol_inspection--reference--group-001.md#canonical-53898f41c9aebb3ff5557e3017bfcec174a3dd1cd5b1cce9cf672bf2286adaaf)
- enable_disable_signatures.disable_signature

<a id="canonical-537214204eb1600708d8ee182e10f8c5ea74ffeb7d7396a06a63d7fd186e1b45"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable signature.

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
disable_signature = {}
```

<a id="canonical-640a6bcb5c7bcf911f7ad3ac55885c580482c2054e936fb164860e3dd759d6f3"></a>

## Direct properties — enable_disable_signatures.disable_signature / bae5dcb64abb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e66dccc7db8a3a7af02694d7f8b7c27b638d27e3a44e08d546e23981f0a702b2"></a>

## Next pages — enable_disable_signatures.disable_signature / bae5dcb64abb / 4

- [enable_disable_signatures](resources--protocol_inspection--reference--group-001.md#canonical-53898f41c9aebb3ff5557e3017bfcec174a3dd1cd5b1cce9cf672bf2286adaaf)
- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)

<a id="canonical-697711c68815f5cbe4fd83ade83e1cc3e0ef4b937e5d52a1740f422ee8c065f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2afd6466ebf8bd49b380da1f46c020bad336400d6b661434e39be98a8025fe1"></a>

## enable_disable_signatures.enable_signature — enable_disable_signatures.enable_signature / 75dd9ddabb25 / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)
- [Property reference](resources--protocol_inspection--reference--group-001.md#canonical-e4364e9cdf919f8070e0c9ff68ce35d1ccf1e8b653ca9312621219bdcca91de8)
- [enable_disable_signatures](resources--protocol_inspection--reference--group-001.md#canonical-53898f41c9aebb3ff5557e3017bfcec174a3dd1cd5b1cce9cf672bf2286adaaf)
- enable_disable_signatures.enable_signature

<a id="canonical-1ad2ef3a28fdceec8c53d6c127ca357254431d16b6275d041e44866eac81101b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable signature.

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
enable_signature = {}
```

<a id="canonical-63b978e61d0d7437c3754970c1f0e5818fcff251b8ec52c3f0b608595287c051"></a>

## Direct properties — enable_disable_signatures.enable_signature / 75dd9ddabb25 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a454974b47025ab6763ac601754c33e9a246bbbe9d5e1470aec33d4e6dd25095"></a>

## Next pages — enable_disable_signatures.enable_signature / 75dd9ddabb25 / 4

- [enable_disable_signatures](resources--protocol_inspection--reference--group-001.md#canonical-53898f41c9aebb3ff5557e3017bfcec174a3dd1cd5b1cce9cf672bf2286adaaf)
- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)

<a id="canonical-a4036f65653afbcd837ce39f90c89b339377eac12c4dfa014da8f20f2ed1ca16"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-725c99225656353d8d8e880f47f15d73c976a19202114b57cdec04abe9b87328"></a>

## timeouts — timeouts / de54c86a1912 / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)
- [Property reference](resources--protocol_inspection--reference--group-001.md#canonical-e4364e9cdf919f8070e0c9ff68ce35d1ccf1e8b653ca9312621219bdcca91de8)
- timeouts

<a id="canonical-8797497dd00379dea37b0259508a9cff83e53f2473e5d54dba04752746ec0504"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0262fd36f787e234e19b51ff48e0f81b41c8ee3bef9b7d5a5d408ab2d67f651b"></a>

## Direct properties — timeouts / de54c86a1912 / 3

<a id="canonical-eada380d86ec84007863be4303783bc7874505923d5217efc80c2faadf9f3a8a"></a>

<a id="canonical-0fa2f10f5555afb3cbd2f4eb972755d7ed9f4aec9a5164c9e0fe601d9e6368ff"></a>

## create property — timeouts / de54c86a1912 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-35321ec37f346bca37df006bb05cbf926f9000d661ae991f2e8e56baa765ebf2"></a>

<a id="canonical-30af40ff50d0071b01627d5a12c7db94bc8f1ccb416a8a4fc8bbc184582e93db"></a>

## delete property — timeouts / de54c86a1912 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-4dddcbebbc17c0377fac813a419b5391083d48a5daac233124ffbf8268ce7a9a"></a>

<a id="canonical-b78a359e7333adb373486d8e8838cdc6af6854ae74a31caaf585011c831b1314"></a>

## read property — timeouts / de54c86a1912 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-5b57f0cf6a4b80d4b639ae897cdd2503e66989d57d45b67fca9323847286384f"></a>

<a id="canonical-43cdfda926a734c3004868bd77ff506a152e5dd7f895dc3b42e8676d7bf658e1"></a>

## update property — timeouts / de54c86a1912 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-d6c9708652eef7f6d61d4e66ffdebfc7cd7f705c141163e421a573ca944d7165"></a>

## Next pages — timeouts / de54c86a1912 / 8

- [Property reference](resources--protocol_inspection--reference--group-001.md#canonical-e4364e9cdf919f8070e0c9ff68ce35d1ccf1e8b653ca9312621219bdcca91de8)
- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)
