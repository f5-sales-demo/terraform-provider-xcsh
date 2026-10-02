---
page_title: "active_service_policies"
subcategory: "Load Balancing"
description: "List of service policies."
xcsh_docs: {"aliases": ["active service policies"], "body_bytes": 2324, "body_sha256": "sha256:cb88ba0974c970daa5f3132b95f3c7e08e6bce6a5c49de9f1a9976dc9b08ed17", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:tcp_loadbalancer:properties:active_service_policies:policies"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:active_service_policies", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:reference", "path": "documentation/resources/tcp_loadbalancer/properties/active_service_policies/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2123332010300020-1023231311030232-1300031122101011-1032023203210220-1200333010322303-0323133123222003-1020130133223133-1112131030203102", "registry_path": "docs/guides/resources--tcp_loadbalancer--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "active_service_policies:RequiredObjectAttributes:policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:active_service_policies:policies", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["active_service_policies"], "schema_version": 1, "sections": [{"aliases": ["policies"], "anchor": "section", "description": "Service Policies is a sequential engine where policies (and rules within the policy) are evaluated one after the other. It's important to define the correct order (policies evaluated from top to bottom in the list) for service policies, to GET the intended result. For each request, its characteristics are evaluated", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:active_service_policies:policies", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-active_service_policies--policies--name", "enforcement": "provider-schema", "group": "active_service_policies.policies:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:active_service_policies:policies", "type": "requires"}], "schema_path": ["active_service_policies", "policies"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/active_service_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of service policies.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_service_policies

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/)
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

- [active_service_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/active_service_policies/#section)
- [no_service_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/no_service_policies/#section)
- [service_policies_from_namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/service_policies_from_namespace/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_service_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/active_service_policies/policies/): complete subsection reference.

## Next pages

- [active_service_policies.policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/active_service_policies/policies/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/)
- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
