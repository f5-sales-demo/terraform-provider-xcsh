---
page_title: "bgp_parameters.local_address"
subcategory: ""
description: "bgp_parameters.local_address for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 929, "body_sha256": "sha256:7733791ac004277d020822ae7c354705084f8da4d404d7b71909af67d0adc741", "canonical_id": "xcsh-docs:resources:bgp:properties:bgp_parameters:local_address", "child_ids": [], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:bgp_parameters:local_address", "parent_id": "xcsh-docs:resources:bgp:properties:bgp_parameters", "path": "docs/guides/resources--bgp--properties--bgp_parameters--local_address.md", "provider_name": "bgp", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bgp_parameters", "local_address"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/bgp_parameters/local_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bgp_parameters.local_address for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bgp_parameters.local_address

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md)
- [Property reference](resources--bgp--reference.md)
- [bgp_parameters](resources--bgp--properties--bgp_parameters.md)
- bgp_parameters.local_address

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
local_address = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [bgp_parameters](resources--bgp--properties--bgp_parameters.md)
- [xcsh_bgp](../resources/bgp.md)
