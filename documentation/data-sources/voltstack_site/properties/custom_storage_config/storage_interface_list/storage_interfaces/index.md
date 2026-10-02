---
page_title: "custom_storage_config.storage_interface_list.storage_interfaces"
subcategory: ""
description: "Configure storage interfaces for this App Stack site."
xcsh_docs: {"aliases": ["custom storage config storage interface list storage interfaces"], "body_bytes": 3428, "body_sha256": "sha256:821fe46269bc6b777d9a6112bd8162aaa0dda3e4c160afc6f4b33627b9cfcd8b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:labels", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list", "path": "documentation/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces"], "schema_version": 1, "sections": [{"aliases": ["description spec"], "anchor": "schema-custom_storage_config--storage_interface_list--storage_interfaces--description_spec", "description": "Interface Description. Description for this Interface.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "section", "description": "Add Labels for this Interface, these labels can be used in firewall policy.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:labels", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "labels"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage interface"], "anchor": "section", "description": "Ethernet Interface Configuration.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure storage interfaces for this App Stack site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_interface_list.storage_interfaces

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.storage_interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/)
- custom_storage_config.storage_interface_list.storage_interfaces

<a id="section"></a>

Type: `"list"`. Computed.

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

## Direct properties

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/labels/): complete subsection reference.

- [storage_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/): complete subsection reference.

## Next pages

- [custom_storage_config.storage_interface_list.storage_interfaces.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/labels/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/)
- [custom_storage_config.storage_interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
