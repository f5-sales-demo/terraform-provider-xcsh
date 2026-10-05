---
page_title: "extensions"
subcategory: ""
description: "Application specific response metadata. Must be set in the first response for streaming APIs."
xcsh_docs: {"aliases": ["extensions"], "body_bytes": 1378, "body_sha256": "sha256:571e444a0af96f02c095cf569427e23085cf0158e5002ae1d3ca0268bc3d907d", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:collection", "completeness": "complete", "id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:properties:extensions", "parent_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:reference", "path": "documentation/ephemeral-resources/kubernetes_manifests/properties/extensions/index.md", "product": "distributed-cloud", "provider_name": "kubernetes_manifests", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "ephemeral-resources", "registry_anchor": "canonical-1212012223132313-3222203233201021-3001002322011022-1321320013031011-2032331313030022-3332122302110221-3231333221313332-2222303212212001", "registry_path": "docs/guides/ephemeral-resources--kubernetes_manifests--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["extensions"], "schema_version": 1, "sections": [{"aliases": ["extensions type url"], "anchor": "schema-extensions--type_url", "description": "URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This string must contain at least one '/' character. The last segment of the URL path must represent the fully qualified name of the type (as in ).", "document_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:properties:extensions", "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["extensions", "type_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["extensions value"], "anchor": "schema-extensions--value", "description": "Must be a valid serialized protocol buffer of the above specified type.", "document_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:properties:extensions", "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["extensions", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/ephemeral-resources/kubernetes_manifests/properties/extensions/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Application specific response metadata. Must be set in the first response for streaming APIs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# extensions

Breadcrumbs:

- [xcsh_kubernetes_manifests](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/properties/)
- extensions

<a id="section"></a>

Type: `"list"`. Computed, Sensitive.

Application specific response metadata. Must be set in the first response for streaming APIs.

## Direct properties

<a id="schema-extensions--type_url"></a>

### type_url property

Type: `"string"`. Computed, Sensitive.

URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This
string must contain at least one '/' character. The last segment of the URL path must represent the
fully qualified name of the type (as in ).

<a id="schema-extensions--value"></a>

### value property

Type: `"string"`. Computed, Sensitive.

Must be a valid serialized protocol buffer of the above specified type.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/properties/)
- [xcsh_kubernetes_manifests](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/)
