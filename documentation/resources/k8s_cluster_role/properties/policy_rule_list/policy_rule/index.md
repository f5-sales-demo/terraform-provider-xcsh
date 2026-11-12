---
page_title: "policy_rule_list.policy_rule"
subcategory: "Container"
description: "List of rules for role permissions."
xcsh_docs: {"aliases": ["policy rule list policy rule"], "body_bytes": 2388, "body_sha256": "sha256:a05f35dacfae65c001e9bb9c59932ed58fb23f4072869c6708d4c5551480bc82", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:non_resource_url_list", "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:k8s_cluster_role:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule", "parent_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list", "path": "documentation/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-3031232101002032-1311120200232222-2332131333022232-2322030333100023-1022010033311002-1333100321120122-0301021330113311-2030031003323331", "registry_path": "docs/guides/resources--k8s_cluster_role--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "policy_rule_list.policy_rule:ConflictingListObjectAttributes:non_resource_url_list,resource_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:non_resource_url_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_rule_list.policy_rule:ConflictingListObjectAttributes:non_resource_url_list,resource_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_rule_list", "policy_rule"], "schema_version": 1, "sections": [{"aliases": ["policy rule list policy rule non resource url list"], "anchor": "section", "description": "Permissions for URL(s) that do not represent K8s resource.", "document_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:non_resource_url_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-policy_rule_list--policy_rule--non_resource_url_list--urls", "enforcement": "provider-schema", "group": "policy_rule_list.policy_rule.non_resource_url_list:RequiredObjectAttributes:urls,verbs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:non_resource_url_list", "type": "requires"}, {"anchor": "schema-policy_rule_list--policy_rule--non_resource_url_list--verbs", "enforcement": "provider-schema", "group": "policy_rule_list.policy_rule.non_resource_url_list:RequiredObjectAttributes:urls,verbs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:non_resource_url_list", "type": "requires"}], "schema_path": ["policy_rule_list", "policy_rule", "non_resource_url_list"], "syntax": "block", "type": "object"}, {"aliases": ["policy rule list policy rule resource list"], "anchor": "section", "description": "List of resources in terms of API groups/resource types/resource instances and verbs allowed.", "document_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-policy_rule_list--policy_rule--resource_list--api_groups", "enforcement": "provider-schema", "group": "policy_rule_list.policy_rule.resource_list:RequiredObjectAttributes:api_groups,resource_types,verbs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list", "type": "requires"}, {"anchor": "schema-policy_rule_list--policy_rule--resource_list--resource_types", "enforcement": "provider-schema", "group": "policy_rule_list.policy_rule.resource_list:RequiredObjectAttributes:api_groups,resource_types,verbs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list", "type": "requires"}, {"anchor": "schema-policy_rule_list--policy_rule--resource_list--verbs", "enforcement": "provider-schema", "group": "policy_rule_list.policy_rule.resource_list:RequiredObjectAttributes:api_groups,resource_types,verbs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list:policy_rule:resource_list", "type": "requires"}], "schema_path": ["policy_rule_list", "policy_rule", "resource_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "List of rules for role permissions.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_rule_list.policy_rule

Breadcrumbs:

- [xcsh_k8s_cluster_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/)
- [policy_rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/)
- policy_rule_list.policy_rule

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Policy Rules. List of rules for role permissions.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("non_resource_url_list",
    "resource_list")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-29T12:14:42+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
policy_rule {
  # Configure direct properties listed below.
}
```

## Direct properties

- [non_resource_url_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/non_resource_url_list/): complete subsection reference.

- [resource_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/resource_list/): complete subsection reference.
