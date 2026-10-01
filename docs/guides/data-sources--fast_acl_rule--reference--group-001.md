---
page_title: "xcsh_fast_acl_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl_rule reference."
---

# xcsh_fast_acl_rule reference

<a id="canonical-1c3c85be785c735afb4ad85cb212fa043eaa753afc4f7076378bdd310dc0b8dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-968a724f4fbe7a295fcf8103abcadff4ed704a34d09892e75e83eb8b53259d99"></a>

## Property reference — Property reference / db0fc039d672 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)
- Property reference

<a id="canonical-0991e3f9513fdcf92d436ed23a8e910f1867a2fa5162ef685e542e0d957743ad"></a>

## Direct properties — Property reference / db0fc039d672 / 3

- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-bbe934173ca437b60ab99b6a81590d4dcf4beb08a0a474c9118dc0883f0355ef): complete subsection reference.

<a id="canonical-edd76168f95fd65e3ba99efa1606a73abdef991558df806bd63860e72fc2460e"></a>

<a id="canonical-064248582afeb90f0f7ac18b3c44de436dcafe474482faec1d0500ae7e881e56"></a>

## annotations property — Property reference / db0fc039d672 / 4

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

<a id="canonical-06e42872ba79a9d8518099dd9b0b87d5d7a0434f89774bf13601317849358d5c"></a>

<a id="canonical-29bed206df2c425c726a903051219e9ec34bdfd7cba889f6ea36978feefce32c"></a>

## description property — Property reference / db0fc039d672 / 5

Type: `"string"`. Computed.

Description of the FastACLRule.

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

<a id="canonical-42292d640da80a8bbe62a2561a527babfe83047c1425e27ec7a754ba0f76afdb"></a>

<a id="canonical-c9ebc90e3fbbcb3d8559143bab818d2dba7d0d00043b9650d0a9d18dbd46018e"></a>

## id property — Property reference / db0fc039d672 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip_prefix_set](data-sources--fast_acl_rule--reference--group-001.md#canonical-514e44fd9fabd7a578d488e007a2f893f58fc57dd209c84937cb81cbaf85f5a3): complete subsection reference.

<a id="canonical-44cc454f1e0c517351a6fba4d6afeac18ce0d2c1a6333df40d34395d834d8a7f"></a>

<a id="canonical-832117bbbd34b1673d1bc4fc18fa212dbf23089f4913708b2c41ac83c6aaf687"></a>

## labels property — Property reference / db0fc039d672 / 7

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

<a id="canonical-9bebdce04c0fc2c84ada4652f6f4f600c973dbcc2370cc69dd4d9d24c09b62ee"></a>

<a id="canonical-d183ed74975b55b626926862933650811d027b3a3cbdc3a9ee222289e388d938"></a>

## name property — Property reference / db0fc039d672 / 8

Type: `"string"`. Required.

Name of the FastACLRule.

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

<a id="canonical-f6ff657a937ceafbee8a2a0edc7d893417ab442085161d75f9ec6766d160762f"></a>

<a id="canonical-16a5cb0bbd465d1964cb5e0a614dae222a216400251855d96a358a6cc7523180"></a>

## namespace property — Property reference / db0fc039d672 / 9

Type: `"string"`. Required.

Namespace where the FastACLRule exists.

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

