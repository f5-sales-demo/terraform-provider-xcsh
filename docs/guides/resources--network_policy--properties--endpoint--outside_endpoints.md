---
page_title: "endpoint.outside_endpoints"
subcategory: "Security"
description: "endpoint.outside_endpoints for xcsh_network_policy."
xcsh_docs: {"aliases": [], "body_bytes": 982, "body_sha256": "sha256:e744015c3d3042eef2449fa8cd2109c08a2eaf0d79b41586b56b02f741298ff0", "canonical_id": "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "child_ids": [], "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "parent_id": "xcsh-docs:resources:network_policy:properties:endpoint", "path": "docs/guides/resources--network_policy--properties--endpoint--outside_endpoints.md", "provider_name": "network_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["endpoint", "outside_endpoints"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/endpoint/outside_endpoints/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "endpoint.outside_endpoints for xcsh_network_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint.outside_endpoints

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md)
- [Property reference](resources--network_policy--reference.md)
- [endpoint](resources--network_policy--properties--endpoint.md)
- endpoint.outside_endpoints

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
outside_endpoints = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [endpoint](resources--network_policy--properties--endpoint.md)
- [xcsh_network_policy](../resources/network_policy.md)
