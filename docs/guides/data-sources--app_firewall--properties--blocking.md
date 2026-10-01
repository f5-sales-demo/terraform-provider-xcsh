---
page_title: "blocking"
subcategory: "Security"
description: "blocking for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1110, "body_sha256": "sha256:95609d63d8fabfcd18bbd1b3ca9b5da6448d9edc06f5004bd599edfff8ff1a09", "canonical_id": "xcsh-docs:data-sources:app_firewall:properties:blocking", "child_ids": [], "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:blocking", "parent_id": "xcsh-docs:data-sources:app_firewall:reference", "path": "docs/guides/data-sources--app_firewall--properties--blocking.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["blocking"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/blocking/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "blocking for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# blocking

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md)
- [Property reference](data-sources--app_firewall--reference.md)
- blocking

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: blocking, monitoring\] Enable this option

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

OneOf alternatives in this subsection:

- [blocking](data-sources--app_firewall--properties--blocking.md#section)
- [monitoring](data-sources--app_firewall--properties--monitoring.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--app_firewall--reference.md)
- [xcsh_app_firewall](../data-sources/app_firewall.md)
