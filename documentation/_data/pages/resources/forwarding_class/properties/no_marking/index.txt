---
page_title: "no_marking"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["no marking"], "body_bytes": 1097, "body_sha256": "sha256:dbed467cd51ee03b9c1307a7999617bdd4a356e376cdbcb72ac3271766d2208e", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:forwarding_class:collection", "completeness": "complete", "id": "xcsh-docs:resources:forwarding_class:properties:no_marking", "parent_id": "xcsh-docs:resources:forwarding_class:reference", "path": "documentation/resources/forwarding_class/properties/no_marking/index.md", "product": "distributed-cloud", "provider_name": "forwarding_class", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1021103130221111-1123033322313003-2032001203200330-0001333002012023-3323021332212211-3113320333001230-1101113133200332-1310201221313033", "registry_path": "docs/guides/resources--forwarding_class--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["no_marking"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forwarding_class/properties/no_marking/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forwarding_classCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# no_marking

Breadcrumbs:

- [xcsh_forwarding_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forwarding_class/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forwarding_class/properties/)
- no_marking

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
no_marking = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forwarding_class/properties/)
- [xcsh_forwarding_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forwarding_class/)
