---
page_title: "client_side_defense"
subcategory: "Load Balancing"
description: "client_side_defense for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1422, "body_sha256": "sha256:713d5058aea1eff9f430bf5021301b41cd06e899b5bf81abeb63a374a0e099e3", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "docs/guides/data-sources--http_loadbalancer--properties--client_side_defense.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["client_side_defense"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/client_side_defense/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "client_side_defense for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# client_side_defense

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- client_side_defense

<a id="section"></a>

Type: `"single"`. Computed.

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

- [client_side_defense](data-sources--http_loadbalancer--properties--client_side_defense.md#section)
- [disable_client_side_defense](data-sources--http_loadbalancer--properties--disable_client_side_defense.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [policy](data-sources--http_loadbalancer--properties--client_side_defense--policy.md): complete subsection reference.

## Next pages

- [client_side_defense.policy](data-sources--http_loadbalancer--properties--client_side_defense--policy.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
