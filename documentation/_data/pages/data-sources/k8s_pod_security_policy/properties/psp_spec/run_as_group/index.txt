---
page_title: "psp_spec.run_as_group"
subcategory: ""
description: "ID ranges and rules."
xcsh_docs: {"aliases": ["psp spec run as group"], "body_bytes": 2487, "body_sha256": "sha256:40327375783516fa9382d50db4161b73aac950d8c2414f121cb6c87efd05707f", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:run_as_group:id_ranges"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:run_as_group", "parent_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "path": "documentation/data-sources/k8s_pod_security_policy/properties/psp_spec/run_as_group/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1000330010132213-0013201221322322-3310132230301122-0231103202203321-3203000002330032-0023030100013121-3203212122101103-2331010222302333", "registry_path": "docs/guides/data-sources--k8s_pod_security_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["psp_spec", "run_as_group"], "schema_version": 1, "sections": [{"aliases": ["psp spec run as group id ranges"], "anchor": "section", "description": "List of range of ID(s)", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:run_as_group:id_ranges", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["psp_spec", "run_as_group", "id_ranges"], "syntax": "attribute", "type": "object"}, {"aliases": ["psp spec run as group rule"], "anchor": "schema-psp_spec--run_as_group--rule", "description": "Rule indicated how the FS group ID range is used.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:run_as_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "run_as_group", "rule"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_policy/properties/psp_spec/run_as_group/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "ID ranges and rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# psp_spec.run_as_group

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/)
- [psp_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/)
- psp_spec.run_as_group

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for run as group.

Upstream description:

ID ranges and rules.

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

## Direct properties

- [id_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/run_as_group/id_ranges/): complete subsection reference.

<a id="schema-psp_spec--run_as_group--rule"></a>

### rule property

Type: `"string"`. Computed.

Rule indicated how the FS group ID range is used.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

## Next pages

- [psp_spec.run_as_group.id_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/run_as_group/id_ranges/)
- [psp_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/)
- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/)
