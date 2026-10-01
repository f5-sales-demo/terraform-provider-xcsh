---
page_title: "site_local_inside_network"
subcategory: "Networking"
description: "site_local_inside_network for xcsh_virtual_network."
xcsh_docs: {"aliases": [], "body_bytes": 896, "body_sha256": "sha256:7af587019c2ffd2decc0be8fb93b12044c5147beea0563256ee36abb9e06dabe", "canonical_id": "xcsh-docs:resources:virtual_network:properties:site_local_inside_network", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_network:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_network:properties:site_local_inside_network", "parent_id": "xcsh-docs:resources:virtual_network:reference", "path": "docs/guides/resources--virtual_network--properties--site_local_inside_network.md", "provider_name": "virtual_network", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["site_local_inside_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_network/properties/site_local_inside_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_local_inside_network for xcsh_virtual_network.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_local_inside_network

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md)
- [Property reference](resources--virtual_network--reference.md)
- site_local_inside_network

<a id="section"></a>

Type: `["object", {}]`. Optional.

Select the site-local inside network for site-internal connectivity.

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
site_local_inside_network = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--virtual_network--reference.md)
- [xcsh_virtual_network](../resources/virtual_network.md)
