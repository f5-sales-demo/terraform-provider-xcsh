---
page_title: "rate_limit.policies"
subcategory: "Load Balancing"
description: "List of rate limiter policies to be applied."
xcsh_docs: {"aliases": ["rate limit policies"], "body_bytes": 1655, "body_sha256": "sha256:b673eba94e85e9c715fb9884583f605fe9ab88a040097dc1bbdcadb1103ec329", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:rate_limit:policies:policies"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:policies", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit", "path": "documentation/resources/http_loadbalancer/properties/rate_limit/policies/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1201121223001130-0102301012320112-0332103302021323-3120320310211121-2331102322201132-0130330012330030-3122013001202010-1320120013120101", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-023.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.policies:RequiredObjectAttributes:policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:policies:policies", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rate_limit", "policies"], "schema_version": 1, "sections": [{"aliases": ["policies"], "anchor": "section", "description": "Ordered list of rate limiter policies.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:policies:policies", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-rate_limit--policies--policies--name", "enforcement": "provider-schema", "group": "rate_limit.policies.policies:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:policies:policies", "type": "requires"}], "schema_path": ["rate_limit", "policies", "policies"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/rate_limit/policies/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of rate limiter policies to be applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.policies

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/)
- rate_limit.policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of rate limiter policies to be applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("policies")}
```

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
policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/policies/policies/): complete subsection reference.

## Next pages

- [rate_limit.policies.policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/policies/policies/)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
