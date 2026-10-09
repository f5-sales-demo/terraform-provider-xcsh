---
page_title: "routes.route_redirect.remove_all_params"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["routes route redirect remove all params"], "body_bytes": 1132, "body_sha256": "sha256:45cd3beb7291cb3e9bf826255a99471b8ac95ae77a5e766b3b0abe9289ad822a", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_redirect:remove_all_params", "parent_id": "xcsh-docs:resources:route:properties:routes:route_redirect", "path": "documentation/resources/route/properties/routes/route_redirect/remove_all_params/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2112000102133131-3330233010323032-2330110232311113-1222013003331011-2113100211123113-0122010001101103-3012000133322130-0130312202323132", "registry_path": "docs/guides/resources--route--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_redirect", "remove_all_params"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_redirect/remove_all_params/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["routeCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_redirect.remove_all_params

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [routes.route_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_redirect/)
- routes.route_redirect.remove_all_params

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for remove all params.

Additional upstream details:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
remove_all_params = {}
```

This is an empty object or choice marker. It has no direct properties.
