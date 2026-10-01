---
page_title: "https_management.advertise_on_internet"
subcategory: ""
description: "https_management.advertise_on_internet for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 1440, "body_sha256": "sha256:46241c7791d09c06cc9a663ed44a4174f170f95ea6acc5aaac5206eef07d4ee2", "canonical_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_internet", "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_internet:public_ip"], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_internet", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:https_management", "path": "docs/guides/data-sources--nfv_service--properties--https_management--advertise_on_internet.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_management", "advertise_on_internet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/https_management/advertise_on_internet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_management.advertise_on_internet for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management.advertise_on_internet

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md)
- [Property reference](data-sources--nfv_service--reference.md)
- [https_management](data-sources--nfv_service--properties--https_management.md)
- https_management.advertise_on_internet

<a id="section"></a>

Type: `"single"`. Computed.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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

- [public_ip](data-sources--nfv_service--properties--https_management--advertise_on_internet--public_ip.md): complete subsection reference.

## Next pages

- [https_management.advertise_on_internet.public_ip](data-sources--nfv_service--properties--https_management--advertise_on_internet--public_ip.md)
- [https_management](data-sources--nfv_service--properties--https_management.md)
- [xcsh_nfv_service](../data-sources/nfv_service.md)
