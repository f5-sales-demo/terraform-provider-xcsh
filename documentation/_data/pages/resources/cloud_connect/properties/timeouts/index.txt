---
page_title: "timeouts"
subcategory: ""
description: "timeouts for xcsh_cloud_connect."
xcsh_docs: {"aliases": ["duration", "lifecycle timeout", "operation timeout", "timeouts"], "body_bytes": 2237, "body_sha256": "sha256:a38d0891a7c34553841b3a73fbbea624a130f18f91ea9ebbefde4de775a9a142", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:timeouts", "parent_id": "xcsh-docs:resources:cloud_connect:reference", "path": "documentation/resources/cloud_connect/properties/timeouts/index.md", "product": "distributed-cloud", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0102000200203213-0133023212110222-3021012023202203-0022100231003022-1321013002322210-0303322031111330-1031321201022011-1032333113131332", "registry_path": "docs/guides/resources--cloud_connect--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["timeouts"], "schema_version": 1, "sections": [{"aliases": ["create operation timeout", "create timeout", "creation duration", "duration", "lifecycle timeout", "operation timeout", "timeouts create"], "anchor": "schema-timeouts--create", "description": "Configures the timeout duration for resource creation.", "document_id": "xcsh-docs:resources:cloud_connect:properties:timeouts", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "create"], "syntax": "attribute", "type": "string"}, {"aliases": ["delete operation timeout", "delete timeout", "deletion duration", "destroy timeout", "duration", "lifecycle timeout", "operation timeout", "timeouts delete"], "anchor": "schema-timeouts--delete", "description": "Configures the timeout duration for resource deletion, applicable only if changes are saved into state before destroy.", "document_id": "xcsh-docs:resources:cloud_connect:properties:timeouts", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "delete"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "lifecycle timeout", "operation timeout", "read duration", "read operation timeout", "read timeout", "refresh timeout", "timeouts read"], "anchor": "schema-timeouts--read", "description": "Configures the timeout duration for read operations during refresh or planning.", "document_id": "xcsh-docs:resources:cloud_connect:properties:timeouts", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "read"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "lifecycle timeout", "modification timeout", "operation timeout", "timeouts update", "update duration", "update operation timeout", "update timeout"], "anchor": "schema-timeouts--update", "description": "Configures the timeout duration for resource update operations.", "document_id": "xcsh-docs:resources:cloud_connect:properties:timeouts", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "update"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/timeouts/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "timeouts for xcsh_cloud_connect.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# timeouts

Breadcrumbs:

- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/)
- timeouts

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-timeouts--create"></a>

### create property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="schema-timeouts--delete"></a>

### delete property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="schema-timeouts--read"></a>

### read property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="schema-timeouts--update"></a>

### update property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/)
- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
