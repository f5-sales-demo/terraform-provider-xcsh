---
page_title: "default_sensitive_data_policy"
subcategory: "Load Balancing"
description: "default_sensitive_data_policy for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1303, "body_sha256": "sha256:1c0d0747590454e5a1db8cd61b17cd952ca28fbf2c56765bf53a3f58f01dd3f0", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:default_sensitive_data_policy", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:default_sensitive_data_policy", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "docs/guides/resources--cdn_loadbalancer--properties--default_sensitive_data_policy.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_sensitive_data_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/default_sensitive_data_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_sensitive_data_policy for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# default_sensitive_data_policy

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- default_sensitive_data_policy

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_sensitive\_data\_policy, sensitive\_data\_policy; Default:
default\_sensitive\_data\_policy\] Policy configuration for this feature.

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

- [default_sensitive_data_policy](resources--cdn_loadbalancer--properties--default_sensitive_data_policy.md#section)
- [sensitive_data_policy](resources--cdn_loadbalancer--properties--sensitive_data_policy.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_sensitive_data_policy = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
