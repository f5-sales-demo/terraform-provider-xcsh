---
page_title: "response_cache"
subcategory: "DNS"
description: "response_cache for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": [], "body_bytes": 2161, "body_sha256": "sha256:0c89fd252ecfaa5676f722d7c45dbb51da313e296236b0b2e5c64d12878cf38b", "canonical_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache", "child_ids": ["xcsh-docs:resources:dns_load_balancer:properties:response_cache:default_response_cache_parameters", "xcsh-docs:resources:dns_load_balancer:properties:response_cache:disable_spec", "xcsh-docs:resources:dns_load_balancer:properties:response_cache:response_cache_parameters"], "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache", "parent_id": "xcsh-docs:resources:dns_load_balancer:reference", "path": "docs/guides/resources--dns_load_balancer--properties--response_cache.md", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["response_cache"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/properties/response_cache/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "response_cache for xcsh_dns_load_balancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# response_cache

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md)
- [Property reference](resources--dns_load_balancer--reference.md)
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

- [default_response_cache_parameters](resources--dns_load_balancer--properties--response_cache--default_response_cache_parameters.md): complete subsection reference.

- [disable_spec](resources--dns_load_balancer--properties--response_cache--disable_spec.md): complete subsection reference.

- [response_cache_parameters](resources--dns_load_balancer--properties--response_cache--response_cache_parameters.md): complete subsection reference.

## Next pages

- [response_cache.default_response_cache_parameters](resources--dns_load_balancer--properties--response_cache--default_response_cache_parameters.md)
- [response_cache.disable_spec](resources--dns_load_balancer--properties--response_cache--disable_spec.md)
- [response_cache.response_cache_parameters](resources--dns_load_balancer--properties--response_cache--response_cache_parameters.md)
- [Property reference](resources--dns_load_balancer--reference.md)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md)
