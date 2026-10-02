---
page_title: "timeouts"
subcategory: ""
description: "timeouts for xcsh_alert_template."
xcsh_docs: {"aliases": ["duration", "operation timeout", "timeouts"], "body_bytes": 2243, "body_sha256": "sha256:f9cf2483bf3c97bad1817ffabfaa16bb380bb20add48198e67a1f30f706648a9", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_template:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_template:properties:timeouts", "parent_id": "xcsh-docs:resources:alert_template:reference", "path": "documentation/resources/alert_template/properties/timeouts/index.md", "product": "distributed-cloud", "provider_name": "alert_template", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2320323033312231-0231121320223030-0122312332202102-2230311233010203-2101322220112001-2022220233000003-3022200323123230-2221323010212232", "registry_path": "docs/guides/resources--alert_template--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["timeouts"], "schema_version": 1, "sections": [{"aliases": ["create"], "anchor": "schema-timeouts--create", "description": "A string that can be (https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).", "document_id": "xcsh-docs:resources:alert_template:properties:timeouts", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "create"], "syntax": "attribute", "type": "string"}, {"aliases": ["delete", "duration", "operation timeout"], "anchor": "schema-timeouts--delete", "description": "A string that can be (https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.", "document_id": "xcsh-docs:resources:alert_template:properties:timeouts", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "delete"], "syntax": "attribute", "type": "string"}, {"aliases": ["read"], "anchor": "schema-timeouts--read", "description": "A string that can be (https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.", "document_id": "xcsh-docs:resources:alert_template:properties:timeouts", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "read"], "syntax": "attribute", "type": "string"}, {"aliases": ["update"], "anchor": "schema-timeouts--update", "description": "A string that can be (https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).", "document_id": "xcsh-docs:resources:alert_template:properties:timeouts", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "update"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_template/properties/timeouts/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "timeouts for xcsh_alert_template.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["alert_templateCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# timeouts

Breadcrumbs:

- [xcsh_alert_template](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_template/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_template/properties/)
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_template/properties/)
- [xcsh_alert_template](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_template/)
