---
page_title: "user_defined_api_discovery_policy.discovery_rules"
subcategory: ""
description: "Define rules to include or exclude endpoints by path, domain, or header. Rules run top to bottom; unmatched endpoints follow the default action."
xcsh_docs: {"aliases": ["user defined api discovery policy discovery rules"], "body_bytes": 2366, "body_sha256": "sha256:393094b8773cb0cb42a54e149362c5ff77f334fcd4748e2344344859c27c2e22", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:labels", "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:metadata", "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules", "parent_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy", "path": "documentation/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/index.md", "product": "distributed-cloud", "provider_name": "api_discovery", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1100332332123333-2101003132012032-3320001300132012-1210013330331020-2321321103303010-1322120010232310-2001131223011002-0010130121220313", "registry_path": "docs/guides/resources--api_discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["user_defined_api_discovery_policy", "discovery_rules"], "schema_version": 1, "sections": [{"aliases": ["user defined api discovery policy discovery rules labels"], "anchor": "section", "description": "Map of string keys and values that can be used to organize and categorize the rule.", "document_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:labels", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "labels"], "syntax": "block", "type": "object"}, {"aliases": ["user defined api discovery policy discovery rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-user_defined_api_discovery_policy--discovery_rules--metadata--name", "enforcement": "provider-schema", "group": "user_defined_api_discovery_policy.discovery_rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:metadata", "type": "requires"}], "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["user defined api discovery policy discovery rules rule properties"], "anchor": "section", "description": "Determines whether matching endpoints are included in API Discovery or excluded.", "document_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--pattern", "enforcement": "provider-schema", "group": "user_defined_api_discovery_policy.discovery_rules.rule_properties:ConflictingObjectAttributes:http_header_criteria,pattern", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "user_defined_api_discovery_policy.discovery_rules.rule_properties:ConflictingObjectAttributes:exclusion,inclusion", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:exclusion", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "user_defined_api_discovery_policy.discovery_rules.rule_properties:ConflictingObjectAttributes:http_header_criteria,pattern", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:http_header_criteria", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "user_defined_api_discovery_policy.discovery_rules.rule_properties:ConflictingObjectAttributes:exclusion,inclusion", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:inclusion", "type": "conflicts"}], "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Define rules to include or exclude endpoints by path, domain, or header. Rules run top to bottom; unmatched endpoints follow the default action.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_defined_api_discovery_policy.discovery_rules

Breadcrumbs:

- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/)
- [user_defined_api_discovery_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/)
- user_defined_api_discovery_policy.discovery_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Define rules to include or exclude endpoints by path, domain, or header. Rules run top to bottom;
unmatched endpoints follow the default action. Defaults to \`\[\]\`. Server applies default when
omitted.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/labels/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/metadata/): complete subsection reference.

- [rule_properties](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/): complete subsection reference.
