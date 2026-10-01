---
page_title: "response_cache"
subcategory: "DNS"
description: "response_cache for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": [], "body_bytes": 2768, "body_sha256": "sha256:ec96d1be57b8d8ceea675e0f8ff8c1577819e03c9fe7df208daf1d9e5d96970c", "child_ids": ["xcsh-docs:resources:dns_load_balancer:properties:response_cache:default_response_cache_parameters", "xcsh-docs:resources:dns_load_balancer:properties:response_cache:disable_spec", "xcsh-docs:resources:dns_load_balancer:properties:response_cache:response_cache_parameters"], "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache", "parent_id": "xcsh-docs:resources:dns_load_balancer:reference", "path": "documentation/resources/dns_load_balancer/properties/response_cache/index.md", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["response_cache"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/properties/response_cache/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "response_cache for xcsh_dns_load_balancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# response_cache

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/)
- response_cache

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for response cache.

Upstream description:

Response Cache x-required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_response_cache_parameters",
    "disable_spec"),
  validators.ConflictingObjectAttributes("default_response_cache_parameters",
    "response_cache_parameters"),
  validators.ConflictingObjectAttributes("disable_spec",
    "response_cache_parameters")}
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
  "x-ves-oneof-field-response_cache_parameters_choice": "[\"default_response_cache_parameters\",\"disable\",\"response_cache_parameters\"]"
}
```

Terraform syntax:

```terraform
response_cache {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_response_cache_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/default_response_cache_parameters/): complete subsection reference.

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/disable_spec/): complete subsection reference.

- [response_cache_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/response_cache_parameters/): complete subsection reference.

## Next pages

- [response_cache.default_response_cache_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/default_response_cache_parameters/)
- [response_cache.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/disable_spec/)
- [response_cache.response_cache_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/response_cache_parameters/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/)
- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/)
