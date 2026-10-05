---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_site_signatures_update."
xcsh_docs: {"aliases": ["site signatures update"], "body_bytes": 1258, "body_sha256": "sha256:fa57919059f176866da0ae44b5f07b33f9eef977a57c35fd6b4d04b24fbf3f1c", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:site_signatures_update:collection", "completeness": "complete", "id": "xcsh-docs:actions:site_signatures_update:reference", "parent_id": "xcsh-docs:actions:site_signatures_update:fundamentals", "path": "documentation/actions/site_signatures_update/properties/index.md", "product": "distributed-cloud", "provider_name": "site_signatures_update", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "actions", "registry_anchor": "canonical-2103311202213331-3322033031022121-3121001201130323-0111123021202301-2313002222033000-2333220012121133-2230132001320330-3233213301303202", "registry_path": "docs/guides/actions--site_signatures_update--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace Namespace of the site to update signatures for.", "document_id": "xcsh-docs:actions:site_signatures_update:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["site names"], "anchor": "schema-site_names", "description": "List of site names to update. If empty, no sites will be updated.", "document_id": "xcsh-docs:actions:site_signatures_update:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_names"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_signatures_update/properties/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Property reference for xcsh_site_signatures_update.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_site_signatures_update](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_signatures_update/)
- Property reference

## Direct properties

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace Namespace of the site to update signatures for.

<a id="schema-site_names"></a>

### site_names property

Type: `["list", "string"]`. Optional.

List of site names to update. If empty, no sites will be updated.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_signatures_update/properties/#schema-namespace) |
| `site_names` | [site_names](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_signatures_update/properties/#schema-site_names) |

## Next pages

- [xcsh_site_signatures_update](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_signatures_update/)
