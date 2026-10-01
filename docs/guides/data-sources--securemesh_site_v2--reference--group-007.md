---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-16d4b0a1511e6575516df474bc8a5d015da785e5aae550e5cc4aa638379f9329"></a>

## deployment_size property — eks_k8s / 426f599ff496 / 4

Type: `"string"`. Computed.

\[Enum: KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM|KUBERNETES\_DEPLOYMENT\_SIZE\_LARGE\] Enum for
Kubernetes deployment size OPTIONS - KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM: Medium Medium deployment
size with moderate resource requirements (8 vCPU, 32 GB memory). Suitable for most deployments. -
KUBERNETES\_DEPLOYMENT\_SIZE\_LARGE: Large Large deployment size with higher resource.. Possible
values are \`KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM\`, \`KUBERNETES\_DEPLOYMENT\_SIZE\_LARGE\`.
Defaults to \`KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM\`.

Upstream description:

Enum for Kubernetes deployment size OPTIONS

&#8203;- KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM: Medium

Medium deployment size with moderate resource requirements (8 vCPU, 32 GB memory). Suitable for most
deployments. &#8203;- KUBERNETES\_DEPLOYMENT\_SIZE\_LARGE: Large

Large deployment size with higher resource requirements (16 vCPU, 64 GB memory) for demanding
workloads requiring additional performance and capacity.

Receipt-pinned upstream constraints:

```json
{
  "default": "KUBERNETES_DEPLOYMENT_SIZE_MEDIUM",
  "enum": [
    "KUBERNETES_DEPLOYMENT_SIZE_MEDIUM",
    "KUBERNETES_DEPLOYMENT_SIZE_LARGE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [disable_anti_affinity](data-sources--securemesh_site_v2--reference--group-007.md#canonical-28b37aa6c2222640058d78d6a938c24fd527e804cbd0fb6cc2fadbf5836c5b1f): complete subsection reference.

- [enable_anti_affinity](data-sources--securemesh_site_v2--reference--group-007.md#canonical-de1dde4c2cb55efc147e488c4fc6678e0d95a47a80bcf1382f7a52693a2ec660): complete subsection reference.

<a id="canonical-5abcb5be4bdc1bc0a5b8553babdd19638182c1d126a57a7772ad4bb916959707"></a>

<a id="canonical-b7f593747218f2c7ef9a05f511e6aee313efa0edbb3c5be37dd56d6cacc240ac"></a>

## labels property — eks_k8s / 426f599ff496 / 5

Type: `["map", "string"]`. Computed.

Add labels to control which Kubernetes nodes the VPM and related pods (etcd, VER, prometheus) are
deployed to. Specify label key-value pairs that match the labels on your Kubernetes nodes. This uses
Kubernetes nodeSelector to schedule pods only on nodes with matching labels.

Upstream description:

Add labels to control which Kubernetes nodes the VPM and related pods (etcd, VER, prometheus) are
deployed to. Specify label key-value pairs that match the labels on your Kubernetes nodes. This uses
Kubernetes nodeSelector to schedule pods only on nodes with matching labels.

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
    "ves.io.schema.rules.map.keys.string.max_len": "253",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "63",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "253",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "63",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9): complete subsection reference.

<a id="canonical-dcb2a84fe70fc2eed5e6d40ffedc575e649c85c2ea047c41dca507a0cc17982e"></a>

## Next pages — eks_k8s / 426f599ff496 / 6

- [eks_k8s.disable_anti_affinity](data-sources--securemesh_site_v2--reference--group-007.md#canonical-28b37aa6c2222640058d78d6a938c24fd527e804cbd0fb6cc2fadbf5836c5b1f)
- [eks_k8s.enable_anti_affinity](data-sources--securemesh_site_v2--reference--group-007.md#canonical-de1dde4c2cb55efc147e488c4fc6678e0d95a47a80bcf1382f7a52693a2ec660)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-28b37aa6c2222640058d78d6a938c24fd527e804cbd0fb6cc2fadbf5836c5b1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59d7defdf25f343b12fb10596e038ce6bba4254680e2a587fa7721b811f8c05c"></a>

## eks_k8s.disable_anti_affinity — eks_k8s.disable_anti_affinity / 0b159fe1cf24 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- eks_k8s.disable_anti_affinity

<a id="canonical-87cedcd38c62e6d148642846388ad8053483d3968a8553b3e262695764d7b59e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable anti affinity.

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

<a id="canonical-cef63074f571f12f92039424033ee0ee2a05901276eefb366bb22d56b7ed2c46"></a>

## Direct properties — eks_k8s.disable_anti_affinity / 0b159fe1cf24 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ea6d10ef214edc0d9f5f1a423962d254a7253764169f38f44907841dfefccbc1"></a>

## Next pages — eks_k8s.disable_anti_affinity / 0b159fe1cf24 / 4

- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-de1dde4c2cb55efc147e488c4fc6678e0d95a47a80bcf1382f7a52693a2ec660"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ff515034dcdb80f2b02b2c5b27425aecb7a126367f279def410986ff0811f74"></a>

## eks_k8s.enable_anti_affinity — eks_k8s.enable_anti_affinity / d3c056ebba74 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- eks_k8s.enable_anti_affinity

<a id="canonical-a51909734c0cd98e80d8703ed7d41fb1cc6d6b705ff716486fd11536f8a07c70"></a>

Type: `"single"`. Computed.

Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different
applications/components are distributed across your Kubernetes cluster.

Upstream description:

Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different
applications/components are distributed across your Kubernetes cluster.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-54600e1070566751e059eca329678a74cc2120fb6858f2503432ba08f2e5791d"></a>

## Direct properties — eks_k8s.enable_anti_affinity / d3c056ebba74 / 3

- [rules](data-sources--securemesh_site_v2--reference--group-007.md#canonical-ede70ee853a1a8b3f6c30f9b8037245cb60bf03ddf06fc998e6d64ab129338a6): complete subsection reference.

<a id="canonical-87bf9a847d2e3dbf95eb37a878d03fe3f233bd4655eadbe05fda58475438621e"></a>

## Next pages — eks_k8s.enable_anti_affinity / d3c056ebba74 / 4

- [eks_k8s.enable_anti_affinity.rules](data-sources--securemesh_site_v2--reference--group-007.md#canonical-ede70ee853a1a8b3f6c30f9b8037245cb60bf03ddf06fc998e6d64ab129338a6)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-ede70ee853a1a8b3f6c30f9b8037245cb60bf03ddf06fc998e6d64ab129338a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6bea255260e039d126844e83c2023d19c65555f0f3b7b3bee769d53915deaea8"></a>

## eks_k8s.enable_anti_affinity.rules — eks_k8s.enable_anti_affinity.rules / 2759fb41d419 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.enable_anti_affinity](data-sources--securemesh_site_v2--reference--group-007.md#canonical-de1dde4c2cb55efc147e488c4fc6678e0d95a47a80bcf1382f7a52693a2ec660)
- eks_k8s.enable_anti_affinity.rules

<a id="canonical-407c9c935c16a2ad9e59b893aba95d2417360234f101e404fc9531d2a058d2c6"></a>

Type: `"list"`. Computed.

Define one or more anti-affinity rules. Each rule specifies which pods (by labels) should be
distributed across which topology domains.

Upstream description:

Define one or more anti-affinity rules. Each rule specifies which pods (by labels) should be
distributed across which topology domains. Example: Rule 1 - Distribute VPM pods across nodes, Rule
2 - Distribute Prometheus pods across zones.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
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
    "ves.io.schema.rules.repeated.max_items": "20",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "20",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3412c6c1589d2a516845d38587903353e9c9b33105df912f1513ddb9db29665f"></a>

## Direct properties — eks_k8s.enable_anti_affinity.rules / 2759fb41d419 / 3

<a id="canonical-f80dcfb1cb6beffb6ce82b589f8f50a1d51b0c707f1c5333fc0d7513768a1527"></a>

<a id="canonical-6625e250b16b1af899aa89c52f29f2c7038321cb8334e69b76b56d83633e4177"></a>

## label_key property — eks_k8s.enable_anti_affinity.rules / 2759fb41d419 / 4

Type: `"string"`. Computed.

Specify the label key of the customer pods that CE pods should avoid being co-scheduled with.
Combined with the label value below, this identifies the target pods.

Upstream description:

Specify the label key of the customer pods that CE pods should avoid being co-scheduled with.
Combined with the label value below, this identifies the target pods.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 253,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 253,
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
    "ves.io.schema.rules.string.max_len": "253",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "253",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-b03e442958f1416a01c504025f0bf29d0ad1c62aaaabaaf424c69f5d11a5c452"></a>

<a id="canonical-4c55e573a76afad268e9affb39eba2c45750b016e3b01774e065a4e5abe0a4ed"></a>

## label_value property — eks_k8s.enable_anti_affinity.rules / 2759fb41d419 / 5

Type: `"string"`. Computed.

Specify the label value that, together with the label key, identifies the customer pods to avoid.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 63,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 63,
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
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3ff704e95cb18009f5c5ab174e71e22c006b389fcd44fcc0fe69ac1acac9a1e7"></a>

<a id="canonical-796bb62853202498734060475c4bc9309d5daff9888e66080509adcce547bfed"></a>

## topology_keys property — eks_k8s.enable_anti_affinity.rules / 2759fb41d419 / 6

Type: `["list", "string"]`. Computed.

Specify one or more node label keys that define the scope of avoidance. For each topology key (e.g.,
Kubernetes.I/O/hostname), CE pods will avoid nodes whose topology value matches a node already
running a pod with the above specified label.

Upstream description:

Specify one or more node label keys that define the scope of avoidance. For each topology key (e.g.,
Kubernetes.I/O/hostname), CE pods will avoid nodes whose topology value matches a node already
running a pod with the above specified label. Example: with Kubernetes.I/O/hostname, CE pods are
kept off any node running the matching customer pod.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "253",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "253",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-32693bac6034287e8a9be5ccea1485512e3603e82f5a8f17221105a5d187eb2c"></a>

## Next pages — eks_k8s.enable_anti_affinity.rules / 2759fb41d419 / 7

- [eks_k8s.enable_anti_affinity](data-sources--securemesh_site_v2--reference--group-007.md#canonical-de1dde4c2cb55efc147e488c4fc6678e0d95a47a80bcf1382f7a52693a2ec660)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e22db97db6152f8c8b24a7219a8f1eb727eed21ded2e35dc127b08e7700de51b"></a>

## eks_k8s.not_managed — eks_k8s.not_managed / 6963e906f0f9 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- eks_k8s.not_managed

<a id="canonical-0fd6caf2aa46c4976adfa041a50366bb86ea0a467cb1ae9c6f19250f4d892fa5"></a>

Type: `"single"`. Computed.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-a4a57fffc50b23cc02ef35f154c7f710b4438d44bd6deade0a1d3eddaf51d15a"></a>

## Direct properties — eks_k8s.not_managed / 6963e906f0f9 / 3

- [node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044): complete subsection reference.

<a id="canonical-857e1d35cde0f72c9602812d767c850e89cdbeeb743cd57833e24776e737cde8"></a>

## Next pages — eks_k8s.not_managed / 6963e906f0f9 / 4

- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-785d707b007a9b55831f17700b1f74650f285c47172a42f0063bca5ff43b19c2"></a>

## eks_k8s.not_managed.node_list — eks_k8s.not_managed.node_list / bc048a3184f5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- eks_k8s.not_managed.node_list

<a id="canonical-0d85cc4c7f0d083d2229f22d270293df8fe967e0f70eba87de6369aa8bee0a15"></a>

Type: `"list"`. Computed.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-c444e284f4527f628e85bf7824991f63e9885349e71c2de1d2cdf4503ed89072"></a>

## Direct properties — eks_k8s.not_managed.node_list / bc048a3184f5 / 3

<a id="canonical-5714f3d41a90946fd9220f9ad5e8d20287486fa79f734b6abcc5d8a117019cbc"></a>

<a id="canonical-c17c07e44fdf4f41cba366d69483a1c81ca2d0e5368766853da183b1984307ef"></a>

## hostname property — eks_k8s.not_managed.node_list / bc048a3184f5 / 4

Type: `"string"`. Computed.

Hostname. Hostname for this Node.

Upstream description:

Hostname for this Node.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

- [interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557): complete subsection reference.

<a id="canonical-3229d6d7afd3b3d1f9cb241873ea197b9d193dc8d3d422981c9ac67f162842aa"></a>

<a id="canonical-c72d2ea9f641b3999f60a072218c7989e4b9fb9f82645de90eef315cf3f838dc"></a>

## public_ip property — eks_k8s.not_managed.node_list / bc048a3184f5 / 5

Type: `"string"`. Computed.

Public IP. Public IP for this Node.

Upstream description:

Public IP for this Node.

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

<a id="canonical-2643d16c6216aca2cab468c41e9af53c6efdb20a74ea7ec260d68720e53629ec"></a>

<a id="canonical-01f8ff40032b7342aeb84c5c0b5a982fc9e59e523d90c5b837c8220bdcf8e664"></a>

## type property — eks_k8s.not_managed.node_list / bc048a3184f5 / 6

Type: `"string"`. Computed.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Upstream description:

Type for this Node, can be Control or Worker.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "Control",
    "Worker"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-1177d205ac77e8b40e5a782ff6ff40287280917b2b3b728dd72c4add63146235"></a>

## Next pages — eks_k8s.not_managed.node_list / bc048a3184f5 / 7

- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-403699730291d4a81ebd227dd88c45b7b6e1f9d231ac2878053cab81fab05860"></a>

## eks_k8s.not_managed.node_list.interface_list — eks_k8s.not_managed.node_list.interface_list / 1dc589f82877 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- eks_k8s.not_managed.node_list.interface_list

<a id="canonical-59585fc19f102a0ef036a2ed5fc44b33f917d4566d7c58d98ec905fc6404dcf7"></a>

Type: `"list"`. Computed.

Manage interfaces belonging to this node.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-851b008612b166a17b7f946d89c7f7beb0bd7fd85f0ef54c069004b9eb4c45a6"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list / 1dc589f82877 / 3

- [bond_interface](data-sources--securemesh_site_v2--reference--group-007.md#canonical-0f7fa3669982ae7642759ecb4e027ab0826fe388cbd6cebc0fa6828881f8f3b7): complete subsection reference.

<a id="canonical-60228a5a16c81e67a77ef5f3f30f75cc03e4a25f82a200503bb437ab971a677e"></a>

<a id="canonical-33d3edec264b726ff5ea61e6f7eeb0e3fb52bdb63b1b2ea8cfde5fc16789c696"></a>

## description_spec property — eks_k8s.not_managed.node_list.interface_list / 1dc589f82877 / 4

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [dhcp_client](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c54dea6836d38e78b71d246d937f9f3b37fced2f613eac1de78bc7ac05df8f43): complete subsection reference.

- [dhcp_server](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c4c1f243d85ccb664e50a678e8000889fea55c0fd9a37a19ad6f87eea1d5ba60): complete subsection reference.

- [ethernet_interface](data-sources--securemesh_site_v2--reference--group-007.md#canonical-5d240ca2a5fd383b44dcc9675e16734cfa7e73f69641f9d2b44b6beec8c3d4de): complete subsection reference.

- [ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8958acd121a1de591d9c295a38c154a6b85f9e6aa7cdaaf4b4bed2d5300a15d1): complete subsection reference.

<a id="canonical-430583ce10350f491a86144dfceac9d73e4ac34a6f90a88cb317cb3211055f4f"></a>

<a id="canonical-b8d630af9ad5e2d569d2bf7acde8545c1ae7e771715b9025ad0b8ff20948c624"></a>

## is_management property — eks_k8s.not_managed.node_list.interface_list / 1dc589f82877 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-1fd3155252699048ba5195b919504dffb95f3f4f2f88f2258387a8ea98ea6d59"></a>

<a id="canonical-62b963094990cd61a48c5ee5469c4ac65feb3c6bb3031490cfb2bf1d98510bc2"></a>

## is_primary property — eks_k8s.not_managed.node_list.interface_list / 1dc589f82877 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-e9979e596a3a632088d6fd86684034735ba706751f00aaaa5aa16899bf572748"></a>

<a id="canonical-3905d3772ac308ad4d1d747d184df0cb69e23ef21a8227dc0192e3b008c059b4"></a>

## labels property — eks_k8s.not_managed.node_list.interface_list / 1dc589f82877 / 7

Type: `["map", "string"]`. Computed.

Add Labels for this Interface, these labels can be used in firewall policy.

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
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [monitor](data-sources--securemesh_site_v2--reference--group-007.md#canonical-89f31d3468d8e20d7b5b0430ae467c78c9725f45cd28956f981826e50d35d2a1): complete subsection reference.

- [monitor_disabled](data-sources--securemesh_site_v2--reference--group-007.md#canonical-156462b8756a9a67940b9725de36116b6fee6831d4f58a84c88444b248c32aaa): complete subsection reference.

<a id="canonical-6c184f329c7f661389577c345bff7a8e7f7a4322abd86c716222e7c5ed2a35b6"></a>

<a id="canonical-94a14a06fd8dbe78bfc2c3b55420b0f155cc206251651704e783e3faad0fb7cb"></a>

## mtu property — eks_k8s.not_managed.node_list.interface_list / 1dc589f82877 / 8

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-4a7c3f96b883a8f41c72dc6281edf833e59eeac42e934d4e30c09946c59de060"></a>

<a id="canonical-64b7185937b1d7a564f46c8b6356712211cabe90f24411e182c51eb2f0b6680b"></a>

## name property — eks_k8s.not_managed.node_list.interface_list / 1dc589f82877 / 9

Type: `"string"`. Computed.

Interface Name. Name of this Interface.

Upstream description:

Name of this Interface.

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

- [network_option](data-sources--securemesh_site_v2--reference--group-007.md#canonical-92d1d61a8e50c720846a9a9dc9ff1f807d88b90899b7d46897205fe2979c50f0): complete subsection reference.

- [no_ipv4_address](data-sources--securemesh_site_v2--reference--group-007.md#canonical-d21ed0b715126b4a844e74e1c1f79b73b2e970771cdcf8963ac37a0565a14660): complete subsection reference.

- [no_ipv6_address](data-sources--securemesh_site_v2--reference--group-007.md#canonical-e1b16b35e297632d2254d99a4b53cc29cde056668d8537ad46283edb252d5bb6): complete subsection reference.

<a id="canonical-c42035337869e5da18fb2144fcf603d9fd623beae354f5cdcb3685617caea70e"></a>

<a id="canonical-cf7c5ec7a46bb0d842da2494abf9e22fb0f391d673e1e8cc840b2f5ada9e43da"></a>

## priority property — eks_k8s.not_managed.node_list.interface_list / 1dc589f82877 / 10

Type: `"number"`. Computed.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Upstream description:

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
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
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-007.md#canonical-a48c3d0552e14088be10a64e10836db4ba356566a3b81fbb55238187eb82d128): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-007.md#canonical-75f691788d847ec3be5898614a862e1c1d35cd4ead0dc9c7b666d8e1338250ef): complete subsection reference.

- [static_ip](data-sources--securemesh_site_v2--reference--group-007.md#canonical-082c90448b4143a96a03cda9f9bc496d6b3ed9ea44783d38c1c9a90858b74488): complete subsection reference.

- [static_ipv6_address](data-sources--securemesh_site_v2--reference--group-008.md#canonical-6fbeb2b610cece275c75e13c9808134564c88d7778450503142d5d3f5e1bf285): complete subsection reference.

- [vlan_interface](data-sources--securemesh_site_v2--reference--group-008.md#canonical-d8ebbe80051ed8c3947c3a6d56a9cead9a4403c159fd18dde4e24fcd07a9a7b3): complete subsection reference.

<a id="canonical-ddf8c4fce087013bf67673ad91d38bff104fbd5df79f71b047460029879f8d1d"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list / 1dc589f82877 / 11

- [eks_k8s.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-007.md#canonical-0f7fa3669982ae7642759ecb4e027ab0826fe388cbd6cebc0fa6828881f8f3b7)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_client](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c54dea6836d38e78b71d246d937f9f3b37fced2f613eac1de78bc7ac05df8f43)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c4c1f243d85ccb664e50a678e8000889fea55c0fd9a37a19ad6f87eea1d5ba60)
- [eks_k8s.not_managed.node_list.interface_list.ethernet_interface](data-sources--securemesh_site_v2--reference--group-007.md#canonical-5d240ca2a5fd383b44dcc9675e16734cfa7e73f69641f9d2b44b6beec8c3d4de)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8958acd121a1de591d9c295a38c154a6b85f9e6aa7cdaaf4b4bed2d5300a15d1)
- [eks_k8s.not_managed.node_list.interface_list.monitor](data-sources--securemesh_site_v2--reference--group-007.md#canonical-89f31d3468d8e20d7b5b0430ae467c78c9725f45cd28956f981826e50d35d2a1)
- [eks_k8s.not_managed.node_list.interface_list.monitor_disabled](data-sources--securemesh_site_v2--reference--group-007.md#canonical-156462b8756a9a67940b9725de36116b6fee6831d4f58a84c88444b248c32aaa)
- [eks_k8s.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-007.md#canonical-92d1d61a8e50c720846a9a9dc9ff1f807d88b90899b7d46897205fe2979c50f0)
- [eks_k8s.not_managed.node_list.interface_list.no_ipv4_address](data-sources--securemesh_site_v2--reference--group-007.md#canonical-d21ed0b715126b4a844e74e1c1f79b73b2e970771cdcf8963ac37a0565a14660)
- [eks_k8s.not_managed.node_list.interface_list.no_ipv6_address](data-sources--securemesh_site_v2--reference--group-007.md#canonical-e1b16b35e297632d2254d99a4b53cc29cde056668d8537ad46283edb252d5bb6)
- [eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-007.md#canonical-a48c3d0552e14088be10a64e10836db4ba356566a3b81fbb55238187eb82d128)
- [eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-007.md#canonical-75f691788d847ec3be5898614a862e1c1d35cd4ead0dc9c7b666d8e1338250ef)
- [eks_k8s.not_managed.node_list.interface_list.static_ip](data-sources--securemesh_site_v2--reference--group-007.md#canonical-082c90448b4143a96a03cda9f9bc496d6b3ed9ea44783d38c1c9a90858b74488)
- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-008.md#canonical-6fbeb2b610cece275c75e13c9808134564c88d7778450503142d5d3f5e1bf285)
- [eks_k8s.not_managed.node_list.interface_list.vlan_interface](data-sources--securemesh_site_v2--reference--group-008.md#canonical-d8ebbe80051ed8c3947c3a6d56a9cead9a4403c159fd18dde4e24fcd07a9a7b3)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0f7fa3669982ae7642759ecb4e027ab0826fe388cbd6cebc0fa6828881f8f3b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b344302c83a2bbb55e4fefb3de20162298315540f1cf9474bf0533661ebfa9b"></a>

## eks_k8s.not_managed.node_list.interface_list.bond_interface — eks_k8s.not_managed.node_list.interface_list.bond_interface / 731832715543 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- eks_k8s.not_managed.node_list.interface_list.bond_interface

<a id="canonical-39929417b3847c9603f5c384d2d8f29d51ac0d78c20f428b6c7da5b8ce489f6c"></a>

Type: `"single"`. Computed.

Configuration parameter for bond interface.

Upstream description:

Bond devices configuration for fleet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-lacp_choice": "[\"active_backup\",\"lacp\"]"
}
```

<a id="canonical-6ec5b45b3f3ff4980dc78996966e7284878c54a5f8c268382c23120ea385a782"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.bond_interface / 731832715543 / 3

- [active_backup](data-sources--securemesh_site_v2--reference--group-007.md#canonical-0903b8a20bff1eef48ae7dce23f908f20472c677aacd4af18cdc16d47db39f6c): complete subsection reference.

<a id="canonical-40d831659525230967a7bbd80c744071fb56de395831162c92e97c886186ead4"></a>

<a id="canonical-008f54fb0a02bdf8acc17d481a99141427d4047f5b9e88e990b1630b9c1ba727"></a>

## devices property — eks_k8s.not_managed.node_list.interface_list.bond_interface / 731832715543 / 4

Type: `["list", "string"]`. Computed.

Ethernet devices that will make up this bond.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [lacp](data-sources--securemesh_site_v2--reference--group-007.md#canonical-0b95bd9095d11332b1995eac59d4b40f2ca18d91f21cafa82b40edbaf665ba3b): complete subsection reference.

<a id="canonical-a5be2709e267cf009c6343f3040677904fcb9de9ffadb5e859d804589a7dc98b"></a>

<a id="canonical-eb6c718d4fe9ca9aa7ec551cbb4ec7088de47ed8a021f071d6eef3edb38a908b"></a>

## link_polling_interval property — eks_k8s.not_managed.node_list.interface_list.bond_interface / 731832715543 / 5

Type: `"number"`. Computed.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 500
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-a476b0f1ddaac9eaa6f30aea5b9e3a270800ddd787166f49a106c13548558450"></a>

<a id="canonical-57428797f0374e68ca1b2e9c48a4ada1cb0dbb31c5ad1a01821f5130f000e584"></a>

## link_up_delay property — eks_k8s.not_managed.node_list.interface_list.bond_interface / 731832715543 / 6

Type: `"number"`. Computed.

Milliseconds wait before link is declared up.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  }
}
```

<a id="canonical-8c954ffa7f39a91040fc3a68505879db8ccbe0fff4d3405f53de2d2e7aa730bd"></a>

<a id="canonical-8d1caa6d51aa11842d956d953b2e6911d4d379b15f5de73adacbceb8b1082498"></a>

## name property — eks_k8s.not_managed.node_list.interface_list.bond_interface / 731832715543 / 7

Type: `"string"`. Computed.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-c007d60e16780ab04d49ecc0c969789609a90d7694750b8504b55294fd22a1b3"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.bond_interface / 731832715543 / 8

- [eks_k8s.not_managed.node_list.interface_list.bond_interface.active_backup](data-sources--securemesh_site_v2--reference--group-007.md#canonical-0903b8a20bff1eef48ae7dce23f908f20472c677aacd4af18cdc16d47db39f6c)
- [eks_k8s.not_managed.node_list.interface_list.bond_interface.lacp](data-sources--securemesh_site_v2--reference--group-007.md#canonical-0b95bd9095d11332b1995eac59d4b40f2ca18d91f21cafa82b40edbaf665ba3b)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0903b8a20bff1eef48ae7dce23f908f20472c677aacd4af18cdc16d47db39f6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb108cb8ad371064a39416190d05180bb7affb060c160cda2c690f9de5b578c5"></a>

## eks_k8s.not_managed.node_list.interface_list.bond_interface.active_backup — eks_k8s.not_managed.node_list.interface_list.bond_interface.active_backup / 67efd40bba6d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-007.md#canonical-0f7fa3669982ae7642759ecb4e027ab0826fe388cbd6cebc0fa6828881f8f3b7)
- eks_k8s.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-34904fde4e886790eca9a81a555e6d1af7e782453d4e925b872f62dd8ab50ae6"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for active backup.

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

<a id="canonical-fbc5b5242e7a1dd1c6dea39a5197e248833119351bc8da46094457abcdf5ef20"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.bond_interface.active_backup / 67efd40bba6d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f9916699f3a1c20ddaf3f62478f3e6ef240b88b5ae7c67cbd0a5ceb0bba1595c"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.bond_interface.active_backup / 67efd40bba6d / 4

- [eks_k8s.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-007.md#canonical-0f7fa3669982ae7642759ecb4e027ab0826fe388cbd6cebc0fa6828881f8f3b7)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0b95bd9095d11332b1995eac59d4b40f2ca18d91f21cafa82b40edbaf665ba3b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a61455cb82ab6c3de45ad1be544591d9c22d14fde41f130dabd4d80c328769f"></a>

## eks_k8s.not_managed.node_list.interface_list.bond_interface.lacp — eks_k8s.not_managed.node_list.interface_list.bond_interface.lacp / e3522d2b9062 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-007.md#canonical-0f7fa3669982ae7642759ecb4e027ab0826fe388cbd6cebc0fa6828881f8f3b7)
- eks_k8s.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-7e299ae1e807fbc908b70b9a527710ffc207df23575ade7ac8800788aae4bd02"></a>

Type: `"single"`. Computed.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-67c1b954fd0bbcb61876ff1b8a0eedd9af539257214b2730b1ece148c2d77246"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.bond_interface.lacp / e3522d2b9062 / 3

<a id="canonical-85d18079a6850ca6ccfb901d3e1590af78d194762d153d59d0e77e4c081b2ac6"></a>

<a id="canonical-f85595263517089776b182e7b0149833e3e8c2dbbf9bbf5addee2da0c5cb1039"></a>

## rate property — eks_k8s.not_managed.node_list.interface_list.bond_interface.lacp / e3522d2b9062 / 4

Type: `"number"`. Computed.

Interval in seconds to transmit LACP packets.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
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
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-3485cc550ba3c4fd8b09b33676f3114cf53f1b82ee7178b52b221f320edc8b44"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.bond_interface.lacp / e3522d2b9062 / 5

- [eks_k8s.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-007.md#canonical-0f7fa3669982ae7642759ecb4e027ab0826fe388cbd6cebc0fa6828881f8f3b7)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c54dea6836d38e78b71d246d937f9f3b37fced2f613eac1de78bc7ac05df8f43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2507ca480d429d5e6ff894fac63ae4356d9bbf6e168c60a59295a09595020bfb"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_client — eks_k8s.not_managed.node_list.interface_list.dhcp_client / 5ccfeea3020e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- eks_k8s.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-2a2c848d4e59ecde3953af9b4de7d3f767ac3181b0896421c95978135d47d7ae"></a>

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

<a id="canonical-f9d61ad6cf5d7c2d85d3613775bd03d378b080c004a0d30a41b86d25064064c0"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.dhcp_client / 5ccfeea3020e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-63bd02668663f2a65a752b2bac070d5a499b24a4d980918c5964e946d208eddf"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.dhcp_client / 5ccfeea3020e / 4

- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c4c1f243d85ccb664e50a678e8000889fea55c0fd9a37a19ad6f87eea1d5ba60"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5dc8b40c51fe7467c00d086ff5eefb70e824e32b15b1d5e9a437a7b38861d697"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_server — eks_k8s.not_managed.node_list.interface_list.dhcp_server / 3a3e005925af / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-8bd4ab3090fae6cfc5ff9c04471fc866a5d98f404d4c2d98a1eb5660345f23a7"></a>

Type: `"single"`. Computed.

DHCPServerParametersType.

Upstream description:

DHCP server configuration for this interface.

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

<a id="canonical-a52ca99d1a8fffcc19dcc929d5133e572597b0195c6ad69084e27bf82ce25604"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.dhcp_server / 3a3e005925af / 3

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60508a7c745a5740878d75273de529291316607c55130b66f6634f307f939e3a): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-007.md#canonical-26188b3f7819ae2648313aefc4e535174c439be8c073cb2323ffe688ddc32e09): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-007.md#canonical-9f806526dc3dd1159f56701dbcbcba74948ac1d6c895818ea0b811cfb6e7dff1): complete subsection reference.

<a id="canonical-4c7556327aadda5beed4c4d588288bab658740388ad04b4c2b854355c1f8db2f"></a>

<a id="canonical-e281b57e2eb1a88d0db222bd5080e80194981cda391b810daa934d486a468a12"></a>

## dhcp_option82_tag property — eks_k8s.not_managed.node_list.interface_list.dhcp_server / 3a3e005925af / 4

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-eeb07af851a394d833141c512a42d79c636616ca13bcefbf5509fe74ee04c89a"></a>

<a id="canonical-5031720560702c433fc00471aa4d95f8267a160494a45115fb5b8debed6aa9e4"></a>

## fixed_ip_map property — eks_k8s.not_managed.node_list.interface_list.dhcp_server / 3a3e005925af / 5

Type: `["map", "string"]`. Computed.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

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
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-007.md#canonical-7cd85f4ecba5e70219d237ad6728f98660d5ebc3137201592fe2feef576ec0f6): complete subsection reference.

<a id="canonical-ea038aeaa0d51405b782a8d8a64db418d164b27341a4fce4dfbdb60094994ce5"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.dhcp_server / 3a3e005925af / 6

- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60508a7c745a5740878d75273de529291316607c55130b66f6634f307f939e3a)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](data-sources--securemesh_site_v2--reference--group-007.md#canonical-26188b3f7819ae2648313aefc4e535174c439be8c073cb2323ffe688ddc32e09)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-007.md#canonical-9f806526dc3dd1159f56701dbcbcba74948ac1d6c895818ea0b811cfb6e7dff1)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](data-sources--securemesh_site_v2--reference--group-007.md#canonical-7cd85f4ecba5e70219d237ad6728f98660d5ebc3137201592fe2feef576ec0f6)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-60508a7c745a5740878d75273de529291316607c55130b66f6634f307f939e3a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0cecf3639198ac6872a7c3a6c398a9cc6c745b7d68b2a33877f5c80f7eea2330"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / 64407abf9e85 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c4c1f243d85ccb664e50a678e8000889fea55c0fd9a37a19ad6f87eea1d5ba60)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-cc22c28026298aa1925cdec304a4590a2facb540e753ba86fae63301d178a1b3"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from end.

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

<a id="canonical-a209c68aabaa47d1825942330446198ed112d171fa4502af4e3c84e54dd9a380"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / 64407abf9e85 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d16f207b8148e07f8a6b760d3a82c55aeb0adc801c41daefabbb924300e8d93a"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / 64407abf9e85 / 4

- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c4c1f243d85ccb664e50a678e8000889fea55c0fd9a37a19ad6f87eea1d5ba60)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-26188b3f7819ae2648313aefc4e535174c439be8c073cb2323ffe688ddc32e09"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3192a25582ad9288aa543a0c6493a9bc6c648517e87d93d221635978f1a8a6f7"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 8168a82b3227 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c4c1f243d85ccb664e50a678e8000889fea55c0fd9a37a19ad6f87eea1d5ba60)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-9d3d8696613d1fa476265a95aeeea39ee9658dd970d469036ffeeca00d0e3c3e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from start.

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

<a id="canonical-340bd456c0e450aaa6414c602b5ffbfcc641e45dd5e80a043f8cee5b0d52eeb6"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 8168a82b3227 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5851cf00ed2bc4ee85084153eb6dc62d62be285a178125e600d2ce93a3941458"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 8168a82b3227 / 4

- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c4c1f243d85ccb664e50a678e8000889fea55c0fd9a37a19ad6f87eea1d5ba60)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9f806526dc3dd1159f56701dbcbcba74948ac1d6c895818ea0b811cfb6e7dff1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d08ef48436cbd2dd174703994eb7b179c360580a4e7d02ca88af313e06a9354d"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / f6b6691f9dad / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c4c1f243d85ccb664e50a678e8000889fea55c0fd9a37a19ad6f87eea1d5ba60)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-66e26e35f740e187e0072f3a58e95184446bca897c0b291bb711cd42170b6945"></a>

Type: `"list"`. Computed.

List of networks from which DHCP Server can allocate IPv4 Addresses.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2e2cbaf7e22072a65ae8e328b243062ccabe79747e49b5811aca79e29d3af94d"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / f6b6691f9dad / 3

<a id="canonical-3d331bb488be498b3f060f52f86ed6d37f4f2412dc5b05589ca01d03dfeeb6f5"></a>

<a id="canonical-479f7ecc465d8f4851a656c3c5261b68833464d13e448ba481ed2e793f02e35c"></a>

## dgw_address property — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / f6b6691f9dad / 4

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

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

<a id="canonical-d39ae8f0ec1b5ea6458773006a05670e3f206da2db82803cd1209ed55097deb9"></a>

<a id="canonical-c647e4e43b7c1f59573b5f01167245d89f48df17030106c91dcb86341c35b23c"></a>

## dns_address property — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / f6b6691f9dad / 5

Type: `"string"`. Computed.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

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

- [first_address](data-sources--securemesh_site_v2--reference--group-007.md#canonical-6beb4bffa4c4732dfa735e2fed5b7c4e342831f331a78a4a5cbe5622a62a557b): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-007.md#canonical-e6cb144b08cf51f474dca72ed2b03aa729e17563705120063d991ca9de14298f): complete subsection reference.

<a id="canonical-e4b8d5a058474b084abcc6a7f888ba6ee9281bc9266f93ac2fbe31a2153b40d3"></a>

<a id="canonical-3b9174f8f2cd5c2f5b329ee9f80a6d12eb6c4d7895c964a38b1f680a68ed523e"></a>

## network_prefix property — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / f6b6691f9dad / 6

Type: `"string"`. Computed.

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

Upstream description:

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

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

<a id="canonical-13331c0cbf46536ca9b848fdc0b894329671e3f964b25776b0ce7cc7a268642d"></a>

<a id="canonical-4211ba014d9673b9d25b76b55f996ce1cb28cfcc2245fa94b2225b372779bb9b"></a>

## pool_settings property — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / f6b6691f9dad / 7

Type: `"string"`. Computed.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](data-sources--securemesh_site_v2--reference--group-007.md#canonical-68b341146872c4d8c0e1a76c0c56755745f4199e0b5cbc70cb10f67dec8f9c9a): complete subsection reference.

- [same_as_dgw](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c85f6b8a18c628a5c1b47aaa29d26761fd9237fe47bc47df8d0f9b4ac1c74524): complete subsection reference.

<a id="canonical-51e803e4ea5fcc6e9e458aba161c5fbbd718b00d91ed44fd84167b4dc1238bd7"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / f6b6691f9dad / 8

- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](data-sources--securemesh_site_v2--reference--group-007.md#canonical-6beb4bffa4c4732dfa735e2fed5b7c4e342831f331a78a4a5cbe5622a62a557b)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](data-sources--securemesh_site_v2--reference--group-007.md#canonical-e6cb144b08cf51f474dca72ed2b03aa729e17563705120063d991ca9de14298f)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](data-sources--securemesh_site_v2--reference--group-007.md#canonical-68b341146872c4d8c0e1a76c0c56755745f4199e0b5cbc70cb10f67dec8f9c9a)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c85f6b8a18c628a5c1b47aaa29d26761fd9237fe47bc47df8d0f9b4ac1c74524)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c4c1f243d85ccb664e50a678e8000889fea55c0fd9a37a19ad6f87eea1d5ba60)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-6beb4bffa4c4732dfa735e2fed5b7c4e342831f331a78a4a5cbe5622a62a557b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6daac072cd37264a8a308df7673f1a317b9212da2beb673ce8c187436990719"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_add / 86bc1ef59187 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c4c1f243d85ccb664e50a678e8000889fea55c0fd9a37a19ad6f87eea1d5ba60)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-007.md#canonical-9f806526dc3dd1159f56701dbcbcba74948ac1d6c895818ea0b811cfb6e7dff1)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-7e26b41f75b9be538e80d8b7f2b27b7abce33e168c4d9fb58a80b5376562ac4f"></a>

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

<a id="canonical-7e89e1a134fa989f3cf225cdaaed359006d4ce2b160d37bc5acae452a95ebb22"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_add / 86bc1ef59187 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b7bf568e5144563dfa14dba8ffc0b61ee1f4ae5f62a4b20423f819fddd1fed85"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_add / 86bc1ef59187 / 4

- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-007.md#canonical-9f806526dc3dd1159f56701dbcbcba74948ac1d6c895818ea0b811cfb6e7dff1)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-e6cb144b08cf51f474dca72ed2b03aa729e17563705120063d991ca9de14298f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c44dc2cff28cf0e2472d7243bc0ed4572d2c2e9d1d71831628687334f172080"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addr / e23cadc3dd9e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c4c1f243d85ccb664e50a678e8000889fea55c0fd9a37a19ad6f87eea1d5ba60)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-007.md#canonical-9f806526dc3dd1159f56701dbcbcba74948ac1d6c895818ea0b811cfb6e7dff1)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-80df51d0c24045123706a0faea2b945376329759e0b23236dc1ce21b7ea4d405"></a>

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

<a id="canonical-a9068278bfc4d186dcc3f88b8e44b3b2e45c60e2d11b145dcdce0b8a1203242d"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addr / e23cadc3dd9e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-80f3fda8aa8fedc97c9b46e58055bb89891697da690e1734ae80f7d58d10b3d7"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addr / e23cadc3dd9e / 4

- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-007.md#canonical-9f806526dc3dd1159f56701dbcbcba74948ac1d6c895818ea0b811cfb6e7dff1)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-68b341146872c4d8c0e1a76c0c56755745f4199e0b5cbc70cb10f67dec8f9c9a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b45ae23cfc930384b391f0c1e586495e29afd352e13835aa28bece38c68f545e"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 32e7bed1d5e6 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c4c1f243d85ccb664e50a678e8000889fea55c0fd9a37a19ad6f87eea1d5ba60)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-007.md#canonical-9f806526dc3dd1159f56701dbcbcba74948ac1d6c895818ea0b811cfb6e7dff1)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-d6f5336d64e5eb0d73473eae35e81e41d998041c6994e5918b4862a88038d7b6"></a>

Type: `"list"`. Computed.

List of non overlapping IP address ranges.

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
    "minItems": 1,
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

<a id="canonical-512dedb10e9381dd53feb263f957c382a23ff7a4dde59ccefdbef87d704fec6d"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 32e7bed1d5e6 / 3

<a id="canonical-025c59f76179f266da2d30074a99273d6346ade34f08f9d530fcd3f5253214aa"></a>

<a id="canonical-bdaaa70790a05b52e2c92fef467d14759cd9755e4f65ba7348510b64a1616ae0"></a>

## end_ip property — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 32e7bed1d5e6 / 4

Type: `"string"`. Computed.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Upstream description:

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

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

<a id="canonical-62a5e98529fad461cb81aa43439fc1a93422cfe8206e5fcdaed08e6d9f2c4826"></a>

<a id="canonical-8cff96a261e30f8790c5d45f56ad33a93206f3528c01894cf0b27ceb8fc1c65c"></a>

## exclude property — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 32e7bed1d5e6 / 5

Type: `"bool"`. Computed.

Exclude this address range from DHCP allocation.

<a id="canonical-0bc805e56beb9ca97b18c726f1b46954ff1e0875fd2154d160d40e83b380f656"></a>

<a id="canonical-52ea2a208da136477bc7846ac20fc2b8cf2bb6863f5706d28d7d765074754f2d"></a>

## start_ip property — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 32e7bed1d5e6 / 6

Type: `"string"`. Computed.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Upstream description:

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

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

<a id="canonical-1b00769b4a5502bc4637a6db0845c5e101882be4a059d4c926227f04615bec59"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 32e7bed1d5e6 / 7

- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-007.md#canonical-9f806526dc3dd1159f56701dbcbcba74948ac1d6c895818ea0b811cfb6e7dff1)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c85f6b8a18c628a5c1b47aaa29d26761fd9237fe47bc47df8d0f9b4ac1c74524"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ef2eae98cc77156de6faed113a2ea0cb2ea663d2ab529b9d14d2c3ac5b3d0f1"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_d / 14aa57ded8d3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c4c1f243d85ccb664e50a678e8000889fea55c0fd9a37a19ad6f87eea1d5ba60)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-007.md#canonical-9f806526dc3dd1159f56701dbcbcba74948ac1d6c895818ea0b811cfb6e7dff1)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-89213016a3042485ee8e156bdff0b5ab54c4e5a4ed53a47f037209eba2361824"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for same as dgw.

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

<a id="canonical-7af1c8dd252772d74aaa6de3a2f4dfa927a9a27eeadd883afa5476329566a2c8"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_d / 14aa57ded8d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-46c0d283b3f1f8a86c5c19ffc6b5e0f01763754cd8fb61a488492ee318b984ac"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_d / 14aa57ded8d3 / 4

- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-007.md#canonical-9f806526dc3dd1159f56701dbcbcba74948ac1d6c895818ea0b811cfb6e7dff1)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-7cd85f4ecba5e70219d237ad6728f98660d5ebc3137201592fe2feef576ec0f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a32b76638351415bed973fb9dcb46e14e8588ecdc40b92f2ee3125ce9606d32d"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — eks_k8s.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 8ccb336a5519 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c4c1f243d85ccb664e50a678e8000889fea55c0fd9a37a19ad6f87eea1d5ba60)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-c71cc5f5d0c78e554ba0b535db092566fbd7f29a069fc003baa7724a1824c331"></a>

Type: `"single"`. Computed.

Interface IPv4 Assignments. Specify static IPv4 addresses per node.

Upstream description:

Specify static IPv4 addresses per node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-bbb613aa781e80d488f5c1936bbe87cd6279d4776fb71b36f74b36cd37ccaff7"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 8ccb336a5519 / 3

<a id="canonical-8b3bbefcb843bc8e9c807b37e0f565136524ef17a40cbbe89ad37f80ee5b1696"></a>

<a id="canonical-fe4f2eef68f586ab05130e95df9b87ae9a00f8609de4f61641a293e98700291b"></a>

## interface_ip_map property — eks_k8s.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 8ccb336a5519 / 4

Type: `["map", "string"]`. Computed.

Specify static IPv4 addresses per site:node.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

<a id="canonical-62e8c0f2b844b5b3232feb3e5e083a827cb82f7fd34d422c830f78cb5f24e81e"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 8ccb336a5519 / 5

- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c4c1f243d85ccb664e50a678e8000889fea55c0fd9a37a19ad6f87eea1d5ba60)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-5d240ca2a5fd383b44dcc9675e16734cfa7e73f69641f9d2b44b6beec8c3d4de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2aa40dede12f0d29383f37d9312f14b7522f6a9c642f18e5becf4d177d1d9057"></a>

## eks_k8s.not_managed.node_list.interface_list.ethernet_interface — eks_k8s.not_managed.node_list.interface_list.ethernet_interface / 687b08963d3b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- eks_k8s.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-2852acf2ac3d2c0f0b69609fd0c01b56bd876dc3a3419d591c0ec22cc354e280"></a>

Type: `"single"`. Computed.

Configuration parameter for ethernet interface.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b9ff2c7c24001077a88a3f6fa94efd0603989325b4c310d5554de1cfd8f796ff"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.ethernet_interface / 687b08963d3b / 3

<a id="canonical-a4b39fed878e9a2157e0e785dbadbe8919187c7abbe2b83af47e2e65534b3943"></a>

<a id="canonical-1b1248356b903648a0d65b22cbac1aba7e7af258e556e9e739f5a53a51671ced"></a>

## device property — eks_k8s.not_managed.node_list.interface_list.ethernet_interface / 687b08963d3b / 4

Type: `"string"`. Computed.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Upstream description:

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

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
    "create": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-ec31a9c0b37e967747cfb49893f62f9ffe7d0e249a14272719f1e2cb6992b726"></a>

<a id="canonical-5a7df5e2053f3d2bbce2fb823223fa5e33a2fad34826b9196a8348b07af73207"></a>

## mac property — eks_k8s.not_managed.node_list.interface_list.ethernet_interface / 687b08963d3b / 5

Type: `"string"`. Computed.

MAC Address. Configuration parameter for mac

Upstream description:

Configuration parameter for mac

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "mac-address",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  }
}
```

<a id="canonical-ea548c27c6fef0228a27fc6803c2dbc1dac2166a826ac7f5b8afcae61299c29b"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.ethernet_interface / 687b08963d3b / 6

- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-8958acd121a1de591d9c295a38c154a6b85f9e6aa7cdaaf4b4bed2d5300a15d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f153242b58ad4cb86b084b869acbf75d99c496226d63964033ed49226040d5c"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config / 8a4214f5daf5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-fc3b01c9335ac8e970b47955daf5d6b7a0c27e1958386d0a25f3b390d98bd9d3"></a>

Type: `"single"`. Computed.

IPV6AutoConfigType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

<a id="canonical-4688d3c618cc13e8a251642979d2ebb79998a20c751eafc86f12bab7bf43b114"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config / 8a4214f5daf5 / 3

- [host](data-sources--securemesh_site_v2--reference--group-007.md#canonical-6bedc1e5e06b351fa64380e1c493065e719eeed2b6b0548e1c7c239fb89c60f6): complete subsection reference.

- [router](data-sources--securemesh_site_v2--reference--group-007.md#canonical-adca2f313eae42c06af5bb050db8941faff9b457179440a1d84daf306fd8e1f9): complete subsection reference.

<a id="canonical-f3eead71746569b110a16242cef8f76a050705f66f5772b3465b7ee984643a10"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config / 8a4214f5daf5 / 4

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.host](data-sources--securemesh_site_v2--reference--group-007.md#canonical-6bedc1e5e06b351fa64380e1c493065e719eeed2b6b0548e1c7c239fb89c60f6)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-007.md#canonical-adca2f313eae42c06af5bb050db8941faff9b457179440a1d84daf306fd8e1f9)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-6bedc1e5e06b351fa64380e1c493065e719eeed2b6b0548e1c7c239fb89c60f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00d4bb9a9056877d9313842c353ff45f3e3e916728cb19806e8d866a0ad8b456"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.host — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.host / c963595e3837 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8958acd121a1de591d9c295a38c154a6b85f9e6aa7cdaaf4b4bed2d5300a15d1)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-bbb5a1e0332c6fa1fc53ccf9e72e8582479b0855f5f6bf1fed0818e90afdc8bb"></a>

Type: `["object", {}]`. Computed.

Hostname or IP address of the target server.

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

<a id="canonical-09425733a46b07dab1f6410c5c1ef5645c842362059b192ea557edb731818240"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.host / c963595e3837 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b9e7d035cdb1c79c1ded205a0491d503092a3eda1f80c9de1303bd17a5c8cbea"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.host / c963595e3837 / 4

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8958acd121a1de591d9c295a38c154a6b85f9e6aa7cdaaf4b4bed2d5300a15d1)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-adca2f313eae42c06af5bb050db8941faff9b457179440a1d84daf306fd8e1f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7a6bae745866cdc2343f1ae85b654457e88cdf0d77c76a51bf0af2295257e1f"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router / 38c9312ba91c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8958acd121a1de591d9c295a38c154a6b85f9e6aa7cdaaf4b4bed2d5300a15d1)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-2b93ffc6cf64e1a5d1edfb258731cf83c9323055a37bb2d442ca2735ba2f616f"></a>

Type: `"single"`. Computed.

IPV6AutoConfigRouterType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-address_choice": "[\"network_prefix\",\"stateful\"]"
}
```

<a id="canonical-e27f903b857662d37a4006bf5571ae010d473d8353faa60e59f2cc663e4aac41"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router / 38c9312ba91c / 3

- [dns_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-f5e3503b0eed7ee2fb9894c55b7363fe81209e445310114e4ba3b7e2cda0ef58): complete subsection reference.

<a id="canonical-af2a8fd3148d499ada6b413a8c4aebddb2224652afafa1eea4a6e373d9b3c516"></a>

<a id="canonical-2137f894595d5ec9eaddada2ef2534bb98308f80a8ec5576d7fac41746f7ee0f"></a>

## network_prefix property — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router / 38c9312ba91c / 4

Type: `"string"`. Computed.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": ".*::/64$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  }
}
```

- [stateful](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8449990c88b1b7e5be4a0c5b0bd8702b8603e3c9c21f498bd9527ce98e021502): complete subsection reference.

<a id="canonical-03badcc9aef11dce53f1ced60df339674024863a33dc2222e63bd8237e8afbd2"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router / 38c9312ba91c / 5

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-f5e3503b0eed7ee2fb9894c55b7363fe81209e445310114e4ba3b7e2cda0ef58)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8449990c88b1b7e5be4a0c5b0bd8702b8603e3c9c21f498bd9527ce98e021502)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8958acd121a1de591d9c295a38c154a6b85f9e6aa7cdaaf4b4bed2d5300a15d1)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-f5e3503b0eed7ee2fb9894c55b7363fe81209e445310114e4ba3b7e2cda0ef58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80eb978e9959a853ac5942fbb5d171347c090be74767a79335aa8323a13fcfc3"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 548f753a076e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8958acd121a1de591d9c295a38c154a6b85f9e6aa7cdaaf4b4bed2d5300a15d1)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-007.md#canonical-adca2f313eae42c06af5bb050db8941faff9b457179440a1d84daf306fd8e1f9)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-4796fd8166d9fa192a79c8bd95502ea09da3e472aea5f64a9662046a6c375f83"></a>

Type: `"single"`. Computed.

IPV6DnsConfig.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dns_choice": "[\"configured_list\",\"local_dns\"]"
}
```

<a id="canonical-f17bf7e39164397d5fbcf52ac51231cdc5c73e423be7edad20494a5b74c6e527"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 548f753a076e / 3

- [configured_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-92228e957f69ec341377a2228adb5f068e552d0d747f1c3e68459373c91098f0): complete subsection reference.

- [local_dns](data-sources--securemesh_site_v2--reference--group-007.md#canonical-64c61bdb54b6d8d85083a86498563c34a805363c103a29151f78ecb6162f1d69): complete subsection reference.

<a id="canonical-d84aef07d75fd34bf6de8f1bd62601a3fe6d175cb024caeb959d953e96b2d477"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 548f753a076e / 4

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-92228e957f69ec341377a2228adb5f068e552d0d747f1c3e68459373c91098f0)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-007.md#canonical-64c61bdb54b6d8d85083a86498563c34a805363c103a29151f78ecb6162f1d69)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-007.md#canonical-adca2f313eae42c06af5bb050db8941faff9b457179440a1d84daf306fd8e1f9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-92228e957f69ec341377a2228adb5f068e552d0d747f1c3e68459373c91098f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0039ff6abece530a574f0ccea748803f926d28f89b2fcee17ef0ff0606a3f83b"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 24d4cabb9ebe / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8958acd121a1de591d9c295a38c154a6b85f9e6aa7cdaaf4b4bed2d5300a15d1)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-007.md#canonical-adca2f313eae42c06af5bb050db8941faff9b457179440a1d84daf306fd8e1f9)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-f5e3503b0eed7ee2fb9894c55b7363fe81209e445310114e4ba3b7e2cda0ef58)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-67299fce1b70afa9fb62ae246b1f2a3c06b4126d3bb880e17f48cdeeaacd640f"></a>

Type: `"single"`. Computed.

IPV6DnsList.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3e281fe8b1e7a8b8d44fa010fd041a9d66d8d2c81317d4ab089120ac57ce2c52"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 24d4cabb9ebe / 3

<a id="canonical-9824c7eeadc610cac74fdd0ef9f2ba28eb86773fdc767d08125389f6a19cdabd"></a>

<a id="canonical-b620d2616460e40c9ea046f9d0f5e24cdfef16cb2e25fc9f526c5ab7ad3431a9"></a>

## dns_list property — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 24d4cabb9ebe / 4

Type: `["list", "string"]`. Computed.

List of IPv6 Addresses acting as DNS servers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2c68b0d56e4afdb03ebac012d2addb507f0f4a0a569c7d0dbcd2709d4e84ae22"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 24d4cabb9ebe / 5

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-f5e3503b0eed7ee2fb9894c55b7363fe81209e445310114e4ba3b7e2cda0ef58)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-64c61bdb54b6d8d85083a86498563c34a805363c103a29151f78ecb6162f1d69"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5dfbd4f68c4686a01aba26a7c75f38d68f57cf8b9f76f6f29f0adeadc8d84f1f"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 06298fbb8b3e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8958acd121a1de591d9c295a38c154a6b85f9e6aa7cdaaf4b4bed2d5300a15d1)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-007.md#canonical-adca2f313eae42c06af5bb050db8941faff9b457179440a1d84daf306fd8e1f9)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-f5e3503b0eed7ee2fb9894c55b7363fe81209e445310114e4ba3b7e2cda0ef58)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-aa313f468c1b15323b93651ad4fd81ce99ec6215a48626c53e047534f7d43735"></a>

Type: `"single"`. Computed.

IPV6LocalDnsAddress.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-local_dns_choice": "[\"configured_address\",\"first_address\",\"last_address\"]"
}
```

<a id="canonical-66f151eb6a5d3d6bd39d4ef2ae14b85362a7b267485e626d7ae704442a975ffb"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 06298fbb8b3e / 3

<a id="canonical-f6bc1ec3d17aa39bc8e691ba6ad5e1f38d9bd2d7d07cf8c5c9bab27aafd50d23"></a>

<a id="canonical-f3c1defc4d4646bd2981bb5f4e17cfab4e63ddd0baffb369cd2917c7e5e08c82"></a>

## configured_address property — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 06298fbb8b3e / 4

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

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

- [first_address](data-sources--securemesh_site_v2--reference--group-007.md#canonical-531b0222f6eb23624ae9ae08dad0f72b7778406bb74b57183357a7df024240cb): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-007.md#canonical-cdc617c75d08ac284d5a17617c0f66f05a3991b994140791d7cf921e8dee6426): complete subsection reference.

<a id="canonical-a3f65132a441c79000dce4dbb27fb1658a46e23d4432493d7689598c7bf9111e"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 06298fbb8b3e / 5

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--securemesh_site_v2--reference--group-007.md#canonical-531b0222f6eb23624ae9ae08dad0f72b7778406bb74b57183357a7df024240cb)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--securemesh_site_v2--reference--group-007.md#canonical-cdc617c75d08ac284d5a17617c0f66f05a3991b994140791d7cf921e8dee6426)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-f5e3503b0eed7ee2fb9894c55b7363fe81209e445310114e4ba3b7e2cda0ef58)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-531b0222f6eb23624ae9ae08dad0f72b7778406bb74b57183357a7df024240cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-437b349e11f1729d09116ef32ab9a4d73e32a6467bdf130d348a978ef9548563"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 69b911d2a0cf / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8958acd121a1de591d9c295a38c154a6b85f9e6aa7cdaaf4b4bed2d5300a15d1)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-007.md#canonical-adca2f313eae42c06af5bb050db8941faff9b457179440a1d84daf306fd8e1f9)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-f5e3503b0eed7ee2fb9894c55b7363fe81209e445310114e4ba3b7e2cda0ef58)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-007.md#canonical-64c61bdb54b6d8d85083a86498563c34a805363c103a29151f78ecb6162f1d69)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-0949e08cd630cc36677fb8c075df5addf8a1c9d5e5df597b9130cb85ea76231a"></a>

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

<a id="canonical-f813fb3889b46f85acd0f75f0208a64f9e3457b99affce031141afc4928100f4"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 69b911d2a0cf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-235cd2ca2a783bcce94093343cc6129c280063bcd62664c06762f70a6bfcfe83"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 69b911d2a0cf / 4

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-007.md#canonical-64c61bdb54b6d8d85083a86498563c34a805363c103a29151f78ecb6162f1d69)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-cdc617c75d08ac284d5a17617c0f66f05a3991b994140791d7cf921e8dee6426"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7272de14be25d5301b36622e6f5e9886bc7542724e760d9bec36ba036ff0253f"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 032b37db58b5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8958acd121a1de591d9c295a38c154a6b85f9e6aa7cdaaf4b4bed2d5300a15d1)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-007.md#canonical-adca2f313eae42c06af5bb050db8941faff9b457179440a1d84daf306fd8e1f9)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-f5e3503b0eed7ee2fb9894c55b7363fe81209e445310114e4ba3b7e2cda0ef58)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-007.md#canonical-64c61bdb54b6d8d85083a86498563c34a805363c103a29151f78ecb6162f1d69)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-b90fba3f4d9fd211f19bdbad73bb370f43a29344a76d9049ead95ba376cef938"></a>

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

<a id="canonical-1801167408ab81002d4b4e4e4e3ecce28a56d14935fbe9fa2c9e09b3513df4a8"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 032b37db58b5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a5675af50abe98a3b3c48928886989d02d979ad579f2be9d549f01b19e7efa62"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 032b37db58b5 / 4

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-007.md#canonical-64c61bdb54b6d8d85083a86498563c34a805363c103a29151f78ecb6162f1d69)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-8449990c88b1b7e5be4a0c5b0bd8702b8603e3c9c21f498bd9527ce98e021502"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d42767d840f3cefcea8a761a3896d383da8707f3e5dec4ccb92dcbb2be10ffd8"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 24f3aa82a49f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8958acd121a1de591d9c295a38c154a6b85f9e6aa7cdaaf4b4bed2d5300a15d1)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-007.md#canonical-adca2f313eae42c06af5bb050db8941faff9b457179440a1d84daf306fd8e1f9)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-c9f8c6a3c71f111092f146a6dcc2b4b03842e9f97bd14b2a1534705a6e11b2fe"></a>

Type: `"single"`. Computed.

DHCPIPV6 Stateful Server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

<a id="canonical-59cec13bb25656ffaac9938d70e138ad3e53c75e332498fff79134066fa3a287"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 24f3aa82a49f / 3

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-007.md#canonical-9128da173ae80400f704c31160ca96f2af66a557c886971b55f259ef182fbb9b): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-007.md#canonical-21be2e10f609ef5bd67c6aedaaedc824ce324ebe1ea7b756be22d4d4cc9170ef): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8479a01520b94cf33c578df6bc86681323d38faf19d3719ababe4dc2a382c8a6): complete subsection reference.

<a id="canonical-ece9f08299879ace42e9b5a36d04deb360b94ec615d74d4d272b0c9241fbd1f3"></a>

<a id="canonical-4ed24abcf72982bfa19b953f98aa5684f9ee759e3250184427007dab59272711"></a>

## fixed_ip_map property — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 24f3aa82a49f / 4

Type: `["map", "string"]`. Computed.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Upstream description:

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

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
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-007.md#canonical-19d7f02c52608217f18aa5e31c52fd217015e5e0ccc96ca3e71484c25d14bb42): complete subsection reference.

<a id="canonical-b1735ded602a54dc6dc1a6b894a37595a6150a5002b2ae9ec130ce7ecb2d08d2"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 24f3aa82a49f / 5

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--securemesh_site_v2--reference--group-007.md#canonical-9128da173ae80400f704c31160ca96f2af66a557c886971b55f259ef182fbb9b)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--securemesh_site_v2--reference--group-007.md#canonical-21be2e10f609ef5bd67c6aedaaedc824ce324ebe1ea7b756be22d4d4cc9170ef)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8479a01520b94cf33c578df6bc86681323d38faf19d3719ababe4dc2a382c8a6)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--securemesh_site_v2--reference--group-007.md#canonical-19d7f02c52608217f18aa5e31c52fd217015e5e0ccc96ca3e71484c25d14bb42)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-007.md#canonical-adca2f313eae42c06af5bb050db8941faff9b457179440a1d84daf306fd8e1f9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9128da173ae80400f704c31160ca96f2af66a557c886971b55f259ef182fbb9b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f6d262493726fbfeb3ca17ea459be487a777d1a3b25b7fd58a32385ca7e8a76"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.au / f2b5232818b1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8958acd121a1de591d9c295a38c154a6b85f9e6aa7cdaaf4b4bed2d5300a15d1)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-007.md#canonical-adca2f313eae42c06af5bb050db8941faff9b457179440a1d84daf306fd8e1f9)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8449990c88b1b7e5be4a0c5b0bd8702b8603e3c9c21f498bd9527ce98e021502)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-55af8cec9882bfc7ed53a26c5ad4bc5c7b92d8e5b404ce0a6621074c178ebe69"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from end.

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

<a id="canonical-616273b7d7fd4d609596ae58f57133eac1ff0ad4c98e16bab11a1c4abd1b8074"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.au / f2b5232818b1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d10bbdfdfeb0a6277b174a97037dba92a8030bec67997b3c4318fca254f9532f"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.au / f2b5232818b1 / 4

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8449990c88b1b7e5be4a0c5b0bd8702b8603e3c9c21f498bd9527ce98e021502)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-21be2e10f609ef5bd67c6aedaaedc824ce324ebe1ea7b756be22d4d4cc9170ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-326de86d2bfc3c0f0737826f0cc6c5fc0f8ec7fce8d4d684c4382a511cf9cef0"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.au / 020efb7fe63f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8958acd121a1de591d9c295a38c154a6b85f9e6aa7cdaaf4b4bed2d5300a15d1)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-007.md#canonical-adca2f313eae42c06af5bb050db8941faff9b457179440a1d84daf306fd8e1f9)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8449990c88b1b7e5be4a0c5b0bd8702b8603e3c9c21f498bd9527ce98e021502)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-ce19f4fedd592c7cb06eaafd34fa05f0aac7f789341e9bfb85c64fbc45644c92"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from start.

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

<a id="canonical-a5b1a6cc4be559f810449e51310e8d40ec8534612e77c638c6e55cd7a2e94532"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.au / 020efb7fe63f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c22a4c53fa9fea6c769411f560d45c4df56d1aa105bc9e1815d967e8becd6690"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.au / 020efb7fe63f / 4

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8449990c88b1b7e5be4a0c5b0bd8702b8603e3c9c21f498bd9527ce98e021502)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-8479a01520b94cf33c578df6bc86681323d38faf19d3719ababe4dc2a382c8a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c70bc8416194195d6549517946ea4a0563d7d88644fa9df7290eba7cf19a5f7b"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / 3862938ce237 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8958acd121a1de591d9c295a38c154a6b85f9e6aa7cdaaf4b4bed2d5300a15d1)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-007.md#canonical-adca2f313eae42c06af5bb050db8941faff9b457179440a1d84daf306fd8e1f9)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8449990c88b1b7e5be4a0c5b0bd8702b8603e3c9c21f498bd9527ce98e021502)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-b72595b3668da75ce1ace0c15c2cf16f69e0e69b8b8bebff2a9d9e2878d84303"></a>

Type: `"list"`. Computed.

List of networks from which DHCP server can allocate IP addresses.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8f6d4e2c9d063a1ab070564d093f504be9eb9cbc95b540efae4b08a0538cc25d"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / 3862938ce237 / 3

<a id="canonical-7b04f7b7fafdd0707f71bb6840daf1b4e287ddabb5e78e2115664f35e17b9d6f"></a>

<a id="canonical-c13d885997bbb7e84c7fdb71b8ca77d919f095645f669252c50f81b0b0ab5764"></a>

## network_prefix property — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / 3862938ce237 / 4

Type: `"string"`. Computed.

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

Upstream description:

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

<a id="canonical-5eddce4d5574b51b3c6692896360f5633b8122b4a7c75ebd06bfb7910e4bbe2b"></a>

<a id="canonical-a99c5cd7ef4ea18f128f845ab0c24509f5a0194608a49268746d73bc59f8f04c"></a>

## pool_settings property — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / 3862938ce237 / 5

Type: `"string"`. Computed.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](data-sources--securemesh_site_v2--reference--group-007.md#canonical-b9ce291f5b30034fae6ee1d046e6df5bef121ef2d7fe6ecb45683496de7b75fc): complete subsection reference.

<a id="canonical-ab724a6e3989e00c428ccdaf64cdcc5cfa495a3ac7e904c67d8165a91420c392"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / 3862938ce237 / 6

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--securemesh_site_v2--reference--group-007.md#canonical-b9ce291f5b30034fae6ee1d046e6df5bef121ef2d7fe6ecb45683496de7b75fc)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8449990c88b1b7e5be4a0c5b0bd8702b8603e3c9c21f498bd9527ce98e021502)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-b9ce291f5b30034fae6ee1d046e6df5bef121ef2d7fe6ecb45683496de7b75fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d7737203654dbcb37262847a8d3e335beea3832524c57a5e5824082e4665009"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / 1b83356aef86 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8958acd121a1de591d9c295a38c154a6b85f9e6aa7cdaaf4b4bed2d5300a15d1)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-007.md#canonical-adca2f313eae42c06af5bb050db8941faff9b457179440a1d84daf306fd8e1f9)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8449990c88b1b7e5be4a0c5b0bd8702b8603e3c9c21f498bd9527ce98e021502)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8479a01520b94cf33c578df6bc86681323d38faf19d3719ababe4dc2a382c8a6)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-06a4448fb0c630dc61ad4666244256643625647d496ba15e746a6257c9d77c3a"></a>

Type: `"list"`. Computed.

List of non overlapping IP address ranges.

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
    "minItems": 1,
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

<a id="canonical-0659dc805cd8a077bca6018021415fb4d40a5971451c66d6f415a72b9ea8ca22"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / 1b83356aef86 / 3

<a id="canonical-1e7c98a7089c3bab42906f19d853e2908a9b560c3e4c1d015fc8d6cb56c05b8c"></a>

<a id="canonical-90fce4f2e276238327d335efbaff7b0f3d31005c176cafce0fb5eaaeb7f6df93"></a>

## end_ip property — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / 1b83356aef86 / 4

Type: `"string"`. Computed.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Upstream description:

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

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

<a id="canonical-5f27daf801a298fed09bbe468e92f52bf5098a377915446011b48120fe94e601"></a>

<a id="canonical-91e9c691f59bc82315fd632746687aed44b636a6a9deb13a0b187aa63d25b805"></a>

## start_ip property — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / 1b83356aef86 / 5

Type: `"string"`. Computed.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Upstream description:

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

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

<a id="canonical-b3a55b4b66329170aa08d75460d5f7abfd92dbf2cfa43e0a6f06ae856f905ef9"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / 1b83356aef86 / 6

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8479a01520b94cf33c578df6bc86681323d38faf19d3719ababe4dc2a382c8a6)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-19d7f02c52608217f18aa5e31c52fd217015e5e0ccc96ca3e71484c25d14bb42"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2dce9ea7c21a0bd9612cb879850dc6fe38ba309ec5ec5fc93b31086383056e0"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.in / f7b132faaa06 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8958acd121a1de591d9c295a38c154a6b85f9e6aa7cdaaf4b4bed2d5300a15d1)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-007.md#canonical-adca2f313eae42c06af5bb050db8941faff9b457179440a1d84daf306fd8e1f9)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8449990c88b1b7e5be4a0c5b0bd8702b8603e3c9c21f498bd9527ce98e021502)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-d7bb36b686df6efd5a95cde5b4ab4993704c3ed3fb6daa9916f290e6e8eddf9b"></a>

Type: `"single"`. Computed.

Map of Interface IPv6 assignments per node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-15c7673e427d2f5af61f3babbe11c62c92e692b725d311f5f4409c53fa5deec2"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.in / f7b132faaa06 / 3

<a id="canonical-999f47e8d64bdd20e03cb9716c1edf6a4467beaa0330c007a80b2c95349d688f"></a>

<a id="canonical-d38856c0494381840617f5b215220c8111c32322d33d713c188aca1a0d91a5e2"></a>

## interface_ip_map property — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.in / f7b132faaa06 / 4

Type: `["map", "string"]`. Computed.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Upstream description:

Map of Site:Node to IPv6 address.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

<a id="canonical-c49fd2caf331d5bb4411bd2035c9a1f3f8427193d231bbe98271a0b6b9a461d3"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.in / f7b132faaa06 / 5

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-007.md#canonical-8449990c88b1b7e5be4a0c5b0bd8702b8603e3c9c21f498bd9527ce98e021502)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-89f31d3468d8e20d7b5b0430ae467c78c9725f45cd28956f981826e50d35d2a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f75c39f561ce7e1a1a0e9d25583b0a8df4934e8cee0a8db8a15fe7bb81e50734"></a>

## eks_k8s.not_managed.node_list.interface_list.monitor — eks_k8s.not_managed.node_list.interface_list.monitor / 6ef2a8ae09ea / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- eks_k8s.not_managed.node_list.interface_list.monitor

<a id="canonical-41e3c891d83f9827b433ead8af09cd25198aecaf308d549184fa9f46eda82763"></a>

Type: `["object", {}]`. Computed.

Link Quality Monitoring configuration for a network interface.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-a8f6006487d756ecd50409ac6d207d2d3fa0a2ae36138d1ffa0c478f0ae39c4c"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.monitor / 6ef2a8ae09ea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5ddfd48dc4160a1697fef2f4f94555a08694af1c96eb4cac0976820db402f756"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.monitor / 6ef2a8ae09ea / 4

- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-156462b8756a9a67940b9725de36116b6fee6831d4f58a84c88444b248c32aaa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36126e5aeda58059a6339a849f45270d50dfd4e3f3e8517c1b2bcd9db28ec05c"></a>

## eks_k8s.not_managed.node_list.interface_list.monitor_disabled — eks_k8s.not_managed.node_list.interface_list.monitor_disabled / 5e903b31ef82 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- eks_k8s.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-ae14551ae0a910f8c12001bd0d91ef7bba1abdf8c0a801d37cd1335b57928193"></a>

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

<a id="canonical-158df34e8245981d759238feaf04d6565f98deb1645fe5e764037cb95487aa10"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.monitor_disabled / 5e903b31ef82 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-86a075ec02ecf669f3af9499baabefd824bfa8aa37ded1718179387399fd3e8c"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.monitor_disabled / 5e903b31ef82 / 4

- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-92d1d61a8e50c720846a9a9dc9ff1f807d88b90899b7d46897205fe2979c50f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30b10f54e210599dda56575499e429f63f9f631b19c5808214540c9066631450"></a>

## eks_k8s.not_managed.node_list.interface_list.network_option — eks_k8s.not_managed.node_list.interface_list.network_option / f28a3dbccaf4 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- eks_k8s.not_managed.node_list.interface_list.network_option

<a id="canonical-05cf87c46425249529788bcc516c7d19064b1bd4acd8311c50ffca20b4d5b7cf"></a>

Type: `"single"`. Computed.

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional.

Upstream description:

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional. Global VRFs are configured via Networking &gt; Segments. A site can have multiple Network
Segments (global VRFs).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

<a id="canonical-788cf29d60da20e1f7f73d09bb19aad4fc239244901c54697d8d605b8596ab11"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.network_option / f28a3dbccaf4 / 3

- [site_local_inside_network](data-sources--securemesh_site_v2--reference--group-007.md#canonical-cb82896d7b34b57e58d21a93179e4329e4365dc90152c91f489cae757622b9d0): complete subsection reference.

- [site_local_network](data-sources--securemesh_site_v2--reference--group-007.md#canonical-074aae476aa4e357c2b2f0e6444382444339e0bab2ae1726b04683003ff0a89e): complete subsection reference.

<a id="canonical-2b6fa046059953ff696bd481e820b1836608d78eedf50084c19d69a2899cad98"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.network_option / f28a3dbccaf4 / 4

- [eks_k8s.not_managed.node_list.interface_list.network_option.site_local_inside_network](data-sources--securemesh_site_v2--reference--group-007.md#canonical-cb82896d7b34b57e58d21a93179e4329e4365dc90152c91f489cae757622b9d0)
- [eks_k8s.not_managed.node_list.interface_list.network_option.site_local_network](data-sources--securemesh_site_v2--reference--group-007.md#canonical-074aae476aa4e357c2b2f0e6444382444339e0bab2ae1726b04683003ff0a89e)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-cb82896d7b34b57e58d21a93179e4329e4365dc90152c91f489cae757622b9d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6e8097f6b846628b6ec5dad030cabdf833216ed4df2bea7b294748b17d8c282"></a>

## eks_k8s.not_managed.node_list.interface_list.network_option.site_local_inside_network — eks_k8s.not_managed.node_list.interface_list.network_option.site_local_inside_ne / 5d07b7d08f6c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-007.md#canonical-92d1d61a8e50c720846a9a9dc9ff1f807d88b90899b7d46897205fe2979c50f0)
- eks_k8s.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-1b5055dece08f968664e87af0e8265d82cf29f33a831b7d11dedb0e8d3deb89a"></a>

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

<a id="canonical-8f3498eb7fbe0c050680cd9c472a4051389db8dd0e743b890735b9d6fccf5a93"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.network_option.site_local_inside_ne / 5d07b7d08f6c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6312c12470e03060e4039fb265aea2538ca97d2173ff17b5d55dcb82d81d594e"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.network_option.site_local_inside_ne / 5d07b7d08f6c / 4

- [eks_k8s.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-007.md#canonical-92d1d61a8e50c720846a9a9dc9ff1f807d88b90899b7d46897205fe2979c50f0)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-074aae476aa4e357c2b2f0e6444382444339e0bab2ae1726b04683003ff0a89e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17b6ed6e6f21f031b22f025113b3f28a8e05fa97c831673ed9379319b9056f6a"></a>

## eks_k8s.not_managed.node_list.interface_list.network_option.site_local_network — eks_k8s.not_managed.node_list.interface_list.network_option.site_local_network / 3fd1e9a35213 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-007.md#canonical-92d1d61a8e50c720846a9a9dc9ff1f807d88b90899b7d46897205fe2979c50f0)
- eks_k8s.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-d2a8a1afd8eaf0f82bca696c0dda68ad56f3262ef8dd6fb53068f9c78b2130b9"></a>

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

<a id="canonical-f0a08996d619c03aa790c0d4ac663d91c544b71e7722cf31dbc4145d5e78ae36"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.network_option.site_local_network / 3fd1e9a35213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d4f0a1d3feeb05b3a6e914d6e56e74a418b506329046b8671c4552fc24c1d0bd"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.network_option.site_local_network / 3fd1e9a35213 / 4

- [eks_k8s.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-007.md#canonical-92d1d61a8e50c720846a9a9dc9ff1f807d88b90899b7d46897205fe2979c50f0)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-d21ed0b715126b4a844e74e1c1f79b73b2e970771cdcf8963ac37a0565a14660"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f2fd11bc21ae78997f8fdd6d0557ec8b21c52de36b35bfad932185644e09eb1"></a>

## eks_k8s.not_managed.node_list.interface_list.no_ipv4_address — eks_k8s.not_managed.node_list.interface_list.no_ipv4_address / b4f2935b2fa9 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- eks_k8s.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-f7b1f5c80e09a840d64b05b6cd7e15ed0a5f320caf4754ed26e2514b8b429373"></a>

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

<a id="canonical-79b27b772b56aeaefc8f6901ff97dbcb98aa53384aee86a68adf9e77b0fd3deb"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.no_ipv4_address / b4f2935b2fa9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6267201cb1654b758875f897a86b1c6bc9fa2e4551a3f646daa5150ce73e2c1a"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.no_ipv4_address / b4f2935b2fa9 / 4

- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-e1b16b35e297632d2254d99a4b53cc29cde056668d8537ad46283edb252d5bb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9d9290ebddaca47b7531312ff354942792f4bc5ba0d62334b2a944e5f6cd001"></a>

## eks_k8s.not_managed.node_list.interface_list.no_ipv6_address — eks_k8s.not_managed.node_list.interface_list.no_ipv6_address / 622b7af76156 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- eks_k8s.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-7387eec75360d70aa9ae01ef230e521fba8bc818cdcfa41bd1580fdaccd631ca"></a>

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

<a id="canonical-8dc44f33ccdb6f92b655ccdf6b71bfceafff73fce65c1403d67a26261d509296"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.no_ipv6_address / 622b7af76156 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cda154bbe996effe30873ad37d8f35a68f0d147110ce644808dc1d258b11dfc2"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.no_ipv6_address / 622b7af76156 / 4

- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-a48c3d0552e14088be10a64e10836db4ba356566a3b81fbb55238187eb82d128"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f85b9206c9100d9eb9baf0852b24c423a26433917f3d3ea3b60658e71ae4d91"></a>

## eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface / 43f833bf41a3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-340e4c649ff6afcaa1613dc753fd944019206e5d83b7b590fcd46b9db528c1c7"></a>

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

<a id="canonical-978591d521f8b7613f7a7675dd1fd5bd0e72a3209c498c1fb275c2331d17c281"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface / 43f833bf41a3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b2fb6ca07ebf5e9d47faae26da10d58dfbff1a24027a01b27898bef6779fe37e"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface / 43f833bf41a3 / 4

- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-75f691788d847ec3be5898614a862e1c1d35cd4ead0dc9c7b666d8e1338250ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0dced427be4299b38ec63b3c9df055d0d935d6bec10157c51129bacfe426e041"></a>

## eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface / fb04b7406626 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-d09859a534350dcc6391016b3a5367cc807283bfb05720ea7171dc51509389d2"></a>

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

<a id="canonical-ca24d9c6a98422c3a5f68a54532c603cb62757650da42c8a91902cd74e49f965"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface / fb04b7406626 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-00361dd8e1d4645d39435b926ccebcffd9814aaaed11bd96ecf17218400685a2"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface / fb04b7406626 / 4

- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-082c90448b4143a96a03cda9f9bc496d6b3ed9ea44783d38c1c9a90858b74488"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40f30ee78ffb20d439077afc5cb41ddc6e43403f8f4f1d30e36317018b00fc90"></a>

## eks_k8s.not_managed.node_list.interface_list.static_ip — eks_k8s.not_managed.node_list.interface_list.static_ip / 1bdb00514b7d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- eks_k8s.not_managed.node_list.interface_list.static_ip

<a id="canonical-c2790dc3475a003dcbd2d560a7fdf3faeb2e70e530b70d95a6ea8957aa7bbf68"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1f86cba0381b80410cadd53489b481f19ebf32249abe3a244b792c8caf479976"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.static_ip / 1bdb00514b7d / 3

<a id="canonical-b4e51290541d99b650af8f2bf21f71f90b927e675f673aee7a9e4c8a9ab03804"></a>

<a id="canonical-5b202e101e235589dac04699bea7a955f3ade534bd104a98b9cf0cedce75bdaf"></a>

## default_gw property — eks_k8s.not_managed.node_list.interface_list.static_ip / 1bdb00514b7d / 4

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-cfaf17446fae87bd4d40b7e0d482e2f12ed0af84660fa2a2c4141f7aabe643a9"></a>
