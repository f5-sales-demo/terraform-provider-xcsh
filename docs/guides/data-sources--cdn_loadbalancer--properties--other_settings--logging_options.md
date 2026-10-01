---
page_title: "other_settings.logging_options"
subcategory: "Load Balancing"
description: "other_settings.logging_options for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1541, "body_sha256": "sha256:d76606f9add0acf961a4825ddc1c17fd47093cfd0bebe6618ba62c4b91a1654b", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:logging_options", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:logging_options:client_log_options", "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:logging_options:origin_log_options"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:logging_options", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--other_settings--logging_options.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["other_settings", "logging_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/other_settings/logging_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "other_settings.logging_options for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# other_settings.logging_options

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [other_settings](data-sources--cdn_loadbalancer--properties--other_settings.md)
- other_settings.logging_options

<a id="section"></a>

Type: `"single"`. Computed.

Defines various OPTIONS related to logging.

Upstream description:

This defines various OPTIONS related to logging.

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

- [client_log_options](data-sources--cdn_loadbalancer--properties--other_settings--logging_options--client_log_options.md): complete subsection reference.

- [origin_log_options](data-sources--cdn_loadbalancer--properties--other_settings--logging_options--origin_log_options.md): complete subsection reference.

## Next pages

- [other_settings.logging_options.client_log_options](data-sources--cdn_loadbalancer--properties--other_settings--logging_options--client_log_options.md)
- [other_settings.logging_options.origin_log_options](data-sources--cdn_loadbalancer--properties--other_settings--logging_options--origin_log_options.md)
- [other_settings](data-sources--cdn_loadbalancer--properties--other_settings.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
