---
page_title: "kubernetes_upgrade_drain.enable_upgrade_drain"
subcategory: "Infrastructure"
description: "Specify batch upgrade settings for worker nodes within a site."
xcsh_docs: {"aliases": ["kubernetes upgrade drain enable upgrade drain"], "body_bytes": 2767, "body_sha256": "sha256:d15693817b26c772c09a20a5d181592cac696e865cd4018169dea531b6ce38e0", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:enable_upgrade_drain:disable_vega_upgrade_mode", "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:enable_upgrade_drain:enable_vega_upgrade_mode"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "parent_id": "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain", "path": "documentation/data-sources/site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1003003101232122-3001012123101232-1013232123113111-3230222011212013-1232232123330001-2322331222000220-3011032010313001-2121132021102313", "registry_path": "docs/guides/data-sources--site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["kubernetes_upgrade_drain", "enable_upgrade_drain"], "schema_version": 1, "sections": [{"aliases": ["disable vega upgrade mode"], "anchor": "section", "description": "Configuration parameter for disable vega upgrade mode.", "document_id": "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:enable_upgrade_drain:disable_vega_upgrade_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kubernetes_upgrade_drain", "enable_upgrade_drain", "disable_vega_upgrade_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["drain max unavailable node count"], "anchor": "schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_count", "description": "Node Batch Size Count. Exclusive with", "document_id": "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kubernetes_upgrade_drain", "enable_upgrade_drain", "drain_max_unavailable_node_count"], "syntax": "attribute", "type": "number"}, {"aliases": ["drain max unavailable node percentage"], "anchor": "schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_percentage", "description": "Maximum percentage of nodes unavailable during upgrade draining.", "document_id": "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kubernetes_upgrade_drain", "enable_upgrade_drain", "drain_max_unavailable_node_percentage"], "syntax": "attribute", "type": "number"}, {"aliases": ["drain node timeout", "duration", "operation timeout"], "anchor": "schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_node_timeout", "description": "Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is..", "document_id": "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kubernetes_upgrade_drain", "enable_upgrade_drain", "drain_node_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["enable vega upgrade mode"], "anchor": "section", "description": "Configuration parameter for enable vega upgrade mode.", "document_id": "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:enable_upgrade_drain:enable_vega_upgrade_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kubernetes_upgrade_drain", "enable_upgrade_drain", "enable_vega_upgrade_mode"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Specify batch upgrade settings for worker nodes within a site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/disable_vega_upgrade_mode/)
- [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/enable_vega_upgrade_mode/)
- [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/kubernetes_upgrade_drain/)
- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
