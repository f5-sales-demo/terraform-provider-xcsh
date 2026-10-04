---
page_title: "custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list"
subcategory: ""
description: "On a multinode site, this list holds the nodes and corresponding networking_interface."
xcsh_docs: {"aliases": ["custom network config sli config static v6 routes static routes node interface list"], "body_bytes": 3963, "body_sha256": "sha256:26077f745f7d629d64638388e730eb76d35228189d9d0681743c029e578d60c1", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:static_v6_routes:static_routes:node_interface:list:interface"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:static_v6_routes:static_routes:node_interface:list", "parent_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:static_v6_routes:static_routes:node_interface", "path": "documentation/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/node_interface/list/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1103201302332003-0030000022000123-1310321232012123-1200020332232312-2332013023203222-2311003111033301-1303312210122320-1310021112133230", "registry_path": "docs/guides/resources--securemesh_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "sli_config", "static_v6_routes", "static_routes", "node_interface", "list"], "schema_version": 1, "sections": [{"aliases": ["custom network config sli config static v6 routes static routes node interface list interface"], "anchor": "section", "description": "Interface reference on this node.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:static_v6_routes:static_routes:node_interface:list:interface", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["custom_network_config", "sli_config", "static_v6_routes", "static_routes", "node_interface", "list", "interface"], "syntax": "block", "type": "object"}, {"aliases": ["custom network config sli config static v6 routes static routes node interface list node"], "anchor": "schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--node", "description": "Node name on this site.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config:static_v6_routes:static_routes:node_interface:list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "sli_config", "static_v6_routes", "static_routes", "node_interface", "list", "node"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/node_interface/list/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "On a multinode site, this list holds the nodes and corresponding networking_interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/)
- [custom_network_config.sli_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/)
- [custom_network_config.sli_config.static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/)
- [custom_network_config.sli_config.static_v6_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/node_interface/)
- custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/node_interface/list/interface/): complete subsection reference.

<a id="schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--node"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/node_interface/list/interface/)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/node_interface/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
