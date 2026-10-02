---
page_title: "rate_limit.rate_limiter.action_block"
subcategory: "Load Balancing"
description: "Action where a user is blocked from making further requests after exceeding rate limit threshold."
xcsh_docs: {"aliases": ["rate limit rate limiter action block"], "body_bytes": 2949, "body_sha256": "sha256:70359d9fdf2c5f4fab481bd4ec851c92731f63e74015150819ee764100da9713", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:hours", "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:minutes", "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:seconds"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter", "path": "documentation/resources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-3103202121112302-1221322032222100-1212030133113120-2323120022200233-3321232113311332-1033310032011022-1300011312223023-0333001032310333", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-023.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter.action_block:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:hours", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter.action_block:ConflictingObjectAttributes:hours,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:hours", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter.action_block:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:minutes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter.action_block:ConflictingObjectAttributes:minutes,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:minutes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter.action_block:ConflictingObjectAttributes:hours,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:seconds", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter.action_block:ConflictingObjectAttributes:minutes,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:seconds", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rate_limit", "rate_limiter", "action_block"], "schema_version": 1, "sections": [{"aliases": ["hours"], "anchor": "section", "description": "Input Duration Hours.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:hours", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "action_block", "hours"], "syntax": "block", "type": "object"}, {"aliases": ["minutes"], "anchor": "section", "description": "Input Duration Minutes.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:minutes", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "action_block", "minutes"], "syntax": "block", "type": "object"}, {"aliases": ["seconds"], "anchor": "section", "description": "Input Duration Seconds.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:seconds", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "action_block", "seconds"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Action where a user is blocked from making further requests after exceeding rate limit threshold.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.rate_limiter.action_block

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/)
- [rate_limit.rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/rate_limiter/)
- rate_limit.rate_limiter.action_block

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Action where a user is blocked from making further requests after exceeding rate limit threshold.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("hours",
    "minutes"),
  validators.ConflictingObjectAttributes("hours",
    "seconds"),
  validators.ConflictingObjectAttributes("minutes",
    "seconds")}
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
  "x-ves-oneof-field-block_duration_choice": "[\"hours\",\"minutes\",\"seconds\"]"
}
```

Terraform syntax:

```terraform
action_block {
  # Configure direct properties listed below.
}
```

## Direct properties

- [hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/hours/): complete subsection reference.

- [minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/minutes/): complete subsection reference.

- [seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/seconds/): complete subsection reference.

## Next pages

- [rate_limit.rate_limiter.action_block.hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/hours/)
- [rate_limit.rate_limiter.action_block.minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/minutes/)
- [rate_limit.rate_limiter.action_block.seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/seconds/)
- [rate_limit.rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/rate_limiter/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
