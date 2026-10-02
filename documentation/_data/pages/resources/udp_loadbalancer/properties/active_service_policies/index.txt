---
page_title: "active_service_policies"
subcategory: ""
description: "List of service policies."
xcsh_docs: {"aliases": ["active service policies"], "body_bytes": 2324, "body_sha256": "sha256:825b38901b0df562768361a7e38b8e9ff1ab083aca81bae03d971a60805bd9f4", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:udp_loadbalancer:properties:active_service_policies:policies"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:udp_loadbalancer:properties:active_service_policies", "parent_id": "xcsh-docs:resources:udp_loadbalancer:reference", "path": "documentation/resources/udp_loadbalancer/properties/active_service_policies/index.md", "product": "distributed-cloud", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0232231312113132-0011113013331010-3331133001111101-2303123300322311-0312001220202312-1122123223310111-2113233111100023-2020111032211132", "registry_path": "docs/guides/resources--udp_loadbalancer--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "active_service_policies:RequiredObjectAttributes:policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:udp_loadbalancer:properties:active_service_policies:policies", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["active_service_policies"], "schema_version": 1, "sections": [{"aliases": ["policies"], "anchor": "section", "description": "Service Policies is a sequential engine where policies (and rules within the policy) are evaluated one after the other. It's important to define the correct order (policies evaluated from top to bottom in the list) for service policies, to GET the intended result. For each request, its characteristics are evaluated", "document_id": "xcsh-docs:resources:udp_loadbalancer:properties:active_service_policies:policies", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-active_service_policies--policies--name", "enforcement": "provider-schema", "group": "active_service_policies.policies:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:udp_loadbalancer:properties:active_service_policies:policies", "type": "requires"}], "schema_path": ["active_service_policies", "policies"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/udp_loadbalancer/properties/active_service_policies/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of service policies.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_service_policies

Breadcrumbs:

- [xcsh_udp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/)
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

- [active_service_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/active_service_policies/#section)
- [no_service_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/no_service_policies/#section)
- [service_policies_from_namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/service_policies_from_namespace/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_service_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/active_service_policies/policies/): complete subsection reference.

## Next pages

- [active_service_policies.policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/active_service_policies/policies/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/)
- [xcsh_udp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/)
