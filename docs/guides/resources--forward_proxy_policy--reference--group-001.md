---
page_title: "xcsh_forward_proxy_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_forward_proxy_policy reference."
---

# xcsh_forward_proxy_policy reference

<a id="canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4ad7acc2028e2315cfba796f9d94bbf3af21aaa60364bb0d27752161ea269c3"></a>

## Property reference — Property reference / 43657993bc01 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- Property reference

<a id="canonical-fb4211865e5dbd49267b53e8ae0056e2a0210db11f0bfabdaae66b728d321c84"></a>

## Direct properties — Property reference / 43657993bc01 / 3

- [allow_all](resources--forward_proxy_policy--reference--group-001.md#canonical-a89b8573386a1817d284b571978314bfef110bcffaf2120c86b4c1ed6dbf4fcf): complete subsection reference.

- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3a4b9e3835285edbd30645838b1bf42cae3bb2e7ce37b35c2f044a7dc5358ff2): complete subsection reference.

<a id="canonical-8b988d7391c72a3c223c4f79e53c2178edfa03f805b5bc5f42c29c73e3939b26"></a>

<a id="canonical-1182a575b5a4eb96c188b7025ac7ba046fdd4ab08291e2c9fbb74701e9bbf0a5"></a>

## annotations property — Property reference / 43657993bc01 / 4

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

- [any_proxy](resources--forward_proxy_policy--reference--group-001.md#canonical-beff0c6888143c7d35734369320d73ef5754f2bd58fdd8491108920308693c60): complete subsection reference.

- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-c860ce79984b910de2a0147d30cbe23a0618f33ee051a2893019776104c24d14): complete subsection reference.

<a id="canonical-9112b389aa1caa8a510645553f253c2302c9fb33e13542f6195469086853e506"></a>

<a id="canonical-8a7b9df337c9d286e643c6e79796873c1f49d44bdf7dc480847a7a347093d00e"></a>

## description property — Property reference / 43657993bc01 / 5

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

<a id="canonical-760b2130d2ae82729e8e84efe61afd4616226f3c310a45466dc730928f12b1b5"></a>

<a id="canonical-e479deeed267627e71ca980f80a487e7ca4a3595c493158ea0ed7a8a4826768a"></a>

## disable property — Property reference / 43657993bc01 / 6

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

