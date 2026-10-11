---
page_title: "main_nodes"
subcategory: "Infrastructure"
description: "Connectivity information of main/master nodes to create a full mesh of Phobos services across all CEs in a site-mesh-group or dc-cluster-group."
xcsh_docs: {"aliases": ["main nodes"], "body_bytes": 1013, "body_sha256": "sha256:d59875b270abccba9bcee3452b4419480df3a54129427359edfb216d7bd83ca4", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:main_nodes", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "documentation/data-sources/site/properties/main_nodes/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-3212112120010131-2002332321130213-3322012031333310-0120331130321030-1202311023111202-1211011121110110-2222233233322233-2132132020113122", "registry_path": "docs/guides/data-sources--site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["main_nodes"], "schema_version": 1, "sections": [{"aliases": ["main nodes name"], "anchor": "schema-main_nodes--name", "description": "Name of the master/main node on the site.", "document_id": "xcsh-docs:data-sources:site:properties:main_nodes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["main_nodes", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["main nodes sli address"], "anchor": "schema-main_nodes--sli_address", "description": "Site Local Inside IP addresses. Site Local Inside IP address.", "document_id": "xcsh-docs:data-sources:site:properties:main_nodes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["main_nodes", "sli_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["main nodes slo address"], "anchor": "schema-main_nodes--slo_address", "description": "Site Local Outside IP addresses. Site Local Outside IP address.", "document_id": "xcsh-docs:data-sources:site:properties:main_nodes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["main_nodes", "slo_address"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/main_nodes/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Connectivity information of main/master nodes to create a full mesh of Phobos services across all CEs in a site-mesh-group or dc-cluster-group.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
