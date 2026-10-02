---
page_title: "timeouts"
subcategory: ""
description: "timeouts for xcsh_cminstance."
xcsh_docs: {"aliases": ["duration", "operation timeout", "timeouts"], "body_bytes": 2219, "body_sha256": "sha256:f881c757da5dd3cfe009257f94a36d254ca71253f20269e46ffefd3473d8a8f7", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cminstance:collection", "completeness": "complete", "id": "xcsh-docs:resources:cminstance:properties:timeouts", "parent_id": "xcsh-docs:resources:cminstance:reference", "path": "documentation/resources/cminstance/properties/timeouts/index.md", "product": "distributed-cloud", "provider_name": "cminstance", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3012012022112030-3310230301133232-1001311133212011-3332112331333221-3032000313020313-0020131223101003-2012031130303010-3030332222211210", "registry_path": "docs/guides/resources--cminstance--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["timeouts"], "schema_version": 1, "sections": [{"aliases": ["create"], "anchor": "schema-timeouts--create", "description": "A string that can be (https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).", "document_id": "xcsh-docs:resources:cminstance:properties:timeouts", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "create"], "syntax": "attribute", "type": "string"}, {"aliases": ["delete", "duration", "operation timeout"], "anchor": "schema-timeouts--delete", "description": "A string that can be (https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.", "document_id": "xcsh-docs:resources:cminstance:properties:timeouts", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "delete"], "syntax": "attribute", "type": "string"}, {"aliases": ["read"], "anchor": "schema-timeouts--read", "description": "A string that can be (https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.", "document_id": "xcsh-docs:resources:cminstance:properties:timeouts", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "read"], "syntax": "attribute", "type": "string"}, {"aliases": ["update"], "anchor": "schema-timeouts--update", "description": "A string that can be (https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).", "document_id": "xcsh-docs:resources:cminstance:properties:timeouts", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "update"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cminstance/properties/timeouts/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "timeouts for xcsh_cminstance.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cminstanceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# timeouts

Breadcrumbs:

- [xcsh_cminstance](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/)
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/)
- [xcsh_cminstance](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/)
