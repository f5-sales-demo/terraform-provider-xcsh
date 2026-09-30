---
page_title: "ingress_egress_gw.accelerated_networking"
subcategory: "Infrastructure"
description: "ingress_egress_gw.accelerated_networking for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2351, "body_sha256": "sha256:91e9c785c9df7487a261e18aaca660fe8899c4ec72a4b2b55371a459b13f4892", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:accelerated_networking:disable_spec", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:accelerated_networking:enable"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:accelerated_networking", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw", "path": "documentation/resources/azure_vnet_site/properties/ingress_egress_gw/accelerated_networking/index.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["ingress_egress_gw", "accelerated_networking"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw/accelerated_networking/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.accelerated_networking for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw.accelerated_networking

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/)
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

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/accelerated_networking/disable_spec/): complete subsection reference.

- [enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/accelerated_networking/enable/): complete subsection reference.

## Next pages

- [ingress_egress_gw.accelerated_networking.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/accelerated_networking/disable_spec/)
- [ingress_egress_gw.accelerated_networking.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/accelerated_networking/enable/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
