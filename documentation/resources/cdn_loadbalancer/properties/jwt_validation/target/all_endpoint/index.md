---
page_title: "jwt_validation.target.all_endpoint"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["jwt validation target all endpoint"], "body_bytes": 1446, "body_sha256": "sha256:35b9ffbb369dfb1445cbe5bc5ca6bfcec1f386842d20b7cb09d2db0e76e30687", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:target:all_endpoint", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:target", "path": "documentation/resources/cdn_loadbalancer/properties/jwt_validation/target/all_endpoint/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0030331312230121-3211212230231021-2323301133300032-3223131222122023-2130213301221001-0210031113111310-2300111333003032-3300212301030330", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation", "target", "all_endpoint"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/jwt_validation/target/all_endpoint/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.target.all_endpoint

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/)
- [jwt_validation.target](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/target/)
- jwt_validation.target.all_endpoint

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
all_endpoint = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [jwt_validation.target](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/target/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
