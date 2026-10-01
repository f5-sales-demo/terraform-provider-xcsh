---
page_title: "xcsh_gcp_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site reference."
---

# xcsh_gcp_vpc_site reference

<a id="canonical-a8e4ad766d17c1897854b6e02d7e1a12b3adfc09544b43fab65eadf575bce24e"></a>

## volterra_software_version property — sw / 761d5743321b / 4

Type: `"string"`. Optional.

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Upstream description:

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-feb055ae7dd9058bb3afc71bf0ddaa656ce028ae6c5826f996bb401021997e69"></a>

## Next pages — sw / 761d5743321b / 5

- [sw.default_sw_version](resources--gcp_vpc_site--reference--group-004.md#canonical-84b59a58878fd650e747bc25e45272cdd659e0b6d53b9d1e4d65fb45118fbf52)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-84b59a58878fd650e747bc25e45272cdd659e0b6d53b9d1e4d65fb45118fbf52"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33eaa41c64f8ff08e47bb9046277ef5c84477827fc718134a3947794f6eee79c"></a>

## sw.default_sw_version — sw.default_sw_version / 9ad305ff5030 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [sw](resources--gcp_vpc_site--reference--group-003.md#canonical-267f967c717532fd9b7a7d530db900ad6f3ed5285284fb0cc81876b19418dec2)
- sw.default_sw_version

<a id="canonical-4acb0e49762c1b77f784b71c216bf0747e5d5bf71c9fb49e802f14a4b79a5ae2"></a>

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
default_sw_version = {}
```

<a id="canonical-f5addd1b388e157dbd6658667c3fecdcdb4ddb49e223cfa4ae4719e4eb5303dd"></a>

## Direct properties — sw.default_sw_version / 9ad305ff5030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-849461f9defdbe9aec1ac1f00047b1bd1cba141f5cf5f7025fbe28661d8a6c3d"></a>

## Next pages — sw.default_sw_version / 9ad305ff5030 / 4

- [sw](resources--gcp_vpc_site--reference--group-003.md#canonical-267f967c717532fd9b7a7d530db900ad6f3ed5285284fb0cc81876b19418dec2)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-2298b57a32aa2303d946d8a0a735921119bf066b4e9cbcf62828b9ca2db0aad9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-376c342af370ccc194ccde26b3d4247dc615270616536dec3f0f81265d556c84"></a>

## timeouts — timeouts / eb1f25b2db68 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- timeouts

<a id="canonical-8eae28ef2cad76a6257738eb65a3ced02f012a7344774ac5d6165414e4c0739e"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-bba1674aab43a1055c6373ff8f24e3697f295871e77793b90f2572fb03565495"></a>

## Direct properties — timeouts / eb1f25b2db68 / 3

<a id="canonical-38844ebfad91166c0a71af7cac77e0d5189b23093dfca2397684076ffb99decf"></a>

<a id="canonical-d93a4fd2aefeb2f2e96477d7b50c37288cb0d55f9b2ee36ae60e31156096e0e7"></a>

## create property — timeouts / eb1f25b2db68 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-c85e6158c1df5fd539adb6f62f1e6ca9279420478642c2580b68abf4beb6380b"></a>

<a id="canonical-507ffc8b4d92cb27c6fb903cf4803341ed7713facf14e4f26c6586defc1a5117"></a>

## delete property — timeouts / eb1f25b2db68 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-36f4b47c9f3598d4fc1327d5795881fa5eec2867966e59f5c023049b21fa9cf4"></a>

<a id="canonical-54bf2658ee65fc1576136f9bba6cc17fd8a8dd407661d14d2e7ab3b699a42687"></a>

## read property — timeouts / eb1f25b2db68 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1f9c8e19946da3aa6cd54f2ebe2adbd18bf95e97a4f86716025f54e4be5f2066"></a>

<a id="canonical-ed70b4007866340253c32ec6810a82731ac49f6e49dff11f05ce93b3cd80c9d8"></a>

## update property — timeouts / eb1f25b2db68 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-a646710daea0cacb4e29d8bde25eea3022e4335ad534d65fcb9e582400f25cd2"></a>

## Next pages — timeouts / eb1f25b2db68 / 8

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-423f3d67d0a6a63daf5e1f9f30ba591d43cbf948195a0c3c0e5708b1274f325e"></a>

## voltstack_cluster — voltstack_cluster / bfea3f9e7117 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- voltstack_cluster

<a id="canonical-760f6ba6855f451d0477a5feddf0087ff8d9892bce89b49960bd8724d3ce6eba"></a>

Type: `"object"`. single nested block, Optional.

App Stack cluster of single interface GCP site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("gcp_certified_hw",
    "gcp_zone_names"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "active_network_policies"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "forward_proxy_allow_all"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("active_network_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("dc_cluster_group",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("default_storage",
    "storage_class_list"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
  validators.ConflictingObjectAttributes("k8s_cluster",
    "no_k8s_cluster"),
  validators.ConflictingObjectAttributes("no_outside_static_routes",
    "outside_static_routes"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
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
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-k8s_cluster_choice": "[\"k8s_cluster\",\"no_k8s_cluster\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]",
  "x-ves-oneof-field-storage_class_choice": "[\"default_storage\",\"storage_class_list\"]"
}
```

Terraform syntax:

```terraform
voltstack_cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-c45f7248c363b8f625d5093f226db3ab0e6c5222555f2e797482715b11878255"></a>

## Direct properties — voltstack_cluster / bfea3f9e7117 / 3

- [active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-073a4b64dcbaf7db8c63a3f055811f46d57a44a780003d15d7b079253e819370): complete subsection reference.

- [active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-f6544c431c2ce1fe8ae7b46159b1fd0860071dbb78ea4a9435d14e47a6755e8b): complete subsection reference.

- [active_network_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-19771103bad834902c434d806b8a57448d4454866dea11cb9120b35b3178dc74): complete subsection reference.

- [dc_cluster_group](resources--gcp_vpc_site--reference--group-004.md#canonical-5c542b8e588b90bd8383cd85037e27de6993feb1423058de1ed14a94c74702d1): complete subsection reference.

- [default_storage](resources--gcp_vpc_site--reference--group-004.md#canonical-4ebc746a187e264dea51f81121997d660ad87e85e543fa545311c932fda7348b): complete subsection reference.

- [forward_proxy_allow_all](resources--gcp_vpc_site--reference--group-004.md#canonical-b548d25632f2323044b216f69dacc3e786b6a774b5d0d5a9a0dabe87ab97961c): complete subsection reference.

<a id="canonical-6c38174d6557a61d5381f449c0bf0c73daf0c4b2908b862cd3b087c068c3da20"></a>

<a id="canonical-829188fbc31a2b81aef2d6bac1b98f08f9bb06aeb57b6302ca2bc131d5c206aa"></a>

## gcp_certified_hw property — voltstack_cluster / bfea3f9e7117 / 4

Type: `"string"`. Optional.

\[Enum: gcp-byol-voltstack-combo\] GCP Certified Hardware. Name for GCP certified hardware. The only
possible value is \`gcp-byol-voltstack-combo\`.

Upstream description:

Name for GCP certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("gcp-byol-voltstack-combo"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "gcp-byol-voltstack-combo"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-8e2e4d67b9eff80eeb9025f1b49b8b9f0b000de3d5839f950e0e8b5a5ac88e41"></a>

<a id="canonical-7a311b245b54ea1f1b3d2f546732cffd7cf5d1f2fb206027af504b96dbcae930"></a>

## gcp_zone_names property — voltstack_cluster / bfea3f9e7117 / 5

Type: `["list", "string"]`. Optional.

X-required List of zones when instances will be created, needs to match with region selected.

Upstream description:

X-required List of zones when instances will be created, needs to match with region selected.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(3),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [global_network_list](resources--gcp_vpc_site--reference--group-004.md#canonical-7f7b911b7737de6d66f5805b10981b5da71f85fae44df698f77fc3d79ccdc908): complete subsection reference.

- [k8s_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-6ab5dfec17019798407b97ffe626c7deaa095d98081ae3cbba570ceb068204ae): complete subsection reference.

- [no_dc_cluster_group](resources--gcp_vpc_site--reference--group-004.md#canonical-4176f808cef5388ec183126f11b664dbf0b4f8be9686450dfa5f0a0c5f9f5c6a): complete subsection reference.

- [no_forward_proxy](resources--gcp_vpc_site--reference--group-004.md#canonical-d0e8998e48affcaa5ac09252943e3c30a4136b272ff3c5b5d1f8cf91e1010cfe): complete subsection reference.

- [no_global_network](resources--gcp_vpc_site--reference--group-004.md#canonical-b5541b65805f5fcb7e007ec8a55ff04911dfc2a02ba9f85b1aa3061d3d57aa58): complete subsection reference.

- [no_k8s_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-e8176eadfa957c1e92e1aaf5996293989efcf098de8ccda58961c6c6a0648228): complete subsection reference.

- [no_network_policy](resources--gcp_vpc_site--reference--group-004.md#canonical-193ae660e21c2dc90147e84bcb3396af6e2d7d7c219241ac7e1e7254ad55b4f7): complete subsection reference.

- [no_outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-e96b9e9e1b31efdc5030732c65d2691e3a5c5f1a3d28565c89c5e805a36374d9): complete subsection reference.

<a id="canonical-0c4eac7ef1c52788f352810d377ccbc461e7c242b0894f80d41868c08042e311"></a>

<a id="canonical-0e0683418077699790ed68ad5b5db5cb72a4e02065b0cd54fb704dfb01049627"></a>

## node_number property — voltstack_cluster / bfea3f9e7117 / 6

Type: `"number"`. Optional.

Number of main nodes to create, either 1 or 3.

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
    "ves.io.schema.rules.uint32.in": "[1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.in": "[1,3]"
  }
}
```

- [outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-5c3c562c7707d83d434086992a2febf9d66463cf41d247a396db55a64f80c18a): complete subsection reference.

- [site_local_network](resources--gcp_vpc_site--reference--group-004.md#canonical-0d117f633e2656829fba928f7b9b2a047ea066556f60ce1d03c0e04f47f29fa0): complete subsection reference.

- [site_local_subnet](resources--gcp_vpc_site--reference--group-004.md#canonical-447430b7a7fb6788106f2e09e8f5fd7089adab0de64c47852e87c1abcc56deaf): complete subsection reference.

- [sm_connection_public_ip](resources--gcp_vpc_site--reference--group-004.md#canonical-94dfc3b3d4fea661a8cea3b72218841e075fbd02a96681d4d90e3f4604ba3e4a): complete subsection reference.

- [sm_connection_pvt_ip](resources--gcp_vpc_site--reference--group-004.md#canonical-3d45010f2a154d078ec9f6f6287c97b091b78fc80801589773e9e5cc9e1ebbf8): complete subsection reference.

- [storage_class_list](resources--gcp_vpc_site--reference--group-004.md#canonical-56ffa539204782f5f466a7bc0f01f7b06e0be7bd95c35b7f6529e6d577be27d9): complete subsection reference.

<a id="canonical-ca763c391b0bb2550f3940ea73ee9576e410c5d51b11fc068362e8709d1dd705"></a>

## Next pages — voltstack_cluster / bfea3f9e7117 / 7

- [voltstack_cluster.active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-073a4b64dcbaf7db8c63a3f055811f46d57a44a780003d15d7b079253e819370)
- [voltstack_cluster.active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-f6544c431c2ce1fe8ae7b46159b1fd0860071dbb78ea4a9435d14e47a6755e8b)
- [voltstack_cluster.active_network_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-19771103bad834902c434d806b8a57448d4454866dea11cb9120b35b3178dc74)
- [voltstack_cluster.dc_cluster_group](resources--gcp_vpc_site--reference--group-004.md#canonical-5c542b8e588b90bd8383cd85037e27de6993feb1423058de1ed14a94c74702d1)
- [voltstack_cluster.default_storage](resources--gcp_vpc_site--reference--group-004.md#canonical-4ebc746a187e264dea51f81121997d660ad87e85e543fa545311c932fda7348b)
- [voltstack_cluster.forward_proxy_allow_all](resources--gcp_vpc_site--reference--group-004.md#canonical-b548d25632f2323044b216f69dacc3e786b6a774b5d0d5a9a0dabe87ab97961c)
- [voltstack_cluster.global_network_list](resources--gcp_vpc_site--reference--group-004.md#canonical-7f7b911b7737de6d66f5805b10981b5da71f85fae44df698f77fc3d79ccdc908)
- [voltstack_cluster.k8s_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-6ab5dfec17019798407b97ffe626c7deaa095d98081ae3cbba570ceb068204ae)
- [voltstack_cluster.no_dc_cluster_group](resources--gcp_vpc_site--reference--group-004.md#canonical-4176f808cef5388ec183126f11b664dbf0b4f8be9686450dfa5f0a0c5f9f5c6a)
- [voltstack_cluster.no_forward_proxy](resources--gcp_vpc_site--reference--group-004.md#canonical-d0e8998e48affcaa5ac09252943e3c30a4136b272ff3c5b5d1f8cf91e1010cfe)
- [voltstack_cluster.no_global_network](resources--gcp_vpc_site--reference--group-004.md#canonical-b5541b65805f5fcb7e007ec8a55ff04911dfc2a02ba9f85b1aa3061d3d57aa58)
- [voltstack_cluster.no_k8s_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-e8176eadfa957c1e92e1aaf5996293989efcf098de8ccda58961c6c6a0648228)
- [voltstack_cluster.no_network_policy](resources--gcp_vpc_site--reference--group-004.md#canonical-193ae660e21c2dc90147e84bcb3396af6e2d7d7c219241ac7e1e7254ad55b4f7)
- [voltstack_cluster.no_outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-e96b9e9e1b31efdc5030732c65d2691e3a5c5f1a3d28565c89c5e805a36374d9)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-5c3c562c7707d83d434086992a2febf9d66463cf41d247a396db55a64f80c18a)
- [voltstack_cluster.site_local_network](resources--gcp_vpc_site--reference--group-004.md#canonical-0d117f633e2656829fba928f7b9b2a047ea066556f60ce1d03c0e04f47f29fa0)
- [voltstack_cluster.site_local_subnet](resources--gcp_vpc_site--reference--group-004.md#canonical-447430b7a7fb6788106f2e09e8f5fd7089adab0de64c47852e87c1abcc56deaf)
- [voltstack_cluster.sm_connection_public_ip](resources--gcp_vpc_site--reference--group-004.md#canonical-94dfc3b3d4fea661a8cea3b72218841e075fbd02a96681d4d90e3f4604ba3e4a)
- [voltstack_cluster.sm_connection_pvt_ip](resources--gcp_vpc_site--reference--group-004.md#canonical-3d45010f2a154d078ec9f6f6287c97b091b78fc80801589773e9e5cc9e1ebbf8)
- [voltstack_cluster.storage_class_list](resources--gcp_vpc_site--reference--group-004.md#canonical-56ffa539204782f5f466a7bc0f01f7b06e0be7bd95c35b7f6529e6d577be27d9)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-073a4b64dcbaf7db8c63a3f055811f46d57a44a780003d15d7b079253e819370"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9094e6ff314eabd220de21ae602b777ed2a0ae06817412bd6e5be242f06289c1"></a>

## voltstack_cluster.active_enhanced_firewall_policies — voltstack_cluster.active_enhanced_firewall_policies / 7dba12012b98 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.active_enhanced_firewall_policies

<a id="canonical-3ba32b782e67df3af8ceede722bd2d1b7bf8edd4a8ee221fd71b218666222fff"></a>

Type: `"object"`. single nested block, Optional.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("enhanced_firewall_policies")}
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
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2b8eaf0bd8be6bd27d86a372cb10f469a0d2eb0404a2506760a722d1d65946ca"></a>

## Direct properties — voltstack_cluster.active_enhanced_firewall_policies / 7dba12012b98 / 3

- [enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-3d556a2975c3698cf5fb0348a81797955225c8cf50c08d79fd1932c94f5046f6): complete subsection reference.

<a id="canonical-3fbc8ea14dc12f1f71d75fbc792e38042444c6f04951fbf442f5e0f14e7a111c"></a>

## Next pages — voltstack_cluster.active_enhanced_firewall_policies / 7dba12012b98 / 4

- [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-3d556a2975c3698cf5fb0348a81797955225c8cf50c08d79fd1932c94f5046f6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-3d556a2975c3698cf5fb0348a81797955225c8cf50c08d79fd1932c94f5046f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13b720684bd2d19461c4d2b81e09cd1a754ab487551b4ae5f3a0884fa33147e1"></a>

## voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies — voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies / 9d473335130d / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-073a4b64dcbaf7db8c63a3f055811f46d57a44a780003d15d7b079253e819370)
- voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-1a67cca6960532ec5d4ab6d48bfe4e6f84b1f4b8b73099ef459c693a20ebf3b7"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Enhanced Firewall Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-ea197894702ed155c6031d3c07614e554ad78e7eae9133901f31fe12b1f74a4c"></a>

## Direct properties — voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies / 9d473335130d / 3

<a id="canonical-cebb3702efb5f1df66128a68cd662390c387ebd37f7ddd6b6ce461294785bed0"></a>

<a id="canonical-d971b42bf8f6655231af3c777b547ecf87b6a081b40dd638f8e737891a661f70"></a>

## name property — voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies / 9d473335130d / 4

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

<a id="canonical-5922b67a7d2422b90b1fc35f191fc6777793aa1beb4d934f3cdc26d89c6c5383"></a>

<a id="canonical-f5cb23f685d526b81e3e6ee911e2bb11faee9579cd59a8b4dde77ab9a1165818"></a>

## namespace property — voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies / 9d473335130d / 5

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

<a id="canonical-2a953ea73d73fef189a305d13d645d0c16984f98cbb1100d09e13232b4dba11b"></a>

<a id="canonical-081ab9eb43ba28cadcaf76b2e4c0d0f8159599538aa83743a6ca65f85b07f2d3"></a>

## tenant property — voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies / 9d473335130d / 6

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

<a id="canonical-c64feb3201fb22502696614ac2cb927db89a024f386ceeb28e9b6eba1a640592"></a>

## Next pages — voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies / 9d473335130d / 7

- [voltstack_cluster.active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-073a4b64dcbaf7db8c63a3f055811f46d57a44a780003d15d7b079253e819370)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-f6544c431c2ce1fe8ae7b46159b1fd0860071dbb78ea4a9435d14e47a6755e8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4f73386889872d2e1e121fd5bfa8d54ad9e8fdc572fc2241c973b958584b5c7"></a>

## voltstack_cluster.active_forward_proxy_policies — voltstack_cluster.active_forward_proxy_policies / fcb9fbe25992 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.active_forward_proxy_policies

<a id="canonical-35170f83f20c897ab9ec9641b9d6859a2bfc13681ffbe5368a23bf03a607adbd"></a>

Type: `"object"`. single nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-7d1913511345d2062ed9662eb8e0ba73008dbdf5181ef51b81c7d740272eeed0"></a>

## Direct properties — voltstack_cluster.active_forward_proxy_policies / fcb9fbe25992 / 3

- [forward_proxy_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-be7a63ba87c3a9c0977e53146dc6332539168144c501f4a377a0b82e74b6a406): complete subsection reference.

<a id="canonical-49a0b5b433af3cd547799a67b1a304478ee59f2309013137ac551b763f6904d3"></a>

## Next pages — voltstack_cluster.active_forward_proxy_policies / fcb9fbe25992 / 4

- [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-be7a63ba87c3a9c0977e53146dc6332539168144c501f4a377a0b82e74b6a406)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-be7a63ba87c3a9c0977e53146dc6332539168144c501f4a377a0b82e74b6a406"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e06363afe8766df16e31bbac53dbb432de2d812fb8f5e0d06b9e4c29c55aafa"></a>

## voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies — voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies / 3b1914c05417 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-f6544c431c2ce1fe8ae7b46159b1fd0860071dbb78ea4a9435d14e47a6755e8b)
- voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-671bf46083feb34481edd240361521c0a8f083d9b3122bac6b213c70a877e3e1"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-57d07d5877893167b9c126f2b88c33c7376a1bb92ce93021d829c398988b0089"></a>

## Direct properties — voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies / 3b1914c05417 / 3

<a id="canonical-174e8ac715b4ffdc7cab62329df4226baae2ecb73c1781519868edef9cd49c2d"></a>

<a id="canonical-bd059fd2e0a0b49c4337608fd23249d5795dd983874ef6bcf3424e4b5fa698a2"></a>

## name property — voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies / 3b1914c05417 / 4

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

<a id="canonical-c90eb76da83b43d7df895a68588d5b16e330344cda15fc95b1233bba7d67687f"></a>

<a id="canonical-ddd7c4e87a4f1825ddd84504290546f9947e007beb435651d9451bc34b069e57"></a>

## namespace property — voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies / 3b1914c05417 / 5

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

<a id="canonical-7e2525226c04fae220224868ec0f36dd59c8ca49494ba2e2b5fbda193cec0505"></a>

<a id="canonical-c26739bd6080954a41a23eba633942c7ac79ed54293d0285279f270c598876c4"></a>

## tenant property — voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies / 3b1914c05417 / 6

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

<a id="canonical-e25a284fa121d73ee888b55a02437221666f8361105724d743efcd09ec42b9bf"></a>

## Next pages — voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies / 3b1914c05417 / 7

- [voltstack_cluster.active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-f6544c431c2ce1fe8ae7b46159b1fd0860071dbb78ea4a9435d14e47a6755e8b)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-19771103bad834902c434d806b8a57448d4454866dea11cb9120b35b3178dc74"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f09e86425a2ca5bff459bfef6308b90e52b7f5f3c026f92c3c849324fe246a20"></a>

## voltstack_cluster.active_network_policies — voltstack_cluster.active_network_policies / 0ed5ebb995b9 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.active_network_policies

<a id="canonical-8a739ec992939e6c8cd2604af0fbc3ad2158dff1d2015788af793cce8418fdc9"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
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
active_network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-8aed3d2e36af94d1d928d4b30e0cec1f55fe5cdea02b6a71a6cbbd716532426f"></a>

## Direct properties — voltstack_cluster.active_network_policies / 0ed5ebb995b9 / 3

- [network_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-c99e784277b5b90d220bf1ee65f9305d4cfed8dcd0f38640d6d9411b45ad436a): complete subsection reference.

<a id="canonical-c7ac06e3fdd3c4c5f3910c3191c26224c2882f215512b72a3e9b0d0a17630749"></a>

## Next pages — voltstack_cluster.active_network_policies / 0ed5ebb995b9 / 4

- [voltstack_cluster.active_network_policies.network_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-c99e784277b5b90d220bf1ee65f9305d4cfed8dcd0f38640d6d9411b45ad436a)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-c99e784277b5b90d220bf1ee65f9305d4cfed8dcd0f38640d6d9411b45ad436a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-336d9286df0bb837781093d22d5c3e3ff4f22f7fbd2f4b689ac0205cbd143864"></a>

## voltstack_cluster.active_network_policies.network_policies — voltstack_cluster.active_network_policies.network_policies / a61759045236 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.active_network_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-19771103bad834902c434d806b8a57448d4454866dea11cb9120b35b3178dc74)
- voltstack_cluster.active_network_policies.network_policies

<a id="canonical-99640778439a9a333ff988ab515ec375f576d96d03c220596a26ef11e1882c16"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Firewall Policies active for this network firewall.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-373f56bf704a18f11e9983426a68a7ef6daf31c682d8e2c7d597d0f41bd859ba"></a>

## Direct properties — voltstack_cluster.active_network_policies.network_policies / a61759045236 / 3

<a id="canonical-61b2be690f775792f4571c7d5a35917cbedd15cbc5eee850cace31f4b694192d"></a>

<a id="canonical-bb31eb06fbedd86b2381b7f143f36ba4c43c7e2302a5e5e66c31d5b683d5f237"></a>

## name property — voltstack_cluster.active_network_policies.network_policies / a61759045236 / 4

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

<a id="canonical-05a8c9ef0c2c963dd798de691c74a7d513a59f60ba740175fa8bc7cf71ba2af2"></a>

<a id="canonical-da685f01625d358dbd253d4102a4aba9af9ad1236328a3cf1897a7427965bfaf"></a>

## namespace property — voltstack_cluster.active_network_policies.network_policies / a61759045236 / 5

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

<a id="canonical-7783eb4daf25d95e089ab6e238146410e4b314c1aec8cd4fac079a29a068022d"></a>

<a id="canonical-b5775a9df35f9cf1a3dcf4386e2bc611914688b78712b07e1186aa667886c016"></a>

## tenant property — voltstack_cluster.active_network_policies.network_policies / a61759045236 / 6

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

<a id="canonical-b08cca9f15b2d17c4e23d65b55df91bb03c7ea31c1434b4b23460952ff23c2c6"></a>

## Next pages — voltstack_cluster.active_network_policies.network_policies / a61759045236 / 7

- [voltstack_cluster.active_network_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-19771103bad834902c434d806b8a57448d4454866dea11cb9120b35b3178dc74)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-5c542b8e588b90bd8383cd85037e27de6993feb1423058de1ed14a94c74702d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5adc2e48e212b567505f8d02643f0ec10ca40830233740770a2debd761d2b3e9"></a>

## voltstack_cluster.dc_cluster_group — voltstack_cluster.dc_cluster_group / c3ad50db2d44 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.dc_cluster_group

<a id="canonical-45928c72e712d901c584aaec9317c761f2f2a2c40956d4abb305e748d2163a4c"></a>

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
dc_cluster_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-2a69a598fb1d858a3211ab290ea6646cb668909dac115d9eb5bd1518bd8b07f2"></a>

## Direct properties — voltstack_cluster.dc_cluster_group / c3ad50db2d44 / 3

<a id="canonical-7f91495e7cbcc5464b4388edc0361872933a1fb5d354d14b1d6dd0114bea51fc"></a>

<a id="canonical-3722d8b61cdbdf5a7f0739b1cf9fa44936c2897b3c47340548a5ffd8605e74c8"></a>

## name property — voltstack_cluster.dc_cluster_group / c3ad50db2d44 / 4

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

<a id="canonical-5178170977e9133e4edcb137f2f1d560c56a8a395b7dafc0efa032f3e92c8140"></a>

<a id="canonical-b302efc0a9166708f7d473bbe0561f563d35614e091f767ad299ce6677831487"></a>

## namespace property — voltstack_cluster.dc_cluster_group / c3ad50db2d44 / 5

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

<a id="canonical-7325fb7aa15c1f2beb8fcafc01348ccc9c6bc37e4082e2268a4ccd242ed661a3"></a>

<a id="canonical-62f9ab41ba542a394089061e1a2cdf41cf28ec9f6f7566a30e6e33bdbf1d3f27"></a>

## tenant property — voltstack_cluster.dc_cluster_group / c3ad50db2d44 / 6

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

<a id="canonical-594916ca904bdb8b6898c153dea74fc6d3e39cc07f41ead2a53c59288e8e73a1"></a>

## Next pages — voltstack_cluster.dc_cluster_group / c3ad50db2d44 / 7

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-4ebc746a187e264dea51f81121997d660ad87e85e543fa545311c932fda7348b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f4b229f592a9a615f413906a09d114a99cf492a6b48814c6559e3120de44355"></a>

## voltstack_cluster.default_storage — voltstack_cluster.default_storage / 7b5c3de00712 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.default_storage

<a id="canonical-e4934706ddb00479eb420cd5e501e6e5077f0ef36c2fc301c5a2e752f3d01d17"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default storage.

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
default_storage = {}
```

<a id="canonical-9080d39a05ad19bbedb9d382e287df5e0f78eb15facc24b25656820d82d1574a"></a>

## Direct properties — voltstack_cluster.default_storage / 7b5c3de00712 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4a44760b1b105c0134964c8786888e6aa3a7f4af3bca4572a832dd71d5f04e3d"></a>

## Next pages — voltstack_cluster.default_storage / 7b5c3de00712 / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-b548d25632f2323044b216f69dacc3e786b6a774b5d0d5a9a0dabe87ab97961c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e38c759a99ad69d278eb09ced12c6a03f5e10ac9154c426db34ef0f430f7dcd"></a>

## voltstack_cluster.forward_proxy_allow_all — voltstack_cluster.forward_proxy_allow_all / 9c59f6f41a0e / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.forward_proxy_allow_all

<a id="canonical-b270bff65d2b871fd504af6f6ad2e59a8f63ba12959bd14b00a0643baa5f8ef2"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for forward proxy allow all.

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
forward_proxy_allow_all = {}
```

<a id="canonical-1ce68880bc7ab62fc054672ff054535e1bfc563085259cbdd35fe9b8d5ba8050"></a>

## Direct properties — voltstack_cluster.forward_proxy_allow_all / 9c59f6f41a0e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-15868ed8c224336cb02a30265aa17f8e324b94c4cdc5c67a18f47f8ae9053e7e"></a>

## Next pages — voltstack_cluster.forward_proxy_allow_all / 9c59f6f41a0e / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-7f7b911b7737de6d66f5805b10981b5da71f85fae44df698f77fc3d79ccdc908"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4af95ebc6b34edc55001ee9b6ef90f7724eaf62896612d99a03c8381864f03b"></a>

## voltstack_cluster.global_network_list — voltstack_cluster.global_network_list / 8868028f15d5 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.global_network_list

<a id="canonical-250995e2a50a5aedccc225c32904a6ff53936fb9bbe7d81d29f7437d0b45d614"></a>

Type: `"object"`. single nested block, Optional.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("global_network_connections")}
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
global_network_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-eda6d9f6b824dce45097e26c3f31251cd53f8e211db23e3152660304bfd2ea16"></a>

## Direct properties — voltstack_cluster.global_network_list / 8868028f15d5 / 3

- [global_network_connections](resources--gcp_vpc_site--reference--group-004.md#canonical-3fe609189efa5a9ade4fedbc3f0522eba9d315551337bb3c00058bf55b3c58eb): complete subsection reference.

<a id="canonical-6bf056d4c68b03b577f156822b99812a8fd5829ef407c0625417c052f121230c"></a>

## Next pages — voltstack_cluster.global_network_list / 8868028f15d5 / 4

- [voltstack_cluster.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-004.md#canonical-3fe609189efa5a9ade4fedbc3f0522eba9d315551337bb3c00058bf55b3c58eb)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-3fe609189efa5a9ade4fedbc3f0522eba9d315551337bb3c00058bf55b3c58eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ea92f8514ddbacc43fbd8cfe61b9bb1344be9364802e15c42d82db31b14a073"></a>

## voltstack_cluster.global_network_list.global_network_connections — voltstack_cluster.global_network_list.global_network_connections / f47fe18e138e / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.global_network_list](resources--gcp_vpc_site--reference--group-004.md#canonical-7f7b911b7737de6d66f5805b10981b5da71f85fae44df698f77fc3d79ccdc908)
- voltstack_cluster.global_network_list.global_network_connections

<a id="canonical-5ddafdefdd62543430af79cd3df7b29849e7c7dcd3bea0d56632a643ea288353"></a>

Type: `"object"`. list nested block, Optional.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("sli_to_global_dr",
    "slo_to_global_dr")}
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
global_network_connections {
  # Configure direct properties listed below.
}
```

<a id="canonical-88c4bc6f2d50a99a219cc2712af79d39daec569d9fef7da929700817ad3472d7"></a>

## Direct properties — voltstack_cluster.global_network_list.global_network_connections / f47fe18e138e / 3

- [sli_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-6b260537ce296e32dee12e38503d30fc5bb1a4eb2f1306b6c274accfd09575b8): complete subsection reference.

- [slo_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-9a30171c28b1cec3eead5c30b2c8cbc2add7ccb91703d194286fe574dc2cb35d): complete subsection reference.

<a id="canonical-582b01e013da9003b40c27487f62632089a629e1aa917f9eadbcda5aeae6617d"></a>

## Next pages — voltstack_cluster.global_network_list.global_network_connections / f47fe18e138e / 4

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-6b260537ce296e32dee12e38503d30fc5bb1a4eb2f1306b6c274accfd09575b8)
- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-9a30171c28b1cec3eead5c30b2c8cbc2add7ccb91703d194286fe574dc2cb35d)
- [voltstack_cluster.global_network_list](resources--gcp_vpc_site--reference--group-004.md#canonical-7f7b911b7737de6d66f5805b10981b5da71f85fae44df698f77fc3d79ccdc908)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-6b260537ce296e32dee12e38503d30fc5bb1a4eb2f1306b6c274accfd09575b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b948d42a6eed8cdcd03b5139d8485b48022f5cce59d9187c302417d26c3a5b0"></a>

## voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / ba09d2d638df / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.global_network_list](resources--gcp_vpc_site--reference--group-004.md#canonical-7f7b911b7737de6d66f5805b10981b5da71f85fae44df698f77fc3d79ccdc908)
- [voltstack_cluster.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-004.md#canonical-3fe609189efa5a9ade4fedbc3f0522eba9d315551337bb3c00058bf55b3c58eb)
- voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-291b72062cfee93fe587afb128866ef7fa87c0487de63dc04006707b4a34112b"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
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
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-41af80c02265c2f220ad8917c4ef4f2db79b2b33eb05ae0e898898e8aac13e10"></a>

## Direct properties — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / ba09d2d638df / 3

- [global_vn](resources--gcp_vpc_site--reference--group-004.md#canonical-97a261aa1836bd16a45c4792e077c33b0df165f8476e436919073a08f451a04c): complete subsection reference.

<a id="canonical-2332829106a52af4ced5a61d83a576bc8a2365178afc0c16cfabd1b44572728a"></a>

## Next pages — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / ba09d2d638df / 4

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--gcp_vpc_site--reference--group-004.md#canonical-97a261aa1836bd16a45c4792e077c33b0df165f8476e436919073a08f451a04c)
- [voltstack_cluster.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-004.md#canonical-3fe609189efa5a9ade4fedbc3f0522eba9d315551337bb3c00058bf55b3c58eb)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-97a261aa1836bd16a45c4792e077c33b0df165f8476e436919073a08f451a04c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2a610008901c9f38065f39c1bf8fca2446dc8a259ad4ce1ea20495b81f78ed0"></a>

## voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / b34e67259d81 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.global_network_list](resources--gcp_vpc_site--reference--group-004.md#canonical-7f7b911b7737de6d66f5805b10981b5da71f85fae44df698f77fc3d79ccdc908)
- [voltstack_cluster.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-004.md#canonical-3fe609189efa5a9ade4fedbc3f0522eba9d315551337bb3c00058bf55b3c58eb)
- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-6b260537ce296e32dee12e38503d30fc5bb1a4eb2f1306b6c274accfd09575b8)
- voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-c408f83aa1811ce927b3ac0faa3fb06509244675145ca969cd7dd3d9a853b9b6"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-d0cfe5b55b4e71e84734cf017e17f86c66661a8a8e9a0f1c6fadaa5717568158"></a>

## Direct properties — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / b34e67259d81 / 3

<a id="canonical-02d612711a27b6e76e946f7c59fbb048f2837082575b8bdc7efecaa6165a696b"></a>

<a id="canonical-daaeaf9ceb7158d26ff1ea2a30e936e625151673afbf018edf5be37a82a3c4cd"></a>

## name property — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / b34e67259d81 / 4

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

<a id="canonical-c145ee1ea8f3c3c83c70c04b893bd7950f0fdb1c03d39d29abbe949f5de999a3"></a>

<a id="canonical-c3c86362f28e39b8af24972f6ea474933351149de5324e191d6da1ba107ebb2b"></a>

## namespace property — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / b34e67259d81 / 5

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

<a id="canonical-5632a2fde54683908df229083bb7f9a3ec627c2c966cf90dfa1a12571a90d859"></a>

<a id="canonical-aef365497325a6991f74d3a77dce5165d641f3a8358fd414432c20bf95275a46"></a>

## tenant property — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / b34e67259d81 / 6

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

<a id="canonical-1f016f7ced4e07e528ad2de817e39deaab0a71157d45d1d1b712381e401142cd"></a>

## Next pages — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / b34e67259d81 / 7

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-6b260537ce296e32dee12e38503d30fc5bb1a4eb2f1306b6c274accfd09575b8)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-9a30171c28b1cec3eead5c30b2c8cbc2add7ccb91703d194286fe574dc2cb35d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-56fe665901ead3b2677ec8ce4d0db8eaeeba28697d3733e9cbe13f91d4193b8b"></a>

## voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / 70b12d968851 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.global_network_list](resources--gcp_vpc_site--reference--group-004.md#canonical-7f7b911b7737de6d66f5805b10981b5da71f85fae44df698f77fc3d79ccdc908)
- [voltstack_cluster.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-004.md#canonical-3fe609189efa5a9ade4fedbc3f0522eba9d315551337bb3c00058bf55b3c58eb)
- voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-d004df18e13d13da8115ee8d893608080f099aa9d2fe372edfb60c4ce2fb4f64"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
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
slo_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-098522aff2405c430f5b10272bca8ccd6d10b4d32eabcc701663da65ea2b4d48"></a>

## Direct properties — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / 70b12d968851 / 3

- [global_vn](resources--gcp_vpc_site--reference--group-004.md#canonical-8310edad4caa3080989a1a714257510f6e17d6f53bd38b5c027385223239289c): complete subsection reference.

<a id="canonical-6767568ac04741178a9997ff0647d7410cd47d4382dd9b37ed89a2759a84de77"></a>

## Next pages — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / 70b12d968851 / 4

- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--gcp_vpc_site--reference--group-004.md#canonical-8310edad4caa3080989a1a714257510f6e17d6f53bd38b5c027385223239289c)
- [voltstack_cluster.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-004.md#canonical-3fe609189efa5a9ade4fedbc3f0522eba9d315551337bb3c00058bf55b3c58eb)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-8310edad4caa3080989a1a714257510f6e17d6f53bd38b5c027385223239289c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9dad43f42ca27dda1e55e837aee5c4243e573c64ec0d23c3f0d63c935a29360a"></a>

## voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / d5738c38c0f7 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.global_network_list](resources--gcp_vpc_site--reference--group-004.md#canonical-7f7b911b7737de6d66f5805b10981b5da71f85fae44df698f77fc3d79ccdc908)
- [voltstack_cluster.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-004.md#canonical-3fe609189efa5a9ade4fedbc3f0522eba9d315551337bb3c00058bf55b3c58eb)
- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-9a30171c28b1cec3eead5c30b2c8cbc2add7ccb91703d194286fe574dc2cb35d)
- voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-73486476f74bc188c2d336f365128aa66e03b1fae6c9407f2437fd720014a38d"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-08e6c46c1f264e445fef9712041624823bfd3a726297a309c44dc74feb2c5007"></a>

## Direct properties — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / d5738c38c0f7 / 3

<a id="canonical-6d2e19994c60263a0d89d4a3cecc655565d87997364eded12809b53774bad99c"></a>

<a id="canonical-2659fb94fd66ebfe5d718e5b1e6844448214110d5aed1fb4360f4ad9a220b9e3"></a>

## name property — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / d5738c38c0f7 / 4

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

<a id="canonical-77ae2eb54889df98777dce829cf26ca19d8ca402c045a418b8589b7a11a4cea2"></a>

<a id="canonical-c92c09a2476ad367aaf09d1fee0c9b2c7782d213d546c3d899c7544c5e9e3b29"></a>

## namespace property — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / d5738c38c0f7 / 5

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

<a id="canonical-6bb17bcaade8d00203487d451d355300735a8c329ccf5e82ce0d84240ade731c"></a>

<a id="canonical-038638ff22c796c21ff28285c4822db7becc00c92ce0f78c0cca2ef91f24fd2c"></a>

## tenant property — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / d5738c38c0f7 / 6

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

<a id="canonical-5a5c71554652c55f4144abd3d59326574256e240d124de46d25c9cbc3cd55be2"></a>

## Next pages — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / d5738c38c0f7 / 7

- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-9a30171c28b1cec3eead5c30b2c8cbc2add7ccb91703d194286fe574dc2cb35d)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-6ab5dfec17019798407b97ffe626c7deaa095d98081ae3cbba570ceb068204ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06deed15c42787f2cc69deba3b3e092cb0982f0a39f14ea0fc6d745fcef1c78a"></a>

## voltstack_cluster.k8s_cluster — voltstack_cluster.k8s_cluster / 301d929c76e6 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.k8s_cluster

<a id="canonical-2bf47800db46259ad52ac94e9319c52820f937344c07c766342ce6d457f23a96"></a>

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
k8s_cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-267bce1d1285c9ce2ca5631b211eda1e1a0b05d1614845fcfd7e0a7478e61524"></a>

## Direct properties — voltstack_cluster.k8s_cluster / 301d929c76e6 / 3

<a id="canonical-dfc9c74019b1f6802985371eea12d70001e7b599de5cd561b74ae24d5b9f16db"></a>

<a id="canonical-36123fa33424178be825dec990e271d9afc4c74c222a9bb288ccc96460e39fe2"></a>

## name property — voltstack_cluster.k8s_cluster / 301d929c76e6 / 4

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

<a id="canonical-29915eee296f345040d6fd82d764be1d4e31904ff3ef15e2d92442f48ee94118"></a>

<a id="canonical-c5373187a41a48feb3dc16ed560b632d349f94c9b102b0654656047481e63539"></a>

## namespace property — voltstack_cluster.k8s_cluster / 301d929c76e6 / 5

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

<a id="canonical-52d9563136f8de026a87d7b4fb883381605bdf4d15300f2fb947e19c8dbe321b"></a>

<a id="canonical-19acf24dc2bf3f152aa52cb7b830a74ffede139dbd9c7fe620171edf7f767ce5"></a>

## tenant property — voltstack_cluster.k8s_cluster / 301d929c76e6 / 6

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

<a id="canonical-59bff33a28a0e7977c244599146d3af8cbb5e47c7aa372b9642fed3838ab80d9"></a>

## Next pages — voltstack_cluster.k8s_cluster / 301d929c76e6 / 7

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-4176f808cef5388ec183126f11b664dbf0b4f8be9686450dfa5f0a0c5f9f5c6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd6436d2d6a685efa6ff4fa7bcc96551b07a402ac5f6d4f1c27e892c027a0173"></a>

## voltstack_cluster.no_dc_cluster_group — voltstack_cluster.no_dc_cluster_group / a2adc66a6d69 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.no_dc_cluster_group

<a id="canonical-e8878600451a685834f2706b24e0c8f7e5b5e1f3a59028efc4e8836a39516456"></a>

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

<a id="canonical-d1ee9f539d6ac2ce124a4f631f812ab49a2647c35146932903609b1759803e53"></a>

## Direct properties — voltstack_cluster.no_dc_cluster_group / a2adc66a6d69 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-11294bab2ea29a28fef5d82d029558ef34b8d651ee675f6cd18a3bc667bfc139"></a>

## Next pages — voltstack_cluster.no_dc_cluster_group / a2adc66a6d69 / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-d0e8998e48affcaa5ac09252943e3c30a4136b272ff3c5b5d1f8cf91e1010cfe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed52a500b8a830f89111c664a972dfbb6656f8b097a751b1dfd53c7acae2f222"></a>

## voltstack_cluster.no_forward_proxy — voltstack_cluster.no_forward_proxy / 9534438d84c3 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.no_forward_proxy

<a id="canonical-981570e351655185177fac094af43bc06253137065e723867cd157d1f5425f0a"></a>

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

<a id="canonical-182a2c57617a758b9e6ba6b82ec48c12e56faa1ad1446bebe1038493f8453f24"></a>

## Direct properties — voltstack_cluster.no_forward_proxy / 9534438d84c3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bb78117e162843f445f79cbc9eecd705667c3dd37a5d61aa8ae10c1acc7cdcaa"></a>

## Next pages — voltstack_cluster.no_forward_proxy / 9534438d84c3 / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-b5541b65805f5fcb7e007ec8a55ff04911dfc2a02ba9f85b1aa3061d3d57aa58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3cc289f0eb07d2b9c89f9e1fb8f01dea9f3361a461bd67d0cd7fd55f77a86067"></a>

## voltstack_cluster.no_global_network — voltstack_cluster.no_global_network / ff31eeae3995 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.no_global_network

<a id="canonical-0dc62f3fc25df0223740f2aedf805c5672a0e8cc65b659b0ee2836c87e6a51b7"></a>

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

<a id="canonical-821e138e8acc03e4d28e935821c739744579c019cdc9815511d36633bae35b01"></a>

## Direct properties — voltstack_cluster.no_global_network / ff31eeae3995 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-09879106e081b9851f6a986d8a68ef8c20049d5a11623c469888acf4b7bf36f8"></a>

## Next pages — voltstack_cluster.no_global_network / ff31eeae3995 / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-e8176eadfa957c1e92e1aaf5996293989efcf098de8ccda58961c6c6a0648228"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8fe21804b7dd1c4dd4bc2020545e37b32b6964c7273233bd08eda6aa198eb46c"></a>

## voltstack_cluster.no_k8s_cluster — voltstack_cluster.no_k8s_cluster / a060c4be3d29 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.no_k8s_cluster

<a id="canonical-5a986bc0ab426a8dc93c4c643f6dc94236c5e62ae14730948ef59a749068d752"></a>

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
no_k8s_cluster = {}
```

<a id="canonical-7bb9408eb2c5e8f85e26979b79edde822bdbbdb5b7666a730dbabb7bee14b017"></a>

## Direct properties — voltstack_cluster.no_k8s_cluster / a060c4be3d29 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-71e0e9920e8c8c54654a53fdfb4c63420bfdcf534ae507b1e8065900b6894296"></a>

## Next pages — voltstack_cluster.no_k8s_cluster / a060c4be3d29 / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-193ae660e21c2dc90147e84bcb3396af6e2d7d7c219241ac7e1e7254ad55b4f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d4018cca19641f72f84027f7dc121e27dff52b4b5050b3afb4545037594b20a"></a>

## voltstack_cluster.no_network_policy — voltstack_cluster.no_network_policy / 6ec0db046e3e / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.no_network_policy

<a id="canonical-b899c4074c916ef91648ffe0310cb2432c6a4da1a9010971b0dfa32f91adb25c"></a>

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

<a id="canonical-88793a9e87ac536126d7d4ec3a0254c9d62d5dff12d84025029af8e840b4ba96"></a>

## Direct properties — voltstack_cluster.no_network_policy / 6ec0db046e3e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-727ea4a4ce845edd5d7f4eeb7fea11c8c71b6b69dc8f15edfe89eea7d8fd0d11"></a>

## Next pages — voltstack_cluster.no_network_policy / 6ec0db046e3e / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-e96b9e9e1b31efdc5030732c65d2691e3a5c5f1a3d28565c89c5e805a36374d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fae4184d5b968e861596ae0f4e86617065d99921b0f73b642aaf3da72d3851e"></a>

## voltstack_cluster.no_outside_static_routes — voltstack_cluster.no_outside_static_routes / abf7aa8b8bee / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.no_outside_static_routes

<a id="canonical-1375c16ad0b028bc75ae2ce39634a682cf1f7309a0861bb22261785b6e3ffdc6"></a>

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

<a id="canonical-85f64376dde2b4ae46a4bbe2ec1e9f8abe62070d3ac5534bccce3c67da51cee9"></a>

## Direct properties — voltstack_cluster.no_outside_static_routes / abf7aa8b8bee / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d202add2b2cf791c4fcb0b74eac2f5247dda8b4f2e1458364eb65100488b8808"></a>

## Next pages — voltstack_cluster.no_outside_static_routes / abf7aa8b8bee / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-5c3c562c7707d83d434086992a2febf9d66463cf41d247a396db55a64f80c18a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5ca1d64a626b9be7a0983988444e4bfc75fd62b3700b8a05e6663e6e5f95d20"></a>

## voltstack_cluster.outside_static_routes — voltstack_cluster.outside_static_routes / 5aacea086ea3 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.outside_static_routes

<a id="canonical-d5bcf5569367f54a95054ce45e587f464533e941209cd060f1542c6a096056f1"></a>

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

<a id="canonical-99430eab2211aa76675b94b68bbf5c8a57531f661f77de12fd56b734ee249b02"></a>

## Direct properties — voltstack_cluster.outside_static_routes / 5aacea086ea3 / 3

- [static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-bbc31abb997b7c3b5aa9e04ae42305aa4d8ebb238c7af872395398e21a7e0715): complete subsection reference.

<a id="canonical-0173319a88912802e4059326ca1244e49336cbd55038c7070121044a5e49cf34"></a>

## Next pages — voltstack_cluster.outside_static_routes / 5aacea086ea3 / 4

- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-bbc31abb997b7c3b5aa9e04ae42305aa4d8ebb238c7af872395398e21a7e0715)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-bbc31abb997b7c3b5aa9e04ae42305aa4d8ebb238c7af872395398e21a7e0715"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-988f9875c4f6695f477da39b3956f70513d2fc133a85acb715c8f34a913131a3"></a>

## voltstack_cluster.outside_static_routes.static_route_list — voltstack_cluster.outside_static_routes.static_route_list / 406667de1821 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-5c3c562c7707d83d434086992a2febf9d66463cf41d247a396db55a64f80c18a)
- voltstack_cluster.outside_static_routes.static_route_list

<a id="canonical-c27e825aae2c35d9fafeaca8adbf4f0a53e6f055124b080ca5bf3f049cd2dccb"></a>

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

<a id="canonical-14095a5b148542ec30e0543087e12715d38e48d26c2b65d1fe01c9d9d1202aa2"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list / 406667de1821 / 3

- [custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-33ef5592b6827447c99cd4007c40eebc8c7e4451ab7d1dca86a0830e6c082cf9): complete subsection reference.

<a id="canonical-3c49603c4c2c5a9034e1097061e22e7b4e590a0589d129996331e9bf58864bb1"></a>

<a id="canonical-c741836f79313f3dd14c3ccc91063f336e2f1522cb94794efbbb033335eb0465"></a>

## simple_static_route property — voltstack_cluster.outside_static_routes.static_route_list / 406667de1821 / 4

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

<a id="canonical-931c333da863bf7f8dabaee5219685c923ab048dc8b2de5846676840bf0c2fb7"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list / 406667de1821 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-33ef5592b6827447c99cd4007c40eebc8c7e4451ab7d1dca86a0830e6c082cf9)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-5c3c562c7707d83d434086992a2febf9d66463cf41d247a396db55a64f80c18a)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-33ef5592b6827447c99cd4007c40eebc8c7e4451ab7d1dca86a0830e6c082cf9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e191717028b1c756480de9226b92d791628c9528168c3cc8269706414636b89"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route / 3e4ef7267b8e / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-5c3c562c7707d83d434086992a2febf9d66463cf41d247a396db55a64f80c18a)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-bbc31abb997b7c3b5aa9e04ae42305aa4d8ebb238c7af872395398e21a7e0715)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-59d2cb72733d728f978a2e335317dfa7d36dc31008ad9e3adf2179b8c9f7b3b5"></a>

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

<a id="canonical-4c8220bb76f03b4f393d4bef66795946aaf86005aae60105e4fc3da486cf5315"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route / 3e4ef7267b8e / 3

<a id="canonical-79083fc6040b8ffa85dbefb5da7a2a41e66219c47976fded4ca59a2f7a154179"></a>

<a id="canonical-2fe4ae415064905038493671a8be303cf509fca41098ea84490857faadae0c9a"></a>

## attrs property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route / 3e4ef7267b8e / 4

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

- [labels](resources--gcp_vpc_site--reference--group-004.md#canonical-31f56b2544dd49904e35aee6027c71fde06298f5d731309cb9708b999e23c354): complete subsection reference.

- [nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-f91734707a4e9dc9b605ea59281da529084ad43d6046e1eee365eb0ce259dd29): complete subsection reference.

- [subnets](resources--gcp_vpc_site--reference--group-004.md#canonical-976d8c4c6aaae4fa797d5ed45a1fb857fd9c1e2fa95fa6278f0ba6bbafc1fa55): complete subsection reference.

<a id="canonical-5e8333cb2453bd4e596c7d6c5fc65352b10b7b4ae23253a0b751e1ab3f238fd0"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route / 3e4ef7267b8e / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](resources--gcp_vpc_site--reference--group-004.md#canonical-31f56b2544dd49904e35aee6027c71fde06298f5d731309cb9708b999e23c354)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-f91734707a4e9dc9b605ea59281da529084ad43d6046e1eee365eb0ce259dd29)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-004.md#canonical-976d8c4c6aaae4fa797d5ed45a1fb857fd9c1e2fa95fa6278f0ba6bbafc1fa55)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-bbc31abb997b7c3b5aa9e04ae42305aa4d8ebb238c7af872395398e21a7e0715)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-31f56b2544dd49904e35aee6027c71fde06298f5d731309cb9708b999e23c354"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1e1680b5c4482d34b6e0c2c73039184285667fb8fa2b845643bc2b1d16a1f19"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.la / 3dbf71f90e2e / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-5c3c562c7707d83d434086992a2febf9d66463cf41d247a396db55a64f80c18a)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-bbc31abb997b7c3b5aa9e04ae42305aa4d8ebb238c7af872395398e21a7e0715)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-33ef5592b6827447c99cd4007c40eebc8c7e4451ab7d1dca86a0830e6c082cf9)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-d155129d55816cf875cc180bc9831941887f82187031a433eef90c3aca9d7b54"></a>

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

<a id="canonical-7e59d9e01919d1f50b2e7aea5654ae5807e7dee79848fe947dc4de90847693fc"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.la / 3dbf71f90e2e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c57be633e17df2365dbb3606130bce6b2751d1e6ab94894427da54407bd7816f"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.la / 3dbf71f90e2e / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-33ef5592b6827447c99cd4007c40eebc8c7e4451ab7d1dca86a0830e6c082cf9)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-f91734707a4e9dc9b605ea59281da529084ad43d6046e1eee365eb0ce259dd29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-586ef4c227e6b55bb1b30b3d729b415ea6b6f3576f6954642ddd7e96b77344f3"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 19ec69ef2322 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-5c3c562c7707d83d434086992a2febf9d66463cf41d247a396db55a64f80c18a)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-bbc31abb997b7c3b5aa9e04ae42305aa4d8ebb238c7af872395398e21a7e0715)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-33ef5592b6827447c99cd4007c40eebc8c7e4451ab7d1dca86a0830e6c082cf9)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-aeb6f82ed0d3c87b4376b0f877aaa23fa3123205eb298dd9c5408a2987088adf"></a>

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

<a id="canonical-0d4174b849eb75e8c97640fbdc36f1303626cc8c458ac60341030295dc0780d1"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 19ec69ef2322 / 3

- [interface](resources--gcp_vpc_site--reference--group-004.md#canonical-dd500f19fca6312c21f2464ef3148b20242ab7e5d9acbe12bf42cb979be48f73): complete subsection reference.

- [nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-80a352763c6f8889257d2662d6de5695aecfeb4a0a1578d35753054b0238e819): complete subsection reference.

<a id="canonical-335941e8353a5d3e85755115d465cd04ecfdab59b44cc5eac5a753f20c92ac94"></a>

<a id="canonical-3ae0c2ef3e19bd75748e441a3421f7bc5b2ce64a964a65d656e66e28e86beeaf"></a>

## type property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 19ec69ef2322 / 4

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

<a id="canonical-d1da2b3c4ff0509b6b7d8e2db5b206d8abfe884f3ab8a9ad4fef052aff654d2b"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 19ec69ef2322 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--gcp_vpc_site--reference--group-004.md#canonical-dd500f19fca6312c21f2464ef3148b20242ab7e5d9acbe12bf42cb979be48f73)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-80a352763c6f8889257d2662d6de5695aecfeb4a0a1578d35753054b0238e819)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-33ef5592b6827447c99cd4007c40eebc8c7e4451ab7d1dca86a0830e6c082cf9)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-dd500f19fca6312c21f2464ef3148b20242ab7e5d9acbe12bf42cb979be48f73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4f434f9c828b4a493aafdebbbf5d2429f87045dc134a49ed1db83ded7eb4b7f"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 99e93cbe8ddd / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-5c3c562c7707d83d434086992a2febf9d66463cf41d247a396db55a64f80c18a)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-bbc31abb997b7c3b5aa9e04ae42305aa4d8ebb238c7af872395398e21a7e0715)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-33ef5592b6827447c99cd4007c40eebc8c7e4451ab7d1dca86a0830e6c082cf9)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-f91734707a4e9dc9b605ea59281da529084ad43d6046e1eee365eb0ce259dd29)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-3244e851722af24c4619a34896da1794ed4c7ec1dc55a6995e051777d0be8115"></a>

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

<a id="canonical-9fdebfe77d1d62b4a044eafd2de86442975e01e32052246542d5463491e2e2fc"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 99e93cbe8ddd / 3

<a id="canonical-ea4418e1dd97338ab16dadcd61bd067a92aff19d09933eff2c820c7502a13044"></a>

<a id="canonical-efc4e9340f3ad1090d03d52b70d8d652ebfc046597608a80d53ee0b4f51a7a09"></a>

## kind property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 99e93cbe8ddd / 4

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

<a id="canonical-c49afbadc73d70db63c8eb7165c1853fbbc1234eedb4f23807a9031e6f1b16b6"></a>

<a id="canonical-b0a973af18284c51834c1fdbe148c9b65e3ef739195fdc9e7afa7cc405d520cb"></a>

## name property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 99e93cbe8ddd / 5

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

<a id="canonical-14d3055b1ec7516694e8d7b6289ba11d7c1d6c4be52346f659c58294ff91dd5c"></a>

<a id="canonical-272727cd23973d819e132d4f8bff86b97d0a012454566dd4fe604823ecafa2a3"></a>

## namespace property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 99e93cbe8ddd / 6

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

<a id="canonical-09967afd5dde12f56f7e1ef7dccf9e240d632980f0bccf8a758b6b84ebf76672"></a>

<a id="canonical-4e266262f8a9b0a4193fa2ca950328500a1a1acebc4408b58447020d7a0f97e3"></a>

## tenant property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 99e93cbe8ddd / 7

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

<a id="canonical-c8ad32bdb63c2b4edad6e99fc1f0f2c307e12544e8f7dcebdde651a8b55a2a4f"></a>

<a id="canonical-85ff17e5bf90b76e126114246928bbb16e4d461ac072831caf7368a9c83e7800"></a>

## uid property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 99e93cbe8ddd / 8

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

<a id="canonical-c0e0476578ed8e6b54002d0cf83e2bed622cb6d0cf72e6e1e772d6e4bcfb2444"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 99e93cbe8ddd / 9

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-f91734707a4e9dc9b605ea59281da529084ad43d6046e1eee365eb0ce259dd29)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-80a352763c6f8889257d2662d6de5695aecfeb4a0a1578d35753054b0238e819"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78399877cb64cac33aef974cbcbb9a1601d5a0964cba36f0d8e764ff3acc7e36"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / f6e462de525b / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-5c3c562c7707d83d434086992a2febf9d66463cf41d247a396db55a64f80c18a)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-bbc31abb997b7c3b5aa9e04ae42305aa4d8ebb238c7af872395398e21a7e0715)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-33ef5592b6827447c99cd4007c40eebc8c7e4451ab7d1dca86a0830e6c082cf9)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-f91734707a4e9dc9b605ea59281da529084ad43d6046e1eee365eb0ce259dd29)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-59d654a088d6925f97de3e48fab7f5db8a9b1cf42fd156c2724419bb4c84a6f3"></a>

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

<a id="canonical-5dd166d4cb4da2d32ed03302edc582d04559414087d5192cc670cee4831f752e"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / f6e462de525b / 3

- [dual_stack](resources--gcp_vpc_site--reference--group-004.md#canonical-cbb983bdb4fb307cb6dfe8eafec36f77f0b725c66e762380ba473b4ca12b039e): complete subsection reference.

- [ipv4](resources--gcp_vpc_site--reference--group-004.md#canonical-09618a886f1fb443197362336080232b920b56ee329b8ed16b7d44c14f87d92e): complete subsection reference.

- [ipv6](resources--gcp_vpc_site--reference--group-004.md#canonical-3fa390ab7cbd56ceea0be0528cc17e3cf97c24a6688fa9bddba955a7456417e7): complete subsection reference.

<a id="canonical-b70920a01147c9e95fd5bcb39a3bb225611db4c3abfc9824d2b6d54455bd6293"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / f6e462de525b / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-004.md#canonical-cbb983bdb4fb307cb6dfe8eafec36f77f0b725c66e762380ba473b4ca12b039e)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--gcp_vpc_site--reference--group-004.md#canonical-09618a886f1fb443197362336080232b920b56ee329b8ed16b7d44c14f87d92e)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--gcp_vpc_site--reference--group-004.md#canonical-3fa390ab7cbd56ceea0be0528cc17e3cf97c24a6688fa9bddba955a7456417e7)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-f91734707a4e9dc9b605ea59281da529084ad43d6046e1eee365eb0ce259dd29)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-cbb983bdb4fb307cb6dfe8eafec36f77f0b725c66e762380ba473b4ca12b039e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ce4beb095eb0ceb5398026a5068ba1c4bbc1fa4fe54200dc1a9ad2e55d1b532"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 47a9735ece2a / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-5c3c562c7707d83d434086992a2febf9d66463cf41d247a396db55a64f80c18a)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-bbc31abb997b7c3b5aa9e04ae42305aa4d8ebb238c7af872395398e21a7e0715)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-33ef5592b6827447c99cd4007c40eebc8c7e4451ab7d1dca86a0830e6c082cf9)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-f91734707a4e9dc9b605ea59281da529084ad43d6046e1eee365eb0ce259dd29)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-80a352763c6f8889257d2662d6de5695aecfeb4a0a1578d35753054b0238e819)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-2a24b26f4b58f020024562bd07e823c492b4e0b297f7a9650af842df3b85bf7e"></a>

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

<a id="canonical-b4a1186602f10b27f8f720e299f483240d446650a2a9da6fc466174647b14f63"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 47a9735ece2a / 3

- [ipv4](resources--gcp_vpc_site--reference--group-004.md#canonical-854162f6cefb4ce5e56707d6eb5e7fbfe34368a4082479022cfb08aa1dce2804): complete subsection reference.

- [ipv6](resources--gcp_vpc_site--reference--group-004.md#canonical-2d743a1e4c9a9d5b654d86ccf1b6a7ec480875c80a0ef73e189bf651e2a759c9): complete subsection reference.

<a id="canonical-7dafd06676f2da92204799957e243a4aa8bd93e9ffda45f5b64f3285fe044cfd"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 47a9735ece2a / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--gcp_vpc_site--reference--group-004.md#canonical-854162f6cefb4ce5e56707d6eb5e7fbfe34368a4082479022cfb08aa1dce2804)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--gcp_vpc_site--reference--group-004.md#canonical-2d743a1e4c9a9d5b654d86ccf1b6a7ec480875c80a0ef73e189bf651e2a759c9)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-80a352763c6f8889257d2662d6de5695aecfeb4a0a1578d35753054b0238e819)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-854162f6cefb4ce5e56707d6eb5e7fbfe34368a4082479022cfb08aa1dce2804"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1a0f95404c2de2e0ceb4bd8b4bad759e4c8b39bd1cd4b8c62d147fe213bf330"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 736c59d55d66 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-5c3c562c7707d83d434086992a2febf9d66463cf41d247a396db55a64f80c18a)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-bbc31abb997b7c3b5aa9e04ae42305aa4d8ebb238c7af872395398e21a7e0715)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-33ef5592b6827447c99cd4007c40eebc8c7e4451ab7d1dca86a0830e6c082cf9)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-f91734707a4e9dc9b605ea59281da529084ad43d6046e1eee365eb0ce259dd29)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-80a352763c6f8889257d2662d6de5695aecfeb4a0a1578d35753054b0238e819)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-004.md#canonical-cbb983bdb4fb307cb6dfe8eafec36f77f0b725c66e762380ba473b4ca12b039e)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-cf60c6ee02bc0def55a94f3199b7f8c22279ebd4d479f22a019fd1847a7f9a14"></a>

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

<a id="canonical-de353c71fd2e2ad75a3760daec94116f1d6e885355e055bd9cd88331052aa7c7"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 736c59d55d66 / 3

<a id="canonical-a0aa7d924e0e1cb413cafdb01f6016aa2b7dcfd861c91cdeae2fcb00ca422386"></a>

<a id="canonical-2cdc91c0c50b7a2fc8e68cb2436c39bd7b02414759f39a013d53649cb2bf195a"></a>

## addr property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 736c59d55d66 / 4

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

<a id="canonical-f2c934e0cf9c3cb8d343a6bc29cbce1eda3db41f27db5e759ddac8db2dd61fcc"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 736c59d55d66 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-004.md#canonical-cbb983bdb4fb307cb6dfe8eafec36f77f0b725c66e762380ba473b4ca12b039e)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-2d743a1e4c9a9d5b654d86ccf1b6a7ec480875c80a0ef73e189bf651e2a759c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-117ac1ba8d0a1a2bcc56f9b7243881f0a3acf2c9202a37adb0f0e8ee999aa61b"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 6f98aa0accbf / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-5c3c562c7707d83d434086992a2febf9d66463cf41d247a396db55a64f80c18a)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-bbc31abb997b7c3b5aa9e04ae42305aa4d8ebb238c7af872395398e21a7e0715)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-33ef5592b6827447c99cd4007c40eebc8c7e4451ab7d1dca86a0830e6c082cf9)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-f91734707a4e9dc9b605ea59281da529084ad43d6046e1eee365eb0ce259dd29)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-80a352763c6f8889257d2662d6de5695aecfeb4a0a1578d35753054b0238e819)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-004.md#canonical-cbb983bdb4fb307cb6dfe8eafec36f77f0b725c66e762380ba473b4ca12b039e)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-297dc78fc661acb899c262a2e5714d58fe6702ff4486c59962b2851a2d46f5dc"></a>

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

<a id="canonical-f2f42f563254a13bc284122e93ce7f4c429268ea226ccd0d299a94bd2c31059f"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 6f98aa0accbf / 3

<a id="canonical-484dcf9533a112387fa504ae7498e2b5b25aab68a775518bf4e84965d96bb25e"></a>

<a id="canonical-a2bf87391f2149de2117f43933312d8c387b0f271c1ac72877da8873d1785ea0"></a>

## addr property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 6f98aa0accbf / 4

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

<a id="canonical-e9d5a6eec43b0c83b9877bfb120b50024177c69eb810c95f14deb919621047af"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 6f98aa0accbf / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-004.md#canonical-cbb983bdb4fb307cb6dfe8eafec36f77f0b725c66e762380ba473b4ca12b039e)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-09618a886f1fb443197362336080232b920b56ee329b8ed16b7d44c14f87d92e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c928ee6f972de9d5929beea32d50e3ac8dfcb1c82b59566b547b41375d7cc225"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 27364e2ba339 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-5c3c562c7707d83d434086992a2febf9d66463cf41d247a396db55a64f80c18a)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-bbc31abb997b7c3b5aa9e04ae42305aa4d8ebb238c7af872395398e21a7e0715)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-33ef5592b6827447c99cd4007c40eebc8c7e4451ab7d1dca86a0830e6c082cf9)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-f91734707a4e9dc9b605ea59281da529084ad43d6046e1eee365eb0ce259dd29)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-80a352763c6f8889257d2662d6de5695aecfeb4a0a1578d35753054b0238e819)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-117e9fc0a8dcd4b43032a76674e51f12d14415a63655e224f1c1fe3386c559bb"></a>

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

<a id="canonical-f579b2f906a6cebc693978a28d4ebb64c6bf9be85a7cfb1d6bd8cabe91e11606"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 27364e2ba339 / 3

<a id="canonical-acb0a7e9d7f2cc6e48e3ab0ec331a2263b87cb782f5cf5dff0ab6796c1cb8ab4"></a>

<a id="canonical-19d836909903b69e664d3aef283f32757efb2a642bc3f2197a9260859de60c18"></a>

## addr property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 27364e2ba339 / 4

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

<a id="canonical-335e3677d95565485ba1b965a63bcb7e4d1ab396ae797af9b1e9df0ff582f786"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 27364e2ba339 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-80a352763c6f8889257d2662d6de5695aecfeb4a0a1578d35753054b0238e819)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-3fa390ab7cbd56ceea0be0528cc17e3cf97c24a6688fa9bddba955a7456417e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8135449d700fcdd01d35c316e1378e18d094c3358b7daced467f0429eebb2dcb"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / dda9c1b30947 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-5c3c562c7707d83d434086992a2febf9d66463cf41d247a396db55a64f80c18a)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-bbc31abb997b7c3b5aa9e04ae42305aa4d8ebb238c7af872395398e21a7e0715)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-33ef5592b6827447c99cd4007c40eebc8c7e4451ab7d1dca86a0830e6c082cf9)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-f91734707a4e9dc9b605ea59281da529084ad43d6046e1eee365eb0ce259dd29)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-80a352763c6f8889257d2662d6de5695aecfeb4a0a1578d35753054b0238e819)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-1a1f1c3234e7ecffbb08cf9dc95947219dca7764d8894c84fa496383d966f60c"></a>

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

<a id="canonical-e11831428de1087bc4ee4b135cef59372544e9b445f627d3893f18318a8e8032"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / dda9c1b30947 / 3

<a id="canonical-4fe406a9d724af87e6290b954ebb2f5bcdb9d13afcf1bac6c82256fcd9c7778b"></a>

<a id="canonical-97f9da7503c26af18f772099ca3ac6aa3a4ad0d1f65789a7552f46d385720027"></a>

## addr property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / dda9c1b30947 / 4

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

<a id="canonical-f4498843c2de549eb09f0797dd37a5689b1b25cd2f032d1d92962dd280994d44"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / dda9c1b30947 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-80a352763c6f8889257d2662d6de5695aecfeb4a0a1578d35753054b0238e819)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-976d8c4c6aaae4fa797d5ed45a1fb857fd9c1e2fa95fa6278f0ba6bbafc1fa55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2318082f3306d6c50e3a15977fc6844d4623814d0f6e4cab5d9e366392b5d92a"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 5afdb8278b82 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-5c3c562c7707d83d434086992a2febf9d66463cf41d247a396db55a64f80c18a)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-bbc31abb997b7c3b5aa9e04ae42305aa4d8ebb238c7af872395398e21a7e0715)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-33ef5592b6827447c99cd4007c40eebc8c7e4451ab7d1dca86a0830e6c082cf9)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-3740612018e269ed2f526e7dc4db014f976a0d2bc73138f8dab123daee31ea21"></a>

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

<a id="canonical-b2e2965b9e885d2728ef8beece76d0ef503eb4657a7b33a458030d0f5c8f156b"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 5afdb8278b82 / 3

- [ipv4](resources--gcp_vpc_site--reference--group-004.md#canonical-ac9a56a40e52901d9c75adbeff71276149531c01a364a553721e5c4cb605bba6): complete subsection reference.

- [ipv6](resources--gcp_vpc_site--reference--group-004.md#canonical-3e231c5503f434e02914bc34b9b966c0fc1632944bea9888308c15acfffe64e1): complete subsection reference.

<a id="canonical-5cedf911869c1cdee5dfa64f24c8e3574c2b9673a37d2dda47d867bf1cb6b8ea"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 5afdb8278b82 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--gcp_vpc_site--reference--group-004.md#canonical-ac9a56a40e52901d9c75adbeff71276149531c01a364a553721e5c4cb605bba6)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--gcp_vpc_site--reference--group-004.md#canonical-3e231c5503f434e02914bc34b9b966c0fc1632944bea9888308c15acfffe64e1)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-33ef5592b6827447c99cd4007c40eebc8c7e4451ab7d1dca86a0830e6c082cf9)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-ac9a56a40e52901d9c75adbeff71276149531c01a364a553721e5c4cb605bba6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a68f6c05817f8366600c4d97412d862565f28e66c9790dc2672fa986b56d9f15"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / d959684524fb / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-5c3c562c7707d83d434086992a2febf9d66463cf41d247a396db55a64f80c18a)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-bbc31abb997b7c3b5aa9e04ae42305aa4d8ebb238c7af872395398e21a7e0715)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-33ef5592b6827447c99cd4007c40eebc8c7e4451ab7d1dca86a0830e6c082cf9)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-004.md#canonical-976d8c4c6aaae4fa797d5ed45a1fb857fd9c1e2fa95fa6278f0ba6bbafc1fa55)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-6d909efbb462fb3a5c3146ff06577fccd99a10e19468a4b5c76bdcb83863ae12"></a>

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

<a id="canonical-acb6d66439968f4782f225ad4a7950d1b9cc1923e37c0aef589231355bdc04d0"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / d959684524fb / 3

<a id="canonical-9e458e64ee558b09911d62280858a0120aae428aee35a82257c72b5438963cb6"></a>

<a id="canonical-9219651621f97d69d6931577fb7efda88670c22201b36e6072b6bf292121f52b"></a>

## plen property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / d959684524fb / 4

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

<a id="canonical-994b89208b6d1db05bf0a2762aadc4f5dd6485b3e00507bb78e66946a8ebb7e3"></a>

<a id="canonical-8815329788dec6dd67eae62aad9f0371e60b4120f09e180a9c5a7df4c79793ec"></a>

## prefix property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / d959684524fb / 5

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

<a id="canonical-54ebc49772dd704c6560c7b9f8b6fbf65015f5cefd8bad7477a8765a9b022d24"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / d959684524fb / 6

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-004.md#canonical-976d8c4c6aaae4fa797d5ed45a1fb857fd9c1e2fa95fa6278f0ba6bbafc1fa55)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-3e231c5503f434e02914bc34b9b966c0fc1632944bea9888308c15acfffe64e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d71dca67eeffda8c5ba242338ae7942a15d32e48e295f32e7f9ff19eab3ff008"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 5f6c2f267913 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-5c3c562c7707d83d434086992a2febf9d66463cf41d247a396db55a64f80c18a)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-bbc31abb997b7c3b5aa9e04ae42305aa4d8ebb238c7af872395398e21a7e0715)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-33ef5592b6827447c99cd4007c40eebc8c7e4451ab7d1dca86a0830e6c082cf9)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-004.md#canonical-976d8c4c6aaae4fa797d5ed45a1fb857fd9c1e2fa95fa6278f0ba6bbafc1fa55)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-5b1fb5985180339ad98209e369bb3878a057b46bcf668dd2b225b73c8c8d6be8"></a>

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

<a id="canonical-9ce8b69c046bd90fe7ca87528b20ba2050e2e8cc6bb64c5c9ab038194a2257b9"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 5f6c2f267913 / 3

<a id="canonical-d897c305925daa3610f9fb162a7f94dc412d1817aa48b564b39ceaf63f0f3248"></a>

<a id="canonical-a38ed2728fd85304536d8ae8556fa90574360cfdbfdcd31c069187ef4c71a095"></a>

## plen property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 5f6c2f267913 / 4

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

<a id="canonical-d5ade6dca7076a1afd37f7e9f6f0e7b49e26d39ebf016260c279a4b237a84d23"></a>

<a id="canonical-73604be9bf5ddcffdb3458af5c33f1de6209b15fe6c97328f88d72d15bd430c1"></a>

## prefix property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 5f6c2f267913 / 5

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

<a id="canonical-5cab038432ae32c4e1f3d6bb464c97b95f69f80742335b377bb52ff39f734f71"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 5f6c2f267913 / 6

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-004.md#canonical-976d8c4c6aaae4fa797d5ed45a1fb857fd9c1e2fa95fa6278f0ba6bbafc1fa55)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-0d117f633e2656829fba928f7b9b2a047ea066556f60ce1d03c0e04f47f29fa0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbe9dfcac2c84103d6b91d7845924f240d425f5a8e2dc092af5c5b1792f4a984"></a>

## voltstack_cluster.site_local_network — voltstack_cluster.site_local_network / c50e2716dbca / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.site_local_network

<a id="canonical-f658f139f77e0428765d3100c68fde04cffb01647313dac0f40e07a1241e3886"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_network",
    "new_network"),
  validators.ConflictingObjectAttributes("existing_network",
    "new_network_autogenerate"),
  validators.ConflictingObjectAttributes("new_network",
    "new_network_autogenerate")}
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
  "x-ves-oneof-field-choice": "[\"existing_network\",\"new_network\",\"new_network_autogenerate\"]"
}
```

Terraform syntax:

```terraform
site_local_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-aaab52c96e3831af191d46719c563a757614ee53e0f9c96d027345c1cfbebd1b"></a>

## Direct properties — voltstack_cluster.site_local_network / c50e2716dbca / 3

- [existing_network](resources--gcp_vpc_site--reference--group-004.md#canonical-925f7925df9001ec9bded593d287cadfedbdc0583c21d790af5435b955be4cb2): complete subsection reference.

- [new_network](resources--gcp_vpc_site--reference--group-004.md#canonical-c9a5d044d970e4e003509bbbf603677b43b47f66d960d431eacebc31041e1e9b): complete subsection reference.

- [new_network_autogenerate](resources--gcp_vpc_site--reference--group-004.md#canonical-a924dc90b4b4f595bffd49e0e6d5b7e185a6d23fdb692ea540f31ae238555d5d): complete subsection reference.

<a id="canonical-517bf846416f10492c76cdf822fc8416cac98548a107635e0f3661060b8e027e"></a>

## Next pages — voltstack_cluster.site_local_network / c50e2716dbca / 4

- [voltstack_cluster.site_local_network.existing_network](resources--gcp_vpc_site--reference--group-004.md#canonical-925f7925df9001ec9bded593d287cadfedbdc0583c21d790af5435b955be4cb2)
- [voltstack_cluster.site_local_network.new_network](resources--gcp_vpc_site--reference--group-004.md#canonical-c9a5d044d970e4e003509bbbf603677b43b47f66d960d431eacebc31041e1e9b)
- [voltstack_cluster.site_local_network.new_network_autogenerate](resources--gcp_vpc_site--reference--group-004.md#canonical-a924dc90b4b4f595bffd49e0e6d5b7e185a6d23fdb692ea540f31ae238555d5d)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-925f7925df9001ec9bded593d287cadfedbdc0583c21d790af5435b955be4cb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ac0c2e309b902a4a2ecb4e7313995ef7c8274820d0cd0993239a73c85021818"></a>

## voltstack_cluster.site_local_network.existing_network — voltstack_cluster.site_local_network.existing_network / 936a7efbf197 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.site_local_network](resources--gcp_vpc_site--reference--group-004.md#canonical-0d117f633e2656829fba928f7b9b2a047ea066556f60ce1d03c0e04f47f29fa0)
- voltstack_cluster.site_local_network.existing_network

<a id="canonical-7305db8ef2592e9ff57af0018db8418a678b6b908cd9191c20d663903d4a3202"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for existing network.

Upstream description:

Name of existing VPC network.

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
  },
  "x-ves-oneof-field-routing_type": "[]"
}
```

Terraform syntax:

```terraform
existing_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-1b6aa6a05ce0007b4ef36a157bba14e5bb4ac653ac2f26b6b06634094785d2d9"></a>

## Direct properties — voltstack_cluster.site_local_network.existing_network / 936a7efbf197 / 3

<a id="canonical-c4c9d13b1bd0ad5af58d1a5bf6e0a8835dde3bc01e1ef67b0234981ad7235716"></a>

<a id="canonical-14934430673e8007576b72c07a9428667b0a92e77777584989572accfeaa409f"></a>

## name property — voltstack_cluster.site_local_network.existing_network / 936a7efbf197 / 4

Type: `"string"`. Optional.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

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
    "maxLength": 64,
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

<a id="canonical-ba48949d1d948602572ded2c9a4bf50d522722a9dd5ce1401b088d3c250dc2ce"></a>

## Next pages — voltstack_cluster.site_local_network.existing_network / 936a7efbf197 / 5

- [voltstack_cluster.site_local_network](resources--gcp_vpc_site--reference--group-004.md#canonical-0d117f633e2656829fba928f7b9b2a047ea066556f60ce1d03c0e04f47f29fa0)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-c9a5d044d970e4e003509bbbf603677b43b47f66d960d431eacebc31041e1e9b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ebd65bbb5de1db2ec0c642293ce1e8abf10fdafd9d76f265417a23fb9920b58c"></a>

## voltstack_cluster.site_local_network.new_network — voltstack_cluster.site_local_network.new_network / eb5537803f09 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.site_local_network](resources--gcp_vpc_site--reference--group-004.md#canonical-0d117f633e2656829fba928f7b9b2a047ea066556f60ce1d03c0e04f47f29fa0)
- voltstack_cluster.site_local_network.new_network

<a id="canonical-23585c10aea578ab42f24fa3bb7624c1a9e61287606e1a56e1c8c1fd34f12e15"></a>

Type: `"object"`. single nested block, Optional.

Parameters to create a new GCP VPC Network.

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
new_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-ef63f153e0b0af362d2f803423ca98b420f448e6df6efe6448f7f20d02981993"></a>

## Direct properties — voltstack_cluster.site_local_network.new_network / eb5537803f09 / 3

<a id="canonical-3dcaa675337d0ae63c322f91e21f872abc21b1d20eef0e78666bcb7ccc76d95b"></a>

<a id="canonical-209d1e196b94fa32cc973a63b8a98627a421d1189a3070a68954b504ecbefafa"></a>

## name property — voltstack_cluster.site_local_network.new_network / eb5537803f09 / 4

Type: `"string"`. Optional.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

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
    "maxLength": 64,
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

<a id="canonical-f4301191586c58cd64edcf0e7f7a68ea3468ceeff2649e0890c1c8291fe37526"></a>

## Next pages — voltstack_cluster.site_local_network.new_network / eb5537803f09 / 5

- [voltstack_cluster.site_local_network](resources--gcp_vpc_site--reference--group-004.md#canonical-0d117f633e2656829fba928f7b9b2a047ea066556f60ce1d03c0e04f47f29fa0)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-a924dc90b4b4f595bffd49e0e6d5b7e185a6d23fdb692ea540f31ae238555d5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6228df4ea1016947283266b5b1f366ee269a15315b4f29fccd2a7feccef3db7b"></a>

## voltstack_cluster.site_local_network.new_network_autogenerate — voltstack_cluster.site_local_network.new_network_autogenerate / d784a5d66d3e / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.site_local_network](resources--gcp_vpc_site--reference--group-004.md#canonical-0d117f633e2656829fba928f7b9b2a047ea066556f60ce1d03c0e04f47f29fa0)
- voltstack_cluster.site_local_network.new_network_autogenerate

<a id="canonical-03c01835d37a0e51f6c25972048c9fd54478b62f81073f74f1c15ebd3d237bb0"></a>

Type: `["object", {}]`. Optional.

Create a new GCP VPC Network with autogenerated name.

Receipt-pinned upstream constraints:

```json
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
new_network_autogenerate = {}
```

<a id="canonical-66e79d489cdc6d57eef7e519d3180c5cae5b5861237ff945735e960b81c53340"></a>

## Direct properties — voltstack_cluster.site_local_network.new_network_autogenerate / d784a5d66d3e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7d5e1b7a5128a52ef018c356872390021be3a08feb616cdea7bc428b9b1fdaa0"></a>

## Next pages — voltstack_cluster.site_local_network.new_network_autogenerate / d784a5d66d3e / 4

- [voltstack_cluster.site_local_network](resources--gcp_vpc_site--reference--group-004.md#canonical-0d117f633e2656829fba928f7b9b2a047ea066556f60ce1d03c0e04f47f29fa0)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-447430b7a7fb6788106f2e09e8f5fd7089adab0de64c47852e87c1abcc56deaf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-468dd971308dd665c214e576eb8ffa1cf6b381a5f9abe088660e0d213e3a3dd1"></a>

## voltstack_cluster.site_local_subnet — voltstack_cluster.site_local_subnet / 87002f869f86 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.site_local_subnet

<a id="canonical-56bf4aadfb600193fd1c6b1684f3e5587e30da2af67d6cd36085ff749220967d"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet",
    "new_subnet")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet\",\"new_subnet\"]"
}
```

Terraform syntax:

```terraform
site_local_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-f40a727dff497e4215e4b9b9eeeb9585c1471d6e037e7914b7d1415251735089"></a>

## Direct properties — voltstack_cluster.site_local_subnet / 87002f869f86 / 3

- [existing_subnet](resources--gcp_vpc_site--reference--group-004.md#canonical-75db298e4cc974e72fc337a173aba68e6bacbc3b00fe460039f90af4eef954bf): complete subsection reference.

- [new_subnet](resources--gcp_vpc_site--reference--group-004.md#canonical-c896e0a384c8a65b85fe567e41bc6812ecdad36793ef1400d65fcfeb3a6eafb4): complete subsection reference.

<a id="canonical-b2c86ebce4332170de41597594203bf2cdb8cf6acba6d440d186fe8b8c85cc43"></a>

## Next pages — voltstack_cluster.site_local_subnet / 87002f869f86 / 4

- [voltstack_cluster.site_local_subnet.existing_subnet](resources--gcp_vpc_site--reference--group-004.md#canonical-75db298e4cc974e72fc337a173aba68e6bacbc3b00fe460039f90af4eef954bf)
- [voltstack_cluster.site_local_subnet.new_subnet](resources--gcp_vpc_site--reference--group-004.md#canonical-c896e0a384c8a65b85fe567e41bc6812ecdad36793ef1400d65fcfeb3a6eafb4)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-75db298e4cc974e72fc337a173aba68e6bacbc3b00fe460039f90af4eef954bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-823c9a1ac2dae94a739844b19acd9db972151837d70ce5b423890c62ecf99d7f"></a>

## voltstack_cluster.site_local_subnet.existing_subnet — voltstack_cluster.site_local_subnet.existing_subnet / 6ce69e2beaf5 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.site_local_subnet](resources--gcp_vpc_site--reference--group-004.md#canonical-447430b7a7fb6788106f2e09e8f5fd7089adab0de64c47852e87c1abcc56deaf)
- voltstack_cluster.site_local_subnet.existing_subnet

<a id="canonical-049f89ec6068a6ca074106fd52159456f1732fe10b8d877293a7144745f1f76c"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for existing subnet.

Upstream description:

Name of existing GCP subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnet_name")}
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
existing_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-e3b0a4eb94b7d7ca0e7790093f5232ad6f269d16ccc7528c29edf0f4eea479de"></a>

## Direct properties — voltstack_cluster.site_local_subnet.existing_subnet / 6ce69e2beaf5 / 3

<a id="canonical-949c0e46fa330d562bb170c3dbbbcd842edb8a6f4ee75fed653b30d9433d9c1b"></a>

<a id="canonical-c42077ef42ba2ce394a64c5d155d17f00903daa06b25c564b7d58b5ad32096fb"></a>

## subnet_name property — voltstack_cluster.site_local_subnet.existing_subnet / 6ce69e2beaf5 / 4

Type: `"string"`. Optional.

VPC Subnet Name. Name of your subnet in VPC network.

Upstream description:

Name of your subnet in VPC network.

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

<a id="canonical-03f49bb19f5b035a4e0554d28515f5e5c81300c59448b3f229b471c96c0086dc"></a>

## Next pages — voltstack_cluster.site_local_subnet.existing_subnet / 6ce69e2beaf5 / 5

- [voltstack_cluster.site_local_subnet](resources--gcp_vpc_site--reference--group-004.md#canonical-447430b7a7fb6788106f2e09e8f5fd7089adab0de64c47852e87c1abcc56deaf)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-c896e0a384c8a65b85fe567e41bc6812ecdad36793ef1400d65fcfeb3a6eafb4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33e61abbed2578386ff5cda92b07fe52a3b3e6daf25f131037d105431ff4ab7a"></a>

## voltstack_cluster.site_local_subnet.new_subnet — voltstack_cluster.site_local_subnet.new_subnet / e75e753db7c6 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.site_local_subnet](resources--gcp_vpc_site--reference--group-004.md#canonical-447430b7a7fb6788106f2e09e8f5fd7089adab0de64c47852e87c1abcc56deaf)
- voltstack_cluster.site_local_subnet.new_subnet

<a id="canonical-bf8d4e45e95e8717eb0d02218c7161a2e671c5e86155d18a96ca0d3cacdd9080"></a>

Type: `"object"`. single nested block, Optional.

GCP subnet parameters Type. Parameters for GCP subnet.

Upstream description:

Parameters for GCP subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("primary_ipv4")}
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
new_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-ef3fbab97b306bc9d335da3e2fd85d7b3033f66edcaf0f1093d5bcaf0726eac3"></a>

## Direct properties — voltstack_cluster.site_local_subnet.new_subnet / e75e753db7c6 / 3

<a id="canonical-19dd63bb4c1963a3e54a4cf48ffbb4d5c88d71961ab209b076060cd8e609ed42"></a>

<a id="canonical-8a14a410c582a249af9be5c4e66c4909efad2b372b104d08dfb3fa815ca776c6"></a>

## primary_ipv4 property — voltstack_cluster.site_local_subnet.new_subnet / e75e753db7c6 / 4

Type: `"string"`. Optional.

IPv4 prefix for this Subnet. It has to be private address space.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "8"
  }
}
```

<a id="canonical-4afd6009eb78ac4ca741fe9a8e6540833ba741e7144417827ab98326b12b06c6"></a>

<a id="canonical-93498d83f99e45ff193ed5a56e22b2e4ffc203a7684b56a6b7e7d9ec53958b15"></a>

## subnet_name property — voltstack_cluster.site_local_subnet.new_subnet / e75e753db7c6 / 5

Type: `"string"`. Optional.

Name of new VPC Subnet, will be autogenerated if empty.

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

<a id="canonical-da9ca5f1f11ae59a811750e1ece6d315c51b0904af2a9869e04ad973b096ab66"></a>

## Next pages — voltstack_cluster.site_local_subnet.new_subnet / e75e753db7c6 / 6

- [voltstack_cluster.site_local_subnet](resources--gcp_vpc_site--reference--group-004.md#canonical-447430b7a7fb6788106f2e09e8f5fd7089adab0de64c47852e87c1abcc56deaf)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-94dfc3b3d4fea661a8cea3b72218841e075fbd02a96681d4d90e3f4604ba3e4a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7bf0b5a3854511d4df3b349371f653790917aeefaf2a2c7dde5b18583cd17530"></a>

## voltstack_cluster.sm_connection_public_ip — voltstack_cluster.sm_connection_public_ip / 394d61b9cc39 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.sm_connection_public_ip

<a id="canonical-7c2d345d7ba13b8cfbc44093a4f28e5f94367fc74436044ae23d5132e3d4e52b"></a>

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
sm_connection_public_ip = {}
```

<a id="canonical-8b23ea19a476472d3145f0ae68346c0610ef8e775f738cfc2dcf23df8042d342"></a>

## Direct properties — voltstack_cluster.sm_connection_public_ip / 394d61b9cc39 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-94d0b797c99a1a5ba4bfcff3194fda6564a41802745444c78f8d9cf2e85e3313"></a>

## Next pages — voltstack_cluster.sm_connection_public_ip / 394d61b9cc39 / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-3d45010f2a154d078ec9f6f6287c97b091b78fc80801589773e9e5cc9e1ebbf8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80018ba2ac29a27e3813440df01fab63c3714695569438ce85e09354bb7c96e8"></a>

## voltstack_cluster.sm_connection_pvt_ip — voltstack_cluster.sm_connection_pvt_ip / 5cee33928857 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.sm_connection_pvt_ip

<a id="canonical-01b61b5f410628eff55ddb1d6934093db1b36bc1d23d4b9e3d157e8a65175b42"></a>

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
sm_connection_pvt_ip = {}
```

<a id="canonical-7712943023249a813cba54ab94d00e2dc715d950d64ceb83d3d4d2e93604ada7"></a>

## Direct properties — voltstack_cluster.sm_connection_pvt_ip / 5cee33928857 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5baa4f596fc74bdcd7c6cd39ea2dc6fb20c81353046ded39c353dd6c776442ae"></a>

## Next pages — voltstack_cluster.sm_connection_pvt_ip / 5cee33928857 / 4

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-56ffa539204782f5f466a7bc0f01f7b06e0be7bd95c35b7f6529e6d577be27d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0dc4b8d08a503d642ba30b5040c60824a93782564879407aeae1f5afe5eb774e"></a>

## voltstack_cluster.storage_class_list — voltstack_cluster.storage_class_list / 4d84201b0590 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- voltstack_cluster.storage_class_list

<a id="canonical-8658a534e6cd77f480f217f65542f6488987b44f41724f719b6941b6079b6c39"></a>

Type: `"object"`. single nested block, Optional.

Add additional custom storage classes in Kubernetes for this site.

Receipt-pinned upstream constraints:

```json
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
storage_class_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-dab205b0e48f2a12ce1f3fe9194f78bf4f7b9a1c83a38ac97913638781a1260e"></a>

## Direct properties — voltstack_cluster.storage_class_list / 4d84201b0590 / 3

- [storage_classes](resources--gcp_vpc_site--reference--group-004.md#canonical-cf8160b0376bcadbd1fba6bede29b8e7e846c1fc6fe97cc450a423806b181e96): complete subsection reference.

<a id="canonical-ae54c8544fa3b003ed352e6369a0d18971ccabe416c8a590ae0f9294f050279b"></a>

## Next pages — voltstack_cluster.storage_class_list / 4d84201b0590 / 4

- [voltstack_cluster.storage_class_list.storage_classes](resources--gcp_vpc_site--reference--group-004.md#canonical-cf8160b0376bcadbd1fba6bede29b8e7e846c1fc6fe97cc450a423806b181e96)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-cf8160b0376bcadbd1fba6bede29b8e7e846c1fc6fe97cc450a423806b181e96"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c71558f028b94e8c7cae4977086b029732afab1a42631c6906ca595bafaf11ee"></a>

## voltstack_cluster.storage_class_list.storage_classes — voltstack_cluster.storage_class_list.storage_classes / 8a4645be96e9 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [voltstack_cluster.storage_class_list](resources--gcp_vpc_site--reference--group-004.md#canonical-56ffa539204782f5f466a7bc0f01f7b06e0be7bd95c35b7f6529e6d577be27d9)
- voltstack_cluster.storage_class_list.storage_classes

<a id="canonical-6fd7b37cd5297ba229262cf9a4c1a96d0c2fd2bfdd617882da061f2d58604583"></a>

Type: `"object"`. list nested block, Optional.

List of Storage Classes. List of custom storage classes.

Upstream description:

List of custom storage classes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("storage_class_name")}
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
