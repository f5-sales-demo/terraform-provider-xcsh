---
page_title: "routes.route_destination.hash_policy.cookie.samesite_strict"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["routes route destination hash policy cookie samesite strict"], "body_bytes": 1488, "body_sha256": "sha256:fa8bcbe100fd726cc7a3810b3dd37e931eeca615bb59019e1ec418a17ef84bdc", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:samesite_strict", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie", "path": "documentation/resources/route/properties/routes/route_destination/hash_policy/cookie/samesite_strict/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-2020131323033213-0301333210023221-1002031200003031-2001230130111230-2233132233221110-0213230023200021-1320332023011013-0312100101121001", "registry_path": "docs/guides/resources--route--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "hash_policy", "cookie", "samesite_strict"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/hash_policy/cookie/samesite_strict/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["routeCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.hash_policy.cookie.samesite_strict

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- [routes.route_destination.hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/)
- [routes.route_destination.hash_policy.cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/cookie/)
- routes.route_destination.hash_policy.cookie.samesite_strict

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
samesite_strict = {}
```

This is an empty object or choice marker. It has no direct properties.
