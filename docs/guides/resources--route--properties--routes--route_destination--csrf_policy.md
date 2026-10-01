---
page_title: "routes.route_destination.csrf_policy"
subcategory: ""
description: "routes.route_destination.csrf_policy for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 3473, "body_sha256": "sha256:a99c6542017918b4d6357ef60b5813f1bb64c9ba0dcf905e4ca2cd51acefc725", "canonical_id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy", "child_ids": ["xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:all_load_balancer_domains", "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:custom_domain_list", "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:disabled"], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination", "path": "docs/guides/resources--route--properties--routes--route_destination--csrf_policy.md", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "route_destination", "csrf_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/csrf_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.route_destination.csrf_policy for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.csrf_policy

Breadcrumbs:

- [xcsh_route](../resources/route.md)
- [Property reference](resources--route--reference.md)
- [routes](resources--route--properties--routes.md)
- [routes.route_destination](resources--route--properties--routes--route_destination.md)
- routes.route_destination.csrf_policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host. 1. The origin that caused the user
agent..

Upstream description:

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

- [all_load_balancer_domains](resources--route--properties--routes--route_destination--csrf_policy--all_load_balancer_domains.md): complete subsection reference.

- [custom_domain_list](resources--route--properties--routes--route_destination--csrf_policy--custom_domain_list.md): complete subsection reference.

- [disabled](resources--route--properties--routes--route_destination--csrf_policy--disabled.md): complete subsection reference.

## Next pages

- [routes.route_destination.csrf_policy.all_load_balancer_domains](resources--route--properties--routes--route_destination--csrf_policy--all_load_balancer_domains.md)
- [routes.route_destination.csrf_policy.custom_domain_list](resources--route--properties--routes--route_destination--csrf_policy--custom_domain_list.md)
- [routes.route_destination.csrf_policy.disabled](resources--route--properties--routes--route_destination--csrf_policy--disabled.md)
- [routes.route_destination](resources--route--properties--routes--route_destination.md)
- [xcsh_route](../resources/route.md)
