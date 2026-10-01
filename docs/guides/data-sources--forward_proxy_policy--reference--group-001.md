---
page_title: "xcsh_forward_proxy_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_forward_proxy_policy reference."
---

# xcsh_forward_proxy_policy reference

<a id="canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59e450a9e594ceba4925d2de32263ec1cffb9465a2c7f055007bcc5b3b4fe252"></a>

## Property reference — Property reference / 471006cbf4c8 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- Property reference

<a id="canonical-ae0de4c20f7a27bdeb64b76a4de7a90928ae3df2146400aaacc531e9f632b941"></a>

## Direct properties — Property reference / 471006cbf4c8 / 3

- [allow_all](data-sources--forward_proxy_policy--reference--group-001.md#canonical-8716c7f4145bacfb52b96bf63e6d7b975d9d1978e5d1a7400c8ba561138be577): complete subsection reference.

- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d183f0efaebceda4460734ecc8a76f3c16f106da3f79a8ef564fe822923add71): complete subsection reference.

<a id="canonical-fecb7809b443e80403d56ad6ce75cab524e45b4fddabae11e8a72bf17b7a3c63"></a>

<a id="canonical-2e4d0c82907da4718dc0583df12a2c921dd5ec7a3107a739af7c749080e22388"></a>

## annotations property — Property reference / 471006cbf4c8 / 4

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

- [any_proxy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-31c9ad9033679a59af087236658e34f64f1b9398597f5b5d69f98c2253e73c13): complete subsection reference.

- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-9a659a7dd15f7b619c2f958e3e0a23bf6c005723edae551e451c2685739cc28e): complete subsection reference.

<a id="canonical-843a50ac043353adb0e3001c171952be4c4f3ccd18cc6532d7c9454d0ae9e86b"></a>

<a id="canonical-7507adb7612517f58e52e2bbaa256c08d349e5d68c85232f8b9c4f88e9f93533"></a>

## description property — Property reference / 471006cbf4c8 / 5

Type: `"string"`. Computed.

Description of the ForwardProxyPolicy.

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

- [drp_http_connect](data-sources--forward_proxy_policy--reference--group-001.md#canonical-313a9458d3817074ebbbda207e7f37f67bd5c77441fec57a2ca64f72e0f99452): complete subsection reference.

<a id="canonical-76713c79b8dbe027f3b14627eabae98bd9c63b1242eb5f03236af1b37991d843"></a>

<a id="canonical-91322924db3466b3f9f8dc273302d14241eac1bb7d28e35473fd162348bd577b"></a>

## id property — Property reference / 471006cbf4c8 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-4605186a4629d8c8e38218ab93ee9a524a6cc4578d541b40bcbeda28a827cec0"></a>

<a id="canonical-f17b0758293213659b3692c8e365ec371e3fc0796c9614171c21514d2210338e"></a>

## labels property — Property reference / 471006cbf4c8 / 7

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

<a id="canonical-8147c9af331f89ce96a0698941955283812753241d5dc025b5bb18569557e8fb"></a>

<a id="canonical-e3292e0c769668afed5b3e3b19616e8c424aea68ec2079e3bb5f3bc2628552b0"></a>

## name property — Property reference / 471006cbf4c8 / 8

Type: `"string"`. Required.

Name of the ForwardProxyPolicy.

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

<a id="canonical-8b43a16d0827908e319123b33be0a22bf201e2aed05fc73fd44f35e552db21c0"></a>

<a id="canonical-05829389a629dfe09596d0b73da5d887fc9c349125cc3e63c7939463faeafde5"></a>

## namespace property — Property reference / 471006cbf4c8 / 9

Type: `"string"`. Required.

Namespace where the ForwardProxyPolicy exists.

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

- [network_connector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-c36323cc0a7cfcdf49d229fd178c1ff4108821e0a2eacb7b04c9fd02f7cd6a65): complete subsection reference.

- [proxy_label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-4ec4f3f751c2ec7e14eb8d0c4ee6682326e0434c5001b90a0d5df49631571d53): complete subsection reference.

- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e): complete subsection reference.

<a id="canonical-88299e8ca6391dbdbe9a5fc7faa631f7cfdfb5a0e1b692940529fa2ae9028cc4"></a>

