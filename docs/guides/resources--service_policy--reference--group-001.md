---
page_title: "xcsh_service_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_service_policy reference."
---

# xcsh_service_policy reference

<a id="canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-561fc0f09f27a94ab99296fca1ce3155ede4b869c0fd538e28a6d50e9a17f351"></a>

## Property reference — Property reference / c3b918ccbafa / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- Property reference

<a id="canonical-93d2e41554e816c3746779b6dc3023622a55d6d53ebf7770d009b7acc6b4f328"></a>

## Direct properties — Property reference / c3b918ccbafa / 3

- [allow_all_requests](resources--service_policy--reference--group-001.md#canonical-3b69e6598e66a1a089c034f8a0c09638e768e7c7463369245d991c9c5df2e4e5): complete subsection reference.

- [allow_list](resources--service_policy--reference--group-001.md#canonical-93c35c3efe81a429730c6d3e71a61214955b4938adda2886bd96ced391c75992): complete subsection reference.

<a id="canonical-c1edc06afee950e850c10672838909f073007dcab8be61460be319e3ca696ff8"></a>

<a id="canonical-2b6401cd38cedbf2743414161ed0f2ccda9c1896eed22c9e69cf6beac5819644"></a>

## annotations property — Property reference / c3b918ccbafa / 4

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

- [any_server](resources--service_policy--reference--group-001.md#canonical-b6e1a4a11f301b65c95f28a92ddfe69d7d2cdfa290875a3fcf88e3dab50776a7): complete subsection reference.

- [deny_all_requests](resources--service_policy--reference--group-001.md#canonical-e9aaa06e2b0e97cd5cd7e20493c62d2557c9005231f3c38641984235f0dc1cc6): complete subsection reference.

- [deny_list](resources--service_policy--reference--group-001.md#canonical-f0b8d4e9de4ee80d1e3ff1ca4e51e89a004ef1276e52cc19138e08d501d46ee4): complete subsection reference.

<a id="canonical-72d66a71744f262e9ea79466975d96331079a246f6779d42f5b85006985f4429"></a>

<a id="canonical-cdeb9501407b6a7fd742c881aa0fb8492d42b98c0374ba232da24458cb9afff1"></a>

## description property — Property reference / c3b918ccbafa / 5

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

<a id="canonical-5f66ba5b580a28a9025afb5058cf14c8559313dd5e738bb1faaea5b2327167fb"></a>

<a id="canonical-60732bbe05c7c29b27568307aa8adc8a1df15aa2c2cf05c8d33a1c591275bea6"></a>

## disable property — Property reference / c3b918ccbafa / 6

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

<a id="canonical-48464a5d82ee61c357ed5cf342b7ab5454827bcba632849ab85f4e88fa48873b"></a>

<a id="canonical-d54c0ff8e9078190103b9eb79aad8e4529fc60966246b8c5c090e99f90fdb51e"></a>

## id property — Property reference / c3b918ccbafa / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-dec97802c3751df400c39b1cd20d6e73d1b81182116bfd2856330a1456a86d3c"></a>

<a id="canonical-2949ffad6b6a7fd12d5517747750c513aab0fa9cbcf0460926e71cd1c3a66f30"></a>

## labels property — Property reference / c3b918ccbafa / 8

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

<a id="canonical-764987d7401b0171a6ea5577489b89b6a624e90ad3def2506972cccba7d42253"></a>

<a id="canonical-fecd2f25bab5339142f557dee9eb873407b5d2909d516f460dda72c758dde80b"></a>

## name property — Property reference / c3b918ccbafa / 9

Type: `"string"`. Required.

Name of the Service Policy. Must be unique within the namespace.

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

<a id="canonical-f4245eea2baea0ffeab0fe7ab9e6ebdef2d96b10c3840232ba98316430cb739d"></a>

<a id="canonical-713dc8b11a058f5d5e3e4f67fb149b7f0ec7b90418329d1fc284782e4af3a844"></a>

## namespace property — Property reference / c3b918ccbafa / 10

Type: `"string"`. Required.

Namespace where the Service Policy is created.

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

- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490): complete subsection reference.

<a id="canonical-1fac4dfe5494bdfc04e8af123a2f3b82ee09c42dce1cde9971a278281150ff5b"></a>

<a id="canonical-ec6a7b8a1afacc03b7a54c8ccf4b5d67c3ea26a7cea13a6c4a68e52dd68c1d77"></a>

## server_name property — Property reference / c3b918ccbafa / 11

Type: `"string"`. Optional, Computed.

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

- [server_name_matcher](resources--service_policy--reference--group-003.md#canonical-de3685065567f4411ab04f86f8f14f9fc4e21e79652d75a6706a62e3882309c7): complete subsection reference.

- [server_selector](resources--service_policy--reference--group-003.md#canonical-ce09b78b2268dfdf225f8fa7808ee3dc2ee2eaa1ca7f4aef27737f6280575fab): complete subsection reference.

- [timeouts](resources--service_policy--reference--group-003.md#canonical-4c2b79645b86fc0db1b5a5e8348eb1539ad05a418b46eed14d7fd13168c502a6): complete subsection reference.

<a id="canonical-e89eb54d086e10a90d40c37a164b3f997495f79d838c854b4f394d74fa63ab76"></a>

## All schema paths — Property reference / c3b918ccbafa / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all_requests` | [allow_all_requests](resources--service_policy--reference--group-001.md#canonical-e3cb629e4f7ac08c9a96efb83d0e2b8171553728fd0cf8672d524b3505cb9857) |
| `allow_list` | [allow_list](resources--service_policy--reference--group-001.md#canonical-4950530880ce5501a5e517c1ba34304e28870c9dee42078c60b005114342a10b) |
| `allow_list.asn_list` | [allow_list.asn_list](resources--service_policy--reference--group-001.md#canonical-41493156507a550fc9d9c96c2e93f72015b3c652afaea794454c63ae94d87b9c) |
| `allow_list.asn_list.as_numbers` | [allow_list.asn_list.as_numbers](resources--service_policy--reference--group-001.md#canonical-024568803759fb55ea6755131b55fdb5d8752ec17df5f5b0b0f9dfd55b41e408) |
| `allow_list.asn_set` | [allow_list.asn_set](resources--service_policy--reference--group-001.md#canonical-615f9e756cde13d365f1257916860eb471ff83df7e96f4cb6190d2e819b51ffb) |
| `allow_list.asn_set.name` | [allow_list.asn_set.name](resources--service_policy--reference--group-001.md#canonical-925a25724e59e4d20445b4aee5385063a4651727652c7fa44981095cd59dd937) |
| `allow_list.asn_set.namespace` | [allow_list.asn_set.namespace](resources--service_policy--reference--group-001.md#canonical-492efe7890f870bce94d6caddc00f92de00329efc097eac5abb9ec4ca624a00d) |
| `allow_list.asn_set.tenant` | [allow_list.asn_set.tenant](resources--service_policy--reference--group-001.md#canonical-43587761c1f3b83ce560f7b688394ffb0e8c8228c16a2667e9f4e6ce708138da) |
| `allow_list.country_list` | [allow_list.country_list](resources--service_policy--reference--group-001.md#canonical-f03508c854530683649e09917740c18e93c3b61a0729203b332cc0d105be53cb) |
| `allow_list.default_action_allow` | [allow_list.default_action_allow](resources--service_policy--reference--group-001.md#canonical-0ea008c5cfc7ae353905c8a4af9937ed275b87a8fd71c9ae37683f00bbb8b318) |
| `allow_list.default_action_deny` | [allow_list.default_action_deny](resources--service_policy--reference--group-001.md#canonical-e03f0b6d8d42593f11c329fbd5ec6be3d61a6f67df97f141f72b750d1c111c49) |
| `allow_list.default_action_next_policy` | [allow_list.default_action_next_policy](resources--service_policy--reference--group-001.md#canonical-3c23d353b680ee8e7f53bfd8e37bc0d22da142f8ecff1447f12e145602172bd0) |
| `allow_list.ip_prefix_set` | [allow_list.ip_prefix_set](resources--service_policy--reference--group-001.md#canonical-c904dc6c2fe8a0ee24d82b9be9f510813f72d4d3df7aee0de04b8ebc319c0002) |
| `allow_list.ip_prefix_set.name` | [allow_list.ip_prefix_set.name](resources--service_policy--reference--group-001.md#canonical-1423d8d94860249e7263e9a6684d7281695c880dc90cdc3098726ca5f78a7961) |
| `allow_list.ip_prefix_set.namespace` | [allow_list.ip_prefix_set.namespace](resources--service_policy--reference--group-001.md#canonical-bdeaa6d178bb53198b86d068653270360186edffdb59c2c405d25179737d6575) |
| `allow_list.ip_prefix_set.tenant` | [allow_list.ip_prefix_set.tenant](resources--service_policy--reference--group-001.md#canonical-7152d6b1dd97b55887e5041b046525c743ef72f293cc44277889fd83547eca7f) |
| `allow_list.prefix_list` | [allow_list.prefix_list](resources--service_policy--reference--group-001.md#canonical-0f17e5790c628158ccadd71761471fab46dc21fad5332c9b6cfea9ee5dd2fc75) |
| `allow_list.prefix_list.prefixes` | [allow_list.prefix_list.prefixes](resources--service_policy--reference--group-001.md#canonical-a2338325f17dbc0419757bf57251081a19584cde986aa2c77afead476da46cb5) |
| `allow_list.tls_fingerprint_classes` | [allow_list.tls_fingerprint_classes](resources--service_policy--reference--group-001.md#canonical-596bed29b44b3af5f8190b403daf45f1ba694477f1eac13dd2dd226fb1cd17b4) |
| `allow_list.tls_fingerprint_values` | [allow_list.tls_fingerprint_values](resources--service_policy--reference--group-001.md#canonical-cdb161d4e772bd9f064300914cce8dd0b1d1975eb309d54224b9452235ecd677) |
| `annotations` | [annotations](resources--service_policy--reference--group-001.md#canonical-c1edc06afee950e850c10672838909f073007dcab8be61460be319e3ca696ff8) |
| `any_server` | [any_server](resources--service_policy--reference--group-001.md#canonical-442b7a0e07e9bc7dfffeed2add1ded6a367c80ad7fb76a6d890162b2aa05aada) |
| `deny_all_requests` | [deny_all_requests](resources--service_policy--reference--group-001.md#canonical-c0e726e216e23a86cf1e47c6f766758ef08efc2f4daab84f6f0a325181a1bc46) |
| `deny_list` | [deny_list](resources--service_policy--reference--group-001.md#canonical-9ce9bd163210878afb5cbb7de4ba16569a54be058aab14ee1c38a5ba7bc5341a) |
| `deny_list.asn_list` | [deny_list.asn_list](resources--service_policy--reference--group-001.md#canonical-6423e6d7e05b1f2d5cb31c48fbd8f2fa7b8ab41b6f1aa6023c76268b84c534d2) |
| `deny_list.asn_list.as_numbers` | [deny_list.asn_list.as_numbers](resources--service_policy--reference--group-001.md#canonical-142ad56b4ab090426bb0415da3e290dcc30e98efb9eca7d56e9947edf16ba1b3) |
| `deny_list.asn_set` | [deny_list.asn_set](resources--service_policy--reference--group-001.md#canonical-6cbf23466b4e9132f4e56a9150524b8940afc6978410f74075dd324fa69d878a) |
| `deny_list.asn_set.name` | [deny_list.asn_set.name](resources--service_policy--reference--group-001.md#canonical-760789468fd821c70cb6d3c636759a276fbd476df2429933f98a4feaa88760e0) |
| `deny_list.asn_set.namespace` | [deny_list.asn_set.namespace](resources--service_policy--reference--group-001.md#canonical-dccfd0f7b772160df141a154e57e61a752e4724ec5867cad295c6009374f0403) |
| `deny_list.asn_set.tenant` | [deny_list.asn_set.tenant](resources--service_policy--reference--group-001.md#canonical-59ae525382943fe22ed9abae8afdc5aa149f30454edd70a4cb2e914db195c2a4) |
| `deny_list.country_list` | [deny_list.country_list](resources--service_policy--reference--group-001.md#canonical-be2bf781e75005c7ce29f2d78f21f4ec0c56a0a9c9629636c7f502ee9c75c9c7) |
| `deny_list.default_action_allow` | [deny_list.default_action_allow](resources--service_policy--reference--group-001.md#canonical-5b6d88227f6af615b007424e6788b2fd9cec751f694a855199e023e1b85993e2) |
| `deny_list.default_action_deny` | [deny_list.default_action_deny](resources--service_policy--reference--group-001.md#canonical-a58db399c83c51e340623a595c56dbd7cf78e316e222f153d396f73e27985af4) |
| `deny_list.default_action_next_policy` | [deny_list.default_action_next_policy](resources--service_policy--reference--group-001.md#canonical-e56ddbd1b81330c100843d4b0c8a4b7123bb40f3be202f7b60f4a6d8ea55f742) |
| `deny_list.ip_prefix_set` | [deny_list.ip_prefix_set](resources--service_policy--reference--group-001.md#canonical-a23b4c23fce6e6217c17bc6d709efa79691811e37064bacbcb73f643354bdba0) |
| `deny_list.ip_prefix_set.name` | [deny_list.ip_prefix_set.name](resources--service_policy--reference--group-001.md#canonical-868e1c9f903622234134659119c8a8ecd6b0027501db2a52239aeec419fd1415) |
| `deny_list.ip_prefix_set.namespace` | [deny_list.ip_prefix_set.namespace](resources--service_policy--reference--group-001.md#canonical-a639d0441b90e44c997299e21fff249e684243973ff6cb1ba3c60a05ee41f1c0) |
| `deny_list.ip_prefix_set.tenant` | [deny_list.ip_prefix_set.tenant](resources--service_policy--reference--group-001.md#canonical-6e191d1b612878e275dc840661077f6bbb4d50192cba016a71dabd0f2de14862) |
| `deny_list.prefix_list` | [deny_list.prefix_list](resources--service_policy--reference--group-001.md#canonical-46c9c0d7b74a576ac88971d9b34fb4e537ebc1dd58ae57b487883097bc5a16b6) |
| `deny_list.prefix_list.prefixes` | [deny_list.prefix_list.prefixes](resources--service_policy--reference--group-001.md#canonical-3453ff68b8d9b92529c2ecd166d40d8ce65c7f619f800acb6585466749592316) |
| `deny_list.tls_fingerprint_classes` | [deny_list.tls_fingerprint_classes](resources--service_policy--reference--group-001.md#canonical-16653ab353561a617676531b3e33225cce46cc080c9618d8c863275a87caa409) |
| `deny_list.tls_fingerprint_values` | [deny_list.tls_fingerprint_values](resources--service_policy--reference--group-001.md#canonical-19d6c50d9237134938f29869d7df13329cde9f4b51cbb26ed282ee5a996015f0) |
| `description` | [description](resources--service_policy--reference--group-001.md#canonical-72d66a71744f262e9ea79466975d96331079a246f6779d42f5b85006985f4429) |
| `disable` | [disable](resources--service_policy--reference--group-001.md#canonical-5f66ba5b580a28a9025afb5058cf14c8559313dd5e738bb1faaea5b2327167fb) |
| `id` | [id](resources--service_policy--reference--group-001.md#canonical-48464a5d82ee61c357ed5cf342b7ab5454827bcba632849ab85f4e88fa48873b) |
| `labels` | [labels](resources--service_policy--reference--group-001.md#canonical-dec97802c3751df400c39b1cd20d6e73d1b81182116bfd2856330a1456a86d3c) |
| `name` | [name](resources--service_policy--reference--group-001.md#canonical-764987d7401b0171a6ea5577489b89b6a624e90ad3def2506972cccba7d42253) |
| `namespace` | [namespace](resources--service_policy--reference--group-001.md#canonical-f4245eea2baea0ffeab0fe7ab9e6ebdef2d96b10c3840232ba98316430cb739d) |
| `rule_list` | [rule_list](resources--service_policy--reference--group-001.md#canonical-64b41208835a700f1a079c0112584c8f2836b47f75ce6639800b2a28e818fddc) |
| `rule_list.rules` | [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-3461f52b3f822cd36bf5d407c52d0896676e173d04c16218020ad687a2993113) |
| `rule_list.rules.metadata` | [rule_list.rules.metadata](resources--service_policy--reference--group-001.md#canonical-2cd4e0b4923966168b812e9eb4e977d6ea65cdadf4780cc79fc502beefc4a510) |
| `rule_list.rules.metadata.description_spec` | [rule_list.rules.metadata.description_spec](resources--service_policy--reference--group-001.md#canonical-a7acf0c4fbe793bad9b1ca8a6571180ba051f4e56550126a74665e7e10e858d9) |
| `rule_list.rules.metadata.name` | [rule_list.rules.metadata.name](resources--service_policy--reference--group-001.md#canonical-65bc52ceb90388a6de195d6a00a3d8413c71d7c2becc4b7a1cb3a68eacf60dca) |
| `rule_list.rules.spec` | [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-cb65f1ff52f129defe8f85dca3c7915cd2adb03d6d8053ca87a8ffc206c0a0f3) |
| `rule_list.rules.spec.action` | [rule_list.rules.spec.action](resources--service_policy--reference--group-001.md#canonical-d9bbf42061ed1c6f44c169e738c350b87a20007a8f201a3c01333430b8d6fbc1) |
| `rule_list.rules.spec.any_asn` | [rule_list.rules.spec.any_asn](resources--service_policy--reference--group-002.md#canonical-e134d6a5d9a60f93809323edd42a7aae918e9acb15d38876c21f4e80f3d7f8ab) |
| `rule_list.rules.spec.any_client` | [rule_list.rules.spec.any_client](resources--service_policy--reference--group-002.md#canonical-84e7df3a30c686e70de011b78c5a29ff11ee1ab0549549de10a91f4dd5f08cf7) |
| `rule_list.rules.spec.any_ip` | [rule_list.rules.spec.any_ip](resources--service_policy--reference--group-002.md#canonical-1f8b76da9a3eb42707df1868ced979c405bddcc275c4b918516d266e098be4b2) |
| `rule_list.rules.spec.api_group_matcher` | [rule_list.rules.spec.api_group_matcher](resources--service_policy--reference--group-002.md#canonical-f0544e1a563e9ede2f14485626e72f0adc8c704b4389949d176ec4dc42206b65) |
| `rule_list.rules.spec.api_group_matcher.invert_matcher` | [rule_list.rules.spec.api_group_matcher.invert_matcher](resources--service_policy--reference--group-002.md#canonical-1529596393f8ef6a8e5be5f95d8b0b4cff21974c72bd18dd97f76c0e586da6d4) |
| `rule_list.rules.spec.api_group_matcher.match` | [rule_list.rules.spec.api_group_matcher.match](resources--service_policy--reference--group-002.md#canonical-a3ac11ef00cd48b0288d1e7c1b71f42b88f74233d3013622a4bab5e65635e07d) |
| `rule_list.rules.spec.arg_matchers` | [rule_list.rules.spec.arg_matchers](resources--service_policy--reference--group-002.md#canonical-5cd26a99acd917fc84cfc331b398032e52a0361833e18e0017896bb74eb1fb4e) |
| `rule_list.rules.spec.arg_matchers.check_not_present` | [rule_list.rules.spec.arg_matchers.check_not_present](resources--service_policy--reference--group-002.md#canonical-98c62422173687c5cdde478effb1e716e36518aee9961c5c34025230050872f3) |
| `rule_list.rules.spec.arg_matchers.check_present` | [rule_list.rules.spec.arg_matchers.check_present](resources--service_policy--reference--group-002.md#canonical-ce5959a222b142dfdb18e8a3ea4248d7b01ea7504c88b5d1d3c87371ac51b5a3) |
| `rule_list.rules.spec.arg_matchers.invert_matcher` | [rule_list.rules.spec.arg_matchers.invert_matcher](resources--service_policy--reference--group-002.md#canonical-85eed1eaf27b1d61cc441d8d923539f6c62ef13c301ec281bc0892a57a6bd02b) |
| `rule_list.rules.spec.arg_matchers.item` | [rule_list.rules.spec.arg_matchers.item](resources--service_policy--reference--group-002.md#canonical-c3bc2a57ff58bb63f737cdf45c5c0e6cc8bd54f1d47575e1ac71c845bfdee647) |
| `rule_list.rules.spec.arg_matchers.item.exact_values` | [rule_list.rules.spec.arg_matchers.item.exact_values](resources--service_policy--reference--group-002.md#canonical-c77f2781b4402c831e3cb148975acfd77230587ea50169c101cc07d1549a6caf) |
| `rule_list.rules.spec.arg_matchers.item.regex_values` | [rule_list.rules.spec.arg_matchers.item.regex_values](resources--service_policy--reference--group-002.md#canonical-6736858246db6fffd437c66bd9569aa5e39d893d156776ed087dab67511eae3e) |
| `rule_list.rules.spec.arg_matchers.item.transformers` | [rule_list.rules.spec.arg_matchers.item.transformers](resources--service_policy--reference--group-002.md#canonical-03f79a22baa7402b1810bd67653534765995454f5a5c6e10d9ed4c5eeda71cc9) |
| `rule_list.rules.spec.arg_matchers.name` | [rule_list.rules.spec.arg_matchers.name](resources--service_policy--reference--group-002.md#canonical-c534f631186981c5048ffe0a3661c818f5db972367594ce9b21b3ca8891c3004) |
| `rule_list.rules.spec.asn_list` | [rule_list.rules.spec.asn_list](resources--service_policy--reference--group-002.md#canonical-16a1dc2400b97c5c6b955d73e713d1864ce06131d1fc071e06b02d84d5778121) |
| `rule_list.rules.spec.asn_list.as_numbers` | [rule_list.rules.spec.asn_list.as_numbers](resources--service_policy--reference--group-002.md#canonical-9dbeb772e7d44f64f8260b2ddbf5982742a26f83dcf4b5c2061fa36b7d45fb1c) |
| `rule_list.rules.spec.asn_matcher` | [rule_list.rules.spec.asn_matcher](resources--service_policy--reference--group-002.md#canonical-37b81c135b0577f30d260798189eb22be695d83ce6337b61fdfa3c96a61fc4cf) |
| `rule_list.rules.spec.asn_matcher.asn_sets` | [rule_list.rules.spec.asn_matcher.asn_sets](resources--service_policy--reference--group-002.md#canonical-c22bdef86b129fcc6ee59a12659e3c5b69048d63cc7583de40a0672cec917481) |
| `rule_list.rules.spec.asn_matcher.asn_sets.kind` | [rule_list.rules.spec.asn_matcher.asn_sets.kind](resources--service_policy--reference--group-002.md#canonical-c7641521d9f0aff4cedb4ce006f60667dfa6853663ddb84aa2802651f1db9410) |
| `rule_list.rules.spec.asn_matcher.asn_sets.name` | [rule_list.rules.spec.asn_matcher.asn_sets.name](resources--service_policy--reference--group-002.md#canonical-a0d6bb31e592ff8a7e53037be55ca2dcc8ce461fcc2a716059415f0b7ca5729c) |
| `rule_list.rules.spec.asn_matcher.asn_sets.namespace` | [rule_list.rules.spec.asn_matcher.asn_sets.namespace](resources--service_policy--reference--group-002.md#canonical-901e07f6546d7906999d5245ae9bae0c99ed2158714436b188bb1214f4c18ae8) |
| `rule_list.rules.spec.asn_matcher.asn_sets.tenant` | [rule_list.rules.spec.asn_matcher.asn_sets.tenant](resources--service_policy--reference--group-002.md#canonical-07c676f5ce3bb522d1e562b0c3d9e73c690407086318f48a7cbd5ebc3d50c99e) |
| `rule_list.rules.spec.asn_matcher.asn_sets.uid` | [rule_list.rules.spec.asn_matcher.asn_sets.uid](resources--service_policy--reference--group-002.md#canonical-ad0703ace8271ba4c2a8d7c8973202aa1340a9f24f7f54c60ebe1bf668ce4e5a) |
| `rule_list.rules.spec.body_matcher` | [rule_list.rules.spec.body_matcher](resources--service_policy--reference--group-002.md#canonical-5ec3e70ca6f32d6dcdf075a370d4c20e640c83162e938e304257676da6163d78) |
| `rule_list.rules.spec.body_matcher.exact_values` | [rule_list.rules.spec.body_matcher.exact_values](resources--service_policy--reference--group-002.md#canonical-38401f5c7cd1ce1a923b2a47c74d533dc73d6f95c48a8aa298e64751ba3b46a1) |
| `rule_list.rules.spec.body_matcher.regex_values` | [rule_list.rules.spec.body_matcher.regex_values](resources--service_policy--reference--group-002.md#canonical-3b6392573bd68ec716d94422f6f964ff0f1e7449f08870144ec0207f25a3bb1d) |
| `rule_list.rules.spec.body_matcher.transformers` | [rule_list.rules.spec.body_matcher.transformers](resources--service_policy--reference--group-002.md#canonical-03a90b3da27dfe15ef32aa507d39587ec4a27338bb7d1c1c7772fe882f090ac4) |
| `rule_list.rules.spec.bot_action` | [rule_list.rules.spec.bot_action](resources--service_policy--reference--group-002.md#canonical-0f65f89d9f9b32237e6d7d0089db3c8ae7047622d2ad842ddf6e1170a433359d) |
| `rule_list.rules.spec.bot_action.bot_skip_processing` | [rule_list.rules.spec.bot_action.bot_skip_processing](resources--service_policy--reference--group-002.md#canonical-b67a5e86c98732080d0edc69500b4fe4003c6ccae668f4492977cf0eca3e77b2) |
| `rule_list.rules.spec.bot_action.none` | [rule_list.rules.spec.bot_action.none](resources--service_policy--reference--group-002.md#canonical-1cda7c78b3588b778db7bc71f8f1d6234573e30cd71f97da8fad45cbac27d57c) |
| `rule_list.rules.spec.client_name` | [rule_list.rules.spec.client_name](resources--service_policy--reference--group-001.md#canonical-1a00f339fa56fcf62ce5c2303fb12fe95e6bf38a8215c2358121dc17c02f8356) |
| `rule_list.rules.spec.client_name_matcher` | [rule_list.rules.spec.client_name_matcher](resources--service_policy--reference--group-002.md#canonical-5aa5291b8e0851fe2e582e6d9cad4be226bafed1acc1edb5ee2c116121b46a6f) |
| `rule_list.rules.spec.client_name_matcher.exact_values` | [rule_list.rules.spec.client_name_matcher.exact_values](resources--service_policy--reference--group-002.md#canonical-35dde138e0c41e00a459f025ac2f2c5c0a783c6fd1f0dddcdf1125e7713cbe53) |
| `rule_list.rules.spec.client_name_matcher.regex_values` | [rule_list.rules.spec.client_name_matcher.regex_values](resources--service_policy--reference--group-002.md#canonical-901f30afc858b949b74cf56006786371fd708d018790e1b2825b9a205831b345) |
| `rule_list.rules.spec.client_name_matcher.transformers` | [rule_list.rules.spec.client_name_matcher.transformers](resources--service_policy--reference--group-002.md#canonical-38c140709f14382922320edf64cafc88bce09e9cf60a18d8281987d5cbf46188) |
| `rule_list.rules.spec.client_selector` | [rule_list.rules.spec.client_selector](resources--service_policy--reference--group-002.md#canonical-835422e953a2bf8bc7fe2cdaad9c809cc118ccc62fe7ca7a30e578dc823effa5) |
| `rule_list.rules.spec.client_selector.expressions` | [rule_list.rules.spec.client_selector.expressions](resources--service_policy--reference--group-002.md#canonical-7ea48c988e92b2cd1c17c7b937731d8a08bd2bbe03ab4cd626505c9a72090498) |
| `rule_list.rules.spec.cookie_matchers` | [rule_list.rules.spec.cookie_matchers](resources--service_policy--reference--group-002.md#canonical-961425eca51fea6ab0c967bf84c0f930edc00ce0a6a961383e9e7fb56c2ea2fd) |
| `rule_list.rules.spec.cookie_matchers.check_not_present` | [rule_list.rules.spec.cookie_matchers.check_not_present](resources--service_policy--reference--group-002.md#canonical-90f63277484584e8ec6fd4605b023a983362c4163229c075f8d9e9cf8a3a70de) |
| `rule_list.rules.spec.cookie_matchers.check_present` | [rule_list.rules.spec.cookie_matchers.check_present](resources--service_policy--reference--group-002.md#canonical-ce4335b6a83cec6de4869afd4a4065c8df2d643d002831312503a88b136a09d6) |
| `rule_list.rules.spec.cookie_matchers.invert_matcher` | [rule_list.rules.spec.cookie_matchers.invert_matcher](resources--service_policy--reference--group-002.md#canonical-e7f5c26c2fbdb38f3e64260bd8ef20bc6d325e8be5c7c8c16eccfcde137ca4da) |
| `rule_list.rules.spec.cookie_matchers.item` | [rule_list.rules.spec.cookie_matchers.item](resources--service_policy--reference--group-002.md#canonical-ed7b27f920d579aff1b4d5b25f880b3fccc863b5d9184888e2e90c26f27b1072) |
| `rule_list.rules.spec.cookie_matchers.item.exact_values` | [rule_list.rules.spec.cookie_matchers.item.exact_values](resources--service_policy--reference--group-002.md#canonical-0af56230c6262b0c5d66b6d4e3bf9bd4c8c83dd4a2f70d8bb07727a018840ad7) |
| `rule_list.rules.spec.cookie_matchers.item.regex_values` | [rule_list.rules.spec.cookie_matchers.item.regex_values](resources--service_policy--reference--group-002.md#canonical-b4689146daa76f1d6dbea7665dc2908bae8826f51814510ce6b4c072350a5c73) |
| `rule_list.rules.spec.cookie_matchers.item.transformers` | [rule_list.rules.spec.cookie_matchers.item.transformers](resources--service_policy--reference--group-002.md#canonical-d56c7e830dfe31fe0e9ff97cf688519243419bd0b6040d21d8ac72deb0375beb) |
| `rule_list.rules.spec.cookie_matchers.name` | [rule_list.rules.spec.cookie_matchers.name](resources--service_policy--reference--group-002.md#canonical-752492a648a1ebdaff2e3fa8b69366d785954b37aed798f4d990481fbe6376d0) |
| `rule_list.rules.spec.domain_matcher` | [rule_list.rules.spec.domain_matcher](resources--service_policy--reference--group-002.md#canonical-c731aa6939bfce2beba8dad284bcc7cdc44131d5c9b1fcf0a36ee7dadd20304a) |
| `rule_list.rules.spec.domain_matcher.exact_values` | [rule_list.rules.spec.domain_matcher.exact_values](resources--service_policy--reference--group-002.md#canonical-e27765aa7b964b421509d1be23123cdad2d0c8aab8c34c1e21290a07d81795df) |
| `rule_list.rules.spec.domain_matcher.regex_values` | [rule_list.rules.spec.domain_matcher.regex_values](resources--service_policy--reference--group-002.md#canonical-a71379ca24d7da8a7b1d7d0de016b78423eb728dcac5190e72b4c91eb9b3e892) |
| `rule_list.rules.spec.domain_matcher.transformers` | [rule_list.rules.spec.domain_matcher.transformers](resources--service_policy--reference--group-002.md#canonical-f1c43b7af018ab6223cbfaf987314a87c99eed35d349d3ce709bdda847504e3b) |
| `rule_list.rules.spec.expiration_timestamp` | [rule_list.rules.spec.expiration_timestamp](resources--service_policy--reference--group-001.md#canonical-479775e6504f18997965092a297ba11c6d93460c87a9ced56fb69491fe8b72d9) |
| `rule_list.rules.spec.headers` | [rule_list.rules.spec.headers](resources--service_policy--reference--group-002.md#canonical-2762d2bf7d517f14f2b26853261f2a05e0cd7fe422d4208e177845b34a36c9a8) |
| `rule_list.rules.spec.headers.check_not_present` | [rule_list.rules.spec.headers.check_not_present](resources--service_policy--reference--group-002.md#canonical-eafaf60b0d3d9b7742cc7ed92774cdc2fe71ba1c4b12ec36ece7948cc81b2500) |
| `rule_list.rules.spec.headers.check_present` | [rule_list.rules.spec.headers.check_present](resources--service_policy--reference--group-002.md#canonical-7f5ec1d8a049c2924df6a31da5b4b637d304eb64ebe4649948e1596ee699cc10) |
| `rule_list.rules.spec.headers.invert_matcher` | [rule_list.rules.spec.headers.invert_matcher](resources--service_policy--reference--group-002.md#canonical-28dec286d341a7044a05c62c9607c4a53c9d1bf7379f1960540700c6c0b2f3ba) |
| `rule_list.rules.spec.headers.item` | [rule_list.rules.spec.headers.item](resources--service_policy--reference--group-002.md#canonical-76d20c6be44b3974eb36148f24489042f91803501fd1e7eb478415213cd6b2d1) |
| `rule_list.rules.spec.headers.item.exact_values` | [rule_list.rules.spec.headers.item.exact_values](resources--service_policy--reference--group-002.md#canonical-dbfd61eed03210ee2c51c6eb83fb23a769b555245d82bd813f420fe7ada3dbac) |
| `rule_list.rules.spec.headers.item.regex_values` | [rule_list.rules.spec.headers.item.regex_values](resources--service_policy--reference--group-002.md#canonical-f260b42a577e9342f8f21bd32acb2e2628e0d210e07eee056a079ff2f63f08c7) |
| `rule_list.rules.spec.headers.item.transformers` | [rule_list.rules.spec.headers.item.transformers](resources--service_policy--reference--group-002.md#canonical-ab80370a8fe4f350687e93ff02da4df360bb0ef0fc510eaada84311fe0ec0e9c) |
| `rule_list.rules.spec.headers.name` | [rule_list.rules.spec.headers.name](resources--service_policy--reference--group-002.md#canonical-5debd011aec08b3dd2928309a154fe50652f954762299f1345c86762581b63a1) |
| `rule_list.rules.spec.http_method` | [rule_list.rules.spec.http_method](resources--service_policy--reference--group-002.md#canonical-490ad5c826d62b410ba86a3619ed6cac86b0dd1a47245b3fc2c35b788f3eddd7) |
| `rule_list.rules.spec.http_method.invert_matcher` | [rule_list.rules.spec.http_method.invert_matcher](resources--service_policy--reference--group-002.md#canonical-c065ae234771944f521b43e0aac2e03f270954a5debaca651b7fb23c9fb77625) |
| `rule_list.rules.spec.http_method.methods` | [rule_list.rules.spec.http_method.methods](resources--service_policy--reference--group-002.md#canonical-3d77ce952eede9643819f88048640bda45e0a9b73003406d996e482759021f3a) |
| `rule_list.rules.spec.ip_matcher` | [rule_list.rules.spec.ip_matcher](resources--service_policy--reference--group-002.md#canonical-b0c5589590d5eec55615e28d94c8b0669191a539acb4132910a46c20faa1af3b) |
| `rule_list.rules.spec.ip_matcher.invert_matcher` | [rule_list.rules.spec.ip_matcher.invert_matcher](resources--service_policy--reference--group-002.md#canonical-9cba6df57a6a26747e6a4cec3d8f0f1b26a3dbbd274fa32a9a5602022772b008) |
| `rule_list.rules.spec.ip_matcher.prefix_sets` | [rule_list.rules.spec.ip_matcher.prefix_sets](resources--service_policy--reference--group-002.md#canonical-c7f5162dfc6c553c78ba13fa21427d85dc5336583a72e9cc5e069aa4d8ee3f4d) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.kind` | [rule_list.rules.spec.ip_matcher.prefix_sets.kind](resources--service_policy--reference--group-002.md#canonical-0aeb7ee91277cde405dbb3035c2ac118158aa85a4d66fac9be0fed893efaa3e2) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.name` | [rule_list.rules.spec.ip_matcher.prefix_sets.name](resources--service_policy--reference--group-002.md#canonical-f932d1e9a50873d9acb6b8f08ae4443e5fb6869addf363f7ab40cbd288913549) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.namespace` | [rule_list.rules.spec.ip_matcher.prefix_sets.namespace](resources--service_policy--reference--group-002.md#canonical-395bc7c7b5b20e961a0bb55a5978f32bf2e5a8e2da1e34073afa654c53a1c0cc) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.tenant` | [rule_list.rules.spec.ip_matcher.prefix_sets.tenant](resources--service_policy--reference--group-002.md#canonical-6436ec63951e29d5c4bcfde3ef4b6917daa7ae64ff5606ef130cca6fe435a740) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.uid` | [rule_list.rules.spec.ip_matcher.prefix_sets.uid](resources--service_policy--reference--group-002.md#canonical-1661435e68429a5ce26986f76df274dc334df64626c4e14912baa3a0ba8a0a70) |
| `rule_list.rules.spec.ip_prefix_list` | [rule_list.rules.spec.ip_prefix_list](resources--service_policy--reference--group-002.md#canonical-b66a92e9e3980e91370513d934553f820c107cff3dc052b975d67e208c6e7e76) |
| `rule_list.rules.spec.ip_prefix_list.invert_match` | [rule_list.rules.spec.ip_prefix_list.invert_match](resources--service_policy--reference--group-002.md#canonical-a3a05f242d8c7b4c07d88042975e3c6910eb79364a5e205b74627b681aeacc8f) |
| `rule_list.rules.spec.ip_prefix_list.ip_prefixes` | [rule_list.rules.spec.ip_prefix_list.ip_prefixes](resources--service_policy--reference--group-002.md#canonical-0bbf2c6eff4d3eec9781f5c09afac7ed04287ac5852c2d5590ffd91619ac7201) |
| `rule_list.rules.spec.ip_threat_category_list` | [rule_list.rules.spec.ip_threat_category_list](resources--service_policy--reference--group-002.md#canonical-7633cdbdf943e13b8abc6fba4040cb00d1138b91c66883f0492a7cfe04e8ea70) |
| `rule_list.rules.spec.ip_threat_category_list.ip_threat_categories` | [rule_list.rules.spec.ip_threat_category_list.ip_threat_categories](resources--service_policy--reference--group-002.md#canonical-f006a088910c9704a6361ab5a394bab2d7f7ccfc1b6563a381ed70d730525d9f) |
| `rule_list.rules.spec.ja4_tls_fingerprint` | [rule_list.rules.spec.ja4_tls_fingerprint](resources--service_policy--reference--group-002.md#canonical-a0f6543f94dc18f405aead66bfb2be9eb21942dc530f4ce7c6915c83f378895b) |
| `rule_list.rules.spec.ja4_tls_fingerprint.exact_values` | [rule_list.rules.spec.ja4_tls_fingerprint.exact_values](resources--service_policy--reference--group-002.md#canonical-10e0372eaddfafff0e3a9c421a2b6c228e04134afc163bc3239385959f8e2fdd) |
| `rule_list.rules.spec.jwt_claims` | [rule_list.rules.spec.jwt_claims](resources--service_policy--reference--group-002.md#canonical-4087de569318f2f45a30011940c321f651f9ad00c5afc225897cb9ae0eb69d54) |
| `rule_list.rules.spec.jwt_claims.check_not_present` | [rule_list.rules.spec.jwt_claims.check_not_present](resources--service_policy--reference--group-002.md#canonical-4012c3243a6befafd1c8c8bdc704012e633a43dfad881fd243d5c44841bc187a) |
| `rule_list.rules.spec.jwt_claims.check_present` | [rule_list.rules.spec.jwt_claims.check_present](resources--service_policy--reference--group-002.md#canonical-7fc6863380ed05e083df5df3edcae127ad6347d17d80d00d1a3837931b5ecaf9) |
| `rule_list.rules.spec.jwt_claims.invert_matcher` | [rule_list.rules.spec.jwt_claims.invert_matcher](resources--service_policy--reference--group-002.md#canonical-0d7fd34d79131f104f205c966a2877455d121e25339860f8dafb78178abc617d) |
| `rule_list.rules.spec.jwt_claims.item` | [rule_list.rules.spec.jwt_claims.item](resources--service_policy--reference--group-002.md#canonical-649d50015be782ff9bedd83c761f34f610cc707420e44a8b3f4c4d7d8ea2f062) |
| `rule_list.rules.spec.jwt_claims.item.exact_values` | [rule_list.rules.spec.jwt_claims.item.exact_values](resources--service_policy--reference--group-002.md#canonical-67a1823b0f8dfd5f32bf56e82252f4ad6a78658b9564fc9d72de7c57055d10d6) |
| `rule_list.rules.spec.jwt_claims.item.regex_values` | [rule_list.rules.spec.jwt_claims.item.regex_values](resources--service_policy--reference--group-002.md#canonical-fe8466986b477dc16fd362b4fd38ca374c87fc6dede7a6449859e313484c972a) |
| `rule_list.rules.spec.jwt_claims.item.transformers` | [rule_list.rules.spec.jwt_claims.item.transformers](resources--service_policy--reference--group-002.md#canonical-87157a9fe1c585856b897d55f90c37e3ace9f6f7d430fe13370e292c5ef46eea) |
| `rule_list.rules.spec.jwt_claims.name` | [rule_list.rules.spec.jwt_claims.name](resources--service_policy--reference--group-002.md#canonical-eb57b5fc966ea75a73ea1913b43c6a6c5d54645314cb4ae968e203d60776ec58) |
| `rule_list.rules.spec.label_matcher` | [rule_list.rules.spec.label_matcher](resources--service_policy--reference--group-002.md#canonical-ade546fd869f9133b35647217e05703cd0c91aa7312c310f6152a8e5a30c766b) |
| `rule_list.rules.spec.label_matcher.keys` | [rule_list.rules.spec.label_matcher.keys](resources--service_policy--reference--group-002.md#canonical-ba36748a19a0024cfa4f04d1e0aec4c6ece14f1d0873db3c06a7594d3bd56670) |
| `rule_list.rules.spec.log_rule_evaluation` | [rule_list.rules.spec.log_rule_evaluation](resources--service_policy--reference--group-001.md#canonical-7e11a4cc71b9e935dbbd6e9ae21cae7d74826a337e5af0b107b10a7b676ad778) |
| `rule_list.rules.spec.mum_action` | [rule_list.rules.spec.mum_action](resources--service_policy--reference--group-002.md#canonical-dc57a53e43dfc9c22260268e7eeae9bd5cb7af5524a53aad2a76d72b4a2f5b08) |
| `rule_list.rules.spec.mum_action.default` | [rule_list.rules.spec.mum_action.default](resources--service_policy--reference--group-002.md#canonical-639291696de8c6bcc59c7e3a4b9edc1657d1a062bb9a97f74f674c8439916312) |
| `rule_list.rules.spec.mum_action.skip_processing` | [rule_list.rules.spec.mum_action.skip_processing](resources--service_policy--reference--group-002.md#canonical-6829b33fe886015fa7a6623a8e9cbcf6cd8e3928b3b5367805b6c2e98f97a1aa) |
| `rule_list.rules.spec.path` | [rule_list.rules.spec.path](resources--service_policy--reference--group-002.md#canonical-d3b35f64302fc0ffe6eb9a5bbcdf1a36acfb5bfdd8295a24a11256518710a024) |
| `rule_list.rules.spec.path.encoded_path_matcher` | [rule_list.rules.spec.path.encoded_path_matcher](resources--service_policy--reference--group-002.md#canonical-eaeb98317e13c6a8c9c71e1536c00f674b9b41d8388c6f4b46660ad8dbdff1d1) |
| `rule_list.rules.spec.path.exact_values` | [rule_list.rules.spec.path.exact_values](resources--service_policy--reference--group-002.md#canonical-e4fd218fdded0d4cc10b99d6a1fe5b8ef1e5449fc86a0141cae42ffbacc8e26a) |
| `rule_list.rules.spec.path.invert_matcher` | [rule_list.rules.spec.path.invert_matcher](resources--service_policy--reference--group-002.md#canonical-432dfc025802f6b47ca258544cf7d13b92abeb5a7babb44891345aff6a8b9c69) |
| `rule_list.rules.spec.path.prefix_values` | [rule_list.rules.spec.path.prefix_values](resources--service_policy--reference--group-002.md#canonical-bd0b5338acde76cb47aba95cf637d937d77002edbf63ba11bee3fd67c7a5c43d) |
| `rule_list.rules.spec.path.regex_values` | [rule_list.rules.spec.path.regex_values](resources--service_policy--reference--group-002.md#canonical-5b687436354d2cfb7002b165e2ff76ecd3b54cb92859fe718b09214e99b4ea47) |
| `rule_list.rules.spec.path.suffix_values` | [rule_list.rules.spec.path.suffix_values](resources--service_policy--reference--group-002.md#canonical-102ea367c4bf828081c84856ce11bd220bc28aa873b092446768e08a5ac3c9c3) |
| `rule_list.rules.spec.path.transformers` | [rule_list.rules.spec.path.transformers](resources--service_policy--reference--group-002.md#canonical-6cb52f25b5c0f0353aa0850b14cf1bd16a23f9b39bc0a2653a064e3eac1d7e26) |
| `rule_list.rules.spec.port_matcher` | [rule_list.rules.spec.port_matcher](resources--service_policy--reference--group-002.md#canonical-3ac2b777a2f0e13572b6d4a7bb6274b9e4eac01e29eaff7e6121ee0768f87e9c) |
| `rule_list.rules.spec.port_matcher.invert_matcher` | [rule_list.rules.spec.port_matcher.invert_matcher](resources--service_policy--reference--group-002.md#canonical-52aafb0b8378d27b7d0e519114f0579cb9d4f43c9764eda51f01e2cce3001b35) |
| `rule_list.rules.spec.port_matcher.ports` | [rule_list.rules.spec.port_matcher.ports](resources--service_policy--reference--group-002.md#canonical-98fad030fd707c09f7ec5a6578ac05d27da19e6800413d4a750773fd196ce472) |
| `rule_list.rules.spec.query_params` | [rule_list.rules.spec.query_params](resources--service_policy--reference--group-002.md#canonical-f7cf6d700fcf3066f4eee8d309594a58190fa6242d0cbf3bbca401d3684c10d8) |
| `rule_list.rules.spec.query_params.check_not_present` | [rule_list.rules.spec.query_params.check_not_present](resources--service_policy--reference--group-002.md#canonical-a7d5146e75396fed4e663c0772788da963f650c369f59713705355a5e252069a) |
| `rule_list.rules.spec.query_params.check_present` | [rule_list.rules.spec.query_params.check_present](resources--service_policy--reference--group-002.md#canonical-f86a651ddf7c028f065479a5bca7f982f2e30ccccf59ee593e94569bf85abcb6) |
| `rule_list.rules.spec.query_params.invert_matcher` | [rule_list.rules.spec.query_params.invert_matcher](resources--service_policy--reference--group-002.md#canonical-80951dd26fa2af4f83eb4522d55b1112862eff28c96b42d089e5526f6e3cc936) |
| `rule_list.rules.spec.query_params.item` | [rule_list.rules.spec.query_params.item](resources--service_policy--reference--group-002.md#canonical-6278a86336657a50d21609b8a45bcecdd15414a7affa373fa54e125a93389542) |
| `rule_list.rules.spec.query_params.item.exact_values` | [rule_list.rules.spec.query_params.item.exact_values](resources--service_policy--reference--group-002.md#canonical-9253ce2c76e746226e85656ce835a658fb101510ab6282e4f67085998c3f80a5) |
| `rule_list.rules.spec.query_params.item.regex_values` | [rule_list.rules.spec.query_params.item.regex_values](resources--service_policy--reference--group-002.md#canonical-0aa925608d970aa34e84aeac6f8e776a0820a4fa19d40d415d16315d154fa3d5) |
| `rule_list.rules.spec.query_params.item.transformers` | [rule_list.rules.spec.query_params.item.transformers](resources--service_policy--reference--group-002.md#canonical-cf765a31d316ae55a2c4d183f451ccb929d648c246e0ae5e7012c08f11921628) |
| `rule_list.rules.spec.query_params.key` | [rule_list.rules.spec.query_params.key](resources--service_policy--reference--group-002.md#canonical-ca20e05769168b712606afed5512c3d293d397ff13ec7d8e0fccee3e0962f755) |
| `rule_list.rules.spec.request_constraints` | [rule_list.rules.spec.request_constraints](resources--service_policy--reference--group-002.md#canonical-ffcba47b21577a1308a5a113664554418b93660a5e8ddf12c0cb1d4054d83bcb) |
| `rule_list.rules.spec.request_constraints.max_cookie_count_exceeds` | [rule_list.rules.spec.request_constraints.max_cookie_count_exceeds](resources--service_policy--reference--group-002.md#canonical-e3807518e468ee0e4b978d03e18de0b932ed576bb8569ce9a9788d37b608c088) |
| `rule_list.rules.spec.request_constraints.max_cookie_count_none` | [rule_list.rules.spec.request_constraints.max_cookie_count_none](resources--service_policy--reference--group-003.md#canonical-8ff641e27180f2d4cf2e4be032523a75685c9cc7285a087b627eee0ae1c5f35b) |
| `rule_list.rules.spec.request_constraints.max_cookie_key_size_exceeds` | [rule_list.rules.spec.request_constraints.max_cookie_key_size_exceeds](resources--service_policy--reference--group-003.md#canonical-3fefcbe59d253f67c10c83698d1156a1b7adfad37e17c031ab755f0012f87ceb) |
| `rule_list.rules.spec.request_constraints.max_cookie_key_size_none` | [rule_list.rules.spec.request_constraints.max_cookie_key_size_none](resources--service_policy--reference--group-003.md#canonical-7aca03db5d566c338ebc70db39442ba295c86e96fad74d043a199af3e5628013) |
| `rule_list.rules.spec.request_constraints.max_cookie_value_size_exceeds` | [rule_list.rules.spec.request_constraints.max_cookie_value_size_exceeds](resources--service_policy--reference--group-003.md#canonical-1c846dc009d5ede0152c80167ede695bfd043d2d0dfd702ba91b2006908145e1) |
| `rule_list.rules.spec.request_constraints.max_cookie_value_size_none` | [rule_list.rules.spec.request_constraints.max_cookie_value_size_none](resources--service_policy--reference--group-003.md#canonical-d0591c0821a7158ab2f06a1ea558bfed9fd2f70e0d862724868b0d35240ec04b) |
| `rule_list.rules.spec.request_constraints.max_header_count_exceeds` | [rule_list.rules.spec.request_constraints.max_header_count_exceeds](resources--service_policy--reference--group-003.md#canonical-5690aa69b7bc21a9122a7b6da36ee187272698751e7c7d090f2dbec9b34308fd) |
| `rule_list.rules.spec.request_constraints.max_header_count_none` | [rule_list.rules.spec.request_constraints.max_header_count_none](resources--service_policy--reference--group-003.md#canonical-ed77a73aba6179c2fbf790d2ef01cd8ef256a94a00e8f3d4df3df3033ce3dea7) |
| `rule_list.rules.spec.request_constraints.max_header_key_size_exceeds` | [rule_list.rules.spec.request_constraints.max_header_key_size_exceeds](resources--service_policy--reference--group-003.md#canonical-1740941ace534eb5ef460a6469572e60cf157622921537e7bae62e22f5cb36b7) |
| `rule_list.rules.spec.request_constraints.max_header_key_size_none` | [rule_list.rules.spec.request_constraints.max_header_key_size_none](resources--service_policy--reference--group-003.md#canonical-b5e6aa65a21c75df647b982abfdcf39d663b3c7f63e08d0d0358fd8d42fa9529) |
| `rule_list.rules.spec.request_constraints.max_header_value_size_exceeds` | [rule_list.rules.spec.request_constraints.max_header_value_size_exceeds](resources--service_policy--reference--group-003.md#canonical-552e813cb27493da2643a71128ce6c43a8945a3c541c411a55a354d7fd9695d0) |
| `rule_list.rules.spec.request_constraints.max_header_value_size_none` | [rule_list.rules.spec.request_constraints.max_header_value_size_none](resources--service_policy--reference--group-003.md#canonical-208b62dde92d80034a7f1b7e4e7412eb30579d8ed2f2a95815b37eb93801a0a5) |
| `rule_list.rules.spec.request_constraints.max_parameter_count_exceeds` | [rule_list.rules.spec.request_constraints.max_parameter_count_exceeds](resources--service_policy--reference--group-003.md#canonical-e74663534f076e8095e5a5540e89bfbb133999b73b5a06831fcaaba76a876a1c) |
| `rule_list.rules.spec.request_constraints.max_parameter_count_none` | [rule_list.rules.spec.request_constraints.max_parameter_count_none](resources--service_policy--reference--group-003.md#canonical-450cfa03ad827a9671ef4ecca111cf74e26bf0c080d4e272a78c53f4be90efca) |
| `rule_list.rules.spec.request_constraints.max_parameter_name_size_exceeds` | [rule_list.rules.spec.request_constraints.max_parameter_name_size_exceeds](resources--service_policy--reference--group-003.md#canonical-f1a95629f30c789bc9181ccc2d2f59424a3fb241cc0906f8cc61cf95cc55a552) |
| `rule_list.rules.spec.request_constraints.max_parameter_name_size_none` | [rule_list.rules.spec.request_constraints.max_parameter_name_size_none](resources--service_policy--reference--group-003.md#canonical-4283d8d3694530ededa7d23174ec460dfdb28fd523073728f4c29d3c10d3fe3e) |
| `rule_list.rules.spec.request_constraints.max_parameter_value_size_exceeds` | [rule_list.rules.spec.request_constraints.max_parameter_value_size_exceeds](resources--service_policy--reference--group-003.md#canonical-7ebd0879f9618e74440341cb5e3db8547bba20f0004f29864656311874bfcf9b) |
| `rule_list.rules.spec.request_constraints.max_parameter_value_size_none` | [rule_list.rules.spec.request_constraints.max_parameter_value_size_none](resources--service_policy--reference--group-003.md#canonical-6e7d17139e61550a5c63d587b31ec61dd5fba801c24ea986783c592de906effc) |
| `rule_list.rules.spec.request_constraints.max_query_size_exceeds` | [rule_list.rules.spec.request_constraints.max_query_size_exceeds](resources--service_policy--reference--group-003.md#canonical-b78e3d3f096503d8a4e2810e2d4e9ff7f22eeec51562d3e5a388babee428a217) |
| `rule_list.rules.spec.request_constraints.max_query_size_none` | [rule_list.rules.spec.request_constraints.max_query_size_none](resources--service_policy--reference--group-003.md#canonical-a8ae2038b75339cf45de48bc0e3a11770f774ccfb93d25af5eeef052534cbfab) |
| `rule_list.rules.spec.request_constraints.max_request_line_size_exceeds` | [rule_list.rules.spec.request_constraints.max_request_line_size_exceeds](resources--service_policy--reference--group-003.md#canonical-55fc86eb832c7077737fb934a71a2ea34b391c4d2aa08f35324991bd75e30c3d) |
| `rule_list.rules.spec.request_constraints.max_request_line_size_none` | [rule_list.rules.spec.request_constraints.max_request_line_size_none](resources--service_policy--reference--group-003.md#canonical-7992a650ae4b203861ba01a94a8703d826afb54d46c1af01da0a453ab04e6ba3) |
| `rule_list.rules.spec.request_constraints.max_request_size_exceeds` | [rule_list.rules.spec.request_constraints.max_request_size_exceeds](resources--service_policy--reference--group-003.md#canonical-cf7c4ba22ee42a28164bffec1950058750f29ec82ddbb08407002d3ef306550d) |
| `rule_list.rules.spec.request_constraints.max_request_size_none` | [rule_list.rules.spec.request_constraints.max_request_size_none](resources--service_policy--reference--group-003.md#canonical-63d0e8e5e0900ae2623388165ff281237416257af873074c48837e626b978f7e) |
| `rule_list.rules.spec.request_constraints.max_url_size_exceeds` | [rule_list.rules.spec.request_constraints.max_url_size_exceeds](resources--service_policy--reference--group-003.md#canonical-d621245f1d3f7a3994332a8b06552b822ab9eeb64a202f22ab38e3291073e1e7) |
| `rule_list.rules.spec.request_constraints.max_url_size_none` | [rule_list.rules.spec.request_constraints.max_url_size_none](resources--service_policy--reference--group-003.md#canonical-fd444ea04ef8ba434703cdd7b25fc9f4bfeff5cc2a0de7825a37e026d1e1c372) |
| `rule_list.rules.spec.segment_policy` | [rule_list.rules.spec.segment_policy](resources--service_policy--reference--group-003.md#canonical-db18072a0934fb5a4992b0e239f6fd12340de6f0a848985db043717928d70c99) |
| `rule_list.rules.spec.segment_policy.dst_any` | [rule_list.rules.spec.segment_policy.dst_any](resources--service_policy--reference--group-003.md#canonical-46bc78791717eb7c3251bdbbcba0293797621b715879e3e0be1e9dcd7b5fa623) |
| `rule_list.rules.spec.segment_policy.dst_segments` | [rule_list.rules.spec.segment_policy.dst_segments](resources--service_policy--reference--group-003.md#canonical-1e80a79cf0f04eb64e6730599c576f72d2004dbb3f42143d3499bdb6d47784b2) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments` | [rule_list.rules.spec.segment_policy.dst_segments.segments](resources--service_policy--reference--group-003.md#canonical-be51d612a8c72d250d36b367b1c5f6abf86d11b8bb072fa84c630b29b930741b) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments.name` | [rule_list.rules.spec.segment_policy.dst_segments.segments.name](resources--service_policy--reference--group-003.md#canonical-088fe359555b65110d1c3047759e264e3fd20c3d595d38b16aa09d4df3bd7d81) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments.namespace` | [rule_list.rules.spec.segment_policy.dst_segments.segments.namespace](resources--service_policy--reference--group-003.md#canonical-1b06f80bd8e43d048af3a42d2fa302ba9587e77d24c0764a6a25c73d916fcf75) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments.tenant` | [rule_list.rules.spec.segment_policy.dst_segments.segments.tenant](resources--service_policy--reference--group-003.md#canonical-a058872bf4bcc350d0bb43e691d38a16e5ffed94580f55f9afe78d138a572742) |
| `rule_list.rules.spec.segment_policy.intra_segment` | [rule_list.rules.spec.segment_policy.intra_segment](resources--service_policy--reference--group-003.md#canonical-643c24a246e7218169258dad2453861a715148d868fe511c77115d892bed80f9) |
| `rule_list.rules.spec.segment_policy.src_any` | [rule_list.rules.spec.segment_policy.src_any](resources--service_policy--reference--group-003.md#canonical-49521907b1797cfa6666abd47cc8cebe1d7020334af68937cb0bef0dbc4ae2a1) |
| `rule_list.rules.spec.segment_policy.src_segments` | [rule_list.rules.spec.segment_policy.src_segments](resources--service_policy--reference--group-003.md#canonical-feb6871145e3225a736f37b035cbe6259e1c559ba47cdc46a45a51430e1f4bd6) |
| `rule_list.rules.spec.segment_policy.src_segments.segments` | [rule_list.rules.spec.segment_policy.src_segments.segments](resources--service_policy--reference--group-003.md#canonical-122d5a0b00c86ec65928b724732b923f9c9ed424719e415da1201802c3ca29ee) |
| `rule_list.rules.spec.segment_policy.src_segments.segments.name` | [rule_list.rules.spec.segment_policy.src_segments.segments.name](resources--service_policy--reference--group-003.md#canonical-3e75dc27315fbdd2b8cf1c32f50ac89a18351f239964afed81602787a86826af) |
| `rule_list.rules.spec.segment_policy.src_segments.segments.namespace` | [rule_list.rules.spec.segment_policy.src_segments.segments.namespace](resources--service_policy--reference--group-003.md#canonical-2821f5aa43d1069604e3f1a526e14015e37d133585cd3c5cff4228c56f1c7939) |
| `rule_list.rules.spec.segment_policy.src_segments.segments.tenant` | [rule_list.rules.spec.segment_policy.src_segments.segments.tenant](resources--service_policy--reference--group-003.md#canonical-565593c228411b1cf6f7352af2481d86aac3bcbf708fecde9625b66026475613) |
| `rule_list.rules.spec.tls_fingerprint_matcher` | [rule_list.rules.spec.tls_fingerprint_matcher](resources--service_policy--reference--group-003.md#canonical-a5cab8a51fdbe6dac16158e981f636009a68f81723a8a5c7f48ee47d9170c802) |
| `rule_list.rules.spec.tls_fingerprint_matcher.classes` | [rule_list.rules.spec.tls_fingerprint_matcher.classes](resources--service_policy--reference--group-003.md#canonical-3da192ed095934700ddea259e4740fb6bd9b984f39fb1109618f4c3733310ef7) |
| `rule_list.rules.spec.tls_fingerprint_matcher.exact_values` | [rule_list.rules.spec.tls_fingerprint_matcher.exact_values](resources--service_policy--reference--group-003.md#canonical-780a1405e3da85d3b670cfba320783baf51f814c7cd567b07b626417016203a2) |
| `rule_list.rules.spec.tls_fingerprint_matcher.excluded_values` | [rule_list.rules.spec.tls_fingerprint_matcher.excluded_values](resources--service_policy--reference--group-003.md#canonical-1ddbbc0154154787d19219a92604bef4c5ce92bc68c6050f8386e89604079b8b) |
| `rule_list.rules.spec.user_identity_matcher` | [rule_list.rules.spec.user_identity_matcher](resources--service_policy--reference--group-003.md#canonical-6e129b1019240134a7ee3e9f32c1f7668f4dbd5806f53a16566b249705708e32) |
| `rule_list.rules.spec.user_identity_matcher.exact_values` | [rule_list.rules.spec.user_identity_matcher.exact_values](resources--service_policy--reference--group-003.md#canonical-79d03fa30b7d1a99f62e9577db3f948bd6a2dcab4d864abcc3fd2b09050537f1) |
| `rule_list.rules.spec.user_identity_matcher.regex_values` | [rule_list.rules.spec.user_identity_matcher.regex_values](resources--service_policy--reference--group-003.md#canonical-8fde96738ceb16338cd03ec8241aa291c2b52b46bb23d94d8dca14b3296573ec) |
| `rule_list.rules.spec.waf_action` | [rule_list.rules.spec.waf_action](resources--service_policy--reference--group-003.md#canonical-bb74d13ec22c2e1cbaaa5036dc9d05c5b80964f1913f42605c5cf7232fa7fb8b) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control` | [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--reference--group-003.md#canonical-ece11ac74536b3ce1169bc10d08dc22abef3a0dfcf85ecf5a60da52b972ec400) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts](resources--service_policy--reference--group-003.md#canonical-e3546643c2eedb4dac4e07645701fc5fbf0a80d278b03ef02724ff700a5b6610) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context](resources--service_policy--reference--group-003.md#canonical-af9744104cea10643baeb9da4aebed1817021eec78ade07f2f8bd00763a3e3b0) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name](resources--service_policy--reference--group-003.md#canonical-71f4540d9ec543c0f7630b5950f318490830a8c88b4fc6af91ea5f874bbf8ba0) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type](resources--service_policy--reference--group-003.md#canonical-25b76d441c1d596286c3bd457a9faf3c6456acf6d63fc7f28a3e30f9a7d319d3) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts](resources--service_policy--reference--group-003.md#canonical-f05567d934380ea1c7f54218b296280578e4689d5b4d15d2bd8004b092eb921c) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name](resources--service_policy--reference--group-003.md#canonical-807636ee3826b9aedc7e7903f75c30cf584c00451bc08b7429b110abb60b7543) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts](resources--service_policy--reference--group-003.md#canonical-47006d64ec1b0e838dae87121fe7fd29918cb94b8288527ba3d85097c35825ef) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context](resources--service_policy--reference--group-003.md#canonical-4006bb4c5e4721f8759121e5d85e2913bcfad025afb31b857333e1ed14e5c117) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name](resources--service_policy--reference--group-003.md#canonical-e3984328ac6dc620d444a928e540cf071c7d936959d87366aefc74f2f82d2af3) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id](resources--service_policy--reference--group-003.md#canonical-7ca8c19e533362228fd0de50311dd05ff2e35ebe5cd490fa9543b7f2b9e95c89) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts](resources--service_policy--reference--group-003.md#canonical-50a7ae6da673301ca3fec365e62936b2789b8011d8eaec9de856eddca797d74e) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context](resources--service_policy--reference--group-003.md#canonical-71bbf0c7365b6cf9565039f83cb510707f48c910dd672462816776fbe64754c0) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name](resources--service_policy--reference--group-003.md#canonical-730ecb6406476ae187160830005eb1b0ff1ed536b1625de5550c5be3bd180e50) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation](resources--service_policy--reference--group-003.md#canonical-f7d72fd9effab4fe20eeb60febcfc54d3ffa651dc93abaf0fa763c43d626b13b) |
| `rule_list.rules.spec.waf_action.none` | [rule_list.rules.spec.waf_action.none](resources--service_policy--reference--group-003.md#canonical-5c95e10f062fff1bdadaafe2b301e3df8c1ac9441bc29b6b474dc17bfe451018) |
| `rule_list.rules.spec.waf_action.waf_skip_processing` | [rule_list.rules.spec.waf_action.waf_skip_processing](resources--service_policy--reference--group-003.md#canonical-3ef89b84f10229bef7f6f40e1e9819a0fdc3ea058e0753739874f9fe6ada3f8a) |
| `server_name` | [server_name](resources--service_policy--reference--group-001.md#canonical-1fac4dfe5494bdfc04e8af123a2f3b82ee09c42dce1cde9971a278281150ff5b) |
| `server_name_matcher` | [server_name_matcher](resources--service_policy--reference--group-003.md#canonical-80e3e5fd99369410697e0682fa18cd97d32af4c850aa895248f5a56b23e38066) |
| `server_name_matcher.exact_values` | [server_name_matcher.exact_values](resources--service_policy--reference--group-003.md#canonical-3e049eb59da2f103404aef24cb640e4e5a14c4cae92ddf732af6d33fa05aebed) |
| `server_name_matcher.regex_values` | [server_name_matcher.regex_values](resources--service_policy--reference--group-003.md#canonical-5c3098b14d6a661e9dc45ea8f16df74f7d1061f5b85f23ba6d61bb8a67488f04) |
| `server_selector` | [server_selector](resources--service_policy--reference--group-003.md#canonical-904ee2af17440d71cdf493a1ce98c30af916155ceb30517d84c371e051cba909) |
| `server_selector.expressions` | [server_selector.expressions](resources--service_policy--reference--group-003.md#canonical-2b0f1323ad1451b6ecc9179323a2631e1fc220e24f536a6c057573242605d3e8) |
| `timeouts` | [timeouts](resources--service_policy--reference--group-003.md#canonical-bef30de4cea2cb42a40409f558317d61fb10e72631ebed266bc38d1d4292c923) |
| `timeouts.create` | [timeouts.create](resources--service_policy--reference--group-003.md#canonical-1caffda87097fdedc967c3a9b2b835c5e3be69e720682e2ba705339db493200c) |
| `timeouts.delete` | [timeouts.delete](resources--service_policy--reference--group-003.md#canonical-4dc9b459511c80df0eda975950cc73547a7c4aa8c1f31caa73c0f147e50e59aa) |
| `timeouts.read` | [timeouts.read](resources--service_policy--reference--group-003.md#canonical-eafb5d0d70c439a9599f9c6db37cf30f1d62bda11c5c7712ab651497c4caa7e6) |
| `timeouts.update` | [timeouts.update](resources--service_policy--reference--group-003.md#canonical-ab3e7ee70b5f533bacba83b7261465348be4231e2e06e8801981a097b00771fa) |

<a id="canonical-f549a72fe78794df03399bbad2043566115f6cf5e57f8f047e71b1004507295a"></a>

## Next pages — Property reference / c3b918ccbafa / 13

- [allow_all_requests](resources--service_policy--reference--group-001.md#canonical-3b69e6598e66a1a089c034f8a0c09638e768e7c7463369245d991c9c5df2e4e5)
- [allow_list](resources--service_policy--reference--group-001.md#canonical-93c35c3efe81a429730c6d3e71a61214955b4938adda2886bd96ced391c75992)
- [any_server](resources--service_policy--reference--group-001.md#canonical-b6e1a4a11f301b65c95f28a92ddfe69d7d2cdfa290875a3fcf88e3dab50776a7)
- [deny_all_requests](resources--service_policy--reference--group-001.md#canonical-e9aaa06e2b0e97cd5cd7e20493c62d2557c9005231f3c38641984235f0dc1cc6)
- [deny_list](resources--service_policy--reference--group-001.md#canonical-f0b8d4e9de4ee80d1e3ff1ca4e51e89a004ef1276e52cc19138e08d501d46ee4)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [server_name_matcher](resources--service_policy--reference--group-003.md#canonical-de3685065567f4411ab04f86f8f14f9fc4e21e79652d75a6706a62e3882309c7)
- [server_selector](resources--service_policy--reference--group-003.md#canonical-ce09b78b2268dfdf225f8fa7808ee3dc2ee2eaa1ca7f4aef27737f6280575fab)
- [timeouts](resources--service_policy--reference--group-003.md#canonical-4c2b79645b86fc0db1b5a5e8348eb1539ad05a418b46eed14d7fd13168c502a6)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-3b69e6598e66a1a089c034f8a0c09638e768e7c7463369245d991c9c5df2e4e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d74b023445f1bbfe94ef79fe26bf2e76828015974c7482d0dae329c30de6a3ee"></a>

## allow_all_requests — allow_all_requests / 81d652786209 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- allow_all_requests

<a id="canonical-e3cb629e4f7ac08c9a96efb83d0e2b8171553728fd0cf8672d524b3505cb9857"></a>

Type: `["object", {}]`. Optional.

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

- [allow_all_requests](resources--service_policy--reference--group-001.md#canonical-e3cb629e4f7ac08c9a96efb83d0e2b8171553728fd0cf8672d524b3505cb9857)
- [allow_list](resources--service_policy--reference--group-001.md#canonical-4950530880ce5501a5e517c1ba34304e28870c9dee42078c60b005114342a10b)
- [deny_all_requests](resources--service_policy--reference--group-001.md#canonical-c0e726e216e23a86cf1e47c6f766758ef08efc2f4daab84f6f0a325181a1bc46)
- [deny_list](resources--service_policy--reference--group-001.md#canonical-9ce9bd163210878afb5cbb7de4ba16569a54be058aab14ee1c38a5ba7bc5341a)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-64b41208835a700f1a079c0112584c8f2836b47f75ce6639800b2a28e818fddc)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
allow_all_requests = {}
```

<a id="canonical-335b67a575edcfebf941634db7cc3534029a84351c7da67f30bb8c6a01a009af"></a>

## Direct properties — allow_all_requests / 81d652786209 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5c65ddbf87bb848be20b60cc61dbf082be140c6e1843f3d96d8935df73e18d2f"></a>

## Next pages — allow_all_requests / 81d652786209 / 4

- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-93c35c3efe81a429730c6d3e71a61214955b4938adda2886bd96ced391c75992"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ca0d3e9ab111c84e1a723c3fd2ef29c485119085dd0e21e96564caab53d1fce"></a>

## allow_list — allow_list / 7cec7147b005 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- allow_list

<a id="canonical-4950530880ce5501a5e517c1ba34304e28870c9dee42078c60b005114342a10b"></a>

Type: `"object"`. single nested block, Optional.

List of sources. A request belongs to this list if it satisfies any of the match criteria.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_action_allow",
    "default_action_deny"),
  validators.ConflictingObjectAttributes("default_action_allow",
    "default_action_next_policy"),
  validators.ConflictingObjectAttributes("default_action_deny",
    "default_action_next_policy")}
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
  "x-ves-oneof-field-default_action_choice": "[\"default_action_allow\",\"default_action_deny\",\"default_action_next_policy\"]"
}
```

Terraform syntax:

```terraform
allow_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-4ee7f9fe08c8ef443f6a794b87dd29a2158156f2032c36acf01a1beb524bf5d9"></a>

## Direct properties — allow_list / 7cec7147b005 / 3

- [asn_list](resources--service_policy--reference--group-001.md#canonical-de5240e6bd7ae9fd566c8c9c6233e449971e85dad309e1e0a5608bd18eeb3b4f): complete subsection reference.

- [asn_set](resources--service_policy--reference--group-001.md#canonical-fa74a0f9524fd54ec0151d33cab54c8a30e1c9e6de1d3f2104d300bf77c6e9fe): complete subsection reference.

<a id="canonical-f03508c854530683649e09917740c18e93c3b61a0729203b332cc0d105be53cb"></a>

<a id="canonical-8e8c0a1077a210652ef039dc1018507c639489a0df394081fe3c01b808c4a924"></a>

## country_list property — allow_list / 7cec7147b005 / 4

Type: `["list", "string"]`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

- [default_action_allow](resources--service_policy--reference--group-001.md#canonical-ec0aadd9ec7ae2d09ea774e71cceecb43e056181bd721010d980b39cbfb00a52): complete subsection reference.

- [default_action_deny](resources--service_policy--reference--group-001.md#canonical-6b38f1f97561a6aac5d297e6b4b12ea570238e33908d64695e2fe82ba3d6332c): complete subsection reference.

- [default_action_next_policy](resources--service_policy--reference--group-001.md#canonical-29e8f1142e466f8ecb735a8e499ab2baa9c2c0503ea9a5cf51c90e0d40c1ce5a): complete subsection reference.

- [ip_prefix_set](resources--service_policy--reference--group-001.md#canonical-ea964ee12f2977bda524123d9fa3964a56599e158d84a1ddbd29f41f9f4f2f7f): complete subsection reference.

- [prefix_list](resources--service_policy--reference--group-001.md#canonical-b4eb7eef0244c40d08cd503b29e629343f81102be4900772b6ac4e2bdcb171cd): complete subsection reference.

<a id="canonical-596bed29b44b3af5f8190b403daf45f1ba694477f1eac13dd2dd226fb1cd17b4"></a>

<a id="canonical-14fe607134786df5f3ce42d29bceb3812ea3eaf962ce5dd3a6171f38aaf031f4"></a>

## tls_fingerprint_classes property — allow_list / 7cec7147b005 / 5

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-cdb161d4e772bd9f064300914cce8dd0b1d1975eb309d54224b9452235ecd677"></a>

<a id="canonical-19ef0cb627cb00635634ab2ba050060c55e4899cf1faa8fa1201c8c5a52e49e2"></a>

## tls_fingerprint_values property — allow_list / 7cec7147b005 / 6

Type: `["list", "string"]`. Optional.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-a5a47b988f5a93e351e39feb2546279b12d3a026e724cf6f1070c3e8daa96238"></a>

## Next pages — allow_list / 7cec7147b005 / 7

- [allow_list.asn_list](resources--service_policy--reference--group-001.md#canonical-de5240e6bd7ae9fd566c8c9c6233e449971e85dad309e1e0a5608bd18eeb3b4f)
- [allow_list.asn_set](resources--service_policy--reference--group-001.md#canonical-fa74a0f9524fd54ec0151d33cab54c8a30e1c9e6de1d3f2104d300bf77c6e9fe)
- [allow_list.default_action_allow](resources--service_policy--reference--group-001.md#canonical-ec0aadd9ec7ae2d09ea774e71cceecb43e056181bd721010d980b39cbfb00a52)
- [allow_list.default_action_deny](resources--service_policy--reference--group-001.md#canonical-6b38f1f97561a6aac5d297e6b4b12ea570238e33908d64695e2fe82ba3d6332c)
- [allow_list.default_action_next_policy](resources--service_policy--reference--group-001.md#canonical-29e8f1142e466f8ecb735a8e499ab2baa9c2c0503ea9a5cf51c90e0d40c1ce5a)
- [allow_list.ip_prefix_set](resources--service_policy--reference--group-001.md#canonical-ea964ee12f2977bda524123d9fa3964a56599e158d84a1ddbd29f41f9f4f2f7f)
- [allow_list.prefix_list](resources--service_policy--reference--group-001.md#canonical-b4eb7eef0244c40d08cd503b29e629343f81102be4900772b6ac4e2bdcb171cd)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-de5240e6bd7ae9fd566c8c9c6233e449971e85dad309e1e0a5608bd18eeb3b4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab1c464c82c7d7b97033e87e0e0a6c082e1510f7b343844f26b5b3cd4338bb62"></a>

## allow_list.asn_list — allow_list.asn_list / 60a81600dfef / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [allow_list](resources--service_policy--reference--group-001.md#canonical-93c35c3efe81a429730c6d3e71a61214955b4938adda2886bd96ced391c75992)
- allow_list.asn_list

<a id="canonical-41493156507a550fc9d9c96c2e93f72015b3c652afaea794454c63ae94d87b9c"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-8625ca7c01f27c9b299cb10d726794ecf264ccdc1dad76136bc9ebce18f7e41c"></a>

## Direct properties — allow_list.asn_list / 60a81600dfef / 3

<a id="canonical-024568803759fb55ea6755131b55fdb5d8752ec17df5f5b0b0f9dfd55b41e408"></a>

<a id="canonical-8e226eb8fe82d7719b1abdcb616a8fa2fc21ef2ebc37b2077b911c70d0086d02"></a>

## as_numbers property — allow_list.asn_list / 60a81600dfef / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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

<a id="canonical-20b46fba13f638d725bb2e8a2c2482de2ae4fb123903d7e6d865caf4b8bc7cb5"></a>

## Next pages — allow_list.asn_list / 60a81600dfef / 5

- [allow_list](resources--service_policy--reference--group-001.md#canonical-93c35c3efe81a429730c6d3e71a61214955b4938adda2886bd96ced391c75992)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-fa74a0f9524fd54ec0151d33cab54c8a30e1c9e6de1d3f2104d300bf77c6e9fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1ba0d6289971af877571c3e3c3e2f377aa14524bb9012420249f9d078aeddef"></a>

## allow_list.asn_set — allow_list.asn_set / e8bdf129929b / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [allow_list](resources--service_policy--reference--group-001.md#canonical-93c35c3efe81a429730c6d3e71a61214955b4938adda2886bd96ced391c75992)
- allow_list.asn_set

<a id="canonical-615f9e756cde13d365f1257916860eb471ff83df7e96f4cb6190d2e819b51ffb"></a>

Type: `"object"`. list nested block, Optional.

Addresses that belong to the ASNs in the given bgp\_asn\_set The ASN is obtained by performing a
lookup for the source IPv4 Address in a GeoIP DB.

Upstream description:

Addresses that belong to the ASNs in the given bgp\_asn\_set The ASN is obtained by performing a
lookup for the source IPv4 Address in a GeoIP DB.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
asn_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-a990b1ad1bcf9cfae85e5d640ea12f82fb809797bdd8431c9f3695bc6c9b84bc"></a>

## Direct properties — allow_list.asn_set / e8bdf129929b / 3

<a id="canonical-925a25724e59e4d20445b4aee5385063a4651727652c7fa44981095cd59dd937"></a>

<a id="canonical-d275a2fee8536fe4e490310030091c4db3c20e6b331698e50a7a0737128c72b7"></a>

## name property — allow_list.asn_set / e8bdf129929b / 4

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

<a id="canonical-492efe7890f870bce94d6caddc00f92de00329efc097eac5abb9ec4ca624a00d"></a>

<a id="canonical-d22587e038bbb0a5810d57ffce76f86e5eebe201ffa0bf6748818af5b64671b3"></a>

## namespace property — allow_list.asn_set / e8bdf129929b / 5

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

<a id="canonical-43587761c1f3b83ce560f7b688394ffb0e8c8228c16a2667e9f4e6ce708138da"></a>

<a id="canonical-a10d8965556f5a75fc49580d44976b54ebab6a5300de70aa3a022b096cc555ea"></a>

## tenant property — allow_list.asn_set / e8bdf129929b / 6

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

<a id="canonical-5af47e3ca6c1e5a6dd5fe547ee4216e213e4b436c7f9e05f938b2c9d8e929e82"></a>

## Next pages — allow_list.asn_set / e8bdf129929b / 7

- [allow_list](resources--service_policy--reference--group-001.md#canonical-93c35c3efe81a429730c6d3e71a61214955b4938adda2886bd96ced391c75992)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-ec0aadd9ec7ae2d09ea774e71cceecb43e056181bd721010d980b39cbfb00a52"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-acba9eaf752dac5a9d1c9f196b43ee70cf04c7928a6bb79b57c5ad76c4a8118e"></a>

## allow_list.default_action_allow — allow_list.default_action_allow / fa8de3997141 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [allow_list](resources--service_policy--reference--group-001.md#canonical-93c35c3efe81a429730c6d3e71a61214955b4938adda2886bd96ced391c75992)
- allow_list.default_action_allow

<a id="canonical-0ea008c5cfc7ae353905c8a4af9937ed275b87a8fd71c9ae37683f00bbb8b318"></a>

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
default_action_allow = {}
```

<a id="canonical-1639cec4bbd869bf19722582550c5c83e4cf633936ff0830c9fcf5451c5503c5"></a>

## Direct properties — allow_list.default_action_allow / fa8de3997141 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6d3855f9386f44db98402496ee32628af74a16c3222c637003b5b80f92fd0593"></a>

## Next pages — allow_list.default_action_allow / fa8de3997141 / 4

- [allow_list](resources--service_policy--reference--group-001.md#canonical-93c35c3efe81a429730c6d3e71a61214955b4938adda2886bd96ced391c75992)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-6b38f1f97561a6aac5d297e6b4b12ea570238e33908d64695e2fe82ba3d6332c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e6c6f6c3b6552d51c0cf63a04b2b8cf115f4fb40ec9c5184de36235e9a2a6be"></a>

## allow_list.default_action_deny — allow_list.default_action_deny / b149d3c3e3c4 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [allow_list](resources--service_policy--reference--group-001.md#canonical-93c35c3efe81a429730c6d3e71a61214955b4938adda2886bd96ced391c75992)
- allow_list.default_action_deny

<a id="canonical-e03f0b6d8d42593f11c329fbd5ec6be3d61a6f67df97f141f72b750d1c111c49"></a>

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
default_action_deny = {}
```

<a id="canonical-c201d6b01ec22a1b9c8594ff205c6f299877872d7d849997c0f00d57bdfc4f66"></a>

## Direct properties — allow_list.default_action_deny / b149d3c3e3c4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-08d7c648017bd751ab82409f2c42c4541eee1292c77912aa4591c175dc4e30ce"></a>

## Next pages — allow_list.default_action_deny / b149d3c3e3c4 / 4

- [allow_list](resources--service_policy--reference--group-001.md#canonical-93c35c3efe81a429730c6d3e71a61214955b4938adda2886bd96ced391c75992)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-29e8f1142e466f8ecb735a8e499ab2baa9c2c0503ea9a5cf51c90e0d40c1ce5a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7859e0a7cdbe98167413ee1b3285dfdcab07e2ad94471c7fcffb1d73a424e8f"></a>

## allow_list.default_action_next_policy — allow_list.default_action_next_policy / de0e5a1647bc / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [allow_list](resources--service_policy--reference--group-001.md#canonical-93c35c3efe81a429730c6d3e71a61214955b4938adda2886bd96ced391c75992)
- allow_list.default_action_next_policy

<a id="canonical-3c23d353b680ee8e7f53bfd8e37bc0d22da142f8ecff1447f12e145602172bd0"></a>

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
default_action_next_policy = {}
```

<a id="canonical-6d23fb1bf5dcc0e51952a3728094dc56e9300c7ae618d065dfd8b9a8113fb307"></a>

## Direct properties — allow_list.default_action_next_policy / de0e5a1647bc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ab1bbff4814e2ecf72d93a7f27aa32d7c56411098f22d9234bc34a44b6929d05"></a>

## Next pages — allow_list.default_action_next_policy / de0e5a1647bc / 4

- [allow_list](resources--service_policy--reference--group-001.md#canonical-93c35c3efe81a429730c6d3e71a61214955b4938adda2886bd96ced391c75992)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-ea964ee12f2977bda524123d9fa3964a56599e158d84a1ddbd29f41f9f4f2f7f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cca22e91bc2d42b2444945df7521b8739b8fd8508ae165c3c24acd57069a7adc"></a>

## allow_list.ip_prefix_set — allow_list.ip_prefix_set / adc2c23a5d0f / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [allow_list](resources--service_policy--reference--group-001.md#canonical-93c35c3efe81a429730c6d3e71a61214955b4938adda2886bd96ced391c75992)
- allow_list.ip_prefix_set

<a id="canonical-c904dc6c2fe8a0ee24d82b9be9f510813f72d4d3df7aee0de04b8ebc319c0002"></a>

Type: `"object"`. list nested block, Optional.

Addresses that are covered by the prefixes in the given ip\_prefix\_set.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-5deab4fb3b8e80a5c9711335aa5414694b3252f2dbfda437c1c031b40bccc09f"></a>

## Direct properties — allow_list.ip_prefix_set / adc2c23a5d0f / 3

<a id="canonical-1423d8d94860249e7263e9a6684d7281695c880dc90cdc3098726ca5f78a7961"></a>

<a id="canonical-7f4ee058fea09e986c524cee5e240b9c6aefeac54b7704143513a9e33aad46e8"></a>

## name property — allow_list.ip_prefix_set / adc2c23a5d0f / 4

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

<a id="canonical-bdeaa6d178bb53198b86d068653270360186edffdb59c2c405d25179737d6575"></a>

<a id="canonical-80e314f9140c1f4b032e18c5ba77662312cb8589626943df21eb759ea87e321f"></a>

## namespace property — allow_list.ip_prefix_set / adc2c23a5d0f / 5

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

<a id="canonical-7152d6b1dd97b55887e5041b046525c743ef72f293cc44277889fd83547eca7f"></a>

<a id="canonical-1cd64ef7d25feffad0138c54eb8fb7d5186934393de79937e6f05dfa6ed3968e"></a>

## tenant property — allow_list.ip_prefix_set / adc2c23a5d0f / 6

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

<a id="canonical-ac5bb5a3420b2600af42080d9c4611a42b197b10b0fab1a878a49e6c9b34f5e7"></a>

## Next pages — allow_list.ip_prefix_set / adc2c23a5d0f / 7

- [allow_list](resources--service_policy--reference--group-001.md#canonical-93c35c3efe81a429730c6d3e71a61214955b4938adda2886bd96ced391c75992)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-b4eb7eef0244c40d08cd503b29e629343f81102be4900772b6ac4e2bdcb171cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02de3601f7903eb44cbb4f8f687e6b620a4875d5f958c8f86b0998a07f2a30e0"></a>

## allow_list.prefix_list — allow_list.prefix_list / 7dab21f5ad38 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [allow_list](resources--service_policy--reference--group-001.md#canonical-93c35c3efe81a429730c6d3e71a61214955b4938adda2886bd96ced391c75992)
- allow_list.prefix_list

<a id="canonical-0f17e5790c628158ccadd71761471fab46dc21fad5332c9b6cfea9ee5dd2fc75"></a>

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
prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-8425725e1127e572d70777179235c2310944fbc78b02cb975adce8cd7e930fb2"></a>

## Direct properties — allow_list.prefix_list / 7dab21f5ad38 / 3

<a id="canonical-a2338325f17dbc0419757bf57251081a19584cde986aa2c77afead476da46cb5"></a>

<a id="canonical-f3e3aa1ad8af520cacb2fdd5dee2cf71fa8902f09353804567324eec05f39c6b"></a>

## prefixes property — allow_list.prefix_list / 7dab21f5ad38 / 4

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

<a id="canonical-e877b8410a3d4718f664cf9efa1775e3a9d6dc4cd4cbf2ab34711cbd6c1d5b16"></a>

## Next pages — allow_list.prefix_list / 7dab21f5ad38 / 5

- [allow_list](resources--service_policy--reference--group-001.md#canonical-93c35c3efe81a429730c6d3e71a61214955b4938adda2886bd96ced391c75992)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-b6e1a4a11f301b65c95f28a92ddfe69d7d2cdfa290875a3fcf88e3dab50776a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d94c73b077aca15006928d704a302b5ff2d339cd40eec4441df37a14eae2256"></a>

## any_server — any_server / e6a94fe570b2 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- any_server

<a id="canonical-442b7a0e07e9bc7dfffeed2add1ded6a367c80ad7fb76a6d890162b2aa05aada"></a>

Type: `["object", {}]`. Optional, Computed.

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

- [any_server](resources--service_policy--reference--group-001.md#canonical-442b7a0e07e9bc7dfffeed2add1ded6a367c80ad7fb76a6d890162b2aa05aada)
- [server_name](resources--service_policy--reference--group-001.md#canonical-1fac4dfe5494bdfc04e8af123a2f3b82ee09c42dce1cde9971a278281150ff5b)
- [server_name_matcher](resources--service_policy--reference--group-003.md#canonical-80e3e5fd99369410697e0682fa18cd97d32af4c850aa895248f5a56b23e38066)
- [server_selector](resources--service_policy--reference--group-003.md#canonical-904ee2af17440d71cdf493a1ce98c30af916155ceb30517d84c371e051cba909)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
any_server = {}
```

<a id="canonical-f76bc98a733230aaa1d863411a44f527e994463166ab225bf261d06772772a4d"></a>

## Direct properties — any_server / e6a94fe570b2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-02188e0d0b3e9329ee190e9de2998f389d954f61b548d0d1d259b07e16102191"></a>

## Next pages — any_server / e6a94fe570b2 / 4

- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-e9aaa06e2b0e97cd5cd7e20493c62d2557c9005231f3c38641984235f0dc1cc6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e4b664642e1696e1580afa155fec06b6ca5b24f938fe04cfff8289e964635c9"></a>

## deny_all_requests — deny_all_requests / f930d01efd74 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- deny_all_requests

<a id="canonical-c0e726e216e23a86cf1e47c6f766758ef08efc2f4daab84f6f0a325181a1bc46"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
deny_all_requests = {}
```

<a id="canonical-1b15ee6b0900452a5d95b0e526addbb24d7857bb1f62d938e9f618ceb37b8baf"></a>

## Direct properties — deny_all_requests / f930d01efd74 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d476307aa8a48fb882c0e54d0a68542b4e9fe5f40d2ab07ac64dd2d4e74b59a7"></a>

## Next pages — deny_all_requests / f930d01efd74 / 4

- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-f0b8d4e9de4ee80d1e3ff1ca4e51e89a004ef1276e52cc19138e08d501d46ee4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b29831155955968ab55159141c43cfafc11c001078681b6bc94ee8cfc10d1853"></a>

## deny_list — deny_list / 7dcd616db6b9 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- deny_list

<a id="canonical-9ce9bd163210878afb5cbb7de4ba16569a54be058aab14ee1c38a5ba7bc5341a"></a>

Type: `"object"`. single nested block, Optional.

List of sources. A request belongs to this list if it satisfies any of the match criteria.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_action_allow",
    "default_action_deny"),
  validators.ConflictingObjectAttributes("default_action_allow",
    "default_action_next_policy"),
  validators.ConflictingObjectAttributes("default_action_deny",
    "default_action_next_policy")}
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
  "x-ves-oneof-field-default_action_choice": "[\"default_action_allow\",\"default_action_deny\",\"default_action_next_policy\"]"
}
```

Terraform syntax:

```terraform
deny_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-6fa7ef0cca57bcfbfc25ee74b7f3a4835336dd50278f5e8d52eefc5d7432312f"></a>

## Direct properties — deny_list / 7dcd616db6b9 / 3

- [asn_list](resources--service_policy--reference--group-001.md#canonical-e08f01654b5bf4da125c4015d879f5dc30797639b213c3859a7e4887dd8c3283): complete subsection reference.

- [asn_set](resources--service_policy--reference--group-001.md#canonical-7663a6b387243cf5e622bc689d957f73ef73564db439ffb2cc3048eb44a054da): complete subsection reference.

<a id="canonical-be2bf781e75005c7ce29f2d78f21f4ec0c56a0a9c9629636c7f502ee9c75c9c7"></a>

<a id="canonical-bfc3cc32a5caf38d0375b598f93efcbe7b514c770c0df255a095222e7d4c3323"></a>

## country_list property — deny_list / 7dcd616db6b9 / 4

Type: `["list", "string"]`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

- [default_action_allow](resources--service_policy--reference--group-001.md#canonical-606c05587f13487463fe7e48dcb2cf00e697b5c1b908499b621b37a0d2d402c1): complete subsection reference.

- [default_action_deny](resources--service_policy--reference--group-001.md#canonical-aeec559265e1cb94a3b344729d98a2821170dbe3964c6b5e0d1e5798cb480604): complete subsection reference.

- [default_action_next_policy](resources--service_policy--reference--group-001.md#canonical-eb622aad7d603239d6d73dab2cb50964783d2277a27335d4c13eda09dd14b125): complete subsection reference.

- [ip_prefix_set](resources--service_policy--reference--group-001.md#canonical-9b4bb379762d48c364a21078a13706e85c51b587bd939b4bb23a2787faebad5f): complete subsection reference.

- [prefix_list](resources--service_policy--reference--group-001.md#canonical-d3072c4b77350c6296ed3e71c23ba6cebe5de1106e0c7a50334baca0e7c334c4): complete subsection reference.

<a id="canonical-16653ab353561a617676531b3e33225cce46cc080c9618d8c863275a87caa409"></a>

<a id="canonical-59faa74b1e412924f5efbcd2d3febc6a9add5a518c42f2010d82e5588fa767e4"></a>

## tls_fingerprint_classes property — deny_list / 7dcd616db6b9 / 5

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-19d6c50d9237134938f29869d7df13329cde9f4b51cbb26ed282ee5a996015f0"></a>

<a id="canonical-d60ace973092d60eeaccfd720a837dbce4e44d80adde8aa3a55674a3b0032a6d"></a>

## tls_fingerprint_values property — deny_list / 7dcd616db6b9 / 6

Type: `["list", "string"]`. Optional.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-b08e3dec8df6a5c36ebfd835eebf6037fa8f943b0f18ee081b2505bc72135230"></a>

## Next pages — deny_list / 7dcd616db6b9 / 7

- [deny_list.asn_list](resources--service_policy--reference--group-001.md#canonical-e08f01654b5bf4da125c4015d879f5dc30797639b213c3859a7e4887dd8c3283)
- [deny_list.asn_set](resources--service_policy--reference--group-001.md#canonical-7663a6b387243cf5e622bc689d957f73ef73564db439ffb2cc3048eb44a054da)
- [deny_list.default_action_allow](resources--service_policy--reference--group-001.md#canonical-606c05587f13487463fe7e48dcb2cf00e697b5c1b908499b621b37a0d2d402c1)
- [deny_list.default_action_deny](resources--service_policy--reference--group-001.md#canonical-aeec559265e1cb94a3b344729d98a2821170dbe3964c6b5e0d1e5798cb480604)
- [deny_list.default_action_next_policy](resources--service_policy--reference--group-001.md#canonical-eb622aad7d603239d6d73dab2cb50964783d2277a27335d4c13eda09dd14b125)
- [deny_list.ip_prefix_set](resources--service_policy--reference--group-001.md#canonical-9b4bb379762d48c364a21078a13706e85c51b587bd939b4bb23a2787faebad5f)
- [deny_list.prefix_list](resources--service_policy--reference--group-001.md#canonical-d3072c4b77350c6296ed3e71c23ba6cebe5de1106e0c7a50334baca0e7c334c4)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-e08f01654b5bf4da125c4015d879f5dc30797639b213c3859a7e4887dd8c3283"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3234fdc806f498bad82726ac8f4595d05bb31c83fc5f4755a8db58017bec9992"></a>

## deny_list.asn_list — deny_list.asn_list / 7fb193948c94 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [deny_list](resources--service_policy--reference--group-001.md#canonical-f0b8d4e9de4ee80d1e3ff1ca4e51e89a004ef1276e52cc19138e08d501d46ee4)
- deny_list.asn_list

<a id="canonical-6423e6d7e05b1f2d5cb31c48fbd8f2fa7b8ab41b6f1aa6023c76268b84c534d2"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-6c2434cdf2593f0bd33b381c8d0eed2b1991514dc5a494d2b5f6debab6db5ac3"></a>

## Direct properties — deny_list.asn_list / 7fb193948c94 / 3

<a id="canonical-142ad56b4ab090426bb0415da3e290dcc30e98efb9eca7d56e9947edf16ba1b3"></a>

<a id="canonical-f5d171e8e91a76f0918e8d000eb05375a31d40f0a013a6a5b8cd5184fafd3920"></a>

## as_numbers property — deny_list.asn_list / 7fb193948c94 / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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

<a id="canonical-c4271926862c45c6ea3b9623b1bcb6c4d2122a466f8173ba41f7719d5116ed5b"></a>

## Next pages — deny_list.asn_list / 7fb193948c94 / 5

- [deny_list](resources--service_policy--reference--group-001.md#canonical-f0b8d4e9de4ee80d1e3ff1ca4e51e89a004ef1276e52cc19138e08d501d46ee4)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-7663a6b387243cf5e622bc689d957f73ef73564db439ffb2cc3048eb44a054da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9236db92ee524b6e627f0133e3e4defa9391aa38ce8e897081afc0683d685624"></a>

## deny_list.asn_set — deny_list.asn_set / 9a68362939d6 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [deny_list](resources--service_policy--reference--group-001.md#canonical-f0b8d4e9de4ee80d1e3ff1ca4e51e89a004ef1276e52cc19138e08d501d46ee4)
- deny_list.asn_set

<a id="canonical-6cbf23466b4e9132f4e56a9150524b8940afc6978410f74075dd324fa69d878a"></a>

Type: `"object"`. list nested block, Optional.

Addresses that belong to the ASNs in the given bgp\_asn\_set The ASN is obtained by performing a
lookup for the source IPv4 Address in a GeoIP DB.

Upstream description:

Addresses that belong to the ASNs in the given bgp\_asn\_set The ASN is obtained by performing a
lookup for the source IPv4 Address in a GeoIP DB.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
asn_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-3bcfa9ef4019e264819cb1e65615f2b38366e062a32bf0d017a85de248157a6c"></a>

## Direct properties — deny_list.asn_set / 9a68362939d6 / 3

<a id="canonical-760789468fd821c70cb6d3c636759a276fbd476df2429933f98a4feaa88760e0"></a>

<a id="canonical-02757af6a6845cc92cd208b1aa716d535105d96d9d3e31eeab104f530458dea4"></a>

## name property — deny_list.asn_set / 9a68362939d6 / 4

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

<a id="canonical-dccfd0f7b772160df141a154e57e61a752e4724ec5867cad295c6009374f0403"></a>

<a id="canonical-6b7b47ab200d1d5e0bd312d6bfd1cf2424b90f4a7f04fe0eeaa21ef3054b8c24"></a>

## namespace property — deny_list.asn_set / 9a68362939d6 / 5

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

<a id="canonical-59ae525382943fe22ed9abae8afdc5aa149f30454edd70a4cb2e914db195c2a4"></a>

<a id="canonical-ec631e2de2f1b4ea947358ddbbd040695a7e7488a57fa1356a6409588e11ab46"></a>

## tenant property — deny_list.asn_set / 9a68362939d6 / 6

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

<a id="canonical-f0e366733f4f705f9695c25eb4dbc9b8f5d48c111b717f2f87c5cd119cb472f6"></a>

## Next pages — deny_list.asn_set / 9a68362939d6 / 7

- [deny_list](resources--service_policy--reference--group-001.md#canonical-f0b8d4e9de4ee80d1e3ff1ca4e51e89a004ef1276e52cc19138e08d501d46ee4)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-606c05587f13487463fe7e48dcb2cf00e697b5c1b908499b621b37a0d2d402c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71ccb4981451f8ea4e9b0fc6588ef1e8e2b65b3079d026586b5de4f0f4fed5ca"></a>

## deny_list.default_action_allow — deny_list.default_action_allow / a018c9bf8330 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [deny_list](resources--service_policy--reference--group-001.md#canonical-f0b8d4e9de4ee80d1e3ff1ca4e51e89a004ef1276e52cc19138e08d501d46ee4)
- deny_list.default_action_allow

<a id="canonical-5b6d88227f6af615b007424e6788b2fd9cec751f694a855199e023e1b85993e2"></a>

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
default_action_allow = {}
```

<a id="canonical-b5e7caf55be3d1020a18fbdd0fa59b04ef4435cfeb70dd1fdf8cbc8553cb51d6"></a>

## Direct properties — deny_list.default_action_allow / a018c9bf8330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b958b296c09931ff07348bdc062be8662370f1492f23f6c317e3d710589d1ab5"></a>

## Next pages — deny_list.default_action_allow / a018c9bf8330 / 4

- [deny_list](resources--service_policy--reference--group-001.md#canonical-f0b8d4e9de4ee80d1e3ff1ca4e51e89a004ef1276e52cc19138e08d501d46ee4)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-aeec559265e1cb94a3b344729d98a2821170dbe3964c6b5e0d1e5798cb480604"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5f59fedb610485e04a30834306275a4728faf257fea8c0c2f55ce34d8347323"></a>

## deny_list.default_action_deny — deny_list.default_action_deny / 5ca3011cae0a / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [deny_list](resources--service_policy--reference--group-001.md#canonical-f0b8d4e9de4ee80d1e3ff1ca4e51e89a004ef1276e52cc19138e08d501d46ee4)
- deny_list.default_action_deny

<a id="canonical-a58db399c83c51e340623a595c56dbd7cf78e316e222f153d396f73e27985af4"></a>

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
default_action_deny = {}
```

<a id="canonical-3673c858bc4c4300109b2a79e514c48c6e3ce16dfd7a6fa62d37a51c9af5f1b7"></a>

## Direct properties — deny_list.default_action_deny / 5ca3011cae0a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-abfe38e9a928eade5b68c4582578ba6e8a248815b02eb471ff30c1da7253c0ae"></a>

## Next pages — deny_list.default_action_deny / 5ca3011cae0a / 4

- [deny_list](resources--service_policy--reference--group-001.md#canonical-f0b8d4e9de4ee80d1e3ff1ca4e51e89a004ef1276e52cc19138e08d501d46ee4)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-eb622aad7d603239d6d73dab2cb50964783d2277a27335d4c13eda09dd14b125"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-180e9a450a72aa82993ce45ce1d1bb9f2c2a388ec6aa414b1637ea0fa2262e57"></a>

## deny_list.default_action_next_policy — deny_list.default_action_next_policy / 6dcb9f793024 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [deny_list](resources--service_policy--reference--group-001.md#canonical-f0b8d4e9de4ee80d1e3ff1ca4e51e89a004ef1276e52cc19138e08d501d46ee4)
- deny_list.default_action_next_policy

<a id="canonical-e56ddbd1b81330c100843d4b0c8a4b7123bb40f3be202f7b60f4a6d8ea55f742"></a>

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
default_action_next_policy = {}
```

<a id="canonical-19c28a3f8699c234ef4aea7970ab92825610e5391cb5e558041eccaf50af9a1e"></a>

## Direct properties — deny_list.default_action_next_policy / 6dcb9f793024 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7708e9f02d5317ff4e77b575df2a0be70304ffb88578939900e8ecb8cf770e3f"></a>

## Next pages — deny_list.default_action_next_policy / 6dcb9f793024 / 4

- [deny_list](resources--service_policy--reference--group-001.md#canonical-f0b8d4e9de4ee80d1e3ff1ca4e51e89a004ef1276e52cc19138e08d501d46ee4)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-9b4bb379762d48c364a21078a13706e85c51b587bd939b4bb23a2787faebad5f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7cfd488b04800c50849a138712c21c5db79fce0c3e1147098848edb5d07d6100"></a>

## deny_list.ip_prefix_set — deny_list.ip_prefix_set / 6ffcec914d24 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [deny_list](resources--service_policy--reference--group-001.md#canonical-f0b8d4e9de4ee80d1e3ff1ca4e51e89a004ef1276e52cc19138e08d501d46ee4)
- deny_list.ip_prefix_set

<a id="canonical-a23b4c23fce6e6217c17bc6d709efa79691811e37064bacbcb73f643354bdba0"></a>

Type: `"object"`. list nested block, Optional.

Addresses that are covered by the prefixes in the given ip\_prefix\_set.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-26859a502006e05568bb88c8383aa5d3bc936b3157dd7bdfc48ecc44a7855ebe"></a>

## Direct properties — deny_list.ip_prefix_set / 6ffcec914d24 / 3

<a id="canonical-868e1c9f903622234134659119c8a8ecd6b0027501db2a52239aeec419fd1415"></a>

<a id="canonical-98ab110e327fd55dc645b139cdddff68ffc81e2d5c4084fc3aa7df27d4829fd9"></a>

## name property — deny_list.ip_prefix_set / 6ffcec914d24 / 4

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

<a id="canonical-a639d0441b90e44c997299e21fff249e684243973ff6cb1ba3c60a05ee41f1c0"></a>

<a id="canonical-4fbe0dc3fe74c59598ad6b7220e7177efb9bb8d577e93826a1363dc702e0ad3e"></a>

## namespace property — deny_list.ip_prefix_set / 6ffcec914d24 / 5

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

<a id="canonical-6e191d1b612878e275dc840661077f6bbb4d50192cba016a71dabd0f2de14862"></a>

<a id="canonical-e7d2852547d203ee7743566813a5d6eea775b490d6574527f2e65b6e34a54b76"></a>

## tenant property — deny_list.ip_prefix_set / 6ffcec914d24 / 6

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

<a id="canonical-c0c4b9f458ebaf6c724b510b48dc21cb63c62ecc95d242ae9481242099880caf"></a>

## Next pages — deny_list.ip_prefix_set / 6ffcec914d24 / 7

- [deny_list](resources--service_policy--reference--group-001.md#canonical-f0b8d4e9de4ee80d1e3ff1ca4e51e89a004ef1276e52cc19138e08d501d46ee4)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-d3072c4b77350c6296ed3e71c23ba6cebe5de1106e0c7a50334baca0e7c334c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fcf4f73d44d5e1d76f9fab9fe22b7c2967669058acbf3b262633ffeac90f0cd"></a>

## deny_list.prefix_list — deny_list.prefix_list / 584185c7db77 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [deny_list](resources--service_policy--reference--group-001.md#canonical-f0b8d4e9de4ee80d1e3ff1ca4e51e89a004ef1276e52cc19138e08d501d46ee4)
- deny_list.prefix_list

<a id="canonical-46c9c0d7b74a576ac88971d9b34fb4e537ebc1dd58ae57b487883097bc5a16b6"></a>

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
prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-b0114a0c17bf0b01d9eb49d033921a412a3655848e7dc822d75a11600fe8020b"></a>

## Direct properties — deny_list.prefix_list / 584185c7db77 / 3

<a id="canonical-3453ff68b8d9b92529c2ecd166d40d8ce65c7f619f800acb6585466749592316"></a>

<a id="canonical-10b2dc6f7a42f7cb37c68291bd17db318184aeed7c7129c3213a35a3417390e0"></a>

## prefixes property — deny_list.prefix_list / 584185c7db77 / 4

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

<a id="canonical-906d4d6765a87f99007929d7113807503cf795a0012fdde7fc78b2f030935673"></a>

## Next pages — deny_list.prefix_list / 584185c7db77 / 5

- [deny_list](resources--service_policy--reference--group-001.md#canonical-f0b8d4e9de4ee80d1e3ff1ca4e51e89a004ef1276e52cc19138e08d501d46ee4)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57c8ab81bcb8945dd5da03e331e485038a42809929c3cc52b3e9e5bffc0b33e5"></a>

## rule_list — rule_list / 0fc7ebe6245b / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- rule_list

<a id="canonical-64b41208835a700f1a079c0112584c8f2836b47f75ce6639800b2a28e818fddc"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
rule_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-eac43cb14c807230ff009f50aa04948255d38c90c775b98de0f575fd60154852"></a>

## Direct properties — rule_list / 0fc7ebe6245b / 3

- [rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed): complete subsection reference.

<a id="canonical-313aa6bea90dc8cebce702e021d71ab9f7d9ca05a2fe8c0f982617d19e92482c"></a>

## Next pages — rule_list / 0fc7ebe6245b / 4

- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ed48b04631003e4a5f5e012f7f718ac01e5bb2e3312e30e79c47337cb4d222b"></a>

## rule_list.rules — rule_list.rules / cf3bf2d86ab6 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- rule_list.rules

<a id="canonical-3461f52b3f822cd36bf5d407c52d0896676e173d04c16218020ad687a2993113"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-c7c66123154440c7ecb0a5f2fc60bc89890be5e94c0037ca8c62b952f04f2fde"></a>

## Direct properties — rule_list.rules / cf3bf2d86ab6 / 3

- [metadata](resources--service_policy--reference--group-001.md#canonical-5b9aefc006b95894b75d4bb425dc841ddbd0e81a530dfb66ed3e9ba8f1a47a52): complete subsection reference.

- [spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6): complete subsection reference.

<a id="canonical-9b80ed4595ce4992a450cc93b10a6a37b85ec92d8f118b4ab5d171452ae5168c"></a>

## Next pages — rule_list.rules / cf3bf2d86ab6 / 4

- [rule_list.rules.metadata](resources--service_policy--reference--group-001.md#canonical-5b9aefc006b95894b75d4bb425dc841ddbd0e81a530dfb66ed3e9ba8f1a47a52)
- [rule_list.rules.spec](resources--service_policy--reference--group-001.md#canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-5b9aefc006b95894b75d4bb425dc841ddbd0e81a530dfb66ed3e9ba8f1a47a52"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ce77d732c8c7bcb8cadf92000f344099276422dca6741a83ae491f7d20d493e"></a>

## rule_list.rules.metadata — rule_list.rules.metadata / c36ff5b84354 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- rule_list.rules.metadata

<a id="canonical-2cd4e0b4923966168b812e9eb4e977d6ea65cdadf4780cc79fc502beefc4a510"></a>

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

<a id="canonical-ba6296de62e2ba7028ada4115a24b302cba7230d20191d3b769c9629ed74b239"></a>

## Direct properties — rule_list.rules.metadata / c36ff5b84354 / 3

<a id="canonical-a7acf0c4fbe793bad9b1ca8a6571180ba051f4e56550126a74665e7e10e858d9"></a>

<a id="canonical-fe755ed7f9e53eb472527745a0a38ef9b8b658cfa9126988a6ab2a17b399463a"></a>

## description_spec property — rule_list.rules.metadata / c36ff5b84354 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-65bc52ceb90388a6de195d6a00a3d8413c71d7c2becc4b7a1cb3a68eacf60dca"></a>

<a id="canonical-624e0c45827720d51466d2b1db4c10b9e512cdea0e7081dd50fe51feb8ca6656"></a>

## name property — rule_list.rules.metadata / c36ff5b84354 / 5

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

<a id="canonical-05f279b198d65b555f763a9b30cdb7b76a0630fce4ca25dbe1c0142d75ad5f84"></a>

## Next pages — rule_list.rules.metadata / c36ff5b84354 / 6

- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-a02a763e28951c57ce2b73affc73fbcc04514d5756feb6f02edc61043205bea6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4074cfe5b23a518a0d9c761475df5d061780589c8276eec9c61c1a2d6fb87a5c"></a>

## rule_list.rules.spec — rule_list.rules.spec / bd9f0a29db3f / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Property reference](resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [rule_list](resources--service_policy--reference--group-001.md#canonical-373e255ac8c837b6f6820b13bb2b0dd545347db3c4d72533c81bc75a44c91490)
- [rule_list.rules](resources--service_policy--reference--group-001.md#canonical-833e592b6fb878cf671b84c3680e40bae8c0d94183fbf7b03128e5861e27caed)
- rule_list.rules.spec

<a id="canonical-cb65f1ff52f129defe8f85dca3c7915cd2adb03d6d8053ca87a8ffc206c0a0f3"></a>

Type: `"object"`. single nested block, Optional.

Shape of service\_policy\_rule in the storage backend.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("action",
    "waf_action"),
  validators.ConflictingObjectAttributes("any_asn",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_asn",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_client",
    "client_name"),
  validators.ConflictingObjectAttributes("any_client",
    "client_name_matcher"),
  validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_client",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("client_name",
    "client_name_matcher"),
  validators.ConflictingObjectAttributes("client_name",
    "client_selector"),
  validators.ConflictingObjectAttributes("client_name",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("client_name_matcher",
    "client_selector"),
  validators.ConflictingObjectAttributes("client_name_matcher",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("client_selector",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("ja4_tls_fingerprint",
    "tls_fingerprint_matcher")}
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
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_name\",\"client_name_matcher\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-dst_asn_choice": "[]",
  "x-ves-oneof-field-dst_ip_choice": "[]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]",
  "x-ves-oneof-field-tls_fingerprint_choice": "[\"ja4_tls_fingerprint\",\"tls_fingerprint_matcher\"]"
}
```

Terraform syntax:

```terraform
spec {
  # Configure direct properties listed below.
}
```

<a id="canonical-1f1e99063e3cb27525ff225e9fde25c21bf473e6484108e92a0dfe20f7be5668"></a>

## Direct properties — rule_list.rules.spec / bd9f0a29db3f / 3

<a id="canonical-d9bbf42061ed1c6f44c169e738c350b87a20007a8f201a3c01333430b8d6fbc1"></a>

<a id="canonical-fa491ca8cd330ba48ff4f95356338b5923303924aa6ebde523afc9414b604dd8"></a>

## action property — rule_list.rules.spec / bd9f0a29db3f / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DENY",
    "ALLOW",
    "NEXT_POLICY"),
}
```

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

- [any_asn](resources--service_policy--reference--group-002.md#canonical-ca4e4da4c437c7579c5290d5891880b6bdfa936b16ceab8903c102936cad206f): complete subsection reference.

- [any_client](resources--service_policy--reference--group-002.md#canonical-fcdb9ec8f1b14e78c1bd08a0d6943014915c93b9a088166ecb6311868ab78b6b): complete subsection reference.

- [any_ip](resources--service_policy--reference--group-002.md#canonical-5a5b0501e8b50e873b9eda0bc8ee70d5ec75ab6a85281c55ed9f485f553ff3c7): complete subsection reference.

- [api_group_matcher](resources--service_policy--reference--group-002.md#canonical-606c7fa6e3cc41c9e309f1e443dd17a517cafe95aa3c65d715078ba5e39e9575): complete subsection reference.

- [arg_matchers](resources--service_policy--reference--group-002.md#canonical-3a118f18d58bfb1a945f0a9ab837e00ef758abe4c8de5411292b5085065033d6): complete subsection reference.

- [asn_list](resources--service_policy--reference--group-002.md#canonical-4308548e1215b0cc5ef0318de3270e278c5841d130d0ae730ea660a8acada2c5): complete subsection reference.

- [asn_matcher](resources--service_policy--reference--group-002.md#canonical-2e94d799c2e5faecf3907c8e3bc930fe1ddd8ad9cddb2fe9c640f9a0bb437b57): complete subsection reference.

- [body_matcher](resources--service_policy--reference--group-002.md#canonical-0ce427881bd5760fc5c5bbe133fa64a43e54da1e903179945c549cae489dbc52): complete subsection reference.

- [bot_action](resources--service_policy--reference--group-002.md#canonical-328329473d917c4183f71b3fbe74418fd5c1bad2bf05fa12145e4adde19d4459): complete subsection reference.

<a id="canonical-1a00f339fa56fcf62ce5c2303fb12fe95e6bf38a8215c2358121dc17c02f8356"></a>

<a id="canonical-dfd0df300373d27a51a5c2498e923f44a70bc18b58dd43db187b6079c82bc93d"></a>

## client_name property — rule_list.rules.spec / bd9f0a29db3f / 5

Type: `"string"`. Optional.

Exclusive with \[any\_client client\_name\_matcher client\_selector ip\_threat\_category\_list\] The
expected name of the client invoking the request API. The predicate evaluates to true if any of the
actual names is the same as the expected client name.

Upstream description:

Exclusive with \[any\_client client\_name\_matcher client\_selector ip\_threat\_category\_list\] The
expected name of the client invoking the request API. The predicate evaluates to true if any of the
actual names is the same as the expected client name.

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

- [client_name_matcher](resources--service_policy--reference--group-002.md#canonical-47940987a6b987a1f5725e9ef832132df4bca22c4bca0a9f56602ddb4d827184): complete subsection reference.

- [client_selector](resources--service_policy--reference--group-002.md#canonical-25849bb50cc43905b401b3741c3d6d19cebff6d01d5dada3c56f849d23a089ef): complete subsection reference.

- [cookie_matchers](resources--service_policy--reference--group-002.md#canonical-d9998220b18fc62913d6721b8d71a19581cd1a60ca4421b669f4deb494cf2139): complete subsection reference.

- [domain_matcher](resources--service_policy--reference--group-002.md#canonical-0260c1f4c22f54a9a9604bc3162c94ae94c3153ceaafb916422c395c6630e4a5): complete subsection reference.

<a id="canonical-479775e6504f18997965092a297ba11c6d93460c87a9ced56fb69491fe8b72d9"></a>

<a id="canonical-9d0c8d9ac8283174f404ee377f06dc4d4dcb0edefa8efa658430246a9ff286c4"></a>

## expiration_timestamp property — rule_list.rules.spec / bd9f0a29db3f / 6

Type: `"string"`. Optional.

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

- [headers](resources--service_policy--reference--group-002.md#canonical-f3956b2ddda07966dd2b83d15cc6ebe7315b3ea30b5c251d236836ad2c5582b1): complete subsection reference.

- [http_method](resources--service_policy--reference--group-002.md#canonical-e0ce9cf93e6e22950e5f4354ff81bb9cf9f556464c88f3424c72eee056611ff7): complete subsection reference.

- [ip_matcher](resources--service_policy--reference--group-002.md#canonical-4c542e332edb5cf8d429b69a4957680b17da5aed116384eb6ef5ade476deada4): complete subsection reference.

- [ip_prefix_list](resources--service_policy--reference--group-002.md#canonical-7851397e047a4f3ba917583a4ad6fc5ee4a2b01ae3aa02455c47e6e7654fdb6b): complete subsection reference.

- [ip_threat_category_list](resources--service_policy--reference--group-002.md#canonical-ff9627221f94bf61f1dc83c02f21f72eb8cf8f2a84877983546e41005148918f): complete subsection reference.

- [ja4_tls_fingerprint](resources--service_policy--reference--group-002.md#canonical-cc32fc3ba1c705a4610ccdebb8d6cf0387bbe00a1d0261dc592441224542d048): complete subsection reference.

- [jwt_claims](resources--service_policy--reference--group-002.md#canonical-2efe15d0d4166d6da423ccc47958f3f0b6957ad3510f247cf7e9214a87738fbd): complete subsection reference.

- [label_matcher](resources--service_policy--reference--group-002.md#canonical-ee40668264d0744e34eb8323e39f81d0d8438228103ba49d478d09c394e23c35): complete subsection reference.

<a id="canonical-7e11a4cc71b9e935dbbd6e9ae21cae7d74826a337e5af0b107b10a7b676ad778"></a>

<a id="canonical-2e572c1f078ed2aa9bed99e9638766b9ae1b11de55631546abc31814c7bc0158"></a>

## log_rule_evaluation property — rule_list.rules.spec / bd9f0a29db3f / 7

Type: `"bool"`. Optional.

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

- [mum_action](resources--service_policy--reference--group-002.md#canonical-4d8d39517872352b63a802bb206affabfc6406dfaf124f795855615d8d6bf1dd): complete subsection reference.

- [path](resources--service_policy--reference--group-002.md#canonical-d03068e450874d644162cde680757770cb480c9dc36d0fef27a4edd1333c9c61): complete subsection reference.

- [port_matcher](resources--service_policy--reference--group-002.md#canonical-ce4989f736425b8fd6af75a10db6d9fdf2608289007bb408c66d9aebb25407c1): complete subsection reference.

- [query_params](resources--service_policy--reference--group-002.md#canonical-88afe4950f729419b9a5643dcebd108607ae13bc445c5c709363fa4759b4d86d): complete subsection reference.

- [request_constraints](resources--service_policy--reference--group-002.md#canonical-5514709015ce569747f054c4df79f37e404373db06a1704cd2abd7c413577b2a): complete subsection reference.

- [segment_policy](resources--service_policy--reference--group-003.md#canonical-1f04e4ae959912d2c814638d48f95e731eea9ca0199223515f9d0c3f0f7f6e6b): complete subsection reference.

- [tls_fingerprint_matcher](resources--service_policy--reference--group-003.md#canonical-48e21c0002eed854bf2da383f1e5b04e92d13a8e1f01fc1c1b81d84a47db63e2): complete subsection reference.

- [user_identity_matcher](resources--service_policy--reference--group-003.md#canonical-c10fb3e61bf6eba8dd8ec88685f42e0566dcdd35111595de92aba354ff270132): complete subsection reference.

- [waf_action](resources--service_policy--reference--group-003.md#canonical-7c63c340e25f8ad0b323faae57eb85f89de0dcb0c9ce781239ce7d9fd9a01184): complete subsection reference.
