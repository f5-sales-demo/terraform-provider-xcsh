---
page_title: "eks_k8s.enable_anti_affinity.rules"
subcategory: ""
description: "eks_k8s.enable_anti_affinity.rules for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 6485, "body_sha256": "sha256:17279f6f446425a66051ce2524a19199082a72d23a3767c8e5764838e6544aae", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "child_ids": [], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity", "path": "docs/guides/data-sources--securemesh_site_v2--properties--eks_k8s--enable_anti_affinity--rules.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["eks_k8s", "enable_anti_affinity", "rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "eks_k8s.enable_anti_affinity.rules for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# eks_k8s.enable_anti_affinity.rules

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [eks_k8s](data-sources--securemesh_site_v2--properties--eks_k8s.md)
- [eks_k8s.enable_anti_affinity](data-sources--securemesh_site_v2--properties--eks_k8s--enable_anti_affinity.md)
- eks_k8s.enable_anti_affinity.rules

<a id="section"></a>

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

## Direct properties

<a id="schema-eks_k8s--enable_anti_affinity--rules--label_key"></a>

### label_key property

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

<a id="schema-eks_k8s--enable_anti_affinity--rules--label_value"></a>

### label_value property

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

<a id="schema-eks_k8s--enable_anti_affinity--rules--topology_keys"></a>

### topology_keys property

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

## Next pages

- [eks_k8s.enable_anti_affinity](data-sources--securemesh_site_v2--properties--eks_k8s--enable_anti_affinity.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
