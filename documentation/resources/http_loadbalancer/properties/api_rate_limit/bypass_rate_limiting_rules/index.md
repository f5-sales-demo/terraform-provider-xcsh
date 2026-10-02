---
page_title: "api_rate_limit.bypass_rate_limiting_rules"
subcategory: "Load Balancing"
description: "This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting."
xcsh_docs: {"aliases": ["api rate limit bypass rate limiting rules"], "body_bytes": 1916, "body_sha256": "sha256:2b5ba0784e48bb0f256fb8c89ce9ed734cb9cb237128dadd08bcd4b505963bb4", "capabilities": ["load-balancing", "security.rate-limiting"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit", "path": "documentation/resources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "bypass_rate_limiting_rules"], "schema_version": 1, "sections": [{"aliases": ["bypass rate limiting rules"], "anchor": "section", "description": "This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--base_path", "enforcement": "provider-schema", "group": "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules:ConflictingListObjectAttributes:any_url,base_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules", "type": "conflicts"}, {"anchor": "schema-api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--base_path", "enforcement": "provider-schema", "group": "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules:ConflictingListObjectAttributes:api_endpoint,base_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules", "type": "conflicts"}, {"anchor": "schema-api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--base_path", "enforcement": "provider-schema", "group": "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules:ConflictingListObjectAttributes:api_groups,base_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules", "type": "conflicts"}, {"anchor": "schema-api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--specific_domain", "enforcement": "provider-schema", "group": "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules:ConflictingListObjectAttributes:any_domain,specific_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules:ConflictingListObjectAttributes:any_domain,specific_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules:ConflictingListObjectAttributes:any_url,api_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:any_url", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules:ConflictingListObjectAttributes:any_url,api_groups", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:any_url", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules:ConflictingListObjectAttributes:any_url,base_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:any_url", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules:ConflictingListObjectAttributes:any_url,api_endpoint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:api_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules:ConflictingListObjectAttributes:api_endpoint,api_groups", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:api_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules:ConflictingListObjectAttributes:api_endpoint,base_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:api_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules:ConflictingListObjectAttributes:any_url,api_groups", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:api_groups", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules:ConflictingListObjectAttributes:api_endpoint,api_groups", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:api_groups", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules:ConflictingListObjectAttributes:api_groups,base_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:api_groups", "type": "conflicts"}], "schema_path": ["api_rate_limit", "bypass_rate_limiting_rules", "bypass_rate_limiting_rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.bypass_rate_limiting_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/)
- api_rate_limit.bypass_rate_limiting_rules

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

Upstream description:

This category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

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
bypass_rate_limiting_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bypass_rate_limiting_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/): complete subsection reference.

## Next pages

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
