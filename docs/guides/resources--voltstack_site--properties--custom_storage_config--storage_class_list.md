---
page_title: "custom_storage_config.storage_class_list"
subcategory: ""
description: "custom_storage_config.storage_class_list for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1237, "body_sha256": "sha256:a35be96b1f1c589b3656a34e86f820b315024dfbfdddf217989c91c0982f8b1a", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config", "path": "docs/guides/resources--voltstack_site--properties--custom_storage_config--storage_class_list.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "storage_class_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_class_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_class_list for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_storage_config.storage_class_list

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_storage_config](resources--voltstack_site--properties--custom_storage_config.md)
- custom_storage_config.storage_class_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Add additional custom storage classes in Kubernetes for this fleet.

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
storage_class_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [storage_classes](resources--voltstack_site--properties--custom_storage_config--storage_class_list--storage_classes.md): complete subsection reference.

## Next pages

- [custom_storage_config.storage_class_list.storage_classes](resources--voltstack_site--properties--custom_storage_config--storage_class_list--storage_classes.md)
- [custom_storage_config](resources--voltstack_site--properties--custom_storage_config.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
