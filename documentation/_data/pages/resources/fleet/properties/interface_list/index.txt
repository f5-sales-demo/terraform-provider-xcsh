---
page_title: "interface_list"
subcategory: ""
description: "Add all interfaces belonging to this fleet."
xcsh_docs: {"aliases": ["interface list"], "body_bytes": 1423, "body_sha256": "sha256:fd4f5c9cff9176525914256bababdc4e689f1f3749b24330bb97d682bcad51b9", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:interface_list:interfaces"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:interface_list", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/interface_list/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2033313231120123-0013102010311223-1030231302132002-1113111101332321-3323223200011101-2001323010002331-0223230320312012-1023032001121302", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "interface_list:RequiredObjectAttributes:interfaces", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:interface_list:interfaces", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["interface_list"], "schema_version": 1, "sections": [{"aliases": ["interfaces"], "anchor": "section", "description": "Add all interfaces belonging to this fleet.", "document_id": "xcsh-docs:resources:fleet:properties:interface_list:interfaces", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-interface_list--interfaces--name", "enforcement": "provider-schema", "group": "interface_list.interfaces:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:interface_list:interfaces", "type": "requires"}], "schema_path": ["interface_list", "interfaces"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/interface_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Add all interfaces belonging to this fleet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# interface_list

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- interface_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Add all interfaces belonging to this fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interfaces")}
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
interface_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/interface_list/interfaces/): complete subsection reference.

## Next pages

- [interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/interface_list/interfaces/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