## All schema paths — Property reference / 471006cbf4c8 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all` | [allow_all](data-sources--forward_proxy_policy--reference--group-001.md#canonical-a81233db54e5cb34cba0a21ba10d27c7fc42eeed8d2601a53b663f4597353f66) |
| `allow_list` | [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-5d27abb54733201c10bc480eeac20b34b4553ccfc3a00192e403602b56dcce1b) |
| `allow_list.default_action_allow` | [allow_list.default_action_allow](data-sources--forward_proxy_policy--reference--group-001.md#canonical-8a6998a7ebec33cd6bbf206435e99b7ca8eb7004ab246e6f90d26143dc969453) |
| `allow_list.default_action_deny` | [allow_list.default_action_deny](data-sources--forward_proxy_policy--reference--group-001.md#canonical-cc43099840bc57d90ec8af371c029dea849a0743a838bff3943eb9ee38e1d9bc) |
| `allow_list.default_action_next_policy` | [allow_list.default_action_next_policy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-84d27d0516c4292c1ce6f3f53cdc6a46a29722065e3b59a241c4c2a01a523a0e) |
| `allow_list.dest_list` | [allow_list.dest_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e7848c259b2e5c745919fa924a46c7c0ed4a65cd4b732ecad38b91a5a824c3c7) |
| `allow_list.dest_list.ipv6_prefixes` | [allow_list.dest_list.ipv6_prefixes](data-sources--forward_proxy_policy--reference--group-001.md#canonical-dbcd89a2369ad30f310ecdcd971a54a935c7fd5622d691b44d94da6b8d7b1f65) |
| `allow_list.dest_list.port_ranges` | [allow_list.dest_list.port_ranges](data-sources--forward_proxy_policy--reference--group-001.md#canonical-68c230d1d8d332d7239d1b334f266dd92f75ba4aa43fe766a562aeaf682181fa) |
| `allow_list.dest_list.prefixes` | [allow_list.dest_list.prefixes](data-sources--forward_proxy_policy--reference--group-001.md#canonical-ede810db665677d090f97af7cfbf3b0faede00cfd1ca2675131b4d36a1211db4) |
| `allow_list.http_list` | [allow_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-c08c3262a51e3f717c50fc90967f35a907cb8f1464a1b3cdac99035a91846cc5) |
| `allow_list.http_list.any_path` | [allow_list.http_list.any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-77398e6ed822ec383637e4ba47ba64f5120ffb5c2fcdd904e43ad138e4424792) |
| `allow_list.http_list.exact_value` | [allow_list.http_list.exact_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-fa51edac2ebca69d87ad32f4ac441d1cc879ae664cb8b3a955c00bba9a9e7dfd) |
| `allow_list.http_list.path_exact_value` | [allow_list.http_list.path_exact_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-5e778423a2849f7148e092d767266f5089f3274266de134d810576f46422b8cc) |
| `allow_list.http_list.path_prefix_value` | [allow_list.http_list.path_prefix_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-4ab5c3ed14a8ec6425de3e01e7a550b74bc6c8937cfd71c4d55e557842b8b8ff) |
| `allow_list.http_list.path_regex_value` | [allow_list.http_list.path_regex_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-4de8928a6e787a1d3b12a6ee7f6f22fcd5b278f8ea0ea446fd11f0d1e8ccb0f2) |
| `allow_list.http_list.regex_value` | [allow_list.http_list.regex_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e5d5f00c31f73bb5692a221241d577468a92d447364378ed62b69956794d7f77) |
| `allow_list.http_list.suffix_value` | [allow_list.http_list.suffix_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-ba6cd8ffbcd098306a37109e96ee39204d8f008606bc10e10445c2655ac5d00d) |
| `allow_list.tls_list` | [allow_list.tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-743f01c933974bf17f856dac48548fc0ceb17290e9b45c8f6ede454775909e6a) |
| `allow_list.tls_list.exact_value` | [allow_list.tls_list.exact_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-b314ae255f36ac6db5dcc5d161328701363a518a51b8099e8cd2913d1e44a5f1) |
| `allow_list.tls_list.regex_value` | [allow_list.tls_list.regex_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-ce096b871ae4d9c107f464c3b855eba984f9997a3c627cda276e77ad7609cd1b) |
| `allow_list.tls_list.suffix_value` | [allow_list.tls_list.suffix_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-fa2926bc6d049db5731554817636403feda899ee0f1e9e1ddadf846dd999ca67) |
| `annotations` | [annotations](data-sources--forward_proxy_policy--reference--group-001.md#canonical-fecb7809b443e80403d56ad6ce75cab524e45b4fddabae11e8a72bf17b7a3c63) |
| `any_proxy` | [any_proxy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-dbdc995e9c52efce3354b50fd6e69b97e554a51587f6e1550efd4077c23931f9) |
| `deny_list` | [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-41ecd20c3b29261610c585dc9f9e691426a85daf0eea1d64a3f8710278decdb1) |
| `deny_list.default_action_allow` | [deny_list.default_action_allow](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e26dfb58d1e9d2d75beba7791c410bdfc9e8bf0434642e4fdb525898140d58f0) |
| `deny_list.default_action_deny` | [deny_list.default_action_deny](data-sources--forward_proxy_policy--reference--group-001.md#canonical-965488b0aa8bfb864aec63d50cb1d4c49695963a87e4b930e068155230cf133e) |
| `deny_list.default_action_next_policy` | [deny_list.default_action_next_policy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-f714e05f35e35d559de67b09f21e16535d05db697eb448938dca5719e7de2217) |
| `deny_list.dest_list` | [deny_list.dest_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-bd0cf8dc7270a777f6350932af75ffade0ba09c2b956253cc6c239df43ca50f6) |
| `deny_list.dest_list.ipv6_prefixes` | [deny_list.dest_list.ipv6_prefixes](data-sources--forward_proxy_policy--reference--group-001.md#canonical-bc10d319fd26b8bc60f2a621c1d3e568db7639961cb7648295222a9b42f72861) |
| `deny_list.dest_list.port_ranges` | [deny_list.dest_list.port_ranges](data-sources--forward_proxy_policy--reference--group-001.md#canonical-58d8c5d9f395de01067216a1e0ad3db0ad40807942b43df89ac50ad220c48d70) |
| `deny_list.dest_list.prefixes` | [deny_list.dest_list.prefixes](data-sources--forward_proxy_policy--reference--group-001.md#canonical-11eeb721a367e797f8ba92fe6d2f2fb8448e0ff5f8bcdc04c4b65aaa663048cd) |
| `deny_list.http_list` | [deny_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-fe3123633aa8fe730ac80f4dc384cbc71d6a3a92ea90d18e258e35c0ef6c9b25) |
| `deny_list.http_list.any_path` | [deny_list.http_list.any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-ec7ba3459fa31c28866928a9e057d363a5d0ebf1058cc71823b759bee455bceb) |
| `deny_list.http_list.exact_value` | [deny_list.http_list.exact_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-06b3221b3cc6ed7afc0ce6961d859975a121f4187a7c95ced53e77f7369a0466) |
| `deny_list.http_list.path_exact_value` | [deny_list.http_list.path_exact_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2521fd081dfc8609ceea506bbbbfa6b30a884b41c089c293aba8138879edf16d) |
| `deny_list.http_list.path_prefix_value` | [deny_list.http_list.path_prefix_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-c8f791c2067142cf66911322ac25f2097ede848e1f2478f3d10ac719873bfcaf) |
| `deny_list.http_list.path_regex_value` | [deny_list.http_list.path_regex_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-53ae9296c6aefc3d09debf2c20f54ac0b2c2f6fba7c0d4430297c6efc2504cab) |
| `deny_list.http_list.regex_value` | [deny_list.http_list.regex_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-465ba2570b04fb2e1d7d9d2f9da0225574a8c65dcf6a0dcd316764e7b10301d3) |
| `deny_list.http_list.suffix_value` | [deny_list.http_list.suffix_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-84e8d6dd2d222d76b35bebdfb946395ca4e27a75557cb4afa9f7f7ff3c856d16) |
| `deny_list.tls_list` | [deny_list.tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-4072d3b1df34177fbdeadf0bfd77a3faad51511e49a5af6050b3da91f34c9647) |
| `deny_list.tls_list.exact_value` | [deny_list.tls_list.exact_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-a0e6e01063b1d77e71c6c22e43b417775e617f598c14f9eda27dc766a352c439) |
| `deny_list.tls_list.regex_value` | [deny_list.tls_list.regex_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-51ffe7d835c41ff43e8686583e8f72ade3f072a5e3eef79bc85a345810204b52) |
| `deny_list.tls_list.suffix_value` | [deny_list.tls_list.suffix_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-f71955e994f2c88e5cc7d22ffc644687f6610083c1d5de7251eeaa9aa292bd1b) |
| `description` | [description](data-sources--forward_proxy_policy--reference--group-001.md#canonical-843a50ac043353adb0e3001c171952be4c4f3ccd18cc6532d7c9454d0ae9e86b) |
| `drp_http_connect` | [drp_http_connect](data-sources--forward_proxy_policy--reference--group-001.md#canonical-6ece1385aeb59ea9b1c14f8b4e0574d4d454ad6d2ef12e29e81fa2dce3128c31) |
| `id` | [id](data-sources--forward_proxy_policy--reference--group-001.md#canonical-76713c79b8dbe027f3b14627eabae98bd9c63b1242eb5f03236af1b37991d843) |
| `labels` | [labels](data-sources--forward_proxy_policy--reference--group-001.md#canonical-4605186a4629d8c8e38218ab93ee9a524a6cc4578d541b40bcbeda28a827cec0) |
| `name` | [name](data-sources--forward_proxy_policy--reference--group-001.md#canonical-8147c9af331f89ce96a0698941955283812753241d5dc025b5bb18569557e8fb) |
| `namespace` | [namespace](data-sources--forward_proxy_policy--reference--group-001.md#canonical-8b43a16d0827908e319123b33be0a22bf201e2aed05fc73fd44f35e552db21c0) |
| `network_connector` | [network_connector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d91eb516f17abe31d509c50ddf552be6c294e5be1b956e3555c99290e4aa40dd) |
| `network_connector.name` | [network_connector.name](data-sources--forward_proxy_policy--reference--group-001.md#canonical-835f9488c46674f6cd3e1cfd37b406562e8eff2f8e562983cdeb09316325126a) |
| `network_connector.namespace` | [network_connector.namespace](data-sources--forward_proxy_policy--reference--group-001.md#canonical-c10d34de10692e267a85b893126e1dade12e76733be9e8fa837a213bea26cdc0) |
| `network_connector.tenant` | [network_connector.tenant](data-sources--forward_proxy_policy--reference--group-001.md#canonical-b0e8e466921522ae5be0a8b12e51edee92c2808396ceb48f95a99ed4af5206e9) |
| `proxy_label_selector` | [proxy_label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-f8a52a95b6faf6db0a2876b42d9ffa57f1ad6548183095a1641981e0645155a7) |
| `proxy_label_selector.expressions` | [proxy_label_selector.expressions](data-sources--forward_proxy_policy--reference--group-001.md#canonical-702e3116836a66d938b70ef236e016079a41e0c5a601ab44df73ca36354c2d81) |
| `rule_list` | [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-f00d51034a807a3998d8c0093a5347c4b08a0f3192a866e6fa19311cbe524df5) |
| `rule_list.rules` | [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d0c8e98a896ea438e7f400b983a13d4811213f8277fb26273960552b94bdfcde) |
| `rule_list.rules.action` | [rule_list.rules.action](data-sources--forward_proxy_policy--reference--group-001.md#canonical-5d00b4886d12a008c1f60a1b23cfdf364d6d754e40c644b992de02e36072ab7f) |
| `rule_list.rules.all_destinations` | [rule_list.rules.all_destinations](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e25daed81ddd8568fd23adecc6e0a76cffc60e55a49d46ff7cfecc7982632263) |
| `rule_list.rules.all_sources` | [rule_list.rules.all_sources](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2eea4b677d92a3c97605731c95c334be99ff1e9db2f5021264455e70ce9ea1a9) |
| `rule_list.rules.dst_asn_list` | [rule_list.rules.dst_asn_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-333215d9536e5df8896be3f8d4334bc8cad76139029ca951b81b755f1e68a558) |
| `rule_list.rules.dst_asn_list.as_numbers` | [rule_list.rules.dst_asn_list.as_numbers](data-sources--forward_proxy_policy--reference--group-001.md#canonical-83b1a2c5999e46cd5734b6a98bb1fc15db6870fe9df1e0160d44f2195dbf5cff) |
| `rule_list.rules.dst_asn_set` | [rule_list.rules.dst_asn_set](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d5d668f7f54b4e808e6b94dcb2d4e27d51aea2096968a382345497fcd46949fa) |
| `rule_list.rules.dst_asn_set.name` | [rule_list.rules.dst_asn_set.name](data-sources--forward_proxy_policy--reference--group-001.md#canonical-51a7c60f7f23b5937b084a82923ab6b85e5364eaa019b867d6b9594c647bf635) |
| `rule_list.rules.dst_asn_set.namespace` | [rule_list.rules.dst_asn_set.namespace](data-sources--forward_proxy_policy--reference--group-001.md#canonical-982ef63eb7863f99e9cc7032f751bb9eea0cf5ccd61890dac38868cca91e8222) |
| `rule_list.rules.dst_asn_set.tenant` | [rule_list.rules.dst_asn_set.tenant](data-sources--forward_proxy_policy--reference--group-001.md#canonical-5a30db9b460e5c2e4f2e9bafc0fce3195f0071337b9fa6bbe153ec9425017761) |
| `rule_list.rules.dst_ip_prefix_set` | [rule_list.rules.dst_ip_prefix_set](data-sources--forward_proxy_policy--reference--group-001.md#canonical-53742d6e78a01bb376b6d13dfa1a09aeebadadde8e352604fae41932910f2af6) |
| `rule_list.rules.dst_ip_prefix_set.name` | [rule_list.rules.dst_ip_prefix_set.name](data-sources--forward_proxy_policy--reference--group-001.md#canonical-ba4b30c48abb65f17ddca3a85de44beb70cf87c5eddc587f0857ecb45e649fbb) |
| `rule_list.rules.dst_ip_prefix_set.namespace` | [rule_list.rules.dst_ip_prefix_set.namespace](data-sources--forward_proxy_policy--reference--group-001.md#canonical-7c979712df77ab02ed82e52f44d0c68176d6705da55a4273bf25378a0846ac78) |
| `rule_list.rules.dst_ip_prefix_set.tenant` | [rule_list.rules.dst_ip_prefix_set.tenant](data-sources--forward_proxy_policy--reference--group-001.md#canonical-aed70415cdb9e43f4f8bf847c6e260f035ddffa622ef8b6ad876409b295eaec6) |
| `rule_list.rules.dst_label_selector` | [rule_list.rules.dst_label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-21878ccbd04f985a46d0bf7cad9e52a054db197942006ea35928a2b080e2629b) |
| `rule_list.rules.dst_label_selector.expressions` | [rule_list.rules.dst_label_selector.expressions](data-sources--forward_proxy_policy--reference--group-001.md#canonical-b6d3b69e5c2953fd7ac9db2189ecd0dfc4aea01841f1ca50a4990c281e955bcd) |
| `rule_list.rules.dst_prefix_list` | [rule_list.rules.dst_prefix_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2b30a843a08209b542fa1f2c3fbed3ff4ab768fce41c0785547afedda62b1bac) |
| `rule_list.rules.dst_prefix_list.prefixes` | [rule_list.rules.dst_prefix_list.prefixes](data-sources--forward_proxy_policy--reference--group-001.md#canonical-c5fc55803175e601aedaea7b2f2ff579e3b0b27870561ca75a9179ff06e85ace) |
| `rule_list.rules.http_list` | [rule_list.rules.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-37e8a0f64525e0afaf5c3182c0b4d7a1bcaa39bcaef548c91ed346269714819a) |
| `rule_list.rules.http_list.http_list` | [rule_list.rules.http_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-bede8880f643c226d60c2427d30d67aaf77e9074d6fd98c66cf241a77c1d9fda) |
| `rule_list.rules.http_list.http_list.any_path` | [rule_list.rules.http_list.http_list.any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-40a7b91281b804898af9d7e0f095ff45f671588c4b2a0913055a435bda38f89d) |
| `rule_list.rules.http_list.http_list.exact_value` | [rule_list.rules.http_list.http_list.exact_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d55bf7297377d30e61a6861d3783ca249aa181a5466638e89f9464fcea78e8e3) |
| `rule_list.rules.http_list.http_list.path_exact_value` | [rule_list.rules.http_list.http_list.path_exact_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e22165b11cc4ef47eb348192a9a9502add48d683d40e2f032e875cc581aa2e00) |
| `rule_list.rules.http_list.http_list.path_prefix_value` | [rule_list.rules.http_list.http_list.path_prefix_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-23cafbf3d1ffba31da45a6c19dfa6d939457d75b2c30c8dd44c8ac2236c9364f) |
| `rule_list.rules.http_list.http_list.path_regex_value` | [rule_list.rules.http_list.http_list.path_regex_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-910bdbd1996f6327667ca4b043a59e2477b467215a2db31cbcb7f0d08a26598b) |
| `rule_list.rules.http_list.http_list.regex_value` | [rule_list.rules.http_list.http_list.regex_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-ad3c7d317013635cd3f781e53e44f04ea7db6bb3b1b83e30245c58839ad79652) |
| `rule_list.rules.http_list.http_list.suffix_value` | [rule_list.rules.http_list.http_list.suffix_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-aa3e232113da36d95cf0dadaede54ed3e1f2b6c97a3b0c1c007014954698e604) |
| `rule_list.rules.ip_prefix_set` | [rule_list.rules.ip_prefix_set](data-sources--forward_proxy_policy--reference--group-001.md#canonical-a9f64a439e6f48ee201d434b0244f35d697977cbca00ce0dcb03ea138a349d07) |
| `rule_list.rules.ip_prefix_set.name` | [rule_list.rules.ip_prefix_set.name](data-sources--forward_proxy_policy--reference--group-001.md#canonical-6c02534f4804a6ef4636c4eb8d8f186508bd9ae6e98a82816c0bba2983f0a3fe) |
| `rule_list.rules.ip_prefix_set.namespace` | [rule_list.rules.ip_prefix_set.namespace](data-sources--forward_proxy_policy--reference--group-001.md#canonical-ed1fa673e52f4a8666c127d53fde7f71197ae7e293e45f56dbc2f85e65339217) |
| `rule_list.rules.ip_prefix_set.tenant` | [rule_list.rules.ip_prefix_set.tenant](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1b852a1b9c38e195ef2892927674b908c37a1fb5c6a22c6362f04e35f8ff930d) |
| `rule_list.rules.label_selector` | [rule_list.rules.label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-6e67681bcba54071527a731a0bd5b1a12d420ad9af9adb0e38491c5b4fe338e2) |
| `rule_list.rules.label_selector.expressions` | [rule_list.rules.label_selector.expressions](data-sources--forward_proxy_policy--reference--group-001.md#canonical-fcf8587b69b09134be59258c5baae92b9908f4fdc5bb26ac07d0aa5d6d25c618) |
| `rule_list.rules.metadata` | [rule_list.rules.metadata](data-sources--forward_proxy_policy--reference--group-001.md#canonical-64074700e41b5bdfe391da6c9af50b536e0ba854f611c4a801441225161b2974) |
| `rule_list.rules.metadata.description_spec` | [rule_list.rules.metadata.description_spec](data-sources--forward_proxy_policy--reference--group-001.md#canonical-ca6f97210eb701bac3ca2483f011a4d935be5b4eb84954060e90275fc39cc768) |
| `rule_list.rules.metadata.name` | [rule_list.rules.metadata.name](data-sources--forward_proxy_policy--reference--group-001.md#canonical-a0e9efb3bbc7638e5be78a074dc6f86c8875cd5c8781c0b555d94ec5b660b609) |
| `rule_list.rules.no_http_connect_port` | [rule_list.rules.no_http_connect_port](data-sources--forward_proxy_policy--reference--group-001.md#canonical-500e60d5bbc2bc6f4db0cc74e696f493c50b38f01692c788c4a2f7722f67240e) |
| `rule_list.rules.port_matcher` | [rule_list.rules.port_matcher](data-sources--forward_proxy_policy--reference--group-001.md#canonical-045644c82174dbc6ca8f3294f8339f1a7080fc7607dc29ac139f5d353296a421) |
| `rule_list.rules.port_matcher.invert_matcher` | [rule_list.rules.port_matcher.invert_matcher](data-sources--forward_proxy_policy--reference--group-001.md#canonical-21209ac6b113e9e91266692498f14b4e83bd7e8e19ac51584683b8edc07ae3ef) |
| `rule_list.rules.port_matcher.ports` | [rule_list.rules.port_matcher.ports](data-sources--forward_proxy_policy--reference--group-001.md#canonical-b41f779252f5fe842bd97c7d570bb6bf6a2d022b6504681ab02cdd1d4caa0fe4) |
| `rule_list.rules.prefix_list` | [rule_list.rules.prefix_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-05d66cb54e790bf9cc85d6fe33b11507b1f9b815b3dfd5a8ad216dbd5e9eabff) |
| `rule_list.rules.prefix_list.prefixes` | [rule_list.rules.prefix_list.prefixes](data-sources--forward_proxy_policy--reference--group-001.md#canonical-033da1c676e5549095e7cc0b159f4f3d6f19efae70a86a7d2003535fdf61f652) |
| `rule_list.rules.tls_list` | [rule_list.rules.tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-34cdf7e2db22f4ff3ccfe9eed291c07123e3b26e7f1ef2aa8b80e9228b0c3c7c) |
| `rule_list.rules.tls_list.tls_list` | [rule_list.rules.tls_list.tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-a9a6a6162c26da1bf1c19164234120c4ad4a6d6aa861a91ac197e9feca092839) |
| `rule_list.rules.tls_list.tls_list.exact_value` | [rule_list.rules.tls_list.tls_list.exact_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-6dbb2e1c8bea4085b830ef8e456d5fe73fd07775e443dc9fe57d30fafdf50039) |
| `rule_list.rules.tls_list.tls_list.regex_value` | [rule_list.rules.tls_list.tls_list.regex_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-dd64b62407fd5d95baea29990a6940ec517c121c5d1a8a7f4018fcd510d7faa7) |
| `rule_list.rules.tls_list.tls_list.suffix_value` | [rule_list.rules.tls_list.tls_list.suffix_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3b1e2ae35d14f7a00e84f092856387e361677971001e2743323d85c1e4c0212f) |
| `rule_list.rules.url_category_list` | [rule_list.rules.url_category_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e519f0890da74ec26733c12bb8c3e460897063b69853b6bfcd0c125e59ee7b64) |
| `rule_list.rules.url_category_list.url_categories` | [rule_list.rules.url_category_list.url_categories](data-sources--forward_proxy_policy--reference--group-001.md#canonical-9da500269c7802f29d0088384d5a025302bc33761d0f0f2c1e968a106dc1f990) |

<a id="canonical-eef5ff6f329a931f98d3ff98a680e3ce40e5e56e41cbc63d4465fe0305847a47"></a>

## Next pages — Property reference / 471006cbf4c8 / 11

- [allow_all](data-sources--forward_proxy_policy--reference--group-001.md#canonical-8716c7f4145bacfb52b96bf63e6d7b975d9d1978e5d1a7400c8ba561138be577)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d183f0efaebceda4460734ecc8a76f3c16f106da3f79a8ef564fe822923add71)
- [any_proxy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-31c9ad9033679a59af087236658e34f64f1b9398597f5b5d69f98c2253e73c13)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-9a659a7dd15f7b619c2f958e3e0a23bf6c005723edae551e451c2685739cc28e)
- [drp_http_connect](data-sources--forward_proxy_policy--reference--group-001.md#canonical-313a9458d3817074ebbbda207e7f37f67bd5c77441fec57a2ca64f72e0f99452)
- [network_connector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-c36323cc0a7cfcdf49d229fd178c1ff4108821e0a2eacb7b04c9fd02f7cd6a65)
- [proxy_label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-4ec4f3f751c2ec7e14eb8d0c4ee6682326e0434c5001b90a0d5df49631571d53)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-8716c7f4145bacfb52b96bf63e6d7b975d9d1978e5d1a7400c8ba561138be577"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8d75e0ea5466ded750763ba57f6492ae1b7d289d75dadbe49cb7e6891462379"></a>

## allow_all — allow_all / 1e002edc2e78 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- allow_all

<a id="canonical-a81233db54e5cb34cba0a21ba10d27c7fc42eeed8d2601a53b663f4597353f66"></a>

Type: `["object", {}]`. Computed.

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

- [allow_all](data-sources--forward_proxy_policy--reference--group-001.md#canonical-a81233db54e5cb34cba0a21ba10d27c7fc42eeed8d2601a53b663f4597353f66)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-5d27abb54733201c10bc480eeac20b34b4553ccfc3a00192e403602b56dcce1b)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-41ecd20c3b29261610c585dc9f9e691426a85daf0eea1d64a3f8710278decdb1)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-f00d51034a807a3998d8c0093a5347c4b08a0f3192a866e6fa19311cbe524df5)

Select alternatives according to the provider validators above.

<a id="canonical-d8a7fdee4da064fd471270b0e79f7e83860734aee860b0d35d291033b4d21311"></a>

## Direct properties — allow_all / 1e002edc2e78 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b85bdac117fa6ad27c1b7f5e0ad73b1384a1fd3c7d97e80fcd7c0eacda0a48d4"></a>

## Next pages — allow_all / 1e002edc2e78 / 4

- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-d183f0efaebceda4460734ecc8a76f3c16f106da3f79a8ef564fe822923add71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75de9d990fe47dfe6ffdb99f8e79d820830e2674920ff4602ecf41a2b2c916f4"></a>

## allow_list — allow_list / a2ba9289ba5a / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- allow_list

<a id="canonical-5d27abb54733201c10bc480eeac20b34b4553ccfc3a00192e403602b56dcce1b"></a>

Type: `"single"`. Computed.

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP).

Upstream description:

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP)

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

<a id="canonical-39a827fbf873b6948debd6d7cd4d98c507475ba0724aa927288577c557a6b9ac"></a>

## Direct properties — allow_list / a2ba9289ba5a / 3

- [default_action_allow](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0ccd90870cb30ef23d948a748d6b1bdeb2eeb2305539e8e694b88e11aa4c31ce): complete subsection reference.

- [default_action_deny](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d7cb64018bb556e7817869e1e1bb76b973f6618e009ffad63e03d37ce08b6872): complete subsection reference.

- [default_action_next_policy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0f68bd9afb2c0514af5399c3cee13ecd71bdce8a79894258a92e527f6d186199): complete subsection reference.

- [dest_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-7682ef5f35133bbe298c0c29cc0f291039162d15e1913d66e33940820d238663): complete subsection reference.

- [http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-83937ff064eaf0d995b8db58fca1570cf3da9692cefea560681f8abdbabfde72): complete subsection reference.

- [tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1c4bd9bca9aafc58d24254a2b7100b092185171e3beb4cd8e30f41d3bc4ef0e7): complete subsection reference.

<a id="canonical-62b6294249cdc96428cfaab80dbde2ee1584e14ccf33f8f170ce4f3518fb4be5"></a>

## Next pages — allow_list / a2ba9289ba5a / 4

- [allow_list.default_action_allow](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0ccd90870cb30ef23d948a748d6b1bdeb2eeb2305539e8e694b88e11aa4c31ce)
- [allow_list.default_action_deny](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d7cb64018bb556e7817869e1e1bb76b973f6618e009ffad63e03d37ce08b6872)
- [allow_list.default_action_next_policy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0f68bd9afb2c0514af5399c3cee13ecd71bdce8a79894258a92e527f6d186199)
- [allow_list.dest_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-7682ef5f35133bbe298c0c29cc0f291039162d15e1913d66e33940820d238663)
- [allow_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-83937ff064eaf0d995b8db58fca1570cf3da9692cefea560681f8abdbabfde72)
- [allow_list.tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1c4bd9bca9aafc58d24254a2b7100b092185171e3beb4cd8e30f41d3bc4ef0e7)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-0ccd90870cb30ef23d948a748d6b1bdeb2eeb2305539e8e694b88e11aa4c31ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb893c35f2ca23c74f9e3e7442173c8aa7ef4f666014abe77d2c38c19112bd34"></a>

## allow_list.default_action_allow — allow_list.default_action_allow / e3fdab9b2817 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d183f0efaebceda4460734ecc8a76f3c16f106da3f79a8ef564fe822923add71)
- allow_list.default_action_allow

<a id="canonical-8a6998a7ebec33cd6bbf206435e99b7ca8eb7004ab246e6f90d26143dc969453"></a>

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

<a id="canonical-f1a2f56aabb8350a62f631828b8c17e21ef640b70f07f250644ec91fe3ba9e3b"></a>

## Direct properties — allow_list.default_action_allow / e3fdab9b2817 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eb1d1ad11338ee6562f8546dcb7fd2ca06e0d0aed8c0a75aeb9f5efcda24bc59"></a>

## Next pages — allow_list.default_action_allow / e3fdab9b2817 / 4

- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d183f0efaebceda4460734ecc8a76f3c16f106da3f79a8ef564fe822923add71)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-d7cb64018bb556e7817869e1e1bb76b973f6618e009ffad63e03d37ce08b6872"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a40180802ae43aa62feb7c0b39c40ae016214eebb23fe07a1ad4a93e8a55330"></a>

## allow_list.default_action_deny — allow_list.default_action_deny / 813f6526320e / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d183f0efaebceda4460734ecc8a76f3c16f106da3f79a8ef564fe822923add71)
- allow_list.default_action_deny

<a id="canonical-cc43099840bc57d90ec8af371c029dea849a0743a838bff3943eb9ee38e1d9bc"></a>

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

<a id="canonical-a0661c8b39cea27ec905fc54774704009a010ee7051b304822f739536e481e77"></a>

## Direct properties — allow_list.default_action_deny / 813f6526320e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-155942791eda35657d49cc36601d518064e1cbb66d015a0eefe950cf0bac5249"></a>

## Next pages — allow_list.default_action_deny / 813f6526320e / 4

- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d183f0efaebceda4460734ecc8a76f3c16f106da3f79a8ef564fe822923add71)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-0f68bd9afb2c0514af5399c3cee13ecd71bdce8a79894258a92e527f6d186199"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a15d37825c3e4031cfc5a57cee84af2db37694d650044459b59d1dba00040a50"></a>

## allow_list.default_action_next_policy — allow_list.default_action_next_policy / 0aacedf187b3 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d183f0efaebceda4460734ecc8a76f3c16f106da3f79a8ef564fe822923add71)
- allow_list.default_action_next_policy

<a id="canonical-84d27d0516c4292c1ce6f3f53cdc6a46a29722065e3b59a241c4c2a01a523a0e"></a>

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

<a id="canonical-cf811aa8263ace97b95a7109d48e51a46407ce3b460e6599909653d476b3bb35"></a>

## Direct properties — allow_list.default_action_next_policy / 0aacedf187b3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-225ed0efb3e59356c0b8d20803eb07ee40e55a01c10f7f1f659aea16aac2bd3e"></a>

## Next pages — allow_list.default_action_next_policy / 0aacedf187b3 / 4

- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d183f0efaebceda4460734ecc8a76f3c16f106da3f79a8ef564fe822923add71)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-7682ef5f35133bbe298c0c29cc0f291039162d15e1913d66e33940820d238663"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d3912a46ba677fd009c172abeb7727cf0d05771b3bb0f055f8c7d527a426fc8"></a>

## allow_list.dest_list — allow_list.dest_list / 0d57500cff82 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d183f0efaebceda4460734ecc8a76f3c16f106da3f79a8ef564fe822923add71)
- allow_list.dest_list

<a id="canonical-e7848c259b2e5c745919fa924a46c7c0ed4a65cd4b732ecad38b91a5a824c3c7"></a>

Type: `"list"`. Computed.

L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI.

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

<a id="canonical-318ec35783a66373e9b774b1e954a5f203eb834a25b97c311999bf64516acff4"></a>

## Direct properties — allow_list.dest_list / 0d57500cff82 / 3

<a id="canonical-dbcd89a2369ad30f310ecdcd971a54a935c7fd5622d691b44d94da6b8d7b1f65"></a>

<a id="canonical-40694a4a7ff2b39ba73b2ee4d101dfd7eef51324eb86fe97a45678ee5ef0c38b"></a>

## ipv6_prefixes property — allow_list.dest_list / 0d57500cff82 / 4

Type: `["list", "string"]`. Computed.

IPv6 Prefixes. Destination IPv6 prefixes.

Upstream description:

Destination IPv6 prefixes.

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

<a id="canonical-68c230d1d8d332d7239d1b334f266dd92f75ba4aa43fe766a562aeaf682181fa"></a>

<a id="canonical-609be6d2dfd45b0c0437736c5cba023e18806e95a8a609a034342a5733cda888"></a>

## port_ranges property — allow_list.dest_list / 0d57500cff82 / 5

Type: `"string"`. Computed.

String containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by '-'.

Upstream description:

A string containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by "-".

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

<a id="canonical-ede810db665677d090f97af7cfbf3b0faede00cfd1ca2675131b4d36a1211db4"></a>

<a id="canonical-d1ae2553047e47c04900396b94e24448821bf7f5fdf84a60f6b8f2b10ebfa2e6"></a>

## prefixes property — allow_list.dest_list / 0d57500cff82 / 6

Type: `["list", "string"]`. Computed.

IPv4 Prefixes. Destination IPv4 prefixes.

Upstream description:

Destination IPv4 prefixes.

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

<a id="canonical-9fe0cf1c4989c3c5b6cc0dccb5e6d4e88df4c8c0e3995579abc237743c93ac21"></a>

## Next pages — allow_list.dest_list / 0d57500cff82 / 7

- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d183f0efaebceda4460734ecc8a76f3c16f106da3f79a8ef564fe822923add71)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-83937ff064eaf0d995b8db58fca1570cf3da9692cefea560681f8abdbabfde72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eeb0f447bde32f0df8ad03b3cbc3e969574cfff7487a6f0082ff28674d0d4480"></a>

## allow_list.http_list — allow_list.http_list / ae645eccc362 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d183f0efaebceda4460734ecc8a76f3c16f106da3f79a8ef564fe822923add71)
- allow_list.http_list

<a id="canonical-c08c3262a51e3f717c50fc90967f35a907cb8f1464a1b3cdac99035a91846cc5"></a>

Type: `"list"`. Computed.

HTTP URLs. URLs for HTTP connections.

Upstream description:

URLs for HTTP connections.

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

<a id="canonical-348ce905561250f74b26b91225571ce56d43b19830f9c073dcbd7e086ef7f8b0"></a>

## Direct properties — allow_list.http_list / ae645eccc362 / 3

- [any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-16ff232b460a4c44c3628858959499ad2a3145ee118cbaacf849402f2c71488b): complete subsection reference.

<a id="canonical-fa51edac2ebca69d87ad32f4ac441d1cc879ae664cb8b3a955c00bba9a9e7dfd"></a>

<a id="canonical-16c5e0c8c990edc5b8602069cdb74b813e1744d918d6b28b96bd70fce39efa60"></a>

## exact_value property — allow_list.http_list / ae645eccc362 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-5e778423a2849f7148e092d767266f5089f3274266de134d810576f46422b8cc"></a>

<a id="canonical-3bbf2a627e0515a282679346e6e651e564797e40dea8d31c2b770d7a0e93e27b"></a>

## path_exact_value property — allow_list.http_list / ae645eccc362 / 5

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Upstream description:

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

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

<a id="canonical-4ab5c3ed14a8ec6425de3e01e7a550b74bc6c8937cfd71c4d55e557842b8b8ff"></a>

<a id="canonical-f9bf709848063535754e61d3468c23577143ffb28547511feed23bdb689c8296"></a>

## path_prefix_value property — allow_list.http_list / ae645eccc362 / 6

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g "/abc/xyz"
will match "/abc/xyz/.\*"

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

<a id="canonical-4de8928a6e787a1d3b12a6ee7f6f22fcd5b278f8ea0ea446fd11f0d1e8ccb0f2"></a>

<a id="canonical-56c5275f6c213738b957195554b26f5137087426144cf1caeedfb46432b868f3"></a>

## path_regex_value property — allow_list.http_list / ae645eccc362 / 7

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

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

<a id="canonical-e5d5f00c31f73bb5692a221241d577468a92d447364378ed62b69956794d7f77"></a>

<a id="canonical-2144c51c1f2a160f5b63ba3bc1eae220024aa9448cd5a1d89bea4a1b60e7a8ba"></a>

## regex_value property — allow_list.http_list / ae645eccc362 / 8

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-ba6cd8ffbcd098306a37109e96ee39204d8f008606bc10e10445c2655ac5d00d"></a>

<a id="canonical-e653681c23f7b0ea2af2ea4857340f7b01fe991fe8efa5bf7a10ef53627d57a0"></a>

## suffix_value property — allow_list.http_list / ae645eccc362 / 9

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain names e.g "xyz.com" will match
"\*.xyz.com"

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

<a id="canonical-9713c5630c14636bca69b10d03819681474c0868df8e08e2166b5904dd345170"></a>

## Next pages — allow_list.http_list / ae645eccc362 / 10

- [allow_list.http_list.any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-16ff232b460a4c44c3628858959499ad2a3145ee118cbaacf849402f2c71488b)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d183f0efaebceda4460734ecc8a76f3c16f106da3f79a8ef564fe822923add71)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-16ff232b460a4c44c3628858959499ad2a3145ee118cbaacf849402f2c71488b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-603c4669508fa2a40d9d9e00acb8948298d552fe6950c9952df43b00e1189da0"></a>

## allow_list.http_list.any_path — allow_list.http_list.any_path / 97531630020e / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d183f0efaebceda4460734ecc8a76f3c16f106da3f79a8ef564fe822923add71)
- [allow_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-83937ff064eaf0d995b8db58fca1570cf3da9692cefea560681f8abdbabfde72)
- allow_list.http_list.any_path

<a id="canonical-77398e6ed822ec383637e4ba47ba64f5120ffb5c2fcdd904e43ad138e4424792"></a>

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

<a id="canonical-e2c546265dc01ecf171e20c273538f676171ee82eb1eff2429f18bbe8dacb784"></a>

## Direct properties — allow_list.http_list.any_path / 97531630020e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-92f12e8d99f3a4bcef58409761b106bc2ec158cb9d48101098bb814aa1f30a12"></a>

## Next pages — allow_list.http_list.any_path / 97531630020e / 4

- [allow_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-83937ff064eaf0d995b8db58fca1570cf3da9692cefea560681f8abdbabfde72)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-1c4bd9bca9aafc58d24254a2b7100b092185171e3beb4cd8e30f41d3bc4ef0e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5dd13b1ddbd0e3d863ba212c2013fec71f088462f64ee752bf04b9839fc47355"></a>

## allow_list.tls_list — allow_list.tls_list / 9c0d6a6a638e / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d183f0efaebceda4460734ecc8a76f3c16f106da3f79a8ef564fe822923add71)
- allow_list.tls_list

<a id="canonical-743f01c933974bf17f856dac48548fc0ceb17290e9b45c8f6ede454775909e6a"></a>

Type: `"list"`. Computed.

TLS Domains. Domains in SNI for TLS connections.

Upstream description:

Domains in SNI for TLS connections.

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

<a id="canonical-d538edf2c25ec8be1f738570112a9221ec4860dfbf646500d0cf8d281a9d2242"></a>

## Direct properties — allow_list.tls_list / 9c0d6a6a638e / 3

<a id="canonical-b314ae255f36ac6db5dcc5d161328701363a518a51b8099e8cd2913d1e44a5f1"></a>

<a id="canonical-19764fffccc4d62005076b500adbf7316fa60bfa9cc8282fe72da34e6e32ad04"></a>

## exact_value property — allow_list.tls_list / 9c0d6a6a638e / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-ce096b871ae4d9c107f464c3b855eba984f9997a3c627cda276e77ad7609cd1b"></a>

<a id="canonical-c9effc339f70b06a111c19dc28e046a10871781ec51a665d9d3ab7cdc17ffcdb"></a>

## regex_value property — allow_list.tls_list / 9c0d6a6a638e / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-fa2926bc6d049db5731554817636403feda899ee0f1e9e1ddadf846dd999ca67"></a>

<a id="canonical-be47d9a3d130c25d09eec440d984c16c9b2dc9dad8934b914de0df8bdc6e022d"></a>

## suffix_value property — allow_list.tls_list / 9c0d6a6a638e / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-125a16ebd2fddffaf8560c62d74ca3f79413b73f16db2988818786cd4eecd2ec"></a>

## Next pages — allow_list.tls_list / 9c0d6a6a638e / 7

- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d183f0efaebceda4460734ecc8a76f3c16f106da3f79a8ef564fe822923add71)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-31c9ad9033679a59af087236658e34f64f1b9398597f5b5d69f98c2253e73c13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9124bb4d9772847f9e63a86edc509d787d0f25285fe2b83e9dbf2bd789f69e97"></a>

## any_proxy — any_proxy / 23e97d9449b5 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- any_proxy

<a id="canonical-dbdc995e9c52efce3354b50fd6e69b97e554a51587f6e1550efd4077c23931f9"></a>

Type: `["object", {}]`. Computed.

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

- [any_proxy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-dbdc995e9c52efce3354b50fd6e69b97e554a51587f6e1550efd4077c23931f9)
- [drp_http_connect](data-sources--forward_proxy_policy--reference--group-001.md#canonical-6ece1385aeb59ea9b1c14f8b4e0574d4d454ad6d2ef12e29e81fa2dce3128c31)
- [network_connector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d91eb516f17abe31d509c50ddf552be6c294e5be1b956e3555c99290e4aa40dd)
- [proxy_label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-f8a52a95b6faf6db0a2876b42d9ffa57f1ad6548183095a1641981e0645155a7)

Select alternatives according to the provider validators above.

<a id="canonical-4fc37d23a4936302b2d5e51baa42221b0daad26d2d047a236e63371dc49808fc"></a>

## Direct properties — any_proxy / 23e97d9449b5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-15477362edcc15fadf732759a738fc5a6d0aa73dc296e0fa256fa09d958220ed"></a>

## Next pages — any_proxy / 23e97d9449b5 / 4

- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-9a659a7dd15f7b619c2f958e3e0a23bf6c005723edae551e451c2685739cc28e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3bd438d26c3d382697bd334877ae556eec0cd2975445e53f7b1d807cd317cbb"></a>

## deny_list — deny_list / 25da40f84a5f / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- deny_list

<a id="canonical-41ecd20c3b29261610c585dc9f9e691426a85daf0eea1d64a3f8710278decdb1"></a>

Type: `"single"`. Computed.

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP).

Upstream description:

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP)

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

<a id="canonical-1a4ef1ecde525989eda4b09c3fd197ced4f077a5eec457f262711869ba4bb007"></a>

## Direct properties — deny_list / 25da40f84a5f / 3

- [default_action_allow](data-sources--forward_proxy_policy--reference--group-001.md#canonical-b82c81d388e53ec61540b7b77031c166501f0a7d5a0d7a44cbc6063a930d3ae4): complete subsection reference.

- [default_action_deny](data-sources--forward_proxy_policy--reference--group-001.md#canonical-793ce488ab483967827cce706d954cf8a8c894231dcf774a1432d6a59662ae24): complete subsection reference.

- [default_action_next_policy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-65bc1f73175b0a16c770a840832fe9ab0bbf5d394c396941eead958c18e9a3fc): complete subsection reference.

- [dest_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d9295efc985939e509d620feb83595076c137d6fe711517bf7db4629c80dc5bb): complete subsection reference.

- [http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-33b4bcc7804591b73c8a47886dc1f4c9873d798b531cbf0b8b2f7067950d1dde): complete subsection reference.

- [tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-97b9d8257e81ce67e2ad89da6d53704f14603be62fd46f3e6f4902a7a941692e): complete subsection reference.

<a id="canonical-1f59b574309b12a9e90f0153111b601bec1d4338c6f4ac69478de3aefb628c93"></a>

## Next pages — deny_list / 25da40f84a5f / 4

- [deny_list.default_action_allow](data-sources--forward_proxy_policy--reference--group-001.md#canonical-b82c81d388e53ec61540b7b77031c166501f0a7d5a0d7a44cbc6063a930d3ae4)
- [deny_list.default_action_deny](data-sources--forward_proxy_policy--reference--group-001.md#canonical-793ce488ab483967827cce706d954cf8a8c894231dcf774a1432d6a59662ae24)
- [deny_list.default_action_next_policy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-65bc1f73175b0a16c770a840832fe9ab0bbf5d394c396941eead958c18e9a3fc)
- [deny_list.dest_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-d9295efc985939e509d620feb83595076c137d6fe711517bf7db4629c80dc5bb)
- [deny_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-33b4bcc7804591b73c8a47886dc1f4c9873d798b531cbf0b8b2f7067950d1dde)
- [deny_list.tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-97b9d8257e81ce67e2ad89da6d53704f14603be62fd46f3e6f4902a7a941692e)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-b82c81d388e53ec61540b7b77031c166501f0a7d5a0d7a44cbc6063a930d3ae4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc4ca6c673e38476cec06fd32cc05a1e4c7e3cdbc9883e59407ba3666a9d7190"></a>

## deny_list.default_action_allow — deny_list.default_action_allow / 66fda49702f7 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-9a659a7dd15f7b619c2f958e3e0a23bf6c005723edae551e451c2685739cc28e)
- deny_list.default_action_allow

<a id="canonical-e26dfb58d1e9d2d75beba7791c410bdfc9e8bf0434642e4fdb525898140d58f0"></a>

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

<a id="canonical-68cc0c00ef5a5dbb7a24dcd9b33023c00bc825e649395fb97e28cca89bedb5c7"></a>

## Direct properties — deny_list.default_action_allow / 66fda49702f7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cd6a1c056a5e600a730cca23868c7aca39a618de99f07d827e401310e860006e"></a>

## Next pages — deny_list.default_action_allow / 66fda49702f7 / 4

- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-9a659a7dd15f7b619c2f958e3e0a23bf6c005723edae551e451c2685739cc28e)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-793ce488ab483967827cce706d954cf8a8c894231dcf774a1432d6a59662ae24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9ca85f0077fd0e4b04f1094f0d88607ef25c887600a3b80643d2addf682d436"></a>

## deny_list.default_action_deny — deny_list.default_action_deny / 1ba520557074 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-9a659a7dd15f7b619c2f958e3e0a23bf6c005723edae551e451c2685739cc28e)
- deny_list.default_action_deny

<a id="canonical-965488b0aa8bfb864aec63d50cb1d4c49695963a87e4b930e068155230cf133e"></a>

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

<a id="canonical-5d01ac26d199c83c59dd076520d561c2630f5a51ae487a962e15ed25cecf8568"></a>

## Direct properties — deny_list.default_action_deny / 1ba520557074 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-45c65733f30d1e3a77958622a814fa8e29ef96f7d0f6f7ed8b38d58cd0b75034"></a>

## Next pages — deny_list.default_action_deny / 1ba520557074 / 4

- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-9a659a7dd15f7b619c2f958e3e0a23bf6c005723edae551e451c2685739cc28e)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-65bc1f73175b0a16c770a840832fe9ab0bbf5d394c396941eead958c18e9a3fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b3472598a110a37bc67303b4cf6eb6dda19e7899eb6cdef0567ad687929f770"></a>

## deny_list.default_action_next_policy — deny_list.default_action_next_policy / 316f00622e65 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-9a659a7dd15f7b619c2f958e3e0a23bf6c005723edae551e451c2685739cc28e)
- deny_list.default_action_next_policy

<a id="canonical-f714e05f35e35d559de67b09f21e16535d05db697eb448938dca5719e7de2217"></a>

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

<a id="canonical-724b79332ea1c54dad6ec4a1e224f37a3194c8212ba1880855b80f5d2418a1f5"></a>

## Direct properties — deny_list.default_action_next_policy / 316f00622e65 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-67830d26d65142e5f2e655fc1de93dd2d4452579420760a13d1b9f0d1962ed63"></a>

## Next pages — deny_list.default_action_next_policy / 316f00622e65 / 4

- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-9a659a7dd15f7b619c2f958e3e0a23bf6c005723edae551e451c2685739cc28e)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-d9295efc985939e509d620feb83595076c137d6fe711517bf7db4629c80dc5bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fddca05a74b45bc65c372facf75ac31b9a65d36f088be38d192deae0c78b8475"></a>

## deny_list.dest_list — deny_list.dest_list / aef4e2b22873 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-9a659a7dd15f7b619c2f958e3e0a23bf6c005723edae551e451c2685739cc28e)
- deny_list.dest_list

<a id="canonical-bd0cf8dc7270a777f6350932af75ffade0ba09c2b956253cc6c239df43ca50f6"></a>

Type: `"list"`. Computed.

L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI.

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

<a id="canonical-f0842d642c522bcd19dbda4ed134e703c57deb60894b7a91097551da54122c57"></a>

## Direct properties — deny_list.dest_list / aef4e2b22873 / 3

<a id="canonical-bc10d319fd26b8bc60f2a621c1d3e568db7639961cb7648295222a9b42f72861"></a>

<a id="canonical-d9946e60633bb7d52b57ca81a6f0cccad19c4789935cffc75626e7c35438eecf"></a>

## ipv6_prefixes property — deny_list.dest_list / aef4e2b22873 / 4

Type: `["list", "string"]`. Computed.

IPv6 Prefixes. Destination IPv6 prefixes.

Upstream description:

Destination IPv6 prefixes.

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

<a id="canonical-58d8c5d9f395de01067216a1e0ad3db0ad40807942b43df89ac50ad220c48d70"></a>

<a id="canonical-57c4538a9c0d88d31042954951e5326a21b6b1a488e8be5d531d28f5681e5bf3"></a>

## port_ranges property — deny_list.dest_list / aef4e2b22873 / 5

Type: `"string"`. Computed.

String containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by '-'.

Upstream description:

A string containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by "-".

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

<a id="canonical-11eeb721a367e797f8ba92fe6d2f2fb8448e0ff5f8bcdc04c4b65aaa663048cd"></a>

<a id="canonical-cd736f693a56ce2f314a89c57f8c334a2ee78e15aad7ff02f109a9fb55ec2b87"></a>

## prefixes property — deny_list.dest_list / aef4e2b22873 / 6

Type: `["list", "string"]`. Computed.

IPv4 Prefixes. Destination IPv4 prefixes.

Upstream description:

Destination IPv4 prefixes.

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

<a id="canonical-41e1d0a7d93ea32819903879337c9609ff78dab15d4778cfaa5fd7f0810bcdf4"></a>

## Next pages — deny_list.dest_list / aef4e2b22873 / 7

- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-9a659a7dd15f7b619c2f958e3e0a23bf6c005723edae551e451c2685739cc28e)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-33b4bcc7804591b73c8a47886dc1f4c9873d798b531cbf0b8b2f7067950d1dde"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1cbefdff359f35d76ac25eba0c7df1a15689f141bfcb36a21d9672308256730"></a>

## deny_list.http_list — deny_list.http_list / 4faa31ad51fc / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-9a659a7dd15f7b619c2f958e3e0a23bf6c005723edae551e451c2685739cc28e)
- deny_list.http_list

<a id="canonical-fe3123633aa8fe730ac80f4dc384cbc71d6a3a92ea90d18e258e35c0ef6c9b25"></a>

Type: `"list"`. Computed.

HTTP URLs. URLs for HTTP connections.

Upstream description:

URLs for HTTP connections.

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

<a id="canonical-f8409017b5553e13675a18779276f938bc81398868d2693f2e0962b16a638c5a"></a>

## Direct properties — deny_list.http_list / 4faa31ad51fc / 3

- [any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-97d1e2162088ad2ca9f2dd46ada509fd3287684479091f14f8aae9dbddcb1170): complete subsection reference.

<a id="canonical-06b3221b3cc6ed7afc0ce6961d859975a121f4187a7c95ced53e77f7369a0466"></a>

<a id="canonical-6ea694317e335a1924764c6d447493790161711c4fbaddb6c59ff8d240ec8e3c"></a>

## exact_value property — deny_list.http_list / 4faa31ad51fc / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-2521fd081dfc8609ceea506bbbbfa6b30a884b41c089c293aba8138879edf16d"></a>

<a id="canonical-a446a3b66507ca41ff9947af064fa6e0c2d2e4b20b65125c30c551af5fd32695"></a>

## path_exact_value property — deny_list.http_list / 4faa31ad51fc / 5

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Upstream description:

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

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

<a id="canonical-c8f791c2067142cf66911322ac25f2097ede848e1f2478f3d10ac719873bfcaf"></a>

<a id="canonical-83727cc70626e84b015a48d0d0cd5d015b3c1db7fd5cd6b238621ff0df73da22"></a>

## path_prefix_value property — deny_list.http_list / 4faa31ad51fc / 6

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g "/abc/xyz"
will match "/abc/xyz/.\*"

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

<a id="canonical-53ae9296c6aefc3d09debf2c20f54ac0b2c2f6fba7c0d4430297c6efc2504cab"></a>

<a id="canonical-0ce7572e71cb645e1517dd79122ffe62a5bcbf900f32ca03afeaa642f029c087"></a>

## path_regex_value property — deny_list.http_list / 4faa31ad51fc / 7

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

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

<a id="canonical-465ba2570b04fb2e1d7d9d2f9da0225574a8c65dcf6a0dcd316764e7b10301d3"></a>

<a id="canonical-0ec2d42ccfae5fe3c557524af54320b7d3d3437a20622a3427333b9c3ecf62b2"></a>

## regex_value property — deny_list.http_list / 4faa31ad51fc / 8

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-84e8d6dd2d222d76b35bebdfb946395ca4e27a75557cb4afa9f7f7ff3c856d16"></a>

<a id="canonical-78a4f05fa630ba5ebe9b3e76cf0cdd9907c06e5dd34d9238ce68c79c619471e0"></a>

## suffix_value property — deny_list.http_list / 4faa31ad51fc / 9

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain names e.g "xyz.com" will match
"\*.xyz.com"

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

<a id="canonical-b0f50c441e3fd7f8f58ff6e96e5deea9a24258a110e270e64c760cfece9c6f76"></a>

## Next pages — deny_list.http_list / 4faa31ad51fc / 10

- [deny_list.http_list.any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-97d1e2162088ad2ca9f2dd46ada509fd3287684479091f14f8aae9dbddcb1170)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-9a659a7dd15f7b619c2f958e3e0a23bf6c005723edae551e451c2685739cc28e)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-97d1e2162088ad2ca9f2dd46ada509fd3287684479091f14f8aae9dbddcb1170"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1bbe33d4b90530f69f47603fd4e1a9cfe11dbf702c36db4459c73b48252ebf37"></a>

## deny_list.http_list.any_path — deny_list.http_list.any_path / dc03c8840f8f / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-9a659a7dd15f7b619c2f958e3e0a23bf6c005723edae551e451c2685739cc28e)
- [deny_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-33b4bcc7804591b73c8a47886dc1f4c9873d798b531cbf0b8b2f7067950d1dde)
- deny_list.http_list.any_path

<a id="canonical-ec7ba3459fa31c28866928a9e057d363a5d0ebf1058cc71823b759bee455bceb"></a>

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

<a id="canonical-fef4d69e2c0c40e6a1924cfd0e3ce580a5610ad851515900acceac70058e438b"></a>

## Direct properties — deny_list.http_list.any_path / dc03c8840f8f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d6d453edc34b3dc15df17f5545c9aa20dd18ec80958df1b4acb9210354df2590"></a>

## Next pages — deny_list.http_list.any_path / dc03c8840f8f / 4

- [deny_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-33b4bcc7804591b73c8a47886dc1f4c9873d798b531cbf0b8b2f7067950d1dde)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-97b9d8257e81ce67e2ad89da6d53704f14603be62fd46f3e6f4902a7a941692e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c4ae02779967e0225cdd57054d04d8f91b84f0b5683ca11ea00484f93a82b8d"></a>

## deny_list.tls_list — deny_list.tls_list / b76a1e49088d / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-9a659a7dd15f7b619c2f958e3e0a23bf6c005723edae551e451c2685739cc28e)
- deny_list.tls_list

<a id="canonical-4072d3b1df34177fbdeadf0bfd77a3faad51511e49a5af6050b3da91f34c9647"></a>

Type: `"list"`. Computed.

TLS Domains. Domains in SNI for TLS connections.

Upstream description:

Domains in SNI for TLS connections.

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

<a id="canonical-96fa4733bf618dd399a6f354db3c7d04b9a42e2215b35a6afbf1a2fb0acf6e99"></a>

## Direct properties — deny_list.tls_list / b76a1e49088d / 3

<a id="canonical-a0e6e01063b1d77e71c6c22e43b417775e617f598c14f9eda27dc766a352c439"></a>

<a id="canonical-a46b0be242636b53f5a981311fcc4434bd77bc6b85f7128660dc8770c3a80377"></a>

## exact_value property — deny_list.tls_list / b76a1e49088d / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-51ffe7d835c41ff43e8686583e8f72ade3f072a5e3eef79bc85a345810204b52"></a>

<a id="canonical-f66d8c83a1a890769ec92ec7f5fa964bc787b8f78b3abaab5f04c7245ade522a"></a>

## regex_value property — deny_list.tls_list / b76a1e49088d / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-f71955e994f2c88e5cc7d22ffc644687f6610083c1d5de7251eeaa9aa292bd1b"></a>

<a id="canonical-ea6268b6bb6d51a5052dee86eba91ba087215c3ad17f30274a5c3b3837845e73"></a>

## suffix_value property — deny_list.tls_list / b76a1e49088d / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-bc4542434e4dda7d455ad2f3ce4a019a147527db54b5a3a280c1ae67a4ed6243"></a>

## Next pages — deny_list.tls_list / b76a1e49088d / 7

- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-9a659a7dd15f7b619c2f958e3e0a23bf6c005723edae551e451c2685739cc28e)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-313a9458d3817074ebbbda207e7f37f67bd5c77441fec57a2ca64f72e0f99452"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b43e9990b69315927e8bc5d4033e2ebcb27554177cfb44ab706ae640a0dacf7"></a>

## drp_http_connect — drp_http_connect / e7808687a18b / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- drp_http_connect

<a id="canonical-6ece1385aeb59ea9b1c14f8b4e0574d4d454ad6d2ef12e29e81fa2dce3128c31"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-655d78ae3d6cde2bb3f410a1cb3b6d0caecea99d7f4eb64b8c0eae3c83f3b6e4"></a>

## Direct properties — drp_http_connect / e7808687a18b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bba746b9343b934dad3b2fca4ade368bcd477123c5ad5456091f71fd42ba5e26"></a>

## Next pages — drp_http_connect / e7808687a18b / 4

- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-c36323cc0a7cfcdf49d229fd178c1ff4108821e0a2eacb7b04c9fd02f7cd6a65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6967a9aa6bc0c8d6d60e246e082a718dfdab769be9d1e5ac491232bc48c0061"></a>

## network_connector — network_connector / 7d5d04f73ad9 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- network_connector

<a id="canonical-d91eb516f17abe31d509c50ddf552be6c294e5be1b956e3555c99290e4aa40dd"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-08fbf8f189d03467f8c060b3344b0e57cc930d187c05f0fe8bf728ca9b27a0af"></a>

## Direct properties — network_connector / 7d5d04f73ad9 / 3

<a id="canonical-835f9488c46674f6cd3e1cfd37b406562e8eff2f8e562983cdeb09316325126a"></a>

<a id="canonical-ab52aceacb7dad5b38f132421f01073c1eb1ef11d18b970801b89fe4faeb16c7"></a>

## name property — network_connector / 7d5d04f73ad9 / 4

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

<a id="canonical-c10d34de10692e267a85b893126e1dade12e76733be9e8fa837a213bea26cdc0"></a>

<a id="canonical-54a14d16b7b030c92c0663e7f5d6ef051e5a9fc393059d2b765eb0d21c72694e"></a>

## namespace property — network_connector / 7d5d04f73ad9 / 5

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

<a id="canonical-b0e8e466921522ae5be0a8b12e51edee92c2808396ceb48f95a99ed4af5206e9"></a>

<a id="canonical-4beda32876b3d44c42b75c9452cab17187fa315b1bec8e90cc3ca0c7f5feccb3"></a>

## tenant property — network_connector / 7d5d04f73ad9 / 6

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

<a id="canonical-4f1bcb6352dcb238be3e47d5e690ee0aac4c320aa7b21fa5dfd7443e8fabdc78"></a>

## Next pages — network_connector / 7d5d04f73ad9 / 7

- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-4ec4f3f751c2ec7e14eb8d0c4ee6682326e0434c5001b90a0d5df49631571d53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-924a79ea2ec10497e591b202fe389bd89f513a9a4f00c99caf49367413d5ceaf"></a>

## proxy_label_selector — proxy_label_selector / 8a55ef476ba6 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- proxy_label_selector

<a id="canonical-f8a52a95b6faf6db0a2876b42d9ffa57f1ad6548183095a1641981e0645155a7"></a>

Type: `"single"`. Computed.

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

<a id="canonical-a65c55782b3693cd8fbcbf11b4a6f0db5a2ef7fee4083dca53c4f1234cf82e70"></a>

## Direct properties — proxy_label_selector / 8a55ef476ba6 / 3

<a id="canonical-702e3116836a66d938b70ef236e016079a41e0c5a601ab44df73ca36354c2d81"></a>

<a id="canonical-e1a4b4ad531d098eeacbcb121cd4ef95dcb01380ceaa7a8c176501791e203a7d"></a>

## expressions property — proxy_label_selector / 8a55ef476ba6 / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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

<a id="canonical-e6c2e06b9b146748eb221b3c22cab63736c0ff87f17218e9dd01324a49571539"></a>

## Next pages — proxy_label_selector / 8a55ef476ba6 / 5

- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d4c865cade13d79c9d3f98c2421edd522963958ffe2f5ecfe5ee2442cf34ace"></a>

## rule_list — rule_list / 5b458183d09a / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- rule_list

<a id="canonical-f00d51034a807a3998d8c0093a5347c4b08a0f3192a866e6fa19311cbe524df5"></a>

Type: `"single"`. Computed.

Custom Rule List. List of custom rules.

Upstream description:

List of custom rules.

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

<a id="canonical-918089a907ccc07f8aad8c13c6582682979ba7ebc7dd75f11092f0db5dc06f48"></a>

## Direct properties — rule_list / 5b458183d09a / 3

- [rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3): complete subsection reference.

<a id="canonical-9b5967a6f1c688b7dea2f87b143c0c11951e6bb02f9f0115d4aa9c5f5fc6ee52"></a>

## Next pages — rule_list / 5b458183d09a / 4

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f8915dd7a0e53f44016802813ccb634799d1ca2f1963ec6828c402037a1954e"></a>

## rule_list.rules — rule_list.rules / 139d797a0be8 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- rule_list.rules

<a id="canonical-d0c8e98a896ea438e7f400b983a13d4811213f8277fb26273960552b94bdfcde"></a>

Type: `"list"`. Computed.

Custom Rule List. List of custom rules.

Upstream description:

List of custom rules.

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

<a id="canonical-d281cd11783ea97f9ee2542fa3ccb7b79624615fab506b42bd87091aa2b5a883"></a>

## Direct properties — rule_list.rules / 139d797a0be8 / 3

<a id="canonical-5d00b4886d12a008c1f60a1b23cfdf364d6d754e40c644b992de02e36072ab7f"></a>

<a id="canonical-b383950a96a2352dadfde9ff6132cf296cddc9613fb1e5ac8748ca499fc6c43d"></a>

## action property — rule_list.rules / 139d797a0be8 / 4

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [all_destinations](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1e1f9fe39b184e3dac247e475c54c2a85ebbc00ab5248d6cdf79327ad7f2ed93): complete subsection reference.

- [all_sources](data-sources--forward_proxy_policy--reference--group-001.md#canonical-4fb66f72353ad83855bbc0c589bb2a3760a3bc23bb5b1eeac1566181f0f86ff7): complete subsection reference.

- [dst_asn_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-706597db7df570db9b0b4cd7a78fb6e3babe1406bf022494afadc916c7080ab0): complete subsection reference.

- [dst_asn_set](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3f47581361ba06ed76d3deb73ca2d5cf4bc7b261affc60b920ca5520b3acb426): complete subsection reference.

- [dst_ip_prefix_set](data-sources--forward_proxy_policy--reference--group-001.md#canonical-f73c4867739bb963eb34eb55401f1bbd8e31b79daac42de1ae283fb9e96239e5): complete subsection reference.

- [dst_label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3954a30ebacac371bcfc202b6b55edd70767645c81efaeaec4021c0cb55c0c9b): complete subsection reference.

- [dst_prefix_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-72245fe962e5b6889eb3a046890cc1dcc379f4c22ef1c92a1c9593ba70b878af): complete subsection reference.

- [http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-53c0c93d85e18cb9293b46690a42495708c4c986c187adb960a5c59403324aa5): complete subsection reference.

- [ip_prefix_set](data-sources--forward_proxy_policy--reference--group-001.md#canonical-a64dcaf4fc73228e2330d5dc43f9eada4b5479060cd46329952dd3abae9d19c0): complete subsection reference.

- [label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-191cbbf56360c30f7ec31fb039938030568c4ebd5ffb6d63a60e602512e73ade): complete subsection reference.

- [metadata](data-sources--forward_proxy_policy--reference--group-001.md#canonical-eb4a1b3f1fef05e208b45527537b7150b36b9d43d8b078f8235b2d08054489f6): complete subsection reference.

- [no_http_connect_port](data-sources--forward_proxy_policy--reference--group-001.md#canonical-fa532b4ab7079e75055b76fe3dce0c8e703dbd31dd8628e80bc76d3abbd5bedb): complete subsection reference.

- [port_matcher](data-sources--forward_proxy_policy--reference--group-001.md#canonical-f3acd89a9c24f235c071cff02fad14cf28080f8710b3b3987342457db5f5fd02): complete subsection reference.

- [prefix_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e671661533b1c895ca42eaefc2878d61a1f8bcbea5b44128cac37544b9652f6d): complete subsection reference.

- [tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-dc35779a77d04fbbc6df18bd4839b93d8fb3f8193587742cf7f7b4bf4be31c79): complete subsection reference.

- [url_category_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2f88fa959a26b990900806339f765a4bb2976cb74d72bbf01601682ae0d08cee): complete subsection reference.

<a id="canonical-97f540c90e134e9fbffaa0848258f434edbd93b4ecd4c2788daf577c61235f50"></a>

## Next pages — rule_list.rules / 139d797a0be8 / 5

- [rule_list.rules.all_destinations](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1e1f9fe39b184e3dac247e475c54c2a85ebbc00ab5248d6cdf79327ad7f2ed93)
- [rule_list.rules.all_sources](data-sources--forward_proxy_policy--reference--group-001.md#canonical-4fb66f72353ad83855bbc0c589bb2a3760a3bc23bb5b1eeac1566181f0f86ff7)
- [rule_list.rules.dst_asn_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-706597db7df570db9b0b4cd7a78fb6e3babe1406bf022494afadc916c7080ab0)
- [rule_list.rules.dst_asn_set](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3f47581361ba06ed76d3deb73ca2d5cf4bc7b261affc60b920ca5520b3acb426)
- [rule_list.rules.dst_ip_prefix_set](data-sources--forward_proxy_policy--reference--group-001.md#canonical-f73c4867739bb963eb34eb55401f1bbd8e31b79daac42de1ae283fb9e96239e5)
- [rule_list.rules.dst_label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3954a30ebacac371bcfc202b6b55edd70767645c81efaeaec4021c0cb55c0c9b)
- [rule_list.rules.dst_prefix_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-72245fe962e5b6889eb3a046890cc1dcc379f4c22ef1c92a1c9593ba70b878af)
- [rule_list.rules.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-53c0c93d85e18cb9293b46690a42495708c4c986c187adb960a5c59403324aa5)
- [rule_list.rules.ip_prefix_set](data-sources--forward_proxy_policy--reference--group-001.md#canonical-a64dcaf4fc73228e2330d5dc43f9eada4b5479060cd46329952dd3abae9d19c0)
- [rule_list.rules.label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-191cbbf56360c30f7ec31fb039938030568c4ebd5ffb6d63a60e602512e73ade)
- [rule_list.rules.metadata](data-sources--forward_proxy_policy--reference--group-001.md#canonical-eb4a1b3f1fef05e208b45527537b7150b36b9d43d8b078f8235b2d08054489f6)
- [rule_list.rules.no_http_connect_port](data-sources--forward_proxy_policy--reference--group-001.md#canonical-fa532b4ab7079e75055b76fe3dce0c8e703dbd31dd8628e80bc76d3abbd5bedb)
- [rule_list.rules.port_matcher](data-sources--forward_proxy_policy--reference--group-001.md#canonical-f3acd89a9c24f235c071cff02fad14cf28080f8710b3b3987342457db5f5fd02)
- [rule_list.rules.prefix_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e671661533b1c895ca42eaefc2878d61a1f8bcbea5b44128cac37544b9652f6d)
- [rule_list.rules.tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-dc35779a77d04fbbc6df18bd4839b93d8fb3f8193587742cf7f7b4bf4be31c79)
- [rule_list.rules.url_category_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2f88fa959a26b990900806339f765a4bb2976cb74d72bbf01601682ae0d08cee)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-1e1f9fe39b184e3dac247e475c54c2a85ebbc00ab5248d6cdf79327ad7f2ed93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-763ce92ebecb2e8aad791fc0485b9b7c516fe9f302b3475b84d0c8b05379e806"></a>

## rule_list.rules.all_destinations — rule_list.rules.all_destinations / 9024550adf9b / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- rule_list.rules.all_destinations

<a id="canonical-e25daed81ddd8568fd23adecc6e0a76cffc60e55a49d46ff7cfecc7982632263"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-edf210a892d321caf2ffaaeacc90f5a32988070aeeced7337f62040a395de162"></a>

## Direct properties — rule_list.rules.all_destinations / 9024550adf9b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3159c74cd9ff19e9dbdf6fad1ca468f2c1decad7599f5cf57176d45afbe11c97"></a>

## Next pages — rule_list.rules.all_destinations / 9024550adf9b / 4

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-4fb66f72353ad83855bbc0c589bb2a3760a3bc23bb5b1eeac1566181f0f86ff7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-271b73515c25cb892a1d083fed34f65ff4f8597c879fb6ac160cabf700d91b10"></a>

## rule_list.rules.all_sources — rule_list.rules.all_sources / 57d56369356f / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- rule_list.rules.all_sources

<a id="canonical-2eea4b677d92a3c97605731c95c334be99ff1e9db2f5021264455e70ce9ea1a9"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-425688a11ea722f742f87a773dce09a357905fbceeaf472fdd6d747b20d34bd0"></a>

## Direct properties — rule_list.rules.all_sources / 57d56369356f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-21a17e9e626c56f7e03a9b120f2e1ec15c4a4c9195b2d72dd7e4742dd98bd9bf"></a>

## Next pages — rule_list.rules.all_sources / 57d56369356f / 4

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-706597db7df570db9b0b4cd7a78fb6e3babe1406bf022494afadc916c7080ab0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d5ff616daab9f262b1c999ebfb684c8938cf7860c2e0a4b27b751eb6d09512a5"></a>

## rule_list.rules.dst_asn_list — rule_list.rules.dst_asn_list / 39e53f38f1d9 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- rule_list.rules.dst_asn_list

<a id="canonical-333215d9536e5df8896be3f8d4334bc8cad76139029ca951b81b755f1e68a558"></a>

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

<a id="canonical-abc33a917719bd880d7e3cf13b6164b45d5c68a197ffca149502fa9bbe014b56"></a>

## Direct properties — rule_list.rules.dst_asn_list / 39e53f38f1d9 / 3

<a id="canonical-83b1a2c5999e46cd5734b6a98bb1fc15db6870fe9df1e0160d44f2195dbf5cff"></a>

<a id="canonical-df14b576eeddcfa0530be313e56e578fe6398d1088075dcdc9dec682e9488160"></a>

## as_numbers property — rule_list.rules.dst_asn_list / 39e53f38f1d9 / 4

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

<a id="canonical-5e014600619171f278de42c59f94734835a09379671488248add3d345b9f0663"></a>

## Next pages — rule_list.rules.dst_asn_list / 39e53f38f1d9 / 5

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-3f47581361ba06ed76d3deb73ca2d5cf4bc7b261affc60b920ca5520b3acb426"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9bd72ac09310db8aed149a22cc37462d31b9b4959bb3e4e9f27fde9089989d17"></a>

## rule_list.rules.dst_asn_set — rule_list.rules.dst_asn_set / 9193dbe6faac / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- rule_list.rules.dst_asn_set

<a id="canonical-d5d668f7f54b4e808e6b94dcb2d4e27d51aea2096968a382345497fcd46949fa"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-c6ca04e35999a26c60395b79810f4d9494c1f198fc0d54228228b884cbb13dc9"></a>

## Direct properties — rule_list.rules.dst_asn_set / 9193dbe6faac / 3

<a id="canonical-51a7c60f7f23b5937b084a82923ab6b85e5364eaa019b867d6b9594c647bf635"></a>

<a id="canonical-5dfe5e7dfc4330f61514a2e7f22b6391efe184c3e83b99e45d2a6b6874306106"></a>

## name property — rule_list.rules.dst_asn_set / 9193dbe6faac / 4

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

<a id="canonical-982ef63eb7863f99e9cc7032f751bb9eea0cf5ccd61890dac38868cca91e8222"></a>

<a id="canonical-389a3943edb2c2e16496d4645fec5fac05358314f72767734d5e4bbab3240aac"></a>

## namespace property — rule_list.rules.dst_asn_set / 9193dbe6faac / 5

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

<a id="canonical-5a30db9b460e5c2e4f2e9bafc0fce3195f0071337b9fa6bbe153ec9425017761"></a>

<a id="canonical-38e67df63478c17f5bd4d4e406ccf8f20e08bb428d1c0a04eb55482e70957da3"></a>

## tenant property — rule_list.rules.dst_asn_set / 9193dbe6faac / 6

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

<a id="canonical-42dd89ee2d8ee9372361e901880c74ab286b347f2eda328d390d4b67d20fcb33"></a>

## Next pages — rule_list.rules.dst_asn_set / 9193dbe6faac / 7

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-f73c4867739bb963eb34eb55401f1bbd8e31b79daac42de1ae283fb9e96239e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e828b3ec8f3096d2513a65423abddfd85f7f1468842c87b3d79d2992bf86307"></a>

## rule_list.rules.dst_ip_prefix_set — rule_list.rules.dst_ip_prefix_set / f6e8257c0b36 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- rule_list.rules.dst_ip_prefix_set

<a id="canonical-53742d6e78a01bb376b6d13dfa1a09aeebadadde8e352604fae41932910f2af6"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-1b6d4bc02eb236d1023d1219e944ba094ef762ed58d506cc0ee8b6be264d78eb"></a>

## Direct properties — rule_list.rules.dst_ip_prefix_set / f6e8257c0b36 / 3

<a id="canonical-ba4b30c48abb65f17ddca3a85de44beb70cf87c5eddc587f0857ecb45e649fbb"></a>

<a id="canonical-e9900b0d4f109808ca8c326fc81a205fdcd4940ae3d23ec6110fc0756a36b1cc"></a>

## name property — rule_list.rules.dst_ip_prefix_set / f6e8257c0b36 / 4

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

<a id="canonical-7c979712df77ab02ed82e52f44d0c68176d6705da55a4273bf25378a0846ac78"></a>

<a id="canonical-fb8e744b90a7e672031e8faf0f3dbf3ceb356eb154697f7f0f0d621011e040d2"></a>

## namespace property — rule_list.rules.dst_ip_prefix_set / f6e8257c0b36 / 5

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

<a id="canonical-aed70415cdb9e43f4f8bf847c6e260f035ddffa622ef8b6ad876409b295eaec6"></a>

<a id="canonical-f68ff798977a296316943d47003e81bf5cdd07ddf08a5fe7457097bbc065f2d2"></a>

## tenant property — rule_list.rules.dst_ip_prefix_set / f6e8257c0b36 / 6

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

<a id="canonical-a68be234c79d2a7d57ecf1d0ef80f10318cdc34bd608bfb21cf34e70af6ff086"></a>

## Next pages — rule_list.rules.dst_ip_prefix_set / f6e8257c0b36 / 7

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-3954a30ebacac371bcfc202b6b55edd70767645c81efaeaec4021c0cb55c0c9b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21a2c04303d7c7c5a6fba8abf2c9949429ac59cf72fea5cad04cae62b70dda0f"></a>

## rule_list.rules.dst_label_selector — rule_list.rules.dst_label_selector / 9cab6197836d / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- rule_list.rules.dst_label_selector

<a id="canonical-21878ccbd04f985a46d0bf7cad9e52a054db197942006ea35928a2b080e2629b"></a>

Type: `"single"`. Computed.

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

<a id="canonical-7c270acce7c9f64e4c6459648c4d31eb200ef5c1297a7f9220a1e940dadd0c2b"></a>

## Direct properties — rule_list.rules.dst_label_selector / 9cab6197836d / 3

<a id="canonical-b6d3b69e5c2953fd7ac9db2189ecd0dfc4aea01841f1ca50a4990c281e955bcd"></a>

<a id="canonical-ad0f654031635e1248c4005acd6eed5f486ba61ad61f63185785ef767e95ee9a"></a>

## expressions property — rule_list.rules.dst_label_selector / 9cab6197836d / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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

<a id="canonical-afa324ae4ead05dce1d0a6aaf1a5e486d378bee89306e0f7ad9543af9cfd6dde"></a>

## Next pages — rule_list.rules.dst_label_selector / 9cab6197836d / 5

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-72245fe962e5b6889eb3a046890cc1dcc379f4c22ef1c92a1c9593ba70b878af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e2db1a84de739de6c67bf7fc46d41649db65ba163ee72a3a3bc4fd415b07f2d"></a>

## rule_list.rules.dst_prefix_list — rule_list.rules.dst_prefix_list / ba88a130f2b7 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- rule_list.rules.dst_prefix_list

<a id="canonical-2b30a843a08209b542fa1f2c3fbed3ff4ab768fce41c0785547afedda62b1bac"></a>

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

<a id="canonical-381fbeb3d24f12028c745898028f2b3ab1acdf99564976962cfcdc25251fbb9f"></a>

## Direct properties — rule_list.rules.dst_prefix_list / ba88a130f2b7 / 3

<a id="canonical-c5fc55803175e601aedaea7b2f2ff579e3b0b27870561ca75a9179ff06e85ace"></a>

<a id="canonical-d8cabd59679d8a981ee1b50a80aa406d63ecd6f3bc9ae1768a0b6cd1918b8b64"></a>

## prefixes property — rule_list.rules.dst_prefix_list / ba88a130f2b7 / 4

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

<a id="canonical-45586aebbff8908679c09270d5c5b5fe0dad9ce33d85b6ca5fe1779b8c8b6ce3"></a>

## Next pages — rule_list.rules.dst_prefix_list / ba88a130f2b7 / 5

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-53c0c93d85e18cb9293b46690a42495708c4c986c187adb960a5c59403324aa5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d94953fa5adc3e43534e5f83b708342dbc5e79753d26db9c5c39e09e1f7ceafd"></a>

## rule_list.rules.http_list — rule_list.rules.http_list / e25707fad5bf / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- rule_list.rules.http_list

<a id="canonical-37e8a0f64525e0afaf5c3182c0b4d7a1bcaa39bcaef548c91ed346269714819a"></a>

Type: `"single"`. Computed.

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

<a id="canonical-5e22a7167ccf9d6f5a2bcce93acc3919fa2a70cb340ffa22694cd0b749023961"></a>

## Direct properties — rule_list.rules.http_list / e25707fad5bf / 3

- [http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-43e79f013374e3272eeef706a3da3e5f7910018d955e5b8c98cb6485f16668a3): complete subsection reference.

<a id="canonical-e032dd483d9ae0ca35540caca21471ae8496ab6b053ffc7fbff0fd3766daccb9"></a>

## Next pages — rule_list.rules.http_list / e25707fad5bf / 4

- [rule_list.rules.http_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-43e79f013374e3272eeef706a3da3e5f7910018d955e5b8c98cb6485f16668a3)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-43e79f013374e3272eeef706a3da3e5f7910018d955e5b8c98cb6485f16668a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-665018df8ba8d5849582139786eaaf00133e52cc6d3bafb0e65041f957705ce5"></a>

## rule_list.rules.http_list.http_list — rule_list.rules.http_list.http_list / d136462d364a / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [rule_list.rules.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-53c0c93d85e18cb9293b46690a42495708c4c986c187adb960a5c59403324aa5)
- rule_list.rules.http_list.http_list

<a id="canonical-bede8880f643c226d60c2427d30d67aaf77e9074d6fd98c66cf241a77c1d9fda"></a>

Type: `"list"`. Computed.

HTTP URLs. URLs for HTTP connections.

Upstream description:

URLs for HTTP connections.

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

<a id="canonical-3525ffca0c42bbdd86597cae391e7f51dfb5b55a59e0e8b582747c3f20b6881e"></a>

## Direct properties — rule_list.rules.http_list.http_list / d136462d364a / 3

- [any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-ff292158ee620c64412b005c3aa74c8c7fce692027e6edd3eb5c8c68ac33af37): complete subsection reference.

<a id="canonical-d55bf7297377d30e61a6861d3783ca249aa181a5466638e89f9464fcea78e8e3"></a>

<a id="canonical-757b608f29032af7d24daa1f32a044e0190369181b8d0b97a2663769e9ec5dd3"></a>

## exact_value property — rule_list.rules.http_list.http_list / d136462d364a / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-e22165b11cc4ef47eb348192a9a9502add48d683d40e2f032e875cc581aa2e00"></a>

<a id="canonical-d3e2221fdc46170666579e6a487b891fef19163502b4c039f907140c160769d8"></a>

## path_exact_value property — rule_list.rules.http_list.http_list / d136462d364a / 5

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Upstream description:

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

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

<a id="canonical-23cafbf3d1ffba31da45a6c19dfa6d939457d75b2c30c8dd44c8ac2236c9364f"></a>

<a id="canonical-a6e64842ad39ccf645542ca4016a00e7a3f8b20c03a59f09dfc79ae138282f82"></a>

## path_prefix_value property — rule_list.rules.http_list.http_list / d136462d364a / 6

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g "/abc/xyz"
will match "/abc/xyz/.\*"

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

<a id="canonical-910bdbd1996f6327667ca4b043a59e2477b467215a2db31cbcb7f0d08a26598b"></a>

<a id="canonical-0443771a059c87e290a852ecbb469b2aed984b146785bedfb2fb86354a62e416"></a>

## path_regex_value property — rule_list.rules.http_list.http_list / d136462d364a / 7

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

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

<a id="canonical-ad3c7d317013635cd3f781e53e44f04ea7db6bb3b1b83e30245c58839ad79652"></a>

<a id="canonical-167d6d74bf20cd9bb384b6e48bd1944862eea8ba4fdc83a4cacda6cd63755dda"></a>

## regex_value property — rule_list.rules.http_list.http_list / d136462d364a / 8

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-aa3e232113da36d95cf0dadaede54ed3e1f2b6c97a3b0c1c007014954698e604"></a>

<a id="canonical-96d137438d8d237c1c214d2e63e31e43a28715b17060c9ce14b36b0ba8c1784d"></a>

## suffix_value property — rule_list.rules.http_list.http_list / d136462d364a / 9

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain names e.g "xyz.com" will match
"\*.xyz.com"

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

<a id="canonical-33b99719351fa99497542610c8a893c1212bd9bcd716b1bf1b78bfa023a2573b"></a>

## Next pages — rule_list.rules.http_list.http_list / d136462d364a / 10

- [rule_list.rules.http_list.http_list.any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-ff292158ee620c64412b005c3aa74c8c7fce692027e6edd3eb5c8c68ac33af37)
- [rule_list.rules.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-53c0c93d85e18cb9293b46690a42495708c4c986c187adb960a5c59403324aa5)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-ff292158ee620c64412b005c3aa74c8c7fce692027e6edd3eb5c8c68ac33af37"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f25b6ed8a05cff181e9b7a76ba41a59adb2eb7d46f355cda95bde0dc280cf1c"></a>

## rule_list.rules.http_list.http_list.any_path — rule_list.rules.http_list.http_list.any_path / 053f0809585b / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [rule_list.rules.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-53c0c93d85e18cb9293b46690a42495708c4c986c187adb960a5c59403324aa5)
- [rule_list.rules.http_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-43e79f013374e3272eeef706a3da3e5f7910018d955e5b8c98cb6485f16668a3)
- rule_list.rules.http_list.http_list.any_path

<a id="canonical-40a7b91281b804898af9d7e0f095ff45f671588c4b2a0913055a435bda38f89d"></a>

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

<a id="canonical-5af83bf3929310e97414f6394c7aca51104febea1ed19b0e422cdbf1f9a92f5f"></a>

## Direct properties — rule_list.rules.http_list.http_list.any_path / 053f0809585b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8f110d0632a5b82fc43ffa1c06435356ecd162d609bfc839b5fee624ef6f00eb"></a>

## Next pages — rule_list.rules.http_list.http_list.any_path / 053f0809585b / 4

- [rule_list.rules.http_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-43e79f013374e3272eeef706a3da3e5f7910018d955e5b8c98cb6485f16668a3)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-a64dcaf4fc73228e2330d5dc43f9eada4b5479060cd46329952dd3abae9d19c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fcb250e5d7d4390a225fbf53fafe1c3ffb2309ce80da570f32a8e7f9a5f6487"></a>

## rule_list.rules.ip_prefix_set — rule_list.rules.ip_prefix_set / 4361897ced57 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- rule_list.rules.ip_prefix_set

<a id="canonical-a9f64a439e6f48ee201d434b0244f35d697977cbca00ce0dcb03ea138a349d07"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-df3afce107923bf9bd0df1418292ddac647d4f0a9689643478063527471e621f"></a>

## Direct properties — rule_list.rules.ip_prefix_set / 4361897ced57 / 3

<a id="canonical-6c02534f4804a6ef4636c4eb8d8f186508bd9ae6e98a82816c0bba2983f0a3fe"></a>

<a id="canonical-8ea0aa72cb67fabe0d93c42979a69ef51433c7008cf1d354fdb93d9362cbb7f8"></a>

## name property — rule_list.rules.ip_prefix_set / 4361897ced57 / 4

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

<a id="canonical-ed1fa673e52f4a8666c127d53fde7f71197ae7e293e45f56dbc2f85e65339217"></a>

<a id="canonical-135885d3351756709846706e2da74e2d0efda36242009bc61afd0ea43ff08f32"></a>

## namespace property — rule_list.rules.ip_prefix_set / 4361897ced57 / 5

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

<a id="canonical-1b852a1b9c38e195ef2892927674b908c37a1fb5c6a22c6362f04e35f8ff930d"></a>

<a id="canonical-da166c9489b7ac217615bfdaa55c313019ee2bf6ff0824a5b7d58068d8e7d23e"></a>

## tenant property — rule_list.rules.ip_prefix_set / 4361897ced57 / 6

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

<a id="canonical-2ca63c41ae503a9116aa4f33b97572af288a5cd34926bba3f0ef5e7e8d5fe6e1"></a>

## Next pages — rule_list.rules.ip_prefix_set / 4361897ced57 / 7

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-191cbbf56360c30f7ec31fb039938030568c4ebd5ffb6d63a60e602512e73ade"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e305a1ea8c6ffeb573f329a55749fe09a67d1c8d569b9fee5131351ff0e08fa"></a>

## rule_list.rules.label_selector — rule_list.rules.label_selector / c3edbf54cbe6 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- rule_list.rules.label_selector

<a id="canonical-6e67681bcba54071527a731a0bd5b1a12d420ad9af9adb0e38491c5b4fe338e2"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1cd4b35fac0a7eb7f1860739494140f4e6cb0da7f5c513263136a7872ff2741d"></a>

## Direct properties — rule_list.rules.label_selector / c3edbf54cbe6 / 3

<a id="canonical-fcf8587b69b09134be59258c5baae92b9908f4fdc5bb26ac07d0aa5d6d25c618"></a>

<a id="canonical-68041cff9d3c0fa45e17b70c902ab7903b5ec8df97645098a6bf02f39d99ea98"></a>

## expressions property — rule_list.rules.label_selector / c3edbf54cbe6 / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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

<a id="canonical-d37091c2f1f5560e0a9805fa20558425ab8710a5fcc83ddf4ec2d00d41d5fc48"></a>

## Next pages — rule_list.rules.label_selector / c3edbf54cbe6 / 5

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-eb4a1b3f1fef05e208b45527537b7150b36b9d43d8b078f8235b2d08054489f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1da424b943b40b0cbdcbae35d94456309318db645254602520a39d437140c70c"></a>

## rule_list.rules.metadata — rule_list.rules.metadata / b94ee1e0fb70 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- rule_list.rules.metadata

<a id="canonical-64074700e41b5bdfe391da6c9af50b536e0ba854f611c4a801441225161b2974"></a>

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

<a id="canonical-302b1246ea81d6167d4fc4aa5ed3dd29ebd2d18fa09376fc3ac085761a843b35"></a>

## Direct properties — rule_list.rules.metadata / b94ee1e0fb70 / 3

<a id="canonical-ca6f97210eb701bac3ca2483f011a4d935be5b4eb84954060e90275fc39cc768"></a>

<a id="canonical-fd4032ec3314dc92d9c56afcb928d9f6c8055ecf29cb9dbcb72ce4e1d4059690"></a>

## description_spec property — rule_list.rules.metadata / b94ee1e0fb70 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-a0e9efb3bbc7638e5be78a074dc6f86c8875cd5c8781c0b555d94ec5b660b609"></a>

<a id="canonical-974af6cf160687436b645cee7840f423d562551ee818693ff624fefc27cd671b"></a>

## name property — rule_list.rules.metadata / b94ee1e0fb70 / 5

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

<a id="canonical-782b8e1a47a4ccb437bf7e89a10a9491f66f9975fd1e76950c1de6738b0d9404"></a>

## Next pages — rule_list.rules.metadata / b94ee1e0fb70 / 6

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-fa532b4ab7079e75055b76fe3dce0c8e703dbd31dd8628e80bc76d3abbd5bedb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f75f4c14bf57d391215947c09725f6df4b374b531bcfb215d046a3f1a09c142d"></a>

## rule_list.rules.no_http_connect_port — rule_list.rules.no_http_connect_port / ff01a59f30a6 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- rule_list.rules.no_http_connect_port

<a id="canonical-500e60d5bbc2bc6f4db0cc74e696f493c50b38f01692c788c4a2f7722f67240e"></a>

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

<a id="canonical-9323adb6bdb72367d7869564af0d07ac19266f9c88c6bb7159a8932dac34b497"></a>

## Direct properties — rule_list.rules.no_http_connect_port / ff01a59f30a6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3767a2c61db42789ea7f09b5e711f5dc93b8d53eb866486b2198ea9c60a91927"></a>

## Next pages — rule_list.rules.no_http_connect_port / ff01a59f30a6 / 4

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-f3acd89a9c24f235c071cff02fad14cf28080f8710b3b3987342457db5f5fd02"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16b8b71361b1944ed0745f7757be562998bb55c90cca1af8fefacd6a5e84fd2d"></a>

## rule_list.rules.port_matcher — rule_list.rules.port_matcher / a91d8e13fb1a / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- rule_list.rules.port_matcher

<a id="canonical-045644c82174dbc6ca8f3294f8339f1a7080fc7607dc29ac139f5d353296a421"></a>

Type: `"single"`. Computed.

Port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true.

Upstream description:

A port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true.

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

<a id="canonical-f3c6879c440210f596a4f2c567b4a1d247946e9f148b199896d25549328f35b1"></a>

## Direct properties — rule_list.rules.port_matcher / a91d8e13fb1a / 3

<a id="canonical-21209ac6b113e9e91266692498f14b4e83bd7e8e19ac51584683b8edc07ae3ef"></a>

<a id="canonical-a4fc9272957446bb94195772b15306cb4f8ad524ed140f3e6c06701e3848522b"></a>

## invert_matcher property — rule_list.rules.port_matcher / a91d8e13fb1a / 4

Type: `"bool"`. Computed.

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

<a id="canonical-b41f779252f5fe842bd97c7d570bb6bf6a2d022b6504681ab02cdd1d4caa0fe4"></a>

<a id="canonical-209ce333c37ef7834d3722dda4fa1dd3a38f9ce75ed37f06edb2daa9742f1072"></a>

## ports property — rule_list.rules.port_matcher / a91d8e13fb1a / 5

Type: `["list", "string"]`. Computed.

List of strings, each of which is a single port value or a tuple of start and end port values
separated by '-'. The start and end values are considered to be part of the range.

Upstream description:

A list of strings, each of which is a single port value or a tuple of start and end port values
separated by "-". The start and end values are considered to be part of the range.

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

<a id="canonical-d31dcb9ae3d2d6dcf61b2e60af5cacfecfb04c8de1ac8b252462cbc024a8aec2"></a>

## Next pages — rule_list.rules.port_matcher / a91d8e13fb1a / 6

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-e671661533b1c895ca42eaefc2878d61a1f8bcbea5b44128cac37544b9652f6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df7ee2da3f448c37de409782341556de58252cabebcf7ed697f56c55462bef72"></a>

## rule_list.rules.prefix_list — rule_list.rules.prefix_list / f0defe934b44 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- rule_list.rules.prefix_list

<a id="canonical-05d66cb54e790bf9cc85d6fe33b11507b1f9b815b3dfd5a8ad216dbd5e9eabff"></a>

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

<a id="canonical-4f0251912c0996d3cc3d8ad2f0f211a74562898f032410a67d0f7b2f44b7e2c5"></a>

## Direct properties — rule_list.rules.prefix_list / f0defe934b44 / 3

<a id="canonical-033da1c676e5549095e7cc0b159f4f3d6f19efae70a86a7d2003535fdf61f652"></a>

<a id="canonical-8d4d00f9babf7dc8896e2f2d7e714622dd8ed74a9a3b97d281fdbb6ccebe1397"></a>

## prefixes property — rule_list.rules.prefix_list / f0defe934b44 / 4

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

<a id="canonical-9b59b4dd217a4727a903221775fc35a0ae2e43476f0f4fccbcdf3b8beb9253e0"></a>

## Next pages — rule_list.rules.prefix_list / f0defe934b44 / 5

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-dc35779a77d04fbbc6df18bd4839b93d8fb3f8193587742cf7f7b4bf4be31c79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85e061eab7ff48dd774ae43b503bf8004b39438c83ac1fc64f7725e22f125ee7"></a>

## rule_list.rules.tls_list — rule_list.rules.tls_list / d782d19b8ff2 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- rule_list.rules.tls_list

<a id="canonical-34cdf7e2db22f4ff3ccfe9eed291c07123e3b26e7f1ef2aa8b80e9228b0c3c7c"></a>

Type: `"single"`. Computed.

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

<a id="canonical-497f588c7f6edcc7ec4da9014b8a498e5b0f67dacdc3a2fcf0806140ee8f89d4"></a>

## Direct properties — rule_list.rules.tls_list / d782d19b8ff2 / 3

- [tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-c4e1de284a057bb8d277cef7204824aedf667dce4b6b2131e5cc1218766d712f): complete subsection reference.

<a id="canonical-03b22be5c8b2f1f596cfa4c57952b267b8ac369957a7a1da34da9c7907fcc044"></a>

## Next pages — rule_list.rules.tls_list / d782d19b8ff2 / 4

- [rule_list.rules.tls_list.tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-c4e1de284a057bb8d277cef7204824aedf667dce4b6b2131e5cc1218766d712f)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-c4e1de284a057bb8d277cef7204824aedf667dce4b6b2131e5cc1218766d712f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23c5274b721849cc0121d814b3f95ed841d5fe46c96843c5f0838a81014485dd"></a>

## rule_list.rules.tls_list.tls_list — rule_list.rules.tls_list.tls_list / 02280bd550ff / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [rule_list.rules.tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-dc35779a77d04fbbc6df18bd4839b93d8fb3f8193587742cf7f7b4bf4be31c79)
- rule_list.rules.tls_list.tls_list

<a id="canonical-a9a6a6162c26da1bf1c19164234120c4ad4a6d6aa861a91ac197e9feca092839"></a>

Type: `"list"`. Computed.

TLS Domains. Domains in SNI for TLS connections.

Upstream description:

Domains in SNI for TLS connections.

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

<a id="canonical-a8a50d909a66ca06a4e12941361690d751dd0ea79b8c574eab09d4b47d4ec5de"></a>

## Direct properties — rule_list.rules.tls_list.tls_list / 02280bd550ff / 3

<a id="canonical-6dbb2e1c8bea4085b830ef8e456d5fe73fd07775e443dc9fe57d30fafdf50039"></a>

<a id="canonical-0b299c01889729ea709c3d619041f17fac1937983171600b42f45869750a8ea0"></a>

## exact_value property — rule_list.rules.tls_list.tls_list / 02280bd550ff / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-dd64b62407fd5d95baea29990a6940ec517c121c5d1a8a7f4018fcd510d7faa7"></a>

<a id="canonical-9e093956a5805a042eb60246d9095a765a5e5d03bec375103543842cb4959d1d"></a>

## regex_value property — rule_list.rules.tls_list.tls_list / 02280bd550ff / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-3b1e2ae35d14f7a00e84f092856387e361677971001e2743323d85c1e4c0212f"></a>

<a id="canonical-8f852e6c4deb6d273b44031be0ffe4bd75e66912fa7255c90feeb6dcb66ff796"></a>

## suffix_value property — rule_list.rules.tls_list.tls_list / 02280bd550ff / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-7a8d1f490f7a8b0fda8067fcac80444a788d0325298dc644c0bae2e0fe38afec"></a>

## Next pages — rule_list.rules.tls_list.tls_list / 02280bd550ff / 7

- [rule_list.rules.tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-dc35779a77d04fbbc6df18bd4839b93d8fb3f8193587742cf7f7b4bf4be31c79)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)

<a id="canonical-2f88fa959a26b990900806339f765a4bb2976cb74d72bbf01601682ae0d08cee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5aab38f3a5891fd684000c25596dc176263cc21af6bca7736671115cb2d64db1"></a>

## rule_list.rules.url_category_list — rule_list.rules.url_category_list / a4774cc13066 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2177cc807ea38c6a2db1f7084b0f98c89ca98df01797a62075f8b417a4d895f6)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-e410c4aa7d4dc5a78f208caa967f09f45d55ebe0de66cd2110abb6f512cdcb2e)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- rule_list.rules.url_category_list

<a id="canonical-e519f0890da74ec26733c12bb8c3e460897063b69853b6bfcd0c125e59ee7b64"></a>

Type: `"single"`. Computed.

URL Category List Type. List of URL categories.

Upstream description:

List of URL categories.

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

<a id="canonical-39624392804ee34fc1218bdc4cfcdf36ecce7aae7e5e52d9e5b50cf5ada20982"></a>

## Direct properties — rule_list.rules.url_category_list / a4774cc13066 / 3

<a id="canonical-9da500269c7802f29d0088384d5a025302bc33761d0f0f2c1e968a106dc1f990"></a>

<a id="canonical-6cd3b12bfd0ce4f00970a622235a13730ed8e73bc29cc50432e4ffd45b0ba743"></a>

## url_categories property — rule_list.rules.url_category_list / a4774cc13066 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
UNCATEGORIZED|REAL\_ESTATE|COMPUTER\_AND\_INTERNET\_SECURITY|FINANCIAL\_SERVICES|BUSINESS\_AND\_ECONOMY|COMPUTER\_AND\_INTERNET\_INFO|AUCTIONS|SHOPPING|CULT\_AND\_OCCULT|TRAVEL|ABUSED\_DRUGS|ADULT\_AND\_PORNOGRAPHY|HOME\_AND\_GARDEN|MILITARY|SOCIAL\_NETWORKING|DEAD\_SITES|INDIVIDUAL\_STOCK\_ADVICE\_AND\_TOOLS|TRAINING\_AND\_TOOLS|DATING|SEX\_EDUCATION|RELIGION|ENTERTAINMENT\_AND\_ARTS|PERSONAL\_SITES\_AND\_BLOGS|LEGAL|LOCAL\_INFORMATION|STREAMING\_MEDIA|JOB\_SEARCH|GAMBLING|TRANSLATION|REFERENCE\_AND\_RESEARCH|SHAREWARE\_AND\_FREEWARE|PEER\_TO\_PEER|MARIJUANA|HACKING|GAMES|PHILOSOPHY\_AND\_POLITICAL\_ADVOCACY|WEAPONS|PAY\_TO\_SURF|HUNTING\_AND\_FISHING|SOCIETY|EDUCATIONAL\_INSTITUTIONS|ONLINE\_GREETING\_CARDS|SPORTS|SWIMSUITS\_AND\_INTIMATE\_APPAREL|QUESTIONABLE|KIDS|HATE\_AND\_RACISM|PERSONAL\_STORAGE|VIOLENCE|KEYLOGGERS\_AND\_MONITORING|SEARCH\_ENGINES|INTERNET\_PORTALS|WEB\_ADVERTISEMENTS|CHEATING|GROSS|WEB\_BASED\_EMAIL|MALWARE\_SITES|PHISHING\_AND\_OTHER\_FRAUDS|PROXY\_AVOIDANCE\_AND\_ANONYMIZERS|SPYWARE\_AND\_ADWARE|MUSIC|GOVERNMENT|NUDITY|NEWS\_AND\_MEDIA|ILLEGAL|CONTENT\_DELIVERY\_NETWORKS|INTERNET\_COMMUNICATIONS|BOT\_NETS|ABORTION|HEALTH\_AND\_MEDICINE|CONFIRMED\_SPAM\_SOURCES|SPAM\_URLS|UNCONFIRMED\_SPAM\_SOURCES|OPEN\_HTTP\_PROXIES|DYNAMICALLY\_GENERATED\_CONTENT|PARKED\_DOMAINS|ALCOHOL\_AND\_TOBACCO|PRIVATE\_IP\_ADDRESSES|IMAGE\_AND\_VIDEO\_SEARCH|FASHION\_AND\_BEAUTY|RECREATION\_AND\_HOBBIES|MOTOR\_VEHICLES|WEB\_HOSTING\]
URL Categories. List of URL categories to be selected. Possible values are \`UNCATEGORIZED\`,
\`REAL\_ESTATE\`, \`COMPUTER\_AND\_INTERNET\_SECURITY\`, \`FINANCIAL\_SERVICES\`,
\`BUSINESS\_AND\_ECONOMY\`, \`COMPUTER\_AND\_INTERNET\_INFO\`, \`AUCTIONS\`, \`SHOPPING\`,
\`CULT\_AND\_OCCULT\`, \`TRAVEL\`, \`ABUSED\_DRUGS\`, \`ADULT\_AND\_PORNOGRAPHY\`,
\`HOME\_AND\_GARDEN\`, \`MILITARY\`, \`SOCIAL\_NETWORKING\`, \`DEAD\_SITES\`,
\`INDIVIDUAL\_STOCK\_ADVICE\_AND\_TOOLS\`, \`TRAINING\_AND\_TOOLS\`, \`DATING\`, \`SEX\_EDUCATION\`,
\`RELIGION\`, \`ENTERTAINMENT\_AND\_ARTS\`, \`PERSONAL\_SITES\_AND\_BLOGS\`, \`LEGAL\`,
\`LOCAL\_INFORMATION\`, \`STREAMING\_MEDIA\`, \`JOB\_SEARCH\`, \`GAMBLING\`, \`TRANSLATION\`,
\`REFERENCE\_AND\_RESEARCH\`, \`SHAREWARE\_AND\_FREEWARE\`, \`PEER\_TO\_PEER\`, \`MARIJUANA\`,
\`HACKING\`, \`GAMES\`, \`PHILOSOPHY\_AND\_POLITICAL\_ADVOCACY\`, \`WEAPONS\`, \`PAY\_TO\_SURF\`,
\`HUNTING\_AND\_FISHING\`, \`SOCIETY\`, \`EDUCATIONAL\_INSTITUTIONS\`, \`ONLINE\_GREETING\_CARDS\`,
\`SPORTS\`, \`SWIMSUITS\_AND\_INTIMATE\_APPAREL\`, \`QUESTIONABLE\`, \`KIDS\`,
\`HATE\_AND\_RACISM\`, \`PERSONAL\_STORAGE\`, \`VIOLENCE\`, \`KEYLOGGERS\_AND\_MONITORING\`,
\`SEARCH\_ENGINES\`, \`INTERNET\_PORTALS\`, \`WEB\_ADVERTISEMENTS\`, \`CHEATING\`, \`GROSS\`,
\`WEB\_BASED\_EMAIL\`, \`MALWARE\_SITES\`, \`PHISHING\_AND\_OTHER\_FRAUDS\`,
\`PROXY\_AVOIDANCE\_AND\_ANONYMIZERS\`, \`SPYWARE\_AND\_ADWARE\`, \`MUSIC\`, \`GOVERNMENT\`,
\`NUDITY\`, \`NEWS\_AND\_MEDIA\`, \`ILLEGAL\`, \`CONTENT\_DELIVERY\_NETWORKS\`,
\`INTERNET\_COMMUNICATIONS\`, \`BOT\_NETS\`, \`ABORTION\`, \`HEALTH\_AND\_MEDICINE\`,
\`CONFIRMED\_SPAM\_SOURCES\`, \`SPAM\_URLS\`, \`UNCONFIRMED\_SPAM\_SOURCES\`,
\`OPEN\_HTTP\_PROXIES\`, \`DYNAMICALLY\_GENERATED\_CONTENT\`, \`PARKED\_DOMAINS\`,
\`ALCOHOL\_AND\_TOBACCO\`, \`PRIVATE\_IP\_ADDRESSES\`, \`IMAGE\_AND\_VIDEO\_SEARCH\`,
\`FASHION\_AND\_BEAUTY\`, \`RECREATION\_AND\_HOBBIES\`, \`MOTOR\_VEHICLES\`, \`WEB\_HOSTING\`.
Defaults to \`UNCATEGORIZED\`.

Upstream description:

List of URL categories to be selected.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8c5180a550901e6519c7fd8ec519b4002f2e55fba6a110ad3e58abd2a154fb90"></a>

## Next pages — rule_list.rules.url_category_list / a4774cc13066 / 5

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-de8eca191cdfa750f960533b53b6538deed1999ffb75ff014cb1a0d2492acec3)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-793844117e930b2db3a86c5fb4aaf79df8c6c2ac66cfa51deb8da8e10b5602a3)
