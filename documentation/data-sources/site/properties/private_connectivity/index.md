---
page_title: "private_connectivity"
subcategory: "Infrastructure"
description: "Private Connectivity Information like ADN network name and cloud link information."
xcsh_docs: {"aliases": ["private connectivity"], "body_bytes": 1246, "body_sha256": "sha256:f5ead5d467d2a4c13919193c341eb6852b6d9cc62da9d6ff49a65b36460a1e5a", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:site:properties:private_connectivity:cloud_link"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:private_connectivity", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "documentation/data-sources/site/properties/private_connectivity/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3001111232010023-2110131311311201-1132001023210300-3031100301300220-0323023311033033-1201003100322301-0100022023112022-2212112220301111", "registry_path": "docs/guides/data-sources--site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["private_connectivity"], "schema_version": 1, "sections": [{"aliases": ["private connectivity cloud link"], "anchor": "section", "description": "Information related to cloud link used by the site.", "document_id": "xcsh-docs:data-sources:site:properties:private_connectivity:cloud_link", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["private_connectivity", "cloud_link"], "syntax": "attribute", "type": "object"}, {"aliases": ["private connectivity private network name"], "anchor": "schema-private_connectivity--private_network_name", "description": "ADN Network Name for private access connectivity to F5XC ADN.", "document_id": "xcsh-docs:data-sources:site:properties:private_connectivity", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["private_connectivity", "private_network_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/private_connectivity/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Private Connectivity Information like ADN network name and cloud link information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

## Next pages

- [private_connectivity.cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/private_connectivity/cloud_link/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/)
- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
