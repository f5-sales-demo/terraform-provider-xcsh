---
page_title: "local_vrf.sli_config.static_routes.static_routes.node_interface.list"
subcategory: ""
description: "On a multinode site, this list holds the nodes and corresponding networking_interface."
xcsh_docs: {"aliases": ["local vrf sli config static routes static routes node interface list"], "body_bytes": 3075, "body_sha256": "sha256:471f3d322d4214c93fc8e536b0d342a793266096a800998dbedac3aa40281483", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config:static_routes:static_routes:node_interface:list:interface"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config:static_routes:static_routes:node_interface:list", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config:static_routes:static_routes:node_interface", "path": "documentation/resources/securemesh_site_v2/properties/local_vrf/sli_config/static_routes/static_routes/node_interface/list/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0102023020212301-3210010200212333-0320310313011103-2102112123100331-2301120110233112-3210011311112122-0032033303133123-0303103323221123", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_vrf", "sli_config", "static_routes", "static_routes", "node_interface", "list"], "schema_version": 1, "sections": [{"aliases": ["local vrf sli config static routes static routes node interface list interface"], "anchor": "section", "description": "Interface reference on this node.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config:static_routes:static_routes:node_interface:list:interface", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["local_vrf", "sli_config", "static_routes", "static_routes", "node_interface", "list", "interface"], "syntax": "block", "type": "object"}, {"aliases": ["local vrf sli config static routes static routes node interface list node"], "anchor": "schema-local_vrf--sli_config--static_routes--static_routes--node_interface--list--node", "description": "Node name on this site.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:sli_config:static_routes:static_routes:node_interface:list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "sli_config", "static_routes", "static_routes", "node_interface", "list", "node"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/local_vrf/sli_config/static_routes/static_routes/node_interface/list/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "On a multinode site, this list holds the nodes and corresponding networking_interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_vrf.sli_config.static_routes.static_routes.node_interface.list

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [local_vrf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/)
- [local_vrf.sli_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/sli_config/)
- [local_vrf.sli_config.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/sli_config/static_routes/)
- [local_vrf.sli_config.static_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/sli_config/static_routes/static_routes/)
- [local_vrf.sli_config.static_routes.static_routes.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/sli_config/static_routes/static_routes/node_interface/)
- local_vrf.sli_config.static_routes.static_routes.node_interface.list

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
      "validatedAt": "2026-10-21T11:00:09+00:00"
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

- [interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/sli_config/static_routes/static_routes/node_interface/list/interface/): complete subsection reference.

<a id="schema-local_vrf--sli_config--static_routes--static_routes--node_interface--list--node"></a>

### node property

Type: `"string"`. Optional.

Node. Node name on this site.

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
      "validatedAt": "2026-10-21T11:00:09+00:00"
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
