---
page_title: "client_side_defense"
subcategory: "Load Balancing"
description: "client_side_defense for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1615, "body_sha256": "sha256:79eb9745c2efc3537de22ae185f92ed5d07c95cf0c5606e0ac2ecc2f4ba837a7", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "docs/guides/resources--cdn_loadbalancer--properties--client_side_defense.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["client_side_defense"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/client_side_defense/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "client_side_defense for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- client_side_defense

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: client\_side\_defense, disable\_client\_side\_defense; Default:
disable\_client\_side\_defense\] Defines various configuration OPTIONS for Client-Side Defense
Policy.

Upstream description:

This defines various configuration OPTIONS for Client-Side Defense Policy.

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

- [client_side_defense](resources--cdn_loadbalancer--properties--client_side_defense.md#section)
- [disable_client_side_defense](resources--cdn_loadbalancer--properties--disable_client_side_defense.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
client_side_defense {
  # Configure direct properties listed below.
}
```

## Direct properties

- [policy](resources--cdn_loadbalancer--properties--client_side_defense--policy.md): complete subsection reference.

## Next pages

- [client_side_defense.policy](resources--cdn_loadbalancer--properties--client_side_defense--policy.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
