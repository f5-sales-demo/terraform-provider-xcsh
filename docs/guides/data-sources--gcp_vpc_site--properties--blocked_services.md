---
page_title: "blocked_services"
subcategory: "Infrastructure"
description: "blocked_services for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 982, "body_sha256": "sha256:365e8b29fc4599d2b8e7d4301a3cf7e40326e27444632c6276bb57196f763652", "canonical_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:blocked_services", "child_ids": ["xcsh-docs:data-sources:gcp_vpc_site:properties:blocked_services:blocked_service"], "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:blocked_services", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:reference", "path": "docs/guides/data-sources--gcp_vpc_site--properties--blocked_services.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["blocked_services"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/blocked_services/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "blocked_services for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# blocked_services

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
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

- [blocked_service](data-sources--gcp_vpc_site--properties--blocked_services--blocked_service.md): complete subsection reference.

## Next pages

- [blocked_services.blocked_service](data-sources--gcp_vpc_site--properties--blocked_services--blocked_service.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
