---
page_title: "xcsh_service_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_service_policy reference."
---

# xcsh_service_policy reference

<a id="canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c924dc47e3b5d55720ae204cf18fe6b44d0c134a9cafe9e693df464f68ea5de"></a>

## Property reference — Property reference / 40aff57552bd / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- Property reference

<a id="canonical-20b561c69b4bb2549fc7a7d432e4bcce442f5df5b1741bfc51768517a8522d86"></a>

## Direct properties — Property reference / 40aff57552bd / 3

- [allow_all_requests](data-sources--service_policy--reference--group-001.md#canonical-747c4797b995158c8ec52d6bd943c7883673d696b5e3355dd4028d11b00f89f1): complete subsection reference.

- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-e4b9e7feaee766541557450a8cdbaf3082e3abbb5f08b9d29cafef55fe3c5475): complete subsection reference.

<a id="canonical-89081df4bd254feeaaf1cc737e48c4a3918dd997647536ae6079da48213c2a37"></a>

<a id="canonical-b5af2c6a3f35cb6008c0d7b844dc398d55e064c344b4450b3439ce3682bee97a"></a>

## annotations property — Property reference / 40aff57552bd / 4

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

- [any_server](data-sources--service_policy--reference--group-001.md#canonical-d13d1e67c51c8cca88bc26321867ed0fb9e64eb8d097d3bf64f69c4ac4b557ce): complete subsection reference.

- [deny_all_requests](data-sources--service_policy--reference--group-001.md#canonical-7961aa59a587848eed69a42716ee84f69ef58e7bf7d664abc780fef2cef3e212): complete subsection reference.

- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-5f93ca70b2dc6b4430d0a8a8956f8fa491038318e823bd49dea1fb5e09f0bcc7): complete subsection reference.

<a id="canonical-2e521d762c5d76f46beb3e7cc3ba89fdf01f10c203fed1468be3945c6a7c0816"></a>

<a id="canonical-d964d2a69b0a16e7d1c2fb02da5a9fca945f3e2129ff196e43a86b8d1115a999"></a>

## description property — Property reference / 40aff57552bd / 5

Type: `"string"`. Computed.

Description of the ServicePolicy.

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

<a id="canonical-191f43912dc92a29eb799f57c72027804b37c2705968581b97ff9fdbc31a5fa3"></a>

<a id="canonical-ab90dd603f841c78c039168e373cd16b41d7d7906a853a64d308ca3f2e0b8347"></a>

## id property — Property reference / 40aff57552bd / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0ab80a9a8ffc1da217bc9f60842c753f131b2c2449ce3842fc00ca3a54654a31"></a>

<a id="canonical-2adb1c6c8cc301f216f70306331d742c2cf87350787110d55b4cc67e6620dfe4"></a>

## labels property — Property reference / 40aff57552bd / 7

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

<a id="canonical-8a3aa8240af3eb83f70615e954c325f3f4861069c63ec76489e025b01ba3792d"></a>

<a id="canonical-3acf0c721783ec7e6abdc5e1e32a0d4230ade2fd26c7ffa11e39e0aa7f08d12c"></a>

## name property — Property reference / 40aff57552bd / 8

Type: `"string"`. Required.

Name of the ServicePolicy.

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

<a id="canonical-c82f25d2f2c6ef5001402302afa542e6e2312bfb0fa3d1e3fe4937a7e8377496"></a>

<a id="canonical-f6082ef56450be5ab8a893ca7d8550434aa32511c112d00a81e98241273f603d"></a>

## namespace property — Property reference / 40aff57552bd / 9

Type: `"string"`. Required.

Namespace where the ServicePolicy exists.

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

- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-2027263afc4d5af4143efdbba5049cbc80defe477cca37d5d12eae77f21d788f): complete subsection reference.

<a id="canonical-3916f8cc3d7239594d7aee991468d9238b386df038335807413250b38452251f"></a>

<a id="canonical-2d79f661f553941ec2e862f1ada48e6064ef4d563581e536ad530efa6a527fa8"></a>

## server_name property — Property reference / 40aff57552bd / 10

Type: `"string"`. Computed.

Exclusive with \[any\_server server\_name\_matcher server\_selector\] The expected name of the
server to which the request API is directed. The actual names for the server are extracted from the
HTTP Host header and the name of the virtual\_host to which the request is directed. If the request
is..

Upstream description:

