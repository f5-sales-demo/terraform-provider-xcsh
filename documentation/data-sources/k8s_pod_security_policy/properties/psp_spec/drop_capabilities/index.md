---
page_title: "psp_spec.drop_capabilities"
subcategory: ""
description: "List of capabilities that docker container has."
xcsh_docs: {"aliases": ["psp spec drop capabilities"], "body_bytes": 2298, "body_sha256": "sha256:b3e5448d4cca1104f446dbef321a3c7e1cbbe69367f3f2a49fd5bafd1903dd8c", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:drop_capabilities", "parent_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "path": "documentation/data-sources/k8s_pod_security_policy/properties/psp_spec/drop_capabilities/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0300011102223101-3030213220010330-3220230101032201-0032331022322000-3020201202311233-1121013210030111-2303020101303201-0301120300011230", "registry_path": "docs/guides/data-sources--k8s_pod_security_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["psp_spec", "drop_capabilities"], "schema_version": 1, "sections": [{"aliases": ["capabilities"], "anchor": "schema-psp_spec--drop_capabilities--capabilities", "description": "List of capabilities that docker container has.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:drop_capabilities", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "drop_capabilities", "capabilities"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_policy/properties/psp_spec/drop_capabilities/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of capabilities that docker container has.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# psp_spec.drop_capabilities

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/)
- [psp_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/)
- psp_spec.drop_capabilities

<a id="section"></a>

Type: `"single"`. Computed.

List of capabilities that docker container has.

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

<a id="schema-psp_spec--drop_capabilities--capabilities"></a>

### capabilities property

Type: `["list", "string"]`. Computed.

List of capabilities that docker container has.

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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [psp_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/)
- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/)
