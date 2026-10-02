---
page_title: "csrf_policy"
subcategory: "Load Balancing"
description: "To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its destination.the policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request"
xcsh_docs: {"aliases": ["csrf policy"], "body_bytes": 3651, "body_sha256": "sha256:88cd1a48476efe34f66206777bf2e0f54da0597ca8681f00f6204979edcd935b", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:csrf_policy:all_load_balancer_domains", "xcsh-docs:resources:cdn_loadbalancer:properties:csrf_policy:custom_domain_list", "xcsh-docs:resources:cdn_loadbalancer:properties:csrf_policy:disabled"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:csrf_policy", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/csrf_policy/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1000130300033331-3002211233300130-2313130100130200-1302332133032013-0013323333213201-3223012120111001-2100331133200101-2301200003031301", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-009.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,custom_domain_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:csrf_policy:all_load_balancer_domains", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:csrf_policy:all_load_balancer_domains", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,custom_domain_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:csrf_policy:custom_domain_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "csrf_policy:ConflictingObjectAttributes:custom_domain_list,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:csrf_policy:custom_domain_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:csrf_policy:disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "csrf_policy:ConflictingObjectAttributes:custom_domain_list,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:csrf_policy:disabled", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["csrf_policy"], "schema_version": 1, "sections": [{"aliases": ["all load balancer domains"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:csrf_policy:all_load_balancer_domains", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["csrf_policy", "all_load_balancer_domains"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom domain list"], "anchor": "section", "description": "List of domain names used for Host header matching.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:csrf_policy:custom_domain_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-csrf_policy--custom_domain_list--domains", "enforcement": "provider-schema", "group": "csrf_policy.custom_domain_list:RequiredObjectAttributes:domains", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:csrf_policy:custom_domain_list", "type": "requires"}], "schema_path": ["csrf_policy", "custom_domain_list"], "syntax": "block", "type": "object"}, {"aliases": ["disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:csrf_policy:disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["csrf_policy", "disabled"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/csrf_policy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its destination.the policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# csrf_policy

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- csrf_policy

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

- [all_load_balancer_domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/csrf_policy/all_load_balancer_domains/): complete subsection reference.

- [custom_domain_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/csrf_policy/custom_domain_list/): complete subsection reference.

- [disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/csrf_policy/disabled/): complete subsection reference.

## Next pages

- [csrf_policy.all_load_balancer_domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/csrf_policy/all_load_balancer_domains/)
- [csrf_policy.custom_domain_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/csrf_policy/custom_domain_list/)
- [csrf_policy.disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/csrf_policy/disabled/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
