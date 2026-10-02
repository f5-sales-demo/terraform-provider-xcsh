---
page_title: "enable"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["enable"], "body_bytes": 1031, "body_sha256": "sha256:82621bde45bcfac26e574c88539363c54d8f08ad0ce5bc50b5a615363a1cb9e0", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:segment:collection", "completeness": "complete", "id": "xcsh-docs:resources:segment:properties:enable", "parent_id": "xcsh-docs:resources:segment:reference", "path": "documentation/resources/segment/properties/enable/index.md", "product": "distributed-cloud", "provider_name": "segment", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-3100223213111010-2103312210220322-2233203002011122-3223023302011331-2333321113330202-2322330231112103-1333211031112123-0331313121113330", "registry_path": "docs/guides/resources--segment--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/segment/properties/enable/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["segmentCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable

Breadcrumbs:

- [xcsh_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/properties/)
- enable

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
enable = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/properties/)
- [xcsh_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/)
