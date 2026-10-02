---
page_title: "storage_static_routes"
subcategory: ""
description: "List of storage static routes."
xcsh_docs: {"aliases": ["storage static routes"], "body_bytes": 1547, "body_sha256": "sha256:3c7d6b5a77ad141e7a997cc15c85836b813ec54b2f43d4e7dd53742cea782210", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_static_routes", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/storage_static_routes/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3310103001011013-0133232333013322-2230330101201311-3021013203230121-1032321030203020-3003121131331312-2101312003321212-2012211010032330", "registry_path": "docs/guides/resources--fleet--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "storage_static_routes:RequiredObjectAttributes:storage_routes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_static_routes"], "schema_version": 1, "sections": [{"aliases": ["storage routes"], "anchor": "section", "description": "List of storage static routes.", "document_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "storage_static_routes.storage_routes:RequiredListObjectAttributes:subnets", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:subnets", "type": "requires"}], "schema_path": ["storage_static_routes", "storage_routes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_static_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of storage static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_static_routes

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- storage_static_routes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for storage static routes.

Upstream description:

List of storage static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_routes")}
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
storage_static_routes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [storage_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/): complete subsection reference.

## Next pages

- [storage_static_routes.storage_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
