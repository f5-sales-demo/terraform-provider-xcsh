---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_kubernetes_manifests."
xcsh_docs: {"aliases": ["kubernetes manifests"], "body_bytes": 2109, "body_sha256": "sha256:422280ac4c65b993da4c3d3858037c13de80e6f80cd7d7cfabb6dc3bc7a38c7b", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:ephemeral-resources:kubernetes_manifests:properties:extensions"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:collection", "completeness": "complete", "id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:reference", "parent_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:fundamentals", "path": "documentation/ephemeral-resources/kubernetes_manifests/properties/index.md", "product": "distributed-cloud", "provider_name": "kubernetes_manifests", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "ephemeral-resources", "registry_anchor": "canonical-0222301022001323-1002221032030331-0321231112233121-0131212101033121-3023002220100301-3312321000102011-3211011101131320-3232112012103020", "registry_path": "docs/guides/ephemeral-resources--kubernetes_manifests--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["content type"], "anchor": "schema-content_type", "description": "The HTTP Content-Type header value specifying the content type of the body.", "document_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["content_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["data"], "anchor": "schema-data", "description": "The HTTP request/response body as raw binary.", "document_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data"], "syntax": "attribute", "type": "string"}, {"aliases": ["extensions"], "anchor": "section", "description": "Application specific response metadata. Must be set in the first response for streaming APIs.", "document_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:properties:extensions", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["extensions"], "syntax": "attribute", "type": "object"}, {"aliases": ["site"], "anchor": "schema-site", "description": "Site Name. Name of the site.", "document_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/ephemeral-resources/kubernetes_manifests/properties/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Property reference for xcsh_kubernetes_manifests.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_kubernetes_manifests](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/)
- Property reference

## Direct properties

<a id="schema-content_type"></a>

### content_type property

Type: `"string"`. Computed, Sensitive.

The HTTP Content-Type header value specifying the content type of the body.

<a id="schema-data"></a>

### data property

Type: `"string"`. Computed, Sensitive.

The HTTP request/response body as raw binary.

- [extensions](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/properties/extensions/): complete subsection reference.

<a id="schema-site"></a>

### site property

Type: `"string"`. Required.

Site Name. Name of the site.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `content_type` | [content_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/properties/#schema-content_type) |
| `data` | [data](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/properties/#schema-data) |
| `extensions` | [extensions](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/properties/extensions/#section) |
| `extensions.type_url` | [extensions.type_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/properties/extensions/#schema-extensions--type_url) |
| `extensions.value` | [extensions.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/properties/extensions/#schema-extensions--value) |
| `site` | [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/properties/#schema-site) |
