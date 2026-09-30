---
page_title: "global_network"
subcategory: "Networking"
description: "global_network for xcsh_virtual_network."
xcsh_docs: {"aliases": [], "body_bytes": 1362, "body_sha256": "sha256:72a12cf117bc2ebb63ae57ae2da0778df5be59b58da57f1d9a13303dec095f60", "canonical_id": "xcsh-docs:resources:virtual_network:properties:global_network", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_network:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_network:properties:global_network", "parent_id": "xcsh-docs:resources:virtual_network:reference", "path": "docs/guides/resources--virtual_network--properties--global_network.md", "provider_name": "virtual_network", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["global_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_network/properties/global_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "global_network for xcsh_virtual_network.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# global_network

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md)
- [Property reference](resources--virtual_network--reference.md)
- global_network

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: global\_network, site\_local\_inside\_network, site\_local\_network\] Select the global
virtual-network scope for connectivity across participating sites.

Upstream description:

Select the global virtual-network scope for connectivity across participating sites.

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

- [global_network](resources--virtual_network--properties--global_network.md#section)
- [site_local_inside_network](resources--virtual_network--properties--site_local_inside_network.md#section)
- [site_local_network](resources--virtual_network--properties--site_local_network.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
global_network = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--virtual_network--reference.md)
- [xcsh_virtual_network](../resources/virtual_network.md)
