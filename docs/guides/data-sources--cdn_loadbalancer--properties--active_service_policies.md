---
page_title: "active_service_policies"
subcategory: "Load Balancing"
description: "active_service_policies for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1611, "body_sha256": "sha256:e9415f24f63e0bdb69ee316db492f7d1d1678664df9fa390990b6d5399ac73d9", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:active_service_policies", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:active_service_policies:policies"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:active_service_policies", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--active_service_policies.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["active_service_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/active_service_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "active_service_policies for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_service_policies

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- active_service_policies

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: active\_service\_policies, no\_service\_policies, service\_policies\_from\_namespace;
Default: no\_service\_policies\] Configuration parameter for active service policies.

Upstream description:

List of service policies.

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

OneOf alternatives in this subsection:

- [active_service_policies](data-sources--cdn_loadbalancer--properties--active_service_policies.md#section)
- [no_service_policies](data-sources--cdn_loadbalancer--properties--no_service_policies.md#section)
- [service_policies_from_namespace](data-sources--cdn_loadbalancer--properties--service_policies_from_namespace.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [policies](data-sources--cdn_loadbalancer--properties--active_service_policies--policies.md): complete subsection reference.

## Next pages

- [active_service_policies.policies](data-sources--cdn_loadbalancer--properties--active_service_policies--policies.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
