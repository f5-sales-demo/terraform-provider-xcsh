---
page_title: "timeouts"
subcategory: ""
description: "timeouts for xcsh_k8s_cluster_role_binding."
xcsh_docs: {"aliases": ["duration", "operation timeout", "timeouts"], "body_bytes": 2303, "body_sha256": "sha256:0da1ef698cfc7e73dce79f55d14528a6e7c61f38e3a5c0ea001cb77ac954910c", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster_role_binding:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:timeouts", "parent_id": "xcsh-docs:resources:k8s_cluster_role_binding:reference", "path": "documentation/resources/k8s_cluster_role_binding/properties/timeouts/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster_role_binding", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3302013110231321-1022022233001010-3003220330100331-2213210110000123-0301231200201111-2302330331332300-0221300020011210-2212213201233100", "registry_path": "docs/guides/resources--k8s_cluster_role_binding--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["timeouts"], "schema_version": 1, "sections": [{"aliases": ["create"], "anchor": "schema-timeouts--create", "description": "A string that can be (https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).", "document_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:timeouts", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "create"], "syntax": "attribute", "type": "string"}, {"aliases": ["delete", "duration", "operation timeout"], "anchor": "schema-timeouts--delete", "description": "A string that can be (https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.", "document_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:timeouts", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "delete"], "syntax": "attribute", "type": "string"}, {"aliases": ["read"], "anchor": "schema-timeouts--read", "description": "A string that can be (https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.", "document_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:timeouts", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "read"], "syntax": "attribute", "type": "string"}, {"aliases": ["update"], "anchor": "schema-timeouts--update", "description": "A string that can be (https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).", "document_id": "xcsh-docs:resources:k8s_cluster_role_binding:properties:timeouts", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeouts", "update"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster_role_binding/properties/timeouts/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "timeouts for xcsh_k8s_cluster_role_binding.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["k8s_cluster_role_bindingCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# timeouts

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role_binding/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role_binding/properties/)
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role_binding/properties/)
- [xcsh_k8s_cluster_role_binding](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role_binding/)
