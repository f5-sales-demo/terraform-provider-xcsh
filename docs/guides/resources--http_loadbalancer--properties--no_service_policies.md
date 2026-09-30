---
page_title: "no_service_policies"
subcategory: "Load Balancing"
description: "no_service_policies for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 853, "body_sha256": "sha256:ccbb980af01b2c673f94a01d4d8ea27a6e5e9cc4e1c0db52905d7d8c73a7be89", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:no_service_policies", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:no_service_policies", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "docs/guides/resources--http_loadbalancer--properties--no_service_policies.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["no_service_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/no_service_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "no_service_policies for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# no_service_policies

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- no_service_policies

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no service policies.

Upstream description:

This can be used for messages where no values are needed.

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
no_service_policies = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
