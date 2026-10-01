---
page_title: "blocked_services"
subcategory: ""
description: "blocked_services for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1497, "body_sha256": "sha256:b994c58786a9a9a915d235fdb28ab9dfb05f887fce52b6b2dd1aa5080c289f68", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:blocked_services", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:blocked_services:blocked_service"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:blocked_services", "parent_id": "xcsh-docs:data-sources:voltstack_site:reference", "path": "docs/guides/data-sources--voltstack_site--properties--blocked_services.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["blocked_services"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/blocked_services/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "blocked_services for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# blocked_services

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- blocked_services

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: blocked\_services, default\_blocked\_services; Default: default\_blocked\_services\]
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

OneOf alternatives in this subsection:

- [blocked_services](data-sources--voltstack_site--properties--blocked_services.md#section)
- [default_blocked_services](data-sources--voltstack_site--properties--default_blocked_services.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [blocked_service](data-sources--voltstack_site--properties--blocked_services--blocked_service.md): complete subsection reference.

## Next pages

- [blocked_services.blocked_service](data-sources--voltstack_site--properties--blocked_services--blocked_service.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
