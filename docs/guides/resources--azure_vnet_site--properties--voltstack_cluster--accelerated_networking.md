---
page_title: "voltstack_cluster.accelerated_networking"
subcategory: "Infrastructure"
description: "voltstack_cluster.accelerated_networking for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1898, "body_sha256": "sha256:0d34cf4544eaf59ccf07b5c1781f081ea9ead84cf23c06979fe7a95eb0a3b23b", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:accelerated_networking", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:accelerated_networking:disable_spec", "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:accelerated_networking:enable"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:accelerated_networking", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster", "path": "docs/guides/resources--azure_vnet_site--properties--voltstack_cluster--accelerated_networking.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "accelerated_networking"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster/accelerated_networking/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.accelerated_networking for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster.accelerated_networking

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [voltstack_cluster](resources--azure_vnet_site--properties--voltstack_cluster.md)
- voltstack_cluster.accelerated_networking

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Upstream description:

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-accelerated_networking": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
accelerated_networking {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_spec](resources--azure_vnet_site--properties--voltstack_cluster--accelerated_networking--disable_spec.md): complete subsection reference.

- [enable](resources--azure_vnet_site--properties--voltstack_cluster--accelerated_networking--enable.md): complete subsection reference.

## Next pages

- [voltstack_cluster.accelerated_networking.disable_spec](resources--azure_vnet_site--properties--voltstack_cluster--accelerated_networking--disable_spec.md)
- [voltstack_cluster.accelerated_networking.enable](resources--azure_vnet_site--properties--voltstack_cluster--accelerated_networking--enable.md)
- [voltstack_cluster](resources--azure_vnet_site--properties--voltstack_cluster.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
