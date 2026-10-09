---
page_title: "routes.route_destination.csrf_policy"
subcategory: ""
description: "To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its destination.the policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request"
xcsh_docs: {"aliases": ["routes route destination csrf policy"], "body_bytes": 2935, "body_sha256": "sha256:883669555e22c8f8bf62feadfa855bb6f2f7a6fc4029d9d0ade50972e633bdbe", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:all_load_balancer_domains", "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:custom_domain_list", "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:disabled"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination", "path": "documentation/resources/route/properties/routes/route_destination/csrf_policy/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1101321211132233-2202003002310130-3123211010001032-2023310312100303-3310321001200211-2001220013330302-2200133210032112-2213001112022232", "registry_path": "docs/guides/resources--route--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,custom_domain_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:all_load_balancer_domains", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:all_load_balancer_domains", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,custom_domain_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:custom_domain_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.csrf_policy:ConflictingObjectAttributes:custom_domain_list,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:custom_domain_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.csrf_policy:ConflictingObjectAttributes:custom_domain_list,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:disabled", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "csrf_policy"], "schema_version": 1, "sections": [{"aliases": ["routes route destination csrf policy all load balancer domains"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:all_load_balancer_domains", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "csrf_policy", "all_load_balancer_domains"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route destination csrf policy custom domain list"], "anchor": "section", "description": "List of domain names used for Host header matching.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:custom_domain_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--route_destination--csrf_policy--custom_domain_list--domains", "enforcement": "provider-schema", "group": "routes.route_destination.csrf_policy.custom_domain_list:RequiredObjectAttributes:domains", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:custom_domain_list", "type": "requires"}], "schema_path": ["routes", "route_destination", "csrf_policy", "custom_domain_list"], "syntax": "block", "type": "object"}, {"aliases": ["routes route destination csrf policy disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:disabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "csrf_policy", "disabled"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/csrf_policy/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its destination.the policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["routeCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.csrf_policy

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- routes.route_destination.csrf_policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host.

&#8203;1. The origin that caused the user agent to issue the request (source origin). &#8203;2. The
origin that the request is going to (target origin). When the policy evaluating a request, it
ensures both pieces of information are present and compare their values. If the source origin is
missing or origins do not match the request is rejected. The exception to this being if the
source-origin has been added to they policy as valid. Because CSRF attacks specifically target
state-changing requests, the policy only acts on the HTTP requests that have state-changing method
(PUT,POST, etc.).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_load_balancer_domains",
    "custom_domain_list"),
  validators.ConflictingObjectAttributes("all_load_balancer_domains",
    "disabled"),
  validators.ConflictingObjectAttributes("custom_domain_list",
    "disabled")}
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
  "x-ves-oneof-field-allowed_domains": "[\"all_load_balancer_domains\",\"custom_domain_list\",\"disabled\"]"
}
```

Terraform syntax:

```terraform
csrf_policy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [all_load_balancer_domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/csrf_policy/all_load_balancer_domains/): complete subsection reference.

- [custom_domain_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/csrf_policy/custom_domain_list/): complete subsection reference.

- [disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/csrf_policy/disabled/): complete subsection reference.
