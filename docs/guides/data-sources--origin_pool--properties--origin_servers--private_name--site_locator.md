---
page_title: "origin_servers.private_name.site_locator"
subcategory: "Load Balancing"
description: "origin_servers.private_name.site_locator for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1619, "body_sha256": "sha256:1f2c569e6f45cab0516b3944db261c4a6fb9d9892eca488c64ae74597f8d50f6", "canonical_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_name:site_locator", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_name:site_locator:site", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_name:site_locator:virtual_site"], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_name:site_locator", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_name", "path": "docs/guides/data-sources--origin_pool--properties--origin_servers--private_name--site_locator.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "private_name", "site_locator"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/origin_servers/private_name/site_locator/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.private_name.site_locator for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# origin_servers.private_name.site_locator

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Property reference](data-sources--origin_pool--reference.md)
- [origin_servers](data-sources--origin_pool--properties--origin_servers.md)
- [origin_servers.private_name](data-sources--origin_pool--properties--origin_servers--private_name.md)
- origin_servers.private_name.site_locator

<a id="section"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

## Direct properties

- [site](data-sources--origin_pool--properties--origin_servers--private_name--site_locator--site.md): complete subsection reference.

- [virtual_site](data-sources--origin_pool--properties--origin_servers--private_name--site_locator--virtual_site.md): complete subsection reference.

## Next pages

- [origin_servers.private_name.site_locator.site](data-sources--origin_pool--properties--origin_servers--private_name--site_locator--site.md)
- [origin_servers.private_name.site_locator.virtual_site](data-sources--origin_pool--properties--origin_servers--private_name--site_locator--virtual_site.md)
- [origin_servers.private_name](data-sources--origin_pool--properties--origin_servers--private_name.md)
- [xcsh_origin_pool](../data-sources/origin_pool.md)
