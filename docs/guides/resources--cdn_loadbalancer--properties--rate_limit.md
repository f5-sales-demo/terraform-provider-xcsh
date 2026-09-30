---
page_title: "rate_limit"
subcategory: "Load Balancing"
description: "rate_limit for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2695, "body_sha256": "sha256:1589e00bfb2453aa5020ab5213708811c52181fe1a3c34a97da3ae44bc9a7ac3", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:custom_ip_allowed_list", "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:ip_allowed_list", "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:no_ip_allowed_list", "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:no_policies", "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:policies", "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "docs/guides/resources--cdn_loadbalancer--properties--rate_limit.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rate_limit"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/rate_limit/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rate_limit for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rate_limit

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
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

- [custom_ip_allowed_list](resources--cdn_loadbalancer--properties--rate_limit--custom_ip_allowed_list.md): complete subsection reference.

- [ip_allowed_list](resources--cdn_loadbalancer--properties--rate_limit--ip_allowed_list.md): complete subsection reference.

- [no_ip_allowed_list](resources--cdn_loadbalancer--properties--rate_limit--no_ip_allowed_list.md): complete subsection reference.

- [no_policies](resources--cdn_loadbalancer--properties--rate_limit--no_policies.md): complete subsection reference.

- [policies](resources--cdn_loadbalancer--properties--rate_limit--policies.md): complete subsection reference.

- [rate_limiter](resources--cdn_loadbalancer--properties--rate_limit--rate_limiter.md): complete subsection reference.

## Next pages

- [rate_limit.custom_ip_allowed_list](resources--cdn_loadbalancer--properties--rate_limit--custom_ip_allowed_list.md)
- [rate_limit.ip_allowed_list](resources--cdn_loadbalancer--properties--rate_limit--ip_allowed_list.md)
- [rate_limit.no_ip_allowed_list](resources--cdn_loadbalancer--properties--rate_limit--no_ip_allowed_list.md)
- [rate_limit.no_policies](resources--cdn_loadbalancer--properties--rate_limit--no_policies.md)
- [rate_limit.policies](resources--cdn_loadbalancer--properties--rate_limit--policies.md)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--properties--rate_limit--rate_limiter.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
