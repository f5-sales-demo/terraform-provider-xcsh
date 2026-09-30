---
page_title: "https.default_header"
subcategory: "Load Balancing"
description: "https.default_header for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 902, "body_sha256": "sha256:5b65de762e7a32cf6a1d4abe259e8ddfbde7e5f17830ddabed5303096aa8a1b3", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:https:default_header", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https:default_header", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https", "path": "docs/guides/resources--http_loadbalancer--properties--https--default_header.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "default_header"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https/default_header/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.default_header for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# https.default_header

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [https](resources--http_loadbalancer--properties--https.md)
- https.default_header

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default header.

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
default_header = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [https](resources--http_loadbalancer--properties--https.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
