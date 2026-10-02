---
page_title: "kubernetes_upgrade_drain.disable_upgrade_drain"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["kubernetes upgrade drain disable upgrade drain"], "body_bytes": 1316, "body_sha256": "sha256:76314194da744a54e55ee32fcac7363b6cdf3008ca5b32f073bcf0ebed9072e4", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "parent_id": "xcsh-docs:resources:fleet:properties:kubernetes_upgrade_drain", "path": "documentation/resources/fleet/properties/kubernetes_upgrade_drain/disable_upgrade_drain/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3220102311201123-3330002212131021-0130223202031202-1031023220131321-2103003012121133-2103201033221311-3023311122332132-3033021132202010", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["kubernetes_upgrade_drain", "disable_upgrade_drain"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/kubernetes_upgrade_drain/disable_upgrade_drain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kubernetes_upgrade_drain.disable_upgrade_drain

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/kubernetes_upgrade_drain/)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable upgrade drain.

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
disable_upgrade_drain = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/kubernetes_upgrade_drain/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
