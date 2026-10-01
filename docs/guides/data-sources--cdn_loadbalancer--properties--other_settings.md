---
page_title: "other_settings"
subcategory: "Load Balancing"
description: "other_settings for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1771, "body_sha256": "sha256:0572aa15ae121d803440e67d838226aee04159e3d2d36366d17c60e07a25caae", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options", "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:logging_options"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--other_settings.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["other_settings"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/other_settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "other_settings for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# other_settings

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- other_settings

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for other settings.

Upstream description:

Other Settings.

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

<a id="schema-other_settings--add_location"></a>

### add_location property

Type: `"bool"`. Computed.

Add Location. X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt;
in responses.

Upstream description:

X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; in responses.

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

- [header_options](data-sources--cdn_loadbalancer--properties--other_settings--header_options.md): complete subsection reference.

- [logging_options](data-sources--cdn_loadbalancer--properties--other_settings--logging_options.md): complete subsection reference.

## Next pages

- [other_settings.header_options](data-sources--cdn_loadbalancer--properties--other_settings--header_options.md)
- [other_settings.logging_options](data-sources--cdn_loadbalancer--properties--other_settings--logging_options.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