- [port](data-sources--fast_acl_rule--reference--group-001.md#canonical-7a606ffc81ed836667c59b2844c5e00508c569bc1c3df639ea25f9d4c7236367): complete subsection reference.

- [prefix](data-sources--fast_acl_rule--reference--group-001.md#canonical-823f0fc06b588d7a6404a2291bfcacfc0c1f0d30e424bbc749870725a3a74c13): complete subsection reference.

<a id="canonical-6f3aa6356ca1b93afcdbfadd9641dadaabcdb89c59e0e8de790a0addd93a0d3d"></a>

## All schema paths — Property reference / db0fc039d672 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-de406c0d954bfa4a8ae9556dcd32e9bd56bde82f9afd102c2e146d5e0cc14f8c) |
| `action.policer_action` | [action.policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-a48e851af80508b27b0fc0cc1e120d4710e323ba5f13e265d241c72050a504ef) |
| `action.policer_action.ref` | [action.policer_action.ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-462a53faa375134e3e9e80728eaa508c0f7f2d6ffb42718d81e0a4e9f553556e) |
| `action.policer_action.ref.kind` | [action.policer_action.ref.kind](data-sources--fast_acl_rule--reference--group-001.md#canonical-6fe88e2c4ded8981aae1eda083cfa11f1858fe9e2610a94fd781e5981fc7c852) |
| `action.policer_action.ref.name` | [action.policer_action.ref.name](data-sources--fast_acl_rule--reference--group-001.md#canonical-ee669335cb424d41224c6e55b314988c6dde7d7ade3539caf4abefa34cbc405e) |
| `action.policer_action.ref.namespace` | [action.policer_action.ref.namespace](data-sources--fast_acl_rule--reference--group-001.md#canonical-1e3af92d99476e5bfa6d9386610f847479f6d547451bda8088ecf59e93e74854) |
| `action.policer_action.ref.tenant` | [action.policer_action.ref.tenant](data-sources--fast_acl_rule--reference--group-001.md#canonical-cf1ac7140d026c688b978a1cca1d8083999844ab5068f2796decf4a583fcb514) |
| `action.policer_action.ref.uid` | [action.policer_action.ref.uid](data-sources--fast_acl_rule--reference--group-001.md#canonical-ee40dc04a1217916469a9d79f3178c73f6876c5f89c8ec3e2755a396451c98a1) |
| `action.protocol_policer_action` | [action.protocol_policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-da80c26d9f2b83f3fe91bdef2603f4da0cbdaacd56fd63f404962ea265b49e4c) |
| `action.protocol_policer_action.ref` | [action.protocol_policer_action.ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-d98c82889b625d6f0c34d3b604261a8c348363fdf30aea17e6382fa4392cf527) |
| `action.protocol_policer_action.ref.kind` | [action.protocol_policer_action.ref.kind](data-sources--fast_acl_rule--reference--group-001.md#canonical-38e06aa0b2c17ca779bc74629c4eb8e78cf012a3277350bad37f5aa62a825ba8) |
| `action.protocol_policer_action.ref.name` | [action.protocol_policer_action.ref.name](data-sources--fast_acl_rule--reference--group-001.md#canonical-e173d6f8889f7a5515f8aa2def401ba21252512610889d894620e8386d48ee3f) |
| `action.protocol_policer_action.ref.namespace` | [action.protocol_policer_action.ref.namespace](data-sources--fast_acl_rule--reference--group-001.md#canonical-cca999982a852c63fb544544e42eeb6d85bb84471ffead285944ae02295d017d) |
| `action.protocol_policer_action.ref.tenant` | [action.protocol_policer_action.ref.tenant](data-sources--fast_acl_rule--reference--group-001.md#canonical-7ba4c06cd46a347da5c2ae1f3b77957325be5e2f3e1a3f8193e9bd84fbbc4cb4) |
| `action.protocol_policer_action.ref.uid` | [action.protocol_policer_action.ref.uid](data-sources--fast_acl_rule--reference--group-001.md#canonical-af96a0103cec83d794acd5de4adad57893c8941e9137ca8483583b65a35af45d) |
| `action.simple_action` | [action.simple_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-a5bf154f536facf61a8a9256fa575baa7d3669717775e213a9617a3cea5cdc95) |
| `annotations` | [annotations](data-sources--fast_acl_rule--reference--group-001.md#canonical-edd76168f95fd65e3ba99efa1606a73abdef991558df806bd63860e72fc2460e) |
| `description` | [description](data-sources--fast_acl_rule--reference--group-001.md#canonical-06e42872ba79a9d8518099dd9b0b87d5d7a0434f89774bf13601317849358d5c) |
| `id` | [id](data-sources--fast_acl_rule--reference--group-001.md#canonical-42292d640da80a8bbe62a2561a527babfe83047c1425e27ec7a754ba0f76afdb) |
| `ip_prefix_set` | [ip_prefix_set](data-sources--fast_acl_rule--reference--group-001.md#canonical-96f09ba19cfb5401d4a9252675dc955b2fc62773866c5511753cc56d1731ad9e) |
| `ip_prefix_set.ref` | [ip_prefix_set.ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-851ff55b916f5e23e5277447aef5fa3e90fc421f7fddf7250bc111bd165c27b9) |
| `ip_prefix_set.ref.kind` | [ip_prefix_set.ref.kind](data-sources--fast_acl_rule--reference--group-001.md#canonical-94faa8ec691af77198242a255081f0e0bacefb6e8c0ae6f3d3426019548cbdf6) |
| `ip_prefix_set.ref.name` | [ip_prefix_set.ref.name](data-sources--fast_acl_rule--reference--group-001.md#canonical-04aa030f8fcc87c1ea645e83b0f3d59f98f745927f565891fbd051bd1887a043) |
| `ip_prefix_set.ref.namespace` | [ip_prefix_set.ref.namespace](data-sources--fast_acl_rule--reference--group-001.md#canonical-a6b1140e6c5fc5a40ee636edfbd320b8106a755c05c36a4c19ad7afa3ddc492b) |
| `ip_prefix_set.ref.tenant` | [ip_prefix_set.ref.tenant](data-sources--fast_acl_rule--reference--group-001.md#canonical-9387380f295e4749ef7530ce467a8218ee05379fc409f001182f4a5c27dfe8bb) |
| `ip_prefix_set.ref.uid` | [ip_prefix_set.ref.uid](data-sources--fast_acl_rule--reference--group-001.md#canonical-d1f04e3bac47921fe3797726dc88760e2a766c21eafcee8f90d028897c2a9e48) |
| `labels` | [labels](data-sources--fast_acl_rule--reference--group-001.md#canonical-44cc454f1e0c517351a6fba4d6afeac18ce0d2c1a6333df40d34395d834d8a7f) |
| `name` | [name](data-sources--fast_acl_rule--reference--group-001.md#canonical-9bebdce04c0fc2c84ada4652f6f4f600c973dbcc2370cc69dd4d9d24c09b62ee) |
| `namespace` | [namespace](data-sources--fast_acl_rule--reference--group-001.md#canonical-f6ff657a937ceafbee8a2a0edc7d893417ab442085161d75f9ec6766d160762f) |
| `port` | [port](data-sources--fast_acl_rule--reference--group-001.md#canonical-de8224b91fb6f558ee15c3b0f80d71b470da11af4cfe0339a86c68f10e3ad929) |
| `port.all` | [port.all](data-sources--fast_acl_rule--reference--group-001.md#canonical-64e256ef70a2694fac0f728fbe14f9fd04eecd48e221af942a33ed1a23e501a1) |
| `port.dns` | [port.dns](data-sources--fast_acl_rule--reference--group-001.md#canonical-66c91f8f40048ca4ba27ba89c126499841bddd489e8f38e6c8948642c2c834fb) |
| `port.user_defined` | [port.user_defined](data-sources--fast_acl_rule--reference--group-001.md#canonical-28029c896e4b383144ebdc300ded7670d339f07142cb5ebbdbd848098a512e9c) |
| `prefix` | [prefix](data-sources--fast_acl_rule--reference--group-001.md#canonical-4cb76dfb12535874c747668fa8d2ea83a503d1891019e3ad4b8a3c4074650924) |
| `prefix.prefix` | [prefix.prefix](data-sources--fast_acl_rule--reference--group-001.md#canonical-0d8d3ed7477e2c6b902b2861b47ca89e1be78c3ba107c1d752bed165af516f35) |

<a id="canonical-ea6b70d22eafea3683c45f715ea86191adec71b8208bb0783a4cb80b06c05120"></a>

## Next pages — Property reference / db0fc039d672 / 11

- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-bbe934173ca437b60ab99b6a81590d4dcf4beb08a0a474c9118dc0883f0355ef)
- [ip_prefix_set](data-sources--fast_acl_rule--reference--group-001.md#canonical-514e44fd9fabd7a578d488e007a2f893f58fc57dd209c84937cb81cbaf85f5a3)
- [port](data-sources--fast_acl_rule--reference--group-001.md#canonical-7a606ffc81ed836667c59b2844c5e00508c569bc1c3df639ea25f9d4c7236367)
- [prefix](data-sources--fast_acl_rule--reference--group-001.md#canonical-823f0fc06b588d7a6404a2291bfcacfc0c1f0d30e424bbc749870725a3a74c13)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)

<a id="canonical-bbe934173ca437b60ab99b6a81590d4dcf4beb08a0a474c9118dc0883f0355ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2361eb5ae6ac389f56ff2b1511f687bebf88fd44dc8e93a77b13122323f895a"></a>

## action — action / f43ff7f1158f / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-1c3c85be785c735afb4ad85cb212fa043eaa753afc4f7076378bdd310dc0b8dc)
- action

<a id="canonical-de406c0d954bfa4a8ae9556dcd32e9bd56bde82f9afd102c2e146d5e0cc14f8c"></a>

Type: `"single"`. Computed.

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Upstream description:

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action": "[\"policer_action\",\"protocol_policer_action\",\"simple_action\"]"
}
```

<a id="canonical-1d3ff0e385ca0a25f1e4d7d45a549f6951c23fc0a52b7fb4760e054cc403df47"></a>

## Direct properties — action / f43ff7f1158f / 3

- [policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-89d11d024f58e78220149a02e51b5c426f434a8344eda69b6534083798bf8870): complete subsection reference.

- [protocol_policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-7d789db381235e55d7a0bd2270ce1a0374c5719fb4fa03738a2bad67baec4fbe): complete subsection reference.

<a id="canonical-a5bf154f536facf61a8a9256fa575baa7d3669717775e213a9617a3cea5cdc95"></a>

<a id="canonical-f72f2d3aac1bded029d3ed7309d76e10f5dd33c32cb60d5bfc87dc66d37bed5d"></a>

## simple_action property — action / f43ff7f1158f / 4

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the
traffic Forward the traffic. Possible values are \`DENY\`, \`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

FastAclRuleSimpleAction specifies simple action like PASS or DENY

Drop the traffic Forward the traffic.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-64bd9d7a50c7ddfa69d616c757faae90caee6e0f23ae18bf695ca273ce8e0448"></a>

## Next pages — action / f43ff7f1158f / 5

- [action.policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-89d11d024f58e78220149a02e51b5c426f434a8344eda69b6534083798bf8870)
- [action.protocol_policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-7d789db381235e55d7a0bd2270ce1a0374c5719fb4fa03738a2bad67baec4fbe)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-1c3c85be785c735afb4ad85cb212fa043eaa753afc4f7076378bdd310dc0b8dc)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)

<a id="canonical-89d11d024f58e78220149a02e51b5c426f434a8344eda69b6534083798bf8870"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70f28555104929cc45250f8038cffbf31543c3667878b382c06c46b573b3e19d"></a>

## action.policer_action — action.policer_action / a31364f0cb1f / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-1c3c85be785c735afb4ad85cb212fa043eaa753afc4f7076378bdd310dc0b8dc)
- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-bbe934173ca437b60ab99b6a81590d4dcf4beb08a0a474c9118dc0883f0355ef)
- action.policer_action

<a id="canonical-a48e851af80508b27b0fc0cc1e120d4710e323ba5f13e265d241c72050a504ef"></a>

Type: `"single"`. Computed.

Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

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

<a id="canonical-3531747e792f764afd1c36e83b9899cb7c0119a2e2647133b38cfa4c300e950f"></a>

## Direct properties — action.policer_action / a31364f0cb1f / 3

- [ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-91970042138a600b6302b0d7e42a106d06ab1793d8082eb5bc5604982638bb22): complete subsection reference.

<a id="canonical-cd9ac78978eb6fbc507eb138490006319a1382485cb2063c22fbf5f81e1ae536"></a>

## Next pages — action.policer_action / a31364f0cb1f / 4

- [action.policer_action.ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-91970042138a600b6302b0d7e42a106d06ab1793d8082eb5bc5604982638bb22)
- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-bbe934173ca437b60ab99b6a81590d4dcf4beb08a0a474c9118dc0883f0355ef)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)

<a id="canonical-91970042138a600b6302b0d7e42a106d06ab1793d8082eb5bc5604982638bb22"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-326ef173ea77cdc34aa572374abe8269abcbf26cb8b391741e1e52375b6d772c"></a>

## action.policer_action.ref — action.policer_action.ref / 39822e3fd87a / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-1c3c85be785c735afb4ad85cb212fa043eaa753afc4f7076378bdd310dc0b8dc)
- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-bbe934173ca437b60ab99b6a81590d4dcf4beb08a0a474c9118dc0883f0355ef)
- [action.policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-89d11d024f58e78220149a02e51b5c426f434a8344eda69b6534083798bf8870)
- action.policer_action.ref

<a id="canonical-462a53faa375134e3e9e80728eaa508c0f7f2d6ffb42718d81e0a4e9f553556e"></a>

Type: `"list"`. Computed.

Reference. A policer direct reference.

Upstream description:

A policer direct reference.

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

<a id="canonical-7d363823e8fe14e1427f5d75796e470863ce1850c02d819daaee3a1dd0ce587d"></a>

## Direct properties — action.policer_action.ref / 39822e3fd87a / 3

<a id="canonical-6fe88e2c4ded8981aae1eda083cfa11f1858fe9e2610a94fd781e5981fc7c852"></a>

<a id="canonical-0174bf9f6423b327a403f979b6552b905ecce8086c8bca1a9a16f7ca828704bf"></a>

## kind property — action.policer_action.ref / 39822e3fd87a / 4

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

<a id="canonical-ee669335cb424d41224c6e55b314988c6dde7d7ade3539caf4abefa34cbc405e"></a>

<a id="canonical-3aa0393ee88ae76e7366f8a36ab36b62d44b5196b86b8e367c9973a1e1dd4be9"></a>

## name property — action.policer_action.ref / 39822e3fd87a / 5

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

<a id="canonical-1e3af92d99476e5bfa6d9386610f847479f6d547451bda8088ecf59e93e74854"></a>

<a id="canonical-065c935f8cad05c5982676e15c4bae0eab31d667de3af5954771370f430e2449"></a>

## namespace property — action.policer_action.ref / 39822e3fd87a / 6

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

<a id="canonical-cf1ac7140d026c688b978a1cca1d8083999844ab5068f2796decf4a583fcb514"></a>

<a id="canonical-98a4ba11e49b24f72eae1e044e6f5151c318b17ca6966cf3abcc53b245cc5e5b"></a>

## tenant property — action.policer_action.ref / 39822e3fd87a / 7

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

<a id="canonical-ee40dc04a1217916469a9d79f3178c73f6876c5f89c8ec3e2755a396451c98a1"></a>

<a id="canonical-52304961b47c42a07a2a5fd75e72e41595c980d8f26e0149c96ca9ca0598eaae"></a>

## uid property — action.policer_action.ref / 39822e3fd87a / 8

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

<a id="canonical-734ea35c103a06b7d479ae1bee5a51c52261a1941e2bfa2f4bc83a80e89f76fc"></a>

## Next pages — action.policer_action.ref / 39822e3fd87a / 9

- [action.policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-89d11d024f58e78220149a02e51b5c426f434a8344eda69b6534083798bf8870)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)

<a id="canonical-7d789db381235e55d7a0bd2270ce1a0374c5719fb4fa03738a2bad67baec4fbe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93e317cdb78aed996f6a5ccd38aba9a4b78d334388ff9c668abaef375c7869d3"></a>

## action.protocol_policer_action — action.protocol_policer_action / 98354238ae53 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-1c3c85be785c735afb4ad85cb212fa043eaa753afc4f7076378bdd310dc0b8dc)
- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-bbe934173ca437b60ab99b6a81590d4dcf4beb08a0a474c9118dc0883f0355ef)
- action.protocol_policer_action

<a id="canonical-da80c26d9f2b83f3fe91bdef2603f4da0cbdaacd56fd63f404962ea265b49e4c"></a>

Type: `"single"`. Computed.

Protocol Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

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

<a id="canonical-06537e501f6103196946594889cd8d53418ae76502d03bdec2c468152e3aecc6"></a>

## Direct properties — action.protocol_policer_action / 98354238ae53 / 3

- [ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-46447f1f0ea24016212a10a64614c320c085ea28506978ddf1a2b77502e50d34): complete subsection reference.

<a id="canonical-17fd12ad30e7482c1f886794d0f59fe6da9f73888e3f283c91056f98de561f49"></a>

## Next pages — action.protocol_policer_action / 98354238ae53 / 4

- [action.protocol_policer_action.ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-46447f1f0ea24016212a10a64614c320c085ea28506978ddf1a2b77502e50d34)
- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-bbe934173ca437b60ab99b6a81590d4dcf4beb08a0a474c9118dc0883f0355ef)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)

<a id="canonical-46447f1f0ea24016212a10a64614c320c085ea28506978ddf1a2b77502e50d34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f36fee790f8d7b83f805de168e271c9a92815f0d4ffbaab54972bd448407d9f"></a>

## action.protocol_policer_action.ref — action.protocol_policer_action.ref / 83b79507f6e1 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-1c3c85be785c735afb4ad85cb212fa043eaa753afc4f7076378bdd310dc0b8dc)
- [action](data-sources--fast_acl_rule--reference--group-001.md#canonical-bbe934173ca437b60ab99b6a81590d4dcf4beb08a0a474c9118dc0883f0355ef)
- [action.protocol_policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-7d789db381235e55d7a0bd2270ce1a0374c5719fb4fa03738a2bad67baec4fbe)
- action.protocol_policer_action.ref

<a id="canonical-d98c82889b625d6f0c34d3b604261a8c348363fdf30aea17e6382fa4392cf527"></a>

Type: `"list"`. Computed.

Protocol policer Reference. Reference to protocol policer object.

Upstream description:

Reference to protocol policer object.

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

<a id="canonical-d78309702cf4f1efe528475d98908b137fcb420498698bd2bc72ed0bae04df34"></a>

## Direct properties — action.protocol_policer_action.ref / 83b79507f6e1 / 3

<a id="canonical-38e06aa0b2c17ca779bc74629c4eb8e78cf012a3277350bad37f5aa62a825ba8"></a>

<a id="canonical-df523f0f59ccc7e477641d9583ca4f40fd4b8880c5576b5e322c88df0e372d4d"></a>

## kind property — action.protocol_policer_action.ref / 83b79507f6e1 / 4

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

<a id="canonical-e173d6f8889f7a5515f8aa2def401ba21252512610889d894620e8386d48ee3f"></a>

<a id="canonical-bfb8170d5424bc6fe560f64782a69b636e78cdf5e8b3d259d2328dc80d00867c"></a>

## name property — action.protocol_policer_action.ref / 83b79507f6e1 / 5

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

<a id="canonical-cca999982a852c63fb544544e42eeb6d85bb84471ffead285944ae02295d017d"></a>

<a id="canonical-8f55034fac6baa17c508ea7c1eee64b6f7c7012dcf9dde9a9fe075c0935f6cfd"></a>

## namespace property — action.protocol_policer_action.ref / 83b79507f6e1 / 6

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

<a id="canonical-7ba4c06cd46a347da5c2ae1f3b77957325be5e2f3e1a3f8193e9bd84fbbc4cb4"></a>

<a id="canonical-e015627ade5e8b737af46a7247d60a9452f480f105a2abfe52900abb0b1cbd93"></a>

## tenant property — action.protocol_policer_action.ref / 83b79507f6e1 / 7

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

<a id="canonical-af96a0103cec83d794acd5de4adad57893c8941e9137ca8483583b65a35af45d"></a>

<a id="canonical-67c75df441f04d8b3fa1f57182cdf5886e52bf1d71739b2a04fcfbd93417c58f"></a>

## uid property — action.protocol_policer_action.ref / 83b79507f6e1 / 8

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

<a id="canonical-62ec1df681b77898ea4f8d8b441a3f25b9db29b9da74d5affe7068063201d42e"></a>

## Next pages — action.protocol_policer_action.ref / 83b79507f6e1 / 9

- [action.protocol_policer_action](data-sources--fast_acl_rule--reference--group-001.md#canonical-7d789db381235e55d7a0bd2270ce1a0374c5719fb4fa03738a2bad67baec4fbe)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)

<a id="canonical-514e44fd9fabd7a578d488e007a2f893f58fc57dd209c84937cb81cbaf85f5a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-935d41b8b1654b96e4c1c067bb3a0e7f9c636595642d3ec44230d19edb3123d8"></a>

## ip_prefix_set — ip_prefix_set / 25a65e87dd90 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-1c3c85be785c735afb4ad85cb212fa043eaa753afc4f7076378bdd310dc0b8dc)
- ip_prefix_set

<a id="canonical-96f09ba19cfb5401d4a9252675dc955b2fc62773866c5511753cc56d1731ad9e"></a>

Type: `"single"`. Computed.

\[OneOf: ip\_prefix\_set, prefix\] List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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

OneOf alternatives in this subsection:

- [ip_prefix_set](data-sources--fast_acl_rule--reference--group-001.md#canonical-96f09ba19cfb5401d4a9252675dc955b2fc62773866c5511753cc56d1731ad9e)
- [prefix](data-sources--fast_acl_rule--reference--group-001.md#canonical-4cb76dfb12535874c747668fa8d2ea83a503d1891019e3ad4b8a3c4074650924)

Select alternatives according to the provider validators above.

<a id="canonical-0ed300e96d69aad40abe0935841b2a3ea8ca7f74c3875a2b1c067fffce3f47ef"></a>

## Direct properties — ip_prefix_set / 25a65e87dd90 / 3

- [ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-02b3597e03864913496a787f303fbad5438270df162355f5bc54ad54d7d25850): complete subsection reference.

<a id="canonical-3e6187890cb9f1e2bd89ff4782cc54750212b45b90d74adb751eb4a37b58fda4"></a>

## Next pages — ip_prefix_set / 25a65e87dd90 / 4

- [ip_prefix_set.ref](data-sources--fast_acl_rule--reference--group-001.md#canonical-02b3597e03864913496a787f303fbad5438270df162355f5bc54ad54d7d25850)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-1c3c85be785c735afb4ad85cb212fa043eaa753afc4f7076378bdd310dc0b8dc)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)

<a id="canonical-02b3597e03864913496a787f303fbad5438270df162355f5bc54ad54d7d25850"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18bf4cac48ab7421ecd9711911c4dd152e8e6a257a16ab967945133756cb0d9f"></a>

## ip_prefix_set.ref — ip_prefix_set.ref / 608dfc6e66f2 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-1c3c85be785c735afb4ad85cb212fa043eaa753afc4f7076378bdd310dc0b8dc)
- [ip_prefix_set](data-sources--fast_acl_rule--reference--group-001.md#canonical-514e44fd9fabd7a578d488e007a2f893f58fc57dd209c84937cb81cbaf85f5a3)
- ip_prefix_set.ref

<a id="canonical-851ff55b916f5e23e5277447aef5fa3e90fc421f7fddf7250bc111bd165c27b9"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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

<a id="canonical-637b47b661ba8d87aecf0c786f2d5e05925bf013b612e5dc3607979ce51ae4e5"></a>

## Direct properties — ip_prefix_set.ref / 608dfc6e66f2 / 3

<a id="canonical-94faa8ec691af77198242a255081f0e0bacefb6e8c0ae6f3d3426019548cbdf6"></a>

<a id="canonical-4fc2217a55ba3dc6175833106461870954111836e4b7bc64404a2180757877cf"></a>

## kind property — ip_prefix_set.ref / 608dfc6e66f2 / 4

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

<a id="canonical-04aa030f8fcc87c1ea645e83b0f3d59f98f745927f565891fbd051bd1887a043"></a>

<a id="canonical-a30c54f51c2cff3652baf463101b15ab0d522109b3fcbedf5a5b511f223f4003"></a>

## name property — ip_prefix_set.ref / 608dfc6e66f2 / 5

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

<a id="canonical-a6b1140e6c5fc5a40ee636edfbd320b8106a755c05c36a4c19ad7afa3ddc492b"></a>

<a id="canonical-aaffa93a41f86cacc46c1c4d6f5306fb0f55cbc28ff4284385a7b67b5996aa1e"></a>

## namespace property — ip_prefix_set.ref / 608dfc6e66f2 / 6

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

<a id="canonical-9387380f295e4749ef7530ce467a8218ee05379fc409f001182f4a5c27dfe8bb"></a>

<a id="canonical-b081a33f486224db2c6988a1e20ab9062f5356d2031e0b02b4b07a8327a654ff"></a>

## tenant property — ip_prefix_set.ref / 608dfc6e66f2 / 7

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

<a id="canonical-d1f04e3bac47921fe3797726dc88760e2a766c21eafcee8f90d028897c2a9e48"></a>

<a id="canonical-f237957dc43b096105523f2e3a680f329622289ba0e18e41f91c7211cd73b6a9"></a>

## uid property — ip_prefix_set.ref / 608dfc6e66f2 / 8

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

<a id="canonical-0fe4ed97afcfaa4ba37782698096526881868683be6fb65992bc693a0905d66a"></a>

## Next pages — ip_prefix_set.ref / 608dfc6e66f2 / 9

- [ip_prefix_set](data-sources--fast_acl_rule--reference--group-001.md#canonical-514e44fd9fabd7a578d488e007a2f893f58fc57dd209c84937cb81cbaf85f5a3)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)

<a id="canonical-7a606ffc81ed836667c59b2844c5e00508c569bc1c3df639ea25f9d4c7236367"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3bb6e0895a832b94335d726655a0aa3cba8e87d16dd1ce08db52b1ab5c76858e"></a>

## port — port / b84c6f40cf31 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-1c3c85be785c735afb4ad85cb212fa043eaa753afc4f7076378bdd310dc0b8dc)
- port

<a id="canonical-de8224b91fb6f558ee15c3b0f80d71b470da11af4cfe0339a86c68f10e3ad929"></a>

Type: `"list"`. Computed.

Source Ports. L4 port numbers to match.

Upstream description:

L4 port numbers to match.

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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-ab3eea6f43b8749e54611994f7dc6561abd96ecef00dd772f392e181b093f869"></a>

## Direct properties — port / b84c6f40cf31 / 3

- [all](data-sources--fast_acl_rule--reference--group-001.md#canonical-3136eb33fa34966b362f064d809abca225b54657707de91e38ebccbddafc4d1b): complete subsection reference.

- [dns](data-sources--fast_acl_rule--reference--group-001.md#canonical-abbbedac3e5e97b80675b9ab4cf46842cd98edaaa3f1d21afd8244fbf91974ff): complete subsection reference.

<a id="canonical-28029c896e4b383144ebdc300ded7670d339f07142cb5ebbdbd848098a512e9c"></a>

<a id="canonical-7467ddc0a957b72497ede8cd0492c476479e5ef4cf0f7944432eb8f0160ab6dd"></a>

## user_defined property — port / b84c6f40cf31 / 4

Type: `"number"`. Computed.

Exclusive with \[all DNS\] Matches the user defined port.

Upstream description:

Exclusive with \[all DNS\] Matches the user defined port.

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
    "create": false,
    "minimum_config": false,
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

<a id="canonical-0b776cf1f2e83dc0ede09e8ab347c7a795a40f6bed7bc56d9fb24a73983b29f3"></a>

## Next pages — port / b84c6f40cf31 / 5

- [port.all](data-sources--fast_acl_rule--reference--group-001.md#canonical-3136eb33fa34966b362f064d809abca225b54657707de91e38ebccbddafc4d1b)
- [port.dns](data-sources--fast_acl_rule--reference--group-001.md#canonical-abbbedac3e5e97b80675b9ab4cf46842cd98edaaa3f1d21afd8244fbf91974ff)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-1c3c85be785c735afb4ad85cb212fa043eaa753afc4f7076378bdd310dc0b8dc)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)

<a id="canonical-3136eb33fa34966b362f064d809abca225b54657707de91e38ebccbddafc4d1b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e82afa3ae60d199d218f9e473e6a9137ffe563c291c0c7dcf407f510e55537d0"></a>

## port.all — port.all / aaca1f977ca2 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-1c3c85be785c735afb4ad85cb212fa043eaa753afc4f7076378bdd310dc0b8dc)
- [port](data-sources--fast_acl_rule--reference--group-001.md#canonical-7a606ffc81ed836667c59b2844c5e00508c569bc1c3df639ea25f9d4c7236367)
- port.all

<a id="canonical-64e256ef70a2694fac0f728fbe14f9fd04eecd48e221af942a33ed1a23e501a1"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3178b0727dd6562a1d30aa1f61a8e598df235b7ad0fd1e5716c224d0cbaee322"></a>

## Direct properties — port.all / aaca1f977ca2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-afb5628d2d9c817aef1eeee8f04b0e694ace6ba1603a19ffc5c9870f97ae113c"></a>

## Next pages — port.all / aaca1f977ca2 / 4

- [port](data-sources--fast_acl_rule--reference--group-001.md#canonical-7a606ffc81ed836667c59b2844c5e00508c569bc1c3df639ea25f9d4c7236367)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)

<a id="canonical-abbbedac3e5e97b80675b9ab4cf46842cd98edaaa3f1d21afd8244fbf91974ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a395074c3621ea94108bf22787c17a76c98e92bafd9410d20926737a7b71901"></a>

## port.dns — port.dns / 5ceceab37410 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-1c3c85be785c735afb4ad85cb212fa043eaa753afc4f7076378bdd310dc0b8dc)
- [port](data-sources--fast_acl_rule--reference--group-001.md#canonical-7a606ffc81ed836667c59b2844c5e00508c569bc1c3df639ea25f9d4c7236367)
- port.dns

<a id="canonical-66c91f8f40048ca4ba27ba89c126499841bddd489e8f38e6c8948642c2c834fb"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-668c430c8a247e0c96d0615b5c7f137964c33fa3380b4311781c9ec85dc94dd3"></a>

## Direct properties — port.dns / 5ceceab37410 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cf98753abcc1e2f70cbcca88c1732e6aadc87f6342794070d4f10a46dda785ae"></a>

## Next pages — port.dns / 5ceceab37410 / 4

- [port](data-sources--fast_acl_rule--reference--group-001.md#canonical-7a606ffc81ed836667c59b2844c5e00508c569bc1c3df639ea25f9d4c7236367)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)

<a id="canonical-823f0fc06b588d7a6404a2291bfcacfc0c1f0d30e424bbc749870725a3a74c13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95eab7618094a870f9de4341a2dbed8fef72b3e0ebebce61f09735914fd813f5"></a>

## prefix — prefix / 55b405d7e996 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)
- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-1c3c85be785c735afb4ad85cb212fa043eaa753afc4f7076378bdd310dc0b8dc)
- prefix

<a id="canonical-4cb76dfb12535874c747668fa8d2ea83a503d1891019e3ad4b8a3c4074650924"></a>

Type: `"single"`. Computed.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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

<a id="canonical-4c3ab860205f75568d489a4cef150eefd8bdd8d595eefd57de3b58625eb85792"></a>

## Direct properties — prefix / 55b405d7e996 / 3

<a id="canonical-0d8d3ed7477e2c6b902b2861b47ca89e1be78c3ba107c1d752bed165af516f35"></a>

<a id="canonical-d701b4b5259ae026e2a1cb621f1e7822229585f552a2a57cf6a729593082f517"></a>

## prefix property — prefix / 55b405d7e996 / 4

Type: `["list", "string"]`. Computed.

IP Address prefix in string format. String must contain both prefix and prefix-length.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-af9482150749cbb02b8b4a7295636d313b89d209d8a98725670ccf656c906052"></a>

## Next pages — prefix / 55b405d7e996 / 5

- [Property reference](data-sources--fast_acl_rule--reference--group-001.md#canonical-1c3c85be785c735afb4ad85cb212fa043eaa753afc4f7076378bdd310dc0b8dc)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)