Exclusive with \[any\_server server\_name\_matcher server\_selector\] The expected name of the
server to which the request API is directed. The actual names for the server are extracted from the
HTTP Host header and the name of the virtual\_host to which the request is directed. If the request
is directed to a virtual K8s service, the actual names also contain the name of that service. The
predicate evaluates to true if any of the actual names is the same as the expected server name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [server_name_matcher](data-sources--service_policy--reference--group-003.md#canonical-c95d6d05b62f043ee5b7f120193e9d9a17e5cbcff706fe8d002e233605c7743c): complete subsection reference.

- [server_selector](data-sources--service_policy--reference--group-003.md#canonical-dadfba9fca72ad25175c5cee239998db1b6d731d15f6c244d917f629bcea7f4f): complete subsection reference.

<a id="canonical-e87c6580c1c50bd008ef8ad582e1294fed89505d0000bf50b9c89a46177b509d"></a>

## All schema paths — Property reference / 40aff57552bd / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all_requests` | [allow_all_requests](data-sources--service_policy--reference--group-001.md#canonical-dc356a63d99c31d8b0e400dfc7fe27bad469370992d9bdeca64d4dccc9fa51ff) |
| `allow_list` | [allow_list](data-sources--service_policy--reference--group-001.md#canonical-f49e2d8171e4dc87e195ec1fbe958f93b4035290959a675913ec1a0a2ed34437) |
| `allow_list.asn_list` | [allow_list.asn_list](data-sources--service_policy--reference--group-001.md#canonical-7655f1c104ff561aa287fd9e94e9fad1227284ed5a19678d9a5aea6b5ca671d7) |
| `allow_list.asn_list.as_numbers` | [allow_list.asn_list.as_numbers](data-sources--service_policy--reference--group-001.md#canonical-9e94e04afa60a799ab8e04dfff6ccd2ebc4f03ca9a909cbb14b32f8ecf6e493d) |
| `allow_list.asn_set` | [allow_list.asn_set](data-sources--service_policy--reference--group-001.md#canonical-7febeaf417068ba24c419464743cbf222766e825750f81d765a24761c4a178ec) |
| `allow_list.asn_set.name` | [allow_list.asn_set.name](data-sources--service_policy--reference--group-001.md#canonical-a9be0158338b8d4c84867306dd3109cb11853df04feedb646f27f3ede750c1fe) |
| `allow_list.asn_set.namespace` | [allow_list.asn_set.namespace](data-sources--service_policy--reference--group-001.md#canonical-28af9c3095f163280fbf435c3fd45ca2125aae1de14611743108656aad4a6d23) |
| `allow_list.asn_set.tenant` | [allow_list.asn_set.tenant](data-sources--service_policy--reference--group-001.md#canonical-d10d1f3f8d303826b0ead2d5967f8aee5b32be2c7ce9278d182e8a1b88802826) |
| `allow_list.country_list` | [allow_list.country_list](data-sources--service_policy--reference--group-001.md#canonical-e1ae4fa02433499f47e2cb95dfafdcee38bd49b751bac12df1bea86574abb768) |
| `allow_list.default_action_allow` | [allow_list.default_action_allow](data-sources--service_policy--reference--group-001.md#canonical-baa13d5256408661d7f8a7ab6859b8e5b7f433610ed360dabc17fc483f91e876) |
| `allow_list.default_action_deny` | [allow_list.default_action_deny](data-sources--service_policy--reference--group-001.md#canonical-7301bff54d50346e690e01dfa1824568ccf2e742a7e76f1fda1fb7778857577e) |
| `allow_list.default_action_next_policy` | [allow_list.default_action_next_policy](data-sources--service_policy--reference--group-001.md#canonical-06ef2748c3f4d5e6b91801bcf9ffc6afdc7a9daeba1570bc1d6367eaa53b6302) |
| `allow_list.ip_prefix_set` | [allow_list.ip_prefix_set](data-sources--service_policy--reference--group-001.md#canonical-38ddb46b4dfb9e1ba307a08b5a08778f67c1d06483ccce797845ba801241c19c) |
| `allow_list.ip_prefix_set.name` | [allow_list.ip_prefix_set.name](data-sources--service_policy--reference--group-001.md#canonical-f706ac1e161ef578db3693190af4dad9cf6cecbf7dab2d2f02ba91ce287a7ef9) |
| `allow_list.ip_prefix_set.namespace` | [allow_list.ip_prefix_set.namespace](data-sources--service_policy--reference--group-001.md#canonical-fb0a3de889fa48327453df16ba8cf1168fbbcea731be5eb9e38c42d377988e6b) |
| `allow_list.ip_prefix_set.tenant` | [allow_list.ip_prefix_set.tenant](data-sources--service_policy--reference--group-001.md#canonical-12eecde2fc46f3f3122608804b1efbb8dc133fbb60038eebeb8c74e7986f9d75) |
| `allow_list.prefix_list` | [allow_list.prefix_list](data-sources--service_policy--reference--group-001.md#canonical-85b99dd6b45b7f1b3dd867c796ba8ed812d8277ade743f11eab4a32b2919cbfa) |
| `allow_list.prefix_list.prefixes` | [allow_list.prefix_list.prefixes](data-sources--service_policy--reference--group-001.md#canonical-2fa0289706ea7121b001c2181ba1918d7ac38aa8b166284df66ac61b8dfa36a9) |
| `allow_list.tls_fingerprint_classes` | [allow_list.tls_fingerprint_classes](data-sources--service_policy--reference--group-001.md#canonical-a03945426ec9aefc12951f41956100f82a792c7fb83fe77315652b4e55740582) |
| `allow_list.tls_fingerprint_values` | [allow_list.tls_fingerprint_values](data-sources--service_policy--reference--group-001.md#canonical-93ba31aac4cf485aa7a71ab770c61a4c8c9cd949db85873b0250cb8d976d33f1) |
| `annotations` | [annotations](data-sources--service_policy--reference--group-001.md#canonical-89081df4bd254feeaaf1cc737e48c4a3918dd997647536ae6079da48213c2a37) |
| `any_server` | [any_server](data-sources--service_policy--reference--group-001.md#canonical-a76f0424395f2a260b9ee0bdf623cedec4b38732f0eee5173f6cd2331e4371b3) |
| `deny_all_requests` | [deny_all_requests](data-sources--service_policy--reference--group-001.md#canonical-6f0cf978384b21bb1a08e44133a619f2e244d4425b386740cc7d004cec79abc1) |
| `deny_list` | [deny_list](data-sources--service_policy--reference--group-001.md#canonical-4aac9865ac1db7af6863d4c6aed2a8d5fd2a30df7d7d5ad5fc249ecbb673de43) |
| `deny_list.asn_list` | [deny_list.asn_list](data-sources--service_policy--reference--group-001.md#canonical-fea91b6264e848850710afa2a27a20dc59b51813571d3403e8efc857aba50190) |
| `deny_list.asn_list.as_numbers` | [deny_list.asn_list.as_numbers](data-sources--service_policy--reference--group-001.md#canonical-917db75472e86ea28691e1a4b00794be0e62a2a9d7991264f7f2bb99d2ba5fa0) |
| `deny_list.asn_set` | [deny_list.asn_set](data-sources--service_policy--reference--group-001.md#canonical-6b85b60b5921fd09ca70bb916eb0b12bab848b47bbf990010ed32deffb132a2e) |
| `deny_list.asn_set.name` | [deny_list.asn_set.name](data-sources--service_policy--reference--group-001.md#canonical-df1d09ef5b0e04d81e31a1e82d8bb1f5baf37dd39842d1b416d929b5a1a93807) |
| `deny_list.asn_set.namespace` | [deny_list.asn_set.namespace](data-sources--service_policy--reference--group-001.md#canonical-29f2e6c0e764acafc42acc689bb6a35221c1e1d57453a56d36873a85d07d5bf7) |
| `deny_list.asn_set.tenant` | [deny_list.asn_set.tenant](data-sources--service_policy--reference--group-001.md#canonical-d7bbbf59c290809ba32647bc444bb75f383c7aabcf3237c47a132befbed68245) |
| `deny_list.country_list` | [deny_list.country_list](data-sources--service_policy--reference--group-001.md#canonical-b5b80f85a16ae23847f5b24bf463c82c4af20e27336dc69ce9ca1e7f628f6538) |
| `deny_list.default_action_allow` | [deny_list.default_action_allow](data-sources--service_policy--reference--group-001.md#canonical-0aefabfab07f4dae0cd01607e6784f616b3158b5df876076579a989730a7d741) |
| `deny_list.default_action_deny` | [deny_list.default_action_deny](data-sources--service_policy--reference--group-001.md#canonical-716255feb5dd14e2415ad2b593f2c72ca4de47e92e7a6c18d126e90f2ff66d6d) |
| `deny_list.default_action_next_policy` | [deny_list.default_action_next_policy](data-sources--service_policy--reference--group-001.md#canonical-b9784205bd7a0d02ad7fa8a486b98ad874a6d630b9500fe2ce2f8d8fe462fc84) |
| `deny_list.ip_prefix_set` | [deny_list.ip_prefix_set](data-sources--service_policy--reference--group-001.md#canonical-d2af1c619762e5c73f2557c424f689a6b61a8ef2b9b36cb34173956cb1e3d9e6) |
| `deny_list.ip_prefix_set.name` | [deny_list.ip_prefix_set.name](data-sources--service_policy--reference--group-001.md#canonical-a46a4c378239dc6bc27ff2343a2b6579ef80737322766d88278e051d64aa60c1) |
| `deny_list.ip_prefix_set.namespace` | [deny_list.ip_prefix_set.namespace](data-sources--service_policy--reference--group-001.md#canonical-c8ea8f31466b4829c17db53b472900ba0701b70951078c927b3d2821fdcad8b5) |
| `deny_list.ip_prefix_set.tenant` | [deny_list.ip_prefix_set.tenant](data-sources--service_policy--reference--group-001.md#canonical-82def6ae2945491139c48d7f0df873ac4dc45859eedfd41eb903c4577d12777f) |
| `deny_list.prefix_list` | [deny_list.prefix_list](data-sources--service_policy--reference--group-001.md#canonical-1408a295a54b76825e032351ab7ccc55c2470397c43c064c7d501b89a7420806) |
| `deny_list.prefix_list.prefixes` | [deny_list.prefix_list.prefixes](data-sources--service_policy--reference--group-001.md#canonical-c7653de422a32806cd2e29a50415d87a915cc427c077d5c4cf92b8bfd6305079) |
| `deny_list.tls_fingerprint_classes` | [deny_list.tls_fingerprint_classes](data-sources--service_policy--reference--group-001.md#canonical-85070545ea69cabc94efde2320d087188d15f1980c4a7a71ac115e553c59c1f1) |
| `deny_list.tls_fingerprint_values` | [deny_list.tls_fingerprint_values](data-sources--service_policy--reference--group-001.md#canonical-b2ecb4257a91931c870340e2c7a84c5b68f7d65c98f8cd9455ed232d89187c58) |
| `description` | [description](data-sources--service_policy--reference--group-001.md#canonical-2e521d762c5d76f46beb3e7cc3ba89fdf01f10c203fed1468be3945c6a7c0816) |
| `id` | [id](data-sources--service_policy--reference--group-001.md#canonical-191f43912dc92a29eb799f57c72027804b37c2705968581b97ff9fdbc31a5fa3) |
| `labels` | [labels](data-sources--service_policy--reference--group-001.md#canonical-0ab80a9a8ffc1da217bc9f60842c753f131b2c2449ce3842fc00ca3a54654a31) |
| `name` | [name](data-sources--service_policy--reference--group-001.md#canonical-8a3aa8240af3eb83f70615e954c325f3f4861069c63ec76489e025b01ba3792d) |
| `namespace` | [namespace](data-sources--service_policy--reference--group-001.md#canonical-c82f25d2f2c6ef5001402302afa542e6e2312bfb0fa3d1e3fe4937a7e8377496) |
| `rule_list` | [rule_list](data-sources--service_policy--reference--group-001.md#canonical-7b3eacb2d5a557b8f145def9c63d10ec64852bea12d3d445430e16caa2324c2e) |
| `rule_list.rules` | [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-916cf59ab2933f7b1791921bdc25d0cae463186d742a741ed01cd3581bbaca8f) |
| `rule_list.rules.metadata` | [rule_list.rules.metadata](data-sources--service_policy--reference--group-001.md#canonical-9e6a91536f54f310ed4f7834208bf55945e4cb8bdece6094153d7015879a99f8) |
| `rule_list.rules.metadata.description_spec` | [rule_list.rules.metadata.description_spec](data-sources--service_policy--reference--group-001.md#canonical-26a6ed47caaa91ab4b30fc30c8b15fd97643fd08d0be66e5d9fb0694638eac55) |
| `rule_list.rules.metadata.name` | [rule_list.rules.metadata.name](data-sources--service_policy--reference--group-001.md#canonical-4079e765d4401d78b28345c49699c6fceca5a0b09deb318e13b9dbee2fb59ee4) |
| `rule_list.rules.spec` | [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-5f1d894f385e18178ccc42e2a8d64de00aa54e43320b6dc73bd05ee5119deecb) |
| `rule_list.rules.spec.action` | [rule_list.rules.spec.action](data-sources--service_policy--reference--group-001.md#canonical-0cd3937364c0103b3ce598be37c02c4c125f3ec669600e8409eaba936108731a) |
| `rule_list.rules.spec.any_asn` | [rule_list.rules.spec.any_asn](data-sources--service_policy--reference--group-001.md#canonical-48946e969cad3dec694accb2d90dd2dfade21a9470e2d1e8f7e933d89e7a8951) |
| `rule_list.rules.spec.any_client` | [rule_list.rules.spec.any_client](data-sources--service_policy--reference--group-002.md#canonical-70f032ff7f57ea36b4aa360e7052e69de36782adfe2e7b6528a28bafe431f2e8) |
| `rule_list.rules.spec.any_ip` | [rule_list.rules.spec.any_ip](data-sources--service_policy--reference--group-002.md#canonical-4396516adb8744120cdd8cde513875124532db41081317a591c4d39da5cdf220) |
| `rule_list.rules.spec.api_group_matcher` | [rule_list.rules.spec.api_group_matcher](data-sources--service_policy--reference--group-002.md#canonical-959786dc2d3bad93e342af27f8e234833774dff8d032856ce74d9b70fc4e1499) |
| `rule_list.rules.spec.api_group_matcher.invert_matcher` | [rule_list.rules.spec.api_group_matcher.invert_matcher](data-sources--service_policy--reference--group-002.md#canonical-d6bfef650ac82e6ae3eb2ca08aa5b717a0e192f22bf90ccaf9870000cff5ec5e) |
| `rule_list.rules.spec.api_group_matcher.match` | [rule_list.rules.spec.api_group_matcher.match](data-sources--service_policy--reference--group-002.md#canonical-be0457f3c459755712048d98405ffd20969c871260f0dd0fd9636550891424c4) |
| `rule_list.rules.spec.arg_matchers` | [rule_list.rules.spec.arg_matchers](data-sources--service_policy--reference--group-002.md#canonical-2d0954d66d764e295d3a1657348b2dc3a38818d7145a14d664078624a6272ac2) |
| `rule_list.rules.spec.arg_matchers.check_not_present` | [rule_list.rules.spec.arg_matchers.check_not_present](data-sources--service_policy--reference--group-002.md#canonical-d9c03803aa08b460c09cbeb073efe5dcee35e3f10776b1b319451935b8d54f2e) |
| `rule_list.rules.spec.arg_matchers.check_present` | [rule_list.rules.spec.arg_matchers.check_present](data-sources--service_policy--reference--group-002.md#canonical-2267e5bd7379b23532b89578bebcbcb8d136788a49b813b78279ae61709a127b) |
| `rule_list.rules.spec.arg_matchers.invert_matcher` | [rule_list.rules.spec.arg_matchers.invert_matcher](data-sources--service_policy--reference--group-002.md#canonical-33a3b9f1db17fbdb8c937543676bdc88f313e814009738c8654a01f72ba25bd9) |
| `rule_list.rules.spec.arg_matchers.item` | [rule_list.rules.spec.arg_matchers.item](data-sources--service_policy--reference--group-002.md#canonical-be2dc38a5ecb4f1e9d076e4a574f75bfcadeb14c2925aae496e641c13761fb41) |
| `rule_list.rules.spec.arg_matchers.item.exact_values` | [rule_list.rules.spec.arg_matchers.item.exact_values](data-sources--service_policy--reference--group-002.md#canonical-e7e28c6d9723dab7845df6bc80e14ae9c7833a1b114215216f211dfcd11d0673) |
| `rule_list.rules.spec.arg_matchers.item.regex_values` | [rule_list.rules.spec.arg_matchers.item.regex_values](data-sources--service_policy--reference--group-002.md#canonical-71eee976f66387bdab2e6532303ee2459adb67be5857b2db10668d4779697710) |
| `rule_list.rules.spec.arg_matchers.item.transformers` | [rule_list.rules.spec.arg_matchers.item.transformers](data-sources--service_policy--reference--group-002.md#canonical-b377814000f34105d9020079183d6c4c95009c98a80024003f845287017fc701) |
| `rule_list.rules.spec.arg_matchers.name` | [rule_list.rules.spec.arg_matchers.name](data-sources--service_policy--reference--group-002.md#canonical-baba7f409af89a2f687adbc226e2f3dc0d2d33042f899aeb157e71d5b17fa499) |
| `rule_list.rules.spec.asn_list` | [rule_list.rules.spec.asn_list](data-sources--service_policy--reference--group-002.md#canonical-4cbc8902218740e783edd4488b660536221a1cbd87406b491e0e7cb20d6cb9dd) |
| `rule_list.rules.spec.asn_list.as_numbers` | [rule_list.rules.spec.asn_list.as_numbers](data-sources--service_policy--reference--group-002.md#canonical-2c54a218b6b0070aaaed6703e21a7492dc2f6795ecb5d9712d9a7b38344727c1) |
| `rule_list.rules.spec.asn_matcher` | [rule_list.rules.spec.asn_matcher](data-sources--service_policy--reference--group-002.md#canonical-f4c752e5e6fd3b7e71afd09c7b70bc1f601337f86d0bd937c3b6ca74530133eb) |
| `rule_list.rules.spec.asn_matcher.asn_sets` | [rule_list.rules.spec.asn_matcher.asn_sets](data-sources--service_policy--reference--group-002.md#canonical-608a3dd307989a1ab59a62d251617d3ef3cfd2e40fa5a29293d8410bc8c12a83) |
| `rule_list.rules.spec.asn_matcher.asn_sets.kind` | [rule_list.rules.spec.asn_matcher.asn_sets.kind](data-sources--service_policy--reference--group-002.md#canonical-e73f26fe4c70726a50885823eebb6eadfe6080ca958f8d4527739d0b7d971a25) |
| `rule_list.rules.spec.asn_matcher.asn_sets.name` | [rule_list.rules.spec.asn_matcher.asn_sets.name](data-sources--service_policy--reference--group-002.md#canonical-0a5c394f6cfff57e13c626ed667f82373b114a89584ba260db0457b52f62d2b4) |
| `rule_list.rules.spec.asn_matcher.asn_sets.namespace` | [rule_list.rules.spec.asn_matcher.asn_sets.namespace](data-sources--service_policy--reference--group-002.md#canonical-065a3bb50a2a759634f6e422b68a1336f69f12340cdc183cca7784a4ce313983) |
| `rule_list.rules.spec.asn_matcher.asn_sets.tenant` | [rule_list.rules.spec.asn_matcher.asn_sets.tenant](data-sources--service_policy--reference--group-002.md#canonical-ac839a8b1a6d83a6e44475ed0183c4778c84acaeb9585c2f65b249b9635e6ea8) |
| `rule_list.rules.spec.asn_matcher.asn_sets.uid` | [rule_list.rules.spec.asn_matcher.asn_sets.uid](data-sources--service_policy--reference--group-002.md#canonical-4f701efdca8afb2c65830aaff9b9085fb3e52b0dee3956d292ab85be7cf967df) |
| `rule_list.rules.spec.body_matcher` | [rule_list.rules.spec.body_matcher](data-sources--service_policy--reference--group-002.md#canonical-504e9df6769fc0ed8e561305dfbb479a07abf0adebcfa1eb2c5e52b958a1e37e) |
| `rule_list.rules.spec.body_matcher.exact_values` | [rule_list.rules.spec.body_matcher.exact_values](data-sources--service_policy--reference--group-002.md#canonical-0ad03bc8770df5334468cad31c60f6a326df40db528b1af7e39dcad009747830) |
| `rule_list.rules.spec.body_matcher.regex_values` | [rule_list.rules.spec.body_matcher.regex_values](data-sources--service_policy--reference--group-002.md#canonical-0866c662d45c76bc2f622b4467e5fde1b240c845a16f60c08d554c1b957471c9) |
| `rule_list.rules.spec.body_matcher.transformers` | [rule_list.rules.spec.body_matcher.transformers](data-sources--service_policy--reference--group-002.md#canonical-7561db8a061ff7b5f622cca0bc2b3eb0f42a5442dca963b0623b35e07da53b95) |
| `rule_list.rules.spec.bot_action` | [rule_list.rules.spec.bot_action](data-sources--service_policy--reference--group-002.md#canonical-1f7b4f8f594594aa03b459595998212f905877f3443c629c22cce7ba6fcd0c20) |
| `rule_list.rules.spec.bot_action.bot_skip_processing` | [rule_list.rules.spec.bot_action.bot_skip_processing](data-sources--service_policy--reference--group-002.md#canonical-9c3d5c2369b6c05bc04cde5ed14ac99321a0557b2773910c2d3107614059b197) |
| `rule_list.rules.spec.bot_action.none` | [rule_list.rules.spec.bot_action.none](data-sources--service_policy--reference--group-002.md#canonical-b6e82381506f3ca766bece06b35f73dd00be111c1f0204b7202b7b623b182829) |
| `rule_list.rules.spec.client_name` | [rule_list.rules.spec.client_name](data-sources--service_policy--reference--group-001.md#canonical-1a6af37560ccdcef815d85e58964d442037955f00596925434b433a5e2698077) |
| `rule_list.rules.spec.client_name_matcher` | [rule_list.rules.spec.client_name_matcher](data-sources--service_policy--reference--group-002.md#canonical-ec8efb7f35c3257eee37c851e2a77904c218190a44fa0db0951350c1252bf5f0) |
| `rule_list.rules.spec.client_name_matcher.exact_values` | [rule_list.rules.spec.client_name_matcher.exact_values](data-sources--service_policy--reference--group-002.md#canonical-14a662c882f273a35a3c8af82a7efd257282fb20c67744733bf9f3b3240a3118) |
| `rule_list.rules.spec.client_name_matcher.regex_values` | [rule_list.rules.spec.client_name_matcher.regex_values](data-sources--service_policy--reference--group-002.md#canonical-ff4d028723bd8383158c027fcf38760ce170f651d993314d4cc2a7270ef12ad2) |
| `rule_list.rules.spec.client_name_matcher.transformers` | [rule_list.rules.spec.client_name_matcher.transformers](data-sources--service_policy--reference--group-002.md#canonical-10b840f84bd121221fecc22f87362d32e135c53674f9c813a633c81a3305c87d) |
| `rule_list.rules.spec.client_selector` | [rule_list.rules.spec.client_selector](data-sources--service_policy--reference--group-002.md#canonical-090cf2cc4e686feec6f3ef48fd529070f1346a1511bc2a49da4a593ce82646bb) |
| `rule_list.rules.spec.client_selector.expressions` | [rule_list.rules.spec.client_selector.expressions](data-sources--service_policy--reference--group-002.md#canonical-5a74df515e8f10f672a6cbb17480c3adac68762237ecf629336d85461d513993) |
| `rule_list.rules.spec.cookie_matchers` | [rule_list.rules.spec.cookie_matchers](data-sources--service_policy--reference--group-002.md#canonical-cace4e857df5007e703442ce368555882ceb32171a98c242ea958e5f49020073) |
| `rule_list.rules.spec.cookie_matchers.check_not_present` | [rule_list.rules.spec.cookie_matchers.check_not_present](data-sources--service_policy--reference--group-002.md#canonical-81c19acd0810628cdc5ef0436ca731f73f7c0ec6fe7cb98595153a8b37066f2f) |
| `rule_list.rules.spec.cookie_matchers.check_present` | [rule_list.rules.spec.cookie_matchers.check_present](data-sources--service_policy--reference--group-002.md#canonical-a785d91302e03dcf4a8bf7a03a5b51d50a265931ee9db452c4071046c5734a05) |
| `rule_list.rules.spec.cookie_matchers.invert_matcher` | [rule_list.rules.spec.cookie_matchers.invert_matcher](data-sources--service_policy--reference--group-002.md#canonical-62ee4c1f122c67f40d9810eaef84306a8eb81cacb94a760f0925b471cd231817) |
| `rule_list.rules.spec.cookie_matchers.item` | [rule_list.rules.spec.cookie_matchers.item](data-sources--service_policy--reference--group-002.md#canonical-24202fe4b60324cdd9ac391796ed60008a51c781be7cfdc1b0259d8fe7a3734f) |
| `rule_list.rules.spec.cookie_matchers.item.exact_values` | [rule_list.rules.spec.cookie_matchers.item.exact_values](data-sources--service_policy--reference--group-002.md#canonical-7380022ecdee1ea3479977004ddf3b9dbbff52c23f174e9fd4c3ef277678a305) |
| `rule_list.rules.spec.cookie_matchers.item.regex_values` | [rule_list.rules.spec.cookie_matchers.item.regex_values](data-sources--service_policy--reference--group-002.md#canonical-4e3933400c8a1577ea82764efb4cc2e65b9d4d4f18b86767d83414e18ffdb3ca) |
| `rule_list.rules.spec.cookie_matchers.item.transformers` | [rule_list.rules.spec.cookie_matchers.item.transformers](data-sources--service_policy--reference--group-002.md#canonical-7782168acdaa4b1e4dca599647e18a8191b7e4c330b457a0673200f533baba5e) |
| `rule_list.rules.spec.cookie_matchers.name` | [rule_list.rules.spec.cookie_matchers.name](data-sources--service_policy--reference--group-002.md#canonical-8ee8368ebd524d9529356a7e8444dc14946bc720814b77f040959901b36a1e12) |
| `rule_list.rules.spec.domain_matcher` | [rule_list.rules.spec.domain_matcher](data-sources--service_policy--reference--group-002.md#canonical-d99bff1053e1c2f5d5d4ce33415fce7da9249980b0dbd40666894cb00baeb37d) |
| `rule_list.rules.spec.domain_matcher.exact_values` | [rule_list.rules.spec.domain_matcher.exact_values](data-sources--service_policy--reference--group-002.md#canonical-0b44c766c244451d1e22b522e478f8dd813a352ae3d4daad62d8a5c49067d77d) |
| `rule_list.rules.spec.domain_matcher.regex_values` | [rule_list.rules.spec.domain_matcher.regex_values](data-sources--service_policy--reference--group-002.md#canonical-4d7a2b3cf907706d64d717207329e3927aab04fcb05e4b9a3c8b5b5dcf82620c) |
| `rule_list.rules.spec.domain_matcher.transformers` | [rule_list.rules.spec.domain_matcher.transformers](data-sources--service_policy--reference--group-002.md#canonical-7f949e0a833f96597118adeb642cb04af16367298968e813e748d4d6e2000287) |
| `rule_list.rules.spec.expiration_timestamp` | [rule_list.rules.spec.expiration_timestamp](data-sources--service_policy--reference--group-001.md#canonical-a9159a904b96f396792f61be3806afbcc8b057ab625626d3d990f4c8cd128640) |
| `rule_list.rules.spec.headers` | [rule_list.rules.spec.headers](data-sources--service_policy--reference--group-002.md#canonical-41ecdadce0720cde054532bae9bab83ead10d0aab32449fd2adf317dc2feab49) |
| `rule_list.rules.spec.headers.check_not_present` | [rule_list.rules.spec.headers.check_not_present](data-sources--service_policy--reference--group-002.md#canonical-17a0df7ad9726ff28219906b6f95527bc67fbdb9fb5713236ff9960d7105bdd2) |
| `rule_list.rules.spec.headers.check_present` | [rule_list.rules.spec.headers.check_present](data-sources--service_policy--reference--group-002.md#canonical-7b2d41f2db3733712da5667e8063c9dcc9154dc83818e7ad04608f287d5d585d) |
| `rule_list.rules.spec.headers.invert_matcher` | [rule_list.rules.spec.headers.invert_matcher](data-sources--service_policy--reference--group-002.md#canonical-a02308e4926e5fe3400c8b430e385cb20843c152137919d90f38fa923b79801e) |
| `rule_list.rules.spec.headers.item` | [rule_list.rules.spec.headers.item](data-sources--service_policy--reference--group-002.md#canonical-1dc3df46e31592c7ecf4b8373b3b82154b081fa5a4e492c4e63aba99847a138e) |
| `rule_list.rules.spec.headers.item.exact_values` | [rule_list.rules.spec.headers.item.exact_values](data-sources--service_policy--reference--group-002.md#canonical-616d9f2f38f637d82938220b5df1f9d5d1d4dc3879581684d6246af310004cbc) |
| `rule_list.rules.spec.headers.item.regex_values` | [rule_list.rules.spec.headers.item.regex_values](data-sources--service_policy--reference--group-002.md#canonical-e6e8b6b7e3e41429531eb7e4366f534585641a3654f32290ec0dafb5fc156158) |
| `rule_list.rules.spec.headers.item.transformers` | [rule_list.rules.spec.headers.item.transformers](data-sources--service_policy--reference--group-002.md#canonical-ea9ed8590a74fab294e1e02fea21b08b772be5a63bb292228606504186a5c809) |
| `rule_list.rules.spec.headers.name` | [rule_list.rules.spec.headers.name](data-sources--service_policy--reference--group-002.md#canonical-0d69dbbb358686c3d6b9c380415ccae02a51281cece20af94638c1f3083c876d) |
| `rule_list.rules.spec.http_method` | [rule_list.rules.spec.http_method](data-sources--service_policy--reference--group-002.md#canonical-b7e7d81c29dc9d5ca2c48e19cffc6558d09dd329c56b36c05b08cdc1c5707bfa) |
| `rule_list.rules.spec.http_method.invert_matcher` | [rule_list.rules.spec.http_method.invert_matcher](data-sources--service_policy--reference--group-002.md#canonical-c2b7350a365d27817967208c9e7b059e87c53ac848c3ffbcb4039e50c394ce11) |
| `rule_list.rules.spec.http_method.methods` | [rule_list.rules.spec.http_method.methods](data-sources--service_policy--reference--group-002.md#canonical-56b966ee4ff82260370af53a5b107b2ce313dab24805fa06c74a2e2e46db8bdf) |
| `rule_list.rules.spec.ip_matcher` | [rule_list.rules.spec.ip_matcher](data-sources--service_policy--reference--group-002.md#canonical-7d4ce312a20ae9ea59a314e17644f61baa6825a7c5123bb298c0b86369b25d41) |
| `rule_list.rules.spec.ip_matcher.invert_matcher` | [rule_list.rules.spec.ip_matcher.invert_matcher](data-sources--service_policy--reference--group-002.md#canonical-f491280dd71c28678cebdbc28f8351d1004de59b02c6dcbdfd8bdd4c31addcb7) |
| `rule_list.rules.spec.ip_matcher.prefix_sets` | [rule_list.rules.spec.ip_matcher.prefix_sets](data-sources--service_policy--reference--group-002.md#canonical-3c160bd353a3be0575f690582523a47cce5ba16383623f51c831c2d7eca570b7) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.kind` | [rule_list.rules.spec.ip_matcher.prefix_sets.kind](data-sources--service_policy--reference--group-002.md#canonical-ba1ef55f45b1777895a147114221a3ffda5e5ef4a32596f52f248de4f74ebb1e) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.name` | [rule_list.rules.spec.ip_matcher.prefix_sets.name](data-sources--service_policy--reference--group-002.md#canonical-3cc4ddf0e5ee88073da5b196c1d44dfc0113041c7abb74d0fde40ae04ba7a792) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.namespace` | [rule_list.rules.spec.ip_matcher.prefix_sets.namespace](data-sources--service_policy--reference--group-002.md#canonical-c136aa52a901c304f82b5a29a61074fc534fd4557a206cff6c9272239d927f05) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.tenant` | [rule_list.rules.spec.ip_matcher.prefix_sets.tenant](data-sources--service_policy--reference--group-002.md#canonical-9985eaac7f9d860b6fbc61b5c7a3b26474c71ff9b61bd0e49b2a5fbb1b936391) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.uid` | [rule_list.rules.spec.ip_matcher.prefix_sets.uid](data-sources--service_policy--reference--group-002.md#canonical-1ad75bce4a106e318f223a9ddae51f23f1c210e8244dad70bd7c0ffe6fa3197f) |
| `rule_list.rules.spec.ip_prefix_list` | [rule_list.rules.spec.ip_prefix_list](data-sources--service_policy--reference--group-002.md#canonical-3ee147e1d33b5aa7b626a2ee3c8f07d86ea5300e09c004e7a8f3af9351b5a453) |
| `rule_list.rules.spec.ip_prefix_list.invert_match` | [rule_list.rules.spec.ip_prefix_list.invert_match](data-sources--service_policy--reference--group-002.md#canonical-c8ab8a7ee575016f2f271c6bd839679aba63acaf8eca2cd3be25e93c136e7f5d) |
| `rule_list.rules.spec.ip_prefix_list.ip_prefixes` | [rule_list.rules.spec.ip_prefix_list.ip_prefixes](data-sources--service_policy--reference--group-002.md#canonical-3f9bff93cc76fb3afe5951d4073bd86f9eb8102cbe2691973773b5fe1a7fffac) |
| `rule_list.rules.spec.ip_threat_category_list` | [rule_list.rules.spec.ip_threat_category_list](data-sources--service_policy--reference--group-002.md#canonical-3ee1532ef45d980a4df23912a68fe936e119eec0208414cb24d1eed5c4f44050) |
| `rule_list.rules.spec.ip_threat_category_list.ip_threat_categories` | [rule_list.rules.spec.ip_threat_category_list.ip_threat_categories](data-sources--service_policy--reference--group-002.md#canonical-921001776d5f579c40c1769e9e9d755b74e5051f68bab50b7446198021aa54dc) |
| `rule_list.rules.spec.ja4_tls_fingerprint` | [rule_list.rules.spec.ja4_tls_fingerprint](data-sources--service_policy--reference--group-002.md#canonical-9fe6787e133ddd4b1f181b70a47d9979450fbcfab7d45d40f4eddb6219653b44) |
| `rule_list.rules.spec.ja4_tls_fingerprint.exact_values` | [rule_list.rules.spec.ja4_tls_fingerprint.exact_values](data-sources--service_policy--reference--group-002.md#canonical-1140dc105a752db332447172dfc364293ce52e293962ef5da8ef5f5717dcce38) |
| `rule_list.rules.spec.jwt_claims` | [rule_list.rules.spec.jwt_claims](data-sources--service_policy--reference--group-002.md#canonical-7fd7ad3cd2342f7ee538b4ef6b5c65dcd621e0c929763a7c5b73d8a2e4949f1a) |
| `rule_list.rules.spec.jwt_claims.check_not_present` | [rule_list.rules.spec.jwt_claims.check_not_present](data-sources--service_policy--reference--group-002.md#canonical-5e1008f106665650960ce291cafadcb79c000091ad062e4e37f1197e32b592f3) |
| `rule_list.rules.spec.jwt_claims.check_present` | [rule_list.rules.spec.jwt_claims.check_present](data-sources--service_policy--reference--group-002.md#canonical-491550c815e518f97a403d40ac86807968b71801f9f9cb3dceaaef79abc86af5) |
| `rule_list.rules.spec.jwt_claims.invert_matcher` | [rule_list.rules.spec.jwt_claims.invert_matcher](data-sources--service_policy--reference--group-002.md#canonical-54a068716eb8239485b52c7e3e0b9e7077b04352e637a8f859ab2e274d3c133a) |
| `rule_list.rules.spec.jwt_claims.item` | [rule_list.rules.spec.jwt_claims.item](data-sources--service_policy--reference--group-002.md#canonical-f3c2dd06a4a6017acffa3c2b4fbdecda5e6d3006706d6be731548eb4ab37610e) |
| `rule_list.rules.spec.jwt_claims.item.exact_values` | [rule_list.rules.spec.jwt_claims.item.exact_values](data-sources--service_policy--reference--group-002.md#canonical-9d0165c3d2b23985b755216538f5c14a452ccdf59660ac0a52bd6b5cb9c4da08) |
| `rule_list.rules.spec.jwt_claims.item.regex_values` | [rule_list.rules.spec.jwt_claims.item.regex_values](data-sources--service_policy--reference--group-002.md#canonical-fdd0136e2e6999f4391c64b65a2e6d7712f906631f720ff6fe98a23c0a08d2c4) |
| `rule_list.rules.spec.jwt_claims.item.transformers` | [rule_list.rules.spec.jwt_claims.item.transformers](data-sources--service_policy--reference--group-002.md#canonical-3757d4f899fb192f2bae4e2e20a2e3ccadf1fb9ad33e612189df0573978f9320) |
| `rule_list.rules.spec.jwt_claims.name` | [rule_list.rules.spec.jwt_claims.name](data-sources--service_policy--reference--group-002.md#canonical-9fb819433f03ca83a31d5138af6d88b1ce952678deb8149252dd6500ab0c437f) |
| `rule_list.rules.spec.label_matcher` | [rule_list.rules.spec.label_matcher](data-sources--service_policy--reference--group-002.md#canonical-38ac53ec64224579a12a4bcdfb35e82c332709888bf360350f3e1fb4f49326ce) |
| `rule_list.rules.spec.label_matcher.keys` | [rule_list.rules.spec.label_matcher.keys](data-sources--service_policy--reference--group-002.md#canonical-4c4a6dead98cb425ad17852ab8e1433a08bab96c6dd909d555a80c3a95f47f46) |
| `rule_list.rules.spec.log_rule_evaluation` | [rule_list.rules.spec.log_rule_evaluation](data-sources--service_policy--reference--group-001.md#canonical-d032f9e29d9790dc374806d3f4920adef7bfaf9da17715c2b0812f395ea1a66b) |
| `rule_list.rules.spec.mum_action` | [rule_list.rules.spec.mum_action](data-sources--service_policy--reference--group-002.md#canonical-c559e7528f8a6a9ad3ed6e2beda275ab82793aa58dc5c068cd8fc1094323efef) |
| `rule_list.rules.spec.mum_action.default` | [rule_list.rules.spec.mum_action.default](data-sources--service_policy--reference--group-002.md#canonical-589129a2044c840aa9c5c30d159e1055208f33a5e369561ea473452a30dcefca) |
| `rule_list.rules.spec.mum_action.skip_processing` | [rule_list.rules.spec.mum_action.skip_processing](data-sources--service_policy--reference--group-002.md#canonical-3c969de038107988f0f6ecf9b46dbdbeadcd0c9cad61adce6e4926204a9e9eb8) |
| `rule_list.rules.spec.path` | [rule_list.rules.spec.path](data-sources--service_policy--reference--group-002.md#canonical-d68bbffc4215eb173176e132ebd3aa17a5023b61fd8cd315a5350342aec4c0cb) |
| `rule_list.rules.spec.path.encoded_path_matcher` | [rule_list.rules.spec.path.encoded_path_matcher](data-sources--service_policy--reference--group-002.md#canonical-7c41ac3d44dece2a8e061275647e3f96007e26a98b5197ca2cd8b02124be1af5) |
| `rule_list.rules.spec.path.exact_values` | [rule_list.rules.spec.path.exact_values](data-sources--service_policy--reference--group-002.md#canonical-18d885ce4c60c2f9e4265d9458e06c60a37361a8c5bb98891445079c83d0b00b) |
| `rule_list.rules.spec.path.invert_matcher` | [rule_list.rules.spec.path.invert_matcher](data-sources--service_policy--reference--group-002.md#canonical-01e9375c54bc64d757230a7a819f823c21ebb679045a544808293d222f2187bc) |
| `rule_list.rules.spec.path.prefix_values` | [rule_list.rules.spec.path.prefix_values](data-sources--service_policy--reference--group-002.md#canonical-44d83eebe178b2e4e6c6941b6cc02fa5b1da25236bd99b3decaafc28b1d63585) |
| `rule_list.rules.spec.path.regex_values` | [rule_list.rules.spec.path.regex_values](data-sources--service_policy--reference--group-002.md#canonical-784fb7e52666866466417834d7e14cad163a52bccff43828185ddf84bd46b51b) |
| `rule_list.rules.spec.path.suffix_values` | [rule_list.rules.spec.path.suffix_values](data-sources--service_policy--reference--group-002.md#canonical-457a160cb3dc60dd23043fcad368ca8ad809f189090323dd3222cea9325d213b) |
| `rule_list.rules.spec.path.transformers` | [rule_list.rules.spec.path.transformers](data-sources--service_policy--reference--group-002.md#canonical-6bf8f93a4202c7b179625f1858532465812360f30268467e2eaa7302864c3553) |
| `rule_list.rules.spec.port_matcher` | [rule_list.rules.spec.port_matcher](data-sources--service_policy--reference--group-002.md#canonical-12234c1dcf3d0d7bd23bf3fe2323925617227058d6ca7e18d861c2df3ba02eb2) |
| `rule_list.rules.spec.port_matcher.invert_matcher` | [rule_list.rules.spec.port_matcher.invert_matcher](data-sources--service_policy--reference--group-002.md#canonical-5a1122ab77e897f0e16edf4b2a418d1bef1bc211db9ad1c04ec257d3a232ff36) |
| `rule_list.rules.spec.port_matcher.ports` | [rule_list.rules.spec.port_matcher.ports](data-sources--service_policy--reference--group-002.md#canonical-1ebbd88f15d0889674f7c7e42bb94a258e1816d164fae684e65e7aef8a96e62d) |
| `rule_list.rules.spec.query_params` | [rule_list.rules.spec.query_params](data-sources--service_policy--reference--group-002.md#canonical-093a45fa3f1e08e11ce8d40c25471286683634803537f0504dfc9f7dccba1dc4) |
| `rule_list.rules.spec.query_params.check_not_present` | [rule_list.rules.spec.query_params.check_not_present](data-sources--service_policy--reference--group-002.md#canonical-cdc498d1cb4ff0274cdd8d7412eb3b914e025e8b8de7290e2ed397f070a875c8) |
| `rule_list.rules.spec.query_params.check_present` | [rule_list.rules.spec.query_params.check_present](data-sources--service_policy--reference--group-002.md#canonical-787fa08d28662acd44deeda2c977103ce3d1fc7703b8a1199808d3390a5014e4) |
| `rule_list.rules.spec.query_params.invert_matcher` | [rule_list.rules.spec.query_params.invert_matcher](data-sources--service_policy--reference--group-002.md#canonical-28575ee9417b72064f265234308ec6da1da1580ea64a4195e4da8b18cb28652f) |
| `rule_list.rules.spec.query_params.item` | [rule_list.rules.spec.query_params.item](data-sources--service_policy--reference--group-002.md#canonical-ce6f453650fb920a8db11de0da0516332aa62b63f6da14c9c4f1e4203d2974f5) |
| `rule_list.rules.spec.query_params.item.exact_values` | [rule_list.rules.spec.query_params.item.exact_values](data-sources--service_policy--reference--group-002.md#canonical-cd538d21e712321f743480a41a88b53fa520e7650f7f02a0b21dbcdcfec330a8) |
| `rule_list.rules.spec.query_params.item.regex_values` | [rule_list.rules.spec.query_params.item.regex_values](data-sources--service_policy--reference--group-002.md#canonical-8fcf022ca0b10c593926e740a737a544c059120a09cc05040e075dc8311a18b5) |
| `rule_list.rules.spec.query_params.item.transformers` | [rule_list.rules.spec.query_params.item.transformers](data-sources--service_policy--reference--group-002.md#canonical-86b5307ee35175c9b6690455d191b83584af0d422732a7025a962c436f3df566) |
| `rule_list.rules.spec.query_params.key` | [rule_list.rules.spec.query_params.key](data-sources--service_policy--reference--group-002.md#canonical-931759b68e1964a035ee947504db6a3b303afdc0da7952547a5505383b5f4e3d) |
| `rule_list.rules.spec.request_constraints` | [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-966fd25e3d725ab368e2061c9bf9c38e6ef020ce4316839f20fa4107f3bf2e43) |
| `rule_list.rules.spec.request_constraints.max_cookie_count_exceeds` | [rule_list.rules.spec.request_constraints.max_cookie_count_exceeds](data-sources--service_policy--reference--group-002.md#canonical-5d90dae75edc850345ad8ff290f226d9eb15c669e1cdcf1b2376d76f7d2e7a05) |
| `rule_list.rules.spec.request_constraints.max_cookie_count_none` | [rule_list.rules.spec.request_constraints.max_cookie_count_none](data-sources--service_policy--reference--group-002.md#canonical-55ddb830e6a1fba3702f8e611a73f083c1735cb640c6a0777bbdd61c63ab4d6b) |
| `rule_list.rules.spec.request_constraints.max_cookie_key_size_exceeds` | [rule_list.rules.spec.request_constraints.max_cookie_key_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-cdac1e2789d54630ceec20290d3d737d9156f8ef06d7542615e9e35b28f1553d) |
| `rule_list.rules.spec.request_constraints.max_cookie_key_size_none` | [rule_list.rules.spec.request_constraints.max_cookie_key_size_none](data-sources--service_policy--reference--group-003.md#canonical-89a45c4ac598aefac1ad4c2635b0279df46a3b17fe39e18dbe2466fae1a73ae7) |
| `rule_list.rules.spec.request_constraints.max_cookie_value_size_exceeds` | [rule_list.rules.spec.request_constraints.max_cookie_value_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-99716e4d77f53bd28867a192313e6320f16f6f61dd6d3b6c01f1eb4ed87f04a3) |
| `rule_list.rules.spec.request_constraints.max_cookie_value_size_none` | [rule_list.rules.spec.request_constraints.max_cookie_value_size_none](data-sources--service_policy--reference--group-003.md#canonical-19f3497283aa593426b2cc64fc2d3db06879759ef77f572a3f289f43576db03f) |
| `rule_list.rules.spec.request_constraints.max_header_count_exceeds` | [rule_list.rules.spec.request_constraints.max_header_count_exceeds](data-sources--service_policy--reference--group-002.md#canonical-ac7a5252f34a6b719a339c8bc9ecff1a1db621f5b28bbbab8986e76e5eadf923) |
| `rule_list.rules.spec.request_constraints.max_header_count_none` | [rule_list.rules.spec.request_constraints.max_header_count_none](data-sources--service_policy--reference--group-003.md#canonical-75de11d7666cfc4721ea83c1854d264bee7e128573d1d1b9606051398fc6a2dc) |
| `rule_list.rules.spec.request_constraints.max_header_key_size_exceeds` | [rule_list.rules.spec.request_constraints.max_header_key_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-f2b7a96c0a5a1cc1d334cc478e3737087ddc8c9b79a9f6630712e3a6481e9186) |
| `rule_list.rules.spec.request_constraints.max_header_key_size_none` | [rule_list.rules.spec.request_constraints.max_header_key_size_none](data-sources--service_policy--reference--group-003.md#canonical-a90a005be6cdb1526a71b147846ea30bbcb187b4527b3d693f925fc7a8ba8a20) |
| `rule_list.rules.spec.request_constraints.max_header_value_size_exceeds` | [rule_list.rules.spec.request_constraints.max_header_value_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-fd8bb37f7888baf1c16397d1d0dea502af24f03d11f6f4be816162044228759d) |
| `rule_list.rules.spec.request_constraints.max_header_value_size_none` | [rule_list.rules.spec.request_constraints.max_header_value_size_none](data-sources--service_policy--reference--group-003.md#canonical-8f05e81bb932cbada67271e999e56e5bc6674ef2cb7f46b2d011845b03b314d3) |
| `rule_list.rules.spec.request_constraints.max_parameter_count_exceeds` | [rule_list.rules.spec.request_constraints.max_parameter_count_exceeds](data-sources--service_policy--reference--group-002.md#canonical-1151f0de8ea98d3ec2d4f53d7ff17e857039938d11e853c0f8bc92da6128c36b) |
| `rule_list.rules.spec.request_constraints.max_parameter_count_none` | [rule_list.rules.spec.request_constraints.max_parameter_count_none](data-sources--service_policy--reference--group-003.md#canonical-37eaf5696426781beda663e5f6d3c4392138c4120a7dbb25e672456abf41426c) |
| `rule_list.rules.spec.request_constraints.max_parameter_name_size_exceeds` | [rule_list.rules.spec.request_constraints.max_parameter_name_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-8ebb10f400ea8f8f48d64684412ebf56d9639c818a835918dfdd21cedb3e7601) |
| `rule_list.rules.spec.request_constraints.max_parameter_name_size_none` | [rule_list.rules.spec.request_constraints.max_parameter_name_size_none](data-sources--service_policy--reference--group-003.md#canonical-d1d83eee0e5d493db95e83e70b703da1ab98cee871ed9027f8e6f3c6e6fe8ef2) |
| `rule_list.rules.spec.request_constraints.max_parameter_value_size_exceeds` | [rule_list.rules.spec.request_constraints.max_parameter_value_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-8f6bc6059a2589d490534ecaad6130e674024f72ab455b1e0942b529f35f065e) |
| `rule_list.rules.spec.request_constraints.max_parameter_value_size_none` | [rule_list.rules.spec.request_constraints.max_parameter_value_size_none](data-sources--service_policy--reference--group-003.md#canonical-02938a887052edffc6c3478e45cfbff510f72af583fd13b3b19054386dad349b) |
| `rule_list.rules.spec.request_constraints.max_query_size_exceeds` | [rule_list.rules.spec.request_constraints.max_query_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-0e13c8b8c207192dfb655ad9cf8049c5ad5d90bf249529a08a99c295f6a6f741) |
| `rule_list.rules.spec.request_constraints.max_query_size_none` | [rule_list.rules.spec.request_constraints.max_query_size_none](data-sources--service_policy--reference--group-003.md#canonical-b2e764a02fdfa60fb47be38e41568499cc2323bbf95b6efabdcff40883d6b8a1) |
| `rule_list.rules.spec.request_constraints.max_request_line_size_exceeds` | [rule_list.rules.spec.request_constraints.max_request_line_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-6f8872cd22d588bfc6dcc358924547d5e46e546e2d3874208edfba5506cf28f5) |
| `rule_list.rules.spec.request_constraints.max_request_line_size_none` | [rule_list.rules.spec.request_constraints.max_request_line_size_none](data-sources--service_policy--reference--group-003.md#canonical-dd46bfb3c93bd2d1aeec0e562cfdbf65154d5ce07b59e2d2155066f398e38657) |
| `rule_list.rules.spec.request_constraints.max_request_size_exceeds` | [rule_list.rules.spec.request_constraints.max_request_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-05bad50cfbceb8c197a6ef96d9774fca3370f85509c38ef2faebe9ada84a8817) |
| `rule_list.rules.spec.request_constraints.max_request_size_none` | [rule_list.rules.spec.request_constraints.max_request_size_none](data-sources--service_policy--reference--group-003.md#canonical-8ec29de7bb27ba74ac0c33f69174f4fd242370d3f840fc5452c571b2a5791593) |
| `rule_list.rules.spec.request_constraints.max_url_size_exceeds` | [rule_list.rules.spec.request_constraints.max_url_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-d2e7a80821a60206be8392d4c9de9b1ea07a4a53d3a96bea2606aa6dd84afac7) |
| `rule_list.rules.spec.request_constraints.max_url_size_none` | [rule_list.rules.spec.request_constraints.max_url_size_none](data-sources--service_policy--reference--group-003.md#canonical-772cb8f00e0c91d0726d383199ac4d94874829af4200707a19b62b2d27e79f43) |
| `rule_list.rules.spec.segment_policy` | [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-003.md#canonical-fbbae3bc5a97dcdc4bca40fd0769a1d8423703c1bd8ac8c84d5e9b9c42991367) |
| `rule_list.rules.spec.segment_policy.dst_any` | [rule_list.rules.spec.segment_policy.dst_any](data-sources--service_policy--reference--group-003.md#canonical-79dbfb6586ea60fc42d2bfcc6241e1e28c7916cd5c76b8d54bf9893f9ad50d1f) |
| `rule_list.rules.spec.segment_policy.dst_segments` | [rule_list.rules.spec.segment_policy.dst_segments](data-sources--service_policy--reference--group-003.md#canonical-0d9e32be4a5ac0afa3e4260a5f5bd562f0d7d1270a4ba39b4b202ca5c12492c2) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments` | [rule_list.rules.spec.segment_policy.dst_segments.segments](data-sources--service_policy--reference--group-003.md#canonical-42d1e7f421bb0f47316f76bf5d5d3884e86baee1a1ad40c9b3596b1673861100) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments.name` | [rule_list.rules.spec.segment_policy.dst_segments.segments.name](data-sources--service_policy--reference--group-003.md#canonical-6952cc741d866e5a170c3d932e10d35954cb696f8c8b997f23fce8702bbbc0dc) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments.namespace` | [rule_list.rules.spec.segment_policy.dst_segments.segments.namespace](data-sources--service_policy--reference--group-003.md#canonical-988b9ae49a5a6329a927c859472ccc6230664f3f10d5cde5c78abea946e30aa1) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments.tenant` | [rule_list.rules.spec.segment_policy.dst_segments.segments.tenant](data-sources--service_policy--reference--group-003.md#canonical-c5c4e4ce6fc19cd70d6891dfa8fc28c4af7772f541fcf43cce328102ef3904f1) |
| `rule_list.rules.spec.segment_policy.intra_segment` | [rule_list.rules.spec.segment_policy.intra_segment](data-sources--service_policy--reference--group-003.md#canonical-6c51ca960f58a948202e70bc6b6ba96a9ba4bdfec72818a301d898590f9ed29d) |
| `rule_list.rules.spec.segment_policy.src_any` | [rule_list.rules.spec.segment_policy.src_any](data-sources--service_policy--reference--group-003.md#canonical-514cdb122c219c8981fcf04dcb0b6a2e7201daeaf779c3140b343fb4de0bf463) |
| `rule_list.rules.spec.segment_policy.src_segments` | [rule_list.rules.spec.segment_policy.src_segments](data-sources--service_policy--reference--group-003.md#canonical-b17d3873812ffe511e10e43b4c8db0895026026f7847ff84acf44996a4b886b4) |
| `rule_list.rules.spec.segment_policy.src_segments.segments` | [rule_list.rules.spec.segment_policy.src_segments.segments](data-sources--service_policy--reference--group-003.md#canonical-c4bd8147280c3b9068711ce72e8d09c75a7c04d64fa272112469931f4c32724b) |
| `rule_list.rules.spec.segment_policy.src_segments.segments.name` | [rule_list.rules.spec.segment_policy.src_segments.segments.name](data-sources--service_policy--reference--group-003.md#canonical-b87f9570e58b0f00db224a921661272dbc96ca9a50998e04a9bf9efc51e1d184) |
| `rule_list.rules.spec.segment_policy.src_segments.segments.namespace` | [rule_list.rules.spec.segment_policy.src_segments.segments.namespace](data-sources--service_policy--reference--group-003.md#canonical-6cc5a70a23218f0bbc0074e87443fe0fa7404e9c823a843625328ddfbad2e897) |
| `rule_list.rules.spec.segment_policy.src_segments.segments.tenant` | [rule_list.rules.spec.segment_policy.src_segments.segments.tenant](data-sources--service_policy--reference--group-003.md#canonical-797ea9eef7810b82ed29bf2d1ccbd0349f4566e14ceeb75f7ddc84cfb6cc65cb) |
| `rule_list.rules.spec.tls_fingerprint_matcher` | [rule_list.rules.spec.tls_fingerprint_matcher](data-sources--service_policy--reference--group-003.md#canonical-f540ec92bb8807ec8b381fc9619318a56b06e33ec24a3453367041da5f72007b) |
| `rule_list.rules.spec.tls_fingerprint_matcher.classes` | [rule_list.rules.spec.tls_fingerprint_matcher.classes](data-sources--service_policy--reference--group-003.md#canonical-3f66bf34fadce86c0762a4f0c38d8b5b7eb03e81162a346e07b5126db9ffffa8) |
| `rule_list.rules.spec.tls_fingerprint_matcher.exact_values` | [rule_list.rules.spec.tls_fingerprint_matcher.exact_values](data-sources--service_policy--reference--group-003.md#canonical-85b469a289e78ce136f84ab7aa62d09dbc7e88d7096b0dc9c6659d8d2bcc710c) |
| `rule_list.rules.spec.tls_fingerprint_matcher.excluded_values` | [rule_list.rules.spec.tls_fingerprint_matcher.excluded_values](data-sources--service_policy--reference--group-003.md#canonical-60776dfd09c5406d883f12e388422b29308d785c7f2453a176b5bfbf005cd7e9) |
| `rule_list.rules.spec.user_identity_matcher` | [rule_list.rules.spec.user_identity_matcher](data-sources--service_policy--reference--group-003.md#canonical-1099cbd274e7613f8cfc90882b1e498fe17b01e2e8f8ddb1bf8e2182250e6ac6) |
| `rule_list.rules.spec.user_identity_matcher.exact_values` | [rule_list.rules.spec.user_identity_matcher.exact_values](data-sources--service_policy--reference--group-003.md#canonical-af7bfa9d2db5133d6b590e1a7b889c861a09e067b25bb7532ed1f966de28c8a9) |
| `rule_list.rules.spec.user_identity_matcher.regex_values` | [rule_list.rules.spec.user_identity_matcher.regex_values](data-sources--service_policy--reference--group-003.md#canonical-5dd8bd896c464495b5f71d20e1b82a83a2918c24ad372458e9b5b944ef54e42b) |
| `rule_list.rules.spec.waf_action` | [rule_list.rules.spec.waf_action](data-sources--service_policy--reference--group-003.md#canonical-a19c7e55092f2abac6304600376499dc27d55145de05b2840783360104da9544) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control` | [rule_list.rules.spec.waf_action.app_firewall_detection_control](data-sources--service_policy--reference--group-003.md#canonical-089ebdf9987f4e1d859007b6d7c4827b323327dc23174db416ab74c641c49cc2) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts](data-sources--service_policy--reference--group-003.md#canonical-37aea1f0ab127be3c05a65697016e65c469961f2c1f3aa0a61d36fdffcb2905f) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context](data-sources--service_policy--reference--group-003.md#canonical-d94df78c1c8db46debc3f63f0542f2c891b7bb95566e7722c81320f039830c79) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name](data-sources--service_policy--reference--group-003.md#canonical-d4b63fe79079348b82b978ebc5d70f05a8d939a04d9c50f982edc1337a7f28f3) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type](data-sources--service_policy--reference--group-003.md#canonical-8647b5dcbac4b2ff412adb435dd188d69b112c9a7e0723ccbecfd36e18fa54b1) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts](data-sources--service_policy--reference--group-003.md#canonical-102b8726190872fe9929c1bd47609e1453bd304d8898840fc76c260fe75ab9bf) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name](data-sources--service_policy--reference--group-003.md#canonical-6b366c4cf260c955b519de2ce81ea0b647b81bea544a204e46593e4afa8e775d) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts](data-sources--service_policy--reference--group-003.md#canonical-eece4376cf12a8d938402cdd8cc36c149b78dcbbfdf0d7b66aab506b78d32878) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context](data-sources--service_policy--reference--group-003.md#canonical-7e1365f66a3e05dedcd0affa8ac5a389d259a477f481d23e4687731d62bbcec5) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name](data-sources--service_policy--reference--group-003.md#canonical-085457c1efa85aea5d137d61ec02f87d5b4ed9026ecfee82b1e437a9ce4b0ae8) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id](data-sources--service_policy--reference--group-003.md#canonical-df123355cc110c2867f64f10ee1675ed73b3b3dc4a4de39d173a7b2d66480243) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts](data-sources--service_policy--reference--group-003.md#canonical-69dab81149b7a80b382d023bf76ea5dbff6074197ba8254353cf531b7215aba5) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context](data-sources--service_policy--reference--group-003.md#canonical-686fa44b183284c869d0f3941078fd4cdffa63a0835f2ecb9d52f04482fbfd66) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name](data-sources--service_policy--reference--group-003.md#canonical-bd7a0506bb7cc08c538e7a9dc431250148e0934b6bcf4fad5ca79d50c10acddc) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation](data-sources--service_policy--reference--group-003.md#canonical-f97f104aec9d122c138a3404ed393b1278a7281428b6d1c3400866fa2b3a6d9b) |
| `rule_list.rules.spec.waf_action.none` | [rule_list.rules.spec.waf_action.none](data-sources--service_policy--reference--group-003.md#canonical-a3d4f42cd50c5e077e713cda7444999e7852f60fee58df1cf4f3af0f42eaebd9) |
| `rule_list.rules.spec.waf_action.waf_skip_processing` | [rule_list.rules.spec.waf_action.waf_skip_processing](data-sources--service_policy--reference--group-003.md#canonical-1e61f11a495c064feeb525571b4acae79c6f65235e17c97d99d4ea79374c9964) |
| `server_name` | [server_name](data-sources--service_policy--reference--group-001.md#canonical-3916f8cc3d7239594d7aee991468d9238b386df038335807413250b38452251f) |
| `server_name_matcher` | [server_name_matcher](data-sources--service_policy--reference--group-003.md#canonical-4d3f96e2c8e3a8de97582dd3e3ce3501f0a17a72f690d51e1a9db840ac0d811f) |
| `server_name_matcher.exact_values` | [server_name_matcher.exact_values](data-sources--service_policy--reference--group-003.md#canonical-d2558c550b0c4cd459b04c3c3c7d8f68d23ddb0a216dcc5c0f2e0f112232ac7a) |
| `server_name_matcher.regex_values` | [server_name_matcher.regex_values](data-sources--service_policy--reference--group-003.md#canonical-0f8f3feebc8ae8792775b3f8c4060263fea3bc8d382fab80ed0a0bdaebf1bf2a) |
| `server_selector` | [server_selector](data-sources--service_policy--reference--group-003.md#canonical-cd989dd4321adcaa5387f6e956b713a6f6d4b43c39206c6e3df31a38d669f456) |
| `server_selector.expressions` | [server_selector.expressions](data-sources--service_policy--reference--group-003.md#canonical-2877ef73b4084199b20e0a4b5e32cd6284a411335a6303d2faec7c02114f1520) |

<a id="canonical-75938846c318b69aed37fe3e1e5a6ba76b8c3994091d126e9962b5887cdb8912"></a>

## Next pages — Property reference / 40aff57552bd / 12

- [allow_all_requests](data-sources--service_policy--reference--group-001.md#canonical-747c4797b995158c8ec52d6bd943c7883673d696b5e3355dd4028d11b00f89f1)
- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-e4b9e7feaee766541557450a8cdbaf3082e3abbb5f08b9d29cafef55fe3c5475)
- [any_server](data-sources--service_policy--reference--group-001.md#canonical-d13d1e67c51c8cca88bc26321867ed0fb9e64eb8d097d3bf64f69c4ac4b557ce)
- [deny_all_requests](data-sources--service_policy--reference--group-001.md#canonical-7961aa59a587848eed69a42716ee84f69ef58e7bf7d664abc780fef2cef3e212)
- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-5f93ca70b2dc6b4430d0a8a8956f8fa491038318e823bd49dea1fb5e09f0bcc7)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-2027263afc4d5af4143efdbba5049cbc80defe477cca37d5d12eae77f21d788f)
- [server_name_matcher](data-sources--service_policy--reference--group-003.md#canonical-c95d6d05b62f043ee5b7f120193e9d9a17e5cbcff706fe8d002e233605c7743c)
- [server_selector](data-sources--service_policy--reference--group-003.md#canonical-dadfba9fca72ad25175c5cee239998db1b6d731d15f6c244d917f629bcea7f4f)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-747c4797b995158c8ec52d6bd943c7883673d696b5e3355dd4028d11b00f89f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63c647e17fd3f1370c42e7f5e44bb598b0d41d619e6967ccc418079738447aeb"></a>

## allow_all_requests — allow_all_requests / 1d027233d730 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- allow_all_requests

<a id="canonical-dc356a63d99c31d8b0e400dfc7fe27bad469370992d9bdeca64d4dccc9fa51ff"></a>

Type: `["object", {}]`. Computed.

\[OneOf: allow\_all\_requests, allow\_list, deny\_all\_requests, deny\_list, rule\_list\]
Configuration parameter for allow all requests.

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

OneOf alternatives in this subsection:

- [allow_all_requests](data-sources--service_policy--reference--group-001.md#canonical-dc356a63d99c31d8b0e400dfc7fe27bad469370992d9bdeca64d4dccc9fa51ff)
- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-f49e2d8171e4dc87e195ec1fbe958f93b4035290959a675913ec1a0a2ed34437)
- [deny_all_requests](data-sources--service_policy--reference--group-001.md#canonical-6f0cf978384b21bb1a08e44133a619f2e244d4425b386740cc7d004cec79abc1)
- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-4aac9865ac1db7af6863d4c6aed2a8d5fd2a30df7d7d5ad5fc249ecbb673de43)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-7b3eacb2d5a557b8f145def9c63d10ec64852bea12d3d445430e16caa2324c2e)

Select alternatives according to the provider validators above.

<a id="canonical-40f757f55547cfd23f000d7aae356f7f969ee4811b718f26dec7c7484178d1ae"></a>

## Direct properties — allow_all_requests / 1d027233d730 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-431e92cfda6bc42ba9828f122cca2abe65473a807b0f5c771f63d8dd66062ee5"></a>

## Next pages — allow_all_requests / 1d027233d730 / 4

- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-e4b9e7feaee766541557450a8cdbaf3082e3abbb5f08b9d29cafef55fe3c5475"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2d125a2af684787569feb00b71dfdeca51bad49ddfd6fa636fb2013a05aa32a"></a>

## allow_list — allow_list / d5e20bbbbe6a / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- allow_list

<a id="canonical-f49e2d8171e4dc87e195ec1fbe958f93b4035290959a675913ec1a0a2ed34437"></a>

Type: `"single"`. Computed.

List of sources. A request belongs to this list if it satisfies any of the match criteria.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_action_choice": "[\"default_action_allow\",\"default_action_deny\",\"default_action_next_policy\"]"
}
```

<a id="canonical-d9393641c5a1fcf3e3d6fe0936754bbdc3749886f0005b573bb9ce4c79b4acbe"></a>

## Direct properties — allow_list / d5e20bbbbe6a / 3

- [asn_list](data-sources--service_policy--reference--group-001.md#canonical-4c2f3b6e63a10dfa6111260fa6385ce7421b12a480d4d80f477c2239fef27b00): complete subsection reference.

- [asn_set](data-sources--service_policy--reference--group-001.md#canonical-b032263fee050854f67fa4967969e1a1f915c52e3a4a370dde2acdecacd48d00): complete subsection reference.

<a id="canonical-e1ae4fa02433499f47e2cb95dfafdcee38bd49b751bac12df1bea86574abb768"></a>

<a id="canonical-223e10088269663c9c737bf4c36a5dd9350879c70a42ceb9e22f16dbeb117437"></a>

## country_list property — allow_list / d5e20bbbbe6a / 4

Type: `["list", "string"]`. Computed.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Addresses that belong to one of the countries in the given list The country is obtained by
performing a lookup for the source IPv4 Address in a GeoIP DB. Possible values are
\`COUNTRY\_NONE\`, \`COUNTRY\_AD\`, \`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`,
\`COUNTRY\_AI\`, \`COUNTRY\_AL\`, \`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`,
\`COUNTRY\_AQ\`, \`COUNTRY\_AR\`, \`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`,
\`COUNTRY\_AW\`, \`COUNTRY\_AX\`, \`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`,
\`COUNTRY\_BD\`, \`COUNTRY\_BE\`, \`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`,
\`COUNTRY\_BI\`, \`COUNTRY\_BJ\`, \`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`,
\`COUNTRY\_BO\`, \`COUNTRY\_BQ\`, \`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`,
\`COUNTRY\_BV\`, \`COUNTRY\_BW\`, \`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`,
\`COUNTRY\_CC\`, \`COUNTRY\_CD\`, \`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`,
\`COUNTRY\_CI\`, \`COUNTRY\_CK\`, \`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`,
\`COUNTRY\_CO\`, \`COUNTRY\_CR\`, \`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`,
\`COUNTRY\_CW\`, \`COUNTRY\_CX\`, \`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`,
\`COUNTRY\_DJ\`, \`COUNTRY\_DK\`, \`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`,
\`COUNTRY\_EC\`, \`COUNTRY\_EE\`, \`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`,
\`COUNTRY\_ES\`, \`COUNTRY\_ET\`, \`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`,
\`COUNTRY\_FM\`, \`COUNTRY\_FO\`, \`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`,
\`COUNTRY\_GD\`, \`COUNTRY\_GE\`, \`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`,
\`COUNTRY\_GI\`, \`COUNTRY\_GL\`, \`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`,
\`COUNTRY\_GQ\`, \`COUNTRY\_GR\`, \`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`,
\`COUNTRY\_GW\`, \`COUNTRY\_GY\`, \`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`,
\`COUNTRY\_HR\`, \`COUNTRY\_HT\`, \`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`,
\`COUNTRY\_IL\`, \`COUNTRY\_IM\`, \`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`,
\`COUNTRY\_IR\`, \`COUNTRY\_IS\`, \`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`,
\`COUNTRY\_JO\`, \`COUNTRY\_JP\`, \`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`,
\`COUNTRY\_KI\`, \`COUNTRY\_KM\`, \`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`,
\`COUNTRY\_KW\`, \`COUNTRY\_KY\`, \`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`,
\`COUNTRY\_LC\`, \`COUNTRY\_LI\`, \`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`,
\`COUNTRY\_LT\`, \`COUNTRY\_LU\`, \`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`,
\`COUNTRY\_MC\`, \`COUNTRY\_MD\`, \`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`,
\`COUNTRY\_MH\`, \`COUNTRY\_MK\`, \`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`,
\`COUNTRY\_MO\`, \`COUNTRY\_MP\`, \`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`,
\`COUNTRY\_MT\`, \`COUNTRY\_MU\`, \`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`,
\`COUNTRY\_MY\`, \`COUNTRY\_MZ\`, \`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`,
\`COUNTRY\_NF\`, \`COUNTRY\_NG\`, \`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`,
\`COUNTRY\_NP\`, \`COUNTRY\_NR\`, \`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`,
\`COUNTRY\_PA\`, \`COUNTRY\_PE\`, \`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`,
\`COUNTRY\_PK\`, \`COUNTRY\_PL\`, \`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`,
\`COUNTRY\_PS\`, \`COUNTRY\_PT\`, \`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`,
\`COUNTRY\_RE\`, \`COUNTRY\_RO\`, \`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`,
\`COUNTRY\_SA\`, \`COUNTRY\_SB\`, \`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`,
\`COUNTRY\_SG\`, \`COUNTRY\_SH\`, \`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`,
\`COUNTRY\_SL\`, \`COUNTRY\_SM\`, \`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`,
\`COUNTRY\_SS\`, \`COUNTRY\_ST\`, \`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`,
\`COUNTRY\_SZ\`, \`COUNTRY\_TC\`, \`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`,
\`COUNTRY\_TH\`, \`COUNTRY\_TJ\`, \`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`,
\`COUNTRY\_TN\`, \`COUNTRY\_TO\`, \`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`,
\`COUNTRY\_TW\`, \`COUNTRY\_TZ\`, \`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`,
\`COUNTRY\_US\`, \`COUNTRY\_UY\`, \`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`,
\`COUNTRY\_VE\`, \`COUNTRY\_VG\`, \`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`,
\`COUNTRY\_WF\`, \`COUNTRY\_WS\`, \`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`,
\`COUNTRY\_YT\`, \`COUNTRY\_ZA\`, \`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

Upstream description:

Addresses that belong to one of the countries in the given list The country is obtained by
performing a lookup for the source IPv4 Address in a GeoIP DB.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_action_allow](data-sources--service_policy--reference--group-001.md#canonical-713b3ac2bf6453aee9d9b4ec02fa348467163ff5d0ca9925ac45329212314ee4): complete subsection reference.

- [default_action_deny](data-sources--service_policy--reference--group-001.md#canonical-aadfb64f0738e87c2bf319b660f7bc0aaed7b34bf545e8bf203f8d385a3d4240): complete subsection reference.

- [default_action_next_policy](data-sources--service_policy--reference--group-001.md#canonical-1b92746974180bb81433314a5f5f084c02c705eb198ec6fc62072f9d3d96c626): complete subsection reference.

- [ip_prefix_set](data-sources--service_policy--reference--group-001.md#canonical-fdb715529d6440c662fb2006204d88a7430fbd71b109dea755664bda96ae5e87): complete subsection reference.

- [prefix_list](data-sources--service_policy--reference--group-001.md#canonical-3b25aa4b7391a5c1e1fdd24c8996b724c2d8d2e517080cd6d1dc393b3f309093): complete subsection reference.

<a id="canonical-a03945426ec9aefc12951f41956100f82a792c7fb83fe77315652b4e55740582"></a>

<a id="canonical-126e0ce4e872c0c5a074e9f4a8114e958fa7ded57ea46ba45d0005ef7bcd143c"></a>

## tls_fingerprint_classes property — allow_list / d5e20bbbbe6a / 5

Type: `["list", "string"]`. Computed.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-93ba31aac4cf485aa7a71ab770c61a4c8c9cd949db85873b0250cb8d976d33f1"></a>

<a id="canonical-3f5ae8578069d03e8564012cf820e5468096d6f8de8388b4a0039aefa77869b4"></a>

## tls_fingerprint_values property — allow_list / d5e20bbbbe6a / 6

Type: `["list", "string"]`. Computed.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-535fe739dca16e91c73800f649af2ad217eda5a23f789ba065ed1560dd69d210"></a>

## Next pages — allow_list / d5e20bbbbe6a / 7

- [allow_list.asn_list](data-sources--service_policy--reference--group-001.md#canonical-4c2f3b6e63a10dfa6111260fa6385ce7421b12a480d4d80f477c2239fef27b00)
- [allow_list.asn_set](data-sources--service_policy--reference--group-001.md#canonical-b032263fee050854f67fa4967969e1a1f915c52e3a4a370dde2acdecacd48d00)
- [allow_list.default_action_allow](data-sources--service_policy--reference--group-001.md#canonical-713b3ac2bf6453aee9d9b4ec02fa348467163ff5d0ca9925ac45329212314ee4)
- [allow_list.default_action_deny](data-sources--service_policy--reference--group-001.md#canonical-aadfb64f0738e87c2bf319b660f7bc0aaed7b34bf545e8bf203f8d385a3d4240)
- [allow_list.default_action_next_policy](data-sources--service_policy--reference--group-001.md#canonical-1b92746974180bb81433314a5f5f084c02c705eb198ec6fc62072f9d3d96c626)
- [allow_list.ip_prefix_set](data-sources--service_policy--reference--group-001.md#canonical-fdb715529d6440c662fb2006204d88a7430fbd71b109dea755664bda96ae5e87)
- [allow_list.prefix_list](data-sources--service_policy--reference--group-001.md#canonical-3b25aa4b7391a5c1e1fdd24c8996b724c2d8d2e517080cd6d1dc393b3f309093)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-4c2f3b6e63a10dfa6111260fa6385ce7421b12a480d4d80f477c2239fef27b00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-077af0317548807a61a30dec9db3502304662e4bda6d5480d4eb9ceb87ca5884"></a>

## allow_list.asn_list — allow_list.asn_list / 028d944a74d6 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-e4b9e7feaee766541557450a8cdbaf3082e3abbb5f08b9d29cafef55fe3c5475)
- allow_list.asn_list

<a id="canonical-7655f1c104ff561aa287fd9e94e9fad1227284ed5a19678d9a5aea6b5ca671d7"></a>

Type: `"single"`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-9853c35fac3139fa9f9950537a99b02bbc10ddeacc65f46ecb125266c11b6fdd"></a>

## Direct properties — allow_list.asn_list / 028d944a74d6 / 3

<a id="canonical-9e94e04afa60a799ab8e04dfff6ccd2ebc4f03ca9a909cbb14b32f8ecf6e493d"></a>

<a id="canonical-9f45bb9cede827d31d2ba1f6f76366c2f2933f10b6df41067ebd2e83a13c89b4"></a>

## as_numbers property — allow_list.asn_list / 028d944a74d6 / 4

Type: `["list", "number"]`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3d6597d6a298bedd650437c4f8c9e4badb7f18ec1f6e6d77590fd3c1010bbe53"></a>

## Next pages — allow_list.asn_list / 028d944a74d6 / 5

- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-e4b9e7feaee766541557450a8cdbaf3082e3abbb5f08b9d29cafef55fe3c5475)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-b032263fee050854f67fa4967969e1a1f915c52e3a4a370dde2acdecacd48d00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5141c6539ed6b8594626ff23a839445b53cbc93993b1caa1c4b29f9b160cf13"></a>

## allow_list.asn_set — allow_list.asn_set / 757c9c774a79 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-e4b9e7feaee766541557450a8cdbaf3082e3abbb5f08b9d29cafef55fe3c5475)
- allow_list.asn_set

<a id="canonical-7febeaf417068ba24c419464743cbf222766e825750f81d765a24761c4a178ec"></a>

Type: `"list"`. Computed.

Addresses that belong to the ASNs in the given bgp\_asn\_set The ASN is obtained by performing a
lookup for the source IPv4 Address in a GeoIP DB.

Upstream description:

Addresses that belong to the ASNs in the given bgp\_asn\_set The ASN is obtained by performing a
lookup for the source IPv4 Address in a GeoIP DB.

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

<a id="canonical-be762de08fdcf41314bbee3de4286d28fee5bc7e9f2b991c2840886079e042d9"></a>

## Direct properties — allow_list.asn_set / 757c9c774a79 / 3

<a id="canonical-a9be0158338b8d4c84867306dd3109cb11853df04feedb646f27f3ede750c1fe"></a>

<a id="canonical-6844be19725dcc0a3806897df70d679914ce1b87d4784275d5b3c74697af6085"></a>

## name property — allow_list.asn_set / 757c9c774a79 / 4

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

<a id="canonical-28af9c3095f163280fbf435c3fd45ca2125aae1de14611743108656aad4a6d23"></a>

<a id="canonical-878304c4e4ec6a451d320be219e7664df746bb1a4463c2bd2708a6ce87b93630"></a>

## namespace property — allow_list.asn_set / 757c9c774a79 / 5

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

<a id="canonical-d10d1f3f8d303826b0ead2d5967f8aee5b32be2c7ce9278d182e8a1b88802826"></a>

<a id="canonical-c12210e1e69ed178ef53d8e113232597e8a46ec0d8f4083c821a6605c627d062"></a>

## tenant property — allow_list.asn_set / 757c9c774a79 / 6

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

<a id="canonical-56670bb3915a2ff8a530b228eb8b44dffdadc5308c76c470238dd566ea39468f"></a>

## Next pages — allow_list.asn_set / 757c9c774a79 / 7

- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-e4b9e7feaee766541557450a8cdbaf3082e3abbb5f08b9d29cafef55fe3c5475)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-713b3ac2bf6453aee9d9b4ec02fa348467163ff5d0ca9925ac45329212314ee4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aedf67d1dfd952c3d2b4b5a9696b965a2a3c72376946e24037ccec4a258d7b5c"></a>

## allow_list.default_action_allow — allow_list.default_action_allow / de48ab0a1bbe / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-e4b9e7feaee766541557450a8cdbaf3082e3abbb5f08b9d29cafef55fe3c5475)
- allow_list.default_action_allow

<a id="canonical-baa13d5256408661d7f8a7ab6859b8e5b7f433610ed360dabc17fc483f91e876"></a>

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

<a id="canonical-0ee7ecbfb96c254640af06143f03a36fa0874f6db6112a877309058bbee478dc"></a>

## Direct properties — allow_list.default_action_allow / de48ab0a1bbe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7372ece2f82959771773759576bff14c12b6fab6ee8dd15e45b38a8bf0c3b32f"></a>

## Next pages — allow_list.default_action_allow / de48ab0a1bbe / 4

- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-e4b9e7feaee766541557450a8cdbaf3082e3abbb5f08b9d29cafef55fe3c5475)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-aadfb64f0738e87c2bf319b660f7bc0aaed7b34bf545e8bf203f8d385a3d4240"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f297777f1d8ea65f1faa707eca24dc7c282857b44cef91cffa60b63e1e90b63"></a>

## allow_list.default_action_deny — allow_list.default_action_deny / 85e6041ca3be / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-e4b9e7feaee766541557450a8cdbaf3082e3abbb5f08b9d29cafef55fe3c5475)
- allow_list.default_action_deny

<a id="canonical-7301bff54d50346e690e01dfa1824568ccf2e742a7e76f1fda1fb7778857577e"></a>

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

<a id="canonical-7a12a090e9ecba18cf46d30437d8c20db1a4473987e24fa9bc96a28b8d158efc"></a>

## Direct properties — allow_list.default_action_deny / 85e6041ca3be / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0bfc9e7043ba0c6f01336a4b3a8636dcc0117df7b2e17a5f51da03a054654f3b"></a>

## Next pages — allow_list.default_action_deny / 85e6041ca3be / 4

- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-e4b9e7feaee766541557450a8cdbaf3082e3abbb5f08b9d29cafef55fe3c5475)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-1b92746974180bb81433314a5f5f084c02c705eb198ec6fc62072f9d3d96c626"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8eda9f23571b06cfd28066494e1bd48dcfd5e3d1fc2b3d040eba14a1f1a8cde1"></a>

## allow_list.default_action_next_policy — allow_list.default_action_next_policy / 10ce8a60aac4 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-e4b9e7feaee766541557450a8cdbaf3082e3abbb5f08b9d29cafef55fe3c5475)
- allow_list.default_action_next_policy

<a id="canonical-06ef2748c3f4d5e6b91801bcf9ffc6afdc7a9daeba1570bc1d6367eaa53b6302"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1122b6686d72d9dc4d3dd59d7e05ac684e21ccc4c296974023e9683748009b17"></a>

## Direct properties — allow_list.default_action_next_policy / 10ce8a60aac4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0b74ec09a2c7e8363d92707058ad9f494a4af4b31b1b2b4320a8c5bab0d341fc"></a>

## Next pages — allow_list.default_action_next_policy / 10ce8a60aac4 / 4

- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-e4b9e7feaee766541557450a8cdbaf3082e3abbb5f08b9d29cafef55fe3c5475)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-fdb715529d6440c662fb2006204d88a7430fbd71b109dea755664bda96ae5e87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5f9833f7bb98b31e25ee7acb9c0c2450a92fba9c96dadac15c91f6ecf93832f"></a>

## allow_list.ip_prefix_set — allow_list.ip_prefix_set / c182d969ee5b / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-e4b9e7feaee766541557450a8cdbaf3082e3abbb5f08b9d29cafef55fe3c5475)
- allow_list.ip_prefix_set

<a id="canonical-38ddb46b4dfb9e1ba307a08b5a08778f67c1d06483ccce797845ba801241c19c"></a>

Type: `"list"`. Computed.

Addresses that are covered by the prefixes in the given ip\_prefix\_set.

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

<a id="canonical-d08a88d07fcd13072cbf7fcd351665351aedadc939554ed19f508ea7d8fb4a65"></a>

## Direct properties — allow_list.ip_prefix_set / c182d969ee5b / 3

<a id="canonical-f706ac1e161ef578db3693190af4dad9cf6cecbf7dab2d2f02ba91ce287a7ef9"></a>

<a id="canonical-665db688df9b6b318eca0af465d5d013e301dd8a7d5142cc4575ea834515ebb3"></a>

## name property — allow_list.ip_prefix_set / c182d969ee5b / 4

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

<a id="canonical-fb0a3de889fa48327453df16ba8cf1168fbbcea731be5eb9e38c42d377988e6b"></a>

<a id="canonical-62c159961a536a2457577c6f00406ee12be3e5418b92842dc880e87d1614f879"></a>

## namespace property — allow_list.ip_prefix_set / c182d969ee5b / 5

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

<a id="canonical-12eecde2fc46f3f3122608804b1efbb8dc133fbb60038eebeb8c74e7986f9d75"></a>

<a id="canonical-f73498a27cd779733f048bb0f857959570e639f314ad1e8d680c140fb132ef0f"></a>

## tenant property — allow_list.ip_prefix_set / c182d969ee5b / 6

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

<a id="canonical-bec0971dbe5dcb2c192623856bbca79f728027f45df8860b6cb361a001085e8b"></a>

## Next pages — allow_list.ip_prefix_set / c182d969ee5b / 7

- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-e4b9e7feaee766541557450a8cdbaf3082e3abbb5f08b9d29cafef55fe3c5475)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-3b25aa4b7391a5c1e1fdd24c8996b724c2d8d2e517080cd6d1dc393b3f309093"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8da692858f52ce38deb8ba10eb430ede45c1a0251d79a1b10ae3896cbf13057a"></a>

## allow_list.prefix_list — allow_list.prefix_list / 5d05b8d75f7f / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-e4b9e7feaee766541557450a8cdbaf3082e3abbb5f08b9d29cafef55fe3c5475)
- allow_list.prefix_list

<a id="canonical-85b99dd6b45b7f1b3dd867c796ba8ed812d8277ade743f11eab4a32b2919cbfa"></a>

Type: `"single"`. Computed.

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

<a id="canonical-dca185388c30ce5dbe5eedf7437bb66447463ccfb3886636bc7749abf0517307"></a>

## Direct properties — allow_list.prefix_list / 5d05b8d75f7f / 3

<a id="canonical-2fa0289706ea7121b001c2181ba1918d7ac38aa8b166284df66ac61b8dfa36a9"></a>

<a id="canonical-f658ff3397588583616ba8c0236ce7280e2765efc57c010a1bbc8c9f94ef5575"></a>

## prefixes property — allow_list.prefix_list / 5d05b8d75f7f / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-156db7c9e8fe458afc9e09b08a1463966b5d542539539e841184bfdd841b8f17"></a>

## Next pages — allow_list.prefix_list / 5d05b8d75f7f / 5

- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-e4b9e7feaee766541557450a8cdbaf3082e3abbb5f08b9d29cafef55fe3c5475)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-d13d1e67c51c8cca88bc26321867ed0fb9e64eb8d097d3bf64f69c4ac4b557ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1b3a00533e057a45ce9c88ea585718498870df34e36102a50bfd0e199388bdb"></a>

## any_server — any_server / 7d9a6a0d71f1 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- any_server

<a id="canonical-a76f0424395f2a260b9ee0bdf623cedec4b38732f0eee5173f6cd2331e4371b3"></a>

Type: `["object", {}]`. Computed.

\[OneOf: any\_server, server\_name, server\_name\_matcher, server\_selector\] Enable this option.
Defaults to \`map\[\]\`. Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [any_server](data-sources--service_policy--reference--group-001.md#canonical-a76f0424395f2a260b9ee0bdf623cedec4b38732f0eee5173f6cd2331e4371b3)
- [server_name](data-sources--service_policy--reference--group-001.md#canonical-3916f8cc3d7239594d7aee991468d9238b386df038335807413250b38452251f)
- [server_name_matcher](data-sources--service_policy--reference--group-003.md#canonical-4d3f96e2c8e3a8de97582dd3e3ce3501f0a17a72f690d51e1a9db840ac0d811f)
- [server_selector](data-sources--service_policy--reference--group-003.md#canonical-cd989dd4321adcaa5387f6e956b713a6f6d4b43c39206c6e3df31a38d669f456)

Select alternatives according to the provider validators above.

<a id="canonical-57a17b20468a7199331757a8a39038d9ed6468f41fd9564c87bd8ce934297d4f"></a>

## Direct properties — any_server / 7d9a6a0d71f1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cb8277ee53af2461888f896286aa4649d41705d9744cfa7ab76daea7d2c50541"></a>

## Next pages — any_server / 7d9a6a0d71f1 / 4

- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-7961aa59a587848eed69a42716ee84f69ef58e7bf7d664abc780fef2cef3e212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37ebfdc7ebb9ca256058470bb71fe6db368c6d0531400d0d8eb214e8806b76d8"></a>

## deny_all_requests — deny_all_requests / aa5084888ee4 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- deny_all_requests

<a id="canonical-6f0cf978384b21bb1a08e44133a619f2e244d4425b386740cc7d004cec79abc1"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for deny all requests.

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

<a id="canonical-d9d591724f696a7bc53ad1280b2e4b8a1d03b990e938281a217ec7f1cdbf0a78"></a>

## Direct properties — deny_all_requests / aa5084888ee4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-97b228206c478863ed4c7c009c0b5d26a751874ad80e9c20c6e598c12424acad"></a>

## Next pages — deny_all_requests / aa5084888ee4 / 4

- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-5f93ca70b2dc6b4430d0a8a8956f8fa491038318e823bd49dea1fb5e09f0bcc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-879783b961f979523916dcb694db7e65602ce56156dc38de55262c78827af636"></a>

## deny_list — deny_list / 948fae6ddbac / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- deny_list

<a id="canonical-4aac9865ac1db7af6863d4c6aed2a8d5fd2a30df7d7d5ad5fc249ecbb673de43"></a>

Type: `"single"`. Computed.

List of sources. A request belongs to this list if it satisfies any of the match criteria.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_action_choice": "[\"default_action_allow\",\"default_action_deny\",\"default_action_next_policy\"]"
}
```

<a id="canonical-ff151685d296d94f383544b6d7eb53aa3b72e3cd27dc8dfeff8033ec16ff0080"></a>

## Direct properties — deny_list / 948fae6ddbac / 3

- [asn_list](data-sources--service_policy--reference--group-001.md#canonical-057bdd37763b2ef7dc6bab3b36d56f8531a47c9eae337954f8e5afcf85f9cd67): complete subsection reference.

- [asn_set](data-sources--service_policy--reference--group-001.md#canonical-fcd78ad770ae94979038ae81de2c1488ad0008304ea739a1cd7284c8bee74c60): complete subsection reference.

<a id="canonical-b5b80f85a16ae23847f5b24bf463c82c4af20e27336dc69ce9ca1e7f628f6538"></a>

<a id="canonical-57a577b50707ab05b5833d1c2f4074d37a2bc2986abd9d20c578fa99a43bc315"></a>

## country_list property — deny_list / 948fae6ddbac / 4

Type: `["list", "string"]`. Computed.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Addresses that belong to one of the countries in the given list The country is obtained by
performing a lookup for the source IPv4 Address in a GeoIP DB. Possible values are
\`COUNTRY\_NONE\`, \`COUNTRY\_AD\`, \`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`,
\`COUNTRY\_AI\`, \`COUNTRY\_AL\`, \`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`,
\`COUNTRY\_AQ\`, \`COUNTRY\_AR\`, \`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`,
\`COUNTRY\_AW\`, \`COUNTRY\_AX\`, \`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`,
\`COUNTRY\_BD\`, \`COUNTRY\_BE\`, \`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`,
\`COUNTRY\_BI\`, \`COUNTRY\_BJ\`, \`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`,
\`COUNTRY\_BO\`, \`COUNTRY\_BQ\`, \`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`,
\`COUNTRY\_BV\`, \`COUNTRY\_BW\`, \`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`,
\`COUNTRY\_CC\`, \`COUNTRY\_CD\`, \`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`,
\`COUNTRY\_CI\`, \`COUNTRY\_CK\`, \`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`,
\`COUNTRY\_CO\`, \`COUNTRY\_CR\`, \`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`,
\`COUNTRY\_CW\`, \`COUNTRY\_CX\`, \`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`,
\`COUNTRY\_DJ\`, \`COUNTRY\_DK\`, \`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`,
\`COUNTRY\_EC\`, \`COUNTRY\_EE\`, \`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`,
\`COUNTRY\_ES\`, \`COUNTRY\_ET\`, \`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`,
\`COUNTRY\_FM\`, \`COUNTRY\_FO\`, \`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`,
\`COUNTRY\_GD\`, \`COUNTRY\_GE\`, \`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`,
\`COUNTRY\_GI\`, \`COUNTRY\_GL\`, \`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`,
\`COUNTRY\_GQ\`, \`COUNTRY\_GR\`, \`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`,
\`COUNTRY\_GW\`, \`COUNTRY\_GY\`, \`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`,
\`COUNTRY\_HR\`, \`COUNTRY\_HT\`, \`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`,
\`COUNTRY\_IL\`, \`COUNTRY\_IM\`, \`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`,
\`COUNTRY\_IR\`, \`COUNTRY\_IS\`, \`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`,
\`COUNTRY\_JO\`, \`COUNTRY\_JP\`, \`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`,
\`COUNTRY\_KI\`, \`COUNTRY\_KM\`, \`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`,
\`COUNTRY\_KW\`, \`COUNTRY\_KY\`, \`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`,
\`COUNTRY\_LC\`, \`COUNTRY\_LI\`, \`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`,
\`COUNTRY\_LT\`, \`COUNTRY\_LU\`, \`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`,
\`COUNTRY\_MC\`, \`COUNTRY\_MD\`, \`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`,
\`COUNTRY\_MH\`, \`COUNTRY\_MK\`, \`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`,
\`COUNTRY\_MO\`, \`COUNTRY\_MP\`, \`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`,
\`COUNTRY\_MT\`, \`COUNTRY\_MU\`, \`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`,
\`COUNTRY\_MY\`, \`COUNTRY\_MZ\`, \`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`,
\`COUNTRY\_NF\`, \`COUNTRY\_NG\`, \`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`,
\`COUNTRY\_NP\`, \`COUNTRY\_NR\`, \`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`,
\`COUNTRY\_PA\`, \`COUNTRY\_PE\`, \`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`,
\`COUNTRY\_PK\`, \`COUNTRY\_PL\`, \`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`,
\`COUNTRY\_PS\`, \`COUNTRY\_PT\`, \`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`,
\`COUNTRY\_RE\`, \`COUNTRY\_RO\`, \`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`,
\`COUNTRY\_SA\`, \`COUNTRY\_SB\`, \`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`,
\`COUNTRY\_SG\`, \`COUNTRY\_SH\`, \`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`,
\`COUNTRY\_SL\`, \`COUNTRY\_SM\`, \`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`,
\`COUNTRY\_SS\`, \`COUNTRY\_ST\`, \`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`,
\`COUNTRY\_SZ\`, \`COUNTRY\_TC\`, \`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`,
\`COUNTRY\_TH\`, \`COUNTRY\_TJ\`, \`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`,
\`COUNTRY\_TN\`, \`COUNTRY\_TO\`, \`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`,
\`COUNTRY\_TW\`, \`COUNTRY\_TZ\`, \`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`,
\`COUNTRY\_US\`, \`COUNTRY\_UY\`, \`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`,
\`COUNTRY\_VE\`, \`COUNTRY\_VG\`, \`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`,
\`COUNTRY\_WF\`, \`COUNTRY\_WS\`, \`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`,
\`COUNTRY\_YT\`, \`COUNTRY\_ZA\`, \`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

Upstream description:

Addresses that belong to one of the countries in the given list The country is obtained by
performing a lookup for the source IPv4 Address in a GeoIP DB.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_action_allow](data-sources--service_policy--reference--group-001.md#canonical-d63f0da05fa71fd44d567e016d894e232c07d18200657212bb6d60dcead79bc2): complete subsection reference.

- [default_action_deny](data-sources--service_policy--reference--group-001.md#canonical-d0b71679b2fd2cf85b18a04e305a03f6065754fe9290f53f33b921d39e1e8b59): complete subsection reference.

- [default_action_next_policy](data-sources--service_policy--reference--group-001.md#canonical-f90b43e0925a1f1f3b7c4b1fbcacf43e9d65a49e9c5e555672419ee2bed9f7d5): complete subsection reference.

- [ip_prefix_set](data-sources--service_policy--reference--group-001.md#canonical-0f7fcd65ad1acde210a778200b887ce4f6a823cb37e58edac5bc4310ef3f1801): complete subsection reference.

- [prefix_list](data-sources--service_policy--reference--group-001.md#canonical-aa03658908e5fe2a886495417ce0876e8e00de3bc6f1489680312df4e8dac519): complete subsection reference.

<a id="canonical-85070545ea69cabc94efde2320d087188d15f1980c4a7a71ac115e553c59c1f1"></a>

<a id="canonical-a107606d5ba4c0f6553f693855b96d94b4628d19154c2b5c1940049882f6e3e7"></a>

## tls_fingerprint_classes property — deny_list / 948fae6ddbac / 5

Type: `["list", "string"]`. Computed.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-b2ecb4257a91931c870340e2c7a84c5b68f7d65c98f8cd9455ed232d89187c58"></a>

<a id="canonical-ae26a432c33dac0fba599b8886dd987e3b05e9b743c04a694ba7358b8af07ff1"></a>

## tls_fingerprint_values property — deny_list / 948fae6ddbac / 6

Type: `["list", "string"]`. Computed.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-c6263d93ae274e677216caec15bc8502d27b585682a19a013e1dfe1475240550"></a>

## Next pages — deny_list / 948fae6ddbac / 7

- [deny_list.asn_list](data-sources--service_policy--reference--group-001.md#canonical-057bdd37763b2ef7dc6bab3b36d56f8531a47c9eae337954f8e5afcf85f9cd67)
- [deny_list.asn_set](data-sources--service_policy--reference--group-001.md#canonical-fcd78ad770ae94979038ae81de2c1488ad0008304ea739a1cd7284c8bee74c60)
- [deny_list.default_action_allow](data-sources--service_policy--reference--group-001.md#canonical-d63f0da05fa71fd44d567e016d894e232c07d18200657212bb6d60dcead79bc2)
- [deny_list.default_action_deny](data-sources--service_policy--reference--group-001.md#canonical-d0b71679b2fd2cf85b18a04e305a03f6065754fe9290f53f33b921d39e1e8b59)
- [deny_list.default_action_next_policy](data-sources--service_policy--reference--group-001.md#canonical-f90b43e0925a1f1f3b7c4b1fbcacf43e9d65a49e9c5e555672419ee2bed9f7d5)
- [deny_list.ip_prefix_set](data-sources--service_policy--reference--group-001.md#canonical-0f7fcd65ad1acde210a778200b887ce4f6a823cb37e58edac5bc4310ef3f1801)
- [deny_list.prefix_list](data-sources--service_policy--reference--group-001.md#canonical-aa03658908e5fe2a886495417ce0876e8e00de3bc6f1489680312df4e8dac519)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-057bdd37763b2ef7dc6bab3b36d56f8531a47c9eae337954f8e5afcf85f9cd67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86b70e7764efc48d5ddd7cc70c4b197cc9fc93a5a9445afbcbe5c39b0bebd6fb"></a>

## deny_list.asn_list — deny_list.asn_list / 1f1aa8c6fb57 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-5f93ca70b2dc6b4430d0a8a8956f8fa491038318e823bd49dea1fb5e09f0bcc7)
- deny_list.asn_list

<a id="canonical-fea91b6264e848850710afa2a27a20dc59b51813571d3403e8efc857aba50190"></a>

Type: `"single"`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-5805c94df3126ea2b8e96ccd1494b245493c70931a457aad7f955e84f65c70a9"></a>

## Direct properties — deny_list.asn_list / 1f1aa8c6fb57 / 3

<a id="canonical-917db75472e86ea28691e1a4b00794be0e62a2a9d7991264f7f2bb99d2ba5fa0"></a>

<a id="canonical-f7338af7b2723e47c8a95efc35dd6152741d5f3d7b21cd3aae0962688075b17e"></a>

## as_numbers property — deny_list.asn_list / 1f1aa8c6fb57 / 4

Type: `["list", "number"]`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5f707aa350bb56987505473cdb6949887102b1ca0590301c5c2e39f2f503d6c8"></a>

## Next pages — deny_list.asn_list / 1f1aa8c6fb57 / 5

- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-5f93ca70b2dc6b4430d0a8a8956f8fa491038318e823bd49dea1fb5e09f0bcc7)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-fcd78ad770ae94979038ae81de2c1488ad0008304ea739a1cd7284c8bee74c60"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79a4ea4ca23da5013b7bf14bff9ad9d0ff710c32f9e60047458bb558edea0977"></a>

## deny_list.asn_set — deny_list.asn_set / acb93eef2947 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-5f93ca70b2dc6b4430d0a8a8956f8fa491038318e823bd49dea1fb5e09f0bcc7)
- deny_list.asn_set

<a id="canonical-6b85b60b5921fd09ca70bb916eb0b12bab848b47bbf990010ed32deffb132a2e"></a>

Type: `"list"`. Computed.

Addresses that belong to the ASNs in the given bgp\_asn\_set The ASN is obtained by performing a
lookup for the source IPv4 Address in a GeoIP DB.

Upstream description:

Addresses that belong to the ASNs in the given bgp\_asn\_set The ASN is obtained by performing a
lookup for the source IPv4 Address in a GeoIP DB.

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

<a id="canonical-1015d153ede25cc1115372fa1c2834aab0addea3a51e010576254f20893f7593"></a>

## Direct properties — deny_list.asn_set / acb93eef2947 / 3

<a id="canonical-df1d09ef5b0e04d81e31a1e82d8bb1f5baf37dd39842d1b416d929b5a1a93807"></a>

<a id="canonical-815b86a5f83300e0df155dec43cb72bd1be74438e5a9d44fff6c6a1efda06cbc"></a>

## name property — deny_list.asn_set / acb93eef2947 / 4

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

<a id="canonical-29f2e6c0e764acafc42acc689bb6a35221c1e1d57453a56d36873a85d07d5bf7"></a>

<a id="canonical-ac130cfc8a0d5d60983b375bed5e17a78c5b71cb7885d846d0bc3603b7570058"></a>

## namespace property — deny_list.asn_set / acb93eef2947 / 5

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

<a id="canonical-d7bbbf59c290809ba32647bc444bb75f383c7aabcf3237c47a132befbed68245"></a>

<a id="canonical-b13aaee33270a6ce3a9b9be3de3de43b092bbef84b8258514f84597d5dddbafc"></a>

## tenant property — deny_list.asn_set / acb93eef2947 / 6

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

<a id="canonical-ec7cdb1722030de745babe49489b1a6c500b651d58b14932a01f308ddaaad132"></a>

## Next pages — deny_list.asn_set / acb93eef2947 / 7

- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-5f93ca70b2dc6b4430d0a8a8956f8fa491038318e823bd49dea1fb5e09f0bcc7)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-d63f0da05fa71fd44d567e016d894e232c07d18200657212bb6d60dcead79bc2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d89dd819fb06ca153f730a086bfca8f38f53179685438174582960755c7015a5"></a>

## deny_list.default_action_allow — deny_list.default_action_allow / 99a3142d80f3 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-5f93ca70b2dc6b4430d0a8a8956f8fa491038318e823bd49dea1fb5e09f0bcc7)
- deny_list.default_action_allow

<a id="canonical-0aefabfab07f4dae0cd01607e6784f616b3158b5df876076579a989730a7d741"></a>

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

<a id="canonical-2bd4c6078e7342f05b10976e581569a6122c06b5773463cb169d350872a9ab78"></a>

## Direct properties — deny_list.default_action_allow / 99a3142d80f3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f1e4285a7b7f7fa87ee90308c02edb5f6d436f506c4a8df2d8ab8b7a44773780"></a>

## Next pages — deny_list.default_action_allow / 99a3142d80f3 / 4

- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-5f93ca70b2dc6b4430d0a8a8956f8fa491038318e823bd49dea1fb5e09f0bcc7)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-d0b71679b2fd2cf85b18a04e305a03f6065754fe9290f53f33b921d39e1e8b59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-892a04b3ce02353f0aedbd4cbb630bc37facdefc7ca8bfba11240537e28014f8"></a>

## deny_list.default_action_deny — deny_list.default_action_deny / d15431ef86e8 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-5f93ca70b2dc6b4430d0a8a8956f8fa491038318e823bd49dea1fb5e09f0bcc7)
- deny_list.default_action_deny

<a id="canonical-716255feb5dd14e2415ad2b593f2c72ca4de47e92e7a6c18d126e90f2ff66d6d"></a>

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

<a id="canonical-7332244ee23f3f17daf73d253790a8ba220e493503dd957cc9c4284f26b1207d"></a>

## Direct properties — deny_list.default_action_deny / d15431ef86e8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-50b9dbfe4a3b80867f035ecda31b3ab9fd76fe476e72287111a111a502cdb1ac"></a>

## Next pages — deny_list.default_action_deny / d15431ef86e8 / 4

- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-5f93ca70b2dc6b4430d0a8a8956f8fa491038318e823bd49dea1fb5e09f0bcc7)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-f90b43e0925a1f1f3b7c4b1fbcacf43e9d65a49e9c5e555672419ee2bed9f7d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43fb92d08a3b245e7c8dd381bc59861d67f550850151f88548409495c6db90af"></a>

## deny_list.default_action_next_policy — deny_list.default_action_next_policy / b744fee5ab34 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-5f93ca70b2dc6b4430d0a8a8956f8fa491038318e823bd49dea1fb5e09f0bcc7)
- deny_list.default_action_next_policy

<a id="canonical-b9784205bd7a0d02ad7fa8a486b98ad874a6d630b9500fe2ce2f8d8fe462fc84"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-31949655b374130a0c4d11ce92fafc65682929e33f8f8b7a3c898d759115113a"></a>

## Direct properties — deny_list.default_action_next_policy / b744fee5ab34 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-36b1b1acdeab3db6ddd15afb2a23880edddcb1ce86fb90ee0dcf9bc4dcb7f124"></a>

## Next pages — deny_list.default_action_next_policy / b744fee5ab34 / 4

- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-5f93ca70b2dc6b4430d0a8a8956f8fa491038318e823bd49dea1fb5e09f0bcc7)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-0f7fcd65ad1acde210a778200b887ce4f6a823cb37e58edac5bc4310ef3f1801"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b55b95989508aa883b6613383a48c5cb96b46607a038d7acaaa9c5a84884952a"></a>

## deny_list.ip_prefix_set — deny_list.ip_prefix_set / 5b82d40b931a / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-5f93ca70b2dc6b4430d0a8a8956f8fa491038318e823bd49dea1fb5e09f0bcc7)
- deny_list.ip_prefix_set

<a id="canonical-d2af1c619762e5c73f2557c424f689a6b61a8ef2b9b36cb34173956cb1e3d9e6"></a>

Type: `"list"`. Computed.

Addresses that are covered by the prefixes in the given ip\_prefix\_set.

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

<a id="canonical-754ffaf1284da43fa76cde94ac151cfbb97169bb975132d66c4a6707728c5520"></a>

## Direct properties — deny_list.ip_prefix_set / 5b82d40b931a / 3

<a id="canonical-a46a4c378239dc6bc27ff2343a2b6579ef80737322766d88278e051d64aa60c1"></a>

<a id="canonical-9c415f0f3d08c12f10f81e73957ed626a7fe15e77064a3365be1be2b7f07c044"></a>

## name property — deny_list.ip_prefix_set / 5b82d40b931a / 4

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

<a id="canonical-c8ea8f31466b4829c17db53b472900ba0701b70951078c927b3d2821fdcad8b5"></a>

<a id="canonical-96f338b70dc2018600cad8535b5eb65054c0fe8bb1453a119444455770f5804c"></a>

## namespace property — deny_list.ip_prefix_set / 5b82d40b931a / 5

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

<a id="canonical-82def6ae2945491139c48d7f0df873ac4dc45859eedfd41eb903c4577d12777f"></a>

<a id="canonical-3ae59471495f34f38353e8afaf7a16a366065e0f527e715db39169d3104b71f8"></a>

## tenant property — deny_list.ip_prefix_set / 5b82d40b931a / 6

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

<a id="canonical-29e8875d3e3cd91359df514e9bac118dc02041685a7316743e22f374fe572a0e"></a>

## Next pages — deny_list.ip_prefix_set / 5b82d40b931a / 7

- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-5f93ca70b2dc6b4430d0a8a8956f8fa491038318e823bd49dea1fb5e09f0bcc7)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-aa03658908e5fe2a886495417ce0876e8e00de3bc6f1489680312df4e8dac519"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45d144b0899c903e6fa14eeadb4cb8b220cd2c6ca360834add2b821f6fcab2fc"></a>

## deny_list.prefix_list — deny_list.prefix_list / bd2c44bb97bc / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-5f93ca70b2dc6b4430d0a8a8956f8fa491038318e823bd49dea1fb5e09f0bcc7)
- deny_list.prefix_list

<a id="canonical-1408a295a54b76825e032351ab7ccc55c2470397c43c064c7d501b89a7420806"></a>

Type: `"single"`. Computed.

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

<a id="canonical-fe6f3ec13b7e4524cde2e3bb40df674429fe27825a078a65b58cf5b2270198d7"></a>

## Direct properties — deny_list.prefix_list / bd2c44bb97bc / 3

<a id="canonical-c7653de422a32806cd2e29a50415d87a915cc427c077d5c4cf92b8bfd6305079"></a>

<a id="canonical-fc57d2860c9af9f725ab2473fd7f8a6ef5553c5db32d2c9c1e761cc0e6143ab0"></a>

## prefixes property — deny_list.prefix_list / bd2c44bb97bc / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-293d6dfa6a963a06536549c1c7e3d2c6ae5a107e2bce0c8476afc60902bd20a5"></a>

## Next pages — deny_list.prefix_list / bd2c44bb97bc / 5

- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-5f93ca70b2dc6b4430d0a8a8956f8fa491038318e823bd49dea1fb5e09f0bcc7)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-2027263afc4d5af4143efdbba5049cbc80defe477cca37d5d12eae77f21d788f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3f2225d3e12d37c132ec29f9f97b6a32ed9c47f0fc725f6cb9170135edc8ac8"></a>

## rule_list — rule_list / 952868a9fc59 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- rule_list

<a id="canonical-7b3eacb2d5a557b8f145def9c63d10ec64852bea12d3d445430e16caa2324c2e"></a>

Type: `"single"`. Computed.

Ordered service-policy rules for non-geographic predicates and actions. Do not use country\_list for
a geo-only rule here: the platform adds match-all any\_ip and any\_asn selectors on readback, so the
rule can match all traffic. Use deny\_list or allow\_list with country\_list for geographic source..

Upstream description:

Ordered service-policy rules for non-geographic predicates and actions. Do not use country\_list for
a geo-only rule here: the platform adds match-all any\_ip and any\_asn selectors on readback, so the
rule can match all traffic. Use deny\_list or allow\_list with country\_list for geographic source
matching.

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

<a id="canonical-3db5bda8ddbcdf052b59976d2d6313acb41f6ffe64394cca03674aaeeabdb8a8"></a>

## Direct properties — rule_list / 952868a9fc59 / 3

- [rules](data-sources--service_policy--reference--group-001.md#canonical-be906348dc6598ef9fd8b1a6709c80b9246a1561b7c4571e52429192f8fc4c99): complete subsection reference.

<a id="canonical-f13e73e2c2143f5d1e69d29ea8f34f5cdda2772f325c2fcbfd3a8a53cd114f62"></a>

## Next pages — rule_list / 952868a9fc59 / 4

- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-be906348dc6598ef9fd8b1a6709c80b9246a1561b7c4571e52429192f8fc4c99)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-be906348dc6598ef9fd8b1a6709c80b9246a1561b7c4571e52429192f8fc4c99"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b30ab8bcf42f0eecece5fbaed1350d1c0afbbb0e94153001519dc561e2317f6"></a>

## rule_list.rules — rule_list.rules / 35f4bf2d7861 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-2027263afc4d5af4143efdbba5049cbc80defe477cca37d5d12eae77f21d788f)
- rule_list.rules

<a id="canonical-916cf59ab2933f7b1791921bdc25d0cae463186d742a741ed01cd3581bbaca8f"></a>

Type: `"list"`. Computed.

Define the list of rules (with an order) that should be evaluated by this service policy. Rules are
evaluated from top to bottom in the list.

Upstream description:

Define the list of rules (with an order) that should be evaluated by this service policy. Rules are
evaluated from top to bottom in the list.

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
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-4742dce81756307bfce586fabc4e1255373cb17751d353f66857fda59071d2d5"></a>

## Direct properties — rule_list.rules / 35f4bf2d7861 / 3

- [metadata](data-sources--service_policy--reference--group-001.md#canonical-3640cadb11307b6b963e2a530faf027516d05018a047edab294957552510daad): complete subsection reference.

- [spec](data-sources--service_policy--reference--group-001.md#canonical-777c55186289618a36c06244acc1e85f5bfef98a9091f6fb8cff026a94d25264): complete subsection reference.

<a id="canonical-eb47600e428885d64cce653b92258cc84909d7f785db7b70f5f65ee540a9ab14"></a>

## Next pages — rule_list.rules / 35f4bf2d7861 / 4

- [rule_list.rules.metadata](data-sources--service_policy--reference--group-001.md#canonical-3640cadb11307b6b963e2a530faf027516d05018a047edab294957552510daad)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-777c55186289618a36c06244acc1e85f5bfef98a9091f6fb8cff026a94d25264)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-2027263afc4d5af4143efdbba5049cbc80defe477cca37d5d12eae77f21d788f)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-3640cadb11307b6b963e2a530faf027516d05018a047edab294957552510daad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19e52442b6372ab5b6d889b003b22a78c446554f81fb605892afba35917610fd"></a>

## rule_list.rules.metadata — rule_list.rules.metadata / c3999960722e / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-2027263afc4d5af4143efdbba5049cbc80defe477cca37d5d12eae77f21d788f)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-be906348dc6598ef9fd8b1a6709c80b9246a1561b7c4571e52429192f8fc4c99)
- rule_list.rules.metadata

<a id="canonical-9e6a91536f54f310ed4f7834208bf55945e4cb8bdece6094153d7015879a99f8"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-8c7938fc38d65f82c8ebacaedf4826433e08248071778df249fdf33afc1621d3"></a>

## Direct properties — rule_list.rules.metadata / c3999960722e / 3

<a id="canonical-26a6ed47caaa91ab4b30fc30c8b15fd97643fd08d0be66e5d9fb0694638eac55"></a>

<a id="canonical-0f23e3ae14bab47ea39562a59c9d4ac13275f0c7ad7d356935bd214826623356"></a>

## description_spec property — rule_list.rules.metadata / c3999960722e / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-4079e765d4401d78b28345c49699c6fceca5a0b09deb318e13b9dbee2fb59ee4"></a>

<a id="canonical-ddfbfcb5f1246484a5c0a53b3724902ed34832797c05ff95df06ac11c903d390"></a>

## name property — rule_list.rules.metadata / c3999960722e / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-0bfd7b5bf79c3322fef74e3c4466ed04b35f08e2950bc0c63b61c68c5ac0bb70"></a>

## Next pages — rule_list.rules.metadata / c3999960722e / 6

- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-be906348dc6598ef9fd8b1a6709c80b9246a1561b7c4571e52429192f8fc4c99)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-777c55186289618a36c06244acc1e85f5bfef98a9091f6fb8cff026a94d25264"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9896c225b60851151a79f06f67f92418e78540f8a5887e994cb7c9fd0e33e506"></a>

## rule_list.rules.spec — rule_list.rules.spec / 947d374dfdf2 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-2027263afc4d5af4143efdbba5049cbc80defe477cca37d5d12eae77f21d788f)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-be906348dc6598ef9fd8b1a6709c80b9246a1561b7c4571e52429192f8fc4c99)
- rule_list.rules.spec

<a id="canonical-5f1d894f385e18178ccc42e2a8d64de00aa54e43320b6dc73bd05ee5119deecb"></a>

Type: `"single"`. Computed.

Shape of service\_policy\_rule in the storage backend.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_name\",\"client_name_matcher\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-dst_asn_choice": "[]",
  "x-ves-oneof-field-dst_ip_choice": "[]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]",
  "x-ves-oneof-field-tls_fingerprint_choice": "[\"ja4_tls_fingerprint\",\"tls_fingerprint_matcher\"]"
}
```

<a id="canonical-eab0df0ea0ff268555fca43d928f8030088a222cbb2f39084f655c8c69994ee0"></a>

## Direct properties — rule_list.rules.spec / 947d374dfdf2 / 3

<a id="canonical-0cd3937364c0103b3ce598be37c02c4c125f3ec669600e8409eaba936108731a"></a>

<a id="canonical-6e2b9e0c7fa0d6166c02096032fa82c7f0c3ee59230091954fc97f08c9950974"></a>

## action property — rule_list.rules.spec / 947d374dfdf2 / 4

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW|NEXT\_POLICY\] The rule action determines the disposition of the input request
API. If a policy matches a rule with an ALLOW action, the processing of the request proceeds
forward. If it matches a rule with a DENY action, the processing of the request is terminated and an
appropriate message/code returned to.. Possible values are \`DENY\`, \`ALLOW\`, \`NEXT\_POLICY\`.
Defaults to \`DENY\`.

Upstream description:

The rule action determines the disposition of the input request API. If a policy matches a rule with
an ALLOW action, the processing of the request proceeds forward. If it matches a rule with a DENY
action, the processing of the request is terminated and an appropriate message/code returned to the
originator. If it matches a rule with a NEXT\_POLICY\_SET action, evaluation of the current policy
set terminates and evaluation of the next policy set in the chain begins.

&#8203;- DENY: DENY

Deny the request. &#8203;- ALLOW: ALLOW

Allow the request to proceed. &#8203;- NEXT\_POLICY\_SET: NEXT\_POLICY\_SET

Terminate evaluation of the current policy set and begin evaluating the next policy set in the
chain. Note that the evaluation of any remaining policies in the current policy set is skipped.
&#8203;- NEXT\_POLICY: NEXT\_POLICY

Terminate evaluation of the current policy and begin evaluating the next policy in the policy set.
Note that the evaluation of any remaining rules in the current policy is skipped. &#8203;-
LAST\_POLICY: LAST\_POLICY

Terminate evaluation of the current policy and begin evaluating the last policy in the policy set.
Note that the evaluation of any remaining rules in the current policy is skipped. &#8203;-
GOTO\_POLICY: GOTO\_POLICY

Terminate evaluation of the current policy and begin evaluating a specific policy in the policy set.
The policy is specified using the goto\_policy field in the rule and must be after the current
policy in the policy set.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW",
    "NEXT_POLICY"
  ],
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  }
}
```

- [any_asn](data-sources--service_policy--reference--group-001.md#canonical-d5b5e248130f0e306f909f0939b094d305bff0a4a77fcb2a597a4292d4305b8a): complete subsection reference.

- [any_client](data-sources--service_policy--reference--group-001.md#canonical-5944d2668619b378076e640d9113a68841b88cdf2cf418419f7009d7dd913b11): complete subsection reference.

- [any_ip](data-sources--service_policy--reference--group-002.md#canonical-fa167b642530c40398b5c03fd4918204bc8a59f7bbd2696888fe9e791ab08329): complete subsection reference.

- [api_group_matcher](data-sources--service_policy--reference--group-002.md#canonical-00f05c9666db74a329f4418c6f51d5ee1e38141f6f2039826ea34c4dd71b0b3f): complete subsection reference.

- [arg_matchers](data-sources--service_policy--reference--group-002.md#canonical-3a79819c72ce229cff8a7786fea86a875ab12318f3610dc0ddb53021c8bf43e1): complete subsection reference.

- [asn_list](data-sources--service_policy--reference--group-002.md#canonical-0350902cdb8d270fb73759b0e61f43b24cb604f88ef2a17bf846905051cb6fcf): complete subsection reference.

- [asn_matcher](data-sources--service_policy--reference--group-002.md#canonical-c39ace095c594a15e6831d1e1b579d17e70747788d37cb635e2e9201c26e52c9): complete subsection reference.

- [body_matcher](data-sources--service_policy--reference--group-002.md#canonical-d80fcebf073da5e74dcc214719551bea819643ffb7a1ee1105a881c60aed15fa): complete subsection reference.

- [bot_action](data-sources--service_policy--reference--group-002.md#canonical-7674d8063baae9b5149e0a5add535665425da7a54e5a854cd015bab6ae9d5a46): complete subsection reference.

<a id="canonical-1a6af37560ccdcef815d85e58964d442037955f00596925434b433a5e2698077"></a>

<a id="canonical-8ced578094985fb6a918d5c7577f88e673a94bebb6807b0d51c1ba9c1bf13998"></a>

## client_name property — rule_list.rules.spec / 947d374dfdf2 / 5

Type: `"string"`. Computed.

Exclusive with \[any\_client client\_name\_matcher client\_selector ip\_threat\_category\_list\] The
expected name of the client invoking the request API. The predicate evaluates to true if any of the
actual names is the same as the expected client name.

Upstream description:

Exclusive with \[any\_client client\_name\_matcher client\_selector ip\_threat\_category\_list\] The
expected name of the client invoking the request API. The predicate evaluates to true if any of the
actual names is the same as the expected client name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [client_name_matcher](data-sources--service_policy--reference--group-002.md#canonical-1976f213ad75f2f7096c4856fa43126f20ccec02508f946ce1cfa40de97e0cf3): complete subsection reference.

- [client_selector](data-sources--service_policy--reference--group-002.md#canonical-6ce4cff99f881b4a246164f6dc93d4c91b7fd99774dd1be5b0f3488d0e0b63bc): complete subsection reference.

- [cookie_matchers](data-sources--service_policy--reference--group-002.md#canonical-95c2d8b38cc51aebcefd80aa80cb5873c6217b096bc09a21f20976dae7507fdc): complete subsection reference.

- [domain_matcher](data-sources--service_policy--reference--group-002.md#canonical-fa8cff4bafa32368b1b5325e885de92b812372f828d5adb719e2a0f39d2437e7): complete subsection reference.

<a id="canonical-a9159a904b96f396792f61be3806afbcc8b057ab625626d3d990f4c8cd128640"></a>

<a id="canonical-2e4ccac1bb3b9445f0b64412621d968652f4801632783c68fb14b9c4adbbaef5"></a>

## expiration_timestamp property — rule_list.rules.spec / 947d374dfdf2 / 6

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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

- [headers](data-sources--service_policy--reference--group-002.md#canonical-8a6b64b8c9b2e0f3710f12c5a664dcb9dbd7381008f61af1d3bafd66cf85e55a): complete subsection reference.

- [http_method](data-sources--service_policy--reference--group-002.md#canonical-2b735edd3340fface35788e7d2f76205c7583f799aec50d9b772ff2f13704ce2): complete subsection reference.

- [ip_matcher](data-sources--service_policy--reference--group-002.md#canonical-c1d7746015dd29ed3c477bca2406bb1baef73803a77965bdaf2af22ad61d2401): complete subsection reference.

- [ip_prefix_list](data-sources--service_policy--reference--group-002.md#canonical-5f27ab6693ad059ea65f9fe317bfa204039f8af5d641f9c9fb146fb65b06ec79): complete subsection reference.

- [ip_threat_category_list](data-sources--service_policy--reference--group-002.md#canonical-27069f863fb08313609747652defd4e3b4d3ef3cf7d8ad86723d81bd099b4ab4): complete subsection reference.

- [ja4_tls_fingerprint](data-sources--service_policy--reference--group-002.md#canonical-082c793e617110d3435b9be698935e1e39f0f07e936d655647e46a0f4ed516b3): complete subsection reference.

- [jwt_claims](data-sources--service_policy--reference--group-002.md#canonical-32fa0b0c145418d2fe3e39f94757b06cd94d7a972fb3cbbdca903c90db538f0d): complete subsection reference.

- [label_matcher](data-sources--service_policy--reference--group-002.md#canonical-69c99fd683ae22677c123001d177a19c1030f66898f2c889f4640752754c993d): complete subsection reference.

<a id="canonical-d032f9e29d9790dc374806d3f4920adef7bfaf9da17715c2b0812f395ea1a66b"></a>

<a id="canonical-bd7185af062adf7d9cd658cfa6dad080de26dac700d83ae4ecf1c828551bcc77"></a>

## log_rule_evaluation property — rule_list.rules.spec / 947d374dfdf2 / 7

Type: `"bool"`. Computed.

Log the rule match details along with the request and continue to evaluate rules in the sequence.

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

- [mum_action](data-sources--service_policy--reference--group-002.md#canonical-9a4ea6d5346e960fdab01889edfcff93d1f4b58b2388eaab5bac60c653ade17a): complete subsection reference.

- [path](data-sources--service_policy--reference--group-002.md#canonical-38934814e0befe48e2d8318ee92b6309e75e581d398e16e171fef4e43771a41f): complete subsection reference.

- [port_matcher](data-sources--service_policy--reference--group-002.md#canonical-025b99fdb1f4335291706f440472dc66150c3acfdd2ae229e0450f7c8fac845e): complete subsection reference.

- [query_params](data-sources--service_policy--reference--group-002.md#canonical-a398fe1a59f442c171bf9237500b2a38f7d0a0b8284b035129575a63825795c3): complete subsection reference.

- [request_constraints](data-sources--service_policy--reference--group-002.md#canonical-9427d4dccbd993f6adcac6255385f32a2929c045c21d361e7994d340382a6e87): complete subsection reference.

- [segment_policy](data-sources--service_policy--reference--group-003.md#canonical-19eaef05818327e7086de986172355b621e45cc4724d90bf7b986bf464b4d342): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--service_policy--reference--group-003.md#canonical-7f0c83d3b337e7f751fc72dcb744306915b9fc1e858b630046b6edddbfa12272): complete subsection reference.

- [user_identity_matcher](data-sources--service_policy--reference--group-003.md#canonical-f2bb82b371ed002eaaae235a9ec59710e65df3302bb58b10d190ccf6c3fc2cf5): complete subsection reference.

- [waf_action](data-sources--service_policy--reference--group-003.md#canonical-79308c6e3b7f42024dda189ab7f735a4d15ebbbe3aa4d1bbe193b51f990b1a10): complete subsection reference.

<a id="canonical-9fb3e184b1001600dd28be78e67ab767a4a5f2ebe1205a63845b234a7e618bad"></a>

## Next pages — rule_list.rules.spec / 947d374dfdf2 / 8

- [rule_list.rules.spec.any_asn](data-sources--service_policy--reference--group-001.md#canonical-d5b5e248130f0e306f909f0939b094d305bff0a4a77fcb2a597a4292d4305b8a)
- [rule_list.rules.spec.any_client](data-sources--service_policy--reference--group-001.md#canonical-5944d2668619b378076e640d9113a68841b88cdf2cf418419f7009d7dd913b11)
- [rule_list.rules.spec.any_ip](data-sources--service_policy--reference--group-002.md#canonical-fa167b642530c40398b5c03fd4918204bc8a59f7bbd2696888fe9e791ab08329)
- [rule_list.rules.spec.api_group_matcher](data-sources--service_policy--reference--group-002.md#canonical-00f05c9666db74a329f4418c6f51d5ee1e38141f6f2039826ea34c4dd71b0b3f)
- [rule_list.rules.spec.arg_matchers](data-sources--service_policy--reference--group-002.md#canonical-3a79819c72ce229cff8a7786fea86a875ab12318f3610dc0ddb53021c8bf43e1)
- [rule_list.rules.spec.asn_list](data-sources--service_policy--reference--group-002.md#canonical-0350902cdb8d270fb73759b0e61f43b24cb604f88ef2a17bf846905051cb6fcf)
- [rule_list.rules.spec.asn_matcher](data-sources--service_policy--reference--group-002.md#canonical-c39ace095c594a15e6831d1e1b579d17e70747788d37cb635e2e9201c26e52c9)
- [rule_list.rules.spec.body_matcher](data-sources--service_policy--reference--group-002.md#canonical-d80fcebf073da5e74dcc214719551bea819643ffb7a1ee1105a881c60aed15fa)
- [rule_list.rules.spec.bot_action](data-sources--service_policy--reference--group-002.md#canonical-7674d8063baae9b5149e0a5add535665425da7a54e5a854cd015bab6ae9d5a46)
- [rule_list.rules.spec.client_name_matcher](data-sources--service_policy--reference--group-002.md#canonical-1976f213ad75f2f7096c4856fa43126f20ccec02508f946ce1cfa40de97e0cf3)
- [rule_list.rules.spec.client_selector](data-sources--service_policy--reference--group-002.md#canonical-6ce4cff99f881b4a246164f6dc93d4c91b7fd99774dd1be5b0f3488d0e0b63bc)
- [rule_list.rules.spec.cookie_matchers](data-sources--service_policy--reference--group-002.md#canonical-95c2d8b38cc51aebcefd80aa80cb5873c6217b096bc09a21f20976dae7507fdc)
- [rule_list.rules.spec.domain_matcher](data-sources--service_policy--reference--group-002.md#canonical-fa8cff4bafa32368b1b5325e885de92b812372f828d5adb719e2a0f39d2437e7)
- [rule_list.rules.spec.headers](data-sources--service_policy--reference--group-002.md#canonical-8a6b64b8c9b2e0f3710f12c5a664dcb9dbd7381008f61af1d3bafd66cf85e55a)
- [rule_list.rules.spec.http_method](data-sources--service_policy--reference--group-002.md#canonical-2b735edd3340fface35788e7d2f76205c7583f799aec50d9b772ff2f13704ce2)
- [rule_list.rules.spec.ip_matcher](data-sources--service_policy--reference--group-002.md#canonical-c1d7746015dd29ed3c477bca2406bb1baef73803a77965bdaf2af22ad61d2401)
- [rule_list.rules.spec.ip_prefix_list](data-sources--service_policy--reference--group-002.md#canonical-5f27ab6693ad059ea65f9fe317bfa204039f8af5d641f9c9fb146fb65b06ec79)
- [rule_list.rules.spec.ip_threat_category_list](data-sources--service_policy--reference--group-002.md#canonical-27069f863fb08313609747652defd4e3b4d3ef3cf7d8ad86723d81bd099b4ab4)
- [rule_list.rules.spec.ja4_tls_fingerprint](data-sources--service_policy--reference--group-002.md#canonical-082c793e617110d3435b9be698935e1e39f0f07e936d655647e46a0f4ed516b3)
- [rule_list.rules.spec.jwt_claims](data-sources--service_policy--reference--group-002.md#canonical-32fa0b0c145418d2fe3e39f94757b06cd94d7a972fb3cbbdca903c90db538f0d)
- [rule_list.rules.spec.label_matcher](data-sources--service_policy--reference--group-002.md#canonical-69c99fd683ae22677c123001d177a19c1030f66898f2c889f4640752754c993d)
- [rule_list.rules.spec.mum_action](data-sources--service_policy--reference--group-002.md#canonical-9a4ea6d5346e960fdab01889edfcff93d1f4b58b2388eaab5bac60c653ade17a)
- [rule_list.rules.spec.path](data-sources--service_policy--reference--group-002.md#canonical-38934814e0befe48e2d8318ee92b6309e75e581d398e16e171fef4e43771a41f)
- [rule_list.rules.spec.port_matcher](data-sources--service_policy--reference--group-002.md#canonical-025b99fdb1f4335291706f440472dc66150c3acfdd2ae229e0450f7c8fac845e)
- [rule_list.rules.spec.query_params](data-sources--service_policy--reference--group-002.md#canonical-a398fe1a59f442c171bf9237500b2a38f7d0a0b8284b035129575a63825795c3)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-9427d4dccbd993f6adcac6255385f32a2929c045c21d361e7994d340382a6e87)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-003.md#canonical-19eaef05818327e7086de986172355b621e45cc4724d90bf7b986bf464b4d342)
- [rule_list.rules.spec.tls_fingerprint_matcher](data-sources--service_policy--reference--group-003.md#canonical-7f0c83d3b337e7f751fc72dcb744306915b9fc1e858b630046b6edddbfa12272)
- [rule_list.rules.spec.user_identity_matcher](data-sources--service_policy--reference--group-003.md#canonical-f2bb82b371ed002eaaae235a9ec59710e65df3302bb58b10d190ccf6c3fc2cf5)
- [rule_list.rules.spec.waf_action](data-sources--service_policy--reference--group-003.md#canonical-79308c6e3b7f42024dda189ab7f735a4d15ebbbe3aa4d1bbe193b51f990b1a10)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-be906348dc6598ef9fd8b1a6709c80b9246a1561b7c4571e52429192f8fc4c99)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-d5b5e248130f0e306f909f0939b094d305bff0a4a77fcb2a597a4292d4305b8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-709567317c92d0d4764eb6dcb1c786778c974779b947acf52a981e28d5f01171"></a>

## rule_list.rules.spec.any_asn — rule_list.rules.spec.any_asn / 9c81eacaa6ea / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-2027263afc4d5af4143efdbba5049cbc80defe477cca37d5d12eae77f21d788f)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-be906348dc6598ef9fd8b1a6709c80b9246a1561b7c4571e52429192f8fc4c99)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-777c55186289618a36c06244acc1e85f5bfef98a9091f6fb8cff026a94d25264)
- rule_list.rules.spec.any_asn

<a id="canonical-48946e969cad3dec694accb2d90dd2dfade21a9470e2d1e8f7e933d89e7a8951"></a>

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

<a id="canonical-e2998f1f9a3aaf7c77db055a10b66b77fd913ff45a858dca3823a2b20eca359c"></a>

## Direct properties — rule_list.rules.spec.any_asn / 9c81eacaa6ea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2a1753a433d3e84ebbce730052ad5a74a61f30cbdadd136d3e46648f59561a41"></a>

## Next pages — rule_list.rules.spec.any_asn / 9c81eacaa6ea / 4

- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-777c55186289618a36c06244acc1e85f5bfef98a9091f6fb8cff026a94d25264)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-5944d2668619b378076e640d9113a68841b88cdf2cf418419f7009d7dd913b11"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
