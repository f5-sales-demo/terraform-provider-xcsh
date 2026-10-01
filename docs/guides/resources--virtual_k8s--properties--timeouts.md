---
page_title: "timeouts"
subcategory: "Container"
description: "timeouts for xcsh_virtual_k8s."
xcsh_docs: {"aliases": [], "body_bytes": 2017, "body_sha256": "sha256:8e20017014b1741f6ce4d1affb9005170f1a9163d27ff75946b495ff6d749689", "canonical_id": "xcsh-docs:resources:virtual_k8s:properties:timeouts", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_k8s:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_k8s:properties:timeouts", "parent_id": "xcsh-docs:resources:virtual_k8s:reference", "path": "docs/guides/resources--virtual_k8s--properties--timeouts.md", "provider_name": "virtual_k8s", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["timeouts"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_k8s/properties/timeouts/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "timeouts for xcsh_virtual_k8s.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_k8sCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# timeouts

Breadcrumbs:

- [xcsh_virtual_k8s](../resources/virtual_k8s.md)
- [Property reference](resources--virtual_k8s--reference.md)
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

- [Property reference](resources--virtual_k8s--reference.md)
- [xcsh_virtual_k8s](../resources/virtual_k8s.md)
