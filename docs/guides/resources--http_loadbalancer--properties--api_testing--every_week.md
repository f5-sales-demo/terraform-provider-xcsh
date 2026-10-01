---
page_title: "api_testing.every_week"
subcategory: "Load Balancing"
description: "api_testing.every_week for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1000, "body_sha256": "sha256:27e775007fab66297e566342cb51a504637a2925063d47f4e8a558667b10ed2f", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_week", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_week", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing", "path": "docs/guides/resources--http_loadbalancer--properties--api_testing--every_week.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_testing", "every_week"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_testing/every_week/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_testing.every_week for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_testing.every_week

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_testing](resources--http_loadbalancer--properties--api_testing.md)
- api_testing.every_week

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
every_week = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [api_testing](resources--http_loadbalancer--properties--api_testing.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
