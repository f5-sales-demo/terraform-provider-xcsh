---
page_title: "custom_storage_config.storage_interface_list"
subcategory: ""
description: "Configure storage interfaces for this App Stack site."
xcsh_docs: {"aliases": ["custom storage config storage interface list"], "body_bytes": 1870, "body_sha256": "sha256:02f8784e83aec0b2397650b9731b9c5a7508d31d419de7e3246e5feb0c2b53cf", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config", "path": "documentation/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333", "registry_path": "docs/guides/resources--voltstack_site--reference--group-008.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list:RequiredObjectAttributes:storage_interfaces", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_interface_list"], "schema_version": 1, "sections": [{"aliases": ["custom storage config storage interface list storage interfaces"], "anchor": "section", "description": "Configure storage interfaces for this App Stack site.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure storage interfaces for this App Stack site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_interface_list

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/)
- custom_storage_config.storage_interface_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure storage interfaces for this App Stack site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_interfaces")}
```

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
storage_interface_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [storage_interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/): complete subsection reference.

## Next pages

- [custom_storage_config.storage_interface_list.storage_interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
