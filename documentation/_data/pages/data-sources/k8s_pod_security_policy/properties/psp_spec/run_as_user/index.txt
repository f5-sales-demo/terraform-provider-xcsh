---
page_title: "psp_spec.run_as_user"
subcategory: ""
description: "ID ranges and rules."
xcsh_docs: {"aliases": ["psp spec run as user"], "body_bytes": 2043, "body_sha256": "sha256:eeb309e45cfa397e5d435ee0b6b731f966b253d8303799a5aa356283465a6095", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:run_as_user:id_ranges"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:run_as_user", "parent_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "path": "documentation/data-sources/k8s_pod_security_policy/properties/psp_spec/run_as_user/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-2211301322001300-1211102230212212-3222112002333031-0321030003031212-2312222233032220-2020331210031013-2231123302321131-2103033201020311", "registry_path": "docs/guides/data-sources--k8s_pod_security_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["psp_spec", "run_as_user"], "schema_version": 1, "sections": [{"aliases": ["psp spec run as user id ranges"], "anchor": "section", "description": "List of range of ID(s)", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:run_as_user:id_ranges", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["psp_spec", "run_as_user", "id_ranges"], "syntax": "attribute", "type": "object"}, {"aliases": ["psp spec run as user rule"], "anchor": "schema-psp_spec--run_as_user--rule", "description": "Rule indicated how the FS group ID range is used.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:run_as_user", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "run_as_user", "rule"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_policy/properties/psp_spec/run_as_user/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "ID ranges and rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# psp_spec.run_as_user

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/)
- [psp_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/)
- psp_spec.run_as_user

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for run as user.

Additional upstream details:

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

- [id_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/run_as_user/id_ranges/): complete subsection reference.

<a id="schema-psp_spec--run_as_user--rule"></a>

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
      "validatedAt": "2026-10-29T12:14:42+00:00"
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
