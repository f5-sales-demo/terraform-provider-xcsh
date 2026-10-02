---
page_title: "custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list"
subcategory: ""
description: "On a multinode site, this list holds the nodes and corresponding networking_interface."
xcsh_docs: {"aliases": ["custom network config slo config static v6 routes static routes node interface list"], "body_bytes": 3963, "body_sha256": "sha256:c3f0c60b4b5b23927859a14360a9b3fe85aafd5037e6567802b66409693a6b06", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_v6_routes:static_routes:node_interface:list:interface"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_v6_routes:static_routes:node_interface:list", "parent_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_v6_routes:static_routes:node_interface", "path": "documentation/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/list/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2300111200221200-1230013330320233-3103211131111130-1123021202210332-2123233013312233-3122000212032302-3003110032103002-2322202232033313", "registry_path": "docs/guides/resources--securemesh_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "slo_config", "static_v6_routes", "static_routes", "node_interface", "list"], "schema_version": 1, "sections": [{"aliases": ["interface"], "anchor": "section", "description": "Interface reference on this node.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_v6_routes:static_routes:node_interface:list:interface", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["custom_network_config", "slo_config", "static_v6_routes", "static_routes", "node_interface", "list", "interface"], "syntax": "block", "type": "object"}, {"aliases": ["node"], "anchor": "schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--node", "description": "Node name on this site.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_v6_routes:static_routes:node_interface:list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "slo_config", "static_v6_routes", "static_routes", "node_interface", "list", "node"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "On a multinode site, this list holds the nodes and corresponding networking_interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/)
- [custom_network_config.slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/)
- [custom_network_config.slo_config.static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/)
- [custom_network_config.slo_config.static_v6_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/list/interface/): complete subsection reference.

<a id="schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--node"></a>

### node property

Type: `"string"`. Optional.

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

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/list/interface/)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
