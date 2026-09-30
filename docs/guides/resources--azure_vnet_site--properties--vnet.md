---
page_title: "vnet"
subcategory: "Infrastructure"
description: "vnet for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1391, "body_sha256": "sha256:bbe787693dd974fab1861d7d83992d6660214ea16d848df2f783ff0ac51bff24", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:vnet", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:vnet:existing_vnet", "xcsh-docs:resources:azure_vnet_site:properties:vnet:new_vnet"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:vnet", "parent_id": "xcsh-docs:resources:azure_vnet_site:reference", "path": "docs/guides/resources--azure_vnet_site--properties--vnet.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/vnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vnet for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# vnet

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
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

- [existing_vnet](resources--azure_vnet_site--properties--vnet--existing_vnet.md): complete subsection reference.

- [new_vnet](resources--azure_vnet_site--properties--vnet--new_vnet.md): complete subsection reference.

## Next pages

- [vnet.existing_vnet](resources--azure_vnet_site--properties--vnet--existing_vnet.md)
- [vnet.new_vnet](resources--azure_vnet_site--properties--vnet--new_vnet.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
