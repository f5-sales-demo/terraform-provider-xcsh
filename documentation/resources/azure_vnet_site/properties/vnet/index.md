---
page_title: "vnet"
subcategory: "Infrastructure"
description: "vnet for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1799, "body_sha256": "sha256:0f2bd1823be4f252d062773492a0f2875b0e677dff5d9bd2858ef3607c7b1e99", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:vnet:existing_vnet", "xcsh-docs:resources:azure_vnet_site:properties:vnet:new_vnet"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:vnet", "parent_id": "xcsh-docs:resources:azure_vnet_site:reference", "path": "documentation/resources/azure_vnet_site/properties/vnet/index.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["vnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/vnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vnet for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# vnet

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- vnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about Azure VNet for a view.

Upstream description:

This defines choice about Azure VNet for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_vnet",
    "new_vnet")}
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
  "x-ves-oneof-field-choice": "[\"existing_vnet\",\"new_vnet\"]"
}
```

Terraform syntax:

```terraform
vnet {
  # Configure direct properties listed below.
}
```

## Direct properties

- [existing_vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/vnet/existing_vnet/): complete subsection reference.

- [new_vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/vnet/new_vnet/): complete subsection reference.

## Next pages

- [vnet.existing_vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/vnet/existing_vnet/)
- [vnet.new_vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/vnet/new_vnet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
