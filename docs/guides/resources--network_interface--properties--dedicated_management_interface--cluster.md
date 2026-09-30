---
page_title: "dedicated_management_interface.cluster"
subcategory: ""
description: "dedicated_management_interface.cluster for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 1006, "body_sha256": "sha256:8d8c206036b41bd28154b4670c4e2caad8d42fc12ba16eb9f6187ecdf52770fd", "canonical_id": "xcsh-docs:resources:network_interface:properties:dedicated_management_interface:cluster", "child_ids": [], "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:dedicated_management_interface:cluster", "parent_id": "xcsh-docs:resources:network_interface:properties:dedicated_management_interface", "path": "docs/guides/resources--network_interface--properties--dedicated_management_interface--cluster.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dedicated_management_interface", "cluster"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/dedicated_management_interface/cluster/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dedicated_management_interface.cluster for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# dedicated_management_interface.cluster

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md)
- [Property reference](resources--network_interface--reference.md)
- [dedicated_management_interface](resources--network_interface--properties--dedicated_management_interface.md)
- dedicated_management_interface.cluster

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
cluster = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [dedicated_management_interface](resources--network_interface--properties--dedicated_management_interface.md)
- [xcsh_network_interface](../resources/network_interface.md)
