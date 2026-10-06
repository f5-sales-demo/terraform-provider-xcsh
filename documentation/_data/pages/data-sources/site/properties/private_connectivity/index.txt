---
page_title: "private_connectivity"
subcategory: "Infrastructure"
description: "Private Connectivity Information like ADN network name and cloud link information."
xcsh_docs: {"aliases": ["private connectivity"], "body_bytes": 875, "body_sha256": "sha256:4967b6e6a16dfb7828a4e0716bc97a7b4563b8075d2f3071ff320d10694617d7", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:site:properties:private_connectivity:cloud_link"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:private_connectivity", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "documentation/data-sources/site/properties/private_connectivity/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3001111232010023-2110131311311201-1132001023210300-3031100301300220-0323023311033033-1201003100322301-0100022023112022-2212112220301111", "registry_path": "docs/guides/data-sources--site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["private_connectivity"], "schema_version": 1, "sections": [{"aliases": ["private connectivity cloud link"], "anchor": "section", "description": "Information related to cloud link used by the site.", "document_id": "xcsh-docs:data-sources:site:properties:private_connectivity:cloud_link", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["private_connectivity", "cloud_link"], "syntax": "attribute", "type": "object"}, {"aliases": ["private connectivity private network name"], "anchor": "schema-private_connectivity--private_network_name", "description": "ADN Network Name for private access connectivity to F5XC ADN.", "document_id": "xcsh-docs:data-sources:site:properties:private_connectivity", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["private_connectivity", "private_network_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/private_connectivity/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Private Connectivity Information like ADN network name and cloud link information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# private_connectivity

Breadcrumbs:

- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/)
- private_connectivity

<a id="section"></a>

Type: `"single"`. Computed.

Private Connectivity Information like ADN network name and cloud link information.

## Direct properties

- [cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/private_connectivity/cloud_link/): complete subsection reference.

<a id="schema-private_connectivity--private_network_name"></a>

### private_network_name property

Type: `"string"`. Computed.

ADN Network Name for private access connectivity to F5XC ADN.
