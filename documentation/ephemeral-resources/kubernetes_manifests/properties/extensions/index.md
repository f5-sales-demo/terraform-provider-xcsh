---
page_title: "extensions"
subcategory: ""
description: "Application specific response metadata. Must be set in the first response for streaming APIs."
xcsh_docs: {"aliases": ["extensions"], "body_bytes": 1100, "body_sha256": "sha256:6dc9cb1088d4c375fd62b97a493acd7d0dba19c0f35a2d08a15177f257e5cb2e", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:collection", "completeness": "complete", "id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:properties:extensions", "parent_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:reference", "path": "documentation/ephemeral-resources/kubernetes_manifests/properties/extensions/index.md", "product": "distributed-cloud", "provider_name": "kubernetes_manifests", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "ephemeral-resources", "registry_anchor": "canonical-1212012223132313-3222203233201021-3001002322011022-1321320013031011-2032331313030022-3332122302110221-3231333221313332-2222303212212001", "registry_path": "docs/guides/ephemeral-resources--kubernetes_manifests--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["extensions"], "schema_version": 1, "sections": [{"aliases": ["extensions type url"], "anchor": "schema-extensions--type_url", "description": "URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This string must contain at least one '/' character. The last segment of the URL path must represent the fully qualified name of the type (as in ).", "document_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:properties:extensions", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["extensions", "type_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["extensions value"], "anchor": "schema-extensions--value", "description": "Must be a valid serialized protocol buffer of the above specified type.", "document_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:properties:extensions", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["extensions", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/ephemeral-resources/kubernetes_manifests/properties/extensions/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Application specific response metadata. Must be set in the first response for streaming APIs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": [], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
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
