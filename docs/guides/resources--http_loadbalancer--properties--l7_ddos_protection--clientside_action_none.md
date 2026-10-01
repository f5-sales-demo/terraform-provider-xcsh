---
page_title: "l7_ddos_protection.clientside_action_none"
subcategory: "Load Balancing"
description: "l7_ddos_protection.clientside_action_none for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1078, "body_sha256": "sha256:8e4fc0f6d2b9e74f64abbf49573ed09e3c66034f13e6aca17bafebddc4093519", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_none", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:clientside_action_none", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection", "path": "docs/guides/resources--http_loadbalancer--properties--l7_ddos_protection--clientside_action_none.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["l7_ddos_protection", "clientside_action_none"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/l7_ddos_protection/clientside_action_none/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "l7_ddos_protection.clientside_action_none for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# l7_ddos_protection.clientside_action_none

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [l7_ddos_protection](resources--http_loadbalancer--properties--l7_ddos_protection.md)
- l7_ddos_protection.clientside_action_none

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
clientside_action_none = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [l7_ddos_protection](resources--http_loadbalancer--properties--l7_ddos_protection.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
