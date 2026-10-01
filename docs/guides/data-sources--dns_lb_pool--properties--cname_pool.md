---
page_title: "cname_pool"
subcategory: ""
description: "cname_pool for xcsh_dns_lb_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1406, "body_sha256": "sha256:3ebf3a74fdab825a7ff2ddc289466ad52cd19ebd8c56ffc6e9901431a6b875a9", "canonical_id": "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool", "child_ids": ["xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:disable_health_check", "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:health_check", "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:members"], "collection_id": "xcsh-docs:data-sources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool", "parent_id": "xcsh-docs:data-sources:dns_lb_pool:reference", "path": "docs/guides/data-sources--dns_lb_pool--properties--cname_pool.md", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cname_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_pool/properties/cname_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cname_pool for xcsh_dns_lb_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cname_pool

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md)
- [Property reference](data-sources--dns_lb_pool--reference.md)
- cname_pool

<a id="section"></a>

Type: `"single"`. Computed.

Pool for CNAME Record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_choice": "[\"disable_health_check\",\"health_check\"]"
}
```

## Direct properties

- [disable_health_check](data-sources--dns_lb_pool--properties--cname_pool--disable_health_check.md): complete subsection reference.

- [health_check](data-sources--dns_lb_pool--properties--cname_pool--health_check.md): complete subsection reference.

- [members](data-sources--dns_lb_pool--properties--cname_pool--members.md): complete subsection reference.

## Next pages

- [cname_pool.disable_health_check](data-sources--dns_lb_pool--properties--cname_pool--disable_health_check.md)
- [cname_pool.health_check](data-sources--dns_lb_pool--properties--cname_pool--health_check.md)
- [cname_pool.members](data-sources--dns_lb_pool--properties--cname_pool--members.md)
- [Property reference](data-sources--dns_lb_pool--reference.md)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md)
