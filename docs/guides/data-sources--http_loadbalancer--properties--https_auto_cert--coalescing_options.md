---
page_title: "https_auto_cert.coalescing_options"
subcategory: "Load Balancing"
description: "https_auto_cert.coalescing_options for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1618, "body_sha256": "sha256:a75b0f71a121043f5c33189a8338857bb600fa69b2a683f63c5bbcda3632fbcc", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https_auto_cert:coalescing_options", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:https_auto_cert:coalescing_options:default_coalescing", "xcsh-docs:data-sources:http_loadbalancer:properties:https_auto_cert:coalescing_options:strict_coalescing"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:https_auto_cert:coalescing_options", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https_auto_cert", "path": "docs/guides/data-sources--http_loadbalancer--properties--https_auto_cert--coalescing_options.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_auto_cert", "coalescing_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/https_auto_cert/coalescing_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_auto_cert.coalescing_options for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# https_auto_cert.coalescing_options

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [https_auto_cert](data-sources--http_loadbalancer--properties--https_auto_cert.md)
- https_auto_cert.coalescing_options

<a id="section"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

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

## Direct properties

- [default_coalescing](data-sources--http_loadbalancer--properties--https_auto_cert--coalescing_options--default_coalescing.md): complete subsection reference.

- [strict_coalescing](data-sources--http_loadbalancer--properties--https_auto_cert--coalescing_options--strict_coalescing.md): complete subsection reference.

## Next pages

- [https_auto_cert.coalescing_options.default_coalescing](data-sources--http_loadbalancer--properties--https_auto_cert--coalescing_options--default_coalescing.md)
- [https_auto_cert.coalescing_options.strict_coalescing](data-sources--http_loadbalancer--properties--https_auto_cert--coalescing_options--strict_coalescing.md)
- [https_auto_cert](data-sources--http_loadbalancer--properties--https_auto_cert.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
