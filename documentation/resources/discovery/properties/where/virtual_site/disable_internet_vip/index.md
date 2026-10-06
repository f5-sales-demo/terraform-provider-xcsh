---
page_title: "where.virtual_site.disable_internet_vip"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["where virtual site disable internet vip"], "body_bytes": 1119, "body_sha256": "sha256:7e6428b2885af9bea021a53310353baba9d35f429a4a0c4a18cf65389206f8c7", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:where:virtual_site:disable_internet_vip", "parent_id": "xcsh-docs:resources:discovery:properties:where:virtual_site", "path": "documentation/resources/discovery/properties/where/virtual_site/disable_internet_vip/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0012000002320212-1110003030333201-2200132323301111-0031003322022332-2131310131100122-1020102030213221-1322011213212030-0212233130001000", "registry_path": "docs/guides/resources--discovery--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["where", "virtual_site", "disable_internet_vip"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/where/virtual_site/disable_internet_vip/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["discoveryCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where.virtual_site.disable_internet_vip

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/)
- [where.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_site/)
- where.virtual_site.disable_internet_vip

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
disable_internet_vip = {}
```

This is an empty object or choice marker. It has no direct properties.
