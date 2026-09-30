---
page_title: "active_service_policies"
subcategory: ""
description: "active_service_policies for xcsh_udp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1512, "body_sha256": "sha256:94a3c3cbd5f546fa6244398258dbc71fb5d49611936b10da36e8befb14669a3a", "canonical_id": "xcsh-docs:data-sources:udp_loadbalancer:properties:active_service_policies", "child_ids": ["xcsh-docs:data-sources:udp_loadbalancer:properties:active_service_policies:policies"], "collection_id": "xcsh-docs:data-sources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:udp_loadbalancer:properties:active_service_policies", "parent_id": "xcsh-docs:data-sources:udp_loadbalancer:reference", "path": "docs/guides/data-sources--udp_loadbalancer--properties--active_service_policies.md", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["active_service_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/udp_loadbalancer/properties/active_service_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "active_service_policies for xcsh_udp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# active_service_policies

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md)
- [Property reference](data-sources--udp_loadbalancer--reference.md)
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

- [active_service_policies](data-sources--udp_loadbalancer--properties--active_service_policies.md#section)
- [no_service_policies](data-sources--udp_loadbalancer--properties--no_service_policies.md#section)
- [service_policies_from_namespace](data-sources--udp_loadbalancer--properties--service_policies_from_namespace.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [policies](data-sources--udp_loadbalancer--properties--active_service_policies--policies.md): complete subsection reference.

## Next pages

- [active_service_policies.policies](data-sources--udp_loadbalancer--properties--active_service_policies--policies.md)
- [Property reference](data-sources--udp_loadbalancer--reference.md)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md)
