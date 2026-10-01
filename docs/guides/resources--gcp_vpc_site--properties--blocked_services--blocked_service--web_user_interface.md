---
page_title: "blocked_services.blocked_service.web_user_interface"
subcategory: "Infrastructure"
description: "blocked_services.blocked_service.web_user_interface for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1196, "body_sha256": "sha256:b8aeebdb06d79b6c8b436ff3c036f341a79643e8b1aeb01a212162d82e32312f", "canonical_id": "xcsh-docs:resources:gcp_vpc_site:properties:blocked_services:blocked_service:web_user_interface", "child_ids": [], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:blocked_services:blocked_service:web_user_interface", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:blocked_services:blocked_service", "path": "docs/guides/resources--gcp_vpc_site--properties--blocked_services--blocked_service--web_user_interface.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["blocked_services", "blocked_service", "web_user_interface"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/blocked_services/blocked_service/web_user_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "blocked_services.blocked_service.web_user_interface for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# blocked_services.blocked_service.web_user_interface

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
- [Property reference](resources--gcp_vpc_site--reference.md)
- [blocked_services](resources--gcp_vpc_site--properties--blocked_services.md)
- [blocked_services.blocked_service](resources--gcp_vpc_site--properties--blocked_services--blocked_service.md)
- blocked_services.blocked_service.web_user_interface

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
web_user_interface = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [blocked_services.blocked_service](resources--gcp_vpc_site--properties--blocked_services--blocked_service.md)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
