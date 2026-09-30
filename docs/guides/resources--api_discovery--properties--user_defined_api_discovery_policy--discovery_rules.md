---
page_title: "user_defined_api_discovery_policy.discovery_rules"
subcategory: ""
description: "user_defined_api_discovery_policy.discovery_rules for xcsh_api_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 2826, "body_sha256": "sha256:2917fb479550c2702026dbc6213a0607d61ec8435cca8742771b24a91f258e20", "canonical_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules", "child_ids": ["xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:labels", "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:metadata", "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties"], "collection_id": "xcsh-docs:resources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules", "parent_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy", "path": "docs/guides/resources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules.md", "provider_name": "api_discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["user_defined_api_discovery_policy", "discovery_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "user_defined_api_discovery_policy.discovery_rules for xcsh_api_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# user_defined_api_discovery_policy.discovery_rules

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md)
- [Property reference](resources--api_discovery--reference.md)
- [user_defined_api_discovery_policy](resources--api_discovery--properties--user_defined_api_discovery_policy.md)
- user_defined_api_discovery_policy.discovery_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Define rules to include or exclude endpoints by path, domain, or header. Rules run top to bottom;
unmatched endpoints follow the default action. Defaults to \`\[\]\`. Server applies default when
omitted.

Upstream description:

Define rules to include or exclude endpoints by path, domain, or header. Rules run top to bottom;
unmatched endpoints follow the default action.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
discovery_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [labels](resources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--labels.md): complete subsection reference.

- [metadata](resources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--metadata.md): complete subsection reference.

- [rule_properties](resources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--rule_properties.md): complete subsection reference.

## Next pages

- [user_defined_api_discovery_policy.discovery_rules.labels](resources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--labels.md)
- [user_defined_api_discovery_policy.discovery_rules.metadata](resources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--metadata.md)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](resources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--rule_properties.md)
- [user_defined_api_discovery_policy](resources--api_discovery--properties--user_defined_api_discovery_policy.md)
- [xcsh_api_discovery](../resources/api_discovery.md)
