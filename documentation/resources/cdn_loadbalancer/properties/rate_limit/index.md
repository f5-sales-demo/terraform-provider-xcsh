---
page_title: "rate_limit"
subcategory: "Load Balancing"
description: "RateLimitConfigType."
xcsh_docs: {"aliases": ["rate limit"], "body_bytes": 3602, "body_sha256": "sha256:d14146480772c07e93fc8ab9fb3053f1f9473fdef9ad0533223274af9ddf7dca", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:custom_ip_allowed_list", "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:ip_allowed_list", "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:no_ip_allowed_list", "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:no_policies", "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:policies", "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/rate_limit/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-014.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit:ConflictingObjectAttributes:custom_ip_allowed_list,ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:custom_ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit:ConflictingObjectAttributes:custom_ip_allowed_list,no_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:custom_ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit:ConflictingObjectAttributes:custom_ip_allowed_list,ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit:ConflictingObjectAttributes:ip_allowed_list,no_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit:ConflictingObjectAttributes:custom_ip_allowed_list,no_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:no_ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit:ConflictingObjectAttributes:ip_allowed_list,no_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:no_ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit:ConflictingObjectAttributes:no_policies,policies", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:no_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit:ConflictingObjectAttributes:no_policies,policies", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:policies", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rate_limit"], "schema_version": 1, "sections": [{"aliases": ["custom ip allowed list"], "anchor": "section", "description": "IP Allowed list using existing ip_prefix_set objects.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:custom_ip_allowed_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.custom_ip_allowed_list:RequiredObjectAttributes:rate_limiter_allowed_prefixes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:custom_ip_allowed_list:rate_limiter_allowed_prefixes", "type": "requires"}], "schema_path": ["rate_limit", "custom_ip_allowed_list"], "syntax": "block", "type": "object"}, {"aliases": ["ip allowed list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:ip_allowed_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rate_limit", "ip_allowed_list"], "syntax": "block", "type": "object"}, {"aliases": ["no ip allowed list"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:no_ip_allowed_list", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "no_ip_allowed_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["no policies"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:no_policies", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "no_policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["policies"], "anchor": "section", "description": "List of rate limiter policies to be applied.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:policies", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.policies:RequiredObjectAttributes:policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:policies:policies", "type": "requires"}], "schema_path": ["rate_limit", "policies"], "syntax": "block", "type": "object"}, {"aliases": ["rate limiter"], "anchor": "section", "description": "A tuple consisting of a rate limit period unit and the total number of allowed requests for that period.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter:ConflictingObjectAttributes:action_block,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter:ConflictingObjectAttributes:action_block,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter:ConflictingObjectAttributes:leaky_bucket,token_bucket", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:leaky_bucket", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter:ConflictingObjectAttributes:leaky_bucket,token_bucket", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:token_bucket", "type": "conflicts"}, {"anchor": "schema-rate_limit--rate_limiter--total_number", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter:RequiredObjectAttributes:total_number", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter", "type": "requires"}], "schema_path": ["rate_limit", "rate_limiter"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/rate_limit/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "RateLimitConfigType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- rate_limit

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

RateLimitConfigType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "ip_allowed_list"),
  validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("ip_allowed_list",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("no_policies",
    "policies")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]",
  "x-ves-oneof-field-policy_choice": "[\"no_policies\",\"policies\"]"
}
```

Terraform syntax:

```terraform
rate_limit {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/custom_ip_allowed_list/): complete subsection reference.

- [ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/ip_allowed_list/): complete subsection reference.

- [no_ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/no_ip_allowed_list/): complete subsection reference.

- [no_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/no_policies/): complete subsection reference.

- [policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/policies/): complete subsection reference.

- [rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/): complete subsection reference.

## Next pages

- [rate_limit.custom_ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/custom_ip_allowed_list/)
- [rate_limit.ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/ip_allowed_list/)
- [rate_limit.no_ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/no_ip_allowed_list/)
- [rate_limit.no_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/no_policies/)
- [rate_limit.policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/policies/)
- [rate_limit.rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
