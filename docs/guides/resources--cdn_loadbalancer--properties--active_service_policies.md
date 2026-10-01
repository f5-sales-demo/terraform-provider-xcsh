---
page_title: "active_service_policies"
subcategory: "Load Balancing"
description: "active_service_policies for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1863, "body_sha256": "sha256:43d715a10634118d7d96df88ef1076e4d0e7935302b4eededbaa5383b684176a", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:active_service_policies", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:active_service_policies:policies"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:active_service_policies", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "docs/guides/resources--cdn_loadbalancer--properties--active_service_policies.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["active_service_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/active_service_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "active_service_policies for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_service_policies

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
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

- [active_service_policies](resources--cdn_loadbalancer--properties--active_service_policies.md#section)
- [no_service_policies](resources--cdn_loadbalancer--properties--no_service_policies.md#section)
- [service_policies_from_namespace](resources--cdn_loadbalancer--properties--service_policies_from_namespace.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_service_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [policies](resources--cdn_loadbalancer--properties--active_service_policies--policies.md): complete subsection reference.

## Next pages

- [active_service_policies.policies](resources--cdn_loadbalancer--properties--active_service_policies--policies.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
