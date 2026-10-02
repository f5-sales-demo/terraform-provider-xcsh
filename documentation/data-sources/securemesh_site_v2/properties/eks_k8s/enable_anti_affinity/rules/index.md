---
page_title: "eks_k8s.enable_anti_affinity.rules"
subcategory: ""
description: "Define one or more anti-affinity rules. Each rule specifies which pods (by labels) should be distributed across which topology domains. Example: Rule 1 - Distribute VPM pods across nodes, Rule 2 - Distribute Prometheus pods across zones."
xcsh_docs: {"aliases": ["eks k8s enable anti affinity rules"], "body_bytes": 6890, "body_sha256": "sha256:88817673e1c3e4cf9ee0ba5b734436c332ae7c8e263b0226df8c749150869535", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity", "path": "documentation/data-sources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/rules/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3231321300323220-1103220122202303-3312300300332123-2000031302101130-2312002333000331-3133001233302121-2032123112102223-0102210303202212", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["eks_k8s", "enable_anti_affinity", "rules"], "schema_version": 1, "sections": [{"aliases": ["label key"], "anchor": "schema-eks_k8s--enable_anti_affinity--rules--label_key", "description": "Specify the label key of the customer pods that CE pods should avoid being co-scheduled with. Combined with the label value below, this identifies the target pods.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["eks_k8s", "enable_anti_affinity", "rules", "label_key"], "syntax": "attribute", "type": "string"}, {"aliases": ["label value"], "anchor": "schema-eks_k8s--enable_anti_affinity--rules--label_value", "description": "Specify the label value that, together with the label key, identifies the customer pods to avoid.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["eks_k8s", "enable_anti_affinity", "rules", "label_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["topology keys"], "anchor": "schema-eks_k8s--enable_anti_affinity--rules--topology_keys", "description": "Specify one or more node label keys that define the scope of avoidance. For each topology key (e.g., Kubernetes.I/O/hostname), CE pods will avoid nodes whose topology value matches a node already running a pod with the above specified label. Example: with Kubernetes.I/O/hostname, CE pods are kept off any node running", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:enable_anti_affinity:rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["eks_k8s", "enable_anti_affinity", "rules", "topology_keys"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/rules/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Define one or more anti-affinity rules. Each rule specifies which pods (by labels) should be distributed across which topology domains. Example: Rule 1 - Distribute VPM pods across nodes, Rule 2 - Distribute Prometheus pods across zones.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# eks_k8s.enable_anti_affinity.rules

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [eks_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/)
- [eks_k8s.enable_anti_affinity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/)
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

- [eks_k8s.enable_anti_affinity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/enable_anti_affinity/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
