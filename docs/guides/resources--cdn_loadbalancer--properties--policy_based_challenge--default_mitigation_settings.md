---
page_title: "policy_based_challenge.default_mitigation_settings"
subcategory: "Load Balancing"
description: "policy_based_challenge.default_mitigation_settings for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1011, "body_sha256": "sha256:21b5d245972248d894eb52c092fdd0c1b8054a2ac91cb37b6a463cd23a12cd10", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:default_mitigation_settings", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:default_mitigation_settings", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge", "path": "docs/guides/resources--cdn_loadbalancer--properties--policy_based_challenge--default_mitigation_settings.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_based_challenge", "default_mitigation_settings"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/policy_based_challenge/default_mitigation_settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge.default_mitigation_settings for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# policy_based_challenge.default_mitigation_settings

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [policy_based_challenge](resources--cdn_loadbalancer--properties--policy_based_challenge.md)
- policy_based_challenge.default_mitigation_settings

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

- [policy_based_challenge](resources--cdn_loadbalancer--properties--policy_based_challenge.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
