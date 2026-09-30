---
page_title: "enable_challenge.default_mitigation_settings"
subcategory: "Load Balancing"
description: "enable_challenge.default_mitigation_settings for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 982, "body_sha256": "sha256:e581a334757fa1e1d19ee86f9c673306f31f216273642999ea7c46ccb60b8d1d", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:default_mitigation_settings", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge:default_mitigation_settings", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_challenge", "path": "docs/guides/resources--http_loadbalancer--properties--enable_challenge--default_mitigation_settings.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_challenge", "default_mitigation_settings"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/enable_challenge/default_mitigation_settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_challenge.default_mitigation_settings for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# enable_challenge.default_mitigation_settings

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [enable_challenge](resources--http_loadbalancer--properties--enable_challenge.md)
- enable_challenge.default_mitigation_settings

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
default_mitigation_settings = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [enable_challenge](resources--http_loadbalancer--properties--enable_challenge.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
