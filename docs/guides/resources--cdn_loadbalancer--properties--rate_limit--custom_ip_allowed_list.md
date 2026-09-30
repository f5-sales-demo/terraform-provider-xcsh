---
page_title: "rate_limit.custom_ip_allowed_list"
subcategory: "Load Balancing"
description: "rate_limit.custom_ip_allowed_list for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1392, "body_sha256": "sha256:061ddddce81c9c08be921f3eb88949f6e830c85004cc4d178f6e57e7834de031", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:custom_ip_allowed_list", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:custom_ip_allowed_list:rate_limiter_allowed_prefixes"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:custom_ip_allowed_list", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit", "path": "docs/guides/resources--cdn_loadbalancer--properties--rate_limit--custom_ip_allowed_list.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rate_limit", "custom_ip_allowed_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/rate_limit/custom_ip_allowed_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rate_limit.custom_ip_allowed_list for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rate_limit.custom_ip_allowed_list

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [rate_limit](resources--cdn_loadbalancer--properties--rate_limit.md)
- rate_limit.custom_ip_allowed_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IP Allowed list using existing ip\_prefix\_set objects.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rate_limiter_allowed_prefixes")}
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
custom_ip_allowed_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [rate_limiter_allowed_prefixes](resources--cdn_loadbalancer--properties--rate_limit--custom_ip_allowed_list--rate_limiter_allowed_prefixes.md): complete subsection reference.

## Next pages

- [rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes](resources--cdn_loadbalancer--properties--rate_limit--custom_ip_allowed_list--rate_limiter_allowed_prefixes.md)
- [rate_limit](resources--cdn_loadbalancer--properties--rate_limit.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
