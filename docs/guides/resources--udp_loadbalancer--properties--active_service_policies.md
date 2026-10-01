---
page_title: "active_service_policies"
subcategory: ""
description: "active_service_policies for xcsh_udp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1863, "body_sha256": "sha256:aa6b0d62b91da29455d1706da5f91669f88bc58c8c83852b6a9313714ac53a3d", "canonical_id": "xcsh-docs:resources:udp_loadbalancer:properties:active_service_policies", "child_ids": ["xcsh-docs:resources:udp_loadbalancer:properties:active_service_policies:policies"], "collection_id": "xcsh-docs:resources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:udp_loadbalancer:properties:active_service_policies", "parent_id": "xcsh-docs:resources:udp_loadbalancer:reference", "path": "docs/guides/resources--udp_loadbalancer--properties--active_service_policies.md", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["active_service_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/udp_loadbalancer/properties/active_service_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "active_service_policies for xcsh_udp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_service_policies

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md)
- [Property reference](resources--udp_loadbalancer--reference.md)
- active_service_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_service\_policies, no\_service\_policies, service\_policies\_from\_namespace;
Default: no\_service\_policies\] Configuration parameter for active service policies.

Upstream description:

List of service policies.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("policies")}
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

OneOf alternatives in this subsection:

- [active_service_policies](resources--udp_loadbalancer--properties--active_service_policies.md#section)
- [no_service_policies](resources--udp_loadbalancer--properties--no_service_policies.md#section)
- [service_policies_from_namespace](resources--udp_loadbalancer--properties--service_policies_from_namespace.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_service_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [policies](resources--udp_loadbalancer--properties--active_service_policies--policies.md): complete subsection reference.

## Next pages

- [active_service_policies.policies](resources--udp_loadbalancer--properties--active_service_policies--policies.md)
- [Property reference](resources--udp_loadbalancer--reference.md)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md)
