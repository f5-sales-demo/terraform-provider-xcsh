---
page_title: "custom_storage_config.storage_interface_list.storage_interfaces"
subcategory: ""
description: "custom_storage_config.storage_interface_list.storage_interfaces for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 3564, "body_sha256": "sha256:e34fd0bba26cccf50c76733d65057df345ee57c669753c144704b81f24413c76", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:labels", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list", "path": "documentation/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/index.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_interface_list.storage_interfaces for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_storage_config.storage_interface_list.storage_interfaces

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.storage_interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/)
- custom_storage_config.storage_interface_list.storage_interfaces

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Configure storage interfaces for this App Stack site.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage_interfaces {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/labels/): complete subsection reference.

- [storage_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/): complete subsection reference.

## Next pages

- [custom_storage_config.storage_interface_list.storage_interfaces.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/labels/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/)
- [custom_storage_config.storage_interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
