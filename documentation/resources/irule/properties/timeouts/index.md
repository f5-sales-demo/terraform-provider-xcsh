---
page_title: "timeouts"
subcategory: ""
description: "timeouts for xcsh_irule."
xcsh_docs: {"aliases": ["duration", "lifecycle timeout", "operation timeout", "timeouts"], "body_bytes": 1976, "body_sha256": "sha256:90b77579df07d1d5e7ecc04b8e0da447bda9c1d4b1178552102b3557ae0f2ada", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:irule:collection", "completeness": "complete", "id": "xcsh-docs:resources:irule:properties:timeouts", "parent_id": "xcsh-docs:resources:irule:reference", "path": "documentation/resources/irule/properties/timeouts/index.md", "product": "distributed-cloud", "provider_name": "irule", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0031111210001221-1233022110020210-2022201202303023-3321313323310310-1323310012031320-0122202302102210-0000030321013110-1110200301110212", "registry_path": "docs/guides/resources--irule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["timeouts"], "schema_version": 1, "sections": [{"aliases": ["create operation timeout", "create timeout", "creation duration", "duration", "lifecycle timeout", "operation timeout", "timeouts create"], "anchor": "schema-timeouts--create", "description": "Configures the timeout duration for resource creation.", "document_id": "xcsh-docs:resources:irule:properties:timeouts", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "create"], "syntax": "attribute", "type": "string"}, {"aliases": ["delete operation timeout", "delete timeout", "deletion duration", "destroy timeout", "duration", "lifecycle timeout", "operation timeout", "timeouts delete"], "anchor": "schema-timeouts--delete", "description": "Configures the timeout duration for resource deletion, applicable only if changes are saved into state before destroy.", "document_id": "xcsh-docs:resources:irule:properties:timeouts", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "delete"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "lifecycle timeout", "operation timeout", "read duration", "read operation timeout", "read timeout", "refresh timeout", "timeouts read"], "anchor": "schema-timeouts--read", "description": "Configures the timeout duration for read operations during refresh or planning.", "document_id": "xcsh-docs:resources:irule:properties:timeouts", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "read"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "lifecycle timeout", "modification timeout", "operation timeout", "timeouts update", "update duration", "update operation timeout", "update timeout"], "anchor": "schema-timeouts--update", "description": "Configures the timeout duration for resource update operations.", "document_id": "xcsh-docs:resources:irule:properties:timeouts", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "update"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/irule/properties/timeouts/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "timeouts for xcsh_irule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["iruleCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# timeouts

Breadcrumbs:

- [xcsh_irule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/properties/)
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
