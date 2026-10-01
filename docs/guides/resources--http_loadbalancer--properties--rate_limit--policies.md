---
page_title: "rate_limit.policies"
subcategory: "Load Balancing"
description: "rate_limit.policies for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1300, "body_sha256": "sha256:b585d6a76ded0e1fb32b642dd6642cb6899a3f380be5a179bac0cdff144e6208", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:policies", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:rate_limit:policies:policies"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:policies", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit", "path": "docs/guides/resources--http_loadbalancer--properties--rate_limit--policies.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rate_limit", "policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/rate_limit/policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rate_limit.policies for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.policies

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [rate_limit](resources--http_loadbalancer--properties--rate_limit.md)
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

- [policies](resources--http_loadbalancer--properties--rate_limit--policies--policies.md): complete subsection reference.

## Next pages

- [rate_limit.policies.policies](resources--http_loadbalancer--properties--rate_limit--policies--policies.md)
- [rate_limit](resources--http_loadbalancer--properties--rate_limit.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
