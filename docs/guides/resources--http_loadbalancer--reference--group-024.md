---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-637f391c1800b5412a3d5596d8227123be914df63e54115e7230d7679bd41b83"></a>

## Next pages — routes.simple_route / d33da5f40522 / 6

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [routes.simple_route.auto_host_rewrite](resources--http_loadbalancer--reference--group-025.md#canonical-b1ea5bffa88e49df6b6d0edc383d20212e413371ddad4550259a823a9ab5d46c)
- [routes.simple_route.caching_disable](resources--http_loadbalancer--reference--group-025.md#canonical-22a8a65b22f0d17699f0a1b1dc4ee5689121e75a32990312eafe2b45153aaa62)
- [routes.simple_route.caching_inherit](resources--http_loadbalancer--reference--group-025.md#canonical-ccc21c633f54dfea082575296cbd101fecfe565dd6330551de617b91fe2060cb)
- [routes.simple_route.disable_host_rewrite](resources--http_loadbalancer--reference--group-025.md#canonical-327311533726c2683bb8d901a8bf4bfb447c085a6759378e1611b0e7dbe5d8ca)
- [routes.simple_route.headers](resources--http_loadbalancer--reference--group-025.md#canonical-cd2e0b3b7be57d33b023879c4cb0adb7038b3073026c8e66e357f0c17bb46cd5)
- [routes.simple_route.incoming_port](resources--http_loadbalancer--reference--group-025.md#canonical-89e4f0ec7c7bfa4082cd6048a93e7f2a9beac6a7e7d10bb0eec35c9d8ea7977b)
- [routes.simple_route.origin_pools](resources--http_loadbalancer--reference--group-025.md#canonical-9c2802a3244b62539e829ec07a437c7401f779c02df58f552cc1bcb7860deeae)
- [routes.simple_route.path](resources--http_loadbalancer--reference--group-025.md#canonical-f745313ee2be64f172d4dfaf3bdfb5d37247115fde2b8c46b8e963301c23b8d4)
- [routes.simple_route.query_params](resources--http_loadbalancer--reference--group-025.md#canonical-79028ea11feedc89453fd837580235418e7e524944afa89f9e7523949cc38b51)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c0a8f2e8cf89f9a3425bc382d3cdab782863aac1d1acb696f3491b1d44fc7f5"></a>

## routes.simple_route.advanced_options — routes.simple_route.advanced_options / c00684f09de4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- routes.simple_route.advanced_options

<a id="canonical-d67ac21e1627870eba59cd34b2d9054bdbb75d93c4cf325747d4de84f6380d1f"></a>

Type: `"object"`. single nested block, Optional.

Configure advanced OPTIONS for route like path rewrite, hash policy, etc.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("app_firewall",
    "disable_waf"),
  validators.ConflictingObjectAttributes("app_firewall",
    "inherited_waf"),
  validators.ConflictingObjectAttributes("bot_defense_javascript_injection",
    "inherited_bot_defense_javascript_injection"),
  validators.ConflictingObjectAttributes("buffer_policy",
    "common_buffering"),
  validators.ConflictingObjectAttributes("common_hash_policy",
    "specific_hash_policy"),
  validators.ConflictingObjectAttributes("default_retry_policy",
    "no_retry_policy"),
  validators.ConflictingObjectAttributes("default_retry_policy",
    "retry_policy"),
  validators.ConflictingObjectAttributes("disable_mirroring",
    "mirror_policy"),
  validators.ConflictingObjectAttributes("disable_prefix_rewrite",
    "prefix_rewrite"),
  validators.ConflictingObjectAttributes("disable_prefix_rewrite",
    "regex_rewrite"),
  validators.ConflictingObjectAttributes("disable_spdy",
    "enable_spdy"),
  validators.ConflictingObjectAttributes("disable_waf",
    "inherited_waf"),
  validators.ConflictingObjectAttributes("disable_web_socket_config",
    "web_socket_config"),
  validators.ConflictingObjectAttributes("do_not_retract_cluster",
    "retract_cluster"),
  validators.ConflictingObjectAttributes("inherited_waf_exclusion",
    "waf_exclusion_policy"),
  validators.ConflictingObjectAttributes("no_retry_policy",
    "retry_policy"),
  validators.ConflictingObjectAttributes("prefix_rewrite",
    "regex_rewrite")}
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
  "x-ves-oneof-field-bot_defense_javascript_injection_choice": "[\"bot_defense_javascript_injection\",\"inherited_bot_defense_javascript_injection\"]",
  "x-ves-oneof-field-buffer_choice": "[\"buffer_policy\",\"common_buffering\"]",
  "x-ves-oneof-field-cluster_retract_choice": "[\"do_not_retract_cluster\",\"retract_cluster\"]",
  "x-ves-oneof-field-hash_policy_choice": "[\"common_hash_policy\",\"specific_hash_policy\"]",
  "x-ves-oneof-field-mirroring_choice": "[\"disable_mirroring\",\"mirror_policy\"]",
  "x-ves-oneof-field-retry_policy_choice": "[\"default_retry_policy\",\"no_retry_policy\",\"retry_policy\"]",
  "x-ves-oneof-field-rewrite_choice": "[\"disable_prefix_rewrite\",\"prefix_rewrite\",\"regex_rewrite\"]",
  "x-ves-oneof-field-spdy_choice": "[\"disable_spdy\",\"enable_spdy\"]",
  "x-ves-oneof-field-waf_choice": "[\"app_firewall\",\"disable_waf\",\"inherited_waf\"]",
  "x-ves-oneof-field-waf_exclusion_choice": "[\"inherited_waf_exclusion\",\"waf_exclusion_policy\"]",
  "x-ves-oneof-field-websocket_choice": "[\"disable_web_socket_config\",\"web_socket_config\"]"
}
```

Terraform syntax:

```terraform
advanced_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-5f9deb72d5cb0c788d46da7c739e403a60590b238396453338ac13ae7dc40081"></a>

## Direct properties — routes.simple_route.advanced_options / c00684f09de4 / 3

- [app_firewall](resources--http_loadbalancer--reference--group-024.md#canonical-810d8fde52c13db2d166c0c69c5fa62355d7ca5f7f6ec4643f1a034bac417855): complete subsection reference.

- [bot_defense_javascript_injection](resources--http_loadbalancer--reference--group-024.md#canonical-538f35fdb26f20505920e66daa6a491f8209a103ae0f1667e1a76df9228cb599): complete subsection reference.

- [buffer_policy](resources--http_loadbalancer--reference--group-024.md#canonical-6505c45b2227fa5247d82aed0f887acdba1b2867295ab053ec5922dcd2ba29e3): complete subsection reference.

- [common_buffering](resources--http_loadbalancer--reference--group-024.md#canonical-c87fad1ad82b33b27f1cfb4f23ac91807f59d75f512bd75e22ef8809a83d278e): complete subsection reference.

- [common_hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-688a99d983f0f544d634ee5f3debc14e908e48ce76da3e4afdd5e8d17bd352f9): complete subsection reference.

- [cors_policy](resources--http_loadbalancer--reference--group-024.md#canonical-ea2851e52dea0f4cc37b0e3b2b7f2866c5a2353275e13df98aefca621d6da778): complete subsection reference.

- [csrf_policy](resources--http_loadbalancer--reference--group-024.md#canonical-896b37a975892f4ac16ce875dd53f0b91cafde199fa15ef7205646b622351335): complete subsection reference.

- [default_retry_policy](resources--http_loadbalancer--reference--group-024.md#canonical-c70bb96c80e44e8d22a3fc83f8ad88cbf52f1d1c5ba26e2ab3bcdf242a30d2b5): complete subsection reference.

<a id="canonical-4806d23b8eba2f0e666bba7ac101f458b8f07e2c4b77c4286f539b7ced58282e"></a>

<a id="canonical-a602680cf6f13b51e686d8aa4dbc78be6eccaa9bc34ca1acf213d43d47428c26"></a>

## disable_location_add property — routes.simple_route.advanced_options / c00684f09de4 / 4

Type: `"bool"`. Optional.

Disables append of x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; at route level, if it is
configured at virtual-host level. This configuration is ignored on CE sites.

Upstream description:

Disables append of x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; at route level, if it is
configured at virtual-host level. This configuration is ignored on CE sites.

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

- [disable_mirroring](resources--http_loadbalancer--reference--group-024.md#canonical-6b9287834e37574a29fc4deef0b95e7fc82d04b3c0d7f8f5bb29ce9871df086e): complete subsection reference.

- [disable_prefix_rewrite](resources--http_loadbalancer--reference--group-024.md#canonical-d1a99f6d0246449840b524a894b6c6e69506dd9e99962be170a13ac38b86a61b): complete subsection reference.

- [disable_spdy](resources--http_loadbalancer--reference--group-024.md#canonical-7f5bd37ed4b125a95b4a6a6a78ffd31492c0efaac4498820209a25d8e9d9aa06): complete subsection reference.

- [disable_waf](resources--http_loadbalancer--reference--group-024.md#canonical-86ea4481408913913118621faa19d6328f73ba0b1d572e6e53a122ba1cf05e95): complete subsection reference.

- [disable_web_socket_config](resources--http_loadbalancer--reference--group-024.md#canonical-5b1d0aa3c6df13a9bfc188198d2f3ef6bc4bc899f798d9627066e6aeec7cd08c): complete subsection reference.

- [do_not_retract_cluster](resources--http_loadbalancer--reference--group-024.md#canonical-c7416844f81270e4944c8b9d6fc8b2ff78573ac4688507ccbaf3c48455cf1fff): complete subsection reference.

- [enable_spdy](resources--http_loadbalancer--reference--group-024.md#canonical-e8a531cc4f2e5226b7028c8c1699b8eaab5039187c5316bf181e6e208570698d): complete subsection reference.

- [endpoint_subsets](resources--http_loadbalancer--reference--group-024.md#canonical-9bb66d001ca2632fbea9d56bfc0d314b25683f8620bcb151ee2540a7326b56dd): complete subsection reference.

- [inherited_bot_defense_javascript_injection](resources--http_loadbalancer--reference--group-024.md#canonical-d1c42952d17231babf05a5d044253f9523a24b246c3d2dda52595ad9a4611a80): complete subsection reference.

- [inherited_waf](resources--http_loadbalancer--reference--group-024.md#canonical-278571ba5269ca7afe639f5af6e0b7cd47eff71c0f57349adfbb6c210b620d3b): complete subsection reference.

- [inherited_waf_exclusion](resources--http_loadbalancer--reference--group-024.md#canonical-821811c0f4d76445f73d7e659d3d92d74205dd708146d942d0348fba9e1a632d): complete subsection reference.

- [mirror_policy](resources--http_loadbalancer--reference--group-024.md#canonical-47069f6847d19857cb08c8f8cc04bf77d1cdf5df9600e52489a86f8af3f1ebcb): complete subsection reference.

- [no_retry_policy](resources--http_loadbalancer--reference--group-024.md#canonical-a8fd9cd2de2bff8f95eb44dc2caaede6525e0322e23b32ff98f07a0718da8c15): complete subsection reference.

<a id="canonical-22e9f0efbb7299453d1f9892c1d2de0ac8c5b695a1fff6db05f3922981cb35ad"></a>

<a id="canonical-0583665b8f35977c4791659c0f74a6267477cb47b34560f53b3271402fee33fa"></a>

## prefix_rewrite property — routes.simple_route.advanced_options / c00684f09de4 / 5

Type: `"string"`. Optional.

Exclusive with \[disable\_prefix\_rewrite regex\_rewrite\] prefix\_rewrite indicates that during
forwarding, the matched prefix (or path) should be swapped with its value. When using regex path
matching, the entire path (not including the query string) will be swapped with this value.

Upstream description:

Exclusive with \[disable\_prefix\_rewrite regex\_rewrite\] prefix\_rewrite indicates that during
forwarding, the matched prefix (or path) should be swapped with its value. When using regex path
matching, the entire path (not including the query string) will be swapped with this value.

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

<a id="canonical-14e5945f09868e761552e2c75663800474ee12b34227c484caf252fc33eddd8d"></a>

<a id="canonical-6f07701fec46da2e8817fbdeb5480d40589485d7cbddcf005c5c051dae8bd92f"></a>

## priority property — routes.simple_route.advanced_options / c00684f09de4 / 6

Type: `"string"`. Optional.

\[Enum: DEFAULT|HIGH\] Priority routing for each request. Different connection pools are used based
on the priority selected for the request. Also, circuit-breaker configuration at destination cluster
is chosen based on selected priority. Possible values are \`DEFAULT\`, \`HIGH\`. Defaults to
\`DEFAULT\`.

Upstream description:

Priority routing for each request. Different connection pools are used based on the priority
selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based
on selected priority.

Default routing mechanism High-Priority routing mechanism.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DEFAULT",
    "HIGH"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DEFAULT",
  "enum": [
    "DEFAULT",
    "HIGH"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [regex_rewrite](resources--http_loadbalancer--reference--group-024.md#canonical-381158a64be0f1d30cd5a2dd7b735d6bdc6d12ea248a26fc03ccab67f33b613e): complete subsection reference.

- [request_cookies_to_add](resources--http_loadbalancer--reference--group-024.md#canonical-8bae6e2b386d2a401ab31abe669d8b768a7b45d7624d13e25ca80e373d921539): complete subsection reference.

<a id="canonical-c3e68414330ce99a5577b75bb17c912eef49284c574f4c61a55e838e78e2fa1a"></a>

<a id="canonical-39734ef0f4ba1b565d880dd619d8c9e05b40dc4f4ece0e2cc64fa6cc611634bd"></a>

## request_cookies_to_remove property — routes.simple_route.advanced_options / c00684f09de4 / 7

Type: `["list", "string"]`. Optional.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [request_headers_to_add](resources--http_loadbalancer--reference--group-024.md#canonical-fe40cbfb64ecdbc429fb055ad7e50015a785212961578e768a42bccfe4952677): complete subsection reference.

<a id="canonical-93bc3dd9ea01ba3691ed8430b842200321a5e827823a70ac39e592426c72748c"></a>

<a id="canonical-5069b726b224c52f732add6e5ff7302059157fbdece4b246d251c24703d62fe6"></a>

## request_headers_to_remove property — routes.simple_route.advanced_options / c00684f09de4 / 8

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_cookies_to_add](resources--http_loadbalancer--reference--group-024.md#canonical-3985b17e1a5096f3db4c6a657ac616194c986fe82e7d9f82e728fcd6ee04bbc8): complete subsection reference.

<a id="canonical-dd58735459d2030dfdca8db6c1f94dadb8d01dc97548eda58748fef31ee3f7bd"></a>

<a id="canonical-af65c67f9afeea4db5ab35e3e9c9249647f13a7b60e31cdb8aa63dcc69388c55"></a>

## response_cookies_to_remove property — routes.simple_route.advanced_options / c00684f09de4 / 9

Type: `["list", "string"]`. Optional.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-723ee2b749676d9fc55181ef05cd23ac76ed60b0b4cccd6183fcdfea8cf8f2ec): complete subsection reference.

<a id="canonical-272477f6231f7c4f3dd90003d3553480dc30baaa6b348531b8f16eda32516c37"></a>

<a id="canonical-ace23da0e54e60896df4b0c2081de10a4576c5666c9a783a416a485945aad6ec"></a>

## response_headers_to_remove property — routes.simple_route.advanced_options / c00684f09de4 / 10

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [retract_cluster](resources--http_loadbalancer--reference--group-025.md#canonical-4189f3380761bfac84e8c054272d103de5f6e45d674011352958d53bb526d795): complete subsection reference.

- [retry_policy](resources--http_loadbalancer--reference--group-025.md#canonical-8d9c9dbe977b7710eafc755f1a1302947bca687d1f839a82e94291c1054bc96c): complete subsection reference.

- [specific_hash_policy](resources--http_loadbalancer--reference--group-025.md#canonical-be4f556dc189ac1103e911e1b0a118b13cab9c72ef61280501141250bca2659e): complete subsection reference.

<a id="canonical-1e511d67ca7095c23dfad0b90ae85c6a52bc14b53f2dcac93d35626a55c1aafa"></a>

<a id="canonical-fb931e504afd987a37eead7b1624f2ebc5aaa42660bbb03c3c3b5a96aa0a1ca6"></a>

## timeout property — routes.simple_route.advanced_options / c00684f09de4 / 11

Type: `"number"`. Optional.

The timeout for the route including all retries, in milliseconds. Should be set to a high value or 0
(infinite timeout) for server-side streaming.

Upstream description:

The timeout for the route including all retries, in milliseconds. Should be set to a high value or 0
(infinite timeout) for server-side streaming.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [waf_exclusion_policy](resources--http_loadbalancer--reference--group-025.md#canonical-db429de4dc85878a0ef7a4a868529379de94854cc6ebe57d2417fa04bf1362e1): complete subsection reference.

- [web_socket_config](resources--http_loadbalancer--reference--group-025.md#canonical-a9734c2effd6190ca83be8d115daf0f43c20fae1b69870e2bce391b4dea1af48): complete subsection reference.

<a id="canonical-dd80b1b419833ffc35ddb2e54b977f5c5fa0ef2cc57fdbb0dccfc9ef9b67cf93"></a>

## Next pages — routes.simple_route.advanced_options / c00684f09de4 / 12

- [routes.simple_route.advanced_options.app_firewall](resources--http_loadbalancer--reference--group-024.md#canonical-810d8fde52c13db2d166c0c69c5fa62355d7ca5f7f6ec4643f1a034bac417855)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection](resources--http_loadbalancer--reference--group-024.md#canonical-538f35fdb26f20505920e66daa6a491f8209a103ae0f1667e1a76df9228cb599)
- [routes.simple_route.advanced_options.buffer_policy](resources--http_loadbalancer--reference--group-024.md#canonical-6505c45b2227fa5247d82aed0f887acdba1b2867295ab053ec5922dcd2ba29e3)
- [routes.simple_route.advanced_options.common_buffering](resources--http_loadbalancer--reference--group-024.md#canonical-c87fad1ad82b33b27f1cfb4f23ac91807f59d75f512bd75e22ef8809a83d278e)
- [routes.simple_route.advanced_options.common_hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-688a99d983f0f544d634ee5f3debc14e908e48ce76da3e4afdd5e8d17bd352f9)
- [routes.simple_route.advanced_options.cors_policy](resources--http_loadbalancer--reference--group-024.md#canonical-ea2851e52dea0f4cc37b0e3b2b7f2866c5a2353275e13df98aefca621d6da778)
- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-024.md#canonical-896b37a975892f4ac16ce875dd53f0b91cafde199fa15ef7205646b622351335)
- [routes.simple_route.advanced_options.default_retry_policy](resources--http_loadbalancer--reference--group-024.md#canonical-c70bb96c80e44e8d22a3fc83f8ad88cbf52f1d1c5ba26e2ab3bcdf242a30d2b5)
- [routes.simple_route.advanced_options.disable_mirroring](resources--http_loadbalancer--reference--group-024.md#canonical-6b9287834e37574a29fc4deef0b95e7fc82d04b3c0d7f8f5bb29ce9871df086e)
- [routes.simple_route.advanced_options.disable_prefix_rewrite](resources--http_loadbalancer--reference--group-024.md#canonical-d1a99f6d0246449840b524a894b6c6e69506dd9e99962be170a13ac38b86a61b)
- [routes.simple_route.advanced_options.disable_spdy](resources--http_loadbalancer--reference--group-024.md#canonical-7f5bd37ed4b125a95b4a6a6a78ffd31492c0efaac4498820209a25d8e9d9aa06)
- [routes.simple_route.advanced_options.disable_waf](resources--http_loadbalancer--reference--group-024.md#canonical-86ea4481408913913118621faa19d6328f73ba0b1d572e6e53a122ba1cf05e95)
- [routes.simple_route.advanced_options.disable_web_socket_config](resources--http_loadbalancer--reference--group-024.md#canonical-5b1d0aa3c6df13a9bfc188198d2f3ef6bc4bc899f798d9627066e6aeec7cd08c)
- [routes.simple_route.advanced_options.do_not_retract_cluster](resources--http_loadbalancer--reference--group-024.md#canonical-c7416844f81270e4944c8b9d6fc8b2ff78573ac4688507ccbaf3c48455cf1fff)
- [routes.simple_route.advanced_options.enable_spdy](resources--http_loadbalancer--reference--group-024.md#canonical-e8a531cc4f2e5226b7028c8c1699b8eaab5039187c5316bf181e6e208570698d)
- [routes.simple_route.advanced_options.endpoint_subsets](resources--http_loadbalancer--reference--group-024.md#canonical-9bb66d001ca2632fbea9d56bfc0d314b25683f8620bcb151ee2540a7326b56dd)
- [routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection](resources--http_loadbalancer--reference--group-024.md#canonical-d1c42952d17231babf05a5d044253f9523a24b246c3d2dda52595ad9a4611a80)
- [routes.simple_route.advanced_options.inherited_waf](resources--http_loadbalancer--reference--group-024.md#canonical-278571ba5269ca7afe639f5af6e0b7cd47eff71c0f57349adfbb6c210b620d3b)
- [routes.simple_route.advanced_options.inherited_waf_exclusion](resources--http_loadbalancer--reference--group-024.md#canonical-821811c0f4d76445f73d7e659d3d92d74205dd708146d942d0348fba9e1a632d)
- [routes.simple_route.advanced_options.mirror_policy](resources--http_loadbalancer--reference--group-024.md#canonical-47069f6847d19857cb08c8f8cc04bf77d1cdf5df9600e52489a86f8af3f1ebcb)
- [routes.simple_route.advanced_options.no_retry_policy](resources--http_loadbalancer--reference--group-024.md#canonical-a8fd9cd2de2bff8f95eb44dc2caaede6525e0322e23b32ff98f07a0718da8c15)
- [routes.simple_route.advanced_options.regex_rewrite](resources--http_loadbalancer--reference--group-024.md#canonical-381158a64be0f1d30cd5a2dd7b735d6bdc6d12ea248a26fc03ccab67f33b613e)
- [routes.simple_route.advanced_options.request_cookies_to_add](resources--http_loadbalancer--reference--group-024.md#canonical-8bae6e2b386d2a401ab31abe669d8b768a7b45d7624d13e25ca80e373d921539)
- [routes.simple_route.advanced_options.request_headers_to_add](resources--http_loadbalancer--reference--group-024.md#canonical-fe40cbfb64ecdbc429fb055ad7e50015a785212961578e768a42bccfe4952677)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-024.md#canonical-3985b17e1a5096f3db4c6a657ac616194c986fe82e7d9f82e728fcd6ee04bbc8)
- [routes.simple_route.advanced_options.response_headers_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-723ee2b749676d9fc55181ef05cd23ac76ed60b0b4cccd6183fcdfea8cf8f2ec)
- [routes.simple_route.advanced_options.retract_cluster](resources--http_loadbalancer--reference--group-025.md#canonical-4189f3380761bfac84e8c054272d103de5f6e45d674011352958d53bb526d795)
- [routes.simple_route.advanced_options.retry_policy](resources--http_loadbalancer--reference--group-025.md#canonical-8d9c9dbe977b7710eafc755f1a1302947bca687d1f839a82e94291c1054bc96c)
- [routes.simple_route.advanced_options.specific_hash_policy](resources--http_loadbalancer--reference--group-025.md#canonical-be4f556dc189ac1103e911e1b0a118b13cab9c72ef61280501141250bca2659e)
- [routes.simple_route.advanced_options.waf_exclusion_policy](resources--http_loadbalancer--reference--group-025.md#canonical-db429de4dc85878a0ef7a4a868529379de94854cc6ebe57d2417fa04bf1362e1)
- [routes.simple_route.advanced_options.web_socket_config](resources--http_loadbalancer--reference--group-025.md#canonical-a9734c2effd6190ca83be8d115daf0f43c20fae1b69870e2bce391b4dea1af48)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-810d8fde52c13db2d166c0c69c5fa62355d7ca5f7f6ec4643f1a034bac417855"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9133359137655849b44532e3e0d1789557f4838cbe1fda4a9b45b91c0fe98a4"></a>

## routes.simple_route.advanced_options.app_firewall — routes.simple_route.advanced_options.app_firewall / 03515a710cae / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.app_firewall

<a id="canonical-409a7ba290c57b7098447fcbfdbee7b935cd5772660982cbfdc41631c68baa89"></a>

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
app_firewall {
  # Configure direct properties listed below.
}
```

<a id="canonical-2fd05b602f0edf0b453616d925303e4f5d7c79a0f89867434c57a67de12d4d4d"></a>

## Direct properties — routes.simple_route.advanced_options.app_firewall / 03515a710cae / 3

<a id="canonical-157916066b5e8d680a5ea9caed4c904745c27cfde81c6f625ca8b074b89fd5e8"></a>

<a id="canonical-fe87f046e94ded9ed085b9a51511ef9700243cfb3b4728553c87f7e919fef1d8"></a>

## name property — routes.simple_route.advanced_options.app_firewall / 03515a710cae / 4

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

<a id="canonical-998a530126664245beeaffd7568f482be11ee15a0e77f65408d17e1e35dd31d7"></a>

<a id="canonical-cc5a4e0dc1b6d4b4860159a0d4c930c4b42469dfa5093eabb7b5b300ec9da6fe"></a>

## namespace property — routes.simple_route.advanced_options.app_firewall / 03515a710cae / 5

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

<a id="canonical-64f1a4386229db19cfa0e4708bbd8eeb47fe9f47d87c980f23bf7a8a3b4d765c"></a>

<a id="canonical-42f8e534ae068162d2537407f0f302efcf6415a40708050310a4c4407ef9d63a"></a>

## tenant property — routes.simple_route.advanced_options.app_firewall / 03515a710cae / 6

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

<a id="canonical-0461a134ef7a518c26dbf28102640ef54ed2106504427544384f3ca981bf35ae"></a>

## Next pages — routes.simple_route.advanced_options.app_firewall / 03515a710cae / 7

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-538f35fdb26f20505920e66daa6a491f8209a103ae0f1667e1a76df9228cb599"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0565f9c383bf70af78539c06b79fc530dc30262e299bf918a9805d32e940dddd"></a>

## routes.simple_route.advanced_options.bot_defense_javascript_injection — routes.simple_route.advanced_options.bot_defense_javascript_injection / 5223a41a3c13 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.bot_defense_javascript_injection

<a id="canonical-d6abd29c7e6c8232b0f607f3606734f8c49b8fe871ed7c5ebf3a67fb668ef65a"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Javascript Injection Configuration for inline bot defense deployments.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("javascript_tags")}
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
bot_defense_javascript_injection {
  # Configure direct properties listed below.
}
```

<a id="canonical-a41a004d9e2c42d8557b01483339222d05a1a49801164d8110b0830c9d62bbe1"></a>

## Direct properties — routes.simple_route.advanced_options.bot_defense_javascript_injection / 5223a41a3c13 / 3

<a id="canonical-8466a1d81ce4fe8643aac830dd280a8ce218237f53afe3094dd543710da705c9"></a>

<a id="canonical-a744881e5e0c534602b113e385d885a5f827a4b8bfaf30580b45bc9ed5d085a8"></a>

## javascript_location property — routes.simple_route.advanced_options.bot_defense_javascript_injection / 5223a41a3c13 / 4

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [javascript_tags](resources--http_loadbalancer--reference--group-024.md#canonical-2cec61af7ddec873874add25b02a011d0740e3477bdc9d44b127c345c34cf499): complete subsection reference.

<a id="canonical-9d0ddf84be4f414ee0c4bf2a0d406a7555f0540b31e64ba73acdd86f084110f1"></a>

## Next pages — routes.simple_route.advanced_options.bot_defense_javascript_injection / 5223a41a3c13 / 5

- [routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags](resources--http_loadbalancer--reference--group-024.md#canonical-2cec61af7ddec873874add25b02a011d0740e3477bdc9d44b127c345c34cf499)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2cec61af7ddec873874add25b02a011d0740e3477bdc9d44b127c345c34cf499"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6abd194359a661f38926dd3a8ca7ce0a9c1956f0fb88b870bfa0f81d0585604e"></a>

## routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags — routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript / cce73dce47a1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection](resources--http_loadbalancer--reference--group-024.md#canonical-538f35fdb26f20505920e66daa6a491f8209a103ae0f1667e1a76df9228cb599)
- routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags

<a id="canonical-e53e0be6b70b1c56b3334b02649966c0b305901567e78a658c78ad09b2075fe7"></a>

Type: `"object"`. list nested block, Optional.

Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot
Javascript should be added first.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("javascript_url")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
javascript_tags {
  # Configure direct properties listed below.
}
```

<a id="canonical-8ded03b9dc1319261da4e7c3a6510404b356a0051ffe3d48921723cf066a6f59"></a>

## Direct properties — routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript / cce73dce47a1 / 3

<a id="canonical-73d52d7a9e79157c441a810185bdc60430d072afec288149de8f8557d66f54f0"></a>

<a id="canonical-5ce5aab6a61d4a4dd7db5fb4e90bfaa3ca9769b9598e429ecbe3f71c9ad8deba"></a>

## javascript_url property — routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript / cce73dce47a1 / 4

Type: `"string"`. Optional.

Please enter the full URL (include domain and path), or relative path.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 2048,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_bytes": "2048",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "2048",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [tag_attributes](resources--http_loadbalancer--reference--group-024.md#canonical-9746e6a4162a9744b2f430d74f6b9c5990bdcf4743fc5e70a83cb1f359572754): complete subsection reference.

<a id="canonical-01b63f92315228a40e6b57831e60494f1ec828bbbec7ea697ade4d222a56b130"></a>

## Next pages — routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript / cce73dce47a1 / 5

- [routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags.tag_attributes](resources--http_loadbalancer--reference--group-024.md#canonical-9746e6a4162a9744b2f430d74f6b9c5990bdcf4743fc5e70a83cb1f359572754)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection](resources--http_loadbalancer--reference--group-024.md#canonical-538f35fdb26f20505920e66daa6a491f8209a103ae0f1667e1a76df9228cb599)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9746e6a4162a9744b2f430d74f6b9c5990bdcf4743fc5e70a83cb1f359572754"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0aa00b33b8a9b09a335a88b1635ecfab3b1052f6c7de3dad0741e92a2cabef59"></a>

## routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags.tag_attributes — routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript / e0000e079825 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection](resources--http_loadbalancer--reference--group-024.md#canonical-538f35fdb26f20505920e66daa6a491f8209a103ae0f1667e1a76df9228cb599)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags](resources--http_loadbalancer--reference--group-024.md#canonical-2cec61af7ddec873874add25b02a011d0740e3477bdc9d44b127c345c34cf499)
- routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags.tag_attributes

<a id="canonical-d9c172b3442fed8bbc883c9293f57d629276bc4910b9245d1aa44089e22121e3"></a>

Type: `"object"`. list nested block, Optional.

Add the tag attributes you want to include in your Javascript tag.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tag_attributes {
  # Configure direct properties listed below.
}
```

<a id="canonical-d3fa0d6a047e94787fdcd165c6af13723d60c364650d39e0e8436d8745fff1ce"></a>

## Direct properties — routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript / e0000e079825 / 3

<a id="canonical-2e43d281befe37290f8ed8dba7ed016f3946ff789f9ce6c6d9a5b923c41ccf3e"></a>

<a id="canonical-90ff7a84c61686927240e501cb845976ddbfb5763bbdda05462f4c1810b53dc7"></a>

## javascript_tag property — routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript / e0000e079825 / 4

Type: `"string"`. Optional.

\[Enum:
JS\_ATTR\_ID|JS\_ATTR\_CID|JS\_ATTR\_CN|JS\_ATTR\_API\_DOMAIN|JS\_ATTR\_API\_URL|JS\_ATTR\_API\_PATH|JS\_ATTR\_ASYNC|JS\_ATTR\_DEFER\]
Select from one of the predefined tag attributes. Possible values are \`JS\_ATTR\_ID\`,
\`JS\_ATTR\_CID\`, \`JS\_ATTR\_CN\`, \`JS\_ATTR\_API\_DOMAIN\`, \`JS\_ATTR\_API\_URL\`,
\`JS\_ATTR\_API\_PATH\`, \`JS\_ATTR\_ASYNC\`, \`JS\_ATTR\_DEFER\`. Defaults to \`JS\_ATTR\_ID\`.

Upstream description:

Select from one of the predefined tag attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("JS_ATTR_ID",
    "JS_ATTR_CID",
    "JS_ATTR_CN",
    "JS_ATTR_API_DOMAIN",
    "JS_ATTR_API_URL",
    "JS_ATTR_API_PATH",
    "JS_ATTR_ASYNC",
    "JS_ATTR_DEFER"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "JS_ATTR_ID",
  "enum": [
    "JS_ATTR_ID",
    "JS_ATTR_CID",
    "JS_ATTR_CN",
    "JS_ATTR_API_DOMAIN",
    "JS_ATTR_API_URL",
    "JS_ATTR_API_PATH",
    "JS_ATTR_ASYNC",
    "JS_ATTR_DEFER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d18898984a5cd8fbca228dd2cc92321559f5df62eab247ffd05746050ad1f1b4"></a>

<a id="canonical-2e623fa5217902cf06ded77765779a07406e040db6a8650e0003e2c96b3d25c4"></a>

## tag_value property — routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript / e0000e079825 / 5

Type: `"string"`. Optional.

Value. Add the tag attribute value.

Upstream description:

Add the tag attribute value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_bytes": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  }
}
```

<a id="canonical-c97852004b5e8bcb2a915881bd632197f98c483a4854217ea369d58b32c5138f"></a>

## Next pages — routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript / e0000e079825 / 6

- [routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags](resources--http_loadbalancer--reference--group-024.md#canonical-2cec61af7ddec873874add25b02a011d0740e3477bdc9d44b127c345c34cf499)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6505c45b2227fa5247d82aed0f887acdba1b2867295ab053ec5922dcd2ba29e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f1c35a99646527ffef9626d6df0533ff490b4f0217f3af8b2c9395f29a580b3"></a>

## routes.simple_route.advanced_options.buffer_policy — routes.simple_route.advanced_options.buffer_policy / 411d31985be7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.buffer_policy

<a id="canonical-100fce355c1f77bc7e95f464168a9ea2962c46edde1d59ca6807ed312b759f3d"></a>

Type: `"object"`. single nested block, Optional.

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Upstream description:

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer
configuration takes precedence.

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
buffer_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-441ae476268f7c720320c550b2f21bad16f0e30434cd8dae84bf2d3e16ca66fe"></a>

## Direct properties — routes.simple_route.advanced_options.buffer_policy / 411d31985be7 / 3

<a id="canonical-55abc8520488551a2eb0b408e37e10be33ad36a20eb4b6c9a9b5c2e7ae62e289"></a>

<a id="canonical-fe6f8546ca8e2e12c99c103c5dd814ba27a92b63ac4c9311b7d50384ce11ac28"></a>

## disabled property — routes.simple_route.advanced_options.buffer_policy / 411d31985be7 / 4

Type: `"bool"`. Optional.

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

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

<a id="canonical-e4aca8010eb3325be16b0e0654093bd170eed70bd24123f03dbb34c11eb93bf0"></a>

<a id="canonical-11d5835141b62ed3012268ba7760c5bcb2ca84b44c57827048775c57f271dae5"></a>

## max_request_bytes property — routes.simple_route.advanced_options.buffer_policy / 411d31985be7 / 5

Type: `"number"`. Optional.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Upstream description:

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(10485760),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
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
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

<a id="canonical-cea69aa2e3db17a4df963091d4c429c5e55f31d031f61d9d2c55eb710ee538d3"></a>

## Next pages — routes.simple_route.advanced_options.buffer_policy / 411d31985be7 / 6

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c87fad1ad82b33b27f1cfb4f23ac91807f59d75f512bd75e22ef8809a83d278e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-117268b26bbe9014342d1ed71986092ce2bd56db814068e3e2cb61f8b7d4b6dd"></a>

## routes.simple_route.advanced_options.common_buffering — routes.simple_route.advanced_options.common_buffering / 3eb34d845f69 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.common_buffering

<a id="canonical-e1f4f1e75e7e00b5a7779305a7a839bb00b5ef123716e421bb53bc8546d42208"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for common buffering.

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
common_buffering = {}
```

<a id="canonical-6a9d88c81b5bc656e7c08bf40786395a0a37d7ce287159d6f1cfdb0682fcaa22"></a>

## Direct properties — routes.simple_route.advanced_options.common_buffering / 3eb34d845f69 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8c5c34513256ce041a8146e803f6248f591027a79f47c87237891e8f4d74d211"></a>

## Next pages — routes.simple_route.advanced_options.common_buffering / 3eb34d845f69 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-688a99d983f0f544d634ee5f3debc14e908e48ce76da3e4afdd5e8d17bd352f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59cd21f599624b898b07511174372b01f136bd58902c946b650baf2747916315"></a>

## routes.simple_route.advanced_options.common_hash_policy — routes.simple_route.advanced_options.common_hash_policy / 54b33f644e79 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.common_hash_policy

<a id="canonical-0ba9748c9f3d71b23feb7e53e9ff98106e4179e145d29d7e7c0a01c2ce095dad"></a>

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
common_hash_policy = {}
```

<a id="canonical-0e9d83a4d81d70d63ed47c53f19ba7af2d310ca12e92ffc3c1a741f504787277"></a>

## Direct properties — routes.simple_route.advanced_options.common_hash_policy / 54b33f644e79 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a016fc96e7ec65eb25f5f6e80c4be7d4af1f5075ce8755aeb557b2df9b667a49"></a>

## Next pages — routes.simple_route.advanced_options.common_hash_policy / 54b33f644e79 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ea2851e52dea0f4cc37b0e3b2b7f2866c5a2353275e13df98aefca621d6da778"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-937c13f3b539ed5da0ba2f1addc2cbfa1fbcb22e09a51455494e4aba99426abb"></a>

## routes.simple_route.advanced_options.cors_policy — routes.simple_route.advanced_options.cors_policy / ae11b06cd007 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.cors_policy

<a id="canonical-8bf4b04b888de2d1bdd2820e1b2352a194f9b433eb3977f1cc596b20c8ca6f80"></a>

Type: `"object"`. single nested block, Optional.

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence. An example of an Cross origin HTTP request GET
/resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS
X 10.5..

Upstream description:

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence.

An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other
User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130
Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8
Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset:
ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive Referrer:
http&#58;//foo.example/examples/access-control/simplexsinvocation.html Origin:
http&#58;//foo.example

HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61
Access-Control-Allow-Origin: \* Keep-Alive: timeout=2, max=100 Connection: Keep-Alive
Transfer-Encoding: chunked Content-Type: application/XML

An example for cross origin HTTP OPTIONS request with Access-Control-Request-\* header

OPTIONS /resources/POST-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel
MAC OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept:
text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8 Accept-Language: en-us,en;q=0.5
Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive
Origin: http&#58;//foo.example Access-Control-Request-Method: POST Access-Control-Request-Headers:
X-PINGOTHER, Content-Type

HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix)
Access-Control-Allow-Origin: http&#58;//foo.example Access-Control-Allow-Methods: POST, GET, OPTIONS
Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary:
Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive.

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
cors_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-c3d1a33345d4c3b80eaa7c8696bd0d7ebc9dfcc6de83a4375aa7a5a0f5197420"></a>

## Direct properties — routes.simple_route.advanced_options.cors_policy / ae11b06cd007 / 3

<a id="canonical-2a05ff1ce2bfc700f362721861dad8f7fabc4c1c1636faa294d6c5125cff5b2b"></a>

<a id="canonical-2c15f40164152a5f1f220063cd9aaf31364b1f0633a4bcd24bb91c89e11680d0"></a>

## allow_credentials property — routes.simple_route.advanced_options.cors_policy / ae11b06cd007 / 4

Type: `"bool"`. Optional.

Specifies whether the resource allows credentials.

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

<a id="canonical-fc4988df1ccaab22a4c33151dc91d97a6f16c4c36691f989f1cc274e48d021d5"></a>

<a id="canonical-23429d2a2dac926bdc19aad4bb4f54eecc74653247b7be0036ac98ddda20699e"></a>

## allow_headers property — routes.simple_route.advanced_options.cors_policy / ae11b06cd007 / 5

Type: `"string"`. Optional.

Specifies the content for the access-control-allow-headers header.

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

<a id="canonical-d9df92f323459a60ee24dd6b6d8b745820d0bbd07981990d9434e45210be16a7"></a>

<a id="canonical-d6ebbf98e1cd7f3830f125ff8d12cba64d4a42a25a2e7438f6030c3c39f5f583"></a>

## allow_methods property — routes.simple_route.advanced_options.cors_policy / ae11b06cd007 / 6

Type: `"string"`. Optional.

Specifies the content for the access-control-allow-methods header.

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
    "ves.io.schema.rules.string.http_valid_methods": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  }
}
```

<a id="canonical-20e97265a15e9724525fd14bc91d2576189cbfee0a151819df494cc4c2eddbef"></a>

<a id="canonical-e8417709630480d9d809dacc48ad59412deec687f2cb08f50c1af61daa9e0280"></a>

## allow_origin property — routes.simple_route.advanced_options.cors_policy / ae11b06cd007 / 7

Type: `["list", "string"]`. Optional.

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

Upstream description:

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2e5362957d2719ca3d778488059f5ab560621b2fe88219e055da2018f060c477"></a>

<a id="canonical-2dec2126332ab954b7d50d760e9fde8e687a2d0a6f8471626249084f95f809ef"></a>

## allow_origin_regex property — routes.simple_route.advanced_options.cors_policy / ae11b06cd007 / 8

Type: `["list", "string"]`. Optional.

Specifies regex patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

Upstream description:

Specifies regex patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-b972d67a1bd4cfe62405c9f52452a48655a82532868e8f83587b10f474240080"></a>

<a id="canonical-45c7622663559d3ffd3c3aadc1affe0b3562d7e0053be413d3675c7715940d72"></a>

## disabled property — routes.simple_route.advanced_options.cors_policy / ae11b06cd007 / 9

Type: `"bool"`. Optional.

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

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

<a id="canonical-6d739ee10abbfc3c30036b292ce8b7fe508fe1dc5436cb73a2dc67d0ef902c11"></a>

<a id="canonical-7fe22499bd8b5521dcf898c80ae90f7cca1b87c77b82cd6a89a6b7973663e972"></a>

## expose_headers property — routes.simple_route.advanced_options.cors_policy / ae11b06cd007 / 10

Type: `"string"`. Optional.

Specifies the content for the access-control-expose-headers header.

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

<a id="canonical-759377a7e479f8ed4e564df6d2d01cd5df65c346678747ee132f97a4c4dd5317"></a>

<a id="canonical-98c95dbf5244a54c22bf56a9d2e7e142192552a433a3dc505cccab258c0a9b4a"></a>

## maximum_age property — routes.simple_route.advanced_options.cors_policy / ae11b06cd007 / 11

Type: `"number"`. Optional.

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours).

Upstream description:

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(-1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": -1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  }
}
```

<a id="canonical-394d1cd37454978b11a8e065154b1ebc90db3caf2d1e1a2f49e0df78fa26037b"></a>

## Next pages — routes.simple_route.advanced_options.cors_policy / ae11b06cd007 / 12

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-896b37a975892f4ac16ce875dd53f0b91cafde199fa15ef7205646b622351335"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2f9f37cf49084eea33cbb7dfd404191e72d7a5224fa4826d8ead65ec180552d"></a>

## routes.simple_route.advanced_options.csrf_policy — routes.simple_route.advanced_options.csrf_policy / 95c1c9d48470 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.csrf_policy

<a id="canonical-e47c49d5767263c3ca5d747d44488a0f44a2f52eeb78874ef47bc08f97cd3966"></a>

Type: `"object"`. single nested block, Optional.

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host. 1. The origin that caused the user
agent..

Upstream description:

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host.

&#8203;1. The origin that caused the user agent to issue the request (source origin). &#8203;2. The
origin that the request is going to (target origin). When the policy evaluating a request, it
ensures both pieces of information are present and compare their values. If the source origin is
missing or origins do not match the request is rejected. The exception to this being if the
source-origin has been added to they policy as valid. Because CSRF attacks specifically target
state-changing requests, the policy only acts on the HTTP requests that have state-changing method
(PUT,POST, etc.).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_load_balancer_domains",
    "custom_domain_list"),
  validators.ConflictingObjectAttributes("all_load_balancer_domains",
    "disabled"),
  validators.ConflictingObjectAttributes("custom_domain_list",
    "disabled")}
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
  "x-ves-oneof-field-allowed_domains": "[\"all_load_balancer_domains\",\"custom_domain_list\",\"disabled\"]"
}
```

Terraform syntax:

```terraform
csrf_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-bcd5306e833ba815a9a27e6bce379d815af972efdd137feb1e858a17e30f697e"></a>

## Direct properties — routes.simple_route.advanced_options.csrf_policy / 95c1c9d48470 / 3

- [all_load_balancer_domains](resources--http_loadbalancer--reference--group-024.md#canonical-d297370235cad6d6c97c9678b058047a0a9da8e8f98c9de94c6bb4f26571f4e9): complete subsection reference.

- [custom_domain_list](resources--http_loadbalancer--reference--group-024.md#canonical-ae2c46207c5a46d89ac5bf5ee6c9f61a02c50cc0980f0c2ba5c0f1ce34674b9a): complete subsection reference.

- [disabled](resources--http_loadbalancer--reference--group-024.md#canonical-a727f52f392caa0ce97b53bf13010636fa80da75b12a0a2372d3ac6961439f8c): complete subsection reference.

<a id="canonical-c029098ea69bceb1fdad7a6c53467ac7efa00e6dce43fe84ead1ee4268a76688"></a>

## Next pages — routes.simple_route.advanced_options.csrf_policy / 95c1c9d48470 / 4

- [routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains](resources--http_loadbalancer--reference--group-024.md#canonical-d297370235cad6d6c97c9678b058047a0a9da8e8f98c9de94c6bb4f26571f4e9)
- [routes.simple_route.advanced_options.csrf_policy.custom_domain_list](resources--http_loadbalancer--reference--group-024.md#canonical-ae2c46207c5a46d89ac5bf5ee6c9f61a02c50cc0980f0c2ba5c0f1ce34674b9a)
- [routes.simple_route.advanced_options.csrf_policy.disabled](resources--http_loadbalancer--reference--group-024.md#canonical-a727f52f392caa0ce97b53bf13010636fa80da75b12a0a2372d3ac6961439f8c)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d297370235cad6d6c97c9678b058047a0a9da8e8f98c9de94c6bb4f26571f4e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14da62431124a71a63ef01b38979bad63642318f670c356c536c705fc266c6da"></a>

## routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains — routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains / 90cb8af00f59 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-024.md#canonical-896b37a975892f4ac16ce875dd53f0b91cafde199fa15ef7205646b622351335)
- routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains

<a id="canonical-fc1026fd0aafb082e63aa797df8c6d39ed5f4c7283d9aa841a1a22090c5ee57f"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all load balancer domains.

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
all_load_balancer_domains = {}
```

<a id="canonical-6e5fefd5c6e5995425a79d15a0cf8e4ddf462e20d8cd397fe6618b99264d749b"></a>

## Direct properties — routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains / 90cb8af00f59 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a900b4e641e1cc9c0c5ff8d10ef374c8ec3e6e6a25ee8985fd6b8695935aea7f"></a>

## Next pages — routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains / 90cb8af00f59 / 4

- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-024.md#canonical-896b37a975892f4ac16ce875dd53f0b91cafde199fa15ef7205646b622351335)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ae2c46207c5a46d89ac5bf5ee6c9f61a02c50cc0980f0c2ba5c0f1ce34674b9a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cbb1ee4f117b27931eeb0c01721415734bd43b419b7246aed650d6f06be6ee93"></a>

## routes.simple_route.advanced_options.csrf_policy.custom_domain_list — routes.simple_route.advanced_options.csrf_policy.custom_domain_list / ccefcbaae051 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-024.md#canonical-896b37a975892f4ac16ce875dd53f0b91cafde199fa15ef7205646b622351335)
- routes.simple_route.advanced_options.csrf_policy.custom_domain_list

<a id="canonical-9564c6e213d89211c68e37b015769dc37bfcf2fc26d542d78563fc3fc12bb3d4"></a>

Type: `"object"`. single nested block, Optional.

List of domain names used for Host header matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains")}
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
custom_domain_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-4ab1be95527ca66e5c2034cb0790cf05009fed6e73cf30c998519c20fae4486b"></a>

## Direct properties — routes.simple_route.advanced_options.csrf_policy.custom_domain_list / ccefcbaae051 / 3

<a id="canonical-b27fb26d4b6c6b06e83fd27a8f1957d9d499a1f878146c83f34687cbe1a4f3d4"></a>

<a id="canonical-26a4b4b0dd4337f6ecd194ce8b375378f0b82acdc650f22d4674b9e18da9f30e"></a>

## domains property — routes.simple_route.advanced_options.csrf_policy.custom_domain_list / ccefcbaae051 / 4

Type: `["list", "string"]`. Optional.

List of domain names that will be matched to loadbalancer. These domains are not used for SNI match.
Wildcard names are supported in the suffix or prefix form.

Upstream description:

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-bc0a3b07463c9aa4d81eeccffa9025c7a48444f3a482630195b880e8b0304183"></a>

## Next pages — routes.simple_route.advanced_options.csrf_policy.custom_domain_list / ccefcbaae051 / 5

- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-024.md#canonical-896b37a975892f4ac16ce875dd53f0b91cafde199fa15ef7205646b622351335)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a727f52f392caa0ce97b53bf13010636fa80da75b12a0a2372d3ac6961439f8c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aff3bc4a3cabf070ca1e2cb9bbaa3912919244d83ba74c2737a534292ce2aeab"></a>

## routes.simple_route.advanced_options.csrf_policy.disabled — routes.simple_route.advanced_options.csrf_policy.disabled / 2c2f4f3e9a99 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-024.md#canonical-896b37a975892f4ac16ce875dd53f0b91cafde199fa15ef7205646b622351335)
- routes.simple_route.advanced_options.csrf_policy.disabled

<a id="canonical-a07af827fbe9c76f968544fa3d932d2339bce632f5b5f904159ae869700eef25"></a>

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
disabled = {}
```

<a id="canonical-f957ad889316c6bcc01c65668828455f75c17273773bcabeb3e2025648f1608b"></a>

## Direct properties — routes.simple_route.advanced_options.csrf_policy.disabled / 2c2f4f3e9a99 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-18847877a0eb1437176d8825255124e2f76816d928c64c5d68e26d904eeba60e"></a>

## Next pages — routes.simple_route.advanced_options.csrf_policy.disabled / 2c2f4f3e9a99 / 4

- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-024.md#canonical-896b37a975892f4ac16ce875dd53f0b91cafde199fa15ef7205646b622351335)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c70bb96c80e44e8d22a3fc83f8ad88cbf52f1d1c5ba26e2ab3bcdf242a30d2b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2a32ba4e34139d455afbcf2bce882109ce17449737f58e3414ba8db1756c87f"></a>

## routes.simple_route.advanced_options.default_retry_policy — routes.simple_route.advanced_options.default_retry_policy / bc8d8899e1b2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.default_retry_policy

<a id="canonical-d31686b2f9d7b3134e473971d315444d30294c2e136c91c2e179990bb253c578"></a>

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
default_retry_policy = {}
```

<a id="canonical-4692d0d0f85f843b4f9ee5d941d94e7d972a3835216a089c0a7870168425a045"></a>

## Direct properties — routes.simple_route.advanced_options.default_retry_policy / bc8d8899e1b2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-17a1e00c520cfb0292c69d89c06889839903bb78fad90efc4198c1303f5765df"></a>

## Next pages — routes.simple_route.advanced_options.default_retry_policy / bc8d8899e1b2 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6b9287834e37574a29fc4deef0b95e7fc82d04b3c0d7f8f5bb29ce9871df086e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02f2ae815426a925ec6f3a380613463971b505d6bac990c50d83197cf8f4b0b6"></a>

## routes.simple_route.advanced_options.disable_mirroring — routes.simple_route.advanced_options.disable_mirroring / 10ca8de3df8a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.disable_mirroring

<a id="canonical-4934d7948f24c52c2d55fe6add4fdd49fe72823b2a87b82d36113fe8209955a4"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable mirroring.

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
disable_mirroring = {}
```

<a id="canonical-84dafb22a1ff8ec371c5ede03eb72f1371be50ccc7e41a8491262bec8a0f9bdb"></a>

## Direct properties — routes.simple_route.advanced_options.disable_mirroring / 10ca8de3df8a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0d8b5b670abc746071c209361c86c4a5a41ed6af83634c36ecab59adee3166c4"></a>

## Next pages — routes.simple_route.advanced_options.disable_mirroring / 10ca8de3df8a / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d1a99f6d0246449840b524a894b6c6e69506dd9e99962be170a13ac38b86a61b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d095c81e16b9adeb811f090fd792a9a6ccc7bd9344fe803bf040533692f4339"></a>

## routes.simple_route.advanced_options.disable_prefix_rewrite — routes.simple_route.advanced_options.disable_prefix_rewrite / 6e1703c66847 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.disable_prefix_rewrite

<a id="canonical-6aa98074c89dfde80903d744d77e253c6987969ebd7f022aa8a6b8d2a9beb561"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable prefix rewrite.

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
disable_prefix_rewrite = {}
```

<a id="canonical-ca7e21ba512c1c5f9c7b820fcf27f43a01c2157b6c93cbdd0b06a72a4be3b665"></a>

## Direct properties — routes.simple_route.advanced_options.disable_prefix_rewrite / 6e1703c66847 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-28110b75e175f77fb6e3ec86eaa2e60b2dbcd962eb1b9e09baf6e8931b860639"></a>

## Next pages — routes.simple_route.advanced_options.disable_prefix_rewrite / 6e1703c66847 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7f5bd37ed4b125a95b4a6a6a78ffd31492c0efaac4498820209a25d8e9d9aa06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-601e7ecb7aaeb78db8159aa0a27e02036e660c959200e09b7dc54924ac188803"></a>

## routes.simple_route.advanced_options.disable_spdy — routes.simple_route.advanced_options.disable_spdy / ae283f1a0210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.disable_spdy

<a id="canonical-313d5570607d9a882f1970e404d8c9e43cd2e5c3137b906152e4785a4a7550f5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable spdy.

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
disable_spdy = {}
```

<a id="canonical-7f55f9501931d2c40749fe52bbd0d4b06c6aeae2d0e14cb34bad9e5ae661b91d"></a>

## Direct properties — routes.simple_route.advanced_options.disable_spdy / ae283f1a0210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-78b9d57de886ac62edbdd7c1a707833f317f9cb371a1a82384a5a21e266a3a0f"></a>

## Next pages — routes.simple_route.advanced_options.disable_spdy / ae283f1a0210 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-86ea4481408913913118621faa19d6328f73ba0b1d572e6e53a122ba1cf05e95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21c339e3e831cd9fc83b9f965fcd5c3a7d761f68d8dda38fc0e7f3b8db78703b"></a>

## routes.simple_route.advanced_options.disable_waf — routes.simple_route.advanced_options.disable_waf / 48843779a6ba / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.disable_waf

<a id="canonical-f31c5406202f71b2022c3c4ed52b40ac16a19750674c9ebaa4e9ae50fa4c329e"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable waf.

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
disable_waf = {}
```

<a id="canonical-842b39655ea521f715340a5a4d00722be7f0933db9d29d248e35c6d362f792c1"></a>

## Direct properties — routes.simple_route.advanced_options.disable_waf / 48843779a6ba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-74a8194573ed049c346bc8ac8ca6da8d3e415b0fc2df834f9972ebf435e5fd64"></a>

## Next pages — routes.simple_route.advanced_options.disable_waf / 48843779a6ba / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5b1d0aa3c6df13a9bfc188198d2f3ef6bc4bc899f798d9627066e6aeec7cd08c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-deb3f2953383d7282e18e9537c9fd05dc74c849d31523b1c6a13d951d8d30ba7"></a>

## routes.simple_route.advanced_options.disable_web_socket_config — routes.simple_route.advanced_options.disable_web_socket_config / 7777963231c7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.disable_web_socket_config

<a id="canonical-072228d4f6b02c03bcb4212d552098004ed9063921d342eeeb01bd1227be030f"></a>

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
disable_web_socket_config = {}
```

<a id="canonical-913c6db8dbbd2e70079ee847b9eede5f918dee8d71f45ad21ac3dfc6c14585ba"></a>

## Direct properties — routes.simple_route.advanced_options.disable_web_socket_config / 7777963231c7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-920f8a22efd163f7d78b98f6b3673a6da812e2c02d3a162258fad21a25e02aec"></a>

## Next pages — routes.simple_route.advanced_options.disable_web_socket_config / 7777963231c7 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c7416844f81270e4944c8b9d6fc8b2ff78573ac4688507ccbaf3c48455cf1fff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-706e0aef126b6476c671c8ddca6d30454675c51f2f9ab9ab237539a2b9697e73"></a>

## routes.simple_route.advanced_options.do_not_retract_cluster — routes.simple_route.advanced_options.do_not_retract_cluster / fa203d7e7088 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.do_not_retract_cluster

<a id="canonical-ce1240401ec74597eddaaed63842da54772b15adf897b6d9e2fe327d4594d093"></a>

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
do_not_retract_cluster = {}
```

<a id="canonical-a32430b8f8ebcf58ebf6e119017aab2d76f9a3159b58b40426e864771383ac01"></a>

## Direct properties — routes.simple_route.advanced_options.do_not_retract_cluster / fa203d7e7088 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-db676749a7a0ef1c4b65a4f8f807809647543a1d39fa2719deb2ac339fc8e4f2"></a>

## Next pages — routes.simple_route.advanced_options.do_not_retract_cluster / fa203d7e7088 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e8a531cc4f2e5226b7028c8c1699b8eaab5039187c5316bf181e6e208570698d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-975d8c24aa33dc8a8e6e8c895a9b399794284eb95a976d3ed0706adeaeb72f51"></a>

## routes.simple_route.advanced_options.enable_spdy — routes.simple_route.advanced_options.enable_spdy / 7677a37e0d46 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.enable_spdy

<a id="canonical-4644afa2b4db1c96715354eb36966adb38f55ceeec72764f08e365898a5e0c64"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable spdy.

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
enable_spdy = {}
```

<a id="canonical-c06d9468d329a81bd8a35daf4b6258f5df3d0f57c9dd35902363ee675410a5fe"></a>

## Direct properties — routes.simple_route.advanced_options.enable_spdy / 7677a37e0d46 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fb471cb4cd567a50485a7b541d6a8c70f065ceec6e314490364ad957fc46aa8f"></a>

## Next pages — routes.simple_route.advanced_options.enable_spdy / 7677a37e0d46 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9bb66d001ca2632fbea9d56bfc0d314b25683f8620bcb151ee2540a7326b56dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-161b836fbf84700b44cb1740012c11fa5d79aec8202b676a55d8260948a9ec7f"></a>

## routes.simple_route.advanced_options.endpoint_subsets — routes.simple_route.advanced_options.endpoint_subsets / 8f455dd1ca4b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.endpoint_subsets

<a id="canonical-6967aea729b91443dd98deb800eeed9fe114e3bbe0401c9f57b5a17f7e7d193f"></a>

Type: `"object"`. single nested block, Optional.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer For origin servers which are discovered in K8s or Consul..

Upstream description:

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

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
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

Terraform syntax:

```terraform
endpoint_subsets {}
```

<a id="canonical-3b6ab67deb5727b82634a42c1e625889b4e522da8f334e7712720d72e63010ab"></a>

## Direct properties — routes.simple_route.advanced_options.endpoint_subsets / 8f455dd1ca4b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1c0d3471ddf7ebb860a2c6808ed8bb8a6896e468cc91e098fa2c2cfdda54fe8b"></a>

## Next pages — routes.simple_route.advanced_options.endpoint_subsets / 8f455dd1ca4b / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d1c42952d17231babf05a5d044253f9523a24b246c3d2dda52595ad9a4611a80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4cc8234bd58d5eb4aa89b6456b2ef7dce205eaa2d1a07400e1ec3bb44841c4ab"></a>

## routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection — routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection / 1ddd86efc459 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection

<a id="canonical-5a875546fef709df3d06b993a443085a0994862105d16c9e0d05aca4efe85d3e"></a>

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
inherited_bot_defense_javascript_injection = {}
```

<a id="canonical-f9408f87cf03a7556c64434f766f1863f2c6824ac34b516d57b86bca26dc8d16"></a>

## Direct properties — routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection / 1ddd86efc459 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-38f067d686390a318137c49a66c9aa472e0e9f7d81b4ec8ff5c947c7a5beca00"></a>

## Next pages — routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection / 1ddd86efc459 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-278571ba5269ca7afe639f5af6e0b7cd47eff71c0f57349adfbb6c210b620d3b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2215192e1979304f20502f1f02eb684e86b2dd4d929b1f60bbbdb64f1cda2a8"></a>

## routes.simple_route.advanced_options.inherited_waf — routes.simple_route.advanced_options.inherited_waf / 65175dd9c168 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.inherited_waf

<a id="canonical-678b76e4de7ba7fd598dc3bf41b62830a5388770356c0edf36ca3abd31be5dfa"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherited waf.

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
inherited_waf = {}
```

<a id="canonical-b9c692004d5dee295888f52af9b8ad75ac64020ac63abcfa4cb3a5d76f32496c"></a>

## Direct properties — routes.simple_route.advanced_options.inherited_waf / 65175dd9c168 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-502bc511c813f0f72a330e46541205385d998d848ba2a76c555e2bc7eab7cbfb"></a>

## Next pages — routes.simple_route.advanced_options.inherited_waf / 65175dd9c168 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-821811c0f4d76445f73d7e659d3d92d74205dd708146d942d0348fba9e1a632d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a12b82b6ab82b6b7ce3953eb49edcbbdd668c893116db3db30f20174bc87fee2"></a>

## routes.simple_route.advanced_options.inherited_waf_exclusion — routes.simple_route.advanced_options.inherited_waf_exclusion / 720b72f42eb4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.inherited_waf_exclusion

<a id="canonical-eb592bb4f94ee25bfdbd92a807266ce8ec67afc08ca51e474da2eb63208c1efd"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherited waf exclusion.

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
inherited_waf_exclusion = {}
```

<a id="canonical-1dad4fcccca32ac265267b9c0728df28337dda250f46ea4b24034e92402f2a89"></a>

## Direct properties — routes.simple_route.advanced_options.inherited_waf_exclusion / 720b72f42eb4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-99ac887bf040010f594776c34423f82b950e348a56b05bdce40b425d67880bfc"></a>

## Next pages — routes.simple_route.advanced_options.inherited_waf_exclusion / 720b72f42eb4 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-47069f6847d19857cb08c8f8cc04bf77d1cdf5df9600e52489a86f8af3f1ebcb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10fa950cc1873151ad9ad3b02e7bcc6681eaa71dff860986ddcf20216fe4ecdc"></a>

## routes.simple_route.advanced_options.mirror_policy — routes.simple_route.advanced_options.mirror_policy / a954d4556d5d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.mirror_policy

<a id="canonical-e3d34b9e54d6c73bd03bd0347e0de8d787e4e76e69d7e49d6424aac26495de85"></a>

Type: `"object"`. single nested block, Optional.

MirrorPolicy is used for shadowing traffic from one origin pool to another. The approach used is
'fire and forget', meaning it will not wait for the shadow origin pool to respond before returning
the response from the primary origin pool. All normal statistics are collected for the shadow
origin..

Upstream description:

MirrorPolicy is used for shadowing traffic from one origin pool to another. The approach used is
"fire and forget", meaning it will not wait for the shadow origin pool to respond before returning
the response from the primary origin pool. All normal statistics are collected for the shadow origin
pool making this feature useful for testing and troubleshooting.

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
mirror_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-e52e2e7b977032e8bc9b0aef05877b21cd575246be28da31219cb20d2967bd42"></a>

## Direct properties — routes.simple_route.advanced_options.mirror_policy / a954d4556d5d / 3

- [origin_pool](resources--http_loadbalancer--reference--group-024.md#canonical-53ae44effb2ffad3e23e552cdb5689b368e1cc76b350cb3bfba41405a97418e1): complete subsection reference.

- [percent](resources--http_loadbalancer--reference--group-024.md#canonical-b76610b1564cd34ca76c50b93ca0634af2ce6952d98f74859966cdc2d280eccd): complete subsection reference.

<a id="canonical-566e3212258b21da9696117ded52a326b245503729ce0f23e6c7ef0f64cfc4f7"></a>

## Next pages — routes.simple_route.advanced_options.mirror_policy / a954d4556d5d / 4

- [routes.simple_route.advanced_options.mirror_policy.origin_pool](resources--http_loadbalancer--reference--group-024.md#canonical-53ae44effb2ffad3e23e552cdb5689b368e1cc76b350cb3bfba41405a97418e1)
- [routes.simple_route.advanced_options.mirror_policy.percent](resources--http_loadbalancer--reference--group-024.md#canonical-b76610b1564cd34ca76c50b93ca0634af2ce6952d98f74859966cdc2d280eccd)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-53ae44effb2ffad3e23e552cdb5689b368e1cc76b350cb3bfba41405a97418e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34c46e72c6008967c1eec7110b501108484275ac3cd1f7b9d4b433ad08bef101"></a>

## routes.simple_route.advanced_options.mirror_policy.origin_pool — routes.simple_route.advanced_options.mirror_policy.origin_pool / e6764d48818f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [routes.simple_route.advanced_options.mirror_policy](resources--http_loadbalancer--reference--group-024.md#canonical-47069f6847d19857cb08c8f8cc04bf77d1cdf5df9600e52489a86f8af3f1ebcb)
- routes.simple_route.advanced_options.mirror_policy.origin_pool

<a id="canonical-3cd87bf3f2f76ca5dc8f1d03fbfba789d769fc7ed00368dd682f23e7764398bd"></a>

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
origin_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-ef1ecc042d6935828eba6526bcea43a9a304985c1857404231429a35300e4221"></a>

## Direct properties — routes.simple_route.advanced_options.mirror_policy.origin_pool / e6764d48818f / 3

<a id="canonical-e55db12b4b0986d4b35d48128734e6713b2e0d20c707bb6705cdbf156ad9c17f"></a>

<a id="canonical-320fd4a44dc9e99f24e3af7463aefc2e5c81e49f3b2319143ffdb23f571f2786"></a>

## name property — routes.simple_route.advanced_options.mirror_policy.origin_pool / e6764d48818f / 4

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

<a id="canonical-5143ea047ddcb242b24802cb5cd7418a0f9a4985211dcc2eef8f4aa7c5338c4a"></a>

<a id="canonical-32383820a17675a8110ba6c9726e2fe48947336a632df0e6677adcafde5c38a7"></a>

## namespace property — routes.simple_route.advanced_options.mirror_policy.origin_pool / e6764d48818f / 5

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

<a id="canonical-d002e4ae678d0f274cd3f9d4ab5b4429bed4fd799fa39542dfbc15772e0c56ce"></a>

<a id="canonical-0210c6838c0db32001ef92c1e901236b046a727a12928da2d5a047fbc9aa3da2"></a>

## tenant property — routes.simple_route.advanced_options.mirror_policy.origin_pool / e6764d48818f / 6

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

<a id="canonical-fe0f48b4ed57a19afe946752092372aa2d5f3c37461f37a540337c488fab5925"></a>

## Next pages — routes.simple_route.advanced_options.mirror_policy.origin_pool / e6764d48818f / 7

- [routes.simple_route.advanced_options.mirror_policy](resources--http_loadbalancer--reference--group-024.md#canonical-47069f6847d19857cb08c8f8cc04bf77d1cdf5df9600e52489a86f8af3f1ebcb)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b76610b1564cd34ca76c50b93ca0634af2ce6952d98f74859966cdc2d280eccd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65a721ca04cef007e78cbd94bd14c5301ab436d129d4d8c8a71c24f428c28de3"></a>

## routes.simple_route.advanced_options.mirror_policy.percent — routes.simple_route.advanced_options.mirror_policy.percent / 4600c2edd2d5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [routes.simple_route.advanced_options.mirror_policy](resources--http_loadbalancer--reference--group-024.md#canonical-47069f6847d19857cb08c8f8cc04bf77d1cdf5df9600e52489a86f8af3f1ebcb)
- routes.simple_route.advanced_options.mirror_policy.percent

<a id="canonical-ff401e619b9d9e6f26804f2b70f2498b8d369c9fbdf0b4a33e00c964803ab195"></a>

Type: `"object"`. single nested block, Optional.

Fraction used where sampling percentages are needed. Example sampled requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("numerator")}
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
percent {
  # Configure direct properties listed below.
}
```

<a id="canonical-ca8c29db33f49bf8b6613adf9e63b9401c9433b67cdc22f404ef9534ec7e6ef7"></a>

## Direct properties — routes.simple_route.advanced_options.mirror_policy.percent / 4600c2edd2d5 / 3

<a id="canonical-424733673b9c0f26fd001c32a65129541b2d2717899ae420fa9a789c48c673e4"></a>

<a id="canonical-d82e859f335cb8fbbe1bf7af6395a16cce91cf9120495d8b769feed143406f7f"></a>

## denominator property — routes.simple_route.advanced_options.mirror_policy.percent / 4600c2edd2d5 / 4

Type: `"string"`. Optional.

\[Enum: HUNDRED|TEN\_THOUSAND|MILLION\] Denominator used in fraction where sampling percentages are
needed. Example sampled requests Use hundred as denominator Use ten thousand as denominator Use
million as denominator. Possible values are \`HUNDRED\`, \`TEN\_THOUSAND\`, \`MILLION\`. Defaults to
\`HUNDRED\`.

Upstream description:

Denominator used in fraction where sampling percentages are needed. Example sampled requests

Use hundred as denominator Use ten thousand as denominator Use million as denominator.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("HUNDRED",
    "TEN_THOUSAND",
    "MILLION"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "HUNDRED",
  "enum": [
    "HUNDRED",
    "TEN_THOUSAND",
    "MILLION"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c5ed98fd5d732e89d12ae91315eb4199614198528d0360112e102531529a00a9"></a>

<a id="canonical-736d304c1b21d943981c55ee1a062ea8a1254dca8ca708eabb369c5334f739d1"></a>

## numerator property — routes.simple_route.advanced_options.mirror_policy.percent / 4600c2edd2d5 / 5

Type: `"number"`. Optional.

Sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1eecbddce4d3ea2a46fffd3a2fa511afa0b3d660933c0ca59677fd6fffdf6e9f"></a>

## Next pages — routes.simple_route.advanced_options.mirror_policy.percent / 4600c2edd2d5 / 6

- [routes.simple_route.advanced_options.mirror_policy](resources--http_loadbalancer--reference--group-024.md#canonical-47069f6847d19857cb08c8f8cc04bf77d1cdf5df9600e52489a86f8af3f1ebcb)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a8fd9cd2de2bff8f95eb44dc2caaede6525e0322e23b32ff98f07a0718da8c15"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a71da555a1c3aa77f1a7cd7114f247d7d88945985a2bce01e556b7d7f5972691"></a>

## routes.simple_route.advanced_options.no_retry_policy — routes.simple_route.advanced_options.no_retry_policy / 277c4c16aa48 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.no_retry_policy

<a id="canonical-eb510f64350ffef80ee44377b1e186723fa318b7a3d40acfff39b70bca9faafd"></a>

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
no_retry_policy = {}
```

<a id="canonical-1a4816c30baa00e6eebb3332c47eaf7a73241544b2e6d055b2db0e6f9614bb12"></a>

## Direct properties — routes.simple_route.advanced_options.no_retry_policy / 277c4c16aa48 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-89fb96d1b0a8e2728335fa55d578be4913f2954b8c7c1179abf9265846e71595"></a>

## Next pages — routes.simple_route.advanced_options.no_retry_policy / 277c4c16aa48 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-381158a64be0f1d30cd5a2dd7b735d6bdc6d12ea248a26fc03ccab67f33b613e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b47713f9d01975202342812313444664f073de9e482671f3dbad5b548ace674e"></a>

## routes.simple_route.advanced_options.regex_rewrite — routes.simple_route.advanced_options.regex_rewrite / a7c286c422f7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.regex_rewrite

<a id="canonical-7e9e176f754d7bb703517b1a0e0d1f4c7aa178171a25952c4d86a744deaee576"></a>

Type: `"object"`. single nested block, Optional.

RegexMatchRewrite describes how to match a string and then produce a new string using a regular
expression and a substitution string.

Upstream description:

RegexMatchRewrite describes how to match a string and then produce a new string using a regular
expression and a substitution string.

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
regex_rewrite {
  # Configure direct properties listed below.
}
```

<a id="canonical-532f7c115f7c5a7cd967a55490cb603cccbb3f745352d3f6d89fa60da3b28cdf"></a>

## Direct properties — routes.simple_route.advanced_options.regex_rewrite / a7c286c422f7 / 3

<a id="canonical-b6918f794e230480fe12da9b10671db3c73f563e121e77b437f6acbc8b9ac990"></a>

<a id="canonical-ab68a6bffd849e50fd3dfd2b39c01d8ddbb7093b13ddb975f5821e98ffac81ab"></a>

## pattern property — routes.simple_route.advanced_options.regex_rewrite / a7c286c422f7 / 4

Type: `"string"`. Optional.

The regular expression used to find portions of a string that should be replaced.

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

<a id="canonical-d0f67224da8812e41e4f1aa4806e0ea52d2778ac5d01aa207e6b53aa336ba970"></a>

<a id="canonical-5a1b5580ac2fbd83d95e19e110c42af3bb4da781917308b787e52b6412232fe9"></a>

## substitution property — routes.simple_route.advanced_options.regex_rewrite / a7c286c422f7 / 5

Type: `"string"`. Optional.

The string that should be substituted into matching portions of the subject string during a
substitution operation to produce a new string.

Upstream description:

The string that should be substituted into matching portions of the subject string during a
substitution operation to produce a new string.

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

<a id="canonical-897c84c03fb2906196ed84957aab3b70ad6062c1f8213ffe3b952076a3362a70"></a>

## Next pages — routes.simple_route.advanced_options.regex_rewrite / a7c286c422f7 / 6

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8bae6e2b386d2a401ab31abe669d8b768a7b45d7624d13e25ca80e373d921539"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca80047c8c0d895911c64f212f487455fceefef4f7055b2cdea91c92a2b3d814"></a>

## routes.simple_route.advanced_options.request_cookies_to_add — routes.simple_route.advanced_options.request_cookies_to_add / 8aba7045f7ef / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.request_cookies_to_add

<a id="canonical-8b6daf06455a92cab41bca4eacdeabcd0fc948f1fd3af956b25c05c3fc0d7b8b"></a>

Type: `"object"`. list nested block, Optional.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Upstream description:

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
request_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-292a512e2e90f02a967d0f94a0db321f1086b612d554d2b2e4799014c303122f"></a>

## Direct properties — routes.simple_route.advanced_options.request_cookies_to_add / 8aba7045f7ef / 3

<a id="canonical-5d3b3ca367c432c09560800a24aaf956b49175d6166830e84e74f6626544a384"></a>

<a id="canonical-5f4135fdf1813802913c645bc2f0f4480bde4997444704a641ff926fdacb62a7"></a>

## name property — routes.simple_route.advanced_options.request_cookies_to_add / 8aba7045f7ef / 4

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-06c23a2dd5099b1df536c3278a53b7a8944136ff4cf84495d19a0ea1573a60c0"></a>

<a id="canonical-9ae47790638d829d2680c2905a4370cdeb24c506a4d690830cd2a97e6af682f5"></a>

## overwrite property — routes.simple_route.advanced_options.request_cookies_to_add / 8aba7045f7ef / 5

Type: `"bool"`. Optional.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

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

- [secret_value](resources--http_loadbalancer--reference--group-024.md#canonical-ee75a47d7299da238cbddbe0ab904cb51388e92dce362abf3a4b792ce88b5659): complete subsection reference.

<a id="canonical-86bcc03ce8bdb52ebf374f895f1bb27571aa176571972414f76c45e853973ff6"></a>

<a id="canonical-dd009230d972a7c2b42f25bad7134e4e9fa28a89e3c39adaa32cd5213e5fefa1"></a>

## value property — routes.simple_route.advanced_options.request_cookies_to_add / 8aba7045f7ef / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-21edce8ffb9e530c0d28bb6b1bbce94fb00a92ca851d423da89b8682d859fda6"></a>

## Next pages — routes.simple_route.advanced_options.request_cookies_to_add / 8aba7045f7ef / 7

- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-024.md#canonical-ee75a47d7299da238cbddbe0ab904cb51388e92dce362abf3a4b792ce88b5659)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ee75a47d7299da238cbddbe0ab904cb51388e92dce362abf3a4b792ce88b5659"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-987ec9f7113302da1917e0d9e99d095fbaf018b348dec2367a9741f4c134fa3d"></a>

## routes.simple_route.advanced_options.request_cookies_to_add.secret_value — routes.simple_route.advanced_options.request_cookies_to_add.secret_value / 04e09a53486e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [routes.simple_route.advanced_options.request_cookies_to_add](resources--http_loadbalancer--reference--group-024.md#canonical-8bae6e2b386d2a401ab31abe669d8b768a7b45d7624d13e25ca80e373d921539)
- routes.simple_route.advanced_options.request_cookies_to_add.secret_value

<a id="canonical-a959f460b9029dc4ef01b01effa28f86a8c58f9b333054ba294c63daf73c5868"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-97752937e1a056a4b59231bfd40fc87995e3f85bc5debc8fd65c1fbb574b9adc"></a>

## Direct properties — routes.simple_route.advanced_options.request_cookies_to_add.secret_value / 04e09a53486e / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-024.md#canonical-26b67b36d2594a655028336ef83e425fde6268897865a1e9bc4d745512196d9e): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-024.md#canonical-9043920518d6f586d9efa5d3b4431b79cc9161c1b10b51b9244c5c42652de8ee): complete subsection reference.

<a id="canonical-02c0ee424d02f58da89f18b3e240c91abb5cd611575889725699332c7d52222e"></a>

## Next pages — routes.simple_route.advanced_options.request_cookies_to_add.secret_value / 04e09a53486e / 4

- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info](resources--http_loadbalancer--reference--group-024.md#canonical-26b67b36d2594a655028336ef83e425fde6268897865a1e9bc4d745512196d9e)
- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_secret_info](resources--http_loadbalancer--reference--group-024.md#canonical-9043920518d6f586d9efa5d3b4431b79cc9161c1b10b51b9244c5c42652de8ee)
- [routes.simple_route.advanced_options.request_cookies_to_add](resources--http_loadbalancer--reference--group-024.md#canonical-8bae6e2b386d2a401ab31abe669d8b768a7b45d7624d13e25ca80e373d921539)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-26b67b36d2594a655028336ef83e425fde6268897865a1e9bc4d745512196d9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e08a88976f91280e3339b96c2256f03cef004397a206ea92b459ac1ea385e472"></a>

## routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfo / c59f26c7cd4d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [routes.simple_route.advanced_options.request_cookies_to_add](resources--http_loadbalancer--reference--group-024.md#canonical-8bae6e2b386d2a401ab31abe669d8b768a7b45d7624d13e25ca80e373d921539)
- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-024.md#canonical-ee75a47d7299da238cbddbe0ab904cb51388e92dce362abf3a4b792ce88b5659)
- routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-b96f6e1c950253a87f67a7b75d763f27915fce162f33d7e6e701f16875c92343"></a>

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

<a id="canonical-00010b45b19873fad3081eef19fb0e5cef95824c1f811aabf6724ed5842a882e"></a>

## Direct properties — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfo / c59f26c7cd4d / 3

<a id="canonical-729e8ecea7c97a7650e63ff6253243545ff11aed41f549222e66398d7d6cda0f"></a>

<a id="canonical-06aedd5f3f93f79d58eeba7697a0b5a117beb3c9b22820549672264a25d29b29"></a>

## decryption_provider property — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfo / c59f26c7cd4d / 4

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

<a id="canonical-cdb6cbb9a2967c5e843c235a599a198339c2ebd6938b1679ed96595ce5360beb"></a>

<a id="canonical-366a9b62ddbe41df03ea793f09c48d82a5a66e0cc8dca021fef8fbb9a849c235"></a>

## location property — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfo / c59f26c7cd4d / 5

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

<a id="canonical-ad18c0194cda896412701d719825029b8b03766783e3fc87a3d9626ba7640665"></a>

<a id="canonical-acd9b6b0a5168a1c9c23d7384cef20d3d71228c2b5366490ebdee681878d2bfa"></a>

## store_provider property — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfo / c59f26c7cd4d / 6

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

<a id="canonical-24336dcae4194f162f8543d020ad4bae45a093e1fd50a23d7bc0280e939ada66"></a>

## Next pages — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfo / c59f26c7cd4d / 7

- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-024.md#canonical-ee75a47d7299da238cbddbe0ab904cb51388e92dce362abf3a4b792ce88b5659)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9043920518d6f586d9efa5d3b4431b79cc9161c1b10b51b9244c5c42652de8ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30534c8974ef7664181551a57e7b5d94ef3323cc440c2766333218d1444ced38"></a>

## routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_secret_info — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_s / dd3d04ee42cf / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [routes.simple_route.advanced_options.request_cookies_to_add](resources--http_loadbalancer--reference--group-024.md#canonical-8bae6e2b386d2a401ab31abe669d8b768a7b45d7624d13e25ca80e373d921539)
- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-024.md#canonical-ee75a47d7299da238cbddbe0ab904cb51388e92dce362abf3a4b792ce88b5659)
- routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-94a4bfd1bff4b17bf0c63ce4663d4585487143354fdfeb1ab5d84e8a07445fde"></a>

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

<a id="canonical-1548fc3881575b37093286b2c8aaa083a1430da092da52fbf3b8cd3f4069bbfa"></a>

## Direct properties — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_s / dd3d04ee42cf / 3

<a id="canonical-63882f4c72eb1f4e6e3da948f74859aa1c3ecce860bed41a91d050f61a55b0c3"></a>

<a id="canonical-6d97b79ed4d11c4ff6460cceec451e7ec9d1e16fd45b4efd5ec084e259ad8a1e"></a>

## provider_ref property — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_s / dd3d04ee42cf / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3832d7ff5a18ec6840f8e3d04c8de1ac66709050d42944e4e393b0120370a6c8"></a>

<a id="canonical-a7961e9485658ce8c7d5b97f1953d9a80149afd1f82035e93412fef74dd83ecb"></a>

## url property — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_s / dd3d04ee42cf / 5

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

<a id="canonical-16be067f5bd24d60c7e0d6837114651fac318f754825bf6509a25cd13c120fb4"></a>

## Next pages — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_s / dd3d04ee42cf / 6

- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-024.md#canonical-ee75a47d7299da238cbddbe0ab904cb51388e92dce362abf3a4b792ce88b5659)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-fe40cbfb64ecdbc429fb055ad7e50015a785212961578e768a42bccfe4952677"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b12ea82967cd2c7dd48cacf6c7a5990c7a19b1df18a52dda730e709e28670e49"></a>

## routes.simple_route.advanced_options.request_headers_to_add — routes.simple_route.advanced_options.request_headers_to_add / e6cb93fa1c08 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.request_headers_to_add

<a id="canonical-8a17436007ceda3404d76e4c99a7c0ef9e92eb532b36defd33072da0c17986cb"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP request being routed towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
request_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-70c863ecf007f9aee7c861692d54809a7ae5a20bb72168040173a70c8af2dc83"></a>

## Direct properties — routes.simple_route.advanced_options.request_headers_to_add / e6cb93fa1c08 / 3

<a id="canonical-abab8062352d829ce254ea2d0479e1631ef40673f1067f68a5571f27e8aa580b"></a>

<a id="canonical-7aef36a555f09d3f78dc20ee34ecb25fb9cb62adb6bed0eab4cf603b90ee8f4c"></a>

## append property — routes.simple_route.advanced_options.request_headers_to_add / e6cb93fa1c08 / 4

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

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

<a id="canonical-e422ac7ff2a4dad76e943b86d2efa028fe24a9d729cf752082295916d4e97cd8"></a>

<a id="canonical-375b7ac0be55705757abe97c45068f1d9285b3ed46e63aa4f7d63b8b4e37ffa8"></a>

## name property — routes.simple_route.advanced_options.request_headers_to_add / e6cb93fa1c08 / 5

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](resources--http_loadbalancer--reference--group-024.md#canonical-cf798906b25a9c791420618295509f1d545974a0794586626fc2485875587b98): complete subsection reference.

<a id="canonical-b51890be9cb6832057c6792ce23bf307955733f32ed38bb397185d9d0499c3ed"></a>

<a id="canonical-775c0e2828f124a59d902be652b5a4cb8f50c2635342f42f77d5eb2eb28764a5"></a>

## value property — routes.simple_route.advanced_options.request_headers_to_add / e6cb93fa1c08 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-3ddb2d3e9454c6fce402e43f6e08b685e62ff7bcc12990a34ae9c45f4ede3dea"></a>

## Next pages — routes.simple_route.advanced_options.request_headers_to_add / e6cb93fa1c08 / 7

- [routes.simple_route.advanced_options.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-024.md#canonical-cf798906b25a9c791420618295509f1d545974a0794586626fc2485875587b98)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-cf798906b25a9c791420618295509f1d545974a0794586626fc2485875587b98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d17ec5ef32e33a01640a41deb96d0c8f925f1aa7a0ee43201fedeb52bfc13447"></a>

## routes.simple_route.advanced_options.request_headers_to_add.secret_value — routes.simple_route.advanced_options.request_headers_to_add.secret_value / d9f48047ec81 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [routes.simple_route.advanced_options.request_headers_to_add](resources--http_loadbalancer--reference--group-024.md#canonical-fe40cbfb64ecdbc429fb055ad7e50015a785212961578e768a42bccfe4952677)
- routes.simple_route.advanced_options.request_headers_to_add.secret_value

<a id="canonical-13bc739dd2c36790ee87068f070fbe762bca64577f558cc6adf6fd3a7184a0b8"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-426029d11777b4c876deeb70b99635c5ea3d77ba86733bc83fed865896248fd1"></a>

## Direct properties — routes.simple_route.advanced_options.request_headers_to_add.secret_value / d9f48047ec81 / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-024.md#canonical-84c002fa27e20f325dc5c9d9997f58a9682a437c1e211ef53c11eab34b61ea91): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-024.md#canonical-ff8cd53e47c4536e5a5b514fccd2556db81a25b76fbddb2912d242045749c698): complete subsection reference.

<a id="canonical-6a0fc2468493b1956a42ccfbb668942b2c1abea165112531fdd272b351a23d63"></a>

## Next pages — routes.simple_route.advanced_options.request_headers_to_add.secret_value / d9f48047ec81 / 4

- [routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info](resources--http_loadbalancer--reference--group-024.md#canonical-84c002fa27e20f325dc5c9d9997f58a9682a437c1e211ef53c11eab34b61ea91)
- [routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_secret_info](resources--http_loadbalancer--reference--group-024.md#canonical-ff8cd53e47c4536e5a5b514fccd2556db81a25b76fbddb2912d242045749c698)
- [routes.simple_route.advanced_options.request_headers_to_add](resources--http_loadbalancer--reference--group-024.md#canonical-fe40cbfb64ecdbc429fb055ad7e50015a785212961578e768a42bccfe4952677)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-84c002fa27e20f325dc5c9d9997f58a9682a437c1e211ef53c11eab34b61ea91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23b57b347db70950d53ced3b8cc1ce86313d6882be77e3f1b742cb4b2afb505b"></a>

## routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info — routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfo / 3d7d22052ffb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [routes.simple_route.advanced_options.request_headers_to_add](resources--http_loadbalancer--reference--group-024.md#canonical-fe40cbfb64ecdbc429fb055ad7e50015a785212961578e768a42bccfe4952677)
- [routes.simple_route.advanced_options.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-024.md#canonical-cf798906b25a9c791420618295509f1d545974a0794586626fc2485875587b98)
- routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-3d508fef58b43358416dc48646b9dc20268675c3a86ffbf0fbfd113a8ab9bbe9"></a>

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

<a id="canonical-55422508628959012e87c6604552e12a76c0ffe73fd2d6d5571e171ba9c9788a"></a>

## Direct properties — routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfo / 3d7d22052ffb / 3

<a id="canonical-066c7b0e7d617898acdfbc80f9528548511ba7d50c7613335779ee5a63ae2dc4"></a>

<a id="canonical-3991922f68b7374f322a8b8a14f896b86098e1a32c07ba67af860a890bd2a309"></a>

## decryption_provider property — routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfo / 3d7d22052ffb / 4

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

<a id="canonical-0ffd90bbc030aeb8e93dd6e46c6ec624311f66fadcc0fec3c0f0ed21398c5544"></a>

<a id="canonical-dbd3af112eee7e81032cfed09fdacfb95aa0a3c69a2e62d9faf59b14e9a7efdb"></a>

## location property — routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfo / 3d7d22052ffb / 5

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

<a id="canonical-64d67e197eaa5a681cdf5a7e1d2f3663a07e38b0cd458cb24600a95f170fdb85"></a>

<a id="canonical-83520a41bc298d9568facb2188d81629200d6afb75501117df6323bc53f698d7"></a>

## store_provider property — routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfo / 3d7d22052ffb / 6

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

<a id="canonical-05a3926e81e0fc175c673c080114da6b3607388493485b5fb0887c55694cb0b4"></a>

## Next pages — routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfo / 3d7d22052ffb / 7

- [routes.simple_route.advanced_options.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-024.md#canonical-cf798906b25a9c791420618295509f1d545974a0794586626fc2485875587b98)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ff8cd53e47c4536e5a5b514fccd2556db81a25b76fbddb2912d242045749c698"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8666bad4e878d360efe4558d302eb4c36e40e86be4ebdf1497981143f7f57f8a"></a>

## routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_secret_info — routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_s / e5bdafb885af / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [routes.simple_route.advanced_options.request_headers_to_add](resources--http_loadbalancer--reference--group-024.md#canonical-fe40cbfb64ecdbc429fb055ad7e50015a785212961578e768a42bccfe4952677)
- [routes.simple_route.advanced_options.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-024.md#canonical-cf798906b25a9c791420618295509f1d545974a0794586626fc2485875587b98)
- routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-51570bb8261c2deff1089eb3383a0bc5268bbf598df8af0d7f742d0f143544fd"></a>

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

<a id="canonical-a9ed13fff5c475d71a095d20fc3f8acf3a22c39f2bb170aef947db18c3e160d2"></a>

## Direct properties — routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_s / e5bdafb885af / 3

<a id="canonical-74ce52677d6d486083f88d0a300363d44720bfb691e025420306db697eeacedf"></a>

<a id="canonical-3b72952d5d3916ca9fad2b115b5eadcc1357fbbac8653bac5398063607718ef4"></a>

## provider_ref property — routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_s / e5bdafb885af / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-ac1abdfe61e7db04ccfbac39b754a88a82abff53d63e4e97042ca0c7cf243782"></a>

<a id="canonical-cedcfa82ee419490fc2041b68fff552cbef6faf0ce1674eb303d8747fc24928f"></a>

## url property — routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_s / e5bdafb885af / 5

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

<a id="canonical-80017b4e903c377812cb9030caec842188b9bd456d900c6a6819587b5aabef25"></a>

## Next pages — routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_s / e5bdafb885af / 6

- [routes.simple_route.advanced_options.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-024.md#canonical-cf798906b25a9c791420618295509f1d545974a0794586626fc2485875587b98)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3985b17e1a5096f3db4c6a657ac616194c986fe82e7d9f82e728fcd6ee04bbc8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-974c4c50af6419cd4b150d68542648beba3a8350c1310cf136356bb512d51dd2"></a>

## routes.simple_route.advanced_options.response_cookies_to_add — routes.simple_route.advanced_options.response_cookies_to_add / 500da90c231d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- routes.simple_route.advanced_options.response_cookies_to_add

<a id="canonical-a2caa31db43012b71e6ecf72bf9f5232527bfc5e7127ca189a09f94c69cf62da"></a>

Type: `"object"`. list nested block, Optional.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Upstream description:

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("add_domain",
    "ignore_domain"),
  validators.ConflictingListObjectAttributes("add_expiry",
    "ignore_expiry"),
  validators.ConflictingListObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingListObjectAttributes("add_partitioned",
    "ignore_partitioned"),
  validators.ConflictingListObjectAttributes("add_path",
    "ignore_path"),
  validators.ConflictingListObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingListObjectAttributes("ignore_max_age",
    "max_age_value"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "secret_value"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "value"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_none",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
response_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-9a4339de1af8eb3874f63462363efdd3dab77d2f58bd2a3a49a58961106b10da"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add / 500da90c231d / 3

<a id="canonical-abc343778fd758b9dd83e4666f53dbcdc9690fbdd8175fe1db7fac2badf2ff88"></a>

<a id="canonical-217c3036e190d6935d9b55bf1892fa8e1ea9051de9f130b1b2ca783b91fb6a76"></a>

## add_domain property — routes.simple_route.advanced_options.response_cookies_to_add / 500da90c231d / 4

Type: `"string"`. Optional.

Exclusive with \[ignore\_domain\] Add domain attribute.

Upstream description:

Exclusive with \[ignore\_domain\] Add domain attribute.

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

<a id="canonical-b5727653bcf9902adb2431d089be0ec3cdcebda83a9b82418994d95943066f43"></a>

<a id="canonical-2bab58403ffd55e04bf697cf1284d1bc37d04d21b10001bc830cfc6f51d6d4aa"></a>

## add_expiry property — routes.simple_route.advanced_options.response_cookies_to_add / 500da90c231d / 5

Type: `"string"`. Optional.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Upstream description:

Exclusive with \[ignore\_expiry\] Add expiry attribute.

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

- [add_httponly](resources--http_loadbalancer--reference--group-024.md#canonical-b732ad7d34c58ee0e651ab59950072b090c2a3ed94694135c9a6c2782b0eeafb): complete subsection reference.

- [add_partitioned](resources--http_loadbalancer--reference--group-024.md#canonical-7bb8622d7bb8ffd00a83326f5e84bc4a754308e24e861a28a67720c3b218575b): complete subsection reference.

<a id="canonical-9ad2f5690b25d353e202452d804ce04c4a8a29ea2873c60d0a411eedbc9cceec"></a>

<a id="canonical-b7417a5c16f0d474effc4a1211ba81c3cd5c5471e2dd99361e0eb2e8b2054d26"></a>

## add_path property — routes.simple_route.advanced_options.response_cookies_to_add / 500da90c231d / 6

Type: `"string"`. Optional.

Exclusive with \[ignore\_path\] Add path attribute.

Upstream description:

Exclusive with \[ignore\_path\] Add path attribute.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_secure](resources--http_loadbalancer--reference--group-025.md#canonical-270b9ccde0e895dcd8de8bd037f6d8fe31a35b1e0f036bec29d3f053c6ac0f92): complete subsection reference.

- [ignore_domain](resources--http_loadbalancer--reference--group-025.md#canonical-6db34b75cd18371823f41749f366fb3e00e094c99a36485e2999208048458007): complete subsection reference.

- [ignore_expiry](resources--http_loadbalancer--reference--group-025.md#canonical-c0f2d529c41ca25f15f61bde8bcb96eb5f3daf7d74c03ad971c8beac1c4a861a): complete subsection reference.

- [ignore_httponly](resources--http_loadbalancer--reference--group-025.md#canonical-cf0c9cd4a9e2e54ae49d81c6620b775115314554d9b0df04788f2bf1914487ec): complete subsection reference.

- [ignore_max_age](resources--http_loadbalancer--reference--group-025.md#canonical-fe41e19461374f29f006d1bb9b2dbdd47808e691943171e0c0e32dca3779bc04): complete subsection reference.

- [ignore_partitioned](resources--http_loadbalancer--reference--group-025.md#canonical-0518b879a328f8c39c755d83621392c38de3dfd95f729a22dcd95557d61670f1): complete subsection reference.

- [ignore_path](resources--http_loadbalancer--reference--group-025.md#canonical-dccc840b4e944a47051d15039048e26282b9db515636d3d9ed1cc5e2759bf04b): complete subsection reference.

- [ignore_samesite](resources--http_loadbalancer--reference--group-025.md#canonical-570d1d614cef1276f89ada309d1016217011b46549163a38ab09f8cafce0ad53): complete subsection reference.

- [ignore_secure](resources--http_loadbalancer--reference--group-025.md#canonical-19434c68f229ed00022e54ea7db35e3b11e9d3c6e619a3e612e693b36d8e8218): complete subsection reference.

- [ignore_value](resources--http_loadbalancer--reference--group-025.md#canonical-3c3b61e0a7d25532f5557d33488954b2bd98222e4edc2b876502c80663c4f59e): complete subsection reference.

<a id="canonical-bdf3ff8af77d5189a9ceb1f187ee8ee4d12cfc0f8a8934efd81a94c3ff537539"></a>

<a id="canonical-26aa436a66b6815ba1eb22d8f9b27cf350b8d3296ae938f5d0c04d165a151bb7"></a>

## max_age_value property — routes.simple_route.advanced_options.response_cookies_to_add / 500da90c231d / 7

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(34560000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-0a1f3d8301fd24ca0f779a34b9a418835d2fc164df0dedc6d591b06bee4fe0e6"></a>

<a id="canonical-d8eeec5ce3a6cff0716945a3f748e8af7853986ed1e3aba45cb3e131ae280757"></a>

## name property — routes.simple_route.advanced_options.response_cookies_to_add / 500da90c231d / 8

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-fbe5964028c6790a3d9acb6f635feb9876dbbb51b4b735cc39d258a4b96860be"></a>

<a id="canonical-e7d2100553eba3cf0d2273bdcce9499d926f9f783b0415d7bc4b8e7f4a5f5b18"></a>

## overwrite property — routes.simple_route.advanced_options.response_cookies_to_add / 500da90c231d / 9

Type: `"bool"`. Optional.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

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

- [samesite_lax](resources--http_loadbalancer--reference--group-025.md#canonical-9a41c832843e252991afe87976d397ba773d4314bb035da9ccf540130e5b3d9a): complete subsection reference.

- [samesite_none](resources--http_loadbalancer--reference--group-025.md#canonical-e724faa2c6f35a235c34b8537260157dcc76b62ea612d67107f35e166ff3e9ae): complete subsection reference.

- [samesite_strict](resources--http_loadbalancer--reference--group-025.md#canonical-156f1399f25dec1e4d2cb6eb7aa7223513a9c564b210daed8eae93f88f1872b7): complete subsection reference.

- [secret_value](resources--http_loadbalancer--reference--group-025.md#canonical-23987764d6be2d466524adabb157ea49f0b9d54983f821d1db43d1d75e1d9d39): complete subsection reference.

<a id="canonical-89dbd84fef0b5c3f1a7be647b2ecf4283220a070633c808c36379a9b9d26563f"></a>

<a id="canonical-74dbaf5e4bc4e62d32e3877f1a424308f72d58d7c49d29ed6a28e9f43c89a6df"></a>

## value property — routes.simple_route.advanced_options.response_cookies_to_add / 500da90c231d / 10

Type: `"string"`. Optional.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-130e6ab00b9d674190584c01b6f909995e3d1116036d0274de2d39e02b321e5a"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add / 500da90c231d / 11

- [routes.simple_route.advanced_options.response_cookies_to_add.add_httponly](resources--http_loadbalancer--reference--group-024.md#canonical-b732ad7d34c58ee0e651ab59950072b090c2a3ed94694135c9a6c2782b0eeafb)
- [routes.simple_route.advanced_options.response_cookies_to_add.add_partitioned](resources--http_loadbalancer--reference--group-024.md#canonical-7bb8622d7bb8ffd00a83326f5e84bc4a754308e24e861a28a67720c3b218575b)
- [routes.simple_route.advanced_options.response_cookies_to_add.add_secure](resources--http_loadbalancer--reference--group-025.md#canonical-270b9ccde0e895dcd8de8bd037f6d8fe31a35b1e0f036bec29d3f053c6ac0f92)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_domain](resources--http_loadbalancer--reference--group-025.md#canonical-6db34b75cd18371823f41749f366fb3e00e094c99a36485e2999208048458007)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_expiry](resources--http_loadbalancer--reference--group-025.md#canonical-c0f2d529c41ca25f15f61bde8bcb96eb5f3daf7d74c03ad971c8beac1c4a861a)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_httponly](resources--http_loadbalancer--reference--group-025.md#canonical-cf0c9cd4a9e2e54ae49d81c6620b775115314554d9b0df04788f2bf1914487ec)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_max_age](resources--http_loadbalancer--reference--group-025.md#canonical-fe41e19461374f29f006d1bb9b2dbdd47808e691943171e0c0e32dca3779bc04)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_partitioned](resources--http_loadbalancer--reference--group-025.md#canonical-0518b879a328f8c39c755d83621392c38de3dfd95f729a22dcd95557d61670f1)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_path](resources--http_loadbalancer--reference--group-025.md#canonical-dccc840b4e944a47051d15039048e26282b9db515636d3d9ed1cc5e2759bf04b)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_samesite](resources--http_loadbalancer--reference--group-025.md#canonical-570d1d614cef1276f89ada309d1016217011b46549163a38ab09f8cafce0ad53)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_secure](resources--http_loadbalancer--reference--group-025.md#canonical-19434c68f229ed00022e54ea7db35e3b11e9d3c6e619a3e612e693b36d8e8218)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_value](resources--http_loadbalancer--reference--group-025.md#canonical-3c3b61e0a7d25532f5557d33488954b2bd98222e4edc2b876502c80663c4f59e)
- [routes.simple_route.advanced_options.response_cookies_to_add.samesite_lax](resources--http_loadbalancer--reference--group-025.md#canonical-9a41c832843e252991afe87976d397ba773d4314bb035da9ccf540130e5b3d9a)
- [routes.simple_route.advanced_options.response_cookies_to_add.samesite_none](resources--http_loadbalancer--reference--group-025.md#canonical-e724faa2c6f35a235c34b8537260157dcc76b62ea612d67107f35e166ff3e9ae)
- [routes.simple_route.advanced_options.response_cookies_to_add.samesite_strict](resources--http_loadbalancer--reference--group-025.md#canonical-156f1399f25dec1e4d2cb6eb7aa7223513a9c564b210daed8eae93f88f1872b7)
- [routes.simple_route.advanced_options.response_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-025.md#canonical-23987764d6be2d466524adabb157ea49f0b9d54983f821d1db43d1d75e1d9d39)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b732ad7d34c58ee0e651ab59950072b090c2a3ed94694135c9a6c2782b0eeafb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b2e0834015a93a8af205520cf73675666000d76f3cd00ee7c03199a6331982b"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.add_httponly — routes.simple_route.advanced_options.response_cookies_to_add.add_httponly / 820789791b20 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-024.md#canonical-3985b17e1a5096f3db4c6a657ac616194c986fe82e7d9f82e728fcd6ee04bbc8)
- routes.simple_route.advanced_options.response_cookies_to_add.add_httponly

<a id="canonical-678dc77538cfd8f050869834599c207c68e6fa82cfd495e1928322bdb4aed11b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add httponly.

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
add_httponly = {}
```

<a id="canonical-7a85cb82bc695e2f64e505a832d1e46cb3560f5165c150c592cd66364c921b04"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.add_httponly / 820789791b20 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7f6cfad1509427d2a1f1e9fd95bcb02ea21dc0e0973d7c84b04a5ca8ee0bee8b"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add.add_httponly / 820789791b20 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-024.md#canonical-3985b17e1a5096f3db4c6a657ac616194c986fe82e7d9f82e728fcd6ee04bbc8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7bb8622d7bb8ffd00a83326f5e84bc4a754308e24e861a28a67720c3b218575b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd86eafecce353d774d465fa516e5383b21142ee3f60ef8a4ef05c19ede6bf7c"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.add_partitioned — routes.simple_route.advanced_options.response_cookies_to_add.add_partitioned / 307088da0ef8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [routes](resources--http_loadbalancer--reference--group-023.md#canonical-541a918f6992f609cb9368685dbf07fc29f8bb8b9a72694ced412793a3b8fb50)
- [routes.simple_route](resources--http_loadbalancer--reference--group-023.md#canonical-5cdf54c5ccf670c09568c2112380c4aa4707a118ff5be079f79050cc6160dbee)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-cf3d32226eac6b9d59c5658cbe133ccad12b2f083881699d5c4eb63abfcb1d35)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-024.md#canonical-3985b17e1a5096f3db4c6a657ac616194c986fe82e7d9f82e728fcd6ee04bbc8)
- routes.simple_route.advanced_options.response_cookies_to_add.add_partitioned

<a id="canonical-3d13e2d06e5c18a734a39071d90c3b35575869e9bd2580cfcbb88ca6d903829c"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add partitioned.

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
add_partitioned = {}
```

<a id="canonical-59e56bff51be282f3d6abc043dbd9391133260e3d7aa3f6b1fb3b66717681463"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.add_partitioned / 307088da0ef8 / 3

This is an empty object or choice marker. It has no direct properties.
