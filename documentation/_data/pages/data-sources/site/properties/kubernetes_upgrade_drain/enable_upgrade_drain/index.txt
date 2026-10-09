---
page_title: "kubernetes_upgrade_drain.enable_upgrade_drain"
subcategory: "Infrastructure"
description: "Specify batch upgrade settings for worker nodes within a site."
xcsh_docs: {"aliases": ["kubernetes upgrade drain enable upgrade drain"], "body_bytes": 2052, "body_sha256": "sha256:6eb4ebf30b3f7344813b833ad3334e742e06a623064ab5467a2081aa4dbd2584", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:enable_upgrade_drain:disable_vega_upgrade_mode", "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:enable_upgrade_drain:enable_vega_upgrade_mode"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "parent_id": "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain", "path": "documentation/data-sources/site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1003003101232122-3001012123101232-1013232123113111-3230222011212013-1232232123330001-2322331222000220-3011032010313001-2121132021102313", "registry_path": "docs/guides/data-sources--site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["kubernetes_upgrade_drain", "enable_upgrade_drain"], "schema_version": 1, "sections": [{"aliases": ["kubernetes upgrade drain enable upgrade drain disable vega upgrade mode"], "anchor": "section", "description": "Configuration parameter for disable vega upgrade mode.", "document_id": "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:enable_upgrade_drain:disable_vega_upgrade_mode", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kubernetes_upgrade_drain", "enable_upgrade_drain", "disable_vega_upgrade_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["kubernetes upgrade drain enable upgrade drain drain max unavailable node count"], "anchor": "schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_count", "description": "Node Batch Size Count. Exclusive with", "document_id": "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kubernetes_upgrade_drain", "enable_upgrade_drain", "drain_max_unavailable_node_count"], "syntax": "attribute", "type": "number"}, {"aliases": ["kubernetes upgrade drain enable upgrade drain drain max unavailable node percentage"], "anchor": "schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_percentage", "description": "Maximum percentage of nodes unavailable during upgrade draining.", "document_id": "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kubernetes_upgrade_drain", "enable_upgrade_drain", "drain_max_unavailable_node_percentage"], "syntax": "attribute", "type": "number"}, {"aliases": ["duration", "kubernetes upgrade drain enable upgrade drain drain node timeout"], "anchor": "schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_node_timeout", "description": "Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is..", "document_id": "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kubernetes_upgrade_drain", "enable_upgrade_drain", "drain_node_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["kubernetes upgrade drain enable upgrade drain enable vega upgrade mode"], "anchor": "section", "description": "Configuration parameter for enable vega upgrade mode.", "document_id": "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:enable_upgrade_drain:enable_vega_upgrade_mode", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kubernetes_upgrade_drain", "enable_upgrade_drain", "enable_vega_upgrade_mode"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Specify batch upgrade settings for worker nodes within a site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kubernetes_upgrade_drain.enable_upgrade_drain

Breadcrumbs:

- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/)
- [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/kubernetes_upgrade_drain/)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="section"></a>

Type: `"single"`. Computed.

Specify batch upgrade settings for worker nodes within a site.

## Direct properties

- [disable_vega_upgrade_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/disable_vega_upgrade_mode/): complete subsection reference.

<a id="schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_count"></a>

### drain_max_unavailable_node_count property

Type: `"number"`. Computed.

Node Batch Size Count. Exclusive with \[\]

<a id="schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_percentage"></a>

### drain_max_unavailable_node_percentage property

Type: `"number"`. Computed.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_node_timeout"></a>

### drain_node_timeout property

Type: `"number"`. Computed.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It
is..

- [enable_vega_upgrade_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/enable_vega_upgrade_mode/): complete subsection reference.
