---
page_title: "main_nodes"
subcategory: "Infrastructure"
description: "Connectivity information of main/master nodes to create a full mesh of Phobos services across all CEs in a site-mesh-group or dc-cluster-group."
xcsh_docs: {"aliases": ["main nodes"], "body_bytes": 1229, "body_sha256": "sha256:628f7dbf84781463d683c7f0a9546af2047aa83b261c1dcaacf456a2c76df05c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:main_nodes", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "documentation/data-sources/site/properties/main_nodes/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3212112120010131-2002332321130213-3322012031333310-0120331130321030-1202311023111202-1211011121110110-2222233233322233-2132132020113122", "registry_path": "docs/guides/data-sources--site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["main_nodes"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-main_nodes--name", "description": "Name of the master/main node on the site.", "document_id": "xcsh-docs:data-sources:site:properties:main_nodes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["main_nodes", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["sli address"], "anchor": "schema-main_nodes--sli_address", "description": "Site Local Inside IP addresses. Site Local Inside IP address.", "document_id": "xcsh-docs:data-sources:site:properties:main_nodes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["main_nodes", "sli_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["slo address"], "anchor": "schema-main_nodes--slo_address", "description": "Site Local Outside IP addresses. Site Local Outside IP address.", "document_id": "xcsh-docs:data-sources:site:properties:main_nodes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["main_nodes", "slo_address"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/main_nodes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Connectivity information of main/master nodes to create a full mesh of Phobos services across all CEs in a site-mesh-group or dc-cluster-group.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# main_nodes

Breadcrumbs:

- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/)
- main_nodes

<a id="section"></a>

Type: `"list"`. Computed.

Connectivity information of main/master nodes to create a full mesh of Phobos services across all
CEs in a site-mesh-group or dc-cluster-group.

## Direct properties

<a id="schema-main_nodes--name"></a>

### name property

Type: `"string"`. Computed.

Name of the master/main node on the site.

<a id="schema-main_nodes--sli_address"></a>

### sli_address property

Type: `"string"`. Computed.

Site Local Inside IP addresses. Site Local Inside IP address.

<a id="schema-main_nodes--slo_address"></a>

### slo_address property

Type: `"string"`. Computed.

Site Local Outside IP addresses. Site Local Outside IP address.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/)
- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
