---
page_title: "rate_limit.policies"
subcategory: "Load Balancing"
description: "rate_limit.policies for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1192, "body_sha256": "sha256:ecfc11964ef052f87fa3d9af3ef7171fd710f3696f48bccb61c92389afa416ca", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:policies", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:policies:policies"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:policies", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit", "path": "docs/guides/resources--cdn_loadbalancer--properties--rate_limit--policies.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rate_limit", "policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/rate_limit/policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rate_limit.policies for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rate_limit.policies

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [rate_limit](resources--cdn_loadbalancer--properties--rate_limit.md)
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

- [policies](resources--cdn_loadbalancer--properties--rate_limit--policies--policies.md): complete subsection reference.

## Next pages

- [rate_limit.policies.policies](resources--cdn_loadbalancer--properties--rate_limit--policies--policies.md)
- [rate_limit](resources--cdn_loadbalancer--properties--rate_limit.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
