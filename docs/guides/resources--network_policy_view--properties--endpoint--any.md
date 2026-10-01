---
page_title: "endpoint.any"
subcategory: ""
description: "endpoint.any for xcsh_network_policy_view."
xcsh_docs: {"aliases": [], "body_bytes": 975, "body_sha256": "sha256:cf7e710102e3d45e061f25670300d633d27947fe00a1b3c824cfd05a6041fb5e", "canonical_id": "xcsh-docs:resources:network_policy_view:properties:endpoint:any", "child_ids": [], "collection_id": "xcsh-docs:resources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy_view:properties:endpoint:any", "parent_id": "xcsh-docs:resources:network_policy_view:properties:endpoint", "path": "docs/guides/resources--network_policy_view--properties--endpoint--any.md", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["endpoint", "any"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_view/properties/endpoint/any/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "endpoint.any for xcsh_network_policy_view.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint.any

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md)
- [Property reference](resources--network_policy_view--reference.md)
- [endpoint](resources--network_policy_view--properties--endpoint.md)
- endpoint.any

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
any = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [endpoint](resources--network_policy_view--properties--endpoint.md)
- [xcsh_network_policy_view](../resources/network_policy_view.md)
