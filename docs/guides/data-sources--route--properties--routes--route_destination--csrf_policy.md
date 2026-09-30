---
page_title: "routes.route_destination.csrf_policy"
subcategory: ""
description: "routes.route_destination.csrf_policy for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 2930, "body_sha256": "sha256:09002c4c2fa31c983154a6e557cfa56a1e84e4fc65f14f27cbbfd3a119780219", "canonical_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:csrf_policy", "child_ids": ["xcsh-docs:data-sources:route:properties:routes:route_destination:csrf_policy:all_load_balancer_domains", "xcsh-docs:data-sources:route:properties:routes:route_destination:csrf_policy:custom_domain_list", "xcsh-docs:data-sources:route:properties:routes:route_destination:csrf_policy:disabled"], "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:route_destination:csrf_policy", "parent_id": "xcsh-docs:data-sources:route:properties:routes:route_destination", "path": "docs/guides/data-sources--route--properties--routes--route_destination--csrf_policy.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "route_destination", "csrf_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/route_destination/csrf_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.route_destination.csrf_policy for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.route_destination.csrf_policy

Breadcrumbs:

- [xcsh_route](../data-sources/route.md)
- [Property reference](data-sources--route--reference.md)
- [routes](data-sources--route--properties--routes.md)
- [routes.route_destination](data-sources--route--properties--routes--route_destination.md)
- routes.route_destination.csrf_policy

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [all_load_balancer_domains](data-sources--route--properties--routes--route_destination--csrf_policy--all_load_balancer_domains.md): complete subsection reference.

- [custom_domain_list](data-sources--route--properties--routes--route_destination--csrf_policy--custom_domain_list.md): complete subsection reference.

- [disabled](data-sources--route--properties--routes--route_destination--csrf_policy--disabled.md): complete subsection reference.

## Next pages

- [routes.route_destination.csrf_policy.all_load_balancer_domains](data-sources--route--properties--routes--route_destination--csrf_policy--all_load_balancer_domains.md)
- [routes.route_destination.csrf_policy.custom_domain_list](data-sources--route--properties--routes--route_destination--csrf_policy--custom_domain_list.md)
- [routes.route_destination.csrf_policy.disabled](data-sources--route--properties--routes--route_destination--csrf_policy--disabled.md)
- [routes.route_destination](data-sources--route--properties--routes--route_destination.md)
- [xcsh_route](../data-sources/route.md)
