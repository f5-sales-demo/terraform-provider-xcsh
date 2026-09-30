---
page_title: "other_settings.logging_options"
subcategory: "Load Balancing"
description: "other_settings.logging_options for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1539, "body_sha256": "sha256:4e8862b98818ad7681068803f5ac265470a8faf90d82fdfb4143ea64712223a2", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings:logging_options", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:other_settings:logging_options:client_log_options", "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings:logging_options:origin_log_options"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings:logging_options", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings", "path": "docs/guides/resources--cdn_loadbalancer--properties--other_settings--logging_options.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["other_settings", "logging_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/other_settings/logging_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "other_settings.logging_options for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# other_settings.logging_options

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [other_settings](resources--cdn_loadbalancer--properties--other_settings.md)
- other_settings.logging_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
logging_options {
  # Configure direct properties listed below.
}
```

## Direct properties

- [client_log_options](resources--cdn_loadbalancer--properties--other_settings--logging_options--client_log_options.md): complete subsection reference.

- [origin_log_options](resources--cdn_loadbalancer--properties--other_settings--logging_options--origin_log_options.md): complete subsection reference.

## Next pages

- [other_settings.logging_options.client_log_options](resources--cdn_loadbalancer--properties--other_settings--logging_options--client_log_options.md)
- [other_settings.logging_options.origin_log_options](resources--cdn_loadbalancer--properties--other_settings--logging_options--origin_log_options.md)
- [other_settings](resources--cdn_loadbalancer--properties--other_settings.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
