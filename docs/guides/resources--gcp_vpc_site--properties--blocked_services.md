---
page_title: "blocked_services"
subcategory: "Infrastructure"
description: "blocked_services for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1188, "body_sha256": "sha256:3c3405a3349d4576fad5b9230a5ed318f0bc41e6266c1deae987eff760a5fa19", "canonical_id": "xcsh-docs:resources:gcp_vpc_site:properties:blocked_services", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:blocked_services:blocked_service"], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:blocked_services", "parent_id": "xcsh-docs:resources:gcp_vpc_site:reference", "path": "docs/guides/resources--gcp_vpc_site--properties--blocked_services.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["blocked_services"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/blocked_services/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "blocked_services for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# blocked_services

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
- [Property reference](resources--gcp_vpc_site--reference.md)
- blocked_services

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blocked_services {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blocked_service](resources--gcp_vpc_site--properties--blocked_services--blocked_service.md): complete subsection reference.

## Next pages

- [blocked_services.blocked_service](resources--gcp_vpc_site--properties--blocked_services--blocked_service.md)
- [Property reference](resources--gcp_vpc_site--reference.md)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
