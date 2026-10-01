---
page_title: "routes.simple_route.advanced_options.csrf_policy"
subcategory: "Load Balancing"
description: "routes.simple_route.advanced_options.csrf_policy for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3920, "body_sha256": "sha256:5b2080acf792dbe00d645b96f037da685ca6e92f0639fd1ea8bd0233d022ff4f", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy:all_load_balancer_domains", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy:custom_domain_list", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy:disabled"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options", "path": "docs/guides/resources--http_loadbalancer--properties--routes--simple_route--advanced_options--csrf_policy.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "csrf_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/csrf_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.simple_route.advanced_options.csrf_policy for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options.csrf_policy

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [routes](resources--http_loadbalancer--properties--routes.md)
- [routes.simple_route](resources--http_loadbalancer--properties--routes--simple_route.md)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--properties--routes--simple_route--advanced_options.md)
- routes.simple_route.advanced_options.csrf_policy

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

- [all_load_balancer_domains](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--csrf_policy--all_load_balancer_domains.md): complete subsection reference.

- [custom_domain_list](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--csrf_policy--custom_domain_list.md): complete subsection reference.

- [disabled](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--csrf_policy--disabled.md): complete subsection reference.

## Next pages

- [routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--csrf_policy--all_load_balancer_domains.md)
- [routes.simple_route.advanced_options.csrf_policy.custom_domain_list](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--csrf_policy--custom_domain_list.md)
- [routes.simple_route.advanced_options.csrf_policy.disabled](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--csrf_policy--disabled.md)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--properties--routes--simple_route--advanced_options.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
