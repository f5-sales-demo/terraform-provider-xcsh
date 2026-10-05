---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_kubernetes_manifests."
xcsh_docs: {"aliases": ["kubernetes manifests"], "body_bytes": 2390, "body_sha256": "sha256:df20ddd69a6a4b34797275ba4407305028aeaa22ce8277ac7bd874149dcfb292", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:ephemeral-resources:kubernetes_manifests:properties:extensions"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:collection", "completeness": "complete", "id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:reference", "parent_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:fundamentals", "path": "documentation/ephemeral-resources/kubernetes_manifests/properties/index.md", "product": "distributed-cloud", "provider_name": "kubernetes_manifests", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "ephemeral-resources", "registry_anchor": "canonical-0222301022001323-1002221032030331-0321231112233121-0131212101033121-3023002220100301-3312321000102011-3211011101131320-3232112012103020", "registry_path": "docs/guides/ephemeral-resources--kubernetes_manifests--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["content type"], "anchor": "schema-content_type", "description": "The HTTP Content-Type header value specifying the content type of the body.", "document_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:reference", "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["content_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["data"], "anchor": "schema-data", "description": "The HTTP request/response body as raw binary.", "document_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:reference", "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data"], "syntax": "attribute", "type": "string"}, {"aliases": ["extensions"], "anchor": "section", "description": "Application specific response metadata. Must be set in the first response for streaming APIs.", "document_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:properties:extensions", "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["extensions"], "syntax": "attribute", "type": "object"}, {"aliases": ["site"], "anchor": "schema-site", "description": "Site Name. Name of the site.", "document_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/ephemeral-resources/kubernetes_manifests/properties/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Property reference for xcsh_kubernetes_manifests.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

## Next pages

- [extensions](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/properties/extensions/)
- [xcsh_kubernetes_manifests](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/)
