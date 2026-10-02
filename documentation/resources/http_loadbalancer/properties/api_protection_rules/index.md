---
page_title: "api_protection_rules"
subcategory: "Load Balancing"
description: "API Protection Rules."
xcsh_docs: {"aliases": ["api protection rules"], "body_bytes": 1850, "body_sha256": "sha256:bc7acb1341cdaf4008edd9dd329c4c162feacb341111dfd10e8b277c47ee11a6", "capabilities": ["load-balancing", "security.api-protection"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/api_protection_rules/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_protection_rules"], "schema_version": 1, "sections": [{"aliases": ["api endpoint rules"], "anchor": "section", "description": "This category defines specific rules per API endpoints. If request matches any of these rules, skipping second category rules.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-api_protection_rules--api_endpoint_rules--specific_domain", "enforcement": "provider-schema", "group": "api_protection_rules.api_endpoint_rules:ConflictingListObjectAttributes:any_domain,specific_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_endpoint_rules:ConflictingListObjectAttributes:any_domain,specific_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:any_domain", "type": "conflicts"}, {"anchor": "schema-api_protection_rules--api_endpoint_rules--api_endpoint_path", "enforcement": "provider-schema", "group": "api_protection_rules.api_endpoint_rules:RequiredListObjectAttributes:api_endpoint_path", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules", "type": "requires"}], "schema_path": ["api_protection_rules", "api_endpoint_rules"], "syntax": "block", "type": "object"}, {"aliases": ["api groups rules"], "anchor": "section", "description": "This category includes rules per API group or Server URL. For API groups, refer to API Definition which includes API groups derived from uploaded swaggers.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-api_protection_rules--api_groups_rules--specific_domain", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules:ConflictingListObjectAttributes:any_domain,specific_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules:ConflictingListObjectAttributes:any_domain,specific_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:any_domain", "type": "conflicts"}, {"anchor": "schema-api_protection_rules--api_groups_rules--base_path", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules:RequiredListObjectAttributes:base_path", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules", "type": "requires"}], "schema_path": ["api_protection_rules", "api_groups_rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_protection_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "API Protection Rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- api_protection_rules

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

API Protection Rules. API Protection Rules.

Upstream description:

API Protection Rules.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
api_protection_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/): complete subsection reference.

- [api_groups_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/): complete subsection reference.

## Next pages

- [api_protection_rules.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/)
- [api_protection_rules.api_groups_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
