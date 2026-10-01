---
page_title: "rule_list.rules"
subcategory: "Security"
description: "rule_list.rules for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2643, "body_sha256": "sha256:ce5beff5b4269adf47130acd61252db950f87fa8a46e133cbb50a00451885e56", "child_ids": ["xcsh-docs:resources:service_policy:properties:rule_list:rules:metadata", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec"], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list", "path": "documentation/resources/service_policy/properties/rule_list/rules/index.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["rule_list", "rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/)
- rule_list.rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Define the list of rules (with an order) that should be evaluated by this service policy. Rules are
evaluated from top to bottom in the list.

Upstream description:

Define the list of rules (with an order) that should be evaluated by this service policy. Rules are
evaluated from top to bottom in the list.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/metadata/): complete subsection reference.

- [spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/): complete subsection reference.

## Next pages

- [rule_list.rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/metadata/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
