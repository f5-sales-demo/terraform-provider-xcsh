---
page_title: "https.coalescing_options"
subcategory: "Load Balancing"
description: "https.coalescing_options for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1783, "body_sha256": "sha256:fd19d26c6acc39a615c82524ccab85ccff07f8f8e171f6f13772d95fc0a8a0fc", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:https:coalescing_options", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:https:coalescing_options:default_coalescing", "xcsh-docs:resources:http_loadbalancer:properties:https:coalescing_options:strict_coalescing"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https:coalescing_options", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https", "path": "docs/guides/resources--http_loadbalancer--properties--https--coalescing_options.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "coalescing_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https/coalescing_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.coalescing_options for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# https.coalescing_options

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [https](resources--http_loadbalancer--properties--https.md)
- https.coalescing_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_coalescing",
    "strict_coalescing")}
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
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_coalescing](resources--http_loadbalancer--properties--https--coalescing_options--default_coalescing.md): complete subsection reference.

- [strict_coalescing](resources--http_loadbalancer--properties--https--coalescing_options--strict_coalescing.md): complete subsection reference.

## Next pages

- [https.coalescing_options.default_coalescing](resources--http_loadbalancer--properties--https--coalescing_options--default_coalescing.md)
- [https.coalescing_options.strict_coalescing](resources--http_loadbalancer--properties--https--coalescing_options--strict_coalescing.md)
- [https](resources--http_loadbalancer--properties--https.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
