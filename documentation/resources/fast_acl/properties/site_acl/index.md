---
page_title: "site_acl"
subcategory: ""
description: "Fast ACL definition for Site."
xcsh_docs: {"aliases": ["site acl"], "body_bytes": 3464, "body_sha256": "sha256:16225505d2a9c190b056ab0ed10e0016c1fb03fd4aabeb218d6bcf11cf47d29f", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:fast_acl:properties:site_acl:all_services", "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules", "xcsh-docs:resources:fast_acl:properties:site_acl:inside_network", "xcsh-docs:resources:fast_acl:properties:site_acl:interface_services", "xcsh-docs:resources:fast_acl:properties:site_acl:outside_network", "xcsh-docs:resources:fast_acl:properties:site_acl:vip_services"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:site_acl", "parent_id": "xcsh-docs:resources:fast_acl:reference", "path": "documentation/resources/fast_acl/properties/site_acl/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322", "registry_path": "docs/guides/resources--fast_acl--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "site_acl:ConflictingObjectAttributes:all_services,interface_services", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:all_services", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl:ConflictingObjectAttributes:all_services,vip_services", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:all_services", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl:ConflictingObjectAttributes:inside_network,outside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl:ConflictingObjectAttributes:all_services,interface_services", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:interface_services", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl:ConflictingObjectAttributes:interface_services,vip_services", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:interface_services", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl:ConflictingObjectAttributes:inside_network,outside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:outside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl:ConflictingObjectAttributes:all_services,vip_services", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:vip_services", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl:ConflictingObjectAttributes:interface_services,vip_services", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:vip_services", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["site_acl"], "schema_version": 1, "sections": [{"aliases": ["all services"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:fast_acl:properties:site_acl:all_services", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_acl", "all_services"], "syntax": "attribute", "type": "object"}, {"aliases": ["fast acl rules"], "anchor": "section", "description": "Fast ACL rules to match.", "document_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:prefix", "type": "conflicts"}], "schema_path": ["site_acl", "fast_acl_rules"], "syntax": "block", "type": "object"}, {"aliases": ["inside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:fast_acl:properties:site_acl:inside_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_acl", "inside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["interface services"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:fast_acl:properties:site_acl:interface_services", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_acl", "interface_services"], "syntax": "attribute", "type": "object"}, {"aliases": ["outside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:fast_acl:properties:site_acl:outside_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_acl", "outside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["vip services"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:fast_acl:properties:site_acl:vip_services", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_acl", "vip_services"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/site_acl/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Fast ACL definition for Site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fast_aclCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_acl

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/)
- site_acl

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Fast ACL for Site. Fast ACL definition for Site.

Upstream description:

Fast ACL definition for Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_services",
    "interface_services"),
  validators.ConflictingObjectAttributes("all_services",
    "vip_services"),
  validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("interface_services",
    "vip_services")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\"]",
  "x-ves-oneof-field-vip_choice": "[\"all_services\",\"interface_services\",\"vip_services\"]"
}
```

Terraform syntax:

```terraform
site_acl {
  # Configure direct properties listed below.
}
```

## Direct properties

- [all_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/all_services/): complete subsection reference.

- [fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/): complete subsection reference.

- [inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/inside_network/): complete subsection reference.

- [interface_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/interface_services/): complete subsection reference.

- [outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/outside_network/): complete subsection reference.

- [vip_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/vip_services/): complete subsection reference.

## Next pages

- [site_acl.all_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/all_services/)
- [site_acl.fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/)
- [site_acl.inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/inside_network/)
- [site_acl.interface_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/interface_services/)
- [site_acl.outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/outside_network/)
- [site_acl.vip_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/vip_services/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/)
- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
