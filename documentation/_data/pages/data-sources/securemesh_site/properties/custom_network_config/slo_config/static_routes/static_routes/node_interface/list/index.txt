---
page_title: "custom_network_config.slo_config.static_routes.static_routes.node_interface.list"
subcategory: ""
description: "On a multinode site, this list holds the nodes and corresponding networking_interface."
xcsh_docs: {"aliases": ["custom network config slo config static routes static routes node interface list"], "body_bytes": 3841, "body_sha256": "sha256:8ed09bfe89e4ab15d7430151dfe81e6682f1f17e2e9b44d91ae188bacaa668df", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes:node_interface:list:interface"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes:node_interface:list", "parent_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes:node_interface", "path": "documentation/data-sources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/node_interface/list/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1133032110213012-1133100113111131-2312223322331200-1232010003331131-3132022323200011-1301011220003101-2010302130123313-2103111101020133", "registry_path": "docs/guides/data-sources--securemesh_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "slo_config", "static_routes", "static_routes", "node_interface", "list"], "schema_version": 1, "sections": [{"aliases": ["interface"], "anchor": "section", "description": "Interface reference on this node.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes:node_interface:list:interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["custom_network_config", "slo_config", "static_routes", "static_routes", "node_interface", "list", "interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["node"], "anchor": "schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--node", "description": "Node name on this site.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes:node_interface:list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "slo_config", "static_routes", "static_routes", "node_interface", "list", "node"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/node_interface/list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "On a multinode site, this list holds the nodes and corresponding networking_interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.slo_config.static_routes.static_routes.node_interface.list

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/)
- [custom_network_config.slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/)
- [custom_network_config.slo_config.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/static_routes/)
- [custom_network_config.slo_config.static_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/node_interface/)
- custom_network_config.slo_config.static_routes.static_routes.node_interface.list

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

- [interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/node_interface/list/interface/): complete subsection reference.

<a id="schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--node"></a>

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

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/node_interface/list/interface/)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/node_interface/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
