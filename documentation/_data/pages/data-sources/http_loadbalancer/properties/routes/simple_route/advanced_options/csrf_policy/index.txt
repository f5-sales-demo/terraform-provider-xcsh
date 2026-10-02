---
page_title: "routes.simple_route.advanced_options.csrf_policy"
subcategory: "Load Balancing"
description: "To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its destination.the policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request"
xcsh_docs: {"aliases": ["backend servers", "origin servers", "routes simple route advanced options csrf policy", "upstream servers"], "body_bytes": 4115, "body_sha256": "sha256:11b7eb860f06dd575e8b455f289979a7ab1e5ec32824fd5c7b982e3a8b4596a0", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy:all_load_balancer_domains", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy:custom_domain_list", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy:disabled"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options", "path": "documentation/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/csrf_policy/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1200313331233133-2123332112011333-0100120030021321-3021202231111013-3220303021030123-3322301130000330-2121223302332012-3211201202012023", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-024.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "csrf_policy"], "schema_version": 1, "sections": [{"aliases": ["all load balancer domains"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy:all_load_balancer_domains", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "csrf_policy", "all_load_balancer_domains"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom domain list"], "anchor": "section", "description": "List of domain names used for Host header matching.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy:custom_domain_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "csrf_policy", "custom_domain_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy:disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "csrf_policy", "disabled"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/csrf_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its destination.the policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options.csrf_policy

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/)
- [routes.simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/)
- [routes.simple_route.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/)
- routes.simple_route.advanced_options.csrf_policy

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

- [all_load_balancer_domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/csrf_policy/all_load_balancer_domains/): complete subsection reference.

- [custom_domain_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/csrf_policy/custom_domain_list/): complete subsection reference.

- [disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/csrf_policy/disabled/): complete subsection reference.

## Next pages

- [routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/csrf_policy/all_load_balancer_domains/)
- [routes.simple_route.advanced_options.csrf_policy.custom_domain_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/csrf_policy/custom_domain_list/)
- [routes.simple_route.advanced_options.csrf_policy.disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/csrf_policy/disabled/)
- [routes.simple_route.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
