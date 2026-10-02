---
page_title: "sriov_interfaces"
subcategory: ""
description: "List of all custom SR-IOV interfaces configuration."
xcsh_docs: {"aliases": ["sriov interfaces"], "body_bytes": 1386, "body_sha256": "sha256:14d09d09fc807891872cd78018b81d9253db19ad99de17eea8a0d66060410ef3", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:sriov_interfaces:sriov_interface"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:sriov_interfaces", "parent_id": "xcsh-docs:resources:voltstack_site:reference", "path": "documentation/resources/voltstack_site/properties/sriov_interfaces/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3202202012201132-3231030210000322-0123032022001103-3031211120312222-3332121322033010-1220331122103013-1123232303130120-1200100021331233", "registry_path": "docs/guides/resources--voltstack_site--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sriov_interfaces"], "schema_version": 1, "sections": [{"aliases": ["sriov interface"], "anchor": "section", "description": "Use custom SR-IOV interfaces Configuration.", "document_id": "xcsh-docs:resources:voltstack_site:properties:sriov_interfaces:sriov_interface", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-sriov_interfaces--sriov_interface--interface_name", "enforcement": "provider-schema", "group": "sriov_interfaces.sriov_interface:RequiredListObjectAttributes:interface_name,number_of_vfs", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:sriov_interfaces:sriov_interface", "type": "requires"}, {"anchor": "schema-sriov_interfaces--sriov_interface--number_of_vfs", "enforcement": "provider-schema", "group": "sriov_interfaces.sriov_interface:RequiredListObjectAttributes:interface_name,number_of_vfs", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:sriov_interfaces:sriov_interface", "type": "requires"}], "schema_path": ["sriov_interfaces", "sriov_interface"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/sriov_interfaces/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of all custom SR-IOV interfaces configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sriov_interfaces

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- sriov_interfaces

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of all custom SR-IOV interfaces configuration.

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
sriov_interfaces {
  # Configure direct properties listed below.
}
```

## Direct properties

- [sriov_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/sriov_interfaces/sriov_interface/): complete subsection reference.

## Next pages

- [sriov_interfaces.sriov_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/sriov_interfaces/sriov_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
