---
page_title: "user_defined_api_discovery_policy.inclusive"
subcategory: ""
description: "user_defined_api_discovery_policy.inclusive for xcsh_api_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1174, "body_sha256": "sha256:8287cfaf4b5f7de1f2680da5eb54722c0b5694a03920cefc1cc09b3c013b57bb", "canonical_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:inclusive", "child_ids": [], "collection_id": "xcsh-docs:resources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:inclusive", "parent_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy", "path": "docs/guides/resources--api_discovery--properties--user_defined_api_discovery_policy--inclusive.md", "provider_name": "api_discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["user_defined_api_discovery_policy", "inclusive"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_discovery/properties/user_defined_api_discovery_policy/inclusive/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "user_defined_api_discovery_policy.inclusive for xcsh_api_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_defined_api_discovery_policy.inclusive

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md)
- [Property reference](resources--api_discovery--reference.md)
- [user_defined_api_discovery_policy](resources--api_discovery--properties--user_defined_api_discovery_policy.md)
- user_defined_api_discovery_policy.inclusive

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
inclusive = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [user_defined_api_discovery_policy](resources--api_discovery--properties--user_defined_api_discovery_policy.md)
- [xcsh_api_discovery](../resources/api_discovery.md)
