---
page_title: "timeouts"
subcategory: ""
description: "timeouts for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2120, "body_sha256": "sha256:b25fc8b5b7ddbba126f2b6c9115314969b65c2696fe8a6b9b251fe22d1d99c25", "child_ids": [], "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:timeouts", "parent_id": "xcsh-docs:resources:nat_policy:reference", "path": "documentation/resources/nat_policy/properties/timeouts/index.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["timeouts"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/timeouts/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "timeouts for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# timeouts

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/)
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
