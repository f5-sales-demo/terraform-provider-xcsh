---
page_title: "blocked_services"
subcategory: "Infrastructure"
description: "blocked_services for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1105, "body_sha256": "sha256:cc3307ad3119e51fd3082e660a4d953fd99d06c117b9153eb7767ad950591197", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:blocked_services", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:blocked_services:blocked_service"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:blocked_services", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "path": "docs/guides/data-sources--azure_vnet_site--properties--blocked_services.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["blocked_services"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/blocked_services/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "blocked_services for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# blocked_services

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- blocked_services

<a id="section"></a>

Type: `"single"`. Computed.

Disable node local services on this site.

Upstream description:

Disable node local services on this site. Note: The chosen services will GET disabled on all nodes
in the site.

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

## Direct properties

- [blocked_service](data-sources--azure_vnet_site--properties--blocked_services--blocked_service.md): complete subsection reference.

## Next pages

- [blocked_services.blocked_service](data-sources--azure_vnet_site--properties--blocked_services--blocked_service.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