- [drp_http_connect](resources--forward_proxy_policy--reference--group-001.md#canonical-19039d3b28b9fc9064f14a5e567384204a0907dce48350f4c5901eb5593034ad): complete subsection reference.

<a id="canonical-43f089d9e6bc5a21e5c530f4e27f5d9243cde5b9b093dd504adc8dbc7a8fba06"></a>

<a id="canonical-aff35e342330c09ae677fbbab4925020573fe446c08c30d7c89bac67a056d646"></a>

## id property — Property reference / 43657993bc01 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1e095b206672ffba466f09160a4cbd462c67d6286313fc8889507b812cf0ca95"></a>

<a id="canonical-14eeb3c0501ff532991689c7fc84db68aab2aa06528a7f16c7ae6f2b0c2c86a7"></a>

## labels property — Property reference / 43657993bc01 / 8

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

<a id="canonical-c39744ef11cbe3931e1c233f0c016b536ec92ac54e4c2a4ff5b1511fcffc4dc9"></a>

<a id="canonical-98c34896b36eacae715306e130d2db97f01b8289b7ee7659ac0e9b11e1959d8a"></a>

## name property — Property reference / 43657993bc01 / 9

Type: `"string"`. Required.

Name of the Forward Proxy Policy. Must be unique within the namespace.

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

<a id="canonical-4f4e982f72e80712a90571cab4df137eebf24df3fda0d6e5a89d480b06fd335b"></a>

<a id="canonical-6cb88197fbfa479e770db1252fdb9454d388f4464834b3fd2fa5c40c0a1f5fed"></a>

## namespace property — Property reference / 43657993bc01 / 10

Type: `"string"`. Required.

Namespace where the Forward Proxy Policy is created.

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

- [network_connector](resources--forward_proxy_policy--reference--group-001.md#canonical-4b5d2c85ac6f4f47b9125e4302b7ec53b98f8c44b1cef0836862bb19ea2b8bcf): complete subsection reference.

- [proxy_label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-29b3bae5c479042fe469e19a6bf5ed71f10838ccf5bee4ea520c0de00ff4b104): complete subsection reference.

- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18): complete subsection reference.

- [timeouts](resources--forward_proxy_policy--reference--group-002.md#canonical-0a1dba132dcaf5b5fd87600bcd2b4d3970caf054a75249a1bd899cb20db29093): complete subsection reference.

<a id="canonical-47d02940ee9bbb239f0d704608c7efc6160b2b49583027a80478afe5e844b771"></a>

## All schema paths — Property reference / 43657993bc01 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all` | [allow_all](resources--forward_proxy_policy--reference--group-001.md#canonical-19da2c92adbf3945533eee358d523b859f32c0c0b49bc37bc3c907546fdd2a5d) |
| `allow_list` | [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-75b7efd0951f8f96f7f7b9e0f74c5f4de88b288139fe315d50af03e215827f68) |
| `allow_list.default_action_allow` | [allow_list.default_action_allow](resources--forward_proxy_policy--reference--group-001.md#canonical-0bb6d76b2d8f7882602b29262a82bc2aa5aedd5163872dd9f1a00c26b78777aa) |
| `allow_list.default_action_deny` | [allow_list.default_action_deny](resources--forward_proxy_policy--reference--group-001.md#canonical-3a5a587e97bc86a9c70c8b19a18c8f18ba94450d91c7cbf04f698e08d34dc402) |
| `allow_list.default_action_next_policy` | [allow_list.default_action_next_policy](resources--forward_proxy_policy--reference--group-001.md#canonical-8ad275059c71b73dd10244811e13816fb59a1304f223eeea0575a4ea61b34a57) |
| `allow_list.dest_list` | [allow_list.dest_list](resources--forward_proxy_policy--reference--group-001.md#canonical-e10c0a89c2a75b9def604a66d174cea0867a63b100ac9ab5b1eabf725b376f6d) |
| `allow_list.dest_list.ipv6_prefixes` | [allow_list.dest_list.ipv6_prefixes](resources--forward_proxy_policy--reference--group-001.md#canonical-afce85db79582fffdf8148731863e1d39b6c05b7a02ef34fe3e25e46570d79da) |
| `allow_list.dest_list.port_ranges` | [allow_list.dest_list.port_ranges](resources--forward_proxy_policy--reference--group-001.md#canonical-812f85562b98e6ee87552ffd35af2e86cfd075e3a04f81668b0c6aa6829e5b2d) |
| `allow_list.dest_list.prefixes` | [allow_list.dest_list.prefixes](resources--forward_proxy_policy--reference--group-001.md#canonical-e783bd5ba9589aebcfee3c24ae89e250750ed0c5b9d1ada679dae90a80dbf7d5) |
| `allow_list.http_list` | [allow_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-e17f2fb535e6d2fc4e5285d26c16674f67e012f7323c85b892224914d3360256) |
| `allow_list.http_list.any_path` | [allow_list.http_list.any_path](resources--forward_proxy_policy--reference--group-001.md#canonical-a1214fc405cfb8816eeefa055a95f8fb76f9b4365f75bddbd7a632bc51e0ecd4) |
| `allow_list.http_list.exact_value` | [allow_list.http_list.exact_value](resources--forward_proxy_policy--reference--group-001.md#canonical-9a91587d982738d9417ae32b6f0ea1177a9f5f79e78ee48f7090ec118ce90879) |
| `allow_list.http_list.path_exact_value` | [allow_list.http_list.path_exact_value](resources--forward_proxy_policy--reference--group-001.md#canonical-602a2822d6ed3dde89f8f491c40de37abe1ccafcee5329af659e3d277a5c28f7) |
| `allow_list.http_list.path_prefix_value` | [allow_list.http_list.path_prefix_value](resources--forward_proxy_policy--reference--group-001.md#canonical-70aa2d2ed14ae03d3af8e9d65f559192aa40c237db4f06a322413c517579c6e6) |
| `allow_list.http_list.path_regex_value` | [allow_list.http_list.path_regex_value](resources--forward_proxy_policy--reference--group-001.md#canonical-9b7f5e2f7fdbe05e2501afce904ef11509f24b7ac381066e9a69162fbc6a47c2) |
| `allow_list.http_list.regex_value` | [allow_list.http_list.regex_value](resources--forward_proxy_policy--reference--group-001.md#canonical-b2fd85a4a3e2cf31519383d1880e5227737ec6a15d6f463fb540f0a7e3f7b249) |
| `allow_list.http_list.suffix_value` | [allow_list.http_list.suffix_value](resources--forward_proxy_policy--reference--group-001.md#canonical-ba07dbebd6a5814f4eb99258736acd6dfca16e1d5cf7ea0926fbab587ada9440) |
| `allow_list.tls_list` | [allow_list.tls_list](resources--forward_proxy_policy--reference--group-001.md#canonical-15d89fcb6006734ee65cc219dd94f644a07655297ceedcda729481f4e495d281) |
| `allow_list.tls_list.exact_value` | [allow_list.tls_list.exact_value](resources--forward_proxy_policy--reference--group-001.md#canonical-1c6338a2ab85e076e3ed5b7b170373075f635a7a996a265574f276646d319e2f) |
| `allow_list.tls_list.regex_value` | [allow_list.tls_list.regex_value](resources--forward_proxy_policy--reference--group-001.md#canonical-77a9ea63205ee57023b5c8f9a115f07f9a2dd49748a855ca7ad6a871f54f9fe3) |
| `allow_list.tls_list.suffix_value` | [allow_list.tls_list.suffix_value](resources--forward_proxy_policy--reference--group-001.md#canonical-7c2e2245f9c36e34330d8b4cbc955d2a2bc5cc1edec011638332544462044f80) |
| `annotations` | [annotations](resources--forward_proxy_policy--reference--group-001.md#canonical-8b988d7391c72a3c223c4f79e53c2178edfa03f805b5bc5f42c29c73e3939b26) |
| `any_proxy` | [any_proxy](resources--forward_proxy_policy--reference--group-001.md#canonical-e2a6c554ca578147ba3e1115cc5fa70068ff707b77a2e1303a513db4ca4d6473) |
| `deny_list` | [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3b96e71ffd51d1ca2ca17ce34ed160f322bbdb7a0c4221948e25295b9cf6ae58) |
| `deny_list.default_action_allow` | [deny_list.default_action_allow](resources--forward_proxy_policy--reference--group-001.md#canonical-75f3c9523426cdde8f9d13fd41467cffe2e4d9cdb38703d81b5c10db507ee5a4) |
| `deny_list.default_action_deny` | [deny_list.default_action_deny](resources--forward_proxy_policy--reference--group-001.md#canonical-39f9dfba2df0ced4675ac9a917188a0bbff13ca191c2650ad2c2c20331e093bb) |
| `deny_list.default_action_next_policy` | [deny_list.default_action_next_policy](resources--forward_proxy_policy--reference--group-001.md#canonical-6e17b6ed7aa2b08655f32ce14fd6be8478b54d49a12de8a29e483ae0cb245072) |
| `deny_list.dest_list` | [deny_list.dest_list](resources--forward_proxy_policy--reference--group-001.md#canonical-f95ed4a816e891bd018a1d7cbcb6a55cabd4b571a567da7bf64c1399287e748b) |
| `deny_list.dest_list.ipv6_prefixes` | [deny_list.dest_list.ipv6_prefixes](resources--forward_proxy_policy--reference--group-001.md#canonical-e2814f7f2d5c1934c68c06aaf4580d46a74d1dc471ed09f39cd4ad778884255c) |
| `deny_list.dest_list.port_ranges` | [deny_list.dest_list.port_ranges](resources--forward_proxy_policy--reference--group-001.md#canonical-7d02152c6c2c67ac9bea1c484a8c72210df839aeb01f63b4dd8324e956422201) |
| `deny_list.dest_list.prefixes` | [deny_list.dest_list.prefixes](resources--forward_proxy_policy--reference--group-001.md#canonical-19edce89be4321e0d75850f45501e731bdb59298fc2b097dd1e4d3aa01487a81) |
| `deny_list.http_list` | [deny_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-c830612a04d3e027edc5bdc31214efa1db0a92738d178e75b9758ab48edacfa2) |
| `deny_list.http_list.any_path` | [deny_list.http_list.any_path](resources--forward_proxy_policy--reference--group-001.md#canonical-a02109718d464954d586e316c363d8d50033390ba2e77b025536f54fd8994984) |
| `deny_list.http_list.exact_value` | [deny_list.http_list.exact_value](resources--forward_proxy_policy--reference--group-001.md#canonical-7d3acd01905f05e1e4a248262376d4fddb250f5a4aa69a4950a7cbbec6dee186) |
| `deny_list.http_list.path_exact_value` | [deny_list.http_list.path_exact_value](resources--forward_proxy_policy--reference--group-001.md#canonical-b82a7bc63a89554c60619e0417f35f5a4ddb7a64b7e535f3d6008846719f19e2) |
| `deny_list.http_list.path_prefix_value` | [deny_list.http_list.path_prefix_value](resources--forward_proxy_policy--reference--group-001.md#canonical-de6753d9d8c543de12998011f68b5f4d958ab118c88dc45e89a71c8bc75a64f1) |
| `deny_list.http_list.path_regex_value` | [deny_list.http_list.path_regex_value](resources--forward_proxy_policy--reference--group-001.md#canonical-89e8112b526a8dd70214d2ce0b3d639db4cd08c0380ba9045a9ba355e3e73d59) |
| `deny_list.http_list.regex_value` | [deny_list.http_list.regex_value](resources--forward_proxy_policy--reference--group-001.md#canonical-3efd3406b6dd374997a749593d51cac61176f2f4689b27b7577232bad61f50c6) |
| `deny_list.http_list.suffix_value` | [deny_list.http_list.suffix_value](resources--forward_proxy_policy--reference--group-001.md#canonical-b4c303fd00bc7e53f64d8b853598911ff55a93a480e5de5e7efb5141bfed1c19) |
| `deny_list.tls_list` | [deny_list.tls_list](resources--forward_proxy_policy--reference--group-001.md#canonical-243be2b2f2762da91d91ad6640996378d72f92e15aa926133d2cce1e003dc7e1) |
| `deny_list.tls_list.exact_value` | [deny_list.tls_list.exact_value](resources--forward_proxy_policy--reference--group-001.md#canonical-340f46541566bb5f6f1596b02b12412da3b740e06458ac4a07a485c6c1ca28e0) |
| `deny_list.tls_list.regex_value` | [deny_list.tls_list.regex_value](resources--forward_proxy_policy--reference--group-001.md#canonical-33e283642b3f68902e9e0d370c5aaf9dc19c014f2c4405a29c0b930128bac0fb) |
| `deny_list.tls_list.suffix_value` | [deny_list.tls_list.suffix_value](resources--forward_proxy_policy--reference--group-001.md#canonical-f2b52fd612c03ef31e6f85cc0f3e0536ca2a59e6c68638e76149dff197376d16) |
| `description` | [description](resources--forward_proxy_policy--reference--group-001.md#canonical-9112b389aa1caa8a510645553f253c2302c9fb33e13542f6195469086853e506) |
| `disable` | [disable](resources--forward_proxy_policy--reference--group-001.md#canonical-760b2130d2ae82729e8e84efe61afd4616226f3c310a45466dc730928f12b1b5) |
| `drp_http_connect` | [drp_http_connect](resources--forward_proxy_policy--reference--group-001.md#canonical-a781e84f04f1ff13f2fc524e738a6f5b9ea07d60b6a1ce6e4d8ae9575a46d1ba) |
| `id` | [id](resources--forward_proxy_policy--reference--group-001.md#canonical-43f089d9e6bc5a21e5c530f4e27f5d9243cde5b9b093dd504adc8dbc7a8fba06) |
| `labels` | [labels](resources--forward_proxy_policy--reference--group-001.md#canonical-1e095b206672ffba466f09160a4cbd462c67d6286313fc8889507b812cf0ca95) |
| `name` | [name](resources--forward_proxy_policy--reference--group-001.md#canonical-c39744ef11cbe3931e1c233f0c016b536ec92ac54e4c2a4ff5b1511fcffc4dc9) |
| `namespace` | [namespace](resources--forward_proxy_policy--reference--group-001.md#canonical-4f4e982f72e80712a90571cab4df137eebf24df3fda0d6e5a89d480b06fd335b) |
| `network_connector` | [network_connector](resources--forward_proxy_policy--reference--group-001.md#canonical-797227252ec719bab8cda78d85d379a025517110e77566b32dc0c0dbf2ed6e61) |
| `network_connector.name` | [network_connector.name](resources--forward_proxy_policy--reference--group-001.md#canonical-9ae3385bbac942c6e63f69bc9acbcfd7cdf5055f7ed219385c79190c811c69b6) |
| `network_connector.namespace` | [network_connector.namespace](resources--forward_proxy_policy--reference--group-001.md#canonical-364484a0fb81ee82e01b21358e7ba27bcfb1b4e69e58d786f86c8cecb2bbfc27) |
| `network_connector.tenant` | [network_connector.tenant](resources--forward_proxy_policy--reference--group-001.md#canonical-76ae06726477e7ed40fa1323dfe3490be1f6cc06edd472281c47fa7270726c73) |
| `proxy_label_selector` | [proxy_label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-d50b328df7cb876202e90fa64a5b9c48e3e38e7f2a35f2c1b681ab394ee018bd) |
| `proxy_label_selector.expressions` | [proxy_label_selector.expressions](resources--forward_proxy_policy--reference--group-001.md#canonical-e195c8bea2cbc1d030298cf75e25b41520dc0812b705e9a096122a2449af30d6) |
| `rule_list` | [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-d844d24035aacceed498e1bcc0ae8f68afd187f6014672ac9bbd6de49086a30e) |
| `rule_list.rules` | [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-f662a2d2c5224e34d30a484003a12d45d514da4413196da2c9bf683f3fdec816) |
| `rule_list.rules.action` | [rule_list.rules.action](resources--forward_proxy_policy--reference--group-001.md#canonical-64c4e8b9f92ed3d9c6aaad99a702b4bce95bb999a8ee65c131561c4349ddb40c) |
| `rule_list.rules.all_destinations` | [rule_list.rules.all_destinations](resources--forward_proxy_policy--reference--group-001.md#canonical-de3169d0b44fc2a09d0f06f25e7e137293211f9f721e316ac6b1001f7b2e4ab8) |
| `rule_list.rules.all_sources` | [rule_list.rules.all_sources](resources--forward_proxy_policy--reference--group-001.md#canonical-e6494ce847eabae046202f60260639c8bbcf8273fa4b4df9493d6313f692332b) |
| `rule_list.rules.dst_asn_list` | [rule_list.rules.dst_asn_list](resources--forward_proxy_policy--reference--group-001.md#canonical-1f74edaacdc5b8c3b7ccc9c477fa8939c260a13b2280bcd04cc30e8260f457bc) |
| `rule_list.rules.dst_asn_list.as_numbers` | [rule_list.rules.dst_asn_list.as_numbers](resources--forward_proxy_policy--reference--group-001.md#canonical-da4012dc4dd148a3b8ef1dfa91987e1df37c93ef34b170347a3d4dda0c3ee49e) |
| `rule_list.rules.dst_asn_set` | [rule_list.rules.dst_asn_set](resources--forward_proxy_policy--reference--group-001.md#canonical-bd5986b11fdf13a906ce6f44c9e41da6850bff13458d6b2246acb56cda886569) |
| `rule_list.rules.dst_asn_set.name` | [rule_list.rules.dst_asn_set.name](resources--forward_proxy_policy--reference--group-001.md#canonical-ba01ca6a79a4dcdf0e051d089c9a36a6e54017d2406a943356dc1377cecc8b01) |
| `rule_list.rules.dst_asn_set.namespace` | [rule_list.rules.dst_asn_set.namespace](resources--forward_proxy_policy--reference--group-001.md#canonical-8176d9e00336fca804c6db172f2c95ff2a5711c4b97a2f4c198615febebfbd1a) |
| `rule_list.rules.dst_asn_set.tenant` | [rule_list.rules.dst_asn_set.tenant](resources--forward_proxy_policy--reference--group-001.md#canonical-cfc5cf51217814e4c633e5000d5fdce4fdc8d271cb86b33532036341443739b2) |
| `rule_list.rules.dst_ip_prefix_set` | [rule_list.rules.dst_ip_prefix_set](resources--forward_proxy_policy--reference--group-001.md#canonical-05b9f01d05420ed1cc3f32ff8aca3e6932fa45dda741ad9002384eaa23ef0e3f) |
| `rule_list.rules.dst_ip_prefix_set.name` | [rule_list.rules.dst_ip_prefix_set.name](resources--forward_proxy_policy--reference--group-001.md#canonical-b5ae584f0073bfdf785cdfd7b327fc593a2e1197790ca06f9ada4c9e26da1763) |
| `rule_list.rules.dst_ip_prefix_set.namespace` | [rule_list.rules.dst_ip_prefix_set.namespace](resources--forward_proxy_policy--reference--group-001.md#canonical-eec21514f71be4aa70f15fe006a9649bd86eda1f0f0b4644a5687309ef0f338f) |
| `rule_list.rules.dst_ip_prefix_set.tenant` | [rule_list.rules.dst_ip_prefix_set.tenant](resources--forward_proxy_policy--reference--group-001.md#canonical-171260ac257fa30fa4292b213c545de60476396161a988418ba57554d4e92789) |
| `rule_list.rules.dst_label_selector` | [rule_list.rules.dst_label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-ff3720510c50d406e05b0a987c5b44d01ef0bd98aac6bfa27027a290fb91687b) |
| `rule_list.rules.dst_label_selector.expressions` | [rule_list.rules.dst_label_selector.expressions](resources--forward_proxy_policy--reference--group-001.md#canonical-3b55e8b60dce9864c918fe8474fe07d4a26c956b7bf2076f16c8c02843917795) |
| `rule_list.rules.dst_prefix_list` | [rule_list.rules.dst_prefix_list](resources--forward_proxy_policy--reference--group-001.md#canonical-dfc03fe3ec33e7b09edbe9d641d317c837457b2455df29d99b9b1da639bc9965) |
| `rule_list.rules.dst_prefix_list.prefixes` | [rule_list.rules.dst_prefix_list.prefixes](resources--forward_proxy_policy--reference--group-001.md#canonical-5b8674ec26e9ae7c9b43c0f9ba9b6da7438a44b92d5f1749bf5fd8e2dbc83de7) |
| `rule_list.rules.http_list` | [rule_list.rules.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-33a72667f0ab633f215f22682fe10e29de1213431974832897f0d1e0995d6818) |
| `rule_list.rules.http_list.http_list` | [rule_list.rules.http_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-ef440c63a9c94a1ad82da4429330dc7d925bebf89ebc95faf0aca44954d01748) |
| `rule_list.rules.http_list.http_list.any_path` | [rule_list.rules.http_list.http_list.any_path](resources--forward_proxy_policy--reference--group-001.md#canonical-e2e579d61f643f285426e0bc033bef2d8ccfc5440e3826faed2529b77c22eba9) |
| `rule_list.rules.http_list.http_list.exact_value` | [rule_list.rules.http_list.http_list.exact_value](resources--forward_proxy_policy--reference--group-001.md#canonical-14cf6d90f2f8bfc5c79525357ee55ac06532d1ca1cbee0ec32102bdb4203cdaa) |
| `rule_list.rules.http_list.http_list.path_exact_value` | [rule_list.rules.http_list.http_list.path_exact_value](resources--forward_proxy_policy--reference--group-001.md#canonical-4cb11e50788b7ce7135de1c2336c70752c86848f2a770ec8c1a4f87fabec10af) |
| `rule_list.rules.http_list.http_list.path_prefix_value` | [rule_list.rules.http_list.http_list.path_prefix_value](resources--forward_proxy_policy--reference--group-001.md#canonical-4b01d306210f43d10fe03fb2cd87738a41859300b6f06d744729c626100a31ed) |
| `rule_list.rules.http_list.http_list.path_regex_value` | [rule_list.rules.http_list.http_list.path_regex_value](resources--forward_proxy_policy--reference--group-001.md#canonical-b94f750128e25c436d536f54aa9dba1bc75c5da29fe7bf6b52a1b6683d59ba19) |
| `rule_list.rules.http_list.http_list.regex_value` | [rule_list.rules.http_list.http_list.regex_value](resources--forward_proxy_policy--reference--group-001.md#canonical-d36e9809c3bb040a9511ce460b66edb9767ec1603ac1a3f5c840a261eada9c75) |
| `rule_list.rules.http_list.http_list.suffix_value` | [rule_list.rules.http_list.http_list.suffix_value](resources--forward_proxy_policy--reference--group-001.md#canonical-f607b113818d1c70c53cb718153d35e783e4bb84616e9a16d37964c52db1646b) |
| `rule_list.rules.ip_prefix_set` | [rule_list.rules.ip_prefix_set](resources--forward_proxy_policy--reference--group-001.md#canonical-d92ca3f24dd8c41e3ca8e24f4e0c6dc9557aa131617896f7f8cc4607f511978b) |
| `rule_list.rules.ip_prefix_set.name` | [rule_list.rules.ip_prefix_set.name](resources--forward_proxy_policy--reference--group-001.md#canonical-8612e5c40a6b1e39be2191aa8659b24514121b29d0c81310d1a2321aba36d60b) |
| `rule_list.rules.ip_prefix_set.namespace` | [rule_list.rules.ip_prefix_set.namespace](resources--forward_proxy_policy--reference--group-001.md#canonical-3f550b1d5e0c42138768520aa4e3c54ff7d3a2f62f2bb36c8f945670334ca093) |
| `rule_list.rules.ip_prefix_set.tenant` | [rule_list.rules.ip_prefix_set.tenant](resources--forward_proxy_policy--reference--group-001.md#canonical-e0fd2f9bf17349bfda48c4773008e039b5935f9caee10e9246fd1ee47207dc89) |
| `rule_list.rules.label_selector` | [rule_list.rules.label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-d95655595aca3856d2fb5b2ce3cd843dc9aaeb96dff125fb968d0a5cfa4bb53d) |
| `rule_list.rules.label_selector.expressions` | [rule_list.rules.label_selector.expressions](resources--forward_proxy_policy--reference--group-001.md#canonical-b10f63e99ba9e8e6d9064e0b933f952f6c27584082b74106c39adb4ab72bcd79) |
| `rule_list.rules.metadata` | [rule_list.rules.metadata](resources--forward_proxy_policy--reference--group-001.md#canonical-7eb19088e59e932090034d0c8498579d081b7ab2a66a55889c306c7f76e0bd96) |
| `rule_list.rules.metadata.description_spec` | [rule_list.rules.metadata.description_spec](resources--forward_proxy_policy--reference--group-001.md#canonical-2f4d9560928864ae6e70aa49c52196c509771089d00245c8bbacd019f85a3575) |
| `rule_list.rules.metadata.name` | [rule_list.rules.metadata.name](resources--forward_proxy_policy--reference--group-001.md#canonical-6497001b44dad25042d97f8925349d5834d06d7c6840685d6013c6e62b3c85bc) |
| `rule_list.rules.no_http_connect_port` | [rule_list.rules.no_http_connect_port](resources--forward_proxy_policy--reference--group-001.md#canonical-152c76ddd00fce356625de706f83ea7ee6ed76ff5b70b8201c2db5eb93019f23) |
| `rule_list.rules.port_matcher` | [rule_list.rules.port_matcher](resources--forward_proxy_policy--reference--group-001.md#canonical-ab91c37cfa554fc02e61e134b4f2db9efb20666e8fea183c32cae2be4d22e759) |
| `rule_list.rules.port_matcher.invert_matcher` | [rule_list.rules.port_matcher.invert_matcher](resources--forward_proxy_policy--reference--group-001.md#canonical-8db0a3e0687c0dc3f4925344b60e9cfc37b557248e0f57dd12aedfa13a1af4ee) |
| `rule_list.rules.port_matcher.ports` | [rule_list.rules.port_matcher.ports](resources--forward_proxy_policy--reference--group-001.md#canonical-f18e4235ed7356e079b51f0d51bdb6e342e9c3734444795a1256944b7a1fc4f6) |
| `rule_list.rules.prefix_list` | [rule_list.rules.prefix_list](resources--forward_proxy_policy--reference--group-001.md#canonical-656bb2621fc6d4e92866af0985deda0272b739c99ee11614583f215b33f3f349) |
| `rule_list.rules.prefix_list.prefixes` | [rule_list.rules.prefix_list.prefixes](resources--forward_proxy_policy--reference--group-001.md#canonical-e24dc902338baaa1ddeef2004ee46632f3a8ec04745b45f96b798ed9099a37a7) |
| `rule_list.rules.tls_list` | [rule_list.rules.tls_list](resources--forward_proxy_policy--reference--group-001.md#canonical-64a5b9de238a8b4129979dc4dea2b4e1fc7d4d6df7481b24060dd45ee12d21a2) |
| `rule_list.rules.tls_list.tls_list` | [rule_list.rules.tls_list.tls_list](resources--forward_proxy_policy--reference--group-002.md#canonical-38a0d7f71c86d0936d8585817d2e61748feab88e31d617336e9c0fe038ef3522) |
| `rule_list.rules.tls_list.tls_list.exact_value` | [rule_list.rules.tls_list.tls_list.exact_value](resources--forward_proxy_policy--reference--group-002.md#canonical-a122bc266eba651fe3240fca7153413a642a86d9aa642610c8ace4e073b78997) |
| `rule_list.rules.tls_list.tls_list.regex_value` | [rule_list.rules.tls_list.tls_list.regex_value](resources--forward_proxy_policy--reference--group-002.md#canonical-bfdaec6f9be54cf9b0b26e13488f20214083a84d69741f2ea98113ab6a7b346f) |
| `rule_list.rules.tls_list.tls_list.suffix_value` | [rule_list.rules.tls_list.tls_list.suffix_value](resources--forward_proxy_policy--reference--group-002.md#canonical-9de97188dfde27ae9aa81f90d399368453e839b4a6e55fe8b51fcc1882aa206a) |
| `rule_list.rules.url_category_list` | [rule_list.rules.url_category_list](resources--forward_proxy_policy--reference--group-002.md#canonical-519967dc066ccee4814bea113875a753827b12b86f0140d4249eb6c0367221ec) |
| `rule_list.rules.url_category_list.url_categories` | [rule_list.rules.url_category_list.url_categories](resources--forward_proxy_policy--reference--group-002.md#canonical-a439204415809db48b3272540f93588fa1f033e75052c373d17ad693fccfb50c) |
| `timeouts` | [timeouts](resources--forward_proxy_policy--reference--group-002.md#canonical-36fe3e0d81ed9c7197e152bf6d69498796c93e2393100a0aa02999a235072f54) |
| `timeouts.create` | [timeouts.create](resources--forward_proxy_policy--reference--group-002.md#canonical-a76dd655832c1d0aef519c9ffe274d7f6bbaba8da8fa3b024cf6e174a4492e2b) |
| `timeouts.delete` | [timeouts.delete](resources--forward_proxy_policy--reference--group-002.md#canonical-30e90edf21657ac494731b546330f941755f75abbcb9cb72f496349bcd1e2ffa) |
| `timeouts.read` | [timeouts.read](resources--forward_proxy_policy--reference--group-002.md#canonical-735b7252d854f651e3dae862f6611c2a0a562d8992067b87ca04be39e735a565) |
| `timeouts.update` | [timeouts.update](resources--forward_proxy_policy--reference--group-002.md#canonical-5ff281d344f432b0876813f9cf2bed3b04c6fc22167fbf486d9ac627f176f550) |

<a id="canonical-68f656d659b8b3ffbb26a3a5d33308a7075d44d80174f2df2e165a8e29bc5526"></a>

## Next pages — Property reference / 43657993bc01 / 12

- [allow_all](resources--forward_proxy_policy--reference--group-001.md#canonical-a89b8573386a1817d284b571978314bfef110bcffaf2120c86b4c1ed6dbf4fcf)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3a4b9e3835285edbd30645838b1bf42cae3bb2e7ce37b35c2f044a7dc5358ff2)
- [any_proxy](resources--forward_proxy_policy--reference--group-001.md#canonical-beff0c6888143c7d35734369320d73ef5754f2bd58fdd8491108920308693c60)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-c860ce79984b910de2a0147d30cbe23a0618f33ee051a2893019776104c24d14)
- [drp_http_connect](resources--forward_proxy_policy--reference--group-001.md#canonical-19039d3b28b9fc9064f14a5e567384204a0907dce48350f4c5901eb5593034ad)
- [network_connector](resources--forward_proxy_policy--reference--group-001.md#canonical-4b5d2c85ac6f4f47b9125e4302b7ec53b98f8c44b1cef0836862bb19ea2b8bcf)
- [proxy_label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-29b3bae5c479042fe469e19a6bf5ed71f10838ccf5bee4ea520c0de00ff4b104)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [timeouts](resources--forward_proxy_policy--reference--group-002.md#canonical-0a1dba132dcaf5b5fd87600bcd2b4d3970caf054a75249a1bd899cb20db29093)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-a89b8573386a1817d284b571978314bfef110bcffaf2120c86b4c1ed6dbf4fcf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20931eeb8bcfa416a6af83fb4cbb7c8a922e11b5479969df76559e71693c61a8"></a>

## allow_all — allow_all / 703774269490 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- allow_all

<a id="canonical-19da2c92adbf3945533eee358d523b859f32c0c0b49bc37bc3c907546fdd2a5d"></a>

Type: `["object", {}]`. Optional.

\[OneOf: allow\_all, allow\_list, deny\_list, rule\_list\] Enable this option

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

- [allow_all](resources--forward_proxy_policy--reference--group-001.md#canonical-19da2c92adbf3945533eee358d523b859f32c0c0b49bc37bc3c907546fdd2a5d)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-75b7efd0951f8f96f7f7b9e0f74c5f4de88b288139fe315d50af03e215827f68)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3b96e71ffd51d1ca2ca17ce34ed160f322bbdb7a0c4221948e25295b9cf6ae58)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-d844d24035aacceed498e1bcc0ae8f68afd187f6014672ac9bbd6de49086a30e)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
allow_all = {}
```

<a id="canonical-47490f1f3ffd3b9424c11ca20319ee65fd04805cd732902991cb345c30019848"></a>

## Direct properties — allow_all / 703774269490 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d3b481831b83b93e4f2ba7da47a30052a3306266368f14c2dcf20cb20d678fe8"></a>

## Next pages — allow_all / 703774269490 / 4

- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-3a4b9e3835285edbd30645838b1bf42cae3bb2e7ce37b35c2f044a7dc5358ff2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1051e2c5b216d6bf7a1bf9f143e1e2a17df038e567c4abf6d6a420b4e941846"></a>

## allow_list — allow_list / 206f33cfa135 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- allow_list

<a id="canonical-75b7efd0951f8f96f7f7b9e0f74c5f4de88b288139fe315d50af03e215827f68"></a>

Type: `"object"`. single nested block, Optional.

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP).

Upstream description:

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP)

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

<a id="canonical-ab33f19398a386f4c34c7b71d412c972414ed4500734313cec466d9b4fba3610"></a>

## Direct properties — allow_list / 206f33cfa135 / 3

- [default_action_allow](resources--forward_proxy_policy--reference--group-001.md#canonical-2d4cb56b59d47ab5790fdb9802f9c291a9def4cf0c5b69ba8ab04ae8d973fa28): complete subsection reference.

- [default_action_deny](resources--forward_proxy_policy--reference--group-001.md#canonical-338bf829e144b0c5cb0c6d8f2c7cd73ee8699666d3e727d2f8977ad8325c23dd): complete subsection reference.

- [default_action_next_policy](resources--forward_proxy_policy--reference--group-001.md#canonical-122d4b2b797321f3ef6bac5f76dc6d7fe38b170c6b62578a3f68c182a83dea48): complete subsection reference.

- [dest_list](resources--forward_proxy_policy--reference--group-001.md#canonical-9b7c125bf6f07f6a41b56e264adadb177f911d69208694381ed0739d03344e18): complete subsection reference.

- [http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-b4ce9c553d7c4faa39cbe734737362a3076ee559fdf0857f78360ac1219b864d): complete subsection reference.

- [tls_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a4a2f9bc9945c2b491b6da812b8496f6d6a12f099888cc95aa8eef931f5da5d6): complete subsection reference.

<a id="canonical-786826b3c8d3fb3ce3f3e452d9919f8c512af0dd641bf750bc1f378344ee5574"></a>

## Next pages — allow_list / 206f33cfa135 / 4

- [allow_list.default_action_allow](resources--forward_proxy_policy--reference--group-001.md#canonical-2d4cb56b59d47ab5790fdb9802f9c291a9def4cf0c5b69ba8ab04ae8d973fa28)
- [allow_list.default_action_deny](resources--forward_proxy_policy--reference--group-001.md#canonical-338bf829e144b0c5cb0c6d8f2c7cd73ee8699666d3e727d2f8977ad8325c23dd)
- [allow_list.default_action_next_policy](resources--forward_proxy_policy--reference--group-001.md#canonical-122d4b2b797321f3ef6bac5f76dc6d7fe38b170c6b62578a3f68c182a83dea48)
- [allow_list.dest_list](resources--forward_proxy_policy--reference--group-001.md#canonical-9b7c125bf6f07f6a41b56e264adadb177f911d69208694381ed0739d03344e18)
- [allow_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-b4ce9c553d7c4faa39cbe734737362a3076ee559fdf0857f78360ac1219b864d)
- [allow_list.tls_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a4a2f9bc9945c2b491b6da812b8496f6d6a12f099888cc95aa8eef931f5da5d6)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-2d4cb56b59d47ab5790fdb9802f9c291a9def4cf0c5b69ba8ab04ae8d973fa28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-331842bb1c6e8bdc89046478ff9109d393ef040c430d4c6310d87e6b3164107e"></a>

## allow_list.default_action_allow — allow_list.default_action_allow / 6296ab0155e2 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3a4b9e3835285edbd30645838b1bf42cae3bb2e7ce37b35c2f044a7dc5358ff2)
- allow_list.default_action_allow

<a id="canonical-0bb6d76b2d8f7882602b29262a82bc2aa5aedd5163872dd9f1a00c26b78777aa"></a>

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

<a id="canonical-468c3eb13b341e329933b52b9d742658921b845bcc1cb0dbca443cc3e9c65f77"></a>

## Direct properties — allow_list.default_action_allow / 6296ab0155e2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1d4cc55414c4f0cab8633c59a4d09028b8d2e60d931f9118c7e165983f18788c"></a>

## Next pages — allow_list.default_action_allow / 6296ab0155e2 / 4

- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3a4b9e3835285edbd30645838b1bf42cae3bb2e7ce37b35c2f044a7dc5358ff2)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-338bf829e144b0c5cb0c6d8f2c7cd73ee8699666d3e727d2f8977ad8325c23dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7fb8a78496f40b50dbbc1534f1cdaf0a87698a0e54d3571b07e5a1febfab84d1"></a>

## allow_list.default_action_deny — allow_list.default_action_deny / 8f6c292ec09d / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3a4b9e3835285edbd30645838b1bf42cae3bb2e7ce37b35c2f044a7dc5358ff2)
- allow_list.default_action_deny

<a id="canonical-3a5a587e97bc86a9c70c8b19a18c8f18ba94450d91c7cbf04f698e08d34dc402"></a>

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

<a id="canonical-7992d8598d4e68aaee69fc49336bc76337afc4a2003f268080536822c58c0073"></a>

## Direct properties — allow_list.default_action_deny / 8f6c292ec09d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-69eb91a482c41d660ce113c18d5e734a4e8e3b06daed9f2c905e5a3ef1415918"></a>

## Next pages — allow_list.default_action_deny / 8f6c292ec09d / 4

- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3a4b9e3835285edbd30645838b1bf42cae3bb2e7ce37b35c2f044a7dc5358ff2)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-122d4b2b797321f3ef6bac5f76dc6d7fe38b170c6b62578a3f68c182a83dea48"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cecfef520ced68791648fd79f687e79daa4180bd1aa6906e9280dbec477cef3e"></a>

## allow_list.default_action_next_policy — allow_list.default_action_next_policy / 5b7c2ccc502d / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3a4b9e3835285edbd30645838b1bf42cae3bb2e7ce37b35c2f044a7dc5358ff2)
- allow_list.default_action_next_policy

<a id="canonical-8ad275059c71b73dd10244811e13816fb59a1304f223eeea0575a4ea61b34a57"></a>

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

<a id="canonical-c67030e9b1cbc5a01ce3224182fa2792dd6b45b335542b060250f13e2802d0e6"></a>

## Direct properties — allow_list.default_action_next_policy / 5b7c2ccc502d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-16d5ca8dba06bf2eaabca5f6d69e71c82312f0a9b17fae267e34b746a6b3c491"></a>

## Next pages — allow_list.default_action_next_policy / 5b7c2ccc502d / 4

- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3a4b9e3835285edbd30645838b1bf42cae3bb2e7ce37b35c2f044a7dc5358ff2)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-9b7c125bf6f07f6a41b56e264adadb177f911d69208694381ed0739d03344e18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8e3d459bb0c3d7ea559091b498b716f7d0b8c05d41dec7c32cd935bbdedffb9"></a>

## allow_list.dest_list — allow_list.dest_list / 0bb16581768b / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3a4b9e3835285edbd30645838b1bf42cae3bb2e7ce37b35c2f044a7dc5358ff2)
- allow_list.dest_list

<a id="canonical-e10c0a89c2a75b9def604a66d174cea0867a63b100ac9ab5b1eabf725b376f6d"></a>

Type: `"object"`. list nested block, Optional.

L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("port_ranges")}
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

Terraform syntax:

```terraform
dest_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3a4a65687a3010068f03d1416e61fb47036908bcefd718c0faf59f7ec1664ba6"></a>

## Direct properties — allow_list.dest_list / 0bb16581768b / 3

<a id="canonical-afce85db79582fffdf8148731863e1d39b6c05b7a02ef34fe3e25e46570d79da"></a>

<a id="canonical-ee6f5d5eeec5e59745d7225c414d10bac0517870dfd62fb8a2c4e115ab161eb9"></a>

## ipv6_prefixes property — allow_list.dest_list / 0bb16581768b / 4

Type: `["list", "string"]`. Optional.

IPv6 Prefixes. Destination IPv6 prefixes.

Upstream description:

Destination IPv6 prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-812f85562b98e6ee87552ffd35af2e86cfd075e3a04f81668b0c6aa6829e5b2d"></a>

<a id="canonical-b424507b07a27ac30b7bfedf1fee57c6c0446b764ff51f9e335c6320697bf160"></a>

## port_ranges property — allow_list.dest_list / 0bb16581768b / 5

Type: `"string"`. Optional.

String containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by '-'.

Upstream description:

A string containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="canonical-e783bd5ba9589aebcfee3c24ae89e250750ed0c5b9d1ada679dae90a80dbf7d5"></a>

<a id="canonical-409ccb4aaf1668115068dceccd914f0c1ff4f891b2c6f546607b456600fb3d4f"></a>

## prefixes property — allow_list.dest_list / 0bb16581768b / 6

Type: `["list", "string"]`. Optional.

IPv4 Prefixes. Destination IPv4 prefixes.

Upstream description:

Destination IPv4 prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-c075f8e677001f27f19cb9699c661aa41cea0c64d18a1a614e3897ecd5538de6"></a>

## Next pages — allow_list.dest_list / 0bb16581768b / 7

- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3a4b9e3835285edbd30645838b1bf42cae3bb2e7ce37b35c2f044a7dc5358ff2)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-b4ce9c553d7c4faa39cbe734737362a3076ee559fdf0857f78360ac1219b864d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa1fb69ed3215dcc0d9e6e5d0b95c3d762002ff5ff10b3be733e5e209658c9fa"></a>

## allow_list.http_list — allow_list.http_list / e3b34059b4b6 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3a4b9e3835285edbd30645838b1bf42cae3bb2e7ce37b35c2f044a7dc5358ff2)
- allow_list.http_list

<a id="canonical-e17f2fb535e6d2fc4e5285d26c16674f67e012f7323c85b892224914d3360256"></a>

Type: `"object"`. list nested block, Optional.

HTTP URLs. URLs for HTTP connections.

Upstream description:

URLs for HTTP connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_path",
    "path_exact_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_prefix_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("path_exact_value",
    "path_prefix_value"),
  validators.ConflictingListObjectAttributes("path_exact_value",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("path_prefix_value",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("regex_value",
    "suffix_value")}
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

Terraform syntax:

```terraform
http_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-7798b77676ad2f6f82f02731603012b0c4322bd2751b6cd731985bf8f239e9f6"></a>

## Direct properties — allow_list.http_list / e3b34059b4b6 / 3

- [any_path](resources--forward_proxy_policy--reference--group-001.md#canonical-4f2538845e24965dc89ce1d2d55189d66430a592114da53982e4e76b1dfef1f6): complete subsection reference.

<a id="canonical-9a91587d982738d9417ae32b6f0ea1177a9f5f79e78ee48f7090ec118ce90879"></a>

<a id="canonical-44369addad79b01f660aa5904b27bfb51321b00ab6012d8a1327f95f15491646"></a>

## exact_value property — allow_list.http_list / e3b34059b4b6 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-602a2822d6ed3dde89f8f491c40de37abe1ccafcee5329af659e3d277a5c28f7"></a>

<a id="canonical-7fff2d8e42117fa3a97df00f79b51e21581d58a355dcdf6406b51c7912b94bef"></a>

## path_exact_value property — allow_list.http_list / e3b34059b4b6 / 5

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Upstream description:

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-70aa2d2ed14ae03d3af8e9d65f559192aa40c237db4f06a322413c517579c6e6"></a>

<a id="canonical-988043d0e6925142efe9d3fe153c2e8f53630d05ece625da1865106912d4f005"></a>

## path_prefix_value property — allow_list.http_list / e3b34059b4b6 / 6

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g "/abc/xyz"
will match "/abc/xyz/.\*"

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-9b7f5e2f7fdbe05e2501afce904ef11509f24b7ac381066e9a69162fbc6a47c2"></a>

<a id="canonical-a53e70fbf4c2c7666f174bbcea0b1d00e21843cf14fad2759e97a0b310dd6619"></a>

## path_regex_value property — allow_list.http_list / e3b34059b4b6 / 7

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-b2fd85a4a3e2cf31519383d1880e5227737ec6a15d6f463fb540f0a7e3f7b249"></a>

<a id="canonical-ff22ad81fb39ea3416401b2093fabb6ff6e48fe6ede6e70d10023cf867f4cdee"></a>

## regex_value property — allow_list.http_list / e3b34059b4b6 / 8

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-ba07dbebd6a5814f4eb99258736acd6dfca16e1d5cf7ea0926fbab587ada9440"></a>

<a id="canonical-8b2f085c9bc025ea06efe2c7350e506d245894ddd00d0c540202c674ae595e4d"></a>

## suffix_value property — allow_list.http_list / e3b34059b4b6 / 9

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain names e.g "xyz.com" will match
"\*.xyz.com"

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2f7ed0b10027580f562619fecc8c7cfe75817ef366579471103a6c15252994b6"></a>

## Next pages — allow_list.http_list / e3b34059b4b6 / 10

- [allow_list.http_list.any_path](resources--forward_proxy_policy--reference--group-001.md#canonical-4f2538845e24965dc89ce1d2d55189d66430a592114da53982e4e76b1dfef1f6)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3a4b9e3835285edbd30645838b1bf42cae3bb2e7ce37b35c2f044a7dc5358ff2)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-4f2538845e24965dc89ce1d2d55189d66430a592114da53982e4e76b1dfef1f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a728bf1377f903915d667c6ae76c3e05d70961e96fa78120bf8f19caaf01cbab"></a>

## allow_list.http_list.any_path — allow_list.http_list.any_path / 45f5583e2074 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3a4b9e3835285edbd30645838b1bf42cae3bb2e7ce37b35c2f044a7dc5358ff2)
- [allow_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-b4ce9c553d7c4faa39cbe734737362a3076ee559fdf0857f78360ac1219b864d)
- allow_list.http_list.any_path

<a id="canonical-a1214fc405cfb8816eeefa055a95f8fb76f9b4365f75bddbd7a632bc51e0ecd4"></a>

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
any_path = {}
```

<a id="canonical-13cc815956b34dbf106101fdcd1591ba867e45871726a860835d25fe5035273a"></a>

## Direct properties — allow_list.http_list.any_path / 45f5583e2074 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-506037dbef2d085b26a1df086b29fe507b93547bce32691dea66d4fb89ec8cd9"></a>

## Next pages — allow_list.http_list.any_path / 45f5583e2074 / 4

- [allow_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-b4ce9c553d7c4faa39cbe734737362a3076ee559fdf0857f78360ac1219b864d)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-a4a2f9bc9945c2b491b6da812b8496f6d6a12f099888cc95aa8eef931f5da5d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9e5f556721f282847e2f3a7fe3fe162fc0aa945b1e828df9055ebe8b44b3a31"></a>

## allow_list.tls_list — allow_list.tls_list / c0815575ce35 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3a4b9e3835285edbd30645838b1bf42cae3bb2e7ce37b35c2f044a7dc5358ff2)
- allow_list.tls_list

<a id="canonical-15d89fcb6006734ee65cc219dd94f644a07655297ceedcda729481f4e495d281"></a>

Type: `"object"`. list nested block, Optional.

TLS Domains. Domains in SNI for TLS connections.

Upstream description:

Domains in SNI for TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("regex_value",
    "suffix_value")}
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

Terraform syntax:

```terraform
tls_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-5a56d66439f9c7286b0b5f6bf919f8c4268f29899c8f26aaf779530a0c61fb4b"></a>

## Direct properties — allow_list.tls_list / c0815575ce35 / 3

<a id="canonical-1c6338a2ab85e076e3ed5b7b170373075f635a7a996a265574f276646d319e2f"></a>

<a id="canonical-01638fa61c5005a60f0d374a86feb6c1b795ef4ca342e1225ae89b5a14411bc8"></a>

## exact_value property — allow_list.tls_list / c0815575ce35 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-77a9ea63205ee57023b5c8f9a115f07f9a2dd49748a855ca7ad6a871f54f9fe3"></a>

<a id="canonical-bf52849702f2e894295d5a54a95e0eb778866d75eb2baa4aa9b615194d14962e"></a>

## regex_value property — allow_list.tls_list / c0815575ce35 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-7c2e2245f9c36e34330d8b4cbc955d2a2bc5cc1edec011638332544462044f80"></a>

<a id="canonical-22e55d84ee3bec1561ea1a7deb3a925645350cdb6a97f05a520c8bc3af5f0f16"></a>

## suffix_value property — allow_list.tls_list / c0815575ce35 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1100a31d0c7940b25c0e7a0335feddfae4095133167aa31cb0f71878f817691c"></a>

## Next pages — allow_list.tls_list / c0815575ce35 / 7

- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3a4b9e3835285edbd30645838b1bf42cae3bb2e7ce37b35c2f044a7dc5358ff2)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-beff0c6888143c7d35734369320d73ef5754f2bd58fdd8491108920308693c60"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed969cfac320a721a59f52889d76caee03e9bba3365fa5d0ed781327e9d022c2"></a>

## any_proxy — any_proxy / e357add9fd55 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- any_proxy

<a id="canonical-e2a6c554ca578147ba3e1115cc5fa70068ff707b77a2e1303a513db4ca4d6473"></a>

Type: `["object", {}]`. Optional.

\[OneOf: any\_proxy, drp\_http\_connect, network\_connector, proxy\_label\_selector\] Enable this
option

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

- [any_proxy](resources--forward_proxy_policy--reference--group-001.md#canonical-e2a6c554ca578147ba3e1115cc5fa70068ff707b77a2e1303a513db4ca4d6473)
- [drp_http_connect](resources--forward_proxy_policy--reference--group-001.md#canonical-a781e84f04f1ff13f2fc524e738a6f5b9ea07d60b6a1ce6e4d8ae9575a46d1ba)
- [network_connector](resources--forward_proxy_policy--reference--group-001.md#canonical-797227252ec719bab8cda78d85d379a025517110e77566b32dc0c0dbf2ed6e61)
- [proxy_label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-d50b328df7cb876202e90fa64a5b9c48e3e38e7f2a35f2c1b681ab394ee018bd)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
any_proxy = {}
```

<a id="canonical-5d47c47868e5096302592a253716dbf3826f6c76c3b01ff65abb594ce4294715"></a>

## Direct properties — any_proxy / e357add9fd55 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f7d9efef6bd993ed4847951ff4edf126f5a13e2d2c720e7d8e8133d6afaa767d"></a>

## Next pages — any_proxy / e357add9fd55 / 4

- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-c860ce79984b910de2a0147d30cbe23a0618f33ee051a2893019776104c24d14"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d812962a752d2af8a3a1da9a78902f04b255344e6a58622fbce97080d6edeb2"></a>

## deny_list — deny_list / 07f3d79e75e6 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- deny_list

<a id="canonical-3b96e71ffd51d1ca2ca17ce34ed160f322bbdb7a0c4221948e25295b9cf6ae58"></a>

Type: `"object"`. single nested block, Optional.

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP).

Upstream description:

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP)

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

<a id="canonical-9880b94c84361fbfbfd687447697a349c43652aaea8944a8ac0a419b34f097e3"></a>

## Direct properties — deny_list / 07f3d79e75e6 / 3

- [default_action_allow](resources--forward_proxy_policy--reference--group-001.md#canonical-1a897ac77e74ec557f6c65f3ade3d68570d1ae8ffb6495c97c1ab3dafac79c1a): complete subsection reference.

- [default_action_deny](resources--forward_proxy_policy--reference--group-001.md#canonical-ad2e3c1c1674968b9f4fa8dbc3b5ec0cc9ebb9e243d355824dd1c0c71604feaa): complete subsection reference.

- [default_action_next_policy](resources--forward_proxy_policy--reference--group-001.md#canonical-641a9983c751c935ff338c784e6509dd4e97abdd7b05898dc30ddc794c23b74b): complete subsection reference.

- [dest_list](resources--forward_proxy_policy--reference--group-001.md#canonical-892e99d8a8b169a13742a61b5031b69a40c16c304dd612a928abd8dda9b5ce4e): complete subsection reference.

- [http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-e2642a6e41bdd4aafea59ae6152d0db3a8a65a9eb6b2954c0a5b13ecc15c0047): complete subsection reference.

- [tls_list](resources--forward_proxy_policy--reference--group-001.md#canonical-8c15b887882b95e4c06c0e933e46f3aecd83acaff8fd766e137c1acbc8f6ab33): complete subsection reference.

<a id="canonical-11282a8f1d48cdbcc29590815b82b0e8a31c196d34592246dfef7b894a75b588"></a>

## Next pages — deny_list / 07f3d79e75e6 / 4

- [deny_list.default_action_allow](resources--forward_proxy_policy--reference--group-001.md#canonical-1a897ac77e74ec557f6c65f3ade3d68570d1ae8ffb6495c97c1ab3dafac79c1a)
- [deny_list.default_action_deny](resources--forward_proxy_policy--reference--group-001.md#canonical-ad2e3c1c1674968b9f4fa8dbc3b5ec0cc9ebb9e243d355824dd1c0c71604feaa)
- [deny_list.default_action_next_policy](resources--forward_proxy_policy--reference--group-001.md#canonical-641a9983c751c935ff338c784e6509dd4e97abdd7b05898dc30ddc794c23b74b)
- [deny_list.dest_list](resources--forward_proxy_policy--reference--group-001.md#canonical-892e99d8a8b169a13742a61b5031b69a40c16c304dd612a928abd8dda9b5ce4e)
- [deny_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-e2642a6e41bdd4aafea59ae6152d0db3a8a65a9eb6b2954c0a5b13ecc15c0047)
- [deny_list.tls_list](resources--forward_proxy_policy--reference--group-001.md#canonical-8c15b887882b95e4c06c0e933e46f3aecd83acaff8fd766e137c1acbc8f6ab33)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-1a897ac77e74ec557f6c65f3ade3d68570d1ae8ffb6495c97c1ab3dafac79c1a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-138ffb8a17e4b68454f34040495b8ace321f11148818e0612fb22c9921aa8c23"></a>

## deny_list.default_action_allow — deny_list.default_action_allow / 7f11212b434a / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-c860ce79984b910de2a0147d30cbe23a0618f33ee051a2893019776104c24d14)
- deny_list.default_action_allow

<a id="canonical-75f3c9523426cdde8f9d13fd41467cffe2e4d9cdb38703d81b5c10db507ee5a4"></a>

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

<a id="canonical-0a5e47edc16487a837a08131494716ce371e4695cc1b3ab370ea67080ab04567"></a>

## Direct properties — deny_list.default_action_allow / 7f11212b434a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-351cda029858154ef1a50180586d8b2a64f0010f1df02de5c3dfaac6d899b882"></a>

## Next pages — deny_list.default_action_allow / 7f11212b434a / 4

- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-c860ce79984b910de2a0147d30cbe23a0618f33ee051a2893019776104c24d14)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-ad2e3c1c1674968b9f4fa8dbc3b5ec0cc9ebb9e243d355824dd1c0c71604feaa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5abf83bd158270e40997534faae20c95092c37f02d0d2fb4f1fdd44a27152321"></a>

## deny_list.default_action_deny — deny_list.default_action_deny / beb2628be774 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-c860ce79984b910de2a0147d30cbe23a0618f33ee051a2893019776104c24d14)
- deny_list.default_action_deny

<a id="canonical-39f9dfba2df0ced4675ac9a917188a0bbff13ca191c2650ad2c2c20331e093bb"></a>

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

<a id="canonical-36b02b6a006c2e3cd2a612f5f97d4300762a2d7900207c9e604fbc0be660388f"></a>

## Direct properties — deny_list.default_action_deny / beb2628be774 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7dc07e39a6f1815e88a3c509c736198d1a4bc3aab951700c46f8cd49f705018d"></a>

## Next pages — deny_list.default_action_deny / beb2628be774 / 4

- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-c860ce79984b910de2a0147d30cbe23a0618f33ee051a2893019776104c24d14)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-641a9983c751c935ff338c784e6509dd4e97abdd7b05898dc30ddc794c23b74b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a81bbef8a474c32d57468e6ddaf5296cb412802757c02469c1db918981a2391"></a>

## deny_list.default_action_next_policy — deny_list.default_action_next_policy / c22b92bbbd0d / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-c860ce79984b910de2a0147d30cbe23a0618f33ee051a2893019776104c24d14)
- deny_list.default_action_next_policy

<a id="canonical-6e17b6ed7aa2b08655f32ce14fd6be8478b54d49a12de8a29e483ae0cb245072"></a>

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

<a id="canonical-ac4e44d97346ee3b1d99b7886a1505ad8cb0ec01e263b87038eb7daf5ca23ba7"></a>

## Direct properties — deny_list.default_action_next_policy / c22b92bbbd0d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cd8fe16612be83ddb3a7b93dc414ffa882605239ffc5e856b09b6db2011421a4"></a>

## Next pages — deny_list.default_action_next_policy / c22b92bbbd0d / 4

- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-c860ce79984b910de2a0147d30cbe23a0618f33ee051a2893019776104c24d14)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-892e99d8a8b169a13742a61b5031b69a40c16c304dd612a928abd8dda9b5ce4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-697873a710c1f198f0a2d341ab3b69d766bc021b0c4eabe1555618c2463f8cdd"></a>

## deny_list.dest_list — deny_list.dest_list / 3268cdd42600 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-c860ce79984b910de2a0147d30cbe23a0618f33ee051a2893019776104c24d14)
- deny_list.dest_list

<a id="canonical-f95ed4a816e891bd018a1d7cbcb6a55cabd4b571a567da7bf64c1399287e748b"></a>

Type: `"object"`. list nested block, Optional.

L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("port_ranges")}
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

Terraform syntax:

```terraform
dest_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-9c69cc96a75929218e1ef46109fca3239136bb245c513ccd6f21091a7b54e983"></a>

## Direct properties — deny_list.dest_list / 3268cdd42600 / 3

<a id="canonical-e2814f7f2d5c1934c68c06aaf4580d46a74d1dc471ed09f39cd4ad778884255c"></a>

<a id="canonical-20203d5809e58c1f0c25b0808fe8df4c1b43a1dc0b9ada78e9116e66d67d49e9"></a>

## ipv6_prefixes property — deny_list.dest_list / 3268cdd42600 / 4

Type: `["list", "string"]`. Optional.

IPv6 Prefixes. Destination IPv6 prefixes.

Upstream description:

Destination IPv6 prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-7d02152c6c2c67ac9bea1c484a8c72210df839aeb01f63b4dd8324e956422201"></a>

<a id="canonical-ccd5c772e534d06c0c9e9de9d6df765cffcf82835fc9b5415c1552c6c5fa070d"></a>

## port_ranges property — deny_list.dest_list / 3268cdd42600 / 5

Type: `"string"`. Optional.

String containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by '-'.

Upstream description:

A string containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="canonical-19edce89be4321e0d75850f45501e731bdb59298fc2b097dd1e4d3aa01487a81"></a>

<a id="canonical-47af30c5f39060fcc8c3b6adf041f02f1edb8253ad4944aa0a326cd99c874c57"></a>

## prefixes property — deny_list.dest_list / 3268cdd42600 / 6

Type: `["list", "string"]`. Optional.

IPv4 Prefixes. Destination IPv4 prefixes.

Upstream description:

Destination IPv4 prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1063f16ff274786e19c5a36382f6a0ceb0f7889d8d2f393c1843dbccda4d019b"></a>

## Next pages — deny_list.dest_list / 3268cdd42600 / 7

- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-c860ce79984b910de2a0147d30cbe23a0618f33ee051a2893019776104c24d14)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-e2642a6e41bdd4aafea59ae6152d0db3a8a65a9eb6b2954c0a5b13ecc15c0047"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-452ad0695dbdae23f787b5223c53f3425ff9ab43b644a31af9bdb1e4ef261837"></a>

## deny_list.http_list — deny_list.http_list / 5e29f7bd686d / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-c860ce79984b910de2a0147d30cbe23a0618f33ee051a2893019776104c24d14)
- deny_list.http_list

<a id="canonical-c830612a04d3e027edc5bdc31214efa1db0a92738d178e75b9758ab48edacfa2"></a>

Type: `"object"`. list nested block, Optional.

HTTP URLs. URLs for HTTP connections.

Upstream description:

URLs for HTTP connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_path",
    "path_exact_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_prefix_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("path_exact_value",
    "path_prefix_value"),
  validators.ConflictingListObjectAttributes("path_exact_value",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("path_prefix_value",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("regex_value",
    "suffix_value")}
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

Terraform syntax:

```terraform
http_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-159f8884bebca7e3adbf536bbafc92f432652497fbb3ebd661dcd8310799931c"></a>

## Direct properties — deny_list.http_list / 5e29f7bd686d / 3

- [any_path](resources--forward_proxy_policy--reference--group-001.md#canonical-29656825c867dc243a6eebd930b6da9c20d53b3f29d5b9736f5ec2b990155562): complete subsection reference.

<a id="canonical-7d3acd01905f05e1e4a248262376d4fddb250f5a4aa69a4950a7cbbec6dee186"></a>

<a id="canonical-dbe40829644bdcdaa7dde3bed1d2e1b99d3349f9e60c8e0f1bddc73086166b10"></a>

## exact_value property — deny_list.http_list / 5e29f7bd686d / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-b82a7bc63a89554c60619e0417f35f5a4ddb7a64b7e535f3d6008846719f19e2"></a>

<a id="canonical-5426cbcf7038f5ca45485a65b640c6c8dc90c2de65b084a5393b6e6ae889fb9d"></a>

## path_exact_value property — deny_list.http_list / 5e29f7bd686d / 5

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Upstream description:

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-de6753d9d8c543de12998011f68b5f4d958ab118c88dc45e89a71c8bc75a64f1"></a>

<a id="canonical-ec8040f328b4b072197bde325172e871359792d5995c4f299439eff324ccc471"></a>

## path_prefix_value property — deny_list.http_list / 5e29f7bd686d / 6

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g "/abc/xyz"
will match "/abc/xyz/.\*"

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-89e8112b526a8dd70214d2ce0b3d639db4cd08c0380ba9045a9ba355e3e73d59"></a>

<a id="canonical-9377b52a4eac6cfa99f923848ce78f2ffe54b56586f6e3f1102c316907172bc1"></a>

## path_regex_value property — deny_list.http_list / 5e29f7bd686d / 7

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3efd3406b6dd374997a749593d51cac61176f2f4689b27b7577232bad61f50c6"></a>

<a id="canonical-b524f8f8601be0d030e0dc6f1707f36e37d60c89b70fa16473797ae27d795106"></a>

## regex_value property — deny_list.http_list / 5e29f7bd686d / 8

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-b4c303fd00bc7e53f64d8b853598911ff55a93a480e5de5e7efb5141bfed1c19"></a>

<a id="canonical-67c28b66c0755463231ad1cd47b0c312760c41c96cd7599f39766e14aa877f41"></a>

## suffix_value property — deny_list.http_list / 5e29f7bd686d / 9

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain names e.g "xyz.com" will match
"\*.xyz.com"

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-adabd679cc426cfe268fbdffee9e17e4b981ed40395ba77d766728a40ec7aa72"></a>

## Next pages — deny_list.http_list / 5e29f7bd686d / 10

- [deny_list.http_list.any_path](resources--forward_proxy_policy--reference--group-001.md#canonical-29656825c867dc243a6eebd930b6da9c20d53b3f29d5b9736f5ec2b990155562)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-c860ce79984b910de2a0147d30cbe23a0618f33ee051a2893019776104c24d14)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-29656825c867dc243a6eebd930b6da9c20d53b3f29d5b9736f5ec2b990155562"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e841cfcdf7ca55dc58bed31ada4dbe0189de046d79807a5edd0d1474c058b47"></a>

## deny_list.http_list.any_path — deny_list.http_list.any_path / 0723c51f3329 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-c860ce79984b910de2a0147d30cbe23a0618f33ee051a2893019776104c24d14)
- [deny_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-e2642a6e41bdd4aafea59ae6152d0db3a8a65a9eb6b2954c0a5b13ecc15c0047)
- deny_list.http_list.any_path

<a id="canonical-a02109718d464954d586e316c363d8d50033390ba2e77b025536f54fd8994984"></a>

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
any_path = {}
```

<a id="canonical-aeb06795ef897ed0f600591fd44d1f70e45ef269f413b6bb7a44e877b4561fcc"></a>

## Direct properties — deny_list.http_list.any_path / 0723c51f3329 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c300e95e0e13435f3eb759e56248dbe8f2fc36f451f4c3dfe6ea1c0d292900f4"></a>

## Next pages — deny_list.http_list.any_path / 0723c51f3329 / 4

- [deny_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-e2642a6e41bdd4aafea59ae6152d0db3a8a65a9eb6b2954c0a5b13ecc15c0047)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-8c15b887882b95e4c06c0e933e46f3aecd83acaff8fd766e137c1acbc8f6ab33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ac19b42632d5aba7325c0071b93d9cb7ae5cb58717709b91983320e703f9436"></a>

## deny_list.tls_list — deny_list.tls_list / d110b3434206 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-c860ce79984b910de2a0147d30cbe23a0618f33ee051a2893019776104c24d14)
- deny_list.tls_list

<a id="canonical-243be2b2f2762da91d91ad6640996378d72f92e15aa926133d2cce1e003dc7e1"></a>

Type: `"object"`. list nested block, Optional.

TLS Domains. Domains in SNI for TLS connections.

Upstream description:

Domains in SNI for TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("regex_value",
    "suffix_value")}
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

Terraform syntax:

```terraform
tls_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-64f711dec1db7d823794c021558d8bcf405db8e3d972fc20fdddd6c9be8f24db"></a>

## Direct properties — deny_list.tls_list / d110b3434206 / 3

<a id="canonical-340f46541566bb5f6f1596b02b12412da3b740e06458ac4a07a485c6c1ca28e0"></a>

<a id="canonical-f51fb41d8f63a378710e43974e0ecee35fef4ae3a1fa135404cf608310c063b7"></a>

## exact_value property — deny_list.tls_list / d110b3434206 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-33e283642b3f68902e9e0d370c5aaf9dc19c014f2c4405a29c0b930128bac0fb"></a>

<a id="canonical-cbb16746ce9715c6b46efeb2077f07a5367a82222e8017ca98858118b506c91e"></a>

## regex_value property — deny_list.tls_list / d110b3434206 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-f2b52fd612c03ef31e6f85cc0f3e0536ca2a59e6c68638e76149dff197376d16"></a>

<a id="canonical-736903796e4f553935130b078b37d5dc8783d7015ce08ae69dc95c94e7a4679b"></a>

## suffix_value property — deny_list.tls_list / d110b3434206 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3618307db18d4feed70c666832b93fe64f57005778579cfc855aea599701e585"></a>

## Next pages — deny_list.tls_list / d110b3434206 / 7

- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-c860ce79984b910de2a0147d30cbe23a0618f33ee051a2893019776104c24d14)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-19039d3b28b9fc9064f14a5e567384204a0907dce48350f4c5901eb5593034ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ce25563fab84fac549ea3687b176fd587a233490d3119153c4ee957c3760161"></a>

## drp_http_connect — drp_http_connect / 9e3688781851 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- drp_http_connect

<a id="canonical-a781e84f04f1ff13f2fc524e738a6f5b9ea07d60b6a1ce6e4d8ae9575a46d1ba"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for drp http connect.

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
drp_http_connect = {}
```

<a id="canonical-1ede9fdb77527858f116e193a39d97516b66e2d65f9e94b209b095a766798388"></a>

## Direct properties — drp_http_connect / 9e3688781851 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-982401e537fa33b2f0cc2038d58020a95f9b02aac2f6e21fe18ee233e3c09834"></a>

## Next pages — drp_http_connect / 9e3688781851 / 4

- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-4b5d2c85ac6f4f47b9125e4302b7ec53b98f8c44b1cef0836862bb19ea2b8bcf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ceb3aa8e5bf56c81a5e87eecb75bc03d6d497c6f80d79d9615cc4b118589d213"></a>

## network_connector — network_connector / 2c4a3bf4512a / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- network_connector

<a id="canonical-797227252ec719bab8cda78d85d379a025517110e77566b32dc0c0dbf2ed6e61"></a>

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
network_connector {
  # Configure direct properties listed below.
}
```

<a id="canonical-0d7c19d1a9dfc5acca4c73a8e178b96bdfeb46953fdbbeea0dba2337fd365adb"></a>

## Direct properties — network_connector / 2c4a3bf4512a / 3

<a id="canonical-9ae3385bbac942c6e63f69bc9acbcfd7cdf5055f7ed219385c79190c811c69b6"></a>

<a id="canonical-96ad57bb79d34968f0cc43f9573ebc3d94ce6df9d98676ceb1cbb2f207ab2ff0"></a>

## name property — network_connector / 2c4a3bf4512a / 4

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

<a id="canonical-364484a0fb81ee82e01b21358e7ba27bcfb1b4e69e58d786f86c8cecb2bbfc27"></a>

<a id="canonical-2a4ecb99d730fb4bbe855277fbffa50cf8d6a5253e3b58855f303bdb7a3723da"></a>

## namespace property — network_connector / 2c4a3bf4512a / 5

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

<a id="canonical-76ae06726477e7ed40fa1323dfe3490be1f6cc06edd472281c47fa7270726c73"></a>

<a id="canonical-0152e0771842ad2c686ff2125ae8307684feb8b3b05807ab8b93f26e6d7b9d9f"></a>

## tenant property — network_connector / 2c4a3bf4512a / 6

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

<a id="canonical-d0b9fa342287173bd64ef705c843d74550b17d5560b29b3f7ef823e425b13eb4"></a>

## Next pages — network_connector / 2c4a3bf4512a / 7

- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-29b3bae5c479042fe469e19a6bf5ed71f10838ccf5bee4ea520c0de00ff4b104"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5d97d034b798920871f31dcbcc77121e77c87536bba513a21372ab4c3b8a9f8"></a>

## proxy_label_selector — proxy_label_selector / eaa6f93c659f / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- proxy_label_selector

<a id="canonical-d50b328df7cb876202e90fa64a5b9c48e3e38e7f2a35f2c1b681ab394ee018bd"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
proxy_label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-4349fd2fbf46cc6eff750de50b79f8108870a83128c680946abce360fd682dcb"></a>

## Direct properties — proxy_label_selector / eaa6f93c659f / 3

<a id="canonical-e195c8bea2cbc1d030298cf75e25b41520dc0812b705e9a096122a2449af30d6"></a>

<a id="canonical-10bf17de3357bcf375db60a48c9565f97300538d653827a79658e4bf38b34b2d"></a>

## expressions property — proxy_label_selector / eaa6f93c659f / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-b2b94a2a4155d4ee19841896198ec1589721fa0584f283e2fb36d812c436c99b"></a>

## Next pages — proxy_label_selector / eaa6f93c659f / 5

- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab08d00b77e326b730f94a079acf6a7337d2f45dc88630c99dc9c08e49eac6ce"></a>

## rule_list — rule_list / facb7e758e5d / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- rule_list

<a id="canonical-d844d24035aacceed498e1bcc0ae8f68afd187f6014672ac9bbd6de49086a30e"></a>

Type: `"object"`. single nested block, Optional.

Custom Rule List. List of custom rules.

Upstream description:

List of custom rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
rule_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-11882022b67e5fd97ce654655e9fb4010eb4147b0c16f18a938b6adea9344457"></a>

## Direct properties — rule_list / facb7e758e5d / 3

- [rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc): complete subsection reference.

<a id="canonical-5f3f091ccaf9981cfeece13af573cc7984f28e92e9120dd36a9a25d70fb9636c"></a>

## Next pages — rule_list / facb7e758e5d / 4

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0e8f60bb2471a89094a7109720219bdb7e6f176aa5016cab4134772f14c280a"></a>

## rule_list.rules — rule_list.rules / f99972ac1b3b / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- rule_list.rules

<a id="canonical-f662a2d2c5224e34d30a484003a12d45d514da4413196da2c9bf683f3fdec816"></a>

Type: `"object"`. list nested block, Optional.

Custom Rule List. List of custom rules.

Upstream description:

List of custom rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_destinations",
    "dst_asn_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "dst_asn_set"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "dst_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "dst_label_selector"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "dst_prefix_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "http_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "tls_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("all_sources",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_sources",
    "label_selector"),
  validators.ConflictingListObjectAttributes("all_sources",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "dst_asn_set"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "dst_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "dst_label_selector"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "dst_prefix_list"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "http_list"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "tls_list"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("dst_asn_set",
    "dst_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("dst_asn_set",
    "dst_label_selector"),
  validators.ConflictingListObjectAttributes("dst_asn_set",
    "dst_prefix_list"),
  validators.ConflictingListObjectAttributes("dst_asn_set",
    "http_list"),
  validators.ConflictingListObjectAttributes("dst_asn_set",
    "tls_list"),
  validators.ConflictingListObjectAttributes("dst_asn_set",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("dst_ip_prefix_set",
    "dst_label_selector"),
  validators.ConflictingListObjectAttributes("dst_ip_prefix_set",
    "dst_prefix_list"),
  validators.ConflictingListObjectAttributes("dst_ip_prefix_set",
    "http_list"),
  validators.ConflictingListObjectAttributes("dst_ip_prefix_set",
    "tls_list"),
  validators.ConflictingListObjectAttributes("dst_ip_prefix_set",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("dst_label_selector",
    "dst_prefix_list"),
  validators.ConflictingListObjectAttributes("dst_label_selector",
    "http_list"),
  validators.ConflictingListObjectAttributes("dst_label_selector",
    "tls_list"),
  validators.ConflictingListObjectAttributes("dst_label_selector",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("dst_prefix_list",
    "http_list"),
  validators.ConflictingListObjectAttributes("dst_prefix_list",
    "tls_list"),
  validators.ConflictingListObjectAttributes("dst_prefix_list",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("http_list",
    "tls_list"),
  validators.ConflictingListObjectAttributes("http_list",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "label_selector"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("label_selector",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("no_http_connect_port",
    "port_matcher"),
  validators.ConflictingListObjectAttributes("tls_list",
    "url_category_list")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
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

<a id="canonical-0480df52e278cbd3c3f6fe9bb09390b9c76a11707ee9aa66f1f798b1c0840a74"></a>

## Direct properties — rule_list.rules / f99972ac1b3b / 3

<a id="canonical-64c4e8b9f92ed3d9c6aaad99a702b4bce95bb999a8ee65c131561c4349ddb40c"></a>

<a id="canonical-4882a6bacf2a336bd0584d423c02e883c4451f34a8be6ea0bb92a00a763a9c2c"></a>

## action property — rule_list.rules / f99972ac1b3b / 4

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [all_destinations](resources--forward_proxy_policy--reference--group-001.md#canonical-f398aece0a0bea317a2473769e6223fc99774bc81301df5486f1ac0f4d8a5d59): complete subsection reference.

- [all_sources](resources--forward_proxy_policy--reference--group-001.md#canonical-aa08121155baec77189095375af4ae784ff2648f902d41a6aa232f0f73c4abe1): complete subsection reference.

- [dst_asn_list](resources--forward_proxy_policy--reference--group-001.md#canonical-854b7de59daf1afa2fb8680ae7a68709307cb40d3211f95a00ca412832f1bd0f): complete subsection reference.

- [dst_asn_set](resources--forward_proxy_policy--reference--group-001.md#canonical-03791326db696f64eaedb0a48158c08c21f461af5eb2595030264827d12d20a3): complete subsection reference.

- [dst_ip_prefix_set](resources--forward_proxy_policy--reference--group-001.md#canonical-a64bc1078d6deaf11d98a040af00692e1501de10e835c6cedba5f42eb3ae2d65): complete subsection reference.

- [dst_label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-94a1603b6c467fbffbab4538cc2f8c24d4553687177acd973e1dbe4bcb5a7a56): complete subsection reference.

- [dst_prefix_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a81dd4d4396cc72265ecd4f9d845428ab6c284def3527c96b88422a5ef82f725): complete subsection reference.

- [http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a90bb02130242c9bd4823a01656427a321a68a4165df69df34b0b6097d393a28): complete subsection reference.

- [ip_prefix_set](resources--forward_proxy_policy--reference--group-001.md#canonical-34a4cd379ea098cbae54a59cfefe5657cd3cfd7f38913217ec9240823966dc53): complete subsection reference.

- [label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-961cf0783971b01c0f6f92b65b2f286ecb6c412122fe7bb0569c60e9dabce4a2): complete subsection reference.

- [metadata](resources--forward_proxy_policy--reference--group-001.md#canonical-cc5c15f579c3bcbe63294fac5df526eedea4052d3f96ccf1fb4eb60e285f6570): complete subsection reference.

- [no_http_connect_port](resources--forward_proxy_policy--reference--group-001.md#canonical-f4f449d4d19f00eac9609e433f2ff7c98ecc6a857e8e3aac397cd8e0c3ed3ba9): complete subsection reference.

- [port_matcher](resources--forward_proxy_policy--reference--group-001.md#canonical-740620282a98c98dd4c6f9f53db619726af163173ea5d80e00014c734002d239): complete subsection reference.

- [prefix_list](resources--forward_proxy_policy--reference--group-001.md#canonical-367463b61d53b97f1bd41185937156c288550181e5b7fed2d5d5a0526fad1708): complete subsection reference.

- [tls_list](resources--forward_proxy_policy--reference--group-001.md#canonical-39a0f8b5db164c87e09341ae6ccd864e837f7f0cf9e256ada0b7c3e7a50705c3): complete subsection reference.

- [url_category_list](resources--forward_proxy_policy--reference--group-002.md#canonical-df60768b4402f4efd096157d5efcfd1041e5802507dd445736e3aef70d2a8cae): complete subsection reference.

<a id="canonical-73680368164646bfe01cd89b2549b536204f481ff26bc7c8bcd7b94407d62344"></a>

## Next pages — rule_list.rules / f99972ac1b3b / 5

- [rule_list.rules.all_destinations](resources--forward_proxy_policy--reference--group-001.md#canonical-f398aece0a0bea317a2473769e6223fc99774bc81301df5486f1ac0f4d8a5d59)
- [rule_list.rules.all_sources](resources--forward_proxy_policy--reference--group-001.md#canonical-aa08121155baec77189095375af4ae784ff2648f902d41a6aa232f0f73c4abe1)
- [rule_list.rules.dst_asn_list](resources--forward_proxy_policy--reference--group-001.md#canonical-854b7de59daf1afa2fb8680ae7a68709307cb40d3211f95a00ca412832f1bd0f)
- [rule_list.rules.dst_asn_set](resources--forward_proxy_policy--reference--group-001.md#canonical-03791326db696f64eaedb0a48158c08c21f461af5eb2595030264827d12d20a3)
- [rule_list.rules.dst_ip_prefix_set](resources--forward_proxy_policy--reference--group-001.md#canonical-a64bc1078d6deaf11d98a040af00692e1501de10e835c6cedba5f42eb3ae2d65)
- [rule_list.rules.dst_label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-94a1603b6c467fbffbab4538cc2f8c24d4553687177acd973e1dbe4bcb5a7a56)
- [rule_list.rules.dst_prefix_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a81dd4d4396cc72265ecd4f9d845428ab6c284def3527c96b88422a5ef82f725)
- [rule_list.rules.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a90bb02130242c9bd4823a01656427a321a68a4165df69df34b0b6097d393a28)
- [rule_list.rules.ip_prefix_set](resources--forward_proxy_policy--reference--group-001.md#canonical-34a4cd379ea098cbae54a59cfefe5657cd3cfd7f38913217ec9240823966dc53)
- [rule_list.rules.label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-961cf0783971b01c0f6f92b65b2f286ecb6c412122fe7bb0569c60e9dabce4a2)
- [rule_list.rules.metadata](resources--forward_proxy_policy--reference--group-001.md#canonical-cc5c15f579c3bcbe63294fac5df526eedea4052d3f96ccf1fb4eb60e285f6570)
- [rule_list.rules.no_http_connect_port](resources--forward_proxy_policy--reference--group-001.md#canonical-f4f449d4d19f00eac9609e433f2ff7c98ecc6a857e8e3aac397cd8e0c3ed3ba9)
- [rule_list.rules.port_matcher](resources--forward_proxy_policy--reference--group-001.md#canonical-740620282a98c98dd4c6f9f53db619726af163173ea5d80e00014c734002d239)
- [rule_list.rules.prefix_list](resources--forward_proxy_policy--reference--group-001.md#canonical-367463b61d53b97f1bd41185937156c288550181e5b7fed2d5d5a0526fad1708)
- [rule_list.rules.tls_list](resources--forward_proxy_policy--reference--group-001.md#canonical-39a0f8b5db164c87e09341ae6ccd864e837f7f0cf9e256ada0b7c3e7a50705c3)
- [rule_list.rules.url_category_list](resources--forward_proxy_policy--reference--group-002.md#canonical-df60768b4402f4efd096157d5efcfd1041e5802507dd445736e3aef70d2a8cae)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-f398aece0a0bea317a2473769e6223fc99774bc81301df5486f1ac0f4d8a5d59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99c3c3800bfff106a6a30c8f0e256043efec8aecc706786d9dd813ea44f7ed8c"></a>

## rule_list.rules.all_destinations — rule_list.rules.all_destinations / 8995851df3fc / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- rule_list.rules.all_destinations

<a id="canonical-de3169d0b44fc2a09d0f06f25e7e137293211f9f721e316ac6b1001f7b2e4ab8"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all destinations.

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
all_destinations = {}
```

<a id="canonical-4f9ecb139ebea3009a8d155f713c635cf24766080553b1225962aed401ce9809"></a>

## Direct properties — rule_list.rules.all_destinations / 8995851df3fc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-14ce70585f52ca26fd2a56e2131c9ac573af844a711ef8a065b3073d75d2b29a"></a>

## Next pages — rule_list.rules.all_destinations / 8995851df3fc / 4

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-aa08121155baec77189095375af4ae784ff2648f902d41a6aa232f0f73c4abe1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc1061c48e3417b4dae03502cb7b0f8e04c087cea6ed77947e1922af56a0ab38"></a>

## rule_list.rules.all_sources — rule_list.rules.all_sources / 6f5e56677da0 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- rule_list.rules.all_sources

<a id="canonical-e6494ce847eabae046202f60260639c8bbcf8273fa4b4df9493d6313f692332b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all sources.

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
all_sources = {}
```

<a id="canonical-7a5a195c9170a5e55c18ec36744baeba59e370eec605947d2dd1ba5d9dcd2ae2"></a>

## Direct properties — rule_list.rules.all_sources / 6f5e56677da0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5a8ab5fb5db907fa7e8bd3cf8392fe93db1b5b06e28e9a4dba18001d419a15e2"></a>

## Next pages — rule_list.rules.all_sources / 6f5e56677da0 / 4

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-854b7de59daf1afa2fb8680ae7a68709307cb40d3211f95a00ca412832f1bd0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-837b1b9c6006f0be631cd21e5b05d1675c4e9bdf1d6ac600e00dc71fa82bc26e"></a>

## rule_list.rules.dst_asn_list — rule_list.rules.dst_asn_list / 78d5861db25d / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- rule_list.rules.dst_asn_list

<a id="canonical-1f74edaacdc5b8c3b7ccc9c477fa8939c260a13b2280bcd04cc30e8260f457bc"></a>

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
dst_asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-150db85bb2b43d06ebb3ffe323c4b8345fd37e02c6a51fdd44b926970e8d0937"></a>

## Direct properties — rule_list.rules.dst_asn_list / 78d5861db25d / 3

<a id="canonical-da4012dc4dd148a3b8ef1dfa91987e1df37c93ef34b170347a3d4dda0c3ee49e"></a>

<a id="canonical-15145c2dd4a3d59e4396f7c9d208747640b6618f36f162781904a0092a0d822e"></a>

## as_numbers property — rule_list.rules.dst_asn_list / 78d5861db25d / 4

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

<a id="canonical-d52b420d5a1058d94cf1c4034bb84c7fed89f7b96d5e29f2e9d65237fb8bae71"></a>

## Next pages — rule_list.rules.dst_asn_list / 78d5861db25d / 5

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-03791326db696f64eaedb0a48158c08c21f461af5eb2595030264827d12d20a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-945f93d685072d8e203da205c6e7ab08c2622708394cb54afae1ebd1e6d18dc6"></a>

## rule_list.rules.dst_asn_set — rule_list.rules.dst_asn_set / 4986adc8a40c / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- rule_list.rules.dst_asn_set

<a id="canonical-bd5986b11fdf13a906ce6f44c9e41da6850bff13458d6b2246acb56cda886569"></a>

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
dst_asn_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-a4cfa34bd70f334f5dbccb89cc98c9437a6d680f0fc10c723c96ed1593df05e6"></a>

## Direct properties — rule_list.rules.dst_asn_set / 4986adc8a40c / 3

<a id="canonical-ba01ca6a79a4dcdf0e051d089c9a36a6e54017d2406a943356dc1377cecc8b01"></a>

<a id="canonical-c55af7a10ee7bc6d7d8e5c3c988ef8fbde8aa79368f86bd460de00f75a391b98"></a>

## name property — rule_list.rules.dst_asn_set / 4986adc8a40c / 4

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

<a id="canonical-8176d9e00336fca804c6db172f2c95ff2a5711c4b97a2f4c198615febebfbd1a"></a>

<a id="canonical-a70c91b1250f815425e3c09c74b6e3e9cc1ffc81450dde0be740b7ef055d3f4e"></a>

## namespace property — rule_list.rules.dst_asn_set / 4986adc8a40c / 5

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

<a id="canonical-cfc5cf51217814e4c633e5000d5fdce4fdc8d271cb86b33532036341443739b2"></a>

<a id="canonical-2936037d242e3eb6511dd1318b1190882252b46b4a58720234e9ea10cf49abfd"></a>

## tenant property — rule_list.rules.dst_asn_set / 4986adc8a40c / 6

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

<a id="canonical-c70f80b23e47ce5dae1355f11c61b3189e314f25eaa76092e633b7c0487fe762"></a>

## Next pages — rule_list.rules.dst_asn_set / 4986adc8a40c / 7

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-a64bc1078d6deaf11d98a040af00692e1501de10e835c6cedba5f42eb3ae2d65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c805591fb9415002f805ec64c48ea0b045a4b4782c3b5aeb81be91e975238cf1"></a>

## rule_list.rules.dst_ip_prefix_set — rule_list.rules.dst_ip_prefix_set / 673764f0f8f8 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- rule_list.rules.dst_ip_prefix_set

<a id="canonical-05b9f01d05420ed1cc3f32ff8aca3e6932fa45dda741ad9002384eaa23ef0e3f"></a>

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
dst_ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-d510e0b8f2808a50318d4f734335230019334d51fb82307db13168797db525f3"></a>

## Direct properties — rule_list.rules.dst_ip_prefix_set / 673764f0f8f8 / 3

<a id="canonical-b5ae584f0073bfdf785cdfd7b327fc593a2e1197790ca06f9ada4c9e26da1763"></a>

<a id="canonical-f64a647682782db973d304ec8d59f4d47b3689e7ee874012316af1346d0005c0"></a>

## name property — rule_list.rules.dst_ip_prefix_set / 673764f0f8f8 / 4

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

<a id="canonical-eec21514f71be4aa70f15fe006a9649bd86eda1f0f0b4644a5687309ef0f338f"></a>

<a id="canonical-fb387802f6ae89b04eb0e252549c7ae94bba402dd1a691ce500b99ae1415a5fb"></a>

## namespace property — rule_list.rules.dst_ip_prefix_set / 673764f0f8f8 / 5

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

<a id="canonical-171260ac257fa30fa4292b213c545de60476396161a988418ba57554d4e92789"></a>

<a id="canonical-1c1c9844796618bbd161a8afe7a0179ba0ec7f5051fb1d6a7198d9b444642111"></a>

## tenant property — rule_list.rules.dst_ip_prefix_set / 673764f0f8f8 / 6

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

<a id="canonical-8b3cfc7f0ad53f2cffede0ac1781bdcc8681de03df6afbe946265ca7a65fb58c"></a>

## Next pages — rule_list.rules.dst_ip_prefix_set / 673764f0f8f8 / 7

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-94a1603b6c467fbffbab4538cc2f8c24d4553687177acd973e1dbe4bcb5a7a56"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc48eeaebc0c2d44a704f6d6293d5a071bad5e47e1a8b7e32fbceaa2e12eddbd"></a>

## rule_list.rules.dst_label_selector — rule_list.rules.dst_label_selector / 196f5fac7b8d / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- rule_list.rules.dst_label_selector

<a id="canonical-ff3720510c50d406e05b0a987c5b44d01ef0bd98aac6bfa27027a290fb91687b"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
dst_label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-c8706cd81369502ecebd2ff6ca64967c9d79770dd8a725591708a3ebdcce2506"></a>

## Direct properties — rule_list.rules.dst_label_selector / 196f5fac7b8d / 3

<a id="canonical-3b55e8b60dce9864c918fe8474fe07d4a26c956b7bf2076f16c8c02843917795"></a>

<a id="canonical-43d35300fcf2a33e1abf83634bdf8c27a6bc35aece65132a5647048727c52bbc"></a>

## expressions property — rule_list.rules.dst_label_selector / 196f5fac7b8d / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-27d5b1add00030058eda9bbb606af3c5489bb80689634945bdbf2605b118bd49"></a>

## Next pages — rule_list.rules.dst_label_selector / 196f5fac7b8d / 5

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-a81dd4d4396cc72265ecd4f9d845428ab6c284def3527c96b88422a5ef82f725"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50809047e152de25c5c11e2974c707123fac3b911f3ab80d95c158ecb7f12588"></a>

## rule_list.rules.dst_prefix_list — rule_list.rules.dst_prefix_list / 1f6473606281 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- rule_list.rules.dst_prefix_list

<a id="canonical-dfc03fe3ec33e7b09edbe9d641d317c837457b2455df29d99b9b1da639bc9965"></a>

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
dst_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-f4600431bf19f307d4e9a4271ee929bd9409ae95b366bb61cc9b7d0bb355df14"></a>

## Direct properties — rule_list.rules.dst_prefix_list / 1f6473606281 / 3

<a id="canonical-5b8674ec26e9ae7c9b43c0f9ba9b6da7438a44b92d5f1749bf5fd8e2dbc83de7"></a>

<a id="canonical-0bb5fe1a64917c6ec483bd0a31bc06bb1ee3158a2bcf4d512cfd53597c9a6cbf"></a>

## prefixes property — rule_list.rules.dst_prefix_list / 1f6473606281 / 4

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

<a id="canonical-c7948c62ba7523a261b83ebbd4e3610ced8eb649956d5e5021e40349b77b21db"></a>

## Next pages — rule_list.rules.dst_prefix_list / 1f6473606281 / 5

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-a90bb02130242c9bd4823a01656427a321a68a4165df69df34b0b6097d393a28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02e38decd4d12a12e5797f7a154fcf50cb769b873a4d1e7aad2f477d7bcc7d0d"></a>

## rule_list.rules.http_list — rule_list.rules.http_list / e1c7f01746ea / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- rule_list.rules.http_list

<a id="canonical-33a72667f0ab633f215f22682fe10e29de1213431974832897f0d1e0995d6818"></a>

Type: `"object"`. single nested block, Optional.

URLListType.

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
http_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-16e6b52432f6b27d830b02e5d1aa32d97c951f4db5f5a60e0e24a43350a6e8f5"></a>

## Direct properties — rule_list.rules.http_list / e1c7f01746ea / 3

- [http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-afdc5d6116f5ebb572db771ce755f1f429b32a7da94d91b22000d7d177cd939e): complete subsection reference.

<a id="canonical-9ab5dd2d7ddbb9e8250f2206e2e077d5c814966e0852f35abe42f6630ac53595"></a>

## Next pages — rule_list.rules.http_list / e1c7f01746ea / 4

- [rule_list.rules.http_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-afdc5d6116f5ebb572db771ce755f1f429b32a7da94d91b22000d7d177cd939e)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-afdc5d6116f5ebb572db771ce755f1f429b32a7da94d91b22000d7d177cd939e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a538e529f68e952608513c215761ed9f1fcf1fd3a3113874e5b0db44ef82a1a"></a>

## rule_list.rules.http_list.http_list — rule_list.rules.http_list.http_list / b8b144503a58 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [rule_list.rules.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a90bb02130242c9bd4823a01656427a321a68a4165df69df34b0b6097d393a28)
- rule_list.rules.http_list.http_list

<a id="canonical-ef440c63a9c94a1ad82da4429330dc7d925bebf89ebc95faf0aca44954d01748"></a>

Type: `"object"`. list nested block, Optional.

HTTP URLs. URLs for HTTP connections.

Upstream description:

URLs for HTTP connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_path",
    "path_exact_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_prefix_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("path_exact_value",
    "path_prefix_value"),
  validators.ConflictingListObjectAttributes("path_exact_value",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("path_prefix_value",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("regex_value",
    "suffix_value")}
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

Terraform syntax:

```terraform
http_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-ed8a80afeedc2d163ed3af6a86923cb4289e8605bfa02a3ea11490fd69307867"></a>

## Direct properties — rule_list.rules.http_list.http_list / b8b144503a58 / 3

- [any_path](resources--forward_proxy_policy--reference--group-001.md#canonical-833dbf39984ca343041d6deb3fee4ead0a9dbb1df999e62bd2d2494843f5f87f): complete subsection reference.

<a id="canonical-14cf6d90f2f8bfc5c79525357ee55ac06532d1ca1cbee0ec32102bdb4203cdaa"></a>

<a id="canonical-ac33f4f01da2404e39c584b0adea645bd4b7b22005566c8d0313142223a45269"></a>

## exact_value property — rule_list.rules.http_list.http_list / b8b144503a58 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-4cb11e50788b7ce7135de1c2336c70752c86848f2a770ec8c1a4f87fabec10af"></a>

<a id="canonical-f9f23f706c51fc432805f5bd1db45afe855242a604c5e23df4fe4c9c95e137bb"></a>

## path_exact_value property — rule_list.rules.http_list.http_list / b8b144503a58 / 5

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Upstream description:

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-4b01d306210f43d10fe03fb2cd87738a41859300b6f06d744729c626100a31ed"></a>

<a id="canonical-65544445ef734226f7f98c2a41fb7b16a8a65b02bf492067e6a0eff1759d9deb"></a>

## path_prefix_value property — rule_list.rules.http_list.http_list / b8b144503a58 / 6

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g "/abc/xyz"
will match "/abc/xyz/.\*"

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-b94f750128e25c436d536f54aa9dba1bc75c5da29fe7bf6b52a1b6683d59ba19"></a>

<a id="canonical-4345982c96271afed1e12c17fe2df58cd0ec6cd2e020db56acab7f6c0d09330f"></a>

## path_regex_value property — rule_list.rules.http_list.http_list / b8b144503a58 / 7

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-d36e9809c3bb040a9511ce460b66edb9767ec1603ac1a3f5c840a261eada9c75"></a>

<a id="canonical-3617a54a3a2f6d25ae8a98a4a0e4aad2f46814dd9b550fd178c2a13710d287de"></a>

## regex_value property — rule_list.rules.http_list.http_list / b8b144503a58 / 8

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-f607b113818d1c70c53cb718153d35e783e4bb84616e9a16d37964c52db1646b"></a>

<a id="canonical-cdd4700b11380165aff686cb630521a63a793b1d46e78a66456dd67e89d7e81a"></a>

## suffix_value property — rule_list.rules.http_list.http_list / b8b144503a58 / 9

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain names e.g "xyz.com" will match
"\*.xyz.com"

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-6959aa11e77de942599ed4c6a3c267c0a197a3e5b1be8bd55ac9ee73688dbca3"></a>

## Next pages — rule_list.rules.http_list.http_list / b8b144503a58 / 10

- [rule_list.rules.http_list.http_list.any_path](resources--forward_proxy_policy--reference--group-001.md#canonical-833dbf39984ca343041d6deb3fee4ead0a9dbb1df999e62bd2d2494843f5f87f)
- [rule_list.rules.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a90bb02130242c9bd4823a01656427a321a68a4165df69df34b0b6097d393a28)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-833dbf39984ca343041d6deb3fee4ead0a9dbb1df999e62bd2d2494843f5f87f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-216ae2676fe2c67ff64eb06a564192de0771e0b745cd060350c78b189b1f6723"></a>

## rule_list.rules.http_list.http_list.any_path — rule_list.rules.http_list.http_list.any_path / b75d33558272 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [rule_list.rules.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a90bb02130242c9bd4823a01656427a321a68a4165df69df34b0b6097d393a28)
- [rule_list.rules.http_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-afdc5d6116f5ebb572db771ce755f1f429b32a7da94d91b22000d7d177cd939e)
- rule_list.rules.http_list.http_list.any_path

<a id="canonical-e2e579d61f643f285426e0bc033bef2d8ccfc5440e3826faed2529b77c22eba9"></a>

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
any_path = {}
```

<a id="canonical-b8eaf3e2e3fa07c299837bf7525e138ef11b68e91cdcc082a24120499f29bd49"></a>

## Direct properties — rule_list.rules.http_list.http_list.any_path / b75d33558272 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e71d0848b90863b9f497729d037bd75908da4839e8e2e853caecda30b54c0490"></a>

## Next pages — rule_list.rules.http_list.http_list.any_path / b75d33558272 / 4

- [rule_list.rules.http_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-afdc5d6116f5ebb572db771ce755f1f429b32a7da94d91b22000d7d177cd939e)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-34a4cd379ea098cbae54a59cfefe5657cd3cfd7f38913217ec9240823966dc53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28a70020f67233f232b5642cd099e92556708d2f67c27d5a03e3aa3ebeb27cbf"></a>

## rule_list.rules.ip_prefix_set — rule_list.rules.ip_prefix_set / 8d3ccc997c05 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- rule_list.rules.ip_prefix_set

<a id="canonical-d92ca3f24dd8c41e3ca8e24f4e0c6dc9557aa131617896f7f8cc4607f511978b"></a>

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
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-68743aa4de1fb9d464e77319f9ac50a9421584e509f7a8a3bc12ffcec65475fd"></a>

## Direct properties — rule_list.rules.ip_prefix_set / 8d3ccc997c05 / 3

<a id="canonical-8612e5c40a6b1e39be2191aa8659b24514121b29d0c81310d1a2321aba36d60b"></a>

<a id="canonical-3b48e5b1e10c231dbdedf173905d167dab60efabc3e471877f7ac6842e382823"></a>

## name property — rule_list.rules.ip_prefix_set / 8d3ccc997c05 / 4

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

<a id="canonical-3f550b1d5e0c42138768520aa4e3c54ff7d3a2f62f2bb36c8f945670334ca093"></a>

<a id="canonical-8d42114b3c6a524848c6543dfd8d02cf1c08b8e734874352c4afada1f928d8f5"></a>

## namespace property — rule_list.rules.ip_prefix_set / 8d3ccc997c05 / 5

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

<a id="canonical-e0fd2f9bf17349bfda48c4773008e039b5935f9caee10e9246fd1ee47207dc89"></a>

<a id="canonical-134a1701054e210bfb3cf75cd0943c91f434958eeb371f322efe681748f37503"></a>

## tenant property — rule_list.rules.ip_prefix_set / 8d3ccc997c05 / 6

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

<a id="canonical-18dea6a2a9f46c209385fc29f3104d5a4eb8e2acf4b2d4fc14d37d38b245a8d7"></a>

## Next pages — rule_list.rules.ip_prefix_set / 8d3ccc997c05 / 7

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-961cf0783971b01c0f6f92b65b2f286ecb6c412122fe7bb0569c60e9dabce4a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf2510149bbca237d0507cf95991217c2e9e918a966227fc071dc89f5b97b4a1"></a>

## rule_list.rules.label_selector — rule_list.rules.label_selector / d17f812a0c49 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- rule_list.rules.label_selector

<a id="canonical-d95655595aca3856d2fb5b2ce3cd843dc9aaeb96dff125fb968d0a5cfa4bb53d"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-1e197d97e319b81f717a915bb8838df4b44127588bc5fc3d6b5cecc28ccb6dde"></a>

## Direct properties — rule_list.rules.label_selector / d17f812a0c49 / 3

<a id="canonical-b10f63e99ba9e8e6d9064e0b933f952f6c27584082b74106c39adb4ab72bcd79"></a>

<a id="canonical-9ad9bb39708261006cd299b9d30f0c49e8554010526c5766569ffa79dffa0fca"></a>

## expressions property — rule_list.rules.label_selector / d17f812a0c49 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-dff3b14a510d6fc426b7877ec2d35a26d1a25b1ffdf1a2683e4dd163716d5757"></a>

## Next pages — rule_list.rules.label_selector / d17f812a0c49 / 5

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-cc5c15f579c3bcbe63294fac5df526eedea4052d3f96ccf1fb4eb60e285f6570"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bc268e0ef70f5fc33e02ee57a937a7ca0fb0471e4e97fdd8746af2c986723ea7"></a>

## rule_list.rules.metadata — rule_list.rules.metadata / 45efa8755ae4 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- rule_list.rules.metadata

<a id="canonical-7eb19088e59e932090034d0c8498579d081b7ab2a66a55889c306c7f76e0bd96"></a>

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

<a id="canonical-f9424ef218c6cfc62305e5df83160966c82a2a01da995a731a02b3794597e827"></a>

## Direct properties — rule_list.rules.metadata / 45efa8755ae4 / 3

<a id="canonical-2f4d9560928864ae6e70aa49c52196c509771089d00245c8bbacd019f85a3575"></a>

<a id="canonical-f226e930fe321242c93ab76e01126a7c7f1357fe356fd26a7156a23768c489b3"></a>

## description_spec property — rule_list.rules.metadata / 45efa8755ae4 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-6497001b44dad25042d97f8925349d5834d06d7c6840685d6013c6e62b3c85bc"></a>

<a id="canonical-9e169a36e77a35cf3d57474fabfcec2dac63941c326f3434b6c9e140f5f14269"></a>

## name property — rule_list.rules.metadata / 45efa8755ae4 / 5

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

<a id="canonical-fea005397ebf34950bc069495c7b1d993712bc3a1df36751953a222ce2e21c7d"></a>

## Next pages — rule_list.rules.metadata / 45efa8755ae4 / 6

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-f4f449d4d19f00eac9609e433f2ff7c98ecc6a857e8e3aac397cd8e0c3ed3ba9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e51fbf04bd70b8f1b481d40668d789eb1cc9c02a490061f93f5ad102fee622d"></a>

## rule_list.rules.no_http_connect_port — rule_list.rules.no_http_connect_port / 6d79b72dce85 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- rule_list.rules.no_http_connect_port

<a id="canonical-152c76ddd00fce356625de706f83ea7ee6ed76ff5b70b8201c2db5eb93019f23"></a>

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
no_http_connect_port = {}
```

<a id="canonical-e8a2de2c2db7bfd077ef12c4ab52eb9da4ee6f25a3abe3ad1e0caf8bbca18630"></a>

## Direct properties — rule_list.rules.no_http_connect_port / 6d79b72dce85 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ab09cd63960de02a86687459ddd5a5738bb3a2040925e1b66908fb0a7534bd74"></a>

## Next pages — rule_list.rules.no_http_connect_port / 6d79b72dce85 / 4

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-740620282a98c98dd4c6f9f53db619726af163173ea5d80e00014c734002d239"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1f07e2feb7c58ff7d2145bbc88e50f901d7f38b3ae707adba5eaca6353ab3b1"></a>

## rule_list.rules.port_matcher — rule_list.rules.port_matcher / d7cc3ddabf62 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- rule_list.rules.port_matcher

<a id="canonical-ab91c37cfa554fc02e61e134b4f2db9efb20666e8fea183c32cae2be4d22e759"></a>

Type: `"object"`. single nested block, Optional.

Port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true.

Upstream description:

A port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ports")}
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
port_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-6a89b30aea95895e8c01a9578c270eb093f6bf6089a6a6b6169abf0a82eda711"></a>

## Direct properties — rule_list.rules.port_matcher / d7cc3ddabf62 / 3

<a id="canonical-8db0a3e0687c0dc3f4925344b60e9cfc37b557248e0f57dd12aedfa13a1af4ee"></a>

<a id="canonical-36a015d7f65016de979925b99c70013733ae59d13035f25eb228f76a3f30b7da"></a>

## invert_matcher property — rule_list.rules.port_matcher / d7cc3ddabf62 / 4

Type: `"bool"`. Optional.

Invert Port Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-f18e4235ed7356e079b51f0d51bdb6e342e9c3734444795a1256944b7a1fc4f6"></a>

<a id="canonical-51ea770d9297f20119c83b73fa00c2bfdba4c716a51d0ad899337685ee2a33b4"></a>

## ports property — rule_list.rules.port_matcher / d7cc3ddabf62 / 5

Type: `["list", "string"]`. Optional.

List of strings, each of which is a single port value or a tuple of start and end port values
separated by '-'. The start and end values are considered to be part of the range.

Upstream description:

A list of strings, each of which is a single port value or a tuple of start and end port values
separated by "-". The start and end values are considered to be part of the range.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-952a9ba00cf1ae40237ed942ee7d6ac49a6c8bb50acb41234a3fb825af9d7b72"></a>

## Next pages — rule_list.rules.port_matcher / d7cc3ddabf62 / 6

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-367463b61d53b97f1bd41185937156c288550181e5b7fed2d5d5a0526fad1708"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4420ebcdcef1db4634b5eec49c816a78cec544eace81ee955411a62d3ed405e5"></a>

## rule_list.rules.prefix_list — rule_list.rules.prefix_list / d8a4a82648a5 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- rule_list.rules.prefix_list

<a id="canonical-656bb2621fc6d4e92866af0985deda0272b739c99ee11614583f215b33f3f349"></a>

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

<a id="canonical-70eaaba5ba027a565a487292dc4a9474376b4b37761df2056018df550ca4e355"></a>

## Direct properties — rule_list.rules.prefix_list / d8a4a82648a5 / 3

<a id="canonical-e24dc902338baaa1ddeef2004ee46632f3a8ec04745b45f96b798ed9099a37a7"></a>

<a id="canonical-a7d46542e57964cdfb29c89249f5bac00411331ef1a4704f493d5f1e868d26ae"></a>

## prefixes property — rule_list.rules.prefix_list / d8a4a82648a5 / 4

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

<a id="canonical-7bce3af1285df4cc645bfb5579413be7726015a0d6f9c8a46a453381b949c6c8"></a>

## Next pages — rule_list.rules.prefix_list / d8a4a82648a5 / 5

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-39a0f8b5db164c87e09341ae6ccd864e837f7f0cf9e256ada0b7c3e7a50705c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d11b044b9adb49998a75ba536821a86b0202ebc65a927731cf0c242fdabacea5"></a>

## rule_list.rules.tls_list — rule_list.rules.tls_list / e03fdde72b81 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- rule_list.rules.tls_list

<a id="canonical-64a5b9de238a8b4129979dc4dea2b4e1fc7d4d6df7481b24060dd45ee12d21a2"></a>

Type: `"object"`. single nested block, Optional.

DomainListType.

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
tls_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-426f9e7af0c91430aa35655a21902d88329a0aaba64f9ffc449a17d9d37e065d"></a>

## Direct properties — rule_list.rules.tls_list / e03fdde72b81 / 3

- [tls_list](resources--forward_proxy_policy--reference--group-001.md#canonical-5feca49f0de47172bfc1d5656e31349dc9d7a5eebd0c7188da426e08a85d0cc4): complete subsection reference.

<a id="canonical-5133540ce83fb89c495b14bf4230ad1d44656c2c682d47ef6cce666e53133954"></a>

## Next pages — rule_list.rules.tls_list / e03fdde72b81 / 4

- [rule_list.rules.tls_list.tls_list](resources--forward_proxy_policy--reference--group-001.md#canonical-5feca49f0de47172bfc1d5656e31349dc9d7a5eebd0c7188da426e08a85d0cc4)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-5feca49f0de47172bfc1d5656e31349dc9d7a5eebd0c7188da426e08a85d0cc4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
