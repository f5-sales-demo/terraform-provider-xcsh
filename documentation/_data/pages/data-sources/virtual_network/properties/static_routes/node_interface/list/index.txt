---
page_title: "static_routes.node_interface.list"
subcategory: "Networking"
description: "On a multinode site, this list holds the nodes and corresponding networking_interface."
xcsh_docs: {"aliases": ["static routes node interface list"], "body_bytes": 2764, "body_sha256": "sha256:efdaf5a8b4466a2f589d16347236027d6339ab8edb2732f1b340744844da35b4", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:virtual_network:properties:static_routes:node_interface:list:interface"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:virtual_network:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_network:properties:static_routes:node_interface:list", "parent_id": "xcsh-docs:data-sources:virtual_network:properties:static_routes:node_interface", "path": "documentation/data-sources/virtual_network/properties/static_routes/node_interface/list/index.md", "product": "distributed-cloud", "provider_name": "virtual_network", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1211120123103031-0301310100203102-2332232103303120-0130100220101311-0010030033211213-2011102230130022-1313232101210022-2321222333123022", "registry_path": "docs/guides/data-sources--virtual_network--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["static_routes", "node_interface", "list"], "schema_version": 1, "sections": [{"aliases": ["interface"], "anchor": "section", "description": "Interface reference on this node.", "document_id": "xcsh-docs:data-sources:virtual_network:properties:static_routes:node_interface:list:interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["static_routes", "node_interface", "list", "interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["node"], "anchor": "schema-static_routes--node_interface--list--node", "description": "Node name on this site.", "document_id": "xcsh-docs:data-sources:virtual_network:properties:static_routes:node_interface:list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["static_routes", "node_interface", "list", "node"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_network/properties/static_routes/node_interface/list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "On a multinode site, this list holds the nodes and corresponding networking_interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# static_routes.node_interface.list

Breadcrumbs:

- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/properties/)
- [static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/properties/static_routes/)
- [static_routes.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/properties/static_routes/node_interface/)
- static_routes.node_interface.list

<a id="section"></a>

Type: `"list"`. Computed.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

## Direct properties

- [interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/properties/static_routes/node_interface/list/interface/): complete subsection reference.

<a id="schema-static_routes--node_interface--list--node"></a>

### node property

Type: `"string"`. Computed.

Node. Node name on this site.

Upstream description:

Node name on this site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [static_routes.node_interface.list.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/properties/static_routes/node_interface/list/interface/)
- [static_routes.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/properties/static_routes/node_interface/)
- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/)
