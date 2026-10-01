---
page_title: "ingress_egress_gw.accelerated_networking"
subcategory: "Infrastructure"
description: "ingress_egress_gw.accelerated_networking for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1997, "body_sha256": "sha256:ecd597438e7d646800335b7efb69fd566e14bc73921037dcec683a5bd8245307", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:accelerated_networking", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:accelerated_networking:disable_spec", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:accelerated_networking:enable"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:accelerated_networking", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_egress_gw--accelerated_networking.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "accelerated_networking"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw/accelerated_networking/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.accelerated_networking for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.accelerated_networking

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_egress_gw](resources--azure_vnet_site--properties--ingress_egress_gw.md)
- ingress_egress_gw.accelerated_networking

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

- [disable_spec](resources--azure_vnet_site--properties--ingress_egress_gw--accelerated_networking--disable_spec.md): complete subsection reference.

- [enable](resources--azure_vnet_site--properties--ingress_egress_gw--accelerated_networking--enable.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.accelerated_networking.disable_spec](resources--azure_vnet_site--properties--ingress_egress_gw--accelerated_networking--disable_spec.md)
- [ingress_egress_gw.accelerated_networking.enable](resources--azure_vnet_site--properties--ingress_egress_gw--accelerated_networking--enable.md)
- [ingress_egress_gw](resources--azure_vnet_site--properties--ingress_egress_gw.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
