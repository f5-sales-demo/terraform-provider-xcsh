---
page_title: "other_settings"
subcategory: "Load Balancing"
description: "other_settings for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1672, "body_sha256": "sha256:40b29279acf6f585660f0d7bab03fdac3f6d2366cda48d8ba92127db90e0c1d2", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options", "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:logging_options"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--other_settings.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["other_settings"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/other_settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "other_settings for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
